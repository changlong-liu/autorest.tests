# Autorest tests

Generate *.md config files in Azure REST API specification:

https://github.com/Azure/azure-rest-api-specs

## How to Generate Test (TODO)

    autorest .. 

``` yaml

az:
  tests: true
python: {}
track2: {}

try-require:
  - ./readme.python.md
  - ./readme.cli.md
  - ./readme.az.md
  - ./readme.go.md
  - ./readme.typescript.md

# use-extension:
  # "@autorest/modelerfour": 4.15.421
  # "@autorest/typescript": "1.0.0-beta.1"
  # "@autorest/typescript": "C:\\ZZ\\projects\\codegen\\autorest.typescript"
  # "@microsoft.azure/autorest.modeler": "2.3.38"
  # "@autorest/go": 4.0.0-preview.20
  # "@autorest/go": "C:\\ZZ\\projects\\codegen\\autorest.go"

pipeline:
    test-modeler:
        # input: modelerfour/new-transform
        # input: swagger-document/loader-swagger
        # input: typescript
        # input: modelerfour/identity
        input: go-transform
        # input: go/text-transform
        output-artifact: source-file-test-modeler
    go-tester:
        input: test-modeler
        output-artifact: source-file-go-tester
    tests/emitter:
        input: 
            - test-modeler
            - go-tester
        scope: scope-tests/emitter
    # go/emitter:
    #     input: 
    #         - test-modeler
    #         - go-tester


# modelerfour:
#     lenient-model-deduplication: true
#     group-parameters: true
#     flatten-models: true
#     flatten-payloads: true

scope-tests/emitter:
  input-artifact:
      - source-file-test-modeler
      - source-file-go-tester
  output-uri-expr: $key

  output-artifact: source-file-go-tester
```

``` yaml $(csharp)

try-require:
  - ./readme.csharp.md

use-extension:
  # "@autorest/csharp": 3.0.0-beta.20210614.2
  "@autorest/csharp": 3.0.0-beta.20210210.4

pipeline:
    test-modeler:
        input: csharpgen
        output-artifact: source-file
    csharpgen/emitter:
        input: test-modeler
        scope: source-file
```
