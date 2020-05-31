Steps:
1. `go test -race -coverprofile=coverage.txt -covermode=atomic ./...
2. bash <(curl -s https://codecov.io/bash)`

To have these set-up on GitHub so that they are run on each push or pull request,
check the [push.yml](.github/workflows/push.yml) file.