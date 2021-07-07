# Bug reproduction for Goa Issue #2887

## Steps

- Run `go run main.go`

Note the errors originating from `gen/http/cli/*` packages. Also note that `gen/http/cli/calc/cli.go` references methods from zoo service and `gen/http/cli/zoo/cli.go` references methods from calc service.
