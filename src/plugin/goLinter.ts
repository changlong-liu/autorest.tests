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
import { Host, startSession } from '@azure-tools/autorest-extension-base'

export class GoLinter {
    async process(fileName: string): Promise<void> {
        const scriptName = path.join('dist/src/python/install.py') + ' ' + fileName
    }
}

export async function processRequest(host: Host): Promise<void> {
    const config = await host.GetValue('')
    const files = await host.ListInputs()
    const session = await startSession<TestCodeModel>(host, {}, codeModelSchema)
    const config2 = await session.getValue('')
}
