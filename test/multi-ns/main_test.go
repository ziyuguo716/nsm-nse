package multi_ns_test

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	//flags used
	kubeconfig  *string
	clusterName = flag.String("clusterName", "test", "Name of kind cluster for this test")
	ipAddr      = flag.String("REMOTE_IP", "127.0.0.1", "Set ENV to REMOTE_IP")
	nsmPath     = flag.String("nsmPath", "../../scripts/vl3/nsm_install_interdomain.sh", "Path of script to install NSM")
	nsePath     = flag.String("nsePath", "../../scripts/vl3/vl3_interdomain.sh", "Path of script to install NSE")

	clientset *kubernetes.Clientset
)

func TestMain(m *testing.M) {
	//Get kubeconfig
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	//Remove existing cluster which has the same name
	fmt.Println("Remove old cluster")
	removeExistingKindCluster(*clusterName)
	//Create a kind cluster for testing
	fmt.Printf("Creating kind cluster '%s'\n", *clusterName)
	execKindCluster("create", *clusterName)
	//Prepare clientset for K8s API
	clientset, _ = getClientSet()

	//Install NSM
	fmt.Println("Installing NSM...")
	cmd := exec.Command(*nsmPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	code := m.Run()
	//Remove the cluster after tests are done
	execKindCluster("delete", *clusterName)
	fmt.Printf("Deleted kind cluster '%s'\n", *clusterName)
	os.Exit(code)
}

func removeExistingKindCluster(clusterName string) {
	//Get all kind clusters
	out, err := exec.Command("kind", "get", "clusters").Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//Check if input cluster already exists
	//If found, remove it
	s := strings.Fields(string(out))
	for _, name := range s {
		if name == clusterName {
			execKindCluster("delete", clusterName)
		}
	}
}

func execKindCluster(action, clusterName string) {
	nameFlag := "--name=" + clusterName
	cmd := exec.Command("kind", action, "cluster", nameFlag)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func getClientSet() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return clientset, nil
}
