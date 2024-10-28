package kubeutils

import (
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeUtilser interface {
	Create(string) error
	Delete(string, string, *int64) error
	DeleteList(string, []string, *int64) error
	Get(string, string) (interface{}, error)
	List(string, string, string) (interface{}, error)
	Update(string) error
}

type ResourceInstance struct {
	Kubeconfig string
	Clientset  *kubernetes.Clientset
}

func (c *ResourceInstance) Init(kubeconfig string) {
	c.Kubeconfig = kubeconfig
	//生成Clientset
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.Kubeconfig))
	if err != nil {
		msg := "解析kubeconfig 错误：" + err.Error()
		panic(msg)
	}
	//设置超时时间
	restConfig.Timeout = 15 * time.Second
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		msg := "创建clientset失败：" + err.Error()
		panic(msg)
	}
	c.Clientset = clientset
}
