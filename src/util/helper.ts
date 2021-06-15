import * as _ from 'lodash'
import {
    AnySchema,
    ArraySchema,
    ChoiceSchema,
    ChoiceValue,
    CodeModel,
    ConstantSchema,
    DictionarySchema,
    ObjectSchema,
    Operation,
    OperationGroup,
    Parameter,
    Property,
    Schema,
    SchemaType,
    SealedChoiceSchema,
    StringSchema,
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

    public static getFlattenedNames(parameter: Parameter) {
        let ret = undefined
        if (isVirtualParameter(parameter)) {
            ret = (parameter as VirtualParameter).targetProperty.flattenedNames
        }
        if (!ret) {
            ret = [parameter.language.default.name]
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
}
