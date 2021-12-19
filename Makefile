
BIN_DIR := $(shell pwd)/bin
DEEPCOPY_GEN := $(BIN_DIR)/deepcopy-gen
CLIENT_GO_VERSION=v0.23.0

prepare:
	rm -rf ./tmp/client-go
	mkdir -p ./tmp/client-go
	git clone --depth 1 -b $(CLIENT_GO_VERSION) https://github.com/kubernetes/client-go ./tmp/client-go

generate-deepcopy:
	rm -rf ./applyconfigurations
	rm -rf ./tmp/output
	mkdir -p ./tmp/output
	./gen-deepcopy.sh
	cp -r ./tmp/output/github.com/zoetrope/ssa-helper/tmp/client-go/applyconfigurations/ applyconfigurations
	cp -r ./tmp/client-go/applyconfigurations/* applyconfigurations
	find ./applyconfigurations/ -type f | xargs sed -i "s#k8s.io/client-go/applyconfigurations#github.com/zoetrope/ssa-helper/applyconfigurations#g"

generate-extractor:
	./gen-extractor.sh
	find ./applyconfigurations/ -print | grep --regex '.*\.go' | xargs goimports -w
	cp ./hack/apply.go.txt ./applyconfigurations/apply.go

.PHONY: setup
setup: $(DEEPCOPY_GEN)

$(DEEPCOPY_GEN):
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install k8s.io/code-generator/cmd/deepcopy-gen@latest
