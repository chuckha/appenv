package kubernetes

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chuckha/appenv/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	statefulSet = "StatefulSet"
	deployment  = "Deployment"
	pod         = "Pod"
)

type KubernetesObject struct {
	Kind      string
	Name      string
	Namespace string
}

type ContainerVersion struct {
	Name  string
	Image string
}

type SSVersion struct {
	*KubernetesObject
	containerVersions []ContainerVersion
}

func (s SSVersion) Versions() []byte {
	var buf bytes.Buffer
	fmt.Fprintln(&buf, s.KubernetesObject.GetName())
	for _, cv := range s.containerVersions {
		fmt.Fprintf(&buf, "  - %s: %s\n", cv.Name, cv.Image)
	}
	return buf.Bytes()
}

func (k *KubernetesObject) Version() ([]byte, error) {
	switch k.Kind {
	case statefulSet:
		ss, err := clientset.AppsV1().StatefulSets(k.Namespace).Get(k.Name, metav1.GetOptions{})
		if err != nil {
			if k8serrors.IsNotFound(err) {
				return []byte{}, &errors.ErrorNotFound{k.GetName()}
			}
			return []byte{}, err
		}
		ssv := &SSVersion{}
		for _, container := range ss.Spec.Template.Spec.Containers {
			ssv.containerVersions = append(ssv.containerVersions, ContainerVersion{
				Name:  container.Name,
				Image: container.Image,
			})
		}
		return ssv.Versions(), nil
	}
	return nil, nil
}
func (k *KubernetesObject) GetName() string {
	return fmt.Sprintf("%s: %s/%s", k.Kind, k.Namespace, k.Name)
}

var clientset *kubernetes.Clientset

func StatefulSet(name, namespace string) *KubernetesObject {
	return &KubernetesObject{
		Kind:      statefulSet,
		Name:      name,
		Namespace: namespace,
	}
}

func init() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
