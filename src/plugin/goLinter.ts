import * as path from 'path'
import { Helper } from '../util/helper'
import { Host } from '@azure-tools/autorest-extension-base'

export async function processRequest(host: Host): Promise<void> {
    const config = await host.GetValue('')
    const files = await host.ListInputs()
    for (const outputFile of files) {
        if (outputFile.endsWith('.go')) {
            const pathName = path.join(config['output-folder'], outputFile)
            Helper.execSync(`go fmt ${pathName}`)
        }
    }
}
