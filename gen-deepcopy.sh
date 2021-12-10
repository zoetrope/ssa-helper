#!/bin/bash -e

cat <<EOF > ./tmp/client-go/applyconfigurations/core/v1/doc.go
// +k8s:deepcopy-gen=package
// +groupName=core

package v1
EOF

cd ./tmp/client-go/
../../bin/deepcopy-gen --input-dirs ./applyconfigurations/core/v1/ --output-file-base zz_generated.deepcopy --go-header-file ../../hack/boilerplate.txt --output-base ../../tmp/output/ -v 40

