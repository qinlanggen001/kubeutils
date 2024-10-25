package kubeutils

import (
	"context"
	"kubeutils/utils/logs"

	rbacv1 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type ClusterRole struct {
	InstanceInterface typedv1.RbacV1Interface
	Item              *rbacv1.ClusterRole
}

// NEW 函数用于配置一些默认值
func NewClusterRole(kubeconfig string, item *rbacv1.ClusterRole) *ClusterRole {
	// 首先调用instance的init函数，生成一个resourceInstance的实例，并配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	// 定义一个ClusterRole 实例
	resource := ClusterRole{}
	resource.InstanceInterface = instance.Clientset.RbacV1()
	resource.Item = item
	return &resource
}

func (c *ClusterRole) Create() error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "创建ClusterRole")
	_, err := c.InstanceInterface.ClusterRoles().Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *ClusterRole) Update() error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "更新ClusterRole")
	_, err := c.InstanceInterface.ClusterRoles().Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *ClusterRole) Get(name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "查询ClusterRole")
	i, err := c.InstanceInterface.ClusterRoles().Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "v1"
	i.Kind = "ClusterRole"
	item = i
	return item, err
}

func (c *ClusterRole) List(labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "查询ClusterRole列表")
	listOptions := v1.ListOptions{
		LabelSelector: labelSelector,
		FieldSelector: fieldSelector,
	}
	i, err := c.InstanceInterface.ClusterRoles().List(context.TODO(), listOptions)
	items = i.Items
	return items, err
}

func (c *ClusterRole) Delete(name string, gradePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "删除ClusterRole")
	deleteOption := v1.DeleteOptions{}
	if gradePeriodSeconds != nil {
		deleteOption.GracePeriodSeconds = gradePeriodSeconds
	}
	err := c.InstanceInterface.ClusterRoles().Delete(context.TODO(), name, deleteOption)
	return err
}

func (c *ClusterRole) DeleteList(nameList []string, gradePeriodSeconds *int64) error {
	for _, name := range nameList {
		c.Delete(name, gradePeriodSeconds)
	}
	return nil
}
