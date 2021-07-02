package golang

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"golang.org/x/net/http2"
)
import (
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"time"
)

const (
	mockHost = "https://localhost:8443"
)

var (
	ctx            context.Context
	subscriptionId string
	cred           *azidentity.EnvironmentCredential
	err            error
	con            *armcore.Connection
)

func TestOperations_List(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestAvailabilitySets_CreateOrUpdate(t *testing.T) {

	// From example Create an availability set.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAvailabilitySetsClient(con,
		"{subscription-id}")
	_, err := client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"myAvailabilitySet",
		AvailabilitySet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &AvailabilitySetProperties{
				PlatformFaultDomainCount:  to.Int32Ptr(2),
				PlatformUpdateDomainCount: to.Int32Ptr(20),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestAvailabilitySets_Update(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestAvailabilitySets_Delete(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestAvailabilitySets_Get(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestAvailabilitySets_ListBySubscription(t *testing.T) {

	// From example List availability sets in a subscription.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewAvailabilitySetsClient(con,
		"{subscriptionId}")
	client.ListBySubscription(nil)

}

func TestAvailabilitySets_List(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestAvailabilitySets_ListAvailableSizes(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestProximityPlacementGroups_CreateOrUpdate(t *testing.T) {

	// From example Create or Update a proximity placement group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewProximityPlacementGroupsClient(con,
		"{subscription-id}")
	_, err := client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"myProximityPlacementGroup",
		ProximityPlacementGroup{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &ProximityPlacementGroupProperties{
				ProximityPlacementGroupType: ProximityPlacementGroupTypeStandard.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestProximityPlacementGroups_Update(t *testing.T) {

	// From example Create a proximity placement group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewProximityPlacementGroupsClient(con,
		"{subscription-id}")
	_, err := client.Update(ctx,
		"myResourceGroup",
		"myProximityPlacementGroup",
		ProximityPlacementGroupUpdate{
			UpdateResource: UpdateResource{
				Tags: map[string]*string{
					"additionalProp1": to.StringPtr("string"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestProximityPlacementGroups_Delete(t *testing.T) {

	// From example Create a proximity placement group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewProximityPlacementGroupsClient(con,
		"{subscription-id}")
	_, err := client.Delete(ctx,
		"myResourceGroup",
		"myProximityPlacementGroup",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestProximityPlacementGroups_Get(t *testing.T) {

	// From example Create a proximity placement group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewProximityPlacementGroupsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myProximityPlacementGroup",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestProximityPlacementGroups_ListBySubscription(t *testing.T) {

	// From example Create a proximity placement group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewProximityPlacementGroupsClient(con,
		"{subscription-id}")
	client.ListBySubscription(nil)

}

func TestProximityPlacementGroups_ListByResourceGroup(t *testing.T) {

	// From example Create a proximity placement group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewProximityPlacementGroupsClient(con,
		"{subscription-id}")
	client.ListByResourceGroup("myResourceGroup",
		nil)

}

func TestDedicatedHostGroups_CreateOrUpdate(t *testing.T) {

	// From example Create or update a dedicated host group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDedicatedHostGroupsClient(con,
		"{subscription-id}")
	_, err := client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"myDedicatedHostGroup",
		DedicatedHostGroup{
			Resource: Resource{
				Location: to.StringPtr("westus"),
				Tags: map[string]*string{
					"department": to.StringPtr("finance"),
				},
			},
			Properties: &DedicatedHostGroupProperties{
				PlatformFaultDomainCount:  to.Int32Ptr(3),
				SupportAutomaticPlacement: to.BoolPtr(true),
			},
			Zones: []*string{
				to.StringPtr("1")},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDedicatedHostGroups_Update(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestDedicatedHostGroups_Delete(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestDedicatedHostGroups_Get(t *testing.T) {

	// From example Create a dedicated host group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDedicatedHostGroupsClient(con,
		"{subscriptionId}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myDedicatedHostGroup",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDedicatedHostGroups_ListByResourceGroup(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestDedicatedHostGroups_ListBySubscription(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestDedicatedHosts_CreateOrUpdate(t *testing.T) {

	// From example Create or update a dedicated host .
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDedicatedHostsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDedicatedHostGroup",
		"myDedicatedHost",
		DedicatedHost{
			Resource: Resource{
				Location: to.StringPtr("westus"),
				Tags: map[string]*string{
					"department": to.StringPtr("HR"),
				},
			},
			Properties: &DedicatedHostProperties{
				PlatformFaultDomain: to.Int32Ptr(1),
			},
			SKU: &SKU{
				Name: to.StringPtr("DSv3-Type1"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDedicatedHosts_Update(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestDedicatedHosts_Delete(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestDedicatedHosts_Get(t *testing.T) {

	// From example Get a dedicated host.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDedicatedHostsClient(con,
		"{subscriptionId}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myDedicatedHostGroup",
		"myHost",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDedicatedHosts_ListByHostGroup(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestSshPublicKeys_ListBySubscription(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestSshPublicKeys_ListByResourceGroup(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestSshPublicKeys_Create(t *testing.T) {

	// From example Create a new SSH public key resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSSHPublicKeysClient(con,
		"{subscription-id}")
	_, err := client.Create(ctx,
		"myResourceGroup",
		"mySshPublicKeyName",
		SSHPublicKeyResource{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &SSHPublicKeyResourceProperties{
				PublicKey: to.StringPtr("{ssh-rsa public key}"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSshPublicKeys_Update(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestSshPublicKeys_Delete(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestSshPublicKeys_Get(t *testing.T) {

	// From example Get an ssh public key.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSSHPublicKeysClient(con,
		"{subscriptionId}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"mySshPublicKeyName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSshPublicKeys_GenerateKeyPair(t *testing.T) {

	// From example Generate an SSH key pair.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSSHPublicKeysClient(con,
		"{subscription-id}")
	_, err := client.GenerateKeyPair(ctx,
		"myResourceGroup",
		"mySshPublicKeyName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineExtensionImages_Get(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineExtensionImages_ListTypes(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineExtensionImages_ListVersions(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineExtensions_CreateOrUpdate(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineExtensions_Update(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineExtensions_Delete(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineExtensions_Get(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineExtensions_List(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineImages_Get(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineImages_List(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineImages_ListOffers(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineImages_ListPublishers(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineImages_ListSkus(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineImagesEdgeZone_Get(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineImagesEdgeZone_List(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineImagesEdgeZone_ListOffers(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineImagesEdgeZone_ListPublishers(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineImagesEdgeZone_ListSkus(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestUsage_List(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachines_ListByLocation(t *testing.T) {

	// From example Lists all the virtual machines under the specified subscription for the specified location.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscriptionId}")
	client.ListByLocation("eastus",
		nil)

}

func TestVirtualMachines_Capture(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachines_CreateOrUpdate(t *testing.T) {

	// From example Create a Linux vm with a patch setting assessmentMode of ImageDefault.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD2SV3.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
					LinuxConfiguration: &LinuxConfiguration{
						PatchSettings: &LinuxPatchSettings{
							AssessmentMode: LinuxPatchAssessmentModeImageDefault.ToPtr(),
						},
						ProvisionVMAgent: to.BoolPtr(true),
					},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("UbuntuServer"),
						Publisher: to.StringPtr("Canonical"),
						SKU:       to.StringPtr("16.04-LTS"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesPremiumLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a Linux vm with a patch setting patchMode of ImageDefault.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD2SV3.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
					LinuxConfiguration: &LinuxConfiguration{
						PatchSettings: &LinuxPatchSettings{
							PatchMode: LinuxVMGuestPatchModeImageDefault.ToPtr(),
						},
						ProvisionVMAgent: to.BoolPtr(true),
					},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("UbuntuServer"),
						Publisher: to.StringPtr("Canonical"),
						SKU:       to.StringPtr("16.04-LTS"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesPremiumLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a Linux vm with a patch settings patchMode and assessmentMode set to AutomaticByPlatform.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD2SV3.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
					LinuxConfiguration: &LinuxConfiguration{
						PatchSettings: &LinuxPatchSettings{
							AssessmentMode: LinuxPatchAssessmentModeAutomaticByPlatform.ToPtr(),
							PatchMode:      LinuxVMGuestPatchModeAutomaticByPlatform.ToPtr(),
						},
						ProvisionVMAgent: to.BoolPtr(true),
					},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("UbuntuServer"),
						Publisher: to.StringPtr("Canonical"),
						SKU:       to.StringPtr("16.04-LTS"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesPremiumLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a VM with Uefi Settings of secureBoot and vTPM.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD2SV3.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				SecurityProfile: &SecurityProfile{
					SecurityType: to.StringPtr("TrustedLaunch"),
					UefiSettings: &UefiSettings{
						SecureBootEnabled: to.BoolPtr(true),
						VTpmEnabled:       to.BoolPtr(true),
					},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("windowsserver-gen2preview-preview"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("windows10-tvm"),
						Version:   to.StringPtr("18363.592.2001092016"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadOnly.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardSSDLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a VM with UserData
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vm-name}",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				DiagnosticsProfile: &DiagnosticsProfile{
					BootDiagnostics: &BootDiagnostics{
						Enabled:    to.BoolPtr(true),
						StorageURI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net"),
					},
				},
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("{vm-name}"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("vmOSdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
				UserData: to.StringPtr("RXhhbXBsZSBVc2VyRGF0YQ=="),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a VM with network interface configuration
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkAPIVersion: NetworkAPIVersionTwoThousandTwenty1101.ToPtr(),
					NetworkInterfaceConfigurations: []*VirtualMachineNetworkInterfaceConfiguration{
						&VirtualMachineNetworkInterfaceConfiguration{
							Name: to.StringPtr("{nic-config-name}"),
							Properties: &VirtualMachineNetworkInterfaceConfigurationProperties{
								DeleteOption: DeleteOptionsDelete.ToPtr(),
								IPConfigurations: []*VirtualMachineNetworkInterfaceIPConfiguration{
									&VirtualMachineNetworkInterfaceIPConfiguration{
										Name: to.StringPtr("{ip-config-name}"),
										Properties: &VirtualMachineNetworkInterfaceIPConfigurationProperties{
											Primary: to.BoolPtr(true),
											PublicIPAddressConfiguration: &VirtualMachinePublicIPAddressConfiguration{
												Name: to.StringPtr("{publicIP-config-name}"),
												Properties: &VirtualMachinePublicIPAddressConfigurationProperties{
													DeleteOption:             DeleteOptionsDetach.ToPtr(),
													PublicIPAllocationMethod: PublicIPAllocationMethodStatic.ToPtr(),
												},
												SKU: &PublicIPAddressSKU{
													PublicIPAddressSKUName: PublicIPAddressSKUNameBasic.ToPtr(),
													PublicIPAddressSKUTier: PublicIPAddressSKUTierGlobal.ToPtr(),
												},
											},
										},
									}},
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a Windows vm with a patch setting assessmentMode of ImageDefault.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
					WindowsConfiguration: &WindowsConfiguration{
						EnableAutomaticUpdates: to.BoolPtr(true),
						PatchSettings: &PatchSettings{
							AssessmentMode: WindowsPatchAssessmentModeImageDefault.ToPtr(),
						},
						ProvisionVMAgent: to.BoolPtr(true),
					},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesPremiumLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a Windows vm with a patch setting patchMode of AutomaticByOS.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/nsgExistingNic"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
					WindowsConfiguration: &WindowsConfiguration{
						EnableAutomaticUpdates: to.BoolPtr(true),
						PatchSettings: &PatchSettings{
							PatchMode: WindowsVMGuestPatchModeAutomaticByOS.ToPtr(),
						},
						ProvisionVMAgent: to.BoolPtr(true),
					},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesPremiumLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a Windows vm with a patch setting patchMode of AutomaticByPlatform and enableHotpatching set to true.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
					WindowsConfiguration: &WindowsConfiguration{
						EnableAutomaticUpdates: to.BoolPtr(true),
						PatchSettings: &PatchSettings{
							EnableHotpatching: to.BoolPtr(true),
							PatchMode:         WindowsVMGuestPatchModeAutomaticByPlatform.ToPtr(),
						},
						ProvisionVMAgent: to.BoolPtr(true),
					},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesPremiumLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a Windows vm with a patch setting patchMode of Manual.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
					WindowsConfiguration: &WindowsConfiguration{
						EnableAutomaticUpdates: to.BoolPtr(true),
						PatchSettings: &PatchSettings{
							PatchMode: WindowsVMGuestPatchModeManual.ToPtr(),
						},
						ProvisionVMAgent: to.BoolPtr(true),
					},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesPremiumLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a Windows vm with patch settings patchMode and assessmentMode set to AutomaticByPlatform.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
					WindowsConfiguration: &WindowsConfiguration{
						EnableAutomaticUpdates: to.BoolPtr(true),
						PatchSettings: &PatchSettings{
							AssessmentMode: WindowsPatchAssessmentModeAutomaticByPlatform.ToPtr(),
							PatchMode:      WindowsVMGuestPatchModeAutomaticByPlatform.ToPtr(),
						},
						ProvisionVMAgent: to.BoolPtr(true),
					},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesPremiumLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a custom-image vm from an unmanaged generalized os image.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vm-name}",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						Image: &VirtualHardDisk{
							URI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/{existing-generalized-os-image-blob-name}.vhd"),
						},
						OSType: OperatingSystemTypesWindows.ToPtr(),
						Vhd: &VirtualHardDisk{
							URI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a platform-image vm with unmanaged os and data disks.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vm-name}",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD2V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					DataDisks: []*DataDisk{
						&DataDisk{
							CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(0),
							Vhd: &VirtualHardDisk{
								URI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk0.vhd"),
							},
						},
						&DataDisk{
							CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(1),
							Vhd: &VirtualHardDisk{
								URI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk1.vhd"),
							},
						}},
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						Vhd: &VirtualHardDisk{
							URI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm from a custom image.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						SubResource: SubResource{
							ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"),
						},
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm from a generalized shared image.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						SubResource: SubResource{
							ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/mySharedGallery/images/mySharedImage"),
						},
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm from a specialized shared image.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						SubResource: SubResource{
							ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/mySharedGallery/images/mySharedImage"),
						},
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm in a Virtual Machine Scale Set with customer assigned platformFaultDomain.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				PlatformFaultDomain: to.Int32Ptr(1),
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
				VirtualMachineScaleSet: &SubResource{
					ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{existing-flex-vmss-name-with-platformFaultDomainCount-greater-than-1}"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm in an availability set.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				AvailabilitySet: &SubResource{
					ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/{existing-availability-set-name}"),
				},
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with DiskEncryptionSet resource id in the os disk and data disk.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					DataDisks: []*DataDisk{
						&DataDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(0),
							ManagedDisk: &ManagedDiskParameters{
								DiskEncryptionSet: &DiskEncryptionSetParameters{
									SubResource: SubResource{
										ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
									},
								},
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
						&DataDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesAttach.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(1),
							ManagedDisk: &ManagedDiskParameters{
								SubResource: SubResource{
									ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{existing-managed-disk-name}"),
								},
								DiskEncryptionSet: &DiskEncryptionSetParameters{
									SubResource: SubResource{
										ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
									},
								},
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						}},
					ImageReference: &ImageReference{
						SubResource: SubResource{
							ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"),
						},
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							DiskEncryptionSet: &DiskEncryptionSetParameters{
								SubResource: SubResource{
									ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
								},
							},
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with Host Encryption using encryptionAtHost property.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Plan: &Plan{
				Name:      to.StringPtr("windows2016"),
				Product:   to.StringPtr("windows-data-science-vm"),
				Publisher: to.StringPtr("microsoft-ads"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardDS1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				SecurityProfile: &SecurityProfile{
					EncryptionAtHost: to.BoolPtr(true),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("windows-data-science-vm"),
						Publisher: to.StringPtr("microsoft-ads"),
						SKU:       to.StringPtr("windows2016"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadOnly.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with Scheduled Events Profile
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				DiagnosticsProfile: &DiagnosticsProfile{
					BootDiagnostics: &BootDiagnostics{
						Enabled:    to.BoolPtr(true),
						StorageURI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net"),
					},
				},
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				ScheduledEventsProfile: &ScheduledEventsProfile{
					TerminateNotificationProfile: &TerminateNotificationProfile{
						Enable:           to.BoolPtr(true),
						NotBeforeTimeout: to.StringPtr("PT10M"),
					},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with a marketplace image plan.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Plan: &Plan{
				Name:      to.StringPtr("windows2016"),
				Product:   to.StringPtr("windows-data-science-vm"),
				Publisher: to.StringPtr("microsoft-ads"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("windows-data-science-vm"),
						Publisher: to.StringPtr("microsoft-ads"),
						SKU:       to.StringPtr("windows2016"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with an extensions time budget.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				DiagnosticsProfile: &DiagnosticsProfile{
					BootDiagnostics: &BootDiagnostics{
						Enabled:    to.BoolPtr(true),
						StorageURI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net"),
					},
				},
				ExtensionsTimeBudget: to.StringPtr("PT30M"),
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with boot diagnostics.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				DiagnosticsProfile: &DiagnosticsProfile{
					BootDiagnostics: &BootDiagnostics{
						Enabled:    to.BoolPtr(true),
						StorageURI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net"),
					},
				},
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with empty data disks.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD2V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					DataDisks: []*DataDisk{
						&DataDisk{
							CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(0),
						},
						&DataDisk{
							CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(1),
						}},
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with ephemeral os disk provisioning in Cache disk using placement property.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Plan: &Plan{
				Name:      to.StringPtr("windows2016"),
				Product:   to.StringPtr("windows-data-science-vm"),
				Publisher: to.StringPtr("microsoft-ads"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardDS1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("windows-data-science-vm"),
						Publisher: to.StringPtr("microsoft-ads"),
						SKU:       to.StringPtr("windows2016"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadOnly.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						DiffDiskSettings: &DiffDiskSettings{
							Option:    DiffDiskOptionsLocal.ToPtr(),
							Placement: DiffDiskPlacementCacheDisk.ToPtr(),
						},
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with ephemeral os disk provisioning in Resource disk using placement property.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Plan: &Plan{
				Name:      to.StringPtr("windows2016"),
				Product:   to.StringPtr("windows-data-science-vm"),
				Publisher: to.StringPtr("microsoft-ads"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardDS1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("windows-data-science-vm"),
						Publisher: to.StringPtr("microsoft-ads"),
						SKU:       to.StringPtr("windows2016"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadOnly.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						DiffDiskSettings: &DiffDiskSettings{
							Option:    DiffDiskOptionsLocal.ToPtr(),
							Placement: DiffDiskPlacementResourceDisk.ToPtr(),
						},
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with ephemeral os disk.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Plan: &Plan{
				Name:      to.StringPtr("windows2016"),
				Product:   to.StringPtr("windows-data-science-vm"),
				Publisher: to.StringPtr("microsoft-ads"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardDS1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("windows-data-science-vm"),
						Publisher: to.StringPtr("microsoft-ads"),
						SKU:       to.StringPtr("windows2016"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadOnly.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						DiffDiskSettings: &DiffDiskSettings{
							Option: DiffDiskOptionsLocal.ToPtr(),
						},
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with managed boot diagnostics.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				DiagnosticsProfile: &DiagnosticsProfile{
					BootDiagnostics: &BootDiagnostics{
						Enabled: to.BoolPtr(true),
					},
				},
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with password authentication.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with premium storage.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesPremiumLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a vm with ssh authentication.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD1V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
					LinuxConfiguration: &LinuxConfiguration{
						DisablePasswordAuthentication: to.BoolPtr(true),
						SSH: &SSHConfiguration{
							PublicKeys: []*SSHPublicKey{
								&SSHPublicKey{
									Path:    to.StringPtr("/home/{your-username}/.ssh/authorized_keys"),
									KeyData: to.StringPtr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1"),
								}},
						},
					},
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("{image_offer}"),
						Publisher: to.StringPtr("{image_publisher}"),
						SKU:       to.StringPtr("{image_sku}"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_Update(t *testing.T) {

	// From example Update a VM by detaching data disk
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachineUpdate{
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD2V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					DataDisks: []*DataDisk{
						&DataDisk{
							CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(0),
							ToBeDetached: to.BoolPtr(true),
						},
						&DataDisk{
							CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(1),
							ToBeDetached: to.BoolPtr(false),
						}},
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Update a VM by force-detaching data disk
	poller, err = client.BeginUpdate(ctx,
		"myResourceGroup",
		"myVM",
		VirtualMachineUpdate{
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardD2V2.ToPtr(),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: []*NetworkInterfaceReference{
						&NetworkInterfaceReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						}},
				},
				OSProfile: &OSProfile{
					AdminPassword: to.StringPtr("{your-password}"),
					AdminUsername: to.StringPtr("{your-username}"),
					ComputerName:  to.StringPtr("myVM"),
				},
				StorageProfile: &StorageProfile{
					DataDisks: []*DataDisk{
						&DataDisk{
							CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
							DetachOption: DiskDetachOptionTypesForceDetach.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(0),
							ToBeDetached: to.BoolPtr(true),
						},
						&DataDisk{
							CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(1023),
							Lun:          to.Int32Ptr(1),
							ToBeDetached: to.BoolPtr(false),
						}},
					ImageReference: &ImageReference{
						Offer:     to.StringPtr("WindowsServer"),
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						SKU:       to.StringPtr("2016-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
					OSDisk: &OSDisk{
						Name:         to.StringPtr("myVMosdisk"),
						Caching:      CachingTypesReadWrite.ToPtr(),
						CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
						ManagedDisk: &ManagedDiskParameters{
							StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_Delete(t *testing.T) {

	// From example Force delete a VM
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myVM",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_Get(t *testing.T) {

	// From example Get a Virtual Machine.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myVM",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	// From example Get a virtual machine placed on a dedicated host group through automatic placement
	_, err = client.Get(ctx,
		"myResourceGroup",
		"myVM",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_InstanceView(t *testing.T) {

	// From example Get Virtual Machine Instance View.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	_, err := client.InstanceView(ctx,
		"myResourceGroup",
		"myVM",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	// From example Get instance view of a virtual machine placed on a dedicated host group through automatic placement.
	_, err = client.InstanceView(ctx,
		"myResourceGroup",
		"myVM",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_ConvertToManagedDisks(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachines_Deallocate(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachines_Generalize(t *testing.T) {

	// From example Generalize a Virtual Machine.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	_, err := client.Generalize(ctx,
		"myResourceGroup",
		"myVMName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_List(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachines_ListAll(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachines_ListAvailableSizes(t *testing.T) {

	// From example Lists all available virtual machine sizes to which the specified virtual machine can be resized
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	_, err := client.ListAvailableSizes(ctx,
		"myResourceGroup",
		"myVmName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_PowerOff(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachines_Reapply(t *testing.T) {

	// From example Reapply the state of a virtual machine.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	poller, err := client.BeginReapply(ctx,
		"ResourceGroup",
		"VMName",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_Restart(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachines_Start(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachines_Redeploy(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachines_Reimage(t *testing.T) {

	// From example Reimage a Virtual Machine.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	poller, err := client.BeginReimage(ctx,
		"myResourceGroup",
		"myVMName",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_RetrieveBootDiagnosticsData(t *testing.T) {

	// From example RetrieveBootDiagnosticsData of a virtual machine.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	_, err := client.RetrieveBootDiagnosticsData(ctx,
		"ResourceGroup",
		"VMName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_PerformMaintenance(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachines_SimulateEviction(t *testing.T) {

	// From example Simulate Eviction a virtual machine.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	_, err := client.SimulateEviction(ctx,
		"ResourceGroup",
		"VMName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_AssessPatches(t *testing.T) {

	// From example Assess patch state of a virtual machine.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	poller, err := client.BeginAssessPatches(ctx,
		"myResourceGroupName",
		"myVMName",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_InstallPatches(t *testing.T) {

	// From example Install patch state of a virtual machine.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"{subscription-id}")
	poller, err := client.BeginInstallPatches(ctx,
		"myResourceGroupName",
		"myVMName",
		VirtualMachineInstallPatchesParameters{
			MaximumDuration: to.StringPtr("PT4H"),
			RebootSetting:   VMGuestPatchRebootSettingIfRequired.ToPtr(),
			WindowsParameters: &WindowsParameters{
				ClassificationsToInclude: []*VMGuestPatchClassificationWindows{
					VMGuestPatchClassificationWindowsCritical.ToPtr(),
					VMGuestPatchClassificationWindowsSecurity.ToPtr()},
				MaxPatchPublishDate: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339, "2020-11-19T02:36:43.0539904+00:00"); return t }()),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachines_RunCommand(t *testing.T) {

	// From example VirtualMachineRunCommand
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachinesClient(con,
		"24fb23e3-6ba3-41f0-9b6e-e41131d5d61e")
	poller, err := client.BeginRunCommand(ctx,
		"crptestar98131",
		"vm3036",
		RunCommandInput{
			CommandID: to.StringPtr("RunPowerShellScript"),
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSets_ListByLocation(t *testing.T) {

	// From example Lists all the VM scale sets under the specified subscription for the specified location.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetsClient(con,
		"{subscription-id}")
	client.ListByLocation("eastus",
		nil)

}

func TestVirtualMachineScaleSets_CreateOrUpdate(t *testing.T) {

	// From example Create a custom-image scale set from an unmanaged generalized os image.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Name:         to.StringPtr("osDisk"),
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							Image: &VirtualHardDisk{
								URI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/{existing-generalized-os-image-blob-name}.vhd"),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a platform-image scale set with unmanaged os disks.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Name:         to.StringPtr("osDisk"),
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							VhdContainers: []*string{
								to.StringPtr("http://{existing-storage-account-name-0}.blob.core.windows.net/vhdContainer"),
								to.StringPtr("http://{existing-storage-account-name-1}.blob.core.windows.net/vhdContainer"),
								to.StringPtr("http://{existing-storage-account-name-2}.blob.core.windows.net/vhdContainer"),
								to.StringPtr("http://{existing-storage-account-name-3}.blob.core.windows.net/vhdContainer"),
								to.StringPtr("http://{existing-storage-account-name-4}.blob.core.windows.net/vhdContainer")},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set from a custom image.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"),
							},
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set from a generalized shared image.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/mySharedGallery/images/mySharedImage"),
							},
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set from a specialized shared image.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/mySharedGallery/images/mySharedImage"),
							},
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with DiskEncryptionSet resource in os disk and data disk.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						DataDisks: []*VirtualMachineScaleSetDataDisk{
							&VirtualMachineScaleSetDataDisk{
								Caching:      CachingTypesReadWrite.ToPtr(),
								CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
								DiskSizeGB:   to.Int32Ptr(1023),
								Lun:          to.Int32Ptr(0),
								ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
									DiskEncryptionSet: &DiskEncryptionSetParameters{
										SubResource: SubResource{
											ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
										},
									},
									StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
								},
							}},
						ImageReference: &ImageReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"),
							},
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								DiskEncryptionSet: &DiskEncryptionSetParameters{
									SubResource: SubResource{
										ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
									},
								},
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_DS1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with Fpga Network Interfaces.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							},
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{fpgaNic-Name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableAcceleratedNetworking: to.BoolPtr(false),
									EnableFpga:                  to.BoolPtr(true),
									EnableIPForwarding:          to.BoolPtr(false),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{fpgaNic-Name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Primary:                 to.BoolPtr(true),
												PrivateIPAddressVersion: IPVersionIPv4.ToPtr(),
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-fpga-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(false),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							SubResource: SubResource{
								ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"),
							},
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with Host Encryption using encryptionAtHost property.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Plan: &Plan{
				Name:      to.StringPtr("windows2016"),
				Product:   to.StringPtr("windows-data-science-vm"),
				Publisher: to.StringPtr("microsoft-ads"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					SecurityProfile: &SecurityProfile{
						EncryptionAtHost: to.BoolPtr(true),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("windows-data-science-vm"),
							Publisher: to.StringPtr("microsoft-ads"),
							SKU:       to.StringPtr("windows2016"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadOnly.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_DS1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with Uefi Settings of secureBoot and vTPM.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					SecurityProfile: &SecurityProfile{
						SecurityType: to.StringPtr("TrustedLaunch"),
						UefiSettings: &UefiSettings{
							SecureBootEnabled: to.BoolPtr(true),
							VTpmEnabled:       to.BoolPtr(true),
						},
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("windowsserver-gen2preview-preview"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("windows10-tvm"),
							Version:   to.StringPtr("18363.592.2001092016"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadOnly.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardSSDLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D2s_v3"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with a marketplace image plan.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Plan: &Plan{
				Name:      to.StringPtr("windows2016"),
				Product:   to.StringPtr("windows-data-science-vm"),
				Publisher: to.StringPtr("microsoft-ads"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("windows-data-science-vm"),
							Publisher: to.StringPtr("microsoft-ads"),
							SKU:       to.StringPtr("windows2016"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with an azure application gateway.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												ApplicationGatewayBackendAddressPools: []*SubResource{
													&SubResource{
														ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/applicationGateways/{existing-application-gateway-name}/backendAddressPools/{existing-backend-address-pool-name}"),
													}},
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with an azure load balancer.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												LoadBalancerBackendAddressPools: []*SubResource{
													&SubResource{
														ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/loadBalancers/{existing-load-balancer-name}/backendAddressPools/{existing-backend-address-pool-name}"),
													}},
												LoadBalancerInboundNatPools: []*SubResource{
													&SubResource{
														ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/loadBalancers/{existing-load-balancer-name}/inboundNatPools/{existing-nat-pool-name}"),
													}},
												PublicIPAddressConfiguration: &VirtualMachineScaleSetPublicIPAddressConfiguration{
													Name: to.StringPtr("{vmss-name}"),
													Properties: &VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
														PublicIPAddressVersion: IPVersionIPv4.ToPtr(),
													},
												},
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with automatic repairs enabled
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				AutomaticRepairsPolicy: &AutomaticRepairsPolicy{
					Enabled:     to.BoolPtr(true),
					GracePeriod: to.StringPtr("PT30M"),
				},
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with boot diagnostics.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					DiagnosticsProfile: &DiagnosticsProfile{
						BootDiagnostics: &BootDiagnostics{
							Enabled:    to.BoolPtr(true),
							StorageURI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net"),
						},
					},
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with empty data disks on each vm.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						DataDisks: []*VirtualMachineScaleSetDataDisk{
							&VirtualMachineScaleSetDataDisk{
								CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
								DiskSizeGB:   to.Int32Ptr(1023),
								Lun:          to.Int32Ptr(0),
							},
							&VirtualMachineScaleSetDataDisk{
								CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
								DiskSizeGB:   to.Int32Ptr(1023),
								Lun:          to.Int32Ptr(1),
							}},
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(512),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D2_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with ephemeral os disks using placement property.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Plan: &Plan{
				Name:      to.StringPtr("windows2016"),
				Product:   to.StringPtr("windows-data-science-vm"),
				Publisher: to.StringPtr("microsoft-ads"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("windows-data-science-vm"),
							Publisher: to.StringPtr("microsoft-ads"),
							SKU:       to.StringPtr("windows2016"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadOnly.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							DiffDiskSettings: &DiffDiskSettings{
								Option:    DiffDiskOptionsLocal.ToPtr(),
								Placement: DiffDiskPlacementResourceDisk.ToPtr(),
							},
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_DS1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with ephemeral os disks.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Plan: &Plan{
				Name:      to.StringPtr("windows2016"),
				Product:   to.StringPtr("windows-data-science-vm"),
				Publisher: to.StringPtr("microsoft-ads"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("windows-data-science-vm"),
							Publisher: to.StringPtr("microsoft-ads"),
							SKU:       to.StringPtr("windows2016"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadOnly.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							DiffDiskSettings: &DiffDiskSettings{
								Option: DiffDiskOptionsLocal.ToPtr(),
							},
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_DS1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with extension time budget.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					DiagnosticsProfile: &DiagnosticsProfile{
						BootDiagnostics: &BootDiagnostics{
							Enabled:    to.BoolPtr(true),
							StorageURI: to.StringPtr("http://{existing-storage-account-name}.blob.core.windows.net"),
						},
					},
					ExtensionProfile: &VirtualMachineScaleSetExtensionProfile{
						ExtensionsTimeBudget: to.StringPtr("PT1H20M"),
						Extensions: []*VirtualMachineScaleSetExtension{
							&VirtualMachineScaleSetExtension{
								Name: to.StringPtr("{extension-name}"),
								Properties: &VirtualMachineScaleSetExtensionProperties{
									Type:                    to.StringPtr("{extension-Type}"),
									AutoUpgradeMinorVersion: to.BoolPtr(false),
									Publisher:               to.StringPtr("{extension-Publisher}"),
									Settings:                map[string]interface{}{},
									TypeHandlerVersion:      to.StringPtr("{handler-version}"),
								},
							}},
					},
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with managed boot diagnostics.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					DiagnosticsProfile: &DiagnosticsProfile{
						BootDiagnostics: &BootDiagnostics{
							Enabled: to.BoolPtr(true),
						},
					},
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with password authentication.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with premium storage.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesPremiumLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with ssh authentication.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
						LinuxConfiguration: &LinuxConfiguration{
							DisablePasswordAuthentication: to.BoolPtr(true),
							SSH: &SSHConfiguration{
								PublicKeys: []*SSHPublicKey{
									&SSHPublicKey{
										Path:    to.StringPtr("/home/{your-username}/.ssh/authorized_keys"),
										KeyData: to.StringPtr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1"),
									}},
							},
						},
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with terminate scheduled events enabled.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					ScheduledEventsProfile: &ScheduledEventsProfile{
						TerminateNotificationProfile: &TerminateNotificationProfile{
							Enable:           to.BoolPtr(true),
							NotBeforeTimeout: to.StringPtr("PT5M"),
						},
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with userData.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("westus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
					UserData: to.StringPtr("RXhhbXBsZSBVc2VyRGF0YQ=="),
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_D1_v2"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("Standard"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a scale set with virtual machines in different zones.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"{vmss-name}",
		VirtualMachineScaleSet{
			Resource: Resource{
				Location: to.StringPtr("centralus"),
			},
			Properties: &VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &UpgradePolicy{
					Mode: UpgradeModeAutomatic.ToPtr(),
				},
				VirtualMachineProfile: &VirtualMachineScaleSetVMProfile{
					NetworkProfile: &VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*VirtualMachineScaleSetNetworkConfiguration{
							&VirtualMachineScaleSetNetworkConfiguration{
								Name: to.StringPtr("{vmss-name}"),
								Properties: &VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*VirtualMachineScaleSetIPConfiguration{
										&VirtualMachineScaleSetIPConfiguration{
											Name: to.StringPtr("{vmss-name}"),
											Properties: &VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &APIEntityReference{
													ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("{your-password}"),
						AdminUsername:      to.StringPtr("{your-username}"),
						ComputerNamePrefix: to.StringPtr("{vmss-name}"),
					},
					StorageProfile: &VirtualMachineScaleSetStorageProfile{
						DataDisks: []*VirtualMachineScaleSetDataDisk{
							&VirtualMachineScaleSetDataDisk{
								CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
								DiskSizeGB:   to.Int32Ptr(1023),
								Lun:          to.Int32Ptr(0),
							},
							&VirtualMachineScaleSetDataDisk{
								CreateOption: DiskCreateOptionTypesEmpty.ToPtr(),
								DiskSizeGB:   to.Int32Ptr(1023),
								Lun:          to.Int32Ptr(1),
							}},
						ImageReference: &ImageReference{
							Offer:     to.StringPtr("WindowsServer"),
							Publisher: to.StringPtr("MicrosoftWindowsServer"),
							SKU:       to.StringPtr("2016-Datacenter"),
							Version:   to.StringPtr("latest"),
						},
						OSDisk: &VirtualMachineScaleSetOSDisk{
							Caching:      CachingTypesReadWrite.ToPtr(),
							CreateOption: DiskCreateOptionTypesFromImage.ToPtr(),
							DiskSizeGB:   to.Int32Ptr(512),
							ManagedDisk: &VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: StorageAccountTypesStandardLRS.ToPtr(),
							},
						},
					},
				},
			},
			SKU: &SKU{
				Name:     to.StringPtr("Standard_A1_v2"),
				Capacity: to.Int64Ptr(2),
				Tier:     to.StringPtr("Standard"),
			},
			Zones: []*string{
				to.StringPtr("1"),
				to.StringPtr("3")},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSets_Update(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_Delete(t *testing.T) {

	// From example Force Delete a VM scale set.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetsClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSets_Get(t *testing.T) {

	// From example Get a virtual machine scale set placed on a dedicated host group through automatic placement.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myVirtualMachineScaleSet",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	// From example Get a virtual machine scale set with UserData
	_, err = client.Get(ctx,
		"myResourceGroup",
		"myVirtualMachineScaleSet",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSets_Deallocate(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_DeleteInstances(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_GetInstanceView(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_List(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_ListAll(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_ListSkus(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_GetOSUpgradeHistory(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_PowerOff(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_Restart(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_Start(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_Redeploy(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_PerformMaintenance(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_UpdateInstances(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_Reimage(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_ReimageAll(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_ForceRecoveryServiceFabricPlatformUpdateDomainWalk(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_ConvertToSinglePlacementGroup(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSets_SetOrchestrationServiceState(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineSizes_List(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestImages_CreateOrUpdate(t *testing.T) {

	// From example Create a virtual machine image from a blob with DiskEncryptionSet resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewImagesClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myImage",
		Image{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &ImageProperties{
				StorageProfile: &ImageStorageProfile{
					OSDisk: &ImageOSDisk{
						ImageDisk: ImageDisk{
							BlobURI: to.StringPtr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
							DiskEncryptionSet: &DiskEncryptionSetParameters{
								SubResource: SubResource{
									ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
								},
							},
						},
						OSState: OperatingSystemStateTypesGeneralized.ToPtr(),
						OSType:  OperatingSystemTypesLinux.ToPtr(),
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a virtual machine image from a blob.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myImage",
		Image{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &ImageProperties{
				StorageProfile: &ImageStorageProfile{
					OSDisk: &ImageOSDisk{
						ImageDisk: ImageDisk{
							BlobURI: to.StringPtr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
						},
						OSState: OperatingSystemStateTypesGeneralized.ToPtr(),
						OSType:  OperatingSystemTypesLinux.ToPtr(),
					},
					ZoneResilient: to.BoolPtr(true),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a virtual machine image from a managed disk with DiskEncryptionSet resource.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myImage",
		Image{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &ImageProperties{
				StorageProfile: &ImageStorageProfile{
					OSDisk: &ImageOSDisk{
						ImageDisk: ImageDisk{
							DiskEncryptionSet: &DiskEncryptionSetParameters{
								SubResource: SubResource{
									ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
								},
							},
							ManagedDisk: &SubResource{
								ID: to.StringPtr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
							},
						},
						OSState: OperatingSystemStateTypesGeneralized.ToPtr(),
						OSType:  OperatingSystemTypesLinux.ToPtr(),
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a virtual machine image from a managed disk.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myImage",
		Image{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &ImageProperties{
				StorageProfile: &ImageStorageProfile{
					OSDisk: &ImageOSDisk{
						ImageDisk: ImageDisk{
							ManagedDisk: &SubResource{
								ID: to.StringPtr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
							},
						},
						OSState: OperatingSystemStateTypesGeneralized.ToPtr(),
						OSType:  OperatingSystemTypesLinux.ToPtr(),
					},
					ZoneResilient: to.BoolPtr(true),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a virtual machine image from a snapshot with DiskEncryptionSet resource.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myImage",
		Image{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &ImageProperties{
				StorageProfile: &ImageStorageProfile{
					OSDisk: &ImageOSDisk{
						ImageDisk: ImageDisk{
							DiskEncryptionSet: &DiskEncryptionSetParameters{
								SubResource: SubResource{
									ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
								},
							},
							Snapshot: &SubResource{
								ID: to.StringPtr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
							},
						},
						OSState: OperatingSystemStateTypesGeneralized.ToPtr(),
						OSType:  OperatingSystemTypesLinux.ToPtr(),
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a virtual machine image from a snapshot.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myImage",
		Image{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &ImageProperties{
				StorageProfile: &ImageStorageProfile{
					OSDisk: &ImageOSDisk{
						ImageDisk: ImageDisk{
							Snapshot: &SubResource{
								ID: to.StringPtr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
							},
						},
						OSState: OperatingSystemStateTypesGeneralized.ToPtr(),
						OSType:  OperatingSystemTypesLinux.ToPtr(),
					},
					ZoneResilient: to.BoolPtr(false),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a virtual machine image from an existing virtual machine.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myImage",
		Image{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &ImageProperties{
				SourceVirtualMachine: &SubResource{
					ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a virtual machine image that includes a data disk from a blob.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myImage",
		Image{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &ImageProperties{
				StorageProfile: &ImageStorageProfile{
					DataDisks: []*ImageDataDisk{
						&ImageDataDisk{
							ImageDisk: ImageDisk{
								BlobURI: to.StringPtr("https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd"),
							},
							Lun: to.Int32Ptr(1),
						}},
					OSDisk: &ImageOSDisk{
						ImageDisk: ImageDisk{
							BlobURI: to.StringPtr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
						},
						OSState: OperatingSystemStateTypesGeneralized.ToPtr(),
						OSType:  OperatingSystemTypesLinux.ToPtr(),
					},
					ZoneResilient: to.BoolPtr(false),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a virtual machine image that includes a data disk from a managed disk.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myImage",
		Image{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &ImageProperties{
				StorageProfile: &ImageStorageProfile{
					DataDisks: []*ImageDataDisk{
						&ImageDataDisk{
							ImageDisk: ImageDisk{
								ManagedDisk: &SubResource{
									ID: to.StringPtr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk2"),
								},
							},
							Lun: to.Int32Ptr(1),
						}},
					OSDisk: &ImageOSDisk{
						ImageDisk: ImageDisk{
							ManagedDisk: &SubResource{
								ID: to.StringPtr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
							},
						},
						OSState: OperatingSystemStateTypesGeneralized.ToPtr(),
						OSType:  OperatingSystemTypesLinux.ToPtr(),
					},
					ZoneResilient: to.BoolPtr(false),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a virtual machine image that includes a data disk from a snapshot.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myImage",
		Image{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &ImageProperties{
				StorageProfile: &ImageStorageProfile{
					DataDisks: []*ImageDataDisk{
						&ImageDataDisk{
							ImageDisk: ImageDisk{
								Snapshot: &SubResource{
									ID: to.StringPtr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot2"),
								},
							},
							Lun: to.Int32Ptr(1),
						}},
					OSDisk: &ImageOSDisk{
						ImageDisk: ImageDisk{
							Snapshot: &SubResource{
								ID: to.StringPtr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
							},
						},
						OSState: OperatingSystemStateTypesGeneralized.ToPtr(),
						OSType:  OperatingSystemTypesLinux.ToPtr(),
					},
					ZoneResilient: to.BoolPtr(true),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestImages_Update(t *testing.T) {

	// From example Updates tags of an Image.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewImagesClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myImage",
		ImageUpdate{
			UpdateResource: UpdateResource{
				Tags: map[string]*string{
					"department": to.StringPtr("HR"),
				},
			},
			Properties: &ImageProperties{
				HyperVGeneration: HyperVGenerationTypesV1.ToPtr(),
				SourceVirtualMachine: &SubResource{
					ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestImages_Delete(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestImages_Get(t *testing.T) {

	// From example Get information about a virtual machine image.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewImagesClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myImage",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestImages_ListByResourceGroup(t *testing.T) {

	// From example List all virtual machine images in a resource group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewImagesClient(con,
		"{subscription-id}")
	client.ListByResourceGroup("myResourceGroup",
		nil)

}

func TestImages_List(t *testing.T) {

	// From example List all virtual machine images in a subscription.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewImagesClient(con,
		"{subscription-id}")
	client.List(nil)

}

func TestRestorePointCollections_CreateOrUpdate(t *testing.T) {

	// From example Create or update a restore point collection.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRestorePointCollectionsClient(con,
		"{subscription-id}")
	_, err := client.CreateOrUpdate(ctx,
		"myResourceGroup",
		"myRpc",
		RestorePointCollection{
			Resource: Resource{
				Location: to.StringPtr("norwayeast"),
				Tags: map[string]*string{
					"myTag1": to.StringPtr("tagValue1"),
				},
			},
			Properties: &RestorePointCollectionProperties{
				Source: &RestorePointCollectionSourceProperties{
					ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestRestorePointCollections_Update(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestRestorePointCollections_Delete(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestRestorePointCollections_Get(t *testing.T) {

	// From example Get a restore point collection (but not the restore points contained in the restore point collection)
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRestorePointCollectionsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myRpc",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	// From example Get a restore point collection, including the restore points contained in the restore point collection
	_, err = client.Get(ctx,
		"myResourceGroup",
		"rpcName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestRestorePointCollections_List(t *testing.T) {

	// From example Gets the list of restore point collections in a resource group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRestorePointCollectionsClient(con,
		"{subscription-id}")
	client.List("myResourceGroup",
		nil)

}

func TestRestorePointCollections_ListAll(t *testing.T) {

	// From example Gets the list of restore point collections in a subscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRestorePointCollectionsClient(con,
		"{subscription-id}")
	client.ListAll(nil)

}

func TestRestorePoints_Create(t *testing.T) {

	// From example Create a restore point
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRestorePointsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreate(ctx,
		"myResourceGroup",
		"rpcName",
		"rpName",
		RestorePoint{
			ExcludeDisks: []*APIEntityReference{
				&APIEntityReference{
					ID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm8768_disk2_fe6ffde4f69b491ca33fb984d5bcd89f"),
				}},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestRestorePoints_Delete(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestRestorePoints_Get(t *testing.T) {

	// From example Get a restore point
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewRestorePointsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"rpcName",
		"rpName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetExtensions_CreateOrUpdate(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetExtensions_Update(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetExtensions_Delete(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetExtensions_Get(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetExtensions_List(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetRollingUpgrades_Cancel(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetRollingUpgrades_StartOSUpgrade(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetRollingUpgrades_StartExtensionUpgrade(t *testing.T) {

	// From example Start an extension rolling upgrade.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetRollingUpgradesClient(con,
		"{subscription-id}")
	poller, err := client.BeginStartExtensionUpgrade(ctx,
		"myResourceGroup",
		"{vmss-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetRollingUpgrades_GetLatest(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetVMExtensions_CreateOrUpdate(t *testing.T) {

	// From example Create VirtualMachineScaleSet VM extension.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMExtensionsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		"0",
		"myVMExtension",
		VirtualMachineScaleSetVMExtension{
			Properties: &VirtualMachineExtensionProperties{
				Type:                    to.StringPtr("extType"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Publisher:               to.StringPtr("extPublisher"),
				Settings: map[string]interface{}{
					"UserName": `"xyz@microsoft.com"`,
				},
				TypeHandlerVersion: to.StringPtr("1.2"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMExtensions_Update(t *testing.T) {

	// From example Update VirtualMachineScaleSet VM extension.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMExtensionsClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		"0",
		"myVMExtension",
		VirtualMachineScaleSetVMExtensionUpdate{
			Properties: &VirtualMachineExtensionUpdateProperties{
				Type:                    to.StringPtr("extType"),
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				Publisher:               to.StringPtr("extPublisher"),
				Settings: map[string]interface{}{
					"UserName": `"xyz@microsoft.com"`,
				},
				TypeHandlerVersion: to.StringPtr("1.2"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMExtensions_Delete(t *testing.T) {

	// From example Delete VirtualMachineScaleSet VM extension.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMExtensionsClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		"0",
		"myVMExtension",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMExtensions_Get(t *testing.T) {

	// From example Get VirtualMachineScaleSet VM extension.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMExtensionsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		"0",
		"myVMExtension",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMExtensions_List(t *testing.T) {

	// From example List extensions in Vmss instance.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMExtensionsClient(con,
		"{subscription-id}")
	_, err := client.List(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		"0",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMs_Reimage(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetVMs_ReimageAll(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetVMs_Deallocate(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetVMs_Update(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetVMs_Delete(t *testing.T) {

	// From example Force Delete a virtual machine from a VM scale set.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMsClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		"0",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMs_Get(t *testing.T) {

	// From example Get VM scale set VM with UserData
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"{vmss-name}",
		"0",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMs_GetInstanceView(t *testing.T) {

	// From example Get instance view of a virtual machine from a VM scale set placed on a dedicated host group through automatic placement.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMsClient(con,
		"{subscription-id}")
	_, err := client.GetInstanceView(ctx,
		"myResourceGroup",
		"myVirtualMachineScaleSet",
		"0",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMs_List(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetVMs_PowerOff(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetVMs_Restart(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetVMs_Start(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetVMs_Redeploy(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetVMs_RetrieveBootDiagnosticsData(t *testing.T) {

	// From example RetrieveBootDiagnosticsData of a virtual machine.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMsClient(con,
		"{subscription-id}")
	_, err := client.RetrieveBootDiagnosticsData(ctx,
		"ResourceGroup",
		"myvmScaleSet",
		"0",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMs_PerformMaintenance(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestVirtualMachineScaleSetVMs_SimulateEviction(t *testing.T) {

	// From example Simulate Eviction a virtual machine.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMsClient(con,
		"{subscription-id}")
	_, err := client.SimulateEviction(ctx,
		"ResourceGroup",
		"VmScaleSetName",
		"InstanceId",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMs_RunCommand(t *testing.T) {

	// From example VirtualMachineScaleSetVMs_RunCommand
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMsClient(con,
		"{subscription-id}")
	poller, err := client.BeginRunCommand(ctx,
		"myResourceGroup",
		"myVirtualMachineScaleSet",
		"0",
		RunCommandInput{
			CommandID: to.StringPtr("RunPowerShellScript"),
			Script: []*string{
				to.StringPtr("Write-Host Hello World!")},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLogAnalytics_ExportRequestRateByInterval(t *testing.T) {

	// From example Export logs which contain all Api requests made to Compute Resource Provider within the given time period broken down by intervals.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewLogAnalyticsClient(con,
		"{subscription-id}")
	poller, err := client.BeginExportRequestRateByInterval(ctx,
		"westus",
		RequestRateByIntervalInput{
			LogAnalyticsInputBase: LogAnalyticsInputBase{
				BlobContainerSasURI: to.StringPtr("https://somesasuri"),
				FromTime:            to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339, "2018-01-21T01:54:06.862601Z"); return t }()),
				GroupByResourceName: to.BoolPtr(true),
				ToTime:              to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339, "2018-01-23T01:54:06.862601Z"); return t }()),
			},
			IntervalLength: IntervalInMinsFiveMins.ToPtr(),
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestLogAnalytics_ExportThrottledRequests(t *testing.T) {

	// From example Export logs which contain all throttled Api requests made to Compute Resource Provider within the given time period.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewLogAnalyticsClient(con,
		"{subscription-id}")
	poller, err := client.BeginExportThrottledRequests(ctx,
		"westus",
		ThrottledRequestsInput{
			LogAnalyticsInputBase: LogAnalyticsInputBase{
				BlobContainerSasURI:        to.StringPtr("https://somesasuri"),
				FromTime:                   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339, "2018-01-21T01:54:06.862601Z"); return t }()),
				GroupByClientApplicationID: to.BoolPtr(false),
				GroupByOperationName:       to.BoolPtr(true),
				GroupByResourceName:        to.BoolPtr(false),
				GroupByUserAgent:           to.BoolPtr(false),
				ToTime:                     to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339, "2018-01-23T01:54:06.862601Z"); return t }()),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineRunCommands_List(t *testing.T) {

	// From example VirtualMachineRunCommandList
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineRunCommandsClient(con,
		"subid")
	client.List("SoutheastAsia",
		nil)

}

func TestVirtualMachineRunCommands_Get(t *testing.T) {

	// From example VirtualMachineRunCommandGet
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineRunCommandsClient(con,
		"24fb23e3-6ba3-41f0-9b6e-e41131d5d61e")
	_, err := client.Get(ctx,
		"SoutheastAsia",
		"RunPowerShellScript",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineRunCommands_CreateOrUpdate(t *testing.T) {

	// From example Create or update a run command.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineRunCommandsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myVM",
		"myRunCommand",
		VirtualMachineRunCommand{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &VirtualMachineRunCommandProperties{
				AsyncExecution: to.BoolPtr(false),
				Parameters: []*RunCommandInputParameter{
					&RunCommandInputParameter{
						Name:  to.StringPtr("param1"),
						Value: to.StringPtr("value1"),
					},
					&RunCommandInputParameter{
						Name:  to.StringPtr("param2"),
						Value: to.StringPtr("value2"),
					}},
				RunAsPassword: to.StringPtr("<runAsPassword>"),
				RunAsUser:     to.StringPtr("user1"),
				Source: &VirtualMachineRunCommandScriptSource{
					Script: to.StringPtr("Write-Host Hello World!"),
				},
				TimeoutInSeconds: to.Int32Ptr(3600),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineRunCommands_Update(t *testing.T) {

	// From example Update a run command.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineRunCommandsClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myVM",
		"myRunCommand",
		VirtualMachineRunCommandUpdate{
			Properties: &VirtualMachineRunCommandProperties{
				Source: &VirtualMachineRunCommandScriptSource{
					Script: to.StringPtr("Write-Host Script Source Updated!"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineRunCommands_Delete(t *testing.T) {

	// From example Delete a run command.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineRunCommandsClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myVM",
		"myRunCommand",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineRunCommands_GetByVirtualMachine(t *testing.T) {

	// From example Get a run command.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineRunCommandsClient(con,
		"{subscription-id}")
	_, err := client.GetByVirtualMachine(ctx,
		"myResourceGroup",
		"myVM",
		"myRunCommand",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineRunCommands_ListByVirtualMachine(t *testing.T) {

	// From example List run commands in a Virtual Machine.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineRunCommandsClient(con,
		"{subscription-id}")
	client.ListByVirtualMachine("myResourceGroup",
		"myVM",
		nil)

}

func TestVirtualMachineScaleSetVMRunCommands_CreateOrUpdate(t *testing.T) {

	// From example Create VirtualMachineScaleSet VM run command.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMRunCommandsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		"0",
		"myRunCommand",
		VirtualMachineRunCommand{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &VirtualMachineRunCommandProperties{
				AsyncExecution: to.BoolPtr(false),
				Parameters: []*RunCommandInputParameter{
					&RunCommandInputParameter{
						Name:  to.StringPtr("param1"),
						Value: to.StringPtr("value1"),
					},
					&RunCommandInputParameter{
						Name:  to.StringPtr("param2"),
						Value: to.StringPtr("value2"),
					}},
				RunAsPassword: to.StringPtr("<runAsPassword>"),
				RunAsUser:     to.StringPtr("user1"),
				Source: &VirtualMachineRunCommandScriptSource{
					Script: to.StringPtr("Write-Host Hello World!"),
				},
				TimeoutInSeconds: to.Int32Ptr(3600),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMRunCommands_Update(t *testing.T) {

	// From example Update VirtualMachineScaleSet VM run command.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMRunCommandsClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		"0",
		"myRunCommand",
		VirtualMachineRunCommandUpdate{
			Properties: &VirtualMachineRunCommandProperties{
				Source: &VirtualMachineRunCommandScriptSource{
					Script: to.StringPtr("Write-Host Script Source Updated!"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMRunCommands_Delete(t *testing.T) {

	// From example Delete VirtualMachineScaleSet VM run command.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMRunCommandsClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		"0",
		"myRunCommand",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMRunCommands_Get(t *testing.T) {

	// From example Get VirtualMachineScaleSet VM run commands.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMRunCommandsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myvmScaleSet",
		"0",
		"myRunCommand",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestVirtualMachineScaleSetVMRunCommands_List(t *testing.T) {

	// From example List run commands in Vmss instance.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewVirtualMachineScaleSetVMRunCommandsClient(con,
		"{subscription-id}")
	client.List("myResourceGroup",
		"myvmScaleSet",
		"0",
		nil)

}

func TestResourceSkus_List(t *testing.T) {

	// From example Lists all available Resource SKUs
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewResourceSKUsClient(con,
		"{subscription-id}")
	client.List(nil)

	// From example Lists all available Resource SKUs for the specified region
	client.List(nil)

}

func TestDisks_CreateOrUpdate(t *testing.T) {

	// From example Create a managed disk and associate with disk access resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDisksClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption: DiskCreateOptionEmpty.ToPtr(),
				},
				DiskAccessID:        to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskAccesses/{existing-diskAccess-name}"),
				DiskSizeGB:          to.Int32Ptr(200),
				NetworkAccessPolicy: NetworkAccessPolicyAllowPrivate.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a managed disk and associate with disk encryption set.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption: DiskCreateOptionEmpty.ToPtr(),
				},
				DiskSizeGB: to.Int32Ptr(200),
				Encryption: &Encryption{
					DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a managed disk by copying a snapshot.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption:     DiskCreateOptionCopy.ToPtr(),
					SourceResourceID: to.StringPtr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a managed disk by importing an unmanaged blob from a different subscription.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption:     DiskCreateOptionImport.ToPtr(),
					SourceURI:        to.StringPtr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
					StorageAccountID: to.StringPtr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a managed disk by importing an unmanaged blob from the same subscription.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption: DiskCreateOptionImport.ToPtr(),
					SourceURI:    to.StringPtr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a managed disk from a platform image.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption: DiskCreateOptionFromImage.ToPtr(),
					ImageReference: &ImageDiskReference{
						ID: to.StringPtr("/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/westus/Publishers/{publisher}/ArtifactTypes/VMImage/Offers/{offer}/Skus/{sku}/Versions/1.0.0"),
					},
				},
				OSType: OperatingSystemTypesWindows.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a managed disk from an existing managed disk in the same or different subscription.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk2",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption:     DiskCreateOptionCopy.ToPtr(),
					SourceResourceID: to.StringPtr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk1"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a managed disk with security profile
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("North Central US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption: DiskCreateOptionFromImage.ToPtr(),
					ImageReference: &ImageDiskReference{
						ID: to.StringPtr("/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/uswest/Publishers/Microsoft/ArtifactTypes/VMImage/Offers/{offer}"),
					},
				},
				OSType: OperatingSystemTypesWindows.ToPtr(),
				SecurityProfile: &DiskSecurityProfile{
					SecurityType: DiskSecurityTypesTrustedLaunch.ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a managed disk with ssd zrs account type.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption: DiskCreateOptionEmpty.ToPtr(),
				},
				DiskSizeGB: to.Int32Ptr(200),
			},
			SKU: &DiskSKU{
				Name: DiskStorageAccountTypesPremiumZRS.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a managed upload disk.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption:    DiskCreateOptionUpload.ToPtr(),
					UploadSizeBytes: to.Int64Ptr(10737418752),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create an empty managed disk in extended location.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			ExtendedLocation: &ExtendedLocation{
				Name: to.StringPtr("{edge-zone-id}"),
				Type: ExtendedLocationTypesEdgeZone.ToPtr(),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption: DiskCreateOptionEmpty.ToPtr(),
				},
				DiskSizeGB: to.Int32Ptr(200),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create an empty managed disk.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption: DiskCreateOptionEmpty.ToPtr(),
				},
				DiskSizeGB: to.Int32Ptr(200),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create an ultra managed disk with logicalSectorSize 512E
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		Disk{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &DiskProperties{
				CreationData: &CreationData{
					CreateOption:      DiskCreateOptionEmpty.ToPtr(),
					LogicalSectorSize: to.Int32Ptr(512),
				},
				DiskSizeGB: to.Int32Ptr(200),
			},
			SKU: &DiskSKU{
				Name: DiskStorageAccountTypesUltraSSDLRS.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDisks_Update(t *testing.T) {

	// From example Create or update a bursting enabled managed disk.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDisksClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		DiskUpdate{
			Properties: &DiskUpdateProperties{
				BurstingEnabled: to.BoolPtr(true),
				DiskSizeGB:      to.Int32Ptr(1024),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Update a managed disk to add purchase plan.
	poller, err = client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		DiskUpdate{
			Properties: &DiskUpdateProperties{
				PurchasePlan: &DiskPurchasePlan{
					Name:          to.StringPtr("myPurchasePlanName"),
					Product:       to.StringPtr("myPurchasePlanProduct"),
					PromotionCode: to.StringPtr("myPurchasePlanPromotionCode"),
					Publisher:     to.StringPtr("myPurchasePlanPublisher"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Update a managed disk to add supportsHibernation.
	poller, err = client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		DiskUpdate{
			Properties: &DiskUpdateProperties{
				SupportsHibernation: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Update a managed disk to change tier.
	poller, err = client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		DiskUpdate{
			Properties: &DiskUpdateProperties{
				Tier: to.StringPtr("P30"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Update a managed disk to disable bursting.
	poller, err = client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		DiskUpdate{
			Properties: &DiskUpdateProperties{
				BurstingEnabled: to.BoolPtr(false),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Update managed disk to remove disk access resource association.
	poller, err = client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDisk",
		DiskUpdate{
			Properties: &DiskUpdateProperties{
				NetworkAccessPolicy: NetworkAccessPolicyAllowAll.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDisks_Get(t *testing.T) {

	// From example Get information about a managed disk.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDisksClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myManagedDisk",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDisks_Delete(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestDisks_ListByResourceGroup(t *testing.T) {

	// From example List all managed disks in a resource group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDisksClient(con,
		"{subscription-id}")
	client.ListByResourceGroup("myResourceGroup",
		nil)

}

func TestDisks_List(t *testing.T) {

	// From example List all managed disks in a subscription.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDisksClient(con,
		"{subscription-id}")
	client.List(nil)

}

func TestDisks_GrantAccess(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestDisks_RevokeAccess(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestSnapshots_CreateOrUpdate(t *testing.T) {

	// From example Create a snapshot by importing an unmanaged blob from a different subscription.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSnapshotsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"mySnapshot1",
		Snapshot{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &SnapshotProperties{
				CreationData: &CreationData{
					CreateOption:     DiskCreateOptionImport.ToPtr(),
					SourceURI:        to.StringPtr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
					StorageAccountID: to.StringPtr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a snapshot by importing an unmanaged blob from the same subscription.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"mySnapshot1",
		Snapshot{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &SnapshotProperties{
				CreationData: &CreationData{
					CreateOption: DiskCreateOptionImport.ToPtr(),
					SourceURI:    to.StringPtr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a snapshot from an existing snapshot in the same or a different subscription.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"mySnapshot2",
		Snapshot{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &SnapshotProperties{
				CreationData: &CreationData{
					CreateOption:     DiskCreateOptionCopy.ToPtr(),
					SourceResourceID: to.StringPtr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1"),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSnapshots_Update(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestSnapshots_Get(t *testing.T) {

	// From example Get information about a snapshot.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSnapshotsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"mySnapshot",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSnapshots_Delete(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestSnapshots_ListByResourceGroup(t *testing.T) {

	// From example List all snapshots in a resource group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSnapshotsClient(con,
		"{subscription-id}")
	client.ListByResourceGroup("myResourceGroup",
		nil)

}

func TestSnapshots_List(t *testing.T) {

	// From example List all snapshots in a subscription.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSnapshotsClient(con,
		"{subscription-id}")
	client.List(nil)

}

func TestSnapshots_GrantAccess(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestSnapshots_RevokeAccess(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestDiskEncryptionSets_CreateOrUpdate(t *testing.T) {

	// From example Create a disk encryption set with key vault from a different subscription.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskEncryptionSetsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDiskEncryptionSet",
		DiskEncryptionSet{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Identity: &EncryptionSetIdentity{
				Type: DiskEncryptionSetIdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &EncryptionSetProperties{
				ActiveKey: &KeyForDiskEncryptionSet{
					KeyURL: to.StringPtr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/{key}"),
				},
				EncryptionType: DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create a disk encryption set.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDiskEncryptionSet",
		DiskEncryptionSet{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Identity: &EncryptionSetIdentity{
				Type: DiskEncryptionSetIdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &EncryptionSetProperties{
				ActiveKey: &KeyForDiskEncryptionSet{
					KeyURL: to.StringPtr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
					SourceVault: &SourceVault{
						ID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
					},
				},
				EncryptionType: DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskEncryptionSets_Update(t *testing.T) {

	// From example Update a disk encryption set with rotationToLatestKeyVersionEnabled set to true - Succeeded
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskEncryptionSetsClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDiskEncryptionSet",
		DiskEncryptionSetUpdate{
			Identity: &EncryptionSetIdentity{
				Type: DiskEncryptionSetIdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &DiskEncryptionSetUpdateProperties{
				ActiveKey: &KeyForDiskEncryptionSet{
					KeyURL: to.StringPtr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/keyName/keyVersion1"),
				},
				EncryptionType:                    DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey.ToPtr(),
				RotationToLatestKeyVersionEnabled: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Update a disk encryption set with rotationToLatestKeyVersionEnabled set to true - Updating
	poller, err = client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDiskEncryptionSet",
		DiskEncryptionSetUpdate{
			Identity: &EncryptionSetIdentity{
				Type: DiskEncryptionSetIdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &DiskEncryptionSetUpdateProperties{
				ActiveKey: &KeyForDiskEncryptionSet{
					KeyURL: to.StringPtr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/keyName/keyVersion1"),
				},
				EncryptionType:                    DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey.ToPtr(),
				RotationToLatestKeyVersionEnabled: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Update a disk encryption set.
	poller, err = client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDiskEncryptionSet",
		DiskEncryptionSetUpdate{
			Properties: &DiskEncryptionSetUpdateProperties{
				ActiveKey: &KeyForDiskEncryptionSet{
					KeyURL: to.StringPtr("https://myvmvault.vault-int.azure-int.net/keys/keyName/keyVersion"),
					SourceVault: &SourceVault{
						ID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
					},
				},
				EncryptionType: DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey.ToPtr(),
			},
			Tags: map[string]*string{
				"department": to.StringPtr("Development"),
				"project":    to.StringPtr("Encryption"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskEncryptionSets_Get(t *testing.T) {

	// From example Get information about a disk encryption set.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskEncryptionSetsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myDiskEncryptionSet",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskEncryptionSets_Delete(t *testing.T) {

	// From example Delete a disk encryption set.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskEncryptionSetsClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myDiskEncryptionSet",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskEncryptionSets_ListByResourceGroup(t *testing.T) {

	// From example List all disk encryption sets in a resource group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskEncryptionSetsClient(con,
		"{subscription-id}")
	client.ListByResourceGroup("myResourceGroup",
		nil)

}

func TestDiskEncryptionSets_List(t *testing.T) {

	// From example List all disk encryption sets in a subscription.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskEncryptionSetsClient(con,
		"{subscription-id}")
	client.List(nil)

}

func TestDiskEncryptionSets_ListAssociatedResources(t *testing.T) {

	// From example List all resources that are encrypted with this disk encryption set.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskEncryptionSetsClient(con,
		"{subscription-id}")
	client.ListAssociatedResources("myResourceGroup",
		"myDiskEncryptionSet",
		nil)

}

func TestDiskAccesses_CreateOrUpdate(t *testing.T) {

	// From example Create a disk access resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskAccessesClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myDiskAccess",
		DiskAccess{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskAccesses_Update(t *testing.T) {

	// From example Update a disk access resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskAccessesClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myDiskAccess",
		DiskAccessUpdate{
			Tags: map[string]*string{
				"department": to.StringPtr("Development"),
				"project":    to.StringPtr("PrivateEndpoints"),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskAccesses_Get(t *testing.T) {

	// From example Get information about a disk access resource with private endpoints.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskAccessesClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myDiskAccess",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	// From example Get information about a disk access resource.
	_, err = client.Get(ctx,
		"myResourceGroup",
		"myDiskAccess",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskAccesses_Delete(t *testing.T) {

	// From example Delete a disk access resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskAccessesClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myDiskAccess",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskAccesses_ListByResourceGroup(t *testing.T) {

	// From example List all disk access resources in a resource group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskAccessesClient(con,
		"{subscription-id}")
	client.ListByResourceGroup("myResourceGroup",
		nil)

}

func TestDiskAccesses_List(t *testing.T) {

	// From example List all disk access resources in a subscription.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskAccessesClient(con,
		"{subscription-id}")
	client.List(nil)

}

func TestDiskAccesses_GetPrivateLinkResources(t *testing.T) {

	// From example List all possible private link resources under disk access resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskAccessesClient(con,
		"{subscription-id}")
	_, err := client.GetPrivateLinkResources(ctx,
		"myResourceGroup",
		"myDiskAccess",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskAccesses_UpdateAPrivateEndpointConnection(t *testing.T) {

	// From example Approve a Private Endpoint Connection under a disk access resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskAccessesClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdateAPrivateEndpointConnection(ctx,
		"myResourceGroup",
		"myDiskAccess",
		"myPrivateEndpointConnection",
		PrivateEndpointConnection{
			Properties: &PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &PrivateLinkServiceConnectionState{
					Description: to.StringPtr("Approving myPrivateEndpointConnection"),
					Status:      PrivateEndpointServiceConnectionStatusApproved.ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskAccesses_GetAPrivateEndpointConnection(t *testing.T) {

	// From example Get information about a private endpoint connection under a disk access resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskAccessesClient(con,
		"{subscription-id}")
	_, err := client.GetAPrivateEndpointConnection(ctx,
		"myResourceGroup",
		"myDiskAccess",
		"myPrivateEndpointConnection",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskAccesses_DeleteAPrivateEndpointConnection(t *testing.T) {

	// From example Delete a private endpoint connection under a disk access resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskAccessesClient(con,
		"{subscription-id}")
	poller, err := client.BeginDeleteAPrivateEndpointConnection(ctx,
		"myResourceGroup",
		"myDiskAccess",
		"myPrivateEndpointConnection",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskAccesses_ListPrivateEndpointConnections(t *testing.T) {

	// From example Get information about a private endpoint connection under a disk access resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskAccessesClient(con,
		"{subscription-id}")
	client.ListPrivateEndpointConnections("myResourceGroup",
		"myDiskAccess",
		nil)

}

func TestDiskRestorePoint_Get(t *testing.T) {

	// From example Get an incremental disk restorePoint resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskRestorePointClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"rpc",
		"vmrp",
		"TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestDiskRestorePoint_ListByRestorePoint(t *testing.T) {

	// From example Get an incremental disk restorePoint resource.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewDiskRestorePointClient(con,
		"{subscription-id}")
	client.ListByRestorePoint("myResourceGroup",
		"rpc",
		"vmrp",
		nil)

}

func TestGalleries_CreateOrUpdate(t *testing.T) {

	// From example Create or update a simple gallery with sharing profile.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleriesClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		Gallery{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &GalleryProperties{
				Description: to.StringPtr("This is the gallery description."),
				SharingProfile: &SharingProfile{
					Permissions: GallerySharingPermissionTypesGroups.ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create or update a simple gallery.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		Gallery{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &GalleryProperties{
				Description: to.StringPtr("This is the gallery description."),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleries_Update(t *testing.T) {

	// From example Update a simple gallery.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleriesClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		GalleryUpdate{
			Properties: &GalleryProperties{
				Description: to.StringPtr("This is the gallery description."),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleries_Get(t *testing.T) {

	// From example Get a gallery with select permissions.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleriesClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myGalleryName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	// From example Get a gallery.
	_, err = client.Get(ctx,
		"myResourceGroup",
		"myGalleryName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleries_Delete(t *testing.T) {

	// From example Delete a gallery.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleriesClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myGalleryName",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleries_ListByResourceGroup(t *testing.T) {

	// From example List galleries in a resource group.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleriesClient(con,
		"{subscription-id}")
	client.ListByResourceGroup("myResourceGroup",
		nil)

}

func TestGalleries_List(t *testing.T) {

	// From example List galleries in a subscription.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleriesClient(con,
		"{subscription-id}")
	client.List(nil)

}

func TestGalleryImages_CreateOrUpdate(t *testing.T) {

	// From example Create or update a simple gallery image.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryImagesClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		GalleryImage{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &GalleryImageProperties{
				HyperVGeneration: HyperVGenerationV1.ToPtr(),
				Identifier: &GalleryImageIdentifier{
					Offer:     to.StringPtr("myOfferName"),
					Publisher: to.StringPtr("myPublisherName"),
					SKU:       to.StringPtr("mySkuName"),
				},
				OSState: OperatingSystemStateTypesGeneralized.ToPtr(),
				OSType:  OperatingSystemTypesWindows.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryImages_Update(t *testing.T) {

	// From example Update a simple gallery image.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryImagesClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		GalleryImageUpdate{
			Properties: &GalleryImageProperties{
				HyperVGeneration: HyperVGenerationV1.ToPtr(),
				Identifier: &GalleryImageIdentifier{
					Offer:     to.StringPtr("myOfferName"),
					Publisher: to.StringPtr("myPublisherName"),
					SKU:       to.StringPtr("mySkuName"),
				},
				OSState: OperatingSystemStateTypesGeneralized.ToPtr(),
				OSType:  OperatingSystemTypesWindows.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryImages_Get(t *testing.T) {

	// From example Get a gallery image.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryImagesClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryImages_Delete(t *testing.T) {

	// From example Delete a gallery image.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryImagesClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryImages_ListByGallery(t *testing.T) {

	// From example List gallery images in a gallery.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryImagesClient(con,
		"{subscription-id}")
	client.ListByGallery("myResourceGroup",
		"myGalleryName",
		nil)

}

func TestGalleryImageVersions_CreateOrUpdate(t *testing.T) {

	// From example Create or update a simple Gallery Image Version using VM as source.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryImageVersionsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		GalleryImageVersion{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &GalleryImageVersionProperties{
				PublishingProfile: &GalleryImageVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: GalleryArtifactPublishingProfileBase{
						TargetRegions: []*TargetRegion{
							&TargetRegion{
								Name: to.StringPtr("West US"),
								Encryption: &EncryptionImages{
									DataDiskImages: []*DataDiskImageEncryption{
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherWestUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(0),
										},
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &OSDiskImageEncryption{
										DiskImageEncryption: DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(1),
							},
							&TargetRegion{
								Name: to.StringPtr("East US"),
								Encryption: &EncryptionImages{
									DataDiskImages: []*DataDiskImageEncryption{
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherEastUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(0),
										},
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &OSDiskImageEncryption{
										DiskImageEncryption: DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(2),
								StorageAccountType:   StorageAccountTypeStandardZRS.ToPtr(),
							}},
					},
				},
				StorageProfile: &GalleryImageVersionStorageProfile{
					Source: &GalleryArtifactVersionSource{
						ID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create or update a simple Gallery Image Version using managed image as source.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		GalleryImageVersion{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &GalleryImageVersionProperties{
				PublishingProfile: &GalleryImageVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: GalleryArtifactPublishingProfileBase{
						TargetRegions: []*TargetRegion{
							&TargetRegion{
								Name: to.StringPtr("West US"),
								Encryption: &EncryptionImages{
									DataDiskImages: []*DataDiskImageEncryption{
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherWestUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(0),
										},
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &OSDiskImageEncryption{
										DiskImageEncryption: DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(1),
							},
							&TargetRegion{
								Name: to.StringPtr("East US"),
								Encryption: &EncryptionImages{
									DataDiskImages: []*DataDiskImageEncryption{
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherEastUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(0),
										},
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &OSDiskImageEncryption{
										DiskImageEncryption: DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(2),
								StorageAccountType:   StorageAccountTypeStandardZRS.ToPtr(),
							}},
					},
				},
				StorageProfile: &GalleryImageVersionStorageProfile{
					Source: &GalleryArtifactVersionSource{
						ID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"),
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create or update a simple Gallery Image Version using mix of disks and snapshots as a source.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		GalleryImageVersion{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &GalleryImageVersionProperties{
				PublishingProfile: &GalleryImageVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: GalleryArtifactPublishingProfileBase{
						TargetRegions: []*TargetRegion{
							&TargetRegion{
								Name: to.StringPtr("West US"),
								Encryption: &EncryptionImages{
									DataDiskImages: []*DataDiskImageEncryption{
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &OSDiskImageEncryption{
										DiskImageEncryption: DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(1),
							},
							&TargetRegion{
								Name: to.StringPtr("East US"),
								Encryption: &EncryptionImages{
									DataDiskImages: []*DataDiskImageEncryption{
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &OSDiskImageEncryption{
										DiskImageEncryption: DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(2),
								StorageAccountType:   StorageAccountTypeStandardZRS.ToPtr(),
							}},
					},
				},
				StorageProfile: &GalleryImageVersionStorageProfile{
					DataDiskImages: []*GalleryDataDiskImage{
						&GalleryDataDiskImage{
							GalleryDiskImage: GalleryDiskImage{
								HostCaching: HostCachingNone.ToPtr(),
								Source: &GalleryArtifactVersionSource{
									ID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/disks/{dataDiskName}"),
								},
							},
							Lun: to.Int32Ptr(1),
						}},
					OSDiskImage: &GalleryOSDiskImage{
						GalleryDiskImage: GalleryDiskImage{
							HostCaching: HostCachingReadOnly.ToPtr(),
							Source: &GalleryArtifactVersionSource{
								ID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/snapshots/{osSnapshotName}"),
							},
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create or update a simple Gallery Image Version using shared image as source.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		GalleryImageVersion{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &GalleryImageVersionProperties{
				PublishingProfile: &GalleryImageVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: GalleryArtifactPublishingProfileBase{
						TargetRegions: []*TargetRegion{
							&TargetRegion{
								Name: to.StringPtr("West US"),
								Encryption: &EncryptionImages{
									DataDiskImages: []*DataDiskImageEncryption{
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherWestUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(0),
										},
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &OSDiskImageEncryption{
										DiskImageEncryption: DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(1),
							},
							&TargetRegion{
								Name: to.StringPtr("East US"),
								Encryption: &EncryptionImages{
									DataDiskImages: []*DataDiskImageEncryption{
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherEastUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(0),
										},
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &OSDiskImageEncryption{
										DiskImageEncryption: DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(2),
								StorageAccountType:   StorageAccountTypeStandardZRS.ToPtr(),
							}},
					},
				},
				StorageProfile: &GalleryImageVersionStorageProfile{
					Source: &GalleryArtifactVersionSource{
						ID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/galleries/{galleryName}/images/{imageDefinitionName}/versions/{versionName}"),
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create or update a simple Gallery Image Version using snapshots as a source.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		GalleryImageVersion{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &GalleryImageVersionProperties{
				PublishingProfile: &GalleryImageVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: GalleryArtifactPublishingProfileBase{
						TargetRegions: []*TargetRegion{
							&TargetRegion{
								Name: to.StringPtr("West US"),
								Encryption: &EncryptionImages{
									DataDiskImages: []*DataDiskImageEncryption{
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &OSDiskImageEncryption{
										DiskImageEncryption: DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(1),
							},
							&TargetRegion{
								Name: to.StringPtr("East US"),
								Encryption: &EncryptionImages{
									DataDiskImages: []*DataDiskImageEncryption{
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &OSDiskImageEncryption{
										DiskImageEncryption: DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(2),
								StorageAccountType:   StorageAccountTypeStandardZRS.ToPtr(),
							}},
					},
				},
				StorageProfile: &GalleryImageVersionStorageProfile{
					DataDiskImages: []*GalleryDataDiskImage{
						&GalleryDataDiskImage{
							GalleryDiskImage: GalleryDiskImage{
								HostCaching: HostCachingNone.ToPtr(),
								Source: &GalleryArtifactVersionSource{
									ID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/disks/{dataDiskName}"),
								},
							},
							Lun: to.Int32Ptr(1),
						}},
					OSDiskImage: &GalleryOSDiskImage{
						GalleryDiskImage: GalleryDiskImage{
							HostCaching: HostCachingReadOnly.ToPtr(),
							Source: &GalleryArtifactVersionSource{
								ID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/snapshots/{osSnapshotName}"),
							},
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create or update a simple Gallery Image Version using vhd as a source.
	poller, err = client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		GalleryImageVersion{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &GalleryImageVersionProperties{
				PublishingProfile: &GalleryImageVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: GalleryArtifactPublishingProfileBase{
						TargetRegions: []*TargetRegion{
							&TargetRegion{
								Name: to.StringPtr("West US"),
								Encryption: &EncryptionImages{
									DataDiskImages: []*DataDiskImageEncryption{
										&DataDiskImageEncryption{
											DiskImageEncryption: DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &OSDiskImageEncryption{
										DiskImageEncryption: DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(1),
							},
							&TargetRegion{
								Name:                 to.StringPtr("East US"),
								RegionalReplicaCount: to.Int32Ptr(2),
								StorageAccountType:   StorageAccountTypeStandardZRS.ToPtr(),
							}},
					},
				},
				StorageProfile: &GalleryImageVersionStorageProfile{
					DataDiskImages: []*GalleryDataDiskImage{
						&GalleryDataDiskImage{
							GalleryDiskImage: GalleryDiskImage{
								HostCaching: HostCachingNone.ToPtr(),
								Source: &GalleryArtifactVersionSource{
									ID:  to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"),
									URI: to.StringPtr("https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd"),
								},
							},
							Lun: to.Int32Ptr(1),
						}},
					OSDiskImage: &GalleryOSDiskImage{
						GalleryDiskImage: GalleryDiskImage{
							HostCaching: HostCachingReadOnly.ToPtr(),
							Source: &GalleryArtifactVersionSource{
								ID:  to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"),
								URI: to.StringPtr("https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd"),
							},
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryImageVersions_Update(t *testing.T) {

	// From example Update a simple Gallery Image Version (Managed Image as source).
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryImageVersionsClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		GalleryImageVersionUpdate{
			Properties: &GalleryImageVersionProperties{
				PublishingProfile: &GalleryImageVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: GalleryArtifactPublishingProfileBase{
						TargetRegions: []*TargetRegion{
							&TargetRegion{
								Name:                 to.StringPtr("West US"),
								RegionalReplicaCount: to.Int32Ptr(1),
							},
							&TargetRegion{
								Name:                 to.StringPtr("East US"),
								RegionalReplicaCount: to.Int32Ptr(2),
								StorageAccountType:   StorageAccountTypeStandardZRS.ToPtr(),
							}},
					},
				},
				StorageProfile: &GalleryImageVersionStorageProfile{
					Source: &GalleryArtifactVersionSource{
						ID: to.StringPtr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"),
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Update a simple Gallery Image Version without source id.
	poller, err = client.BeginUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		GalleryImageVersionUpdate{
			Properties: &GalleryImageVersionProperties{
				PublishingProfile: &GalleryImageVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: GalleryArtifactPublishingProfileBase{
						TargetRegions: []*TargetRegion{
							&TargetRegion{
								Name:                 to.StringPtr("West US"),
								RegionalReplicaCount: to.Int32Ptr(1),
							},
							&TargetRegion{
								Name:                 to.StringPtr("East US"),
								RegionalReplicaCount: to.Int32Ptr(2),
								StorageAccountType:   StorageAccountTypeStandardZRS.ToPtr(),
							}},
					},
				},
				StorageProfile: &GalleryImageVersionStorageProfile{},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryImageVersions_Get(t *testing.T) {

	// From example Get a gallery image version with replication status.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryImageVersionsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	// From example Get a gallery image version with snapshots as a source.
	_, err = client.Get(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	// From example Get a gallery image version with vhd as a source.
	_, err = client.Get(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	// From example Get a gallery image version.
	_, err = client.Get(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryImageVersions_Delete(t *testing.T) {

	// From example Delete a gallery image version.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryImageVersionsClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		"1.0.0",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryImageVersions_ListByGalleryImage(t *testing.T) {

	// From example List gallery image versions in a gallery image definition.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryImageVersionsClient(con,
		"{subscription-id}")
	client.ListByGalleryImage("myResourceGroup",
		"myGalleryName",
		"myGalleryImageName",
		nil)

}

func TestGalleryApplications_CreateOrUpdate(t *testing.T) {

	// From example Create or update a simple gallery Application.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryApplicationsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryApplicationName",
		GalleryApplication{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &GalleryApplicationProperties{
				Description:         to.StringPtr("This is the gallery application description."),
				Eula:                to.StringPtr("This is the gallery application EULA."),
				PrivacyStatementURI: to.StringPtr("myPrivacyStatementUri}"),
				ReleaseNoteURI:      to.StringPtr("myReleaseNoteUri"),
				SupportedOSType:     OperatingSystemTypesWindows.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryApplications_Update(t *testing.T) {

	// From example Update a simple gallery Application.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryApplicationsClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryApplicationName",
		GalleryApplicationUpdate{
			Properties: &GalleryApplicationProperties{
				Description:         to.StringPtr("This is the gallery application description."),
				Eula:                to.StringPtr("This is the gallery application EULA."),
				PrivacyStatementURI: to.StringPtr("myPrivacyStatementUri}"),
				ReleaseNoteURI:      to.StringPtr("myReleaseNoteUri"),
				SupportedOSType:     OperatingSystemTypesWindows.ToPtr(),
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryApplications_Get(t *testing.T) {

	// From example Get a gallery Application.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryApplicationsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryApplicationName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryApplications_Delete(t *testing.T) {

	// From example Delete a gallery Application.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryApplicationsClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryApplicationName",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryApplications_ListByGallery(t *testing.T) {

	// From example List gallery Applications in a gallery.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryApplicationsClient(con,
		"{subscription-id}")
	client.ListByGallery("myResourceGroup",
		"myGalleryName",
		nil)

}

func TestGalleryApplicationVersions_CreateOrUpdate(t *testing.T) {

	// From example Create or update a simple gallery Application Version.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryApplicationVersionsClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryApplicationName",
		"1.0.0",
		GalleryApplicationVersion{
			Resource: Resource{
				Location: to.StringPtr("West US"),
			},
			Properties: &GalleryApplicationVersionProperties{
				PublishingProfile: &GalleryApplicationVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: GalleryArtifactPublishingProfileBase{
						EndOfLifeDate:      to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339, "2019-07-01T07:00:00Z"); return t }()),
						ReplicaCount:       to.Int32Ptr(1),
						StorageAccountType: StorageAccountTypeStandardLRS.ToPtr(),
						TargetRegions: []*TargetRegion{
							&TargetRegion{
								Name:                 to.StringPtr("West US"),
								RegionalReplicaCount: to.Int32Ptr(1),
								StorageAccountType:   StorageAccountTypeStandardLRS.ToPtr(),
							}},
					},
					ManageActions: &UserArtifactManage{
						Install: to.StringPtr("powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\""),
						Remove:  to.StringPtr("del C:\\package "),
					},
					Source: &UserArtifactSource{
						MediaLink: to.StringPtr("https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"),
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryApplicationVersions_Update(t *testing.T) {

	// From example Update a simple gallery Application Version.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryApplicationVersionsClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryApplicationName",
		"1.0.0",
		GalleryApplicationVersionUpdate{
			Properties: &GalleryApplicationVersionProperties{
				PublishingProfile: &GalleryApplicationVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: GalleryArtifactPublishingProfileBase{
						EndOfLifeDate:      to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339, "2019-07-01T07:00:00Z"); return t }()),
						ReplicaCount:       to.Int32Ptr(1),
						StorageAccountType: StorageAccountTypeStandardLRS.ToPtr(),
						TargetRegions: []*TargetRegion{
							&TargetRegion{
								Name:                 to.StringPtr("West US"),
								RegionalReplicaCount: to.Int32Ptr(1),
								StorageAccountType:   StorageAccountTypeStandardLRS.ToPtr(),
							}},
					},
					ManageActions: &UserArtifactManage{
						Install: to.StringPtr("powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\""),
						Remove:  to.StringPtr("del C:\\package "),
					},
					Source: &UserArtifactSource{
						MediaLink: to.StringPtr("https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"),
					},
				},
			},
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryApplicationVersions_Get(t *testing.T) {

	// From example Get a gallery Application Version with replication status.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryApplicationVersionsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryApplicationName",
		"1.0.0",
		nil)
	if err != nil {
		t.Fatal(err)
	}

	// From example Get a gallery Application Version.
	_, err = client.Get(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryApplicationName",
		"1.0.0",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryApplicationVersions_Delete(t *testing.T) {

	// From example Delete a gallery Application Version.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryApplicationVersionsClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"myResourceGroup",
		"myGalleryName",
		"myGalleryApplicationName",
		"1.0.0",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGalleryApplicationVersions_ListByGalleryApplication(t *testing.T) {

	// From example List gallery Application Versions in a gallery Application Definition.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGalleryApplicationVersionsClient(con,
		"{subscription-id}")
	client.ListByGalleryApplication("myResourceGroup",
		"myGalleryName",
		"myGalleryApplicationName",
		nil)

}

func TestGallerySharingProfile_Update(t *testing.T) {

	// From example Add sharing id to the sharing profile of a gallery.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewGallerySharingProfileClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		SharingUpdate{
			Groups: []*SharingProfileGroup{
				&SharingProfileGroup{
					Type: SharingProfileGroupTypesSubscriptions.ToPtr(),
					IDs: []*string{
						to.StringPtr("34a4ab42-0d72-47d9-bd1a-aed207386dac"),
						to.StringPtr("380fd389-260b-41aa-bad9-0a83108c370b")},
				},
				&SharingProfileGroup{
					Type: SharingProfileGroupTypesAADTenants.ToPtr(),
					IDs: []*string{
						to.StringPtr("c24c76aa-8897-4027-9b03-8f7928b54ff6")},
				}},
			OperationType: SharingUpdateOperationTypesAdd.ToPtr(),
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example reset sharing profile of a gallery.
	poller, err = client.BeginUpdate(ctx,
		"myResourceGroup",
		"myGalleryName",
		SharingUpdate{
			OperationType: SharingUpdateOperationTypesReset.ToPtr(),
		},
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSharedGalleries_List(t *testing.T) {

	// From example Get a gallery.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSharedGalleriesClient(con,
		"{subscription-id}")
	client.List("myLocation",
		nil)

}

func TestSharedGalleries_Get(t *testing.T) {

	// From example Get a gallery.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSharedGalleriesClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myLocation",
		"galleryUniqueName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSharedGalleryImages_List(t *testing.T) {

	// From example Get a gallery.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSharedGalleryImagesClient(con,
		"{subscription-id}")
	client.List("myLocation",
		"galleryUniqueName",
		nil)

}

func TestSharedGalleryImages_Get(t *testing.T) {

	// From example Get a gallery.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSharedGalleryImagesClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myLocation",
		"galleryUniqueName",
		"myGalleryImageName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSharedGalleryImageVersions_List(t *testing.T) {

	// From example Get a gallery.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSharedGalleryImageVersionsClient(con,
		"{subscription-id}")
	client.List("myLocation",
		"galleryUniqueName",
		"myGalleryImageName",
		nil)

}

func TestSharedGalleryImageVersions_Get(t *testing.T) {

	// From example Get a gallery.
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewSharedGalleryImageVersionsClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"myLocation",
		"galleryUniqueName",
		"myGalleryImageName",
		"myGalleryImageVersionName",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServiceRoleInstances_Delete(t *testing.T) {

	// From example Delete Cloud Service Role Instance
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceRoleInstancesClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"{roleInstance-name}",
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServiceRoleInstances_Get(t *testing.T) {

	// From example Get Cloud Service Role Instance
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceRoleInstancesClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"{roleInstance-name}",
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServiceRoleInstances_GetInstanceView(t *testing.T) {

	// From example Get Instance View of Cloud Service Role Instance
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceRoleInstancesClient(con,
		"{subscription-id}")
	_, err := client.GetInstanceView(ctx,
		"{roleInstance-name}",
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServiceRoleInstances_List(t *testing.T) {

	// From example List Role Instances in a Cloud Service
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceRoleInstancesClient(con,
		"{subscription-id}")
	client.List("ConstosoRG",
		"{cs-name}",
		nil)

}

func TestCloudServiceRoleInstances_Restart(t *testing.T) {

	// From example Restart Cloud Service Role Instance
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceRoleInstancesClient(con,
		"{subscription-id}")
	poller, err := client.BeginRestart(ctx,
		"{roleInstance-name}",
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServiceRoleInstances_Reimage(t *testing.T) {

	// From example Reimage Cloud Service Role Instance
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceRoleInstancesClient(con,
		"{subscription-id}")
	poller, err := client.BeginReimage(ctx,
		"{roleInstance-name}",
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServiceRoleInstances_Rebuild(t *testing.T) {

	// From example Rebuild Cloud Service Role Instance
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceRoleInstancesClient(con,
		"{subscription-id}")
	poller, err := client.BeginRebuild(ctx,
		"{roleInstance-name}",
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServiceRoleInstances_GetRemoteDesktopFile(t *testing.T) {

	t.Skip("Warning: No test steps for this operation!")
}

func TestCloudServiceRoles_Get(t *testing.T) {

	// From example Get Cloud Service Role
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceRolesClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"{role-name}",
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServiceRoles_List(t *testing.T) {

	// From example List Roles in a Cloud Service
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceRolesClient(con,
		"{subscription-id}")
	client.List("ConstosoRG",
		"{cs-name}",
		nil)

}

func TestCloudServices_CreateOrUpdate(t *testing.T) {

	// From example Create New Cloud Service with Multiple Roles
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create New Cloud Service with Single Role
	poller, err = client.BeginCreateOrUpdate(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create New Cloud Service with Single Role and Certificate from Key Vault
	poller, err = client.BeginCreateOrUpdate(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// From example Create New Cloud Service with Single Role and RDP Extension
	poller, err = client.BeginCreateOrUpdate(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServices_Update(t *testing.T) {

	// From example Update existing Cloud Service to add tags
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	poller, err := client.BeginUpdate(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServices_Delete(t *testing.T) {

	// From example Delete Cloud Service
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	poller, err := client.BeginDelete(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServices_Get(t *testing.T) {

	// From example Get Cloud Service with Multiple Roles and RDP Extension
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	_, err := client.Get(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServices_GetInstanceView(t *testing.T) {

	// From example Get Cloud Service Instance View with Multiple Roles
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	_, err := client.GetInstanceView(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServices_ListAll(t *testing.T) {

	// From example List Cloud Services in a Subscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	client.ListAll(nil)

}

func TestCloudServices_List(t *testing.T) {

	// From example List Cloud Services in a Resource Group
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	client.List("ConstosoRG",
		nil)

}

func TestCloudServices_Start(t *testing.T) {

	// From example Start Cloud Service
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	poller, err := client.BeginStart(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServices_PowerOff(t *testing.T) {

	// From example Stop or PowerOff Cloud Service
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	poller, err := client.BeginPowerOff(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServices_Restart(t *testing.T) {

	// From example Restart Cloud Service Role Instances
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	poller, err := client.BeginRestart(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServices_Reimage(t *testing.T) {

	// From example Reimage Cloud Service Role Instances
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	poller, err := client.BeginReimage(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServices_Rebuild(t *testing.T) {

	// From example Rebuild Cloud Service Role Instances
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	poller, err := client.BeginRebuild(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServices_DeleteInstances(t *testing.T) {

	// From example Delete Cloud Service Role Instances
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesClient(con,
		"{subscription-id}")
	poller, err := client.BeginDeleteInstances(ctx,
		"ConstosoRG",
		"{cs-name}",
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServicesUpdateDomain_WalkUpdateDomain(t *testing.T) {

	// From example Update Cloud Service to specified Domain
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesUpdateDomainClient(con,
		"{subscription-id}")
	poller, err := client.BeginWalkUpdateDomain(ctx,
		"ConstosoRG",
		"{cs-name}",
		1,
		nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServicesUpdateDomain_GetUpdateDomain(t *testing.T) {

	// From example Get Cloud Service Update Domain
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesUpdateDomainClient(con,
		"{subscription-id}")
	_, err := client.GetUpdateDomain(ctx,
		"ConstosoRG",
		"{cs-name}",
		1,
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServicesUpdateDomain_ListUpdateDomains(t *testing.T) {

	// From example List Update Domains in Cloud Service
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServicesUpdateDomainClient(con,
		"{subscription-id}")
	client.ListUpdateDomains("ConstosoRG",
		"{cs-name}",
		nil)

}

func TestCloudServiceOperatingSystems_GetOSVersion(t *testing.T) {

	// From example Get Cloud Service OS Version
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceOperatingSystemsClient(con,
		"{subscription-id}")
	_, err := client.GetOSVersion(ctx,
		"westus2",
		"WA-GUEST-OS-3.90_202010-02",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServiceOperatingSystems_ListOSVersions(t *testing.T) {

	// From example List Cloud Service OS Versions in a subscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceOperatingSystemsClient(con,
		"{subscription-id}")
	client.ListOSVersions("westus2",
		nil)

}

func TestCloudServiceOperatingSystems_GetOSFamily(t *testing.T) {

	// From example Get Cloud Service OS Family
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceOperatingSystemsClient(con,
		"{subscription-id}")
	_, err := client.GetOSFamily(ctx,
		"westus2",
		"3",
		nil)
	if err != nil {
		t.Fatal(err)
	}

}

func TestCloudServiceOperatingSystems_ListOSFamilies(t *testing.T) {

	// From example List Cloud Service OS Families in a subscription
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	client := NewCloudServiceOperatingSystemsClient(con,
		"{subscription-id}")
	client.ListOSFamilies("westus2",
		nil)

}

// TestMain will exec each test
func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run() // exec test and this returns an exit code to pass to os
	tearDown()
	os.Exit(retCode)
}

func setUp() {
	ctx = context.Background()
	subscriptionId = os.Getenv("SUBSCRIPTION_ID")

	tr := &http.Transport{}
	if err := http2.ConfigureTransport(tr); err != nil {
		fmt.Printf("Failed to configure http2 transport: %v", err)
	}
	tr.TLSClientConfig.InsecureSkipVerify = true
	client := &http.Client{Transport: tr}
	cred, err = azidentity.NewEnvironmentCredential(&azidentity.EnvironmentCredentialOptions{AuthorityHost: mockHost, HTTPClient: client})
	if err != nil {
		panic(err)
	}

	con = armcore.NewConnection(mockHost, cred, &armcore.ConnectionOptions{
		Logging: azcore.LogOptions{
			IncludeBody: true,
		},
		HTTPClient: client,
	})
}

func tearDown() {

}
