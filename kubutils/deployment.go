package kubeutils

import (
	"context"
	"kubeutils/utils/logs"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type Deployment struct {
	InstanceInterface typedv1.AppsV1Interface
	Item              *appsv1.Deployment
}

func NewDeployment(kubeconfig string, item *appsv1.Deployment) *Deployment {
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	resource := Deployment{}
	resource.InstanceInterface = instance.Clientset.AppsV1()
	// 将item转换为具体的类型
	resource.Item = item
	return &resource
}

// 创建deployment
func (c *Deployment) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "创建Deployment")
	_, err := c.InstanceInterface.Deployments(namespace).Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *Deployment) Delete(namespace, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "删除Deployment")
	deleteOptions := v1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOptions.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.Deployments(namespace).Delete(context.TODO(), name, deleteOptions)
	return err
}

func (c *Deployment) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"命名空间": namespace}, "删除Deployment列表")
	for _, name := range nameList {
		c.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (c *Deployment) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "获取Deployment")
	_, err := c.InstanceInterface.Deployments(namespace).Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *Deployment) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "获取Deployment")
	i, err := c.InstanceInterface.Deployments(namespace).Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "apps/v1"
	i.Kind = "Deployment"
	item = i
	return item, err
}

func (c *Deployment) List(namespace, labelSelector, fieldSelector string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{}, "获取deployment列表")
	listOptions := v1.ListOptions{
		LabelSelector: labelSelector,
		FieldSelector: fieldSelector,
	}
	i, err := c.InstanceInterface.Deployments(namespace).List(context.TODO(), listOptions)
	return i.Items, err
}
