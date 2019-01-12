package client

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v2/tokens"
	"github.com/morrisson/load_gen/common"
)

func GetV2Token() *tokens.Token {
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: common.KeystoneEndpointURL,
		Username:         common.GohanUser,
		Password:         common.GohanPassword,
		TenantName:       common.GohanTenant,
		AllowReauth:      true,
	}

	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		panic(err)
	}

	client, err := openstack.NewIdentityV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		panic(err)
	}

	opts := tokens.AuthOptions{
		IdentityEndpoint: common.KeystoneEndpointURL,
		Username:         common.GohanUser,
		Password:         common.GohanPassword,
		TenantName:       common.GohanTenant,
		AllowReauth:      true,
	}

	token, err := tokens.Create(client, opts).ExtractToken()
	if err != nil {
		panic(err)
	}

	return token
}
