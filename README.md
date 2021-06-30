# Autorest tests

Generate *.md config files in Azure REST API specification:

https://github.com/Azure/azure-rest-api-specs

## How to Generate Test (GoLang)

    autorest --version=3.2.1 --use=. d:\projects\codegen\azure-rest-api-specs\specification\agrifood\resource-manager\readme.md --use=@autorest/go@4.0.0-preview.20 --file-prefix="zz_generated_" --track2 --go --output-folder=D:\projects\codegen\generated\azure-sdk-for-go\armagrifood --debug

``` yaml $(go)
clear-output-folder: false
```

``` yaml $(go)

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
    go-linter:
        input:
            - go-tester
            - tests/emitter
    # go/emitter:
    #     input: 
    #         - test-modeler
    #         - go-tester
    go/emitter:
        output-artifact: source-file-go-tester


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

``` yaml $(go) && $(generate-sdk)
pipeline:
    go/emitter:
        output-artifact: source-file-go
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
