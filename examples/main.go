package main

import (
	"fmt"
	v1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/core/v1"
)

func main() {
	tmp := v1.PodTemplate("a", "b").WithClusterName("test")

	var out v1.PodTemplateApplyConfiguration
	tmp.DeepCopyInto(&out)
	out.WithNamespace("c")

	fmt.Printf("src: %s, dst: %s\n", *tmp.Namespace, *out.Namespace)
	fmt.Printf("src: %s, dst: %s\n", *tmp.Name, *out.Name)
}
