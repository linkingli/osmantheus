package petals

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"os/exec"
	"osmantheus/client"
	"strings"
)

//install doc: https://github.com/istio/istio/releases
func TestIstioInject() {
	//https://github.com/linkingli/istio-opentracing-demo/blob/master/k8s/eshop.yaml
	cmd := exec.Command("kubectl", "apply -f istio-demo.yaml")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	//https://istio.io/latest/docs/ops/diagnostic-tools/istioctl-describe/
	k8sClient := client.K8sClient()
	list, _ := k8sClient.CoreV1().Pods("").List(v1.ListOptions{})
	for _, pod := range list.Items {
		status := pod.Name
		if strings.Contains(status, "eshop") || strings.Contains(status, "inventory") || strings.Contains(status, "billing") || strings.Contains(status, "delivery") {
			cmd2 := exec.Command("istioctl  ", "experimental describe pod "+status)
			err2 := cmd2.Run()
			if err != nil {
				log.Fatal(err2)
			}
		}
	}

}

func IstioBasicInfo() {
	//https://istio.io/latest/docs/ops/diagnostic-tools/proxy-cmd/
	cmd := exec.Command("istioctl ", "proxy-status")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
