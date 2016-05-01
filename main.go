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
	disks := *vm.Properties.StorageProfile.DataDisks;
	for _, disk := range disks {
		//FIXME: check nil before reference
		fmt.Printf("lun %d name %s vhd %s size(GB): %d\n", *disk.Lun, *disk.Name, *disk.Vhd.URI, *disk.DiskSizeGB)
	}
	d := disks
	d = d[:len(d) - 1]
	newVM := compute.VirtualMachine {
		Location: vm.Location,
		Properties: &compute.VirtualMachineProperties {
			StorageProfile: &compute.StorageProfile {
				DataDisks: &d,
			},
		},
	}
	res, err := client.CreateOrUpdate(cred["resourceGroup"], cred["vm"],
		newVM, nil)
	fmt.Printf("delete disk: res %#v err %v\n", res, err)
}
