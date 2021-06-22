import {
    ArraySchema,
    ComplexSchema,
    Languages,
    ObjectSchema,
    Operation,
    Parameter,
    Schema,
    SchemaType
} from '@azure-tools/codemodel'
import { Helper } from '../util/helper'

export enum ExtensionName {
    xMsExamples = 'x-ms-examples',
    xMsExampleModels = 'x-ms-example-models'
}
export interface ExampleExtensionResponse {
    body: any
    headers: Record<string, any>
}
export interface ExampleExtension {
    parameters?: Record<string, any>
    responses?: Record<string, ExampleExtensionResponse>
}

export class ExampleValue {
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

    public static createInstance(rawValue: any, schema: Schema, language: Languages): ExampleValue {
        const instance = new ExampleValue(rawValue, schema, language)
        if (!schema) {
            return instance
        }

        if (schema.type === SchemaType.Array && Array.isArray(rawValue)) {
            instance.value = rawValue.map((x) =>
                this.createInstance(x, (schema as ArraySchema).elementType, undefined)
            )
        } else if (schema.type === SchemaType.Object && rawValue === Object(rawValue)) {
            const childSchema: ComplexSchema = Helper.findInDescents(
                schema as ObjectSchema,
                rawValue
            )

            instance.value = {}
            for (const property of Helper.getAllProperties(childSchema)) {
                if (Object.prototype.hasOwnProperty.call(rawValue, property.serializedName)) {
                    instance.value[property.serializedName] = this.createInstance(
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

export class ExampleParameter {
    /** Ref to the Parameter of operations in codeModel */
    parameter: Parameter
    /** Can be object, list, primitive data, ParameterModel*/
    exampleValue: ExampleValue

    public constructor(parameter: Parameter, rawValue: any) {
        this.parameter = parameter
        this.exampleValue = ExampleValue.createInstance(
            rawValue,
            parameter?.schema,
            parameter.language
        )
    }
}

export class ExampleResponse {
    body: ExampleValue
    headers: Record<string, any>

    public static createInstance(
        rawResponse: ExampleExtensionResponse,
        schema: Schema,
        language: Languages
    ): ExampleResponse {
        const instance = new ExampleResponse()
        instance.body = ExampleValue.createInstance(rawResponse.body, schema, language)
        instance.headers = rawResponse.headers
        return instance
    }
}

export class ExampleModel {
    /** Key in x-ms-examples */
    name: string

    operation: Operation
    clientParameters: Array<ExampleParameter> = []
    methodParameters: Array<ExampleParameter> = []
    responses: Record<string, ExampleResponse> = {} // statusCode-->ExampleResponse

    public constructor(name: string, operation: Operation) {
        this.name = name
        this.operation = operation
    }
}

/**
 * Generally a test group should be generated into one test source file.
 */
export class TestGroup {
    name: string
    properties: Record<string, any> = {} //TODO: setup, cleanup, env
    scenarios: Array<TestScenario> = []
    public constructor(name: string) {
        this.name = name
    }
}

export class TestScenario {
    name: string
    properties: Record<string, any> = {} //TODO: setup, cleanup, env
    examples: Array<ExampleModel> = []
    public constructor(name: string) {
        this.name = name
    }
}

export class TestModel {
    mockTests: Array<TestGroup> = []
    scenarioTests: Array<TestGroup> = []
}
