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

package PublicIPAddress

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/test"

	"github.com/crossplane/provider-azure/apis/network/v1alpha3"
	azure "github.com/crossplane/provider-azure/pkg/clients"
	"github.com/crossplane/provider-azure/pkg/clients/network/fake"
)

const (
	name              = "cooPubIP"
	uid               = types.UID("definitely-a-uuid")
	resourceGroupName = "coolRG"
	location          = "coolplace"
	domainNameLabel = "cooldomain"
	fqdn =  "cooldomain.fqdn.net"
)

var (
	ctx       = context.Background()
	errorBoom = errors.New("boom")
	tags      = map[string]string{"one": "test", "two": "test"}
)

type testCase struct {
	name    string
	e       managed.ExternalClient
	r       resource.Managed
	want    resource.Managed
	wantErr error
}

type publicIPAddressModifier func(address *v1alpha3.PublicIPAddress)

func withConditions(c ...runtimev1alpha1.Condition) publicIPAddressModifier {
	return func(r *v1alpha3.PublicIPAddress) { r.Status.ConditionedStatus.Conditions = c }
}

func withState(s string) publicIPAddressModifier {
	return func(r *v1alpha3.PublicIPAddress) { r.Status.State = s }
}

func publicIPAddress(vm ...publicIPAddressModifier) *v1alpha3.PublicIPAddress{
	r := &v1alpha3.PublicIPAddress{
		ObjectMeta: metav1.ObjectMeta{
			Name:       name,
			UID:        uid,
			Finalizers: []string{},
		},
		Spec: v1alpha3.PublicIPAddressSpec{
			ResourceGroupName: resourceGroupName,
			Properties: v1alpha3.PublicIPAddressPropertiesFormat{
				IPAddressSkuName : v1alpha3.PublicIPAddressSkuNameBasic,
				PublicIPAddressVersion: v1alpha3.IPv4,
				DNSSettings: &v1alpha3.PublicIPAddressDNSSettings{
					DomainNameLabel:azure.ToStringPtr(domainNameLabel),
					Fqdn: azure.ToStringPtr(fqdn),
				},
				IdleTimeoutInMinutes  : 4,
				PublicIPAllocationMethod   : v1alpha3.Static,
			},
			Location: location,
			Tags:     tags,
		},
		Status: v1alpha3.PublicIPAddressStatus{},
	}
	meta.SetExternalName(r, name)

	for _, m := range vm {
		m(r)
	}

	return r
}

// Test that our Reconciler implementation satisfies the Reconciler interface.
var _ managed.ExternalClient = &external{}
var _ managed.ExternalConnecter = &connecter{}

func TestCreate(t *testing.T) {
	cases := []testCase{
		{
			name:    "NotPublicIPAddress",
			e:       &external{client: &fake.MockPublicIPAddressClient{}},
			r:       &v1alpha3.Subnet{},
			want:    &v1alpha3.Subnet{},
			wantErr: errors.New(errNotAPublicIPAddress),
		},
		{
			name: "SuccessfulCreate",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockCreateOrUpdate: func(_ context.Context, _ string, _ string, _ network.PublicIPAddress) (result network.PublicIPAddressesCreateOrUpdateFuture, err error) {
					return network.PublicIPAddressesCreateOrUpdateFuture{}, nil
				},
			}},
			r: publicIPAddress(),
			want: publicIPAddress(
				withConditions(runtimev1alpha1.Creating()),
			),
		},
		{
			name: "FailedCreate",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockCreateOrUpdate: func(_ context.Context, _ string, _ string, _ network.PublicIPAddress) (result network.PublicIPAddressesCreateOrUpdateFuture, err error) {
					return network.PublicIPAddressesCreateOrUpdateFuture{}, errorBoom
				},
			}},
			r: publicIPAddress(),
			want: publicIPAddress(
				withConditions(runtimev1alpha1.Creating()),
			),
			wantErr: errors.Wrap(errorBoom, errCreatePublicIPAddress),
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
			name:    "NoPublicIPAddress",
			e:       &external{client: &fake.MockPublicIPAddressClient{}},
			r:       nil,
			want:    nil,
			wantErr: errors.New(errNotAPublicIPAddress),
		},
		{
			name: "SuccessfulObserveNotExist",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockGet: func(_ context.Context, _ string, _ string, _ string) (result network.PublicIPAddress, err error) {
					return network.PublicIPAddress{
							Tags: azure.ToStringPtrMap(tags),
							PublicIPAddressPropertiesFormat : &network.PublicIPAddressPropertiesFormat{
								PublicIPAddressVersion: network.IPv4,
								PublicIPAllocationMethod: network.Static,
							},
						}, autorest.DetailedError{
							StatusCode: http.StatusNotFound,
						}
				},
			}},
			r:    publicIPAddress(),
			want: publicIPAddress(),
		},
		{
			name: "SuccessfulObserveExists",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockGet: func(_ context.Context, _ string, _ string, _ string) (result network.PublicIPAddress, err error) {
					return network.PublicIPAddress{
						Location: azure.ToStringPtr(location),
						Name: azure.ToStringPtr(name),
						Tags: azure.ToStringPtrMap(tags),
						PublicIPAddressPropertiesFormat : &network.PublicIPAddressPropertiesFormat{
							PublicIPAddressVersion: network.IPv4,
							PublicIPAllocationMethod: network.Static,
							ProvisioningState:    azure.ToStringPtr(string(network.Available)),
						},
					}, nil
				},
			}},
			r: publicIPAddress(),
			want: publicIPAddress(
				withConditions(runtimev1alpha1.Available()),
				withState(string(network.Available)),
			),
		},
		{
			name: "FailedObserve",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockGet: func(_ context.Context, _ string, _ string, _ string) (result network.PublicIPAddress, err error) {
					return network.PublicIPAddress{}, errorBoom
				},
			}},
			r:       publicIPAddress(),
			want:    publicIPAddress(),
			wantErr: errors.Wrap(errorBoom, errGetPublicIPAddress),
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

func TestUpdate(t *testing.T) {
	cases := []testCase{
		{
			name: "SuccessfulDoesNotNeedUpdate",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockGet: func(_ context.Context, _ string, _ string, _ string) (result network.PublicIPAddress, err error) {
					return network.PublicIPAddress{
						Location: azure.ToStringPtr(location),
						Name: azure.ToStringPtr(name),
						Tags: azure.ToStringPtrMap(tags),
						PublicIPAddressPropertiesFormat : &network.PublicIPAddressPropertiesFormat{
							PublicIPAddressVersion: network.IPv4,
							PublicIPAllocationMethod: network.Static,
							IdleTimeoutInMinutes: azure.ToInt32Ptr(4),
							DNSSettings: &network.PublicIPAddressDNSSettings{
								DomainNameLabel: azure.ToStringPtr(domainNameLabel),
								Fqdn: azure.ToStringPtr(fqdn),
							},
							},
					}, nil
				},
			}},
			r:    publicIPAddress(),
			want:  publicIPAddress(),
		},
		{
			name:    "NotPublicIPAddress",
			e:       &external{client: &fake.MockPublicIPAddressClient{}},
			r:       &v1alpha3.VirtualNetwork{},
			want:    &v1alpha3.VirtualNetwork{},
			wantErr: errors.New(errNotAPublicIPAddress),
		},

		{
			name: "SuccessfulNeedsUpdate",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockGet: func(_ context.Context, _ string, _ string, _ string) (result network.PublicIPAddress, err error) {
					return network.PublicIPAddress{
						Tags: azure.ToStringPtrMap(tags),
						PublicIPAddressPropertiesFormat : &network.PublicIPAddressPropertiesFormat{
							PublicIPAddressVersion: network.IPv4,
							PublicIPAllocationMethod: network.Static,
						},
					}, nil
				},
				MockCreateOrUpdate: func(_ context.Context, _ string, _ string, _ network.PublicIPAddress) (result network.PublicIPAddressesCreateOrUpdateFuture, err error) {
					return network.PublicIPAddressesCreateOrUpdateFuture{}, nil
				},
			}},
			r:    publicIPAddress(),
			want:  publicIPAddress(),
		},
		{
			name: "UnsuccessfulGet",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockGet: func(_ context.Context, _ string, _ string, _ string) (result network.PublicIPAddress, err error) {
					return network.PublicIPAddress{
						Tags: azure.ToStringPtrMap(tags),
						PublicIPAddressPropertiesFormat : &network.PublicIPAddressPropertiesFormat{
							PublicIPAddressVersion: network.IPv4,
							PublicIPAllocationMethod: network.Static,
						},
					}, errorBoom
				},
			}},
			r:    publicIPAddress(),
			want:  publicIPAddress(),
			wantErr: errors.Wrap(errorBoom, errGetPublicIPAddress),
		},
		{
			name: "UnsuccessfulUpdate",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockGet: func(_ context.Context, _ string, _ string, _ string) (result network.PublicIPAddress, err error) {
					return network.PublicIPAddress{
						Tags: azure.ToStringPtrMap(tags),
						PublicIPAddressPropertiesFormat : &network.PublicIPAddressPropertiesFormat{
							PublicIPAddressVersion: network.IPv4,
							PublicIPAllocationMethod: network.Static,
						},
					}, nil
				},
				MockCreateOrUpdate: func(_ context.Context, _ string, _ string, _ network.PublicIPAddress) (result network.PublicIPAddressesCreateOrUpdateFuture, err error) {
					return network.PublicIPAddressesCreateOrUpdateFuture{}, errorBoom
				},
			}},
			r:    publicIPAddress(),
			want:  publicIPAddress(),
			wantErr: errors.Wrap(errorBoom, errUpdatePublicIPAddress),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.e.Update(ctx, tc.r)

			if diff := cmp.Diff(tc.wantErr, err, test.EquateErrors()); diff != "" {
				t.Errorf("tc.e.Update(...): want error != got error:\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, tc.r, test.EquateConditions()); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	cases := []testCase{
		{
			name:    "NotPublicIPAddress",
			e:       &external{client: &fake.MockPublicIPAddressClient{}},
			r:       &v1alpha3.Subnet{},
			want:    &v1alpha3.Subnet{},
			wantErr: errors.New(errNotAPublicIPAddress),
		},
		{
			name: "Successful",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockDelete: func(_ context.Context, _ string, _ string) (result network.PublicIPAddressesDeleteFuture, err error) {
					return network.PublicIPAddressesDeleteFuture{}, nil
				},
			}},
			r:    publicIPAddress(),
			want:  publicIPAddress(
				withConditions(runtimev1alpha1.Deleting()),
			),
		},
		{
			name: "SuccessfulNotFound",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockDelete: func(_ context.Context, _ string, _ string) (result network.PublicIPAddressesDeleteFuture, err error) {
					return network.PublicIPAddressesDeleteFuture{}, autorest.DetailedError{
						StatusCode: http.StatusNotFound,
					}
				},
			}},
			r:    publicIPAddress(),
			want:  publicIPAddress(
				withConditions(runtimev1alpha1.Deleting()),
			),
		},
		{
			name: "Failed",
			e: &external{client: &fake.MockPublicIPAddressClient{
				MockDelete: func(_ context.Context, _ string, _ string) (result network.PublicIPAddressesDeleteFuture, err error) {
					return network.PublicIPAddressesDeleteFuture{}, errorBoom
				},
			}},
			r:    publicIPAddress(),
			want:  publicIPAddress(
				withConditions(runtimev1alpha1.Deleting()),
			),
			wantErr: errors.Wrap(errorBoom, errDeletePublicIPAddress),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.e.Delete(ctx, tc.r)

			if diff := cmp.Diff(tc.wantErr, err, test.EquateErrors()); diff != "" {
				t.Errorf("tc.e.Delete(...): want error != got error:\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, tc.r, test.EquateConditions()); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}
