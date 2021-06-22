/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license output.pushrmation.
 *--------------------------------------------------------------------------------------------*/

import {
    CodeModel,
    ImplementationLocation,
    SchemaResponse,
    codeModelSchema
} from '@azure-tools/codemodel'
import {
    ExampleExtension,
    ExampleExtensionResponse,
    ExampleModel,
    ExampleParameter,
    ExampleResponse,
    ExtensionName,
    TestGroup,
    TestModel,
    TestScenario
} from '../model/testModel'
import { Helper } from '../util/helper'
import { Host, startSession } from '@azure-tools/autorest-extension-base'

export async function processRequest(host: Host): Promise<void> {
    const session = await startSession<CodeModel>(host, {}, codeModelSchema)
    const files = await session.listInputs()
    const codemodel = await session.readFile('code-model-v4.yaml')
    const model = session.model
    await genExampleModels(session.model)
    const testModel = new TestModel()
    await genMockTests(session.model, testModel)
    session.model['testModel'] = testModel
    await Helper.outputToModelerfour(host, session)
    await Helper.dumpCodeModel(host, session, 'test-modeler.yaml')
}

async function genExampleModels(codeModel: CodeModel) {
    codeModel.operationGroups.forEach((operationGroup) => {
        operationGroup.operations.forEach((operation) => {
            for (const [exampleName, rawValue] of Object.entries(
                operation.extensions?.[ExtensionName.xMsExamples] ?? {}
            )) {
                const exampleExtension = rawValue as ExampleExtension
                const parametersInExample = exampleExtension.parameters
                const exampleModel = new ExampleModel(exampleName, operation)
                for (const parameter of Helper.allParameters(operation)) {
                    const dotPath = Helper.getFlattenedNames(parameter).join('.')
                    if (Helper.isPathDefined(parametersInExample, dotPath)) {
                        const exampleParameter = new ExampleParameter(
                            parameter,
                            parametersInExample[dotPath]
                        )
                        if (parameter.implementation == ImplementationLocation.Method) {
                            exampleModel.methodParameters.push(exampleParameter)
                        } else if (parameter.implementation == ImplementationLocation.Client) {
                            exampleModel.clientParameters.push(exampleParameter)
                        } else {
                            //
                        }
                    }
                }

                for (const [statusCode, response] of Object.entries(exampleExtension.responses)) {
                    const exampleExtensionResponse = response as ExampleExtensionResponse
                    const schemaResponse = operation.responses[0] as SchemaResponse
                    exampleModel.responses[statusCode] = ExampleResponse.createInstance(
                        exampleExtensionResponse,
                        schemaResponse.schema,
                        schemaResponse.language
                    )
                }

                operation.extensions[ExtensionName.xMsExampleModels] =
                    operation.extensions[ExtensionName.xMsExampleModels] || {}
                operation.extensions[ExtensionName.xMsExampleModels][exampleName] = exampleModel
            }
        })
    })
}

async function genMockTests(codeModel: CodeModel, testModel: TestModel) {
    const testGroup = new TestGroup('mock')
    codeModel.operationGroups.forEach((operationGroup) => {
        operationGroup.operations.forEach((operation) => {
            for (const [exampleName, exampleModel] of Object.entries(
                operation.extensions?.[ExtensionName.xMsExampleModels] ?? {}
            )) {
                const testScenario = new TestScenario(exampleName)
                testScenario.examples.push(exampleModel as ExampleModel)
                testGroup.scenarios.push(testScenario)
            }
        })
    })
    testModel.mockTests.push(testGroup)
}
