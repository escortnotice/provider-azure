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

package RouteTable

import (
	"context"
	azurenetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network/networkapi"
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/provider-azure/apis/network/v1alpha3"
	azureclients "github.com/crossplane/provider-azure/pkg/clients"
	"github.com/crossplane/provider-azure/pkg/clients/network"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Error strings.
const (
	errNotRouteTable    = "managed resource is not an RouteTable"
	errCreateRouteTable = "cannot create RouteTable"
	errUpdateRouteTable = "cannot update RouteTable"
	errGetRouteTable    = "cannot get RouteTable"
	errDeleteRouteTable = "cannot delete RouteTable"
)

// Setup adds a controller that reconciles Subnets.
func Setup(mgr ctrl.Manager, l logging.Logger) error {
	name := managed.ControllerName(v1alpha3.RouteTableGroupKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&v1alpha3.RouteTable{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1alpha3.RouteTableGroupVersionKind),
			managed.WithConnectionPublishers(),
			managed.WithExternalConnecter(&connecter{client: mgr.GetClient()}),
			managed.WithReferenceResolver(managed.NewAPISimpleReferenceResolver(mgr.GetClient())),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

type connecter struct {
	client client.Client
}

func (c connecter) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	creds, auth, err := azureclients.GetAuthInfo(ctx, c.client, mg)
	if err != nil {
		return nil, err
	}
	cl := azurenetwork.NewRouteTablesClient(creds[azureclients.CredentialsKeySubscriptionID])
	cl.Authorizer = auth
	return &external{client: cl}, nil
}

type external struct {
	client networkapi.RouteTablesClientAPI
}

func (e external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	rt, ok := mg.(*v1alpha3.RouteTable)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotRouteTable)
	}

	az, err := e.client.Get(ctx, rt.Spec.ResourceGroupName, meta.GetExternalName(rt), "")
	if azureclients.IsNotFound(err) {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, errGetRouteTable)
	}

	network.UpdateRouteTableStatusFromAzure(rt, az)
	rt.SetConditions(runtimev1alpha1.Available())

	o := managed.ExternalObservation{
		ResourceExists:    true,
		ConnectionDetails: managed.ConnectionDetails{},
	}

	return o, nil
}

func (e external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	rt, ok := mg.(*v1alpha3.RouteTable)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotRouteTable)
	}

	rt.Status.SetConditions(runtimev1alpha1.Creating())

	rtable := network.NewRouteTableParameters(rt)
	if _, err := e.client.CreateOrUpdate(ctx, rt.Spec.ResourceGroupName, meta.GetExternalName(rt), rtable); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errCreateRouteTable)
	}

	return managed.ExternalCreation{}, nil
}

func (e external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	panic("implement me")
}

func (e external) Delete(ctx context.Context, mg resource.Managed) error {
	panic("implement me")
}
