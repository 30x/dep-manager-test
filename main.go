package main

import (
	"context"
	"fmt"

	"io/ioutil"

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

	//Get a yaml file
	tempYaml, err := ioutil.ReadFile("temp.yaml")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	//DEBUG
	fmt.Println(string(tempYaml))

	tempDep := deploymentmanager.Deployment{
		Name: "test4",
		Target: &deploymentmanager.TargetConfiguration{
			Config: &deploymentmanager.ConfigFile{
				Content: string(tempYaml),
			},
		},
	}

	resp, err := deploymentManagerClient.Deployments.Insert(project, &tempDep).Context(ctx).Do()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	_ = resp

}
