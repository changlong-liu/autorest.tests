import { ImportManager } from '@autorest/go/dist/generator/imports'
import { Operation, OperationGroup, Parameter, SchemaResponse } from '@azure-tools/codemodel'
import { formatParameterTypeName, getMethodParameters } from '@autorest/go/dist/generator/helpers'
import {
    getResponse,
    isLROOperation,
    isPageableOperation,
    isSchemaResponse
} from '@autorest/go/dist/common/helpers'
import { values } from '@azure-tools/linq'

// homo structureed with getAPIParametersSig() in autorest.go
export function getAPIParametersSig(op: any, imports: ImportManager): Array<[string, string]> {
    const methodParams = getMethodParameters(op)
    const params = new Array<[string, string]>()
    if (!isPageableOperation(op) || isLROOperation(op)) {
        // imports.add('context')       // comment out since has been envolved in .njk
        params.push(['ctx', 'context.Context'])
    }
    for (const methodParam of values(methodParams)) {
        params.push([methodParam.language.go.name, formatParameterTypeName(methodParam)])
    }
    return params
}

// homo structured with logic in generateOperations() in autorest.go
export function getClientParametersSig(
    group: OperationGroup,
    imports: ImportManager
): Array<[string, string]> {
    const params = []
    // imports.add('armcore')       // comment out since has been envolved in .njk
    params.push(['con', '*armcore.Connection'])

    for (const parameter of values((group.language.go?.clientParams || []) as Parameter[])) {
        params.push([parameter.language.go.name, formatParameterTypeName(parameter as any)])
    }
    return params
}

// copied from autorest.go
function isMultiRespOperation(op: Operation): boolean {
    // treat LROs as single-response ops
    if (!op.responses || op.responses.length === 1 || isLROOperation(op as any)) {
        return false
    }
    // count the number of distinct schemas returned by this operation
    const schemaResponses = new Array<SchemaResponse>()
    for (const response of values(op.responses)) {
        // perform the comparison by name as some responses have different objects for the same underlying response type
        if (
            isSchemaResponse(response as any) &&
            !values(schemaResponses)
                .where(
                    (sr) =>
                        sr.schema.language.go?.name ===
                        (response as SchemaResponse).schema.language.go?.name
                )
                .any()
        ) {
            schemaResponses.push(response as any)
        }
    }
    return schemaResponses.length > 1
}

// homo structured with generateReturnsInfo() in autorest.go
export function generateReturnsInfo(op: Operation, apiType: 'api' | 'op' | 'handler'): string[] {
    if (!op.responses) {
        return ['*http.Response', 'error']
    }
    if (isMultiRespOperation(op)) {
        return ['interface{}', 'error']
    }
    const schemaResponse = getResponse(op as any)
    let returnType = '*http.Response'
    if (isLROOperation(op as any)) {
        switch (apiType) {
            case 'handler':
                // we only have a handler for operations that return a schema
                returnType = schemaResponse!.schema.language.go!.responseType.name
                break
            case 'api':
                returnType = 'HTTPPollerResponse'
                if (schemaResponse !== undefined) {
                    returnType = schemaResponse.schema.language.go!.lroResponseType.language.go!
                        .name
                }
                break
            case 'op':
                returnType = '*azcore.Response'
                break
        }
    } else if (isPageableOperation(op as any)) {
        switch (apiType) {
            case 'handler':
                // pageable operations always return a schema
                returnType = schemaResponse!.schema.language.go!.responseType.name
                break
            case 'api':
            case 'op':
                // pager operations don't return an error
                return [op.language.go!.pageableType.name]
        }
    } else if (schemaResponse !== undefined) {
        // simple schema response
        returnType = schemaResponse.schema.language.go!.responseType.name
    } else if (op.language.go!.headAsBoolean === true) {
        // NOTE: this case must come after the hasSchemaResponse() check to properly handle
        //       the intersection of head-as-boolean with modeled response headers
        return ['BooleanResponse', 'error']
    }
    return [returnType, 'error']
}
