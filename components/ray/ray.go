// Package ray provides utility functions to config Ray as part of the stack which makes managing distributed compute infrastructure in the cloud easy and intuitive for Data Scientists
package ray

import (
	dsci "github.com/opendatahub-io/opendatahub-operator/v2/apis/dscinitialization/v1"
	"github.com/opendatahub-io/opendatahub-operator/v2/components"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/deploy"
	operatorv1 "github.com/openshift/api/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ComponentName = "ray"
	RayPath       = deploy.DefaultManifestPath + "/" + "ray/operator/base"
)

var imageParamMap = map[string]string{
	"odh-kuberay-operator-controller-image": "RELATED_IMAGE_ODH_KUBERAY_OPERATOR_CONTROLLER_IMAGE",
}

type Ray struct {
	components.Component `json:""`
}

func (d *Ray) SetImageParamsMap(imageMap map[string]string) map[string]string {
	imageParamMap = imageMap
	return imageParamMap
}

func (d *Ray) GetComponentName() string {
	return ComponentName
}

// Verifies that Ray implements ComponentInterface
var _ components.ComponentInterface = (*Ray)(nil)

func (d *Ray) ReconcileComponent(owner metav1.Object, cli client.Client, scheme *runtime.Scheme, managementState operatorv1.ManagementState, dscispec *dsci.DSCInitializationSpec) error {
	enabled := managementState == operatorv1.Managed
	monitoringEnabled := dscispec.Monitoring.ManagementState == operatorv1.Managed
	platform, err := deploy.GetPlatform(cli)
	if err != nil {
		return err
	}
	if enabled {
		if dscispec.DevFlags.ManifestsUri == "" {
			if err := deploy.ApplyImageParams(RayPath, imageParamMap); err != nil {
				return err
			}
		}
	}
	// Deploy Ray Operator
	if err := deploy.DeployManifestsFromPath(owner, cli, ComponentName,
		RayPath,
		dscispec.ApplicationsNamespace,
		scheme, enabled); err != nil {
		return err
	}

	// CloudService Monitoring handling
	if platform == deploy.ManagedRhods && monitoringEnabled {
		if err := deploy.DeployManifestsFromPath(owner, cli, ComponentName,
			deploy.DefaultManifestPath+"/monitoring/prometheus/components/"+ComponentName,
			dscispec.Monitoring.Namespace,
			scheme, monitoringEnabled); err != nil {
			return err
		}
	}
	return err

}

func (in *Ray) DeepCopyInto(out *Ray) {
	*out = *in
	out.Component = in.Component
}
