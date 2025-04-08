PROTO_DIR=api/protobuf
GEN_DIR=gen/proto

.PHONY: proto
proto:
	@echo "Generating protobuf code..."
	protoc -I $(PROTO_DIR) \
	-I ./third_party \
	--go_out=$(GEN_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(GEN_DIR) \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=$(GEN_DIR) \
	--grpc-gateway_opt=paths=source_relative \
	$(PROTO_DIR)/auth/v1/auth.proto \
	$(PROTO_DIR)/movie/v1/movie.proto \
	$(PROTO_DIR)/booking/v1/booking.proto

