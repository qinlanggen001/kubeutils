package kubeutils

import (
	"context"
	"kubeutils/utils/logs"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

type CronJob struct {
	InstanceInterface typedv1.BatchV1Interface
	Item              *batchv1.CronJob
}

func NewCronjob(kubeconfig string, item *batchv1.CronJob) *CronJob {
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	// 定义一个Cronjob实例
	resource := CronJob{}
	resource.InstanceInterface = instance.Clientset.BatchV1()
	resource.Item = item
	return &resource
}

// 创建cronjob
func (c *CronJob) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "创建ConfigMap")
	_, err := c.InstanceInterface.CronJobs(namespace).Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

// 删除cronjob, 设置*int64 是指针类型，可传空，如果是int64则不行
func (c *CronJob) Delete(namespace, name string, gradePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "删除ConfigMap")
	err := c.InstanceInterface.CronJobs(namespace).Delete(context.TODO(), name, v1.DeleteOptions{})
	return err
}

func (c *CronJob) DeleteList(namespace string, nameList []string, gradePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "删除ConfigMap列表")
	for _, name := range nameList {
		c.Delete(namespace, name, gradePeriodSeconds)
	}
	return nil
}

func (c *CronJob) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "更新ConfigMap")
	_, err := c.InstanceInterface.CronJobs(namespace).Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

func (c *CronJob) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称:": c.Item.Name, "命名空间：": namespace}, "获取configmap")
	i, err := c.InstanceInterface.CronJobs(namespace).Get(context.TODO(), name, v1.GetOptions{})
	i.APIVersion = "batch/v1"
	i.Kind = "CronJob"
	item = i
	return item, err
}

func (c *CronJob) List(namespace, labelSelector, fieldSelector string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称:": c.Item.Name, "命名空间：": namespace}, "获取configmap列表信息")
	listOpertions := v1.ListOptions{
		LabelSelector: labelSelector,
		FieldSelector: fieldSelector,
	}
	list, err := c.InstanceInterface.CronJobs(namespace).List(context.TODO(), listOpertions)
	item = list.Items
	return item, err
}
