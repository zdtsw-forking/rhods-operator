package deploy

import (
	"context"

	ofapi "github.com/operator-framework/api/pkg/operators/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// ManagedRhods defines expected addon catalogsource.
	ManagedRhods Platform = "addon-managed-odh-catalog"
	// SelfManagedRhods defines display name in csv.
	SelfManagedRhods Platform = "Red Hat OpenShift AI"
	// OpenDataHub defines display name in csv.
	OpenDataHub Platform = "Open Data Hub Operator"
	// Unknown indicates that operator is not deployed using OLM.
	Unknown Platform = ""
)

type Platform string

// isSelfManaged detects if it is Self Managed Rhods or OpenDataHub.
func isSelfManaged(cli client.Client) (Platform, error) {
	variants := map[string]Platform{
		"opendatahub-operator": OpenDataHub,
		"rhods-operator":       SelfManagedRhods,
	}

	for k, v := range variants {
		exists, err := OperatorExists(cli, k)
		if err != nil {
			return Unknown, err
		}
		if exists {
			return v, nil
		}
	}

	return Unknown, nil
}

// isManagedRHODS checks if catsrc CR add-on exists ManagedRhods.
func isManagedRHODS(ctx context.Context, cli client.Client) (Platform, error) {
	catalogSource := &ofapi.CatalogSource{}

	err := cli.Get(ctx, client.ObjectKey{Name: "addon-managed-odh-catalog", Namespace: "openshift-marketplace"}, catalogSource)
	if err != nil {
		return Unknown, client.IgnoreNotFound(err)
	}
	return ManagedRhods, nil
}

func GetPlatform(ctx context.Context, cli client.Client) (Platform, error) {
	// First check if its addon installation to return 'ManagedRhods, nil'
	if platform, err := isManagedRHODS(ctx, cli); err != nil {
		return Unknown, err
	} else if platform == ManagedRhods {
		return ManagedRhods, nil
	}

	// check and return whether ODH or self-managed platform
	return isSelfManaged(cli)
}
