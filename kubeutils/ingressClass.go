package kubeutils

import (
	"context"

	"github.com/qinlanggen001/kubeutils/utils/logs"

	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

type IngressClass struct {
	InstanceInterface typedv1.NetworkingV1Interface
	Item              *networkingv1.IngressClass
}

func NetIngressClass(kubeconfig string, item *networkingv1.IngressClass) *IngressClass {
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	resource := IngressClass{}
	resource.InstanceInterface = instance.Clientset.NetworkingV1()
	resource.Item = item
	return &resource
}

// 创建资源
func (c *IngressClass) Create() error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "创建IngressClasses")
	_, err := c.InstanceInterface.IngressClasses().Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

// 删除资源
func (c *IngressClass) Delete(name string, gracePeriodSeconds *int64) error {
	logs.Warning(map[string]interface{}{"名称": name}, "删除IngressClasses")
	deleteOption := v1.DeleteOptions{}
	// gracePeriodSeconds可配置，如果为0代表是强制删除
	if gracePeriodSeconds != nil {
		// 说明传递了gracePeriodSeconds
		deleteOption.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.IngressClasses().Delete(context.TODO(), name, deleteOption)
	return err
}

// 删除多个
func (c *IngressClass) DeleteList(nameList []string, gracePeriodSeconds *int64) error {
	// 删除多个时，结构体会接收一个nameList的切片，循环该切片，然后调用Delete函数即可
	for _, name := range nameList {
		// 调用删除函数
		c.Delete(name, gracePeriodSeconds)
	}
	// 忽略错误
	return nil
}

// 更新资源
func (c *IngressClass) Update() error {
	logs.Warning(map[string]interface{}{"名称": c.Item.Name}, "更新IngressClasses")
	_, err := c.InstanceInterface.IngressClasses().Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

// 获取资源列表
func (c *IngressClass) List(labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(nil, "查询Ingress列表")
	// 有可能是根据查询条件进行查询
	listOptions := v1.ListOptions{
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	list, err := c.InstanceInterface.IngressClasses().List(context.TODO(), listOptions)
	items = list.Items
	return items, err
}

// 获取资源配置
func (c *IngressClass) Get(name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name}, "查询IngressClasses详情")
	i, err := c.InstanceInterface.IngressClasses().Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "networking.k8s.io/v1"
	i.Kind = "IngressClass"
	item = i
	return item, err
}
