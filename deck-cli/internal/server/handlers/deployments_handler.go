package handlers

import (
	"context"
	"net/http"

	"github.com/MrD0511/deck/deck-cli/internal/server/k3sclient"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func GetDeploymentHandler(c *gin.Context){
	client, err := k3sclient.GetClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize K3s client"})
		return
	}

	deployments, err := client.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Deployments"})
	}

	var deploymentDetails []gin.H
	for _, deployment := range deployments.Items {
		deploymentDetails = append(deploymentDetails, gin.H{
			"name":        deployment.Name,
			"namespace":   deployment.Namespace,
			"replicas":    deployment.Spec.Replicas,
			"available":   deployment.Status.AvailableReplicas,
			"pods":        deployment.Status.Replicas,
			"labels":      deployment.Labels,
			"createdAt":   deployment.CreationTimestamp.String(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"deployments": deploymentDetails})

}
