package apis

import (
	"github.com/rajdurvasula/rd-tomcat-operator/tree/master/pkg/apis/cache/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemaBuilder.AddToScheme)
}
