#!/usr/bin/env bash

#protoc --go_out=plugins=grpc:. google/api/*.proto
#*.proto
protoc -I. -I%GOPATH%/src -I%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. api/auth.proto  api/tenant.proto api/role.proto api/user.proto api/domain.proto api/actionlog.proto api/service.proto api/policy.proto api/payment.proto api/trading.proto
protoc -I. -I%GOPATH%/src -I%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:./api/swagger api/auth.proto  api/tenant.proto api/role.proto api/user.proto api/domain.proto api/service.proto api/policy.proto api/payment.proto api/trading.proto
protoc -I. -I%GOPATH%/src -I%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. api/auth.proto  api/tenant.proto api/role.proto api/user.proto api/domain.proto api/service.proto api/policy.proto api/payment.proto api/trading.proto
go run api/swagger/main.go api/swagger/api > static/swagger/api.swagger.json
go-bindata -prefix static/ -pkg static -o internal/static/static_gen.go static/...
