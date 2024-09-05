// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package consul

import (
	"github.com/hashicorp/consul/internal/hcp"
	"github.com/hashicorp/consul/internal/multicluster"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/internal/resource/demo"
)

// NewTypeRegistry returns a registry populated with all supported resource
// types.
//
// Note: the registry includes resource types that may not be suitable for
// production use (e.g. experimental or development resource types) because
// it is used in the CLI, where feature flags and other runtime configuration
// may not be available.
func NewTypeRegistry() resource.Registry {
	registry := resource.NewRegistry()

	demo.RegisterTypes(registry)
	multicluster.RegisterTypes(registry)
	hcp.RegisterTypes(registry)

	return registry
}
