/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license output.pushrmation.
 *--------------------------------------------------------------------------------------------*/

import { CodeModel, codeModelSchema } from '@azure-tools/codemodel'
import { Host, startSession } from '@azure-tools/autorest-extension-base'

export async function processRequest(host: Host): Promise<void> {
    const session = await startSession<CodeModel>(host, {}, codeModelSchema)
    const files = await session.listInputs()
    const codemodel = await session.readFile('code-model-v4.yaml')
    console.log(files)
    const model = session.model
}
