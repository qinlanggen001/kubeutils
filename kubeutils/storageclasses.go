package kubeutils

import (
	"context"

	"github.com/qinlanggen001/kubeutils/utils/logs"

	storagev1 "k8s.io/api/storage/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

// 定义结构体
type StorageClass struct {
	InstanceInterface typedv1.StorageV1Interface
	Item              *storagev1.StorageClass
}

// New函数可以用于配置一些默认值
func NewStorageClass(kubeconfig string, item *storagev1.StorageClass) *StorageClass {
	// 首先调用instance的init函数，生成一个ResourceInstance的实例，并配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	// 定义一个StorageClass实例
	resource := StorageClass{}
	resource.InstanceInterface = instance.Clientset.StorageV1()
	resource.Item = item
	return &resource
}

// 创建资源
func (c *StorageClass) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "创建StorageClass")
	_, err := c.InstanceInterface.StorageClasses().Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

// 删除资源
func (c *StorageClass) Delete(namespace, name string, gracePeriodSeconds *int64) error {
	logs.Warning(map[string]interface{}{"名称": name}, "删除StorageClass")
	deleteOption := v1.DeleteOptions{}
	// gracePeriodSeconds可配置，如果为0代表是强制删除
	if gracePeriodSeconds != nil {
		// 说明传递了gracePeriodSeconds
		deleteOption.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.StorageClasses().Delete(context.TODO(), name, deleteOption)
	return err
}

// 删除多个
func (c *StorageClass) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	// 删除多个时，结构体会接收一个nameList的切片，循环该切片，然后调用Delete函数即可
	for _, name := range nameList {
		// 调用删除函数
		c.Delete("", name, gracePeriodSeconds)
	}
	// 忽略错误
	return nil
}

// 更新资源
func (c *StorageClass) Update(namespace string) error {
	logs.Warning(map[string]interface{}{"名称": c.Item.Name}, "更新StorageClass")
	_, err := c.InstanceInterface.StorageClasses().Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

// 获取资源列表
func (c *StorageClass) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{}, "查询StorageClass列表")
	// 有可能是根据查询条件进行查询
	listOptions := v1.ListOptions{
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	list, err := c.InstanceInterface.StorageClasses().List(context.TODO(), listOptions)
	items = list.Items
	return items, err
}

// 获取资源配置
func (c *StorageClass) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name}, "查询StorageClass详情")
	i, err := c.InstanceInterface.StorageClasses().Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "storage.k8s.io/v1"
	i.Kind = "StorageClass"
	item = i
	return item, err
}
