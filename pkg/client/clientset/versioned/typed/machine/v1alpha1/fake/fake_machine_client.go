// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/gardener/machine-controller-manager-provider-aws/pkg/client/clientset/versioned/typed/machine/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMachineV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMachineV1alpha1) AWSMachineClasses(namespace string) v1alpha1.AWSMachineClassInterface {
	return &FakeAWSMachineClasses{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMachineV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
