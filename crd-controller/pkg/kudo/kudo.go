package kudo

import (
	"context"
	"fmt"
	"github.com/kudobuilder/kudo/pkg/kudoctl/resources/upgrade"
	"os"
	"reflect"

	"github.com/Masterminds/semver"
	"github.com/kudobuilder/kudo/pkg/apis/kudo/v1beta1"
	"github.com/kudobuilder/kudo/pkg/kudoctl/packages"
	"github.com/kudobuilder/kudo/pkg/kudoctl/packages/install"
	"github.com/kudobuilder/kudo/pkg/kudoctl/packages/resolver"
	"github.com/kudobuilder/kudo/pkg/kudoctl/util/kudo"
	"github.com/kudobuilder/kudo/pkg/kudoctl/util/repo"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/kudobuilder/shim/shim-controller/pkg/apis/kudoshim/v1alpha1"
	"github.com/kudobuilder/shim/crd-controller/pkg/client"
)

type KUDOClient struct {
	c *client.Client

	kc                *kudo.Client
	kudoPackageName   string
	version           string
	appVersion        string
	repoURL           string
	inClusterOperator bool

	resources *packages.Resources
	resolver  resolver.Resolver
}

func NewKUDOClient(k *client.Client, bi v1alpha1.ShimInstance) (*KUDOClient, error) {
	kc, err := kudo.NewClient(os.Getenv("KUBECONFIG"), 0, true)
	if err != nil {
		return nil, err
	}
	return &KUDOClient{
		c:                 k,
		kc:                kc,
		kudoPackageName:   bi.Spec.KUDOOperator.Package,
		version:           bi.Spec.KUDOOperator.Version,
		appVersion:        bi.Spec.KUDOOperator.AppVersion,
		repoURL:           bi.Spec.KUDOOperator.KUDORepository,
		inClusterOperator: bi.Spec.KUDOOperator.InClusterOperator,
	}, nil
}

func (k *KUDOClient) GetOVOrInstall(crd *unstructured.Unstructured) (*v1beta1.OperatorVersion, error) {
	if instance, err := k.kc.GetInstance(crd.GetName(), crd.GetNamespace()); err == nil && instance != nil {
		// already installed
		log.Infof("found the KUDO Instance %s/%s", crd.GetNamespace(), crd.GetName())
		log.Infof("fetching the KUDO Instance Operator %s/%s", instance.Spec.OperatorVersion.Name, instance.GetNamespace())
		return k.kc.GetOperatorVersion(instance.Spec.OperatorVersion.Name, instance.GetNamespace())
	}

	// install OV
	return k.InstallOV(crd)
}

func (k *KUDOClient) InstallOV(crd *unstructured.Unstructured) (*v1beta1.OperatorVersion, error) {
	var r resolver.Resolver
	if k.inClusterOperator {
		r = k.getClusterResolver(crd.GetNamespace())
	} else {

		repoConfig := repo.Configuration{
			URL:  k.repoURL,
			Name: "kudoShim",
		}

		repository, err := repo.NewClient(&repoConfig)
		if err != nil {
			return nil, err
		}

		r = resolver.New(repository)
	}
	p, err := r.Resolve(k.kudoPackageName, k.appVersion, k.version)
	if err != nil {
		return nil, err
	}

	k.resources = p.Resources
	k.resolver = r
	installOpts := install.Options{
		SkipInstance:    true,
		CreateNamespace: false,
	}

	parameters := make(map[string]string)
	if err := install.Package(k.kc, crd.GetName(), crd.GetNamespace(), *k.resources, parameters, r, installOpts); err != nil {
		return nil, err
	}
	return k.kc.GetOperatorVersion(k.resources.OperatorVersion.GetName(), k.resources.OperatorVersion.GetNamespace())
}

func (k *KUDOClient) InstallOrUpdateInstance(crd *unstructured.Unstructured, ov *v1beta1.OperatorVersion, params map[string]string) error {
	instance, err := k.kc.GetInstance(crd.GetName(), crd.GetNamespace())
	if instance == nil && err == nil {
		// install Instance
		return k.InstallInstance(crd, ov, params)
	}
	if err != nil {
		return err
	}
	// update existing instance
	return k.upgrade(instance, crd, ov, params)
}

func (k *KUDOClient) InstallInstance(crd *unstructured.Unstructured, ov *v1beta1.OperatorVersion, params map[string]string) error {
	installOpts := install.Options{
		SkipInstance:    false,
		CreateNamespace: false,
	}
	log.Infof("from here %s", crd.GetName())
	if err := install.Package(k.kc, crd.GetName(), crd.GetNamespace(), *k.resources, params, k.resolver, installOpts); err != nil {
		return err
	}
	return k.MarkOwnerReference(crd)

}

func (k *KUDOClient) upgrade(instance *v1beta1.Instance, crd *unstructured.Unstructured, ov *v1beta1.OperatorVersion, params map[string]string) error {
	oldOv, err := k.kc.GetOperatorVersion(instance.Spec.OperatorVersion.Name, instance.GetNamespace())
	if err != nil {
		return err
	}
	if oldOv == nil {
		return fmt.Errorf("no OperatorVersion installed for Instance %s/%s", instance.GetNamespace(), instance.GetName())
	}

	oldVersion, err := semver.NewVersion(oldOv.Spec.Version)
	if err != nil {
		return err
	}
	newVersion, err := semver.NewVersion(ov.Spec.Version)
	if err != nil {
		return err
	}

	if newVersion.Equal(oldVersion) {
		// update the instance values only if parameters are changed
		if !reflect.DeepEqual(instance.Spec.Parameters, params) {
			log.Infof("updating instance %s/%s parameters", instance.GetNamespace(), instance.GetName())
			log.Infof("old parameters: %+v", instance.Spec.Parameters)
			log.Infof("new parameters: %+v", params)
			return k.kc.UpdateInstance(instance.GetName(), instance.GetNamespace(), nil, params, nil, false, 0)
		}
		return nil
	}

	return upgrade.OperatorVersion(k.kc, ov, instance.GetName(), params, k.resolver)
}
func (k *KUDOClient) MarkOwnerReference(crd *unstructured.Unstructured) error {
	instance, err := k.kc.GetInstance(crd.GetName(), crd.GetNamespace())
	if err != nil {
		return err
	}
	instance.OwnerReferences = []metav1.OwnerReference{
		{
			APIVersion: crd.GetAPIVersion(),
			Kind:       crd.GetKind(),
			Name:       crd.GetName(),
			UID:        crd.GetUID(),
		},
	}
	_, err = k.c.KudoClient.KudoV1beta1().Instances(crd.GetNamespace()).Update(context.TODO(), instance, metav1.UpdateOptions{})
	return err
}

func (k *KUDOClient) getClusterResolver(ns string) InClusterResolver {
	return InClusterResolver{
		c:  k.c,
		ns: ns,
	}
}

type InClusterResolver struct {
	c  *client.Client
	ns string
}

func (r InClusterResolver) Resolve(name string, appVersion string, operatorVersion string) (*packages.Package, error) {
	ovn := v1beta1.OperatorVersionName(name, operatorVersion)

	ov, err := r.c.KudoClient.KudoV1beta1().OperatorVersions(r.ns).Get(context.TODO(), ovn, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to resolve operator version %s/%s:%s", r.ns, ovn, appVersion)
	}

	o, err := r.c.KudoClient.KudoV1beta1().Operators(r.ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to resolve operator %s/%s", r.ns, name)
	}

	return &packages.Package{
		Resources: &packages.Resources{
			Operator:        o,
			OperatorVersion: ov,
			Instance: &v1beta1.Instance{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Instance",
					APIVersion: packages.APIVersion,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:   v1beta1.OperatorInstanceName(o.Name),
					Labels: map[string]string{"kudo.dev/operator": o.Name},
				},
				Spec: v1beta1.InstanceSpec{
					OperatorVersion: corev1.ObjectReference{
						Name: v1beta1.OperatorVersionName(o.Name, ov.Spec.Version),
					},
				},
			},
		},
		Files: nil,
	}, nil
}
