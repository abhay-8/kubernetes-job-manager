package cronjobs

import (
	"context"
	"fmt"
	"log"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func deleteCronJob() {

	fmt.Println("deleting cron job", cronJobName)

	cronJobsClient := clientset.BatchV1().CronJobs(apiv1.NamespaceDefault)

	pp := metav1.DeletePropagationBackground

	err := cronJobsClient.Delete(context.Background(), cronJobName, metav1.DeleteOptions{PropagationPolicy: &pp})

	if err != nil {

		log.Fatal("failed to delete cron job", err)

	}

	fmt.Println("deleted cron job successfully")

}
