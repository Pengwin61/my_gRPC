### 
```
protoc --go_out=. --go-grpc_out=. api/proto/adder.proto

protoc --go_out=. --go-grpc_out=plugins=grpc: api/proto/adder.proto

protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. api/proto/adder.proto
```