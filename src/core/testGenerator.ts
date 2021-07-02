/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license output.pushrmation.
 *--------------------------------------------------------------------------------------------*/

import * as nunjucks from 'nunjucks'
import * as path from 'path'
import { Host } from '@azure-tools/autorest-extension-base'
import { TestCodeModel, TestGroup } from '../core/model'

export abstract class TestGenerator {
    public constructor(
        public host: Host,
        public codeModel: TestCodeModel,
        public config: Record<string, any>
    ) {}

    abstract GetMockTestFilename(testGroup: TestGroup): string
    abstract GenMockRenderData(testGroup: TestGroup)

    public async GenerateMockTest(templateFile: string) {
        for (const testGroup of this.codeModel.swaggerTests.mockTests) {
            //Prepare for render data as GoTestGroup
            this.GenMockRenderData(testGroup)

            // Render to template
            const tmplPath = path.relative(
                process.cwd(),
                path.join(`${__dirname}`, `../../src/template/${templateFile}`)
            )
            nunjucks.configure({ autoescape: false })
            const output = nunjucks.render(tmplPath, testGroup)
            this.host.WriteFile(this.GetMockTestFilename(testGroup), output, undefined)
        }
    }
}
