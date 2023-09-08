.PHONY: setup
setup:
	mkdir -p generated
	mkdir -p generated/agent_service

.PHONY: generate
generate: setup agent_service

.PHONY: agent_service
agent_service:
	protoc -I . \
		--go_out ./generated/agent_service --go_opt paths=source_relative \
		--go-grpc_out ./generated/agent_service --go-grpc_opt paths=source_relative \
		v1/agent_service.proto

