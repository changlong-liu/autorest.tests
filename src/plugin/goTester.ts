/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license output.pushrmation.
 *--------------------------------------------------------------------------------------------*/

import * as nunjucks from 'nunjucks'
import * as path from 'path'
import {
    CodeModel,
    ImplementationLocation,
    Metadata,
    Operation,
    SchemaResponse,
    SchemaType,
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
import { formatParameterTypeName, getMethodParameters } from '@autorest/go/dist/generator/helpers'
import { isLROOperation, isPageableOperation } from '@autorest/go/dist/common/helpers'
import { values } from '@azure-tools/linq'

export async function processRequest(host: Host): Promise<void> {
    const session = await startSession<TestCodeModel>(host, {}, codeModelSchema)
    const config = await session.getValue('')
    const generator = await new TestGenerator(host, session.model, config)
    await generator.GenerateMockTest()
    await Helper.outputToModelerfour(host, session)
    // await Helper.dumpCodeModel(host, session, 'test-modeler.yaml')
}

class GoTestData extends TestGroup {
    packageName: string
}

class GoExampleModel extends ExampleModel {
    opName: string
    parametersOutput: string[]
}

// homo structureed with getAPIParametersSig() in autorest.go
function getAPIParametersSig(op: any, imports: ImportManager): Array<string> {
    const methodParams = getMethodParameters(op)
    const params = new Array<string>()
    if (!isPageableOperation(op) || isLROOperation(op)) {
        imports.add('context')
        params.push('ctx')
    }
    for (const methodParam of values(methodParams)) {
        params.push(methodParam.language.go.name)
    }
    return params
}

class TestGenerator {
    public constructor(
        public host: Host,
        public codeModel: TestCodeModel,
        public config: Record<string, any>
    ) {}

    public GetMockRenderData(testGroup: GoTestData) {
        testGroup.packageName = this.codeModel.language.go.packageName
        for (const scenario of testGroup.scenarios) {
            for (const example of scenario.examples as GoExampleModel[]) {
                const op = example.operation
                example.opName = op.language.go.name
                if (isLROOperation(op as any)) {
                    example.opName = op.language.go.protocolNaming.internalMethod
                }
                example.parametersOutput = getAPIParametersSig(op, new ImportManager()).map(
                    (paramName) => {
                        for (const methodParameter of example.methodParameters) {
                            if (this.GetLanguageName(methodParameter.parameter) == paramName) {
                                return this.ExampleValueToString(methodParameter.exampleValue)
                            }
                        }
                    }
                )
            }
        }
    }

    private GetLanguageName(meta: Metadata): string {
        return meta.language.go.name
    }

    public ExampleValueToString(exampleValue: ExampleValue): string {
        if (exampleValue.schema?.type === SchemaType.Array) {
            return (
                '[' +
                (exampleValue.value as any[]).map((x) => this.ExampleValueToString(x)).join(', ') +
                ']'
            )
        } else if (exampleValue.schema?.type === SchemaType.Object) {
        }
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
