/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license output.pushrmation.
 *--------------------------------------------------------------------------------------------*/

import * as nunjucks from 'nunjucks'
import * as path from 'path'
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
    TestCodeModel,
    TestGroup,
    TestModel,
    TestScenario
} from '../model/testModel'
import { Helper } from '../util/helper'
import { Host, startSession } from '@azure-tools/autorest-extension-base'

export async function processRequest(host: Host): Promise<void> {
    const session = await startSession<TestCodeModel>(host, {}, codeModelSchema)
    const config = await session.getValue('')
    host.WriteFile('mock_test.go', await GenerateMockTest(session.model, config), undefined)
    await Helper.outputToModelerfour(host, session)
    // await Helper.dumpCodeModel(host, session, 'test-modeler.yaml')
}

async function GenerateMockTest(codeModel: TestCodeModel, config: Record<string, any>) {
    //Prepare for render data as GoTestData
    const data = codeModel.testModel.mockTests[0] as GoTestData

    // Render to template
    const tmplPath = path.relative(
        process.cwd(),
        path.join(`${__dirname}`, '../../src/template/mockTest.go.njk')
    )
    nunjucks.configure({ autoescape: false })
    const output = nunjucks.render(tmplPath, data)
    return output
}

class GoTestData extends TestGroup {
    packageName: string
}
