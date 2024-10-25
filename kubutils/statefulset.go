package kubeutils

import (
	"context"
	"kubeutils/utils/logs"

	corev1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typev1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type StatefulSet struct {
	InstanceInterface typev1.AppsV1Interface
	Item              *corev1.StatefulSet
}

func NewStatefulSet(kubeconfig string, item *corev1.StatefulSet) *StatefulSet {
	instance := ResourceInstance{}
	instance.Init(kubeconfig)

	resource := StatefulSet{}
	resource.InstanceInterface = instance.Clientset.AppsV1()
	resource.Item = item
	return &resource
}

func (c *StatefulSet) Create(namespace string) error {
	logs.Info(map[string]interface{}{}, "创建StatefulSet")
	_, err := c.InstanceInterface.StatefulSets(namespace).Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *StatefulSet) Update(namespace string) error {
	_, err := c.InstanceInterface.StatefulSets(namespace).Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *StatefulSet) Delete(namespace, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "删除StatefulSet")
	deleteOptions := v1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOptions.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.StatefulSets(namespace).Delete(context.TODO(), name, deleteOptions)
	return err
}

func (c *StatefulSet) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"命名空间": namespace}, "删除StatefulSet列表")
	for _, name := range nameList {
		c.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (c *StatefulSet) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "获取StatefulSet")
	i, err := c.InstanceInterface.StatefulSets(namespace).Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "apps/v1"
	i.Kind = "StatefulSet"
	item = i
	return item, err
}

func (c *StatefulSet) List(namespace, labelSelector, fieldSelector string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{}, "获取StatefulSet列表")
	listOptions := v1.ListOptions{
		LabelSelector: labelSelector,
		FieldSelector: fieldSelector,
	}
	i, err := c.InstanceInterface.StatefulSets(namespace).List(context.TODO(), listOptions)
	return i.Items, err
}
