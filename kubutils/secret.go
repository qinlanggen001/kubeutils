package kubeutils

import (
	"context"
	"kubeutils/utils/logs"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Secret struct {
	InstanceInterface typedv1.CoreV1Interface
	Item              *corev1.Secret
}

// NEW 函数用于配置一些默认值
func NewSecret(kubeconfig string, item *corev1.Secret) *Secret {
	// 首先调用instance的init函数，生成一个resourceInstance的实例，并配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	// 定义一个ClusterRole 实例
	resource := Secret{}
	resource.InstanceInterface = instance.Clientset.CoreV1()
	resource.Item = item
	return &resource
}

func (c *Secret) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "创建Secret")
	_, err := c.InstanceInterface.Secrets(namespace).Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *Secret) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "更新Secret")
	_, err := c.InstanceInterface.Secrets(namespace).Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *Secret) Get(namespace string, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "查询Secret")
	i, err := c.InstanceInterface.Secrets(namespace).Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "v1"
	i.Kind = "Secret"
	item = i
	return item, err
}

func (c *Secret) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "查询Secret列表")
	listOptions := v1.ListOptions{
		LabelSelector: labelSelector,
		FieldSelector: fieldSelector,
	}
	i, err := c.InstanceInterface.Secrets(namespace).List(context.TODO(), listOptions)
	items = i.Items
	return items, err
}

func (c *Secret) Delete(namespace, name string, gradePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "删除Secret")
	deleteOption := v1.DeleteOptions{}
	if gradePeriodSeconds != nil {
		deleteOption.GracePeriodSeconds = gradePeriodSeconds
	}
	err := c.InstanceInterface.Secrets(namespace).Delete(context.TODO(), name, deleteOption)
	return err
}

func (c *Secret) DeleteList(namespace string, nameList []string, gradePeriodSeconds *int64) error {
	for _, name := range nameList {
		c.Delete(namespace, name, gradePeriodSeconds)
	}
	return nil
}
