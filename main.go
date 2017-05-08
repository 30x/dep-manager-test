package main

import (
	"context"
	"fmt"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/deploymentmanager/v2"
)

var (
	project = "turbo-dev"
)

func main() {
	ctx := context.Background()
	httpClient, err := google.DefaultClient(ctx, deploymentmanager.CloudPlatformScope)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	deploymentManagerClient, err := deploymentmanager.New(httpClient)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	tempDep := deploymentmanager.Deployment{
		Name: "test",
	}
	_ = tempDep

	resp, err := deploymentManagerClient.Deployments.Insert(project, &deploymentmanager.Deployment{
	//Fill in
	}).Context(ctx).Do()
	_ = resp

}
