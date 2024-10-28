package kubeutils

import (
	"context"

	"github.com/qinlanggen001/kubeutils/utils/logs"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Node struct {
	InstanceInterface typedv1.CoreV1Interface
	Item              *corev1.Node
}

func NewNode(kubeconfig string, item *corev1.Node) *Node {
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	resource := Node{}
	resource.InstanceInterface = instance.Clientset.CoreV1()
	resource.Item = item
	return &resource
}

func (c *Node) Create(namespace string) error {
	logs.Info(map[string]interface{}{"命名空间": c.Item.Name}, "创建Node")
	_, err := c.InstanceInterface.Nodes().Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *Node) Update(namespace string) error {
	logs.Info(map[string]interface{}{"命名空间": c.Item.Name}, "更新Node")
	_, err := c.InstanceInterface.Nodes().Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *Node) Delete(namespace string, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"命名空间": c.Item.Name}, "删除Node")
	deleteOptions := v1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOptions.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.Nodes().Delete(context.TODO(), name, v1.DeleteOptions{})
	return err
}

func (c *Node) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"命名空间": c.Item.Name}, "删除Node")
	for _, name := range nameList {
		c.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (c *Node) Get(namespace string, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"命名空间": name}, "获取Node")
	i, err := c.InstanceInterface.Nodes().Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "v1"
	i.Kind = "Node"
	item = i
	return item, err
}

func (c *Node) List(namespace string, labelSelector, fieldSelector string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{}, "获取Node列表")
	listOption := v1.ListOptions{
		LabelSelector: labelSelector,
		FieldSelector: fieldSelector,
	}
	i, err := c.InstanceInterface.Nodes().List(context.TODO(), listOption)
	item = i.Items
	return item, err
}
