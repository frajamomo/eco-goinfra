package far

import (
	"context"
	"fmt"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/golang/glog"
	farAlpha1 "github.com/medik8s/fence-agents-remediation/api/v1alpha1"
	"github.com/openshift-kni/eco-goinfra/pkg/clients"
	"github.com/openshift-kni/eco-goinfra/pkg/msg"
)

type FenceAgentsRemediationBuilder struct {
	// FenceAgentsRemediatonBuilder definition. Used to create
	// FenceAgentsRemediatonBuilder object with minimun set of required elements
	Definition *farAlpha1.FenceAgentsRemediation
	// Created FenceAgentsRemediationBuilder object on the cluster.
	Object *farAlpha1.FenceAgentsRemediation
	// api client to interact with the cluster
	apiClient *clients.Settings
	// errorMsg is processed before FenceAgentsRemediationBuilder object is created.
	errorMsg string
}

// PullFenceAgentsRemediation loads an existing fenceagentsremediation into Builder struct.
func PullFenceAgentsRemediation(apiClient *clients.Settings, name, namespace string) (*FenceAgentsRemediationBuilder,
	error) {
	glog.V(100).Infof("Pulling existing Fence Agents Remediation name %s in namespace %s", name, namespace)

	builder := FenceAgentsRemediationBuilder{
		apiClient: apiClient,
		Definition: &farAlpha1.FenceAgentsRemediation{
			ObjectMeta: metaV1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
		},
	}

	if name == "" {
		builder.errorMsg = "fence agents remediation 'name' cannot be empty"
	}

	if namespace == "" {
		builder.errorMsg = "fence agents remediation 'namespace' cannot be empty"
	}

	if !builder.Exists() {
		return nil, fmt.Errorf("fence agents remediation object %s doesn't exist in namespace %s", name, namespace)
	}

	builder.Definition = builder.Object

	return &builder, nil
}

// Exists checks whether the given fenceagentsremediation exists.
func (builder *FenceAgentsRemediationBuilder) Exists() bool {
	if valid, _ := builder.validate(); !valid {
		return false
	}

	glog.V(100).Infof(
		"Checking if fenceagentsremediation %s exists",
		builder.Definition.Name)

	var err error
	builder.Object, err = builder.apiClient.OperatorsV1alpha1Interface.FenceAgentsRemediations(
		builder.Definition.Namespace).Get(
		context.Background(), builder.Definition.Name, metaV1.GetOptions{})

	return err == nil || !k8serrors.IsNotFound(err)
}

// Delete removes a fenceagentsremediation
func (builder *FenceAgentsRemediationBuilder) Delete() error {
	if valid, err := builder.validate(); !valid {
		return err
	}

	glog.V(100).Infof("Deleting fenceagentsremediation %s in namespace %s", builder.Definition.Name,
		builder.Definition.Namespace)

	if !builder.Exists() {
		return nil
	}

	err := builder.apiClient.FenceAgentsRemediations(builder.Definition.Namespace).Delete(context.TODO(),
		builder.Object.Name, metaV1.DeleteOptions{})

	if err != nil {
		return err
	}

	builder.Object = nil

	return err
}

// validate will check that the builder and builder definition are properly initialized before
// accessing any member fields.
func (builder *FenceAgentsRemediationBuilder) validate() (bool, error) {
	resourceCRD := "FenceAgentsRemediation"

	if builder == nil {
		glog.V(100).Infof("The %s builder is uninitialized", resourceCRD)

		return false, fmt.Errorf("error: received nil %s builder", resourceCRD)
	}

	if builder.Definition == nil {
		glog.V(100).Infof("The %s is undefined", resourceCRD)

		builder.errorMsg = msg.UndefinedCrdObjectErrString(resourceCRD)
	}

	if builder.apiClient == nil {
		glog.V(100).Infof("The %s builder apiclient is nil", resourceCRD)

		builder.errorMsg = fmt.Sprintf("%s builder cannot have nil apiClient", resourceCRD)
	}

	if builder.errorMsg != "" {
		glog.V(100).Infof("The %s builder has error message: %s", resourceCRD, builder.errorMsg)

		return false, fmt.Errorf(builder.errorMsg)
	}

	return true, nil
}
