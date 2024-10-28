package kubeutils

import (
	"context"

	"github.com/qinlanggen001/kubeutils/utils/logs"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type PersistentVolumeClaim struct {
	InstanceInterface typev1.CoreV1Interface
	Item              *corev1.PersistentVolumeClaim
}

// New函数用于配置一些初始化值
func NewPersistentVolumeClaim(kubeconfig string, item *corev1.PersistentVolumeClaim) *PersistentVolumeClaim {
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	//定义一个PersistentVolumeClaim实例
	resource := PersistentVolumeClaim{}
	resource.InstanceInterface = instance.Clientset.CoreV1()
	resource.Item = item
	return &resource
}

func (c *PersistentVolumeClaim) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name}, "创建Persistentvolumeclaim")
	_, err := c.InstanceInterface.PersistentVolumeClaims(namespace).Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

func (c *PersistentVolumeClaim) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "更新PersistentVolumeCliam")
	_, err := c.InstanceInterface.PersistentVolumeClaims(namespace).Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *PersistentVolumeClaim) Delete(namespace string, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "删除PersistentVolumeClaim")
	deleteOptions := v1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOptions.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.PersistentVolumeClaims(namespace).Delete(context.TODO(), name, deleteOptions)
	return err
}

func (c *PersistentVolumeClaim) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	for _, name := range nameList {
		logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "pi删除PersistentVolumeClaim")
		c.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (c *PersistentVolumeClaim) Get(namespace string, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "根据名字搜索对应的PersistentVolumeClaim")
	i, err := c.InstanceInterface.PersistentVolumeClaims(namespace).Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "v1"
	i.Kind = "PersistentVolumeClaim"
	item = i
	return item, err
}

func (c *PersistentVolumeClaim) List(namespace string, labelSelector string, fieldSelector string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"命名空间": namespace}, "搜索对应的PersistentVolumeClaim 队列")
	listOptions := v1.ListOptions{
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	i, err := c.InstanceInterface.PersistentVolumeClaims(namespace).List(context.TODO(), listOptions)
	return i.Items, err
}
