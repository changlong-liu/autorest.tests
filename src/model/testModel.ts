import { Operation, Parameter } from '@azure-tools/codemodel'

export enum ExtensionName {
    xMsExamples = 'x-ms-examples',
    xMsExampleModels = 'x-ms-example-models'
}

export class ParameterModel {
    /** Ref to the Parameter of operations in codeModel */
    parameter: Parameter
    /** Can be object, list, primitive data, ParameterModel*/
    value: any
}

export class ExampleModel {
    /** Key in x-ms-examples */
    name: string

    operation: Operation
    clientParameters: Array<ParameterModel>
    methodParameters: Array<ParameterModel>

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
