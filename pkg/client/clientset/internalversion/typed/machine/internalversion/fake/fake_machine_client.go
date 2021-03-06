// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	internalversion "github.com/gardener/machine-controller-manager-provider-aws/pkg/client/clientset/internalversion/typed/machine/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMachine struct {
	*testing.Fake
}

func (c *FakeMachine) AWSMachineClasses(namespace string) internalversion.AWSMachineClassInterface {
	return &FakeAWSMachineClasses{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMachine) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
