spot_proto:
	protoc \
		-I=./Protobuf/proto \
		--go-grpc_out=./Protobuf/gen \
		--go-grpc_opt=paths=source_relative \
		--go_out=./Protobuf/gen \
		--go_opt=paths=source_relative \
		--validate_out="lang=go:Protobuf/gen" \
		--validate_opt=paths=source_relative \
		spot_service/spot_service.proto

order_proto:
	protoc \
		-I=./Protobuf/proto \
		--go-grpc_out=./Protobuf/gen \
		--go-grpc_opt=paths=source_relative \
		--go_out=./Protobuf/gen \
		--go_opt=paths=source_relative \
		--validate_out="lang=go:Protobuf/gen" \
		--validate_opt=paths=source_relative \
		order_service/order_service.proto

user_proto:
	protoc \
		-I=./Protobuf/proto \
		--go-grpc_out=./Protobuf/gen \
		--go-grpc_opt=paths=source_relative \
		--go_out=./Protobuf/gen \
		--go_opt=paths=source_relative \
		--validate_out="lang=go:Protobuf/gen" \
		--validate_opt=paths=source_relative \
		user_service/user_service.proto