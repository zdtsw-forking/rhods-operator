package cluster_test

import (
	"context"
	"testing"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/cluster"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/cluster/gvk"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/utils/test/fakeclient"

	. "github.com/onsi/gomega"
)

func TestHasCRDWithVersion(t *testing.T) {
	ctx := context.Background()

	t.Run("should succeed if version is present", func(t *testing.T) {
		g := NewWithT(t)

		cli, err := fakeclient.New()
		g.Expect(err).ShouldNot(HaveOccurred())

		crd := apiextensionsv1.CustomResourceDefinition{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dashboards.components.platform.opendatahub.io",
			},
			Status: apiextensionsv1.CustomResourceDefinitionStatus{
				StoredVersions: []string{gvk.Dashboard.Version},
			},
		}

		err = cli.Create(ctx, &crd)
		g.Expect(err).ShouldNot(HaveOccurred())

		hasCRD, err := cluster.HasCRDWithVersion(ctx, cli, gvk.Dashboard.GroupKind(), gvk.Dashboard.Version)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(hasCRD).Should(BeTrue())
	})

	t.Run("should fails if version is not present", func(t *testing.T) {
		g := NewWithT(t)

		cli, err := fakeclient.New()
		g.Expect(err).ShouldNot(HaveOccurred())

		crd := apiextensionsv1.CustomResourceDefinition{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dashboards.components.platform.opendatahub.io",
			},
			Status: apiextensionsv1.CustomResourceDefinitionStatus{
				StoredVersions: []string{"v1alpha2"},
			},
		}

		err = cli.Create(ctx, &crd)
		g.Expect(err).ShouldNot(HaveOccurred())

		hasCRD, err := cluster.HasCRDWithVersion(ctx, cli, gvk.Dashboard.GroupKind(), gvk.Dashboard.Version)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(hasCRD).Should(BeFalse())
	})

	t.Run("should fails if terminating", func(t *testing.T) {
		g := NewWithT(t)

		cli, err := fakeclient.New()
		g.Expect(err).ShouldNot(HaveOccurred())

		crd := apiextensionsv1.CustomResourceDefinition{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dashboards.components.platform.opendatahub.io",
			},
			Status: apiextensionsv1.CustomResourceDefinitionStatus{
				StoredVersions: []string{gvk.Dashboard.Version},
				Conditions: []apiextensionsv1.CustomResourceDefinitionCondition{{
					Type:   apiextensionsv1.Terminating,
					Status: apiextensionsv1.ConditionTrue,
				}},
			},
		}

		err = cli.Create(ctx, &crd)
		g.Expect(err).ShouldNot(HaveOccurred())

		hasCRD, err := cluster.HasCRDWithVersion(ctx, cli, gvk.Dashboard.GroupKind(), gvk.Dashboard.Version)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(hasCRD).Should(BeFalse())
	})
}
