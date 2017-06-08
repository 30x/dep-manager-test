package main

import (
	"context"
	"fmt"
	"time"

	yaml "gopkg.in/yaml.v2"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/deploymentmanager/v2"
)

var (
	project = "turbo-dev"
	depName = "gotest12"
)

// Config ...
type Config struct {
	Resources []Resource `yaml:"resources"`
}

// Resource ...
type Resource struct {
	Name       string      `yaml:"name"`
	Properties *Properties `yaml:"properties"`
	Type       string      `yaml:"type"`
}

// Properties ...
type Properties struct {
	Apis               []string          `yaml:"apis"`
	Labels             map[string]string `yaml:"labels"`
	BillingAccountName string            `yaml:"billing-account-name"`
	IAMPolicy          IAMPolicy         `yaml:"iam-policy"`
	OrganizationID     string            `yaml:"organization-id"`
	Region             string            `yaml:"region"`
	ServiceAccounts    []string          `yaml:"service-accounts"`
}

// IAMPolicy ...
type IAMPolicy struct {
	Bindings []Bindings `yaml:"bindings"`
}

// Bindings ...
type Bindings struct {
	Members []string `yaml:"members"`
	Role    string   `yaml:"role"`
}

//Take in a name and generate the relevant yaml file
func genYaml(name string) error {
	return nil
}

func doCall(file []byte) error {
	ctx := context.Background()
	httpClient, err := google.DefaultClient(ctx, deploymentmanager.CloudPlatformScope)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	deploymentManagerClient, err := deploymentmanager.New(httpClient)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	//TODO: Generate hash for deployment name
	tempDep := deploymentmanager.Deployment{
		Name: depName,
		Target: &deploymentmanager.TargetConfiguration{
			Config: &deploymentmanager.ConfigFile{
				Content: string(file),
			},
		},
	}

	op, err := deploymentManagerClient.Deployments.Insert(project, &tempDep).Context(ctx).Do()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	fmt.Printf("%v\n", op.Name)

	done := false
	for done != true {
		getOp, err := deploymentManagerClient.Operations.Get(project, op.Name).Context(ctx).Do()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return err
		}
		if getOp.Status == "DONE" {
			done = true
		}
		fmt.Printf("%v\n", getOp.Status)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func main() {
	// TODO: Generate this config

	tempStore := Config{
		Resources: []Resource{
			{
				Name: "go-bin-test-3",
				Type: "turbo-dev/composite:project-test26",
				Properties: &Properties{
					Apis: []string{
						"deploymentmanager.googleapis.com",
						"logging.googleapis.com",
						"appengine.googleapis.com",
					},
					Labels: map[string]string{
						"test1": "val1",
					},
					BillingAccountName: "billingAccounts/00A5B6-123E67-6EEAD6",
					IAMPolicy: IAMPolicy{
						Bindings: []Bindings{
							{
								Members: []string{
									"serviceAccount:438049019500@cloudservices.gserviceaccount.com",
									"group:edgek8s@google.com",
								},
								Role: "roles/owner",
							},
						},
					},
					OrganizationID: "433637338589",
					Region:         "us-central",
					ServiceAccounts: []string{
						"my-service-account-1",
					},
				},
			},
		},
	}

	y, err := yaml.Marshal(tempStore)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// ioutil.WriteFile("gotest1.yaml", y, 0644)

	err = doCall(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}
