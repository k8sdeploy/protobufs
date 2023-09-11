# All .proto files
PROTO_FILES := $(wildcard v1/*.proto)

# Output directories
OUT_DIRS := $(patsubst v1/%.proto, generated/%, $(PROTO_FILES))

# Sentinel files to track completion
SENTINELS := $(patsubst v1/%.proto, generated/%/.done, $(PROTO_FILES))

# Default target
.PHONY: all
all: generate clean-sentinels

.PHONY: setup
setup:
	mkdir -p $(OUT_DIRS)
	$(foreach dir,$(OUT_DIRS),[ -e $(dir)/.keep ] || touch $(dir)/.keep;)

.PHONY: generate
generate: setup $(SENTINELS)

# Generate Go code from .proto files and use a sentinel file to signal completion
generated/%/.done: v1/%.proto
	mkdir -p $(dir $@)
	protoc -I . \
		--go_out $(dir $@) --go_opt paths=source_relative \
		--go-grpc_out $(dir $@) --go-grpc_opt paths=source_relative \
		$<
	touch $@
	[ -e $(dir $@).keep ] || touch $(dir $@).keep

.PHONY: clean-sentinels
clean-sentinels:
	rm -f $(SENTINELS)

# Clean generated files but leave .keep files
.PHONY: clean
clean:
	find generated -type d -name "v1" -exec rm -rf {} +
