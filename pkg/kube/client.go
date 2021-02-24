package kube

import (
	"path/filepath"

	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	Kclientset            kubernetes.Interface
	ApiExtensionClientset *apiextensionsclientset.Clientset
	Dclientset            dynamic.Interface
	MysqlclusterRes       schema.GroupVersionResource
)

func Init(kubeconfig string) error {
	config, err := getKubeConfig(kubeconfig)
	if err != nil {
		return err
	}
	Kclientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	ApiExtensionClientset, err = apiextensionsclientset.NewForConfig(config)
	if err != nil {
		return err
	}
	Dclientset, err = dynamic.NewForConfig(config)
	if err != nil {
		return err
	}

	MysqlclusterRes = schema.GroupVersionResource{Group: "mysql.presslabs.org", Version: "v1alpha1", Resource: "mysqlclusters"}
	return nil
}

func getKubeConfig(kubeConf string) (*rest.Config, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		if kubeConf != "" {
			config, err := clientcmd.BuildConfigFromFlags("", kubeConf)
			if err != nil {
				return nil, err
			}
			return config, nil
		}
		var kubeconfig *string
		var tmp string
		if home := homedir.HomeDir(); home != "" {
			tmp = filepath.Join(home, ".kube", "config")
			kubeconfig = &tmp
		}
		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			return nil, err
		}
		return config, nil
	}
	return config, nil
}
