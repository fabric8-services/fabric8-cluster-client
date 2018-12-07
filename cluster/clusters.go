// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cluster": clusters Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-cluster/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-cluster-client
// --pkg=cluster
// --version=v1.3.0

package cluster

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateClustersPayload is the clusters create action payload.
type CreateClustersPayload struct {
	Data *CreateClusterData `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// CreateClustersPath computes a request path to the create action of clusters.
func CreateClustersPath() string {

	return fmt.Sprintf("/api/clusters/")
}

// Add a cluster configuration
func (c *Client) CreateClusters(ctx context.Context, path string, payload *CreateClustersPayload) (*http.Response, error) {
	req, err := c.NewCreateClustersRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateClustersRequest create the request corresponding to the create action endpoint of the clusters resource.
func (c *Client) NewCreateClustersRequest(ctx context.Context, path string, payload *CreateClustersPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ShowClustersPath computes a request path to the show action of clusters.
func ShowClustersPath() string {

	return fmt.Sprintf("/api/clusters/")
}

// Get clusters configuration
func (c *Client) ShowClusters(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowClustersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowClustersRequest create the request corresponding to the show action endpoint of the clusters resource.
func (c *Client) NewShowClustersRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ShowAuthClientClustersPath computes a request path to the showAuthClient action of clusters.
func ShowAuthClientClustersPath() string {

	return fmt.Sprintf("/api/clusters/auth")
}

// Get full cluster configuration including Auth information
func (c *Client) ShowAuthClientClusters(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowAuthClientClustersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowAuthClientClustersRequest create the request corresponding to the showAuthClient action endpoint of the clusters resource.
func (c *Client) NewShowAuthClientClustersRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
