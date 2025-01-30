package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/MrD0511/deck/deck-cli/internal/server/k3sclient"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func GetPodsHandler(c *gin.Context) {
	client, err := k3sclient.GetClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize K3s client"})
		return
	}

	pods, err := client.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error fetching pods: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pods"})
		return
	}

	var podDetails []gin.H
	for _, pod := range pods.Items {
		// Extract container names
		var containers []string
		for _, container := range pod.Spec.Containers {
			containers = append(containers, container.Name)
		}

		// Add detailed pod info
		podDetails = append(podDetails, gin.H{
			"name":        pod.Name,
			"namespace":   pod.Namespace,
			"status":      pod.Status.Phase,
			"node":        pod.Spec.NodeName,
			"containers":  containers,
			"podIP":       pod.Status.PodIP,
			"createdAt":   pod.CreationTimestamp.String(),
			"hostIP":      pod.Status.HostIP,
			"restarts":    pod.Status.ContainerStatuses[0].RestartCount,
		})
	}

	c.JSON(http.StatusOK, gin.H{"pods": podDetails})
}

