// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package deploy

import (
	"context"
	"strings"

	"github.com/elastic/e2e-testing/internal/shell"
	"go.elastic.co/apm"
)

// remoteDeploymentManifest deploy manifest for docker
type remoteDeploymentManifest struct {
	Context context.Context
}

func newRemoteDeploy() Deployment {
	return &remoteDeploymentManifest{Context: context.Background()}
}

// Add - stub for remote deployment
func (c *remoteDeploymentManifest) Add(ctx context.Context, profile ServiceRequest, services []ServiceRequest, env map[string]string) error {
	return nil
}

// Bootstrap - stub for remote deployment
func (c *remoteDeploymentManifest) Bootstrap(ctx context.Context, profile ServiceRequest, env map[string]string, waitCB func() error) error {
	return nil
}

// AddFiles - add files to service
func (c *remoteDeploymentManifest) AddFiles(ctx context.Context, profile ServiceRequest, service ServiceRequest, files []string) error {
	return nil
}

// Destroy teardown environment
func (c *remoteDeploymentManifest) Destroy(ctx context.Context, profile ServiceRequest) error {
	return nil
}

// ExecIn execute command in service
func (c *remoteDeploymentManifest) ExecIn(ctx context.Context, profile ServiceRequest, service ServiceRequest, cmd []string) (string, error) {
	span, _ := apm.StartSpanOptions(ctx, "Executing command in remote deployment", "remote.manifest.execIn", apm.SpanOptions{
		Parent: apm.SpanFromContext(ctx).TraceContext(),
	})
	span.Context.SetLabel("profile", profile)
	span.Context.SetLabel("service", service)
	span.Context.SetLabel("arguments", cmd)
	defer span.End()

	output, err := shell.Execute(ctx, ".", cmd[0], cmd[1:]...)
	if err != nil {
		return "", err
	}
	return output, nil
}

// Inspect inspects a service
func (c *remoteDeploymentManifest) Inspect(ctx context.Context, service ServiceRequest) (*ServiceManifest, error) {
	// TODO: convert to a platform agnostic command structure
	hostname, _ := shell.Execute(ctx, ".", "powershell.exe", "hostname")
	return &ServiceManifest{
		Hostname:   strings.TrimSpace(hostname),
		Connection: service.Name,
		Alias:      service.Name,
		Platform:   "windows",
	}, nil
}

// Logs print logs of service
func (c *remoteDeploymentManifest) Logs(ctx context.Context, service ServiceRequest) error {
	return nil
}

// PreBootstrap sets up environment
func (c *remoteDeploymentManifest) PreBootstrap(ctx context.Context) error {
	return nil
}

// Remove remove services from deployment
func (c *remoteDeploymentManifest) Remove(ctx context.Context, profile ServiceRequest, services []ServiceRequest, env map[string]string) error {
	return nil
}

// Start a container
func (c *remoteDeploymentManifest) Start(ctx context.Context, service ServiceRequest) error {
	return nil
}

// Stop a container
func (c *remoteDeploymentManifest) Stop(ctx context.Context, service ServiceRequest) error {
	return nil
}
