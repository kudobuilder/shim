package watcher

import (
	"context"
	"errors"
	"fmt"
	"github.com/kudobuilder/shim/crd-controller/pkg/utils"

	"github.com/devopsfaith/flatmap"
	"github.com/kudobuilder/kudo/pkg/apis/kudo/v1beta1"
	log "github.com/sirupsen/logrus"
	"github.com/kudobuilder/shim/crd-controller/pkg/client"
	"github.com/kudobuilder/shim/crd-controller/pkg/kudo"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func Process(client *client.Client, item runtime.Object) error {
	if item == nil {
		// Event was deleted
		return nil
	}

	crd, ok := item.(*unstructured.Unstructured)
	if !ok {
		return errors.New("the CRD doesn't have unstructured.Unstructured spec")
	}
	labelSelector := fmt.Sprintf("%s=%s,%s=%s,%s=%s",
		"version", crd.GroupVersionKind().Version, "kind", crd.GroupVersionKind().Kind, "group", crd.GroupVersionKind().Group)

	//find shim instance for the current CRD
	shimInstanceList, err := client.Shim.KudoshimV1alpha1().ShimInstances(crd.GetNamespace()).List(context.TODO(), v1.ListOptions{
		LabelSelector: labelSelector,
	})

	if err != nil {
		log.Errorf("error retrieving KUDO Shim Instances for %s : %v", labelSelector, err)
		return err
	}

	if len(shimInstanceList.Items) != 1 {
		log.Errorf("Expecting 1 Shim Instance but found %d", len(shimInstanceList.Items))
		return err
	}

	kc, err := kudo.NewKUDOClient(client, shimInstanceList.Items[0])
	if err != nil {
		log.Errorf("Error initializing KUDO Client :%v", err)
		return err
	}

	log.Infof("checking if the KUDO Instance %s/%s is already installed", crd.GetNamespace(), crd.GetName())
	// get the operatorversion using shimInstance reference
	ov, err := kc.GetOVOrInstall(crd)
	if err != nil {
		log.Errorf("Error initializing OV :%v", err)
		return err
	}

	crdFlatMap, _ := utils.Flatten(crd.UnstructuredContent(), flatmap.DefaultTokenizer)
	shimInstanceFlatMap, _ := utils.Flatten(shimInstanceList.Items[0].Spec.CRDSpec.UnstructuredContent(), flatmap.DefaultTokenizer)
	ovParamsMap, _ := getParamsMapFromOV(ov.Spec.Parameters)
	instanceParamsToUpdate := make(map[string]string)
	for key, val := range shimInstanceFlatMap.M {
		if _, exists := ovParamsMap[fmt.Sprintf("%v", val)]; exists {
			if crdVal, ok := crdFlatMap.M[key]; ok {
				instanceParamsToUpdate[val.(string)] = fmt.Sprintf("%v", crdVal)
			}
		}
	}
	// OV is already installed
	// Install Instance or Update/Upgrade the instance
	return kc.InstallOrUpdateInstance(crd, ov, instanceParamsToUpdate)

}

func getParamsMapFromOV(parameters []v1beta1.Parameter) (map[string]bool, error) {
	paramMap := make(map[string]bool)
	for _, val := range parameters {
		paramMap[val.Name] = true
	}
	return paramMap, nil
}
