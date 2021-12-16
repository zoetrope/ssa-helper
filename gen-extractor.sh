#!/bin/bash -e

function gen () {
  read type_name orig_type <<< $(awk 'match($0, /^func extract(\w+)\(\w+ \*(\w+\.\w+), fieldManager string, subresource string\) \(\*(\w+)ApplyConfiguration, error\) {$/, p){print p[1], p[2]}' $file)

  cat <<EOF >> $file
func (b *${type_name}ApplyConfiguration) Original() client.Object {
	return &${orig_type}{}
}

func (b *${type_name}ApplyConfiguration) Extract(obj client.Object, fieldManager string, subresource string) (*${type_name}ApplyConfiguration, error) {
	return extract${type_name}(obj.(*${orig_type}), fieldManager, subresource)
}
EOF

  if grep -q "b.WithNamespace(" $file; then
    cat <<EOF >> $file
func (b *${type_name}ApplyConfiguration) ObjectKey() (client.ObjectKey, error) {
	if b.Namespace == nil {
		return client.ObjectKey{}, errors.New("The ${type_name}ApplyConfiguration namespace should not be empty.")
	}
	if b.Name == nil {
		return client.ObjectKey{}, errors.New("The ${type_name}ApplyConfiguration name should not be empty.")
	}
	return client.ObjectKey{
		Name:      *b.Name,
		Namespace: *b.Namespace,
	}, nil
}
EOF
  else
    cat <<EOF >> $file
func (b *${type_name}ApplyConfiguration) ObjectKey() (client.ObjectKey, error) {
	if b.Name == nil {
		return client.ObjectKey{}, errors.New("The ${type_name}ApplyConfiguration name should not be empty.")
	}
	return client.ObjectKey{
		Name:      *b.Name,
	}, nil
}
EOF
  fi
}

files=$(grep -rl "func extract" ./applyconfigurations/)
for file in $files; do
  gen $file
done

