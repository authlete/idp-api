    docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.11.0 batch /local/go-client-idp.yml
    go install golang.org/x/tools/cmd/goimports@latest
    goimports -w *.go
    go mod tidy
    go mod download