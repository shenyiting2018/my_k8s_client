package k8s_lib

import (
	"context"
	"errors"
	"flag"
	"path/filepath"
	"strings"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func connect() *kubernetes.Clientset {
	var kubeconfig *string
	home := homedir.HomeDir()
	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

// list namespaces
func ListNamespaces() *apiv1.NamespaceList {
	clientset := connect()
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	return namespaces
}

func validateNewNamespace(newNamespace string) error {
	if strings.HasPrefix(newNamespace, "kube-") {
		return errors.New("we do not support a namespace prefixed with kube-")
	}
	return nil
}

func CreateNewNamespace(newNamespace string) error {
	error := validateNewNamespace(newNamespace)
	if error != nil {
		panic(error)
	}

	clientset := connect()
	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: newNamespace,
		},
	}
	clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	return nil
}
