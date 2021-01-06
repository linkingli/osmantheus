package petals

import (
	"bufio"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"os"
	"osmantheus/client"
)

func ListNotRunngPod() {
	k8sClient := client.K8sClient()
	list, _ := k8sClient.CoreV1().Pods("").List(v1.ListOptions{})
	fmt.Println("--- not running pod ---")
	for _, pod := range list.Items {
		status := pod.Status.Phase
		if status != "Running" {
			fmt.Println(pod.Namespace, pod.Name)
		}
	}
	prompt()
}

func ListEoughController() {
	k8sClient := client.K8sClient()

	deploymentList, _ := k8sClient.AppsV1().Deployments("").List(v1.ListOptions{})
	rcList, _ := k8sClient.CoreV1().ReplicationControllers("").List(v1.ListOptions{})
	daemonSetList, _ := k8sClient.AppsV1().DaemonSets("").List(v1.ListOptions{})
	statefulSetList, _ := k8sClient.AppsV1().StatefulSets("").List(v1.ListOptions{})

	fmt.Println("--- not enough  controller ---")
	for _, deployment := range deploymentList.Items {
		applyedpods := deployment.Spec.Replicas
		currentpods := deployment.Status.AvailableReplicas
		if *applyedpods != currentpods {
			log.Println(deployment.Namespace, deployment.Name)
		}

	}
	for _, rc := range rcList.Items {
		applyedpods := rc.Spec.Replicas
		currentpods := rc.Status.AvailableReplicas
		if *applyedpods != currentpods {
			log.Println(rc.Namespace, rc.Name)
		}

	}
	for _, ds := range daemonSetList.Items {
		i := ds.Status.CurrentNumberScheduled
		if i != 1 {
			log.Println(ds.Namespace, ds.Name)
		}

	}
	for _, sts := range statefulSetList.Items {
		applyedpods := sts.Spec.Replicas
		currentpods := sts.Status.ReadyReplicas
		if *applyedpods != currentpods {
			log.Println(sts.Namespace, sts.Name)
		}

	}

	prompt()
}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}
