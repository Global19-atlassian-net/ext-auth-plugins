package main

import (
	"context"
	v2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"github.com/solo-io/ext-auth-plugins/api"
)

func CreatePlugin() api.ExtauthPlugin {
	return &AuthorizeAllPlugin{}
}

func main() {}

var _ api.ExtauthPlugin = new(AuthorizeAllPlugin)

type AuthorizeAllPlugin struct {}

func (p *AuthorizeAllPlugin) NewConfigInstance(ctx context.Context) interface{} {
	return nil
}

func (p *AuthorizeAllPlugin) GetAuthClient(ctx context.Context, configInstance interface{}) (api.AuthClient, error) {
	return &AuthorizeAllClient{}, nil
}

type AuthorizeAllClient struct {}

func (c *AuthorizeAllClient) Start() {
	// no-op
}

func (c *AuthorizeAllClient) Authorize(ctx context.Context, request *v2.CheckRequest) (*api.AuthorizationResponse, error) {
	return api.AuthorizedResponse(), nil
}






