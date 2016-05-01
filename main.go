package main

import (
	"fmt"
	
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/azure-sdk-for-go/arm/compute"
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
		return
	}
	//fmt.Printf("token: %+v\n", token)
	client := compute.NewVirtualMachinesClient(cred["subscriptionID"])
    client.Authorizer = token
	vm, err := client.Get(cred["resourceGroup"], cred["vm"], "")
	if err != nil {
		fmt.Printf("failed to get VM: %+v\n", err)
		return
	}
	disks := vm.Properties.StorageProfile.DataDisks;
	for _, disk := range *disks {
		//FIXME: check nil before reference
		fmt.Printf("lun %d name %s vhd %s size(GB): %d\n", *disk.Lun, *disk.Name, *disk.Vhd.URI, *disk.DiskSizeGB)
	}
}
