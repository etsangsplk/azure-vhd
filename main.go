package main

import (
	"fmt"
	
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/azure-sdk-for-go/arm/examples/helpers"
)

func main() {
	cred, err := helpers.LoadCredentials()
	if err != nil {
		fmt.Println("azure: failed to acquire service principal token")
		return
	}
	token, err := helpers.NewServicePrincipalTokenFromCredentials(cred,
		azure.PublicCloud.ServiceManagementEndpoint)
	if err != nil {
		fmt.Println("azure: failed to acquire service principal token")
	}
	fmt.Printf("token: %+v\n", token)
}
