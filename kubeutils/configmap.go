package kubeutils

import (
	"context"

	"github.com/qinlanggen001/kubeutils/utils/logs"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type ConfigMap struct {
	InstanceInterface typedv1.CoreV1Interface
	Item              *corev1.ConfigMap
}

// NEW 函数用于配置一些默认值
func NewConfigMap(kubeconfig string, item *corev1.ConfigMap) *ConfigMap {
	// 首先调用instance的init函数，生成一个resourceInstance的实例，并配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	// 定义一个ClusterRole 实例
	resource := ConfigMap{}
	resource.InstanceInterface = instance.Clientset.CoreV1()
	resource.Item = item
	return &resource
}

func (c *ConfigMap) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "创建ConfigMap")
	_, err := c.InstanceInterface.ConfigMaps(namespace).Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *ConfigMap) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "更新ConfigMap")
	_, err := c.InstanceInterface.ConfigMaps(namespace).Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *ConfigMap) Get(namespace string, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "查询ConfigMap")
	i, err := c.InstanceInterface.ConfigMaps(namespace).Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "v1"
	i.Kind = "ConfigMap"
	item = i
	return item, err
}

func (c *ConfigMap) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "查询ConfigMap列表")
	listOptions := v1.ListOptions{
		LabelSelector: labelSelector,
		FieldSelector: fieldSelector,
	}
	i, err := c.InstanceInterface.ConfigMaps(namespace).List(context.TODO(), listOptions)
	items = i.Items
	return items, err
}

func (c *ConfigMap) Delete(namespace, name string, gradePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "删除ConfigMap")
	deleteOption := v1.DeleteOptions{}
	if gradePeriodSeconds != nil {
		deleteOption.GracePeriodSeconds = gradePeriodSeconds
	}
	err := c.InstanceInterface.ConfigMaps(namespace).Delete(context.TODO(), name, deleteOption)
	return err
}

func (c *ConfigMap) DeleteList(namespace string, nameList []string, gradePeriodSeconds *int64) error {
	for _, name := range nameList {
		c.Delete(namespace, name, gradePeriodSeconds)
	}
	return nil
}
