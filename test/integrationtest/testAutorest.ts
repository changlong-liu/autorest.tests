import * as assert from 'assert'
import * as fs from 'fs'
import * as path from 'path'
import { Helper } from '../../src/util/helper'
import { exec } from 'child_process'

async function compare(dir1: string, dir2: string) {
    const cmd = 'diff -r --exclude=gen.zip --strip-trailing-cr ' + dir1 + ' ' + dir2
    console.log(cmd)
    return await new Promise<boolean>((resolve, reject) => {
        exec(cmd, function (error, stdout) {
            if (error) {
                console.log('exec error: ' + error + ', ' + stdout)
                // Reject if there is an error:
                return reject(false)
            }
            // Otherwise resolve the promise:
            return resolve(true)
        })
    })
}

async function runAutorest(readmePath: string, extraOption: string[]) {
    const cmd =
        path.join(`${__dirname}`, '..', '..' + '/node_modules', '.bin', 'autorest') +
        ' --version=3.2.1 --use=' +
        path.join(`${__dirname}`, '..', '..') +
        ' ' +
        readmePath +
        ' ' +
        extraOption.join(' ')
    console.log(cmd)
    return await new Promise<boolean>((resolve, reject) => {
        exec(cmd, function (error, stdout, stderr) {
            if (error) {
                console.error('exec error: ' + stderr)
                // Reject if there is an error:
                return reject(false)
            }
            // Otherwise resolve the promise:
            console.log(stdout)
            return resolve(true)
        })
    })
}
function getExtraOption(outputFolder: string, lang: string) {
    let ret = [`--output-folder=${outputFolder}`]
    if (lang == 'golang') {
        ret = ret.concat([
            '--use=@autorest/go@4.0.0-preview.22',
            '--file-prefix="zz_generated_"',
            '--track2',
            '--go',
            '--debug'
        ])
    }
    return ret
}

async function runSingleTest(
    swaggerDir: string,
    rp: string,
    extraOption: string[],
    outputFolder: string,
    tempOutputFolder: string
) {
    let result = true
    let msg = ''
    const readmePath = path.join(swaggerDir, 'specification', rp, 'resource-manager/readme.md')
    await runAutorest(readmePath, extraOption)
        .then((res) => {
            if (!res) {
                msg = 'Run autorest not successfully!'
            }
            result = res
        })
        .catch((e) => {
            msg = `Run autorest failed! ${e}`
            result = false
        })
    if (result) {
        await compare(outputFolder, tempOutputFolder)
            .then((res1) => {
                if (res1 === false) {
                    msg = 'The generated files have changed!'
                }
                result = res1
            })
            .catch((e) => {
                msg = `The diff has some error ${e}`
                result = false
            })
    } else {
        console.error(msg)
    }
    return result
}

describe('Run autorest and compare the output', () => {
    beforeAll(async () => {
        //
    })

    afterAll(async () => {
        //
    })

    it('', async () => {
        jest.setTimeout(8 * 60 * 60 * 1000)
        const swaggerDir = path.join(`${__dirname}`, '../swagger')
        const outputDir = path.join(`${__dirname}`, 'output')
        const tempoutputDir = path.join(`${__dirname}`, 'tempoutput')
        const msg = ''
        let finalResult = true
        const allTests: boolean[] = []
        for (const rp of ['compute']) {
            for (const lang of ['golang']) {
                let result = true
                console.log('Start Processing: ' + rp)

                // Remove tmpoutput
                const outputFolder = path.join(outputDir, rp, lang)
                const tempOutputFolder = path.join(tempoutputDir, rp, lang)
                Helper.deleteFolderRecursive(tempOutputFolder)
                fs.mkdirSync(tempOutputFolder, { recursive: true })

                try {
                    const test = await runSingleTest(
                        swaggerDir,
                        rp,
                        getExtraOption(tempOutputFolder, lang),
                        outputFolder,
                        tempOutputFolder
                    )
                    allTests.push(test)
                } catch (error) {
                    console.log(msg)
                    result = false
                    break
                }
                if (!result) {
                    finalResult = false
                }
            }
        }
        finalResult = (await Promise.all(allTests)).every((x) => x)
        assert.strictEqual(finalResult, true, msg)
    })
})
