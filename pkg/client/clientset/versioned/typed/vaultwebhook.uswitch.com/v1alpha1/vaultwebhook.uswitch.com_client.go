// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/uswitch/vault-webhook/pkg/apis/vaultwebhook.uswitch.com/v1alpha1"
	"github.com/uswitch/vault-webhook/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type VaultwebhookV1alpha1Interface interface {
	RESTClient() rest.Interface
	DatabaseCredentialBindingsGetter
}

// VaultwebhookV1alpha1Client is used to interact with features provided by the vaultwebhook.uswitch.com group.
type VaultwebhookV1alpha1Client struct {
	restClient rest.Interface
}

func (c *VaultwebhookV1alpha1Client) DatabaseCredentialBindings(namespace string) DatabaseCredentialBindingInterface {
	return newDatabaseCredentialBindings(c, namespace)
}

// NewForConfig creates a new VaultwebhookV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*VaultwebhookV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &VaultwebhookV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new VaultwebhookV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *VaultwebhookV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new VaultwebhookV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *VaultwebhookV1alpha1Client {
	return &VaultwebhookV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *VaultwebhookV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
