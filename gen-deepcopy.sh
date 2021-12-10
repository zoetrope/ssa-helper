#!/bin/bash -e

function gen () {
  pkg=$1
  group=$(echo $pkg | awk -F "/" '{ print $(NF-1) }')
  version=$(echo $pkg | awk -F "/" '{ print $NF }')

  cat <<EOF > $pkg/doc.go
// +k8s:deepcopy-gen=package
// +groupName=$group

package $version
EOF

  ../../bin/deepcopy-gen --input-dirs $pkg --output-file-base zz_generated.deepcopy --go-header-file ../../hack/boilerplate.txt --output-base ../../tmp/output/ -v 40
}

cd ./tmp/client-go/
packages="./applyconfigurations/*/*"
for pkg in $packages; do
  if [ -d $pkg ]; then
    gen $pkg
  fi
done

