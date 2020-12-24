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
package virtualmachine

//import (
//	"context"
//	azurenetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
//
//	//azurenetwork "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
//	azurecompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
//	//securitygroup "github.com/crossplane/provider-azure/pkg/clients/network"
//	virtualmachine "github.com/crossplane/provider-azure/pkg/clients/compute"
//
//	//"github.com/crossplane/provider-azure/pkg/clients/network"
//
//	//"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-12-01/network/networkapi"
//	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
//	"github.com/crossplane/crossplane-runtime/pkg/event"
//	"github.com/crossplane/crossplane-runtime/pkg/logging"
//	"github.com/crossplane/crossplane-runtime/pkg/meta"
//	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
//	"github.com/crossplane/crossplane-runtime/pkg/resource"
//	"github.com/crossplane/provider-azure/apis/compute/v1alpha3"
//	azureclients "github.com/crossplane/provider-azure/pkg/clients"
//	"github.com/pkg/errors"
//	ctrl "sigs.k8s.io/controller-runtime"
//	"sigs.k8s.io/controller-runtime/pkg/client"
//)
//
//// Error strings.
//const (
//	errNotSecurityGroup    = "managed resource is not an SecurityGroup"
//	errCreateSecurityGroup = "cannot create SecurityGroup"
//	errUpdateSecurityGroup = "cannot update SecurityGroup"
//	errGetSecurityGroup    = "cannot get SecurityGroup"
//	errDeleteSecurityGroup = "cannot delete SecurityGroup"
//)
//
//// Setup adds a controller that reconciles Security Group.
//func Setup(mgr ctrl.Manager, l logging.Logger) error {
//	name := managed.ControllerName(v1alpha3.VirtualMachineKind)
//
//	return ctrl.NewControllerManagedBy(mgr).
//		Named(name).
//		For(&v1alpha3.VirtualMachine{}).
//		Complete(managed.NewReconciler(mgr,
//			resource.ManagedKind(v1alpha3.VirtualMachineGroupVersionKind),
//			managed.WithConnectionPublishers(),
//			managed.WithExternalConnecter(&connecter{client: mgr.GetClient()}),
//			managed.WithReferenceResolver(managed.NewAPISimpleReferenceResolver(mgr.GetClient())),
//			managed.WithLogger(l.WithValues("controller", name)),
//			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
//}
//
//type connecter struct {
//	client client.Client
//}
//
//func (c *connecter) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
//	creds, auth, err := azureclients.GetAuthInfo(ctx, c.client, mg)
//	if err != nil {
//		return nil, err
//	}
//	cl := azurecompute.NewVirtualMachinesClient(creds[azureclients.CredentialsKeySubscriptionID])
//	cl.Authorizer = auth
//	return &external{client: cl}, nil
//}
//
//type external struct {
//	//client computeapi.VirtualMachinesClientAPI
//	client azurecompute.VirtualMachinesClient
//}
//
//func (e *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
//	v, ok := mg.(*v1alpha3.VirtualMachine)
//	if !ok {
//		return managed.ExternalObservation{}, errors.New(errNotSecurityGroup)
//	}
//	az, err := e.client.Get(ctx, v.Spec.ResourceGroupName, meta.GetExternalName(v), "")
//	if azureclients.IsNotFound(err) {
//		return managed.ExternalObservation{ResourceExists: false}, nil
//	}
//	if err != nil {
//		return managed.ExternalObservation{}, errors.Wrap(err, errGetSecurityGroup)
//	}
//
//	virtualmachine.UpdateSecurityGroupStatusFromAzure(v, az)
//
//	v.SetConditions(runtimev1alpha1.Available())
//
//	o := managed.ExternalObservation{
//		ResourceExists:    true,
//		ConnectionDetails: managed.ConnectionDetails{},
//	}
//	return o, nil
//}
