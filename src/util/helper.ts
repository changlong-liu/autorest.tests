import * as _ from 'lodash'
import * as child_process from 'child_process'
import * as fs from 'fs'
import * as path from 'path'
import {
    ChoiceSchema,
    CodeModel,
    ComplexSchema,
    ObjectSchema,
    Operation,
    Parameter,
    Property,
    VirtualParameter,
    codeModelSchema,
    isVirtualParameter
} from '@azure-tools/codemodel'
import { Host, Session } from '@azure-tools/autorest-extension-base'
import { serialize } from '@azure-tools/codegen'

export class Helper {
    public static async outputToModelerfour(
        host: Host,
        session: Session<CodeModel>
    ): Promise<void> {
        // write the final result first which is hardcoded in the Session class to use to build the model..
        // overwrite the modelerfour which should be fine considering our change is backward compatible
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const modelerfourOptions = <any>await session.getValue('modelerfour', {})
        if (modelerfourOptions['emit-yaml-tags'] !== false) {
            host.WriteFile(
                'code-model-v4.yaml',
                serialize(session.model, codeModelSchema),
                undefined,
                'code-model-v4'
            )
        }
        if (modelerfourOptions['emit-yaml-tags'] !== true) {
            host.WriteFile(
                'code-model-v4-no-tags.yaml',
                serialize(session.model),
                undefined,
                'code-model-v4-no-tags'
            )
        }
    }

    public static async dumpCodeModel(
        host: Host,
        session: Session<CodeModel>,
        fileName: string
    ): Promise<void> {
        host.WriteFile(fileName, serialize(session.model), undefined)
    }

    public static allParameters(operation: Operation) {
        const ret: Parameter[] = []
        if (operation.parameters) {
            ret.push(...operation.parameters)
        }
        if (operation.requests[0]?.parameters) {
            ret.push(...operation.requests[0].parameters)
        }
        return ret
    }

    public static getParameterSerializedName(parameter: Parameter) {
        return (
            parameter?.language?.default?.['serializedName'] ||
            parameter?.language?.default?.['name']
        )
    }

    public static getFlattenedNames(parameter: Parameter) {
        let ret = undefined
        if (isVirtualParameter(parameter)) {
            ret = (parameter as VirtualParameter).targetProperty.flattenedNames
        }
        if (!ret) {
            ret = [this.getParameterSerializedName(parameter)]
        }
        return ret
    }

    public static isPathDefined(object: any, dotPath: string): boolean {
        const IMPOSSIBLE_VALUE = 'impossible-value'
        return _.get(object, dotPath, IMPOSSIBLE_VALUE) !== IMPOSSIBLE_VALUE
    }

    public static getPathValue(object: any, dotPath: string): any {
        return _.get(object, dotPath)
    }

    public static findInDescents(schema: ObjectSchema, value: Record<string, any>): ComplexSchema {
        if (schema.discriminator) {
            const discriminatorKey = schema.discriminator.property.serializedName
            if (
                Object.prototype.hasOwnProperty.call(value, discriminatorKey) &&
                Object.prototype.hasOwnProperty.call(
                    schema.discriminator.all,
                    value[discriminatorKey]
                )
            ) {
                return schema.discriminator.all[value[discriminatorKey]]
            }
        }

        //TODO: find most matched child by properties if no discriminator
        return schema
    }

    public static getAllProperties(schema: ComplexSchema, withParents = false): Array<Property> {
        const ret: Array<Property> = []
        if (Object.prototype.hasOwnProperty.call(schema, 'properties')) {
            ret.push(...(schema as ObjectSchema).properties)
        }
        if (withParents && Object.prototype.hasOwnProperty.call(schema, 'parents')) {
            for (const parent of (schema as ObjectSchema).parents.immediate) {
                ret.push(...this.getAllProperties(parent))
            }
        }
        return ret
    }

    public static escapeString(str: string): string {
        return str.split('\\').join('\\\\').split('"').join('\\"')
    }

    public static findChoiceValue(schema: ChoiceSchema, rawValue: any) {
        for (const choiceValue of schema.choices) {
            if (choiceValue.value == rawValue) {
                return choiceValue
            }
        }

        if (schema.choices?.length > 0) {
            console.warn(`${rawValue} is NOT a valid ${schema.language.default.name} value`)
            return schema.choices[0]
        }
        throw `${schema.language.default.name} has no choices!`
    }

    public static execSync(command: string) {
        child_process.execSync(command)
    }

    public static deleteFolderRecursive(target) {
        if (fs.existsSync(target)) {
            fs.readdirSync(target).forEach((file, _) => {
                const curPath = path.join(target, file)
                if (fs.lstatSync(curPath).isDirectory()) {
                    // recurse
                    this.deleteFolderRecursive(curPath)
                } else {
                    // delete file
                    fs.unlinkSync(curPath)
                }
            })
            fs.rmdirSync(target)
        }
    }
}
