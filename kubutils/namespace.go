package kubeutils

import (
	"context"
	"kubeutils/utils/logs"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Namespace struct {
	InstanceInterface typedv1.CoreV1Interface
	Item              *corev1.Namespace
}

func NewNamespace(kubeconfig string, item *corev1.Namespace) *Namespace {
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	resource := Namespace{}
	resource.InstanceInterface = instance.Clientset.CoreV1()
	resource.Item = item
	return &resource
}

func (c *Namespace) Create() error {
	logs.Info(map[string]interface{}{"命名空间": c.Item.Name}, "创建Namespace")
	_, err := c.InstanceInterface.Namespaces().Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *Namespace) Update() error {
	logs.Info(map[string]interface{}{"命名空间": c.Item.Name}, "更新Namespace")
	_, err := c.InstanceInterface.Namespaces().Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *Namespace) Delete(name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"命名空间": c.Item.Name}, "删除Namespace")
	deleteOptions := v1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOptions.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.Namespaces().Delete(context.TODO(), name, v1.DeleteOptions{})
	return err
}

func (c *Namespace) DeleteList(nameList []string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"命名空间": c.Item.Name}, "删除Namespace")
	for _, name := range nameList {
		c.Delete(name, gracePeriodSeconds)
	}
	return nil
}

func (c *Namespace) GET(name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"命名空间": name}, "获取Namespace")
	i, err := c.InstanceInterface.Namespaces().Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "v1"
	i.Kind = "Namespace"
	item = i
	return item, err
}

func (c *Namespace) List(labelSelector, fieldSelector string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{}, "获取Namespace列表")
	listOption := v1.ListOptions{
		LabelSelector: labelSelector,
		FieldSelector: fieldSelector,
	}
	i, err := c.InstanceInterface.Namespaces().List(context.TODO(), listOption)
	item = i.Items
	return item, err
}
