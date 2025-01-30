package k3sclient

import (
	"log"
	"os"
	"sync"
	"errors"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clientset *kubernetes.Clientset
	once      sync.Once
)

// GetClient initializes and returns a singleton Kubernetes clientset for K3s
func GetClient() (*kubernetes.Clientset, error) {
	var err error
	once.Do(func() {
		clientset, err = newClient()
		if err != nil {
			log.Printf("Error initializing K3s client: %v", err)
		}
	})

	if clientset == nil {
		return nil, errors.New("failed to initialize K3s client")
	}

	return clientset, err
}

// newClient initializes a K3s clientset
func newClient() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error

	// Check if running inside a cluster (K3s)
	if _, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount"); err == nil {
		log.Println("Running inside a K3s cluster, using in-cluster config")
		config, err = rest.InClusterConfig()
	} else {
		// Load KUBECONFIG from K3s default location
		kubeconfig := os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			kubeconfig = "/etc/rancher/k3s/k3s.yaml" // Default path for K3s
		}

		if _, err := os.Stat(kubeconfig); os.IsNotExist(err) {
			log.Fatalf("K3s config file not found: %s", kubeconfig)
			return nil, err
		}

		log.Printf("Using K3s KUBECONFIG: %s\n", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		log.Fatalf("Failed to create K3s config: %v", err)
		return nil, err
	}

	if config == nil {
		log.Fatal("K3s config is nil, cannot proceed")
		return nil, errors.New("K3s config is nil")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create K3s client: %v", err)
		return nil, err
	}

	log.Println("K3s client initialized successfully")
	return clientset, nil
}
