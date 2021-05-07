# go-commit-log

A simple `commit log` example for distributed computing.

See: [Distributed Services with Go](https://pragprog.com/titles/tjgo/distributed-services-with-go/)

---

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
