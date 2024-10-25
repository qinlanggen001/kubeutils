package kubeutils

import (
	"context"
	"kubeutils/utils/logs"

	rbacv1 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typev1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type RoleBinding struct {
	InstanceInterface typev1.RbacV1Interface
	Item              *rbacv1.RoleBinding
}

func NewRoleBinding(kubeconfig string, item *rbacv1.RoleBinding) *RoleBinding {
	instance := ResourceInstance{}
	instance.Init(kubeconfig)

	resource := RoleBinding{}
	resource.InstanceInterface = instance.Clientset.RbacV1()
	resource.Item = item
	return &resource
}

func (c *RoleBinding) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "创建rolebinding")
	_, err := c.InstanceInterface.RoleBindings(namespace).Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *RoleBinding) Delete(namespace string, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "删除rolebinding")
	deleteoptions := v1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteoptions.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.RoleBindings(namespace).Delete(context.TODO(), name, deleteoptions)
	return err
}

func (c *RoleBinding) DeleteCollection(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	for _, name := range nameList {
		logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "删除rolebinding数组")
		c.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (c *RoleBinding) Get(namespace string, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "查找rolebinding")
	i, err := c.InstanceInterface.RoleBindings(namespace).Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "v1"
	i.Kind = "RoleBinding"
	item = i
	return item, err
}

func (c *RoleBinding) List(namespace string, labelSelector string, fieldSelectot string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"命名空间": namespace}, "查找rolebinding")
	listOptions := v1.ListOptions{
		LabelSelector: labelSelector,
		FieldSelector: fieldSelectot,
	}
	i, err := c.InstanceInterface.RoleBindings(namespace).List(context.TODO(), listOptions)
	item = i.Items
	return item, err
}
