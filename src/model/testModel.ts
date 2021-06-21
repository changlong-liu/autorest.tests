import {
    ArraySchema,
    Languages,
    ObjectSchema,
    Operation,
    Parameter,
    Schema,
    SchemaType
} from '@azure-tools/codemodel'

export enum ExtensionName {
    xMsExamples = 'x-ms-examples',
    xMsExampleModels = 'x-ms-example-models'
}

export interface ExampleExtension {
    parameters?: Record<string, any>
    responses?: Record<string, any>
}

export class ExampleValueModel {
    language: Languages
    schema: Schema
    value: any

    public constructor(
        value: any = undefined,
        schema: Schema = undefined,
        language: Languages = undefined
    ) {
        this.language = language
        this.schema = schema
        this.value = value
    }

    public static createInstance(
        rawValue: any,
        schema: Schema,
        language: Languages
    ): ExampleValueModel {
        const instance = new ExampleValueModel(rawValue, schema, language)
        if (!schema) {
            // keep raw value
        } else if (schema.type === SchemaType.Array && Array.isArray(rawValue)) {
            instance.value = rawValue.map((x) =>
                this.createInstance(x, (schema as ArraySchema).elementType, undefined)
            )
        } else if (schema.type === SchemaType.Object && rawValue === Object(rawValue)) {
            instance.value = {}
            for (const property of (schema as ObjectSchema).properties) {
                if (Object.prototype.hasOwnProperty.call(rawValue, property.serializedName)) {
                    instance[property.serializedName] = this.createInstance(
                        rawValue[property.serializedName],
                        property.schema,
                        property.language
                    )
                }
            }
        }
        return instance
    }
}

export class ExampleParameterModel {
    /** Ref to the Parameter of operations in codeModel */
    parameter: Parameter
    /** Can be object, list, primitive data, ParameterModel*/
    exampleValue: ExampleValueModel

    public constructor(parameter: Parameter, rawValue: any) {
        this.parameter = parameter
        this.exampleValue = ExampleValueModel.createInstance(
            rawValue,
            parameter?.schema,
            parameter.language
        )
    }
}

export class ExampleModel {
    /** Key in x-ms-examples */
    name: string

    operation: Operation
    clientParameters: Array<ExampleParameterModel>
    methodParameters: Array<ExampleParameterModel>

    public constructor(name: string, operation: Operation) {
        this.name = name
        ;(this.operation = operation), (this.clientParameters = [])
        this.methodParameters = []
    }
}

export class ScenarioModel {
    name: string
    examples: Array<ExampleModel>
}

export interface RawExample {
    name: string
    examples: Array<ExampleModel>
}
