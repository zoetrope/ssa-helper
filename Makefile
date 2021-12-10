
BIN_DIR := $(shell pwd)/bin
DEEPCOPY_GEN := $(BIN_DIR)/deepcopy-gen

prepare:
	rm -rf ./tmp/client-go
	mkdir -p ./tmp/client-go
	git clone --depth 1 https://github.com/kubernetes/client-go ./tmp/client-go

generate:
	rm -rf ./tmp/output
	mkdir -p ./tmp/output
	./gen-deepcopy.sh

.PHONY: setup
setup: $(DEEPCOPY_GEN)

$(DEEPCOPY_GEN):
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install k8s.io/code-generator/cmd/deepcopy-gen@latest
