package fenceAgentsRemediation

import (
	"context"
	"fmt"

	fenceAgentsRemediation "github.com/medik8s/fence-agents-remediation/api/v1alpha1"
	"github.com/golang/glog"
	"github.com/openshift-kni/eco-goinfra/pkg/clients"
	"github.com/openshift-kni/eco-goinfra/pkg/msg"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	goclient "sigs.k8s.io/controller-runtime/pkg/client"
}

// Builder provides a struct for ClusterPolicy object
// from the cluster and a ClusterPolicy definition.
type FenceAgentsRemediationBuilder struct {
	// Builder definition. Used to create
	// Builder object with minimum set of required elements.
	Definition *fenceAgfenceAgentsRemediation.FenceAgentsRemediation
	// Created Builder object on the cluster.
	Object *fenceAfenceAgentsRemediation.FenceAgentsRemediation
	// api client to interact with the cluster.
	apiClient *clients.Settings
	// errorMsg is processed before Builder object is created.
	errorMsg string
}


func NewFenceAgentsRemediationBuilder(
	apiClient *clients.Settings, name) *FenceAgentsRemediationBuilder{
	glog.V(100).Infof(
		"Initializing new FenceAgentsRemediation structure with the following params: "+
			"name: %s", name)

	builder := FenceAgentsRemediationBuilder{
		apiClient: apiClient,
		Definition: &fenceAgentsRemediation.FenceAgentsRemediation{
			ObjectMeta: metaV1.ObjectMeta{
				Name: name,
			}
		},
	}
}

