# go-commit-log

A simple `commit log` example for distributed computing.

See: [Distributed Services with Go](https://pragprog.com/titles/tjgo/distributed-services-with-go/)

---

## Build

### Run Server

```bash
go run cmd/server/main.go
```

## Add Data

```bash
$ curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzEK"}}'
$ curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzIK"}}'
$ curl -X POST localhost:8080 -d \
'{"record": {"value": "TGV0J3MgR28gIzMK"}}'
```

---

### Read Data

```bash
$ curl -X GET localhost:8080 -d '{"offset": 0}'
$ curl -X GET localhost:8080 -d '{"offset": 1}'
$ curl -X GET localhost:8080 -d '{"offset": 2}'
```

---

## Protobuf

1. Download the latest version of [protoc](https://github.com/protocolbuffers/protobuf/releases) for the target system.

    ```bash
    sudo apt install protobuf-compiler
    sudo apt install golang-goprotobuf-dev
    ```

2. Add `proto` files

    ```bash
    mkdir -p api/v1
    touch api/v1/log.protot
    ```

3. Install protobuf into project

    ```bash
    go get google.golang.org/protobuf/...@v1.25.0
    ```

4. Generate

    ```bash
    protoc api/v1/*.proto \
        --go_out=. \
        --go_opt=paths=source_relative \
        --proto_path=.
    ```
