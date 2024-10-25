package kubeutils

import (
	"context"
	"kubeutils/utils/logs"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Service struct {
	InstenceInterface typev1.CoreV1Interface
	Item              *corev1.Service
}

func NewService(kubeconfig string, item *corev1.Service) *Service {
	instance := ResourceInstance{}
	instance.Init(kubeconfig)

	resource := Service{}
	resource.InstenceInterface = instance.Clientset.CoreV1()
	resource.Item = item
	return &resource
}

// 创建svc
func (c *Service) Create(namespace string) error {
	logs.Info(map[string]interface{}{"命名空间": namespace}, "创建svc")
	_, err := c.InstenceInterface.Services(namespace).Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

// 删除svc
func (c *Service) Delete(namespace string, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"命名空间": namespace, "svc名称": name}, "删除单个svc")
	deleteOptions := v1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOptions.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstenceInterface.Services(namespace).Delete(context.TODO(), name, deleteOptions)
	return err
}

func (c *Service) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	for _, name := range nameList {
		logs.Info(map[string]interface{}{"命名空间": namespace, "svc名称": name}, "删除svc数组")
		c.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (c *Service) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间：": namespace}, "更新Svc")
	_, err := c.InstenceInterface.Services(namespace).Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *Service) Get(namespace string, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间：": namespace}, "获取Svc")
	i, err := c.InstenceInterface.Services(namespace).Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "v1"
	i.Kind = "Service"
	item = i
	return item, err
}

func (c *Service) List(namespace string, labelSelector string, fieldSelector string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"命名空间：": namespace}, "获取Svc列表")
	listOptions := v1.ListOptions{
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	i, err := c.InstenceInterface.Services(namespace).List(context.TODO(), listOptions)
	item = i.Items
	return item, err
}
