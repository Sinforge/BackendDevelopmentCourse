grpc_prep:
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go

grpc_gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/protobuf/de_logic.proto
	
