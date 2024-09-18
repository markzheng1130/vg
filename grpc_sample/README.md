# Install the Protocol Buffers compiler
brew install protobuf  # For macOS, or use the appropriate package manager for your system

# Install the Go Protocol Buffers plugin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Install the Go gRPC plugin
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH="$PATH:$(go env GOPATH)/bin"

protoc --go_out=. --go-grpc_out=. greet.proto

go run server.go

go run client.go
