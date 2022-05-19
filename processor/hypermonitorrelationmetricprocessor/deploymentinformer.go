package hypermonitorrelationmetricprocessor

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
	"time"
)

var logger *zap.Logger
var DeploymentMap map[string]string

func intLogger() {
	logger, _ = zap.NewProduction()
	DeploymentMap = make(map[string]string)
}

func KubernetesStart() {
	intLogger()

	//config, err := clientcmd.BuildConfigFromFlags("", "/Users/williamlee/Desktop/code/pml/Informers-example/admin.conf")
	config, err := clientcmd.BuildConfigFromFlags("", "/conf/config")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)
	//表示每分钟进行一次resync，resync会周期性地执行List操作
	sharedInformers := informers.NewSharedInformerFactory(clientset, time.Minute)

	informerDeployment := sharedInformers.Apps().V1().Deployments().Informer()

	informerDeployment.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			syncDeploymentPodName(obj, clientset, DeploymentMap)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			syncDeploymentPodName(newObj, clientset, DeploymentMap)
		},
		DeleteFunc: func(obj interface{}) {
			syncDeploymentPodName(obj, clientset, DeploymentMap)
		},
	})

	informerDeployment.Run(stopCh)
}

func syncDeploymentPodName(obj interface{}, clientset *kubernetes.Clientset, deploymentMap map[string]string) {
	mObj := obj.(*appsv1.Deployment)
	logger.Info("get the deployment name", zap.String("name:", mObj.GetName()))

	pods, err := clientset.CoreV1().Pods(mObj.GetNamespace()).List(context.TODO(), v1.ListOptions{LabelSelector: convertMapToSelector(mObj.Spec.Selector.MatchLabels)})
	if err != nil {
		return
	}

	for _, pod := range pods.Items {
		deploymentMap[pod.GetName()] = mObj.GetName()
	}
}

// convertMapToSelector convert map to string, use comma connection: k1=v1,k2=v2
func convertMapToSelector(labels map[string]string) string {
	var l []string
	for k, v := range labels {
		l = append(l, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(l, ",")
}
