/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license output.pushrmation.
 *--------------------------------------------------------------------------------------------*/

import * as nunjucks from 'nunjucks'
import * as path from 'path'
import {
    ArraySchema,
    ChoiceSchema,
    CodeModel,
    DictionarySchema,
    ImplementationLocation,
    Metadata,
    Operation,
    OperationGroup,
    Parameter,
    Schema,
    SchemaResponse,
    SchemaType,
    SealedChoiceSchema,
    codeModelSchema
} from '@azure-tools/codemodel'
import {
    ExampleExtension,
    ExampleExtensionResponse,
    ExampleModel,
    ExampleParameter,
    ExampleResponse,
    ExampleValue,
    ExtensionName,
    TestCodeModel,
    TestGroup,
    TestModel,
    TestScenario
} from '../model/testModel'
import { Helper } from '../util/helper'
import { Host, startSession } from '@azure-tools/autorest-extension-base'
import { ImportManager } from '@autorest/go/dist/generator/imports'
import {
    formatParameterTypeName,
    getMethodParameters,
    sortParametersByRequired
} from '@autorest/go/dist/generator/helpers'
import { formatWithOptions } from 'util'
import {
    getResponse,
    isLROOperation,
    isPageableOperation,
    isSchemaResponse
} from '@autorest/go/dist/common/helpers'
import { values } from '@azure-tools/linq'

export async function processRequest(host: Host): Promise<void> {
    const session = await startSession<TestCodeModel>(host, {}, codeModelSchema)
    const config = await session.getValue('')
    const generator = await new TestGenerator(host, session.model, config)
    await generator.GenerateMockTest()
    await Helper.outputToModelerfour(host, session)
    await Helper.dumpCodeModel(host, session, 'go-tester.yaml')
}

class GoTestData extends TestGroup {
    packageName: string
    imports: string
}

class GoExampleModel extends ExampleModel {
    opName: string
    methodParametersOutput: string
    clientParametersOutput: string
    returnInfo: string[]
}

// homo structureed with getAPIParametersSig() in autorest.go
function getAPIParametersSig(op: any, imports: ImportManager): Array<[string, string]> {
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
function getClientParametersSig(
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
function generateReturnsInfo(op: Operation, apiType: 'api' | 'op' | 'handler'): string[] {
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

class TestGenerator {
    public constructor(
        public host: Host,
        public codeModel: TestCodeModel,
        public config: Record<string, any>
    ) {}
    importManager: ImportManager

    public GetMockRenderData(testGroup: GoTestData) {
        this.importManager = new ImportManager()
        testGroup.packageName = this.codeModel.language.go.packageName
        for (const scenario of testGroup.scenarios) {
            for (const example of scenario.examples as GoExampleModel[]) {
                const op = example.operation
                example.opName = op.language.go.name
                if (isLROOperation(op as any)) {
                    example.opName = op.language.go.protocolNaming.internalMethod
                }
                example.methodParametersOutput = getAPIParametersSig(op, this.importManager)
                    .map(([paramName, typeName]) => {
                        if (paramName === 'ctx') {
                            return 'ctx'
                        }
                        for (const methodParameter of example.methodParameters) {
                            if (this.GetLanguageName(methodParameter.parameter) == paramName) {
                                return this.ExampleValueToString(
                                    methodParameter.exampleValue,
                                    !methodParameter.parameter.required
                                )
                            }
                        }
                        return 'nil'
                    })
                    .join(',\n')
                example.clientParametersOutput = getClientParametersSig(
                    example.operationGroup,
                    this.importManager
                )
                    .map(([paramName, typeName]) => {
                        if (paramName === 'con') {
                            return 'con'
                        }
                        for (const clientParameter of example.clientParameters) {
                            if (this.GetLanguageName(clientParameter.parameter) == paramName) {
                                return this.ExampleValueToString(
                                    clientParameter.exampleValue,
                                    !clientParameter.parameter.required
                                )
                            }
                        }
                        return 'nil'
                    })
                    .join(',\n')
                example.returnInfo = generateReturnsInfo(op, 'op')
            }
        }
        testGroup.imports = this.importManager.text()
    }

    private GetLanguageName(meta: any): string {
        return (meta as Metadata).language.go.name
    }

    public ExampleValueToString(exampleValue: ExampleValue, isPtr: boolean | undefined): string {
        const ptr = exampleValue.language?.go?.byValue || isPtr === false ? '' : '&'
        if (exampleValue.schema?.type === SchemaType.Array) {
            const elementPtr = exampleValue.schema.language.go.elementIsPtr ? '*' : ''
            return (
                `${ptr}[]${elementPtr}${this.GetLanguageName(
                    (exampleValue.schema as ArraySchema).elementType
                )}{\n` +
                (exampleValue.value as any[])
                    .map((x) =>
                        this.ExampleValueToString(x, exampleValue.schema.language.go.elementIsPtr)
                    )
                    .join(',\n') +
                '}'
            )
        } else if (exampleValue.schema?.type === SchemaType.Object) {
            let output = ''
            output += `${ptr}${this.GetLanguageName(exampleValue.schema)}{\n`
            for (const [_, parentValue] of Object.entries(exampleValue.parentsValue)) {
                output += `${this.GetLanguageName(parentValue)}: ${this.ExampleValueToString(
                    parentValue,
                    false
                )},\n`
            }
            for (const [_, value] of Object.entries(exampleValue.value)) {
                output += `${this.GetLanguageName(value)}: ${this.ExampleValueToString(
                    value as ExampleValue,
                    undefined
                )},\n`
            }
            output += '}'
            return output
        } else if (exampleValue.schema?.type === SchemaType.Dictionary) {
            let output = `${ptr}map[string]${
                exampleValue.schema.language.go.elementIsPtr ? '*' : ''
            }${(exampleValue.schema as DictionarySchema).elementType.language.go.name}{\n`
            for (const [key, value] of Object.entries(exampleValue.value)) {
                output += `"${key}": ${this.ExampleValueToString(
                    value as ExampleValue,
                    exampleValue.schema.language.go.elementIsPtr
                )},\n`
            }
            output += '}'
            return output
        }
        return this.RawValueToString(
            exampleValue.value,
            exampleValue.schema,
            isPtr === undefined ? !exampleValue.language.go.byValue : isPtr
        )
    }

    public RawValueToString(rawValue: any, schema: Schema, isPtr: boolean): string {
        let ret = JSON.stringify(rawValue)
        if (Object.getPrototypeOf(rawValue) === Object.prototype) ret = '`' + ret + '`'
        const goType = this.GetLanguageName(schema)
        if ([SchemaType.Choice, SchemaType.SealedChoice].indexOf(schema.type) >= 0) {
            const choiceValue = Helper.findChoiceValue(schema as ChoiceSchema, rawValue)
            ret = this.GetLanguageName(choiceValue)
        }
        if (schema.type === SchemaType.Constant || goType === 'string') {
            ret = `"${Helper.escapeString(rawValue)}"`
        } else if (goType === 'time.Time') {
            this.importManager.add('time')
            ret = `func() time.Time { t, _ := time.Parse(time.RFC3339, "${rawValue}"); return t}()`
        } else if (goType === 'map[string]interface{}') {
            ret = `map[string]interface{}{\n`
            for (const [key, value] of Object.entries(rawValue)) {
                ret += `"${key}": \`${JSON.stringify(value)}\`,\n`
            }
            ret += '}'
        }

        if (isPtr) {
            const PTR_CONVERTS = {
                string: 'StringPtr',
                bool: 'BoolPtr',
                'time.Time': 'TimePtr',
                int32: 'Int32Ptr',
                int64: 'Int64Ptr',
                float32: 'Float32Ptr',
                float64: 'Float64Ptr'
            }

            if (schema.type === SchemaType.Constant) {
                ret = `to.StringPtr(${ret})`
                this.importManager.add('github.com/Azure/azure-sdk-for-go/sdk/to')
            } else if ([SchemaType.Choice, SchemaType.SealedChoice].indexOf(schema.type) >= 0) {
                ret += '.ToPtr()'
            } else if (Object.prototype.hasOwnProperty.call(PTR_CONVERTS, goType)) {
                ret = `to.${PTR_CONVERTS[goType]}(${ret})`
                this.importManager.add('github.com/Azure/azure-sdk-for-go/sdk/to')
            } else {
                ret = '&' + ret
            }
        }

        return ret
    }

    public async GenerateMockTest() {
        for (const testGroup of this.codeModel.testModel.mockTests) {
            //Prepare for render data as GoTestData
            this.GetMockRenderData(testGroup as GoTestData)

            // Render to template
            const tmplPath = path.relative(
                process.cwd(),
                path.join(`${__dirname}`, '../../src/template/mockTest.go.njk')
            )
            nunjucks.configure({ autoescape: false })
            const output = nunjucks.render(tmplPath, testGroup)
            this.host.WriteFile(`${testGroup.name}_test.go`, output, undefined)
        }
    }
}
