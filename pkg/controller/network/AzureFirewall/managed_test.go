/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package AzureFirewall

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
	"github.com/Azure/go-autorest/autorest"
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/crossplane/provider-azure/apis/network/v1alpha3"
	azure "github.com/crossplane/provider-azure/pkg/clients"
	"github.com/crossplane/provider-azure/pkg/clients/network/fake"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"net/http"
	"testing"
)

const (
	name              = "coolAzureFirewall"
	uid               = types.UID("definitely-a-uuid")
	resourceGroupName = "coolRG"
	location          = "coolLocation"
	firewallType      = "coolType"
	etag              = "coolEtag"
	id                = "coolID"
	privateIpAddress  = "coolAddress"
	ipAddress         = "coolPublicAddress"
)

var (
	ctx       = context.Background()
	errorBoom = errors.New("boom")
	tags      = map[string]string{"one": "test", "two": "test"}
	zones     = []string{"one", "two"}
)

type testCase struct {
	name    string
	e       managed.ExternalClient
	r       resource.Managed
	want    resource.Managed
	wantErr error
}

type azureFirewallModifier func(firewall *v1alpha3.AzureFirewall)

func withConditions(c ...runtimev1alpha1.Condition) azureFirewallModifier {
	return func(r *v1alpha3.AzureFirewall) { r.Status.ConditionedStatus.Conditions = c }
}

func withState(s string) azureFirewallModifier {
	return func(r *v1alpha3.AzureFirewall) { r.Status.State = s }
}

func azureFirewall(sm ...azureFirewallModifier) *v1alpha3.AzureFirewall {
	r := &v1alpha3.AzureFirewall{
		ObjectMeta: metav1.ObjectMeta{
			Name:       name,
			UID:        uid,
			Finalizers: []string{},
		},
		Spec: v1alpha3.AzureFirewallSpec{
			ResourceSpec:      runtimev1alpha1.ResourceSpec{},
			ResourceGroupName: resourceGroupName,
			Location:          location,
			AzureFirewallPropertiesFormat: v1alpha3.AzureFirewallPropertiesFormat{
				ApplicationRuleCollections: setApplicationRuleCollection(),
				NatRuleCollections:         setNatRuleCollection(),
				NetworkRuleCollections:     setNetworkRuleCollection(),
				IPConfigurations:           setIpConfigurations(),
				ProvisioningState:          "",
				ThreatIntelMode:            "Alert",
				VirtualHub:                 setSubResource(),
				FirewallPolicy:             setSubResource(),
				HubIPAddresses:             setHubIPAddresses(),
			},
			Tags:  tags,
			Zones: zones,
			Etag:  etag,
			ID:    string(uid),
			Name:  name,
			Type:  firewallType,
		},
		Status: v1alpha3.AzureFirewallStatus{},
	}

	meta.SetExternalName(r, name)

	for _, m := range sm {
		m(r)
	}

	return r
}

func setNetworkRuleCollection() *[]v1alpha3.AzureFirewallNetworkRuleCollection {
	var afnrcs = new([]v1alpha3.AzureFirewallNetworkRuleCollection)
	var afnrc = new(v1alpha3.AzureFirewallNetworkRuleCollection)
	afnrc.Name = "NetworkRule Name"
	afnrc.Etag = etag
	afnrc.ID = "ID"
	afnrc.Properties = v1alpha3.AzureFirewallNetworkRuleCollectionPropertiesFormat{
		Priority:          0,
		Action:            "Allow",
		Rules:             setAzureFirewallNetworkRule(),
		ProvisioningState: "",
	}
	*afnrcs = append(*afnrcs, *afnrc)
	return afnrcs
}

func setAzureFirewallNetworkRule() []v1alpha3.AzureFirewallNetworkRule {
	var afrns = new([]v1alpha3.AzureFirewallNetworkRule)
	var afrn = new(v1alpha3.AzureFirewallNetworkRule)
	afrn.Name = "Network Rule Name"
	afrn.Description = "Network rule description"
	afrn.DestinationPorts = []string{"80", "8081"}
	afrn.DestinationAddresses = []string{"address1", "address2"}
	afrn.SourceAddresses = []string{"source1", "source2"}
	afrn.Protocols = []string{"TCP", "UDP"}
	*afrns = append(*afrns, *afrn)
	return *afrns
}

func setNatRuleCollection() *[]v1alpha3.AzureFirewallNatRuleCollection {
	var afnrcs = new([]v1alpha3.AzureFirewallNatRuleCollection)
	var afnrc = new(v1alpha3.AzureFirewallNatRuleCollection)
	afnrc.Name = "Nat Rule Name"
	afnrc.ID = "ID"
	afnrc.Etag = etag
	afnrc.Properties = v1alpha3.AzureFirewallNatRuleCollectionProperties{
		Priority:          0,
		Action:            "Allow",
		Rules:             setAzureFirewallNATRule(),
		ProvisioningState: "",
	}
	*afnrcs = append(*afnrcs, *afnrc)
	return afnrcs
}

func setAzureFirewallNATRule() []v1alpha3.AzureFirewallNatRule {
	var afnrs = new([]v1alpha3.AzureFirewallNatRule)
	var afnr = new(v1alpha3.AzureFirewallNatRule)
	afnr.Name = "Nat Rule Name"
	afnr.Description = "Nat Rule Description"
	afnr.TranslatedAddress = "Translated Address"
	afnr.TranslatedPort = "Translated Port"
	afnr.Protocols = []string{"TCP", "UDP"}
	afnr.SourceAddresses = []string{"source1", "source2"}
	afnr.DestinationAddresses = []string{"destinationAddress1", "destinationAddress2"}
	afnr.DestinationPorts = []string{"80", "8080", "8443"}
	*afnrs = append(*afnrs, *afnr)
	return *afnrs
}

func setApplicationRuleCollection() *[]v1alpha3.AzureFirewallApplicationRuleCollection {
	var afarc = new([]v1alpha3.AzureFirewallApplicationRuleCollection)
	var afar = new(v1alpha3.AzureFirewallApplicationRuleCollection)
	afar.Name = "name"
	afar.Etag = etag
	afar.ID = "ID"
	afar.Properties = v1alpha3.AzureFirewallApplicationRuleCollectionPropertiesFormat{
		Priority:          0,
		Action:            "Allow",
		Rules:             setAzureFirewallApplicationRule(),
		ProvisioningState: "succeeded",
	}
	*afarc = append(*afarc, *afar)
	return afarc
}

func setAzureFirewallApplicationRule() []v1alpha3.AzureFirewallApplicationRule {
	var afars = new([]v1alpha3.AzureFirewallApplicationRule)
	var afar = new(v1alpha3.AzureFirewallApplicationRule)
	afar.Name = "Rule Name"
	afar.Description = "Rule Description"
	afar.FqdnTags = []string{"one", "two", "three"}
	afar.SourceAddresses = []string{"source1", "source2", "source3"}
	afar.TargetFqdns = []string{"target1", "target2"}
	afar.Protocols = setApplicationRuleProtocol()
	*afars = append(*afars, *afar)
	return *afars
}

func setApplicationRuleProtocol() []v1alpha3.AzureFirewallApplicationRuleProtocol {
	var afarpcs = new([]v1alpha3.AzureFirewallApplicationRuleProtocol)
	var afarpc = new(v1alpha3.AzureFirewallApplicationRuleProtocol)
	afarpc.Port = 0
	afarpc.ProtocolType = "Protocol Type"
	*afarpcs = append(*afarpcs, *afarpc)
	return *afarpcs
}

func setIpConfigurations() *[]v1alpha3.AzureFirewallIPConfiguration {
	var afipc = new([]v1alpha3.AzureFirewallIPConfiguration)
	var afip = new(v1alpha3.AzureFirewallIPConfiguration)
	afip.ID = azure.ToStringPtr("ID")
	afip.Etag = azure.ToStringPtr(etag)
	afip.Name = azure.ToStringPtr("name")
	afip.AzureFirewallIPConfigurationPropertiesFormat = v1alpha3.AzureFirewallIPConfigurationPropertiesFormat{
		PrivateIPAddress:  azure.ToStringPtr(privateIpAddress),
		Subnet:            setSubResource(),
		PublicIPAddress:   setSubResource(),
		ProvisioningState: nil,
	}
	*afipc = append(*afipc, *afip)
	return afipc
}

func setHubIPAddresses() *v1alpha3.HubIPAddresses {
	var res = new(v1alpha3.HubIPAddresses)
	res.PrivateIPAddress = privateIpAddress
	var publicIpAddress = new([]v1alpha3.AzureFirewallPublicIPAddress)
	var pIAdd = new(v1alpha3.AzureFirewallPublicIPAddress)
	pIAdd.Address = azure.ToStringPtr(ipAddress)
	*publicIpAddress = append(*publicIpAddress, *pIAdd)
	res.PublicIPAddresses = *publicIpAddress
	return res
}

func setSubResource() *v1alpha3.SubResource {
	var res = new(v1alpha3.SubResource)
	res.ID = id
	return res
}

func TestCreate(t *testing.T) {
	cases := []testCase{
		{
			name:    "NotFirewall",
			e:       &external{client: &fake.MockAzureFirewallClient{}},
			r:       &v1alpha3.VirtualNetwork{},
			want:    &v1alpha3.VirtualNetwork{},
			wantErr: errors.New(errNotAzureFirewall),
		},
		{
			name: "SuccessfulCreate",
			e: &external{client: &fake.MockAzureFirewallClient{
				MockCreateOrUpdate: func(_ context.Context, _ string, _ string, _ network.AzureFirewall) (network.AzureFirewallsCreateOrUpdateFuture, error) {
					return network.AzureFirewallsCreateOrUpdateFuture{}, nil
				},
			}},
			r: azureFirewall(),
			want: azureFirewall(
				withConditions(runtimev1alpha1.Creating()),
			),
		},
		{
			name: "FailedCreate",
			e: &external{client: &fake.MockAzureFirewallClient{
				MockCreateOrUpdate: func(_ context.Context, _ string, _ string, _ network.AzureFirewall) (network.AzureFirewallsCreateOrUpdateFuture, error) {
					return network.AzureFirewallsCreateOrUpdateFuture{}, errorBoom
				},
			}},
			r: azureFirewall(),
			want: azureFirewall(
				withConditions(runtimev1alpha1.Creating()),
			),
			wantErr: errors.Wrap(errorBoom, errCreateAzureFirewall),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.e.Create(ctx, tc.r)

			if diff := cmp.Diff(tc.wantErr, err, test.EquateErrors()); diff != "" {
				t.Errorf("tc.e.Create(...): want error != got error:\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, tc.r, test.EquateConditions()); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestObserve(t *testing.T) {
	cases := []testCase{
		{
			name:    "NotAzureFirewall",
			e:       &external{client: &fake.MockAzureFirewallClient{}},
			r:       &v1alpha3.VirtualNetwork{},
			want:    &v1alpha3.VirtualNetwork{},
			wantErr: errors.New(errNotAzureFirewall),
		},
		{
			name: "SuccessfulObserveNotExist",
			e: &external{client: &fake.MockAzureFirewallClient{
				MockGet: func(_ context.Context, _ string, _ string) (result network.AzureFirewall, err error) {
					return network.AzureFirewall{
							AzureFirewallPropertiesFormat: &network.AzureFirewallPropertiesFormat{
								ThreatIntelMode: "Alert",
							},
							Zones:    azure.ToStringArrayPtr(zones),
							Etag:     azure.ToStringPtr(etag),
							Type:     azure.ToStringPtr(firewallType),
							Location: azure.ToStringPtr(location),
							Tags:     azure.ToStringPtrMap(tags),
						}, autorest.DetailedError{
							StatusCode: http.StatusNotFound,
						}
				},
			}},
			r:    azureFirewall(),
			want: azureFirewall(),
		},
		{
			name: "SuccessfulObserveExists",
			e: &external{client: &fake.MockAzureFirewallClient{
				MockGet: func(_ context.Context, _ string, _ string) (result network.AzureFirewall, err error) {
					return network.AzureFirewall{
						Response: autorest.Response{},
						AzureFirewallPropertiesFormat: &network.AzureFirewallPropertiesFormat{
							ApplicationRuleCollections: nil,
							NatRuleCollections:         nil,
							NetworkRuleCollections:     nil,
							IPConfigurations:           nil,
							ProvisioningState:          network.ProvisioningState(string(network.Available)),
							ThreatIntelMode:            "",
							VirtualHub:                 nil,
							FirewallPolicy:             nil,
							HubIPAddresses:             nil,
						},
						Zones:    azure.ToStringArrayPtr(zones),
						Name:     azure.ToStringPtr(name),
						Location: azure.ToStringPtr(location),
						Tags:     azure.ToStringPtrMap(tags),
					}, nil
				},
			}},
			r: azureFirewall(),
			want: azureFirewall(
				withConditions(runtimev1alpha1.Available()),
				withState(string(network.Available)),
			),
		},

		{
			name: "FailedObserve",
			e: &external{client: &fake.MockAzureFirewallClient{
				MockGet: func(_ context.Context, _ string, _ string) (result network.AzureFirewall, err error) {
					return network.AzureFirewall{}, errorBoom
				},
			}},
			r:       azureFirewall(),
			want:    azureFirewall(),
			wantErr: errors.Wrap(errorBoom, errGetAzureFirewall),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.e.Observe(ctx, tc.r)

			if diff := cmp.Diff(tc.wantErr, err, test.EquateErrors()); diff != "" {
				t.Errorf("tc.e.Observe(...): want error != got error:\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, tc.r, test.EquateConditions()); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}
