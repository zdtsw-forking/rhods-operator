package gvr

import "k8s.io/apimachinery/pkg/runtime/schema"

var (
	KnativeServing = schema.GroupVersionResource{
		Group:    "operator.knative.dev",
		Version:  "v1beta1",
		Resource: "knativeservings",
	}

	OpenshiftIngress = schema.GroupVersionResource{
		Group:    "config.openshift.io",
		Version:  "v1",
		Resource: "ingresses",
	}

	ResourceTracker = schema.GroupVersionResource{
		Group:    "features.opendatahub.io",
		Version:  "v1",
		Resource: "featuretrackers",
	}

	SMCP = schema.GroupVersionResource{
		Group:    "maistra.io",
		Version:  "v2",
		Resource: "servicemeshcontrolplanes",
	}

	// general core API
	Secret = schema.GroupVersionResource{
		Resource: "Secret",
	}

	ConfigMap = schema.GroupVersionResource{
		Resource: "ConfigMap",
	}

	Service = schema.GroupVersionResource{
		Resource: "Service",
	}

	ServiceAccount = schema.GroupVersionResource{
		Resource: "ServiceAccount",
	}

	Deployment = schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "Deployment",
	}

	NetworkPolicy = schema.GroupVersionResource{
		Group:    "networking.k8s.io",
		Version:  "v1",
		Resource: "NetworkPolicy",
	}

	Role = schema.GroupVersionResource{
		Group:    "rbac.authorization.k8s.io",
		Version:  "v1",
		Resource: "featuretrackers",
	}

	RoleBinding = schema.GroupVersionResource{
		Group:    "rbac.authorization.k8s.io",
		Version:  "v1",
		Resource: "RoleBinding",
	}

	ClusterRole = schema.GroupVersionResource{
		Group:    "rbac.authorization.k8s.io",
		Version:  "v1",
		Resource: "ClusterRole",
	}

	ClusterRoleBinding = schema.GroupVersionResource{
		Group:    "rbac.authorization.k8s.io",
		Version:  "v1",
		Resource: "ClusterRoleBinding",
	}

	ImageStream = schema.GroupVersionResource{
		Group:    "image.openshift.io",
		Version:  "v1",
		Resource: "ImageStream",
	}

	MutatingWebhookConfiguration = schema.GroupVersionResource{
		Group:    "admissionregistration.k8s.io",
		Version:  "v1",
		Resource: "MutatingWebhookConfiguration",
	}

	ValidatingWebhookConfiguration = schema.GroupVersionResource{
		Group:    "admissionregistration.k8s.io",
		Version:  "v1",
		Resource: "ValidatingWebhookConfiguration",
	}
)

