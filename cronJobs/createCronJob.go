package cronjobs

import (
	"context"
	"fmt"
	"log"

	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var clientset = &kubernetes.Clientset{}

const cronJobName = "test-cronjob-clientgo"

func createCronJob() {

	fmt.Println("creating cron job", cronJobName)
	cronJobsClient := clientset.BatchV1().CronJobs(apiv1.NamespaceDefault)

	cronJob := &batchv1.CronJob{ObjectMeta: metav1.ObjectMeta{
		Name: cronJobName},
		Spec: batchv1.CronJobSpec{
			Schedule: "* * * * *",
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: apiv1.PodTemplateSpec{
						Spec: apiv1.PodSpec{
							Containers: []apiv1.Container{
								{
									Name:    "hello",
									Image:   "busybox:1.28",
									Command: []string{"/bin/sh", "-c", "date; echo Hello from the Kubernetes cluster"},
								},
							},
							RestartPolicy: apiv1.RestartPolicyOnFailure,
						},
					},
				},
			},
		},
	}
	_, err := cronJobsClient.Create(context.Background(), cronJob, metav1.CreateOptions{})
	if err != nil {

		log.Fatal("failed to create cron job", err)

	}

	fmt.Println("created cron job successfully")
}
