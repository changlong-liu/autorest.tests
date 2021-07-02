/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license output.pushrmation.
 *--------------------------------------------------------------------------------------------*/

import { Helper } from '../util/helper'
import { Host, startSession } from '@azure-tools/autorest-extension-base'
import { TestCodeModel } from './model'
import { codeModelSchema } from '@azure-tools/codemodel'

export async function processRequest(host: Host): Promise<void> {
    const session = await startSession<TestCodeModel>(host, {}, codeModelSchema)
    // const files = await session.listInputs()
    // const codemodel = await session.readFile('code-model-v4.yaml')
    const codeModel = new TestCodeModel(session.model)
    await codeModel.genExampleModels()
    await codeModel.genMockTests()
    await Helper.outputToModelerfour(host, session)
    if (session.getValue('debug', false)) {
        await Helper.dumpCodeModel(host, session, 'test-modeler.yaml')
    }
}
