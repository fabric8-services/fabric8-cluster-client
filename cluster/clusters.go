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
	uuid "github.com/goadesign/goa/uuid"
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

// DeleteClustersPath computes a request path to the delete action of clusters.
func DeleteClustersPath(clusterID uuid.UUID) string {
	param0 := clusterID.String()

	return fmt.Sprintf("/api/clusters/%s", param0)
}

// Delete a cluster configuration
func (c *Client) DeleteClusters(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteClustersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteClustersRequest create the request corresponding to the delete action endpoint of the clusters resource.
func (c *Client) NewDeleteClustersRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// LinkIdentityToClusterClustersPath computes a request path to the linkIdentityToCluster action of clusters.
func LinkIdentityToClusterClustersPath() string {

	return fmt.Sprintf("/api/clusters/identities")
}

// create a identitycluster using a service account
func (c *Client) LinkIdentityToClusterClusters(ctx context.Context, path string, payload *LinkIdentityToClusterData) (*http.Response, error) {
	req, err := c.NewLinkIdentityToClusterClustersRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLinkIdentityToClusterClustersRequest create the request corresponding to the linkIdentityToCluster action endpoint of the clusters resource.
func (c *Client) NewLinkIdentityToClusterClustersRequest(ctx context.Context, path string, payload *LinkIdentityToClusterData) (*http.Request, error) {
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

// ListClustersPath computes a request path to the list action of clusters.
func ListClustersPath() string {

	return fmt.Sprintf("/api/clusters/")
}

// Get all cluster configurations. If the 'cluster-url' query parameter is set, then a single cluster is returned. If the 'type' query parameter is set then only the clusters with the matchin type are returned
func (c *Client) ListClusters(ctx context.Context, path string, clusterURL *string, type_ *string) (*http.Response, error) {
	req, err := c.NewListClustersRequest(ctx, path, clusterURL, type_)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListClustersRequest create the request corresponding to the list action endpoint of the clusters resource.
func (c *Client) NewListClustersRequest(ctx context.Context, path string, clusterURL *string, type_ *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if clusterURL != nil {
		values.Set("cluster-url", *clusterURL)
	}
	if type_ != nil {
		values.Set("type", *type_)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListForAuthClientClustersPath computes a request path to the listForAuthClient action of clusters.
func ListForAuthClientClustersPath() string {

	return fmt.Sprintf("/api/clusters/auth")
}

// Get all cluster configurations unless the 'cluster-url' is specified. This endpoint returns all sensitive information
func (c *Client) ListForAuthClientClusters(ctx context.Context, path string, clusterURL *string, type_ *string) (*http.Response, error) {
	req, err := c.NewListForAuthClientClustersRequest(ctx, path, clusterURL, type_)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListForAuthClientClustersRequest create the request corresponding to the listForAuthClient action endpoint of the clusters resource.
func (c *Client) NewListForAuthClientClustersRequest(ctx context.Context, path string, clusterURL *string, type_ *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if clusterURL != nil {
		values.Set("cluster-url", *clusterURL)
	}
	if type_ != nil {
		values.Set("type", *type_)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// RemoveIdentityToClusterLinkClustersPath computes a request path to the removeIdentityToClusterLink action of clusters.
func RemoveIdentityToClusterLinkClustersPath() string {

	return fmt.Sprintf("/api/clusters/identities")
}

// Remove a identity cluster relation using a service account
func (c *Client) RemoveIdentityToClusterLinkClusters(ctx context.Context, path string, payload *UnLinkIdentityToClusterdata) (*http.Response, error) {
	req, err := c.NewRemoveIdentityToClusterLinkClustersRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRemoveIdentityToClusterLinkClustersRequest create the request corresponding to the removeIdentityToClusterLink action endpoint of the clusters resource.
func (c *Client) NewRemoveIdentityToClusterLinkClustersRequest(ctx context.Context, path string, payload *UnLinkIdentityToClusterdata) (*http.Request, error) {
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
	req, err := http.NewRequest("DELETE", u.String(), &body)
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
func ShowClustersPath(clusterID uuid.UUID) string {
	param0 := clusterID.String()

	return fmt.Sprintf("/api/clusters/%s", param0)
}

// Get single cluster configuration
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

// ShowForAuthClientClustersPath computes a request path to the showForAuthClient action of clusters.
func ShowForAuthClientClustersPath(clusterID uuid.UUID) string {
	param0 := clusterID.String()

	return fmt.Sprintf("/api/clusters/%s/auth", param0)
}

// Get single cluster configuration (including Auth information)
func (c *Client) ShowForAuthClientClusters(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowForAuthClientClustersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowForAuthClientClustersRequest create the request corresponding to the showForAuthClient action endpoint of the clusters resource.
func (c *Client) NewShowForAuthClientClustersRequest(ctx context.Context, path string) (*http.Request, error) {
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
