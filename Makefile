
BIN_DIR := $(shell pwd)/bin
DEEPCOPY_GEN := $(BIN_DIR)/deepcopy-gen
CLIENT_GO_VERSION=v0.22.4

prepare:
	rm -rf ./tmp/client-go
	mkdir -p ./tmp/client-go
	git clone --depth 1 -b $(CLIENT_GO_VERSION) https://github.com/kubernetes/client-go ./tmp/client-go

generate:
	rm -rf ./applyconfigurations
	rm -rf ./tmp/output
	mkdir -p ./tmp/output
	./gen-deepcopy.sh
	cp -r ./tmp/output/github.com/zoetrope/ac-deepcopy/tmp/client-go/applyconfigurations  applyconfigurations

.PHONY: setup
setup: $(DEEPCOPY_GEN)

$(DEEPCOPY_GEN):
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install k8s.io/code-generator/cmd/deepcopy-gen@latest
