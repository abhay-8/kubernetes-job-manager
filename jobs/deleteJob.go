package jobs

import (
	"context"
	"fmt"
	"log"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DeleteJob() {
	fmt.Println("deleting job", jobName)
	jobsClient := clientset.BatchV1().Jobs(apiv1.NamespaceDefault)

	pp := metav1.DeletePropagationBackground
	err := jobsClient.Delete(context.Background(), jobName, metav1.DeleteOptions{PropagationPolicy: &pp})

	if err != nil {

		log.Fatal("failed to delete job", err)

	}

	fmt.Println("deleted job successfully")
}
