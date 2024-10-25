package kubeutils

import (
	"context"
	"kubeutils/utils/logs"

	rbacv1 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typev1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type Role struct {
	InstanceInterface typev1.RbacV1Interface
	Item              *rbacv1.Role
}

func NewRole(kubeconfig string, item *rbacv1.Role) *Role {
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	//定义一个role实例
	resource := Role{}
	resource.InstanceInterface = instance.Clientset.RbacV1()
	resource.Item = item
	return &resource
}

func (c *Role) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "创建role")
	_, err := c.InstanceInterface.Roles(namespace).Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *Role) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "更新role")
	_, err := c.InstanceInterface.Roles(namespace).Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *Role) Get(namespace string, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "获取role")
	i, err := c.InstanceInterface.Roles(namespace).Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "v1"
	i.Kind = "Role"
	item = i
	return item, err
}

func (c *Role) List(namespace string, labelSelector string, fieldSelector string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"命名空间": namespace}, "获取role")
	listOptions := v1.ListOptions{
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	i, err := c.InstanceInterface.Roles(namespace).List(context.TODO(), listOptions)
	item = i.Items
	return item, err
}

func (c *Role) Delete(namespace string, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "获取role")
	deleteOptions := v1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOptions.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.Roles(namespace).Delete(context.TODO(), name, deleteOptions)
	return err
}

func (c *Role) DeleteCollection(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	for _, name := range nameList {
		logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "获取role")
		c.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}
