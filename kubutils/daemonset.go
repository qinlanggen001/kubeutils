package kubeutils

import (
	"context"
	"kubeutils/utils/logs"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type DaemonSet struct {
	InstanceInterface typedv1.AppsV1Interface
	Item              *appsv1.DaemonSet
}

func NewDaemonSet(kubeconfig string, item *appsv1.DaemonSet) *DaemonSet {
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	resource := DaemonSet{}
	resource.InstanceInterface = instance.Clientset.AppsV1()
	resource.Item = item
	return &resource
}

func (c *DaemonSet) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "获取daemonset")
	i, err := c.InstanceInterface.DaemonSets(namespace).Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "apps/v1"
	i.Kind = "DaemonSet"
	item = i
	return item, err
}

func (c *DaemonSet) List(namespace, labelSelector, fieldSelector string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "获取daemonset列表")
	listOptions := v1.ListOptions{
		LabelSelector: labelSelector,
		FieldSelector: fieldSelector,
	}
	i, err := c.InstanceInterface.DaemonSets(namespace).List(context.TODO(), listOptions)
	return i.Items, err
}

func (c *DaemonSet) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "创建daemonset")
	_, err := c.InstanceInterface.DaemonSets(namespace).Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *DaemonSet) Delete(namespace, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "删除daemonset")
	deleteOptions := v1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOptions.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.DaemonSets(namespace).Delete(context.TODO(), name, v1.DeleteOptions{})
	return err
}

func (c *DaemonSet) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "删除daemonset列表")
	for _, name := range nameList {
		c.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (c *DaemonSet) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "更新daemonset列表")
	_, err := c.InstanceInterface.DaemonSets(namespace).Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}
