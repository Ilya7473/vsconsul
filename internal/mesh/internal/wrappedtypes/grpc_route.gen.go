// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package wrappedtypes

// Code generated by wrapped-protobuf-types/generate.sh. DO NOT EDIT.

import (
	"github.com/hashicorp/consul/internal/resource"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v1alpha1"
	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v1alpha1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

// Avoid unused imports in generated code.
var _ *pbmesh.ParentReference
var _ *pbcatalog.Service

var _ WrappedRoute = (*GRPCRoute)(nil)

type GRPCRoute struct {
	Resource *pbresource.Resource
	*pbmesh.GRPCRoute
}

func NewGRPCRoute(dec *resource.DecodedResource[pbmesh.GRPCRoute, *pbmesh.GRPCRoute]) *GRPCRoute {
	if dec == nil {
		return nil
	}
	return &GRPCRoute{
		Resource:  dec.Resource,
		GRPCRoute: dec.Data,
	}
}

func (r *GRPCRoute) GetResource() *pbresource.Resource { return r.Resource }

func (r *GRPCRoute) ToDecodedResource() *resource.DecodedResource[pbmesh.GRPCRoute, *pbmesh.GRPCRoute] {
	if r == nil {
		return nil
	}
	return &resource.DecodedResource[pbmesh.GRPCRoute, *pbmesh.GRPCRoute]{
		Resource: r.Resource,
		Data:     r.GRPCRoute,
	}
}
