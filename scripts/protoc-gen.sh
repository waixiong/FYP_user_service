# # reserve proxy (number 5 https://github.com/grpc-ecosystem/grpc-gateway)

# protoc --proto_path=api/proto/v1 --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --go-grpc_out=plugins=grpc:pkg/api/v1 user.proto
protoc --proto_path=api/proto/v1 \
    --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    --go_out ./pkg/api/v1\
    --go-grpc_out ./pkg/api/v1 service.proto

# protoc --proto_path=api/proto/v1 --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:pkg/api/v1 user.proto
protoc --proto_path=api/proto/v1 \
    --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    --grpc-gateway_out ./pkg/api/v1 \
    --grpc-gateway_opt logtostderr=true\
    service.proto

protoc --proto_path=api/proto/v1 --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:api/swagger/v1 service.proto
# protoc --proto_path=api/proto/v1 --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --doc_out=./docs/api/grpc user.proto
# spectacle ./api/swagger/v1/user.swagger.json -t ./docs/api/swagger/

# protoc --proto_path=api/proto/v1 --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --dart_out=grpc:pkg/api/v1/dart user.proto
