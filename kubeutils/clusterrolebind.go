package kubeutils

import (
	"context"

	"github.com/qinlanggen001/kubeutils/utils/logs"

	rbacv1 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type ClusterRoleBind struct {
	InstanceInterface typedv1.RbacV1Interface
	Item              *rbacv1.ClusterRoleBinding
}

// NEW 函数用于配置一些默认值
func NewClusterRoleBind(kubeconfig string, item *rbacv1.ClusterRoleBinding) *ClusterRoleBind {
	// 首先调用instance的init函数，生成一个resourceInstance的实例，并配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	// 定义一个ClusterRoleBind 实例
	resource := ClusterRoleBind{}
	resource.InstanceInterface = instance.Clientset.RbacV1()
	resource.Item = item
	return &resource
}

func (c *ClusterRoleBind) Create() error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "创建ClusterRoleBind")
	_, err := c.InstanceInterface.ClusterRoleBindings().Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *ClusterRoleBind) Update() error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "更新ClusterRoleBind")
	_, err := c.InstanceInterface.ClusterRoleBindings().Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *ClusterRoleBind) Get(name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "查询ClusterRoleBind")
	i, err := c.InstanceInterface.ClusterRoleBindings().Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "v1"
	i.Kind = "ClusterRoleBind"
	item = i
	return item, err
}

func (c *ClusterRoleBind) List(labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "查询ClusterRoleBind列表")
	listOptions := v1.ListOptions{
		LabelSelector: labelSelector,
		FieldSelector: fieldSelector,
	}
	i, err := c.InstanceInterface.ClusterRoleBindings().List(context.TODO(), listOptions)
	items = i.Items
	return items, err
}

func (c *ClusterRoleBind) Delete(name string, gradePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "删除ClusterRoleBind")
	deleteOption := v1.DeleteOptions{}
	if gradePeriodSeconds != nil {
		deleteOption.GracePeriodSeconds = gradePeriodSeconds
	}
	err := c.InstanceInterface.ClusterRoleBindings().Delete(context.TODO(), name, deleteOption)
	return err
}

func (c *ClusterRoleBind) DeleteList(nameList []string, gradePeriodSeconds *int64) error {
	for _, name := range nameList {
		c.Delete(name, gradePeriodSeconds)
	}
	return nil
}
