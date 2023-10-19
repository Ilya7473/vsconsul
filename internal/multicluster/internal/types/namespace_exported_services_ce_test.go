// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent
// +build !consulent

package types

import (
	multiclusterv1alpha1 "github.com/hashicorp/consul/proto-public/pbmulticluster/v1alpha1"
)

func validNamespaceExportedServicesWithPartition() *multiclusterv1alpha1.NamespaceExportedServices {
	consumers := []*multiclusterv1alpha1.ExportedServicesConsumer{
		{
			ConsumerTenancy: &multiclusterv1alpha1.ExportedServicesConsumer_Partition{
				Partition: "",
			},
		},
	}
	return &multiclusterv1alpha1.NamespaceExportedServices{
		Consumers: consumers,
	}
}

func validNamespaceExportedServicesWithSamenessGroup() *multiclusterv1alpha1.NamespaceExportedServices {
	consumers := []*multiclusterv1alpha1.ExportedServicesConsumer{
		{
			ConsumerTenancy: &multiclusterv1alpha1.ExportedServicesConsumer_SamenessGroup{
				SamenessGroup: "",
			},
		},
	}
	return &multiclusterv1alpha1.NamespaceExportedServices{
		Consumers: consumers,
	}
}
