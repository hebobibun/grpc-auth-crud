proto:
	protoc --proto_path=pkg \
			--go_out=pkg --go_opt=paths=source_relative \
			--go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
			pkg/**/pb/*.proto

run:
	go run cmd/main.go