import {
    ArraySchema,
    CodeModel,
    ComplexSchema,
    DictionarySchema,
    ImplementationLocation,
    Languages,
    ObjectSchema,
    Operation,
    OperationGroup,
    Parameter,
    Schema,
    SchemaResponse,
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
    /**The value is {string:ExampleValue}/[ExampleValue] instance if schema.type==Object/Array, otherwise it's rawValue */
    value: any
    parentsValue?: Record<string, ExampleValue> // parent class Name--> value

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

            instance.parentsValue = {}
            if (Object.prototype.hasOwnProperty.call(childSchema, 'parents')) {
                for (const parent of (schema as ObjectSchema).parents.immediate) {
                    if (schema.type === SchemaType.Object) {
                        const parentValue = this.createInstance(rawValue, parent, parent.language)
                        if (
                            Object.keys(parentValue.value).length !== 0 ||
                            Object.keys(parentValue.parentsValue).length !== 0
                        ) {
                            instance.parentsValue[parent.language.default.name] = parentValue
                        }
                    } else {
                        console.warn(
                            `${parent.language.default.name} is NOT a object type of parent of ${childSchema.language.default.name}!`
                        )
                    }
                }
            }
        } else if (schema.type === SchemaType.Dictionary && rawValue === Object(rawValue)) {
            instance.value = {}
            for (const [key, value] of Object.entries(rawValue)) {
                instance.value[key] = this.createInstance(
                    value,
                    (<DictionarySchema>schema).elementType,
                    undefined
                )
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
    operationGroup: OperationGroup
    operation: Operation
    clientParameters: Array<ExampleParameter> = []
    methodParameters: Array<ExampleParameter> = []
    responses: Record<string, ExampleResponse> = {} // statusCode-->ExampleResponse

    public constructor(name: string, operation: Operation, operationGroup: OperationGroup) {
        this.name = name
        this.operation = operation
        this.operationGroup = operationGroup
    }
}

/**
 * Generally a test group should be generated into one test source file.
 */
export class TestGroup {
    name: string
    properties: Record<string, any> = {} //TODO: setup/cleanup/env with oav
    scenarios: Array<TestScenario> = []
    public constructor(name: string) {
        this.name = name
    }
}

export class TestScenario {
    name: string
    properties: Record<string, any> = {} //TODO: setup/cleanup/env with oav
    examples: Array<ExampleModel> = []
    public constructor(name: string) {
        this.name = name
    }
}

export class Tests {
    mockTests: Array<TestGroup> = []
    scenarioTests: Array<TestGroup> = []
}

export interface TestCodeModel extends CodeModel {
    tests?: Tests
}

export class TestCodeModel {
    public constructor(public codeModel: TestCodeModel) {}

    public async genExampleModels() {
        this.codeModel.operationGroups.forEach((operationGroup) => {
            operationGroup.operations.forEach((operation) => {
                for (const [exampleName, rawValue] of Object.entries(
                    operation.extensions?.[ExtensionName.xMsExamples] ?? {}
                )) {
                    const exampleExtension = rawValue as ExampleExtension
                    const parametersInExample = exampleExtension.parameters
                    const exampleModel = new ExampleModel(exampleName, operation, operationGroup)
                    for (const parameter of Helper.allParameters(operation)) {
                        const dotPath = Helper.getFlattenedNames(parameter).join('.')
                        if (Helper.isPathDefined(parametersInExample, dotPath)) {
                            const exampleParameter = new ExampleParameter(
                                parameter,
                                parametersInExample[dotPath]
                            )
                            if (parameter.implementation == ImplementationLocation.Method) {
                                exampleModel.methodParameters.push(exampleParameter)
                            } else if (parameter.implementation == ImplementationLocation.Client) {
                                exampleModel.clientParameters.push(exampleParameter)
                            } else {
                                //
                            }
                        }
                    }

                    for (const [statusCode, response] of Object.entries(
                        exampleExtension.responses
                    )) {
                        const exampleExtensionResponse = response as ExampleExtensionResponse
                        const schemaResponse = operation.responses[0] as SchemaResponse
                        exampleModel.responses[statusCode] = ExampleResponse.createInstance(
                            exampleExtensionResponse,
                            schemaResponse.schema,
                            schemaResponse.language
                        )
                    }

                    operation.extensions[ExtensionName.xMsExampleModels] =
                        operation.extensions[ExtensionName.xMsExampleModels] || {}
                    operation.extensions[ExtensionName.xMsExampleModels][exampleName] = exampleModel
                }
            })
        })
    }

    public initiateTests() {
        if (!this.codeModel.tests) {
            this.codeModel.tests = new Tests()
        }
    }

    public async genMockTests() {
        this.initiateTests()
        const testGroup = new TestGroup('mock')
        this.codeModel.operationGroups.forEach((operationGroup) => {
            operationGroup.operations.forEach((operation) => {
                const testScenario = new TestScenario(
                    operationGroup.language.default.name + '_' + operation.language.default.name
                )
                for (const [_, exampleModel] of Object.entries(
                    operation.extensions?.[ExtensionName.xMsExampleModels] ?? {}
                )) {
                    testScenario.examples.push(exampleModel as ExampleModel)
                }
                testGroup.scenarios.push(testScenario)
            })
        })
        this.codeModel.tests.mockTests.push(testGroup)
    }
}
