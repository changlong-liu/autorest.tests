/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license output.pushrmation.
 *--------------------------------------------------------------------------------------------*/

import { CodeModel, ImplementationLocation, codeModelSchema } from '@azure-tools/codemodel'
import {
    ExampleExtension,
    ExampleModel,
    ExampleParameterModel,
    ExtensionName
} from '../model/testModel'
import { Helper } from '../util/helper'
import { Host, startSession } from '@azure-tools/autorest-extension-base'

export async function processRequest(host: Host): Promise<void> {
    const session = await startSession<CodeModel>(host, {}, codeModelSchema)
    const files = await session.listInputs()
    const codemodel = await session.readFile('code-model-v4.yaml')
    const model = session.model
    await AddTestModel(session.model)
    await Helper.outputToModelerfour(host, session)
    await Helper.dumpCodeModel(host, session, 'test-modeler.yaml')
}

async function AddTestModel(codeModel: CodeModel) {
    codeModel.operationGroups.forEach((operationGroup) => {
        operationGroup.operations.forEach((operation) => {
            for (const [exampleName, rawValue] of Object.entries(
                operation.extensions?.[ExtensionName.xMsExamples] ?? {}
            )) {
                const parametersInExample = (rawValue as ExampleExtension).parameters
                const exampleModel = new ExampleModel(exampleName, operation)
                for (const parameter of Helper.allParameters(operation)) {
                    const dotPath = Helper.getFlattenedNames(parameter).join('.')
                    if (Helper.isPathDefined(parametersInExample, dotPath)) {
                        const exampleParameter = new ExampleParameterModel(
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

                operation.extensions[ExtensionName.xMsExampleModels] =
                    operation.extensions[ExtensionName.xMsExampleModels] || {}
                operation.extensions[ExtensionName.xMsExampleModels][exampleName] = exampleModel
            }
        })
    })
}
