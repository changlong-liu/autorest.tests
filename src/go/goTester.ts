/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license output.pushrmation.
 *--------------------------------------------------------------------------------------------*/

import {
    ArraySchema,
    ChoiceSchema,
    DictionarySchema,
    Metadata,
    Schema,
    SchemaType,
    codeModelSchema
} from '@azure-tools/codemodel'
import { ExampleModel, ExampleValue, TestCodeModel, TestGroup } from '../core/model'
import { Helper } from '../util/helper'
import { Host, startSession } from '@azure-tools/autorest-extension-base'
import { ImportManager } from '@autorest/go/dist/generator/imports'
import { TestGenerator } from '../core/testGenerator'
import { generateReturnsInfo, getAPIParametersSig, getClientParametersSig } from './codegenBridge'
import { isLROOperation } from '@autorest/go/dist/common/helpers'

export async function processRequest(host: Host): Promise<void> {
    const session = await startSession<TestCodeModel>(host, {}, codeModelSchema)
    const config = await session.getValue('')
    const generator = await new GoTestGenerator(host, session.model, config)
    await generator.GenerateMockTest('mockTest.go.njk')
    await Helper.outputToModelerfour(host, session)
    if (session.getValue('debug', false)) {
        await Helper.dumpCodeModel(host, session, 'go-tester.yaml')
    }
}

class GoTestGroup extends TestGroup {
    packageName: string
    imports: string
}

class GoExampleModel extends ExampleModel {
    opName: string
    isLRO: boolean
    methodParametersOutput: string
    clientParametersOutput: string
    returnInfo: string[]
}
class GoTestGenerator extends TestGenerator {
    importManager: ImportManager

    public GetMockTestFilename(testGroup: TestGroup) {
        return `${testGroup.name}_test.go`
    }
    public GetLanguageName(meta: any): string {
        return (meta as Metadata).language.go.name
    }
    public GenMockRenderData(testGroup: GoTestGroup) {
        this.importManager = new ImportManager()
        testGroup.packageName = this.codeModel.language.go.packageName
        for (const scenario of testGroup.scenarios) {
            for (const example of scenario.examples as GoExampleModel[]) {
                const op = example.operation
                example.opName = op.language.go.name
                if (isLROOperation(op as any)) {
                    // example.opName = op.language.go.protocolNaming.internalMethod
                    example.opName = 'Begin' + example.opName
                    example.isLRO = true
                } else {
                    example.isLRO = false
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
}
