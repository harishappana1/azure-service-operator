// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import (
	"fmt"
	alpha20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20201201storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Deprecated version of RedisLinkedServer. Use v1beta20201201.RedisLinkedServer instead
type RedisLinkedServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Redis_LinkedServer_Spec                `json:"spec,omitempty"`
	Status            RedisLinkedServerWithProperties_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisLinkedServer{}

// GetConditions returns the conditions of the resource
func (server *RedisLinkedServer) GetConditions() conditions.Conditions {
	return server.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (server *RedisLinkedServer) SetConditions(conditions conditions.Conditions) {
	server.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisLinkedServer{}

// ConvertFrom populates our RedisLinkedServer from the provided hub RedisLinkedServer
func (server *RedisLinkedServer) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20201201s.RedisLinkedServer

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = server.AssignProperties_From_RedisLinkedServer(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to server")
	}

	return nil
}

// ConvertTo populates the provided hub RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20201201s.RedisLinkedServer
	err := server.AssignProperties_To_RedisLinkedServer(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from server")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1alpha1api20201201-redislinkedserver,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redislinkedservers,verbs=create;update,versions=v1alpha1api20201201,name=default.v1alpha1api20201201.redislinkedservers.cache.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &RedisLinkedServer{}

// Default applies defaults to the RedisLinkedServer resource
func (server *RedisLinkedServer) Default() {
	server.defaultImpl()
	var temp interface{} = server
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (server *RedisLinkedServer) defaultAzureName() {
	if server.Spec.AzureName == "" {
		server.Spec.AzureName = server.Name
	}
}

// defaultImpl applies the code generated defaults to the RedisLinkedServer resource
func (server *RedisLinkedServer) defaultImpl() { server.defaultAzureName() }

var _ genruntime.KubernetesResource = &RedisLinkedServer{}

// AzureName returns the Azure name of the resource
func (server *RedisLinkedServer) AzureName() string {
	return server.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (server RedisLinkedServer) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (server *RedisLinkedServer) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (server *RedisLinkedServer) GetSpec() genruntime.ConvertibleSpec {
	return &server.Spec
}

// GetStatus returns the status of this resource
func (server *RedisLinkedServer) GetStatus() genruntime.ConvertibleStatus {
	return &server.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (server *RedisLinkedServer) GetType() string {
	return "Microsoft.Cache/redis/linkedServers"
}

// NewEmptyStatus returns a new empty (blank) status
func (server *RedisLinkedServer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisLinkedServerWithProperties_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (server *RedisLinkedServer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(server.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  server.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (server *RedisLinkedServer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisLinkedServerWithProperties_STATUS); ok {
		server.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisLinkedServerWithProperties_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	server.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1alpha1api20201201-redislinkedserver,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redislinkedservers,verbs=create;update,versions=v1alpha1api20201201,name=validate.v1alpha1api20201201.redislinkedservers.cache.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &RedisLinkedServer{}

// ValidateCreate validates the creation of the resource
func (server *RedisLinkedServer) ValidateCreate() error {
	validations := server.createValidations()
	var temp interface{} = server
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (server *RedisLinkedServer) ValidateDelete() error {
	validations := server.deleteValidations()
	var temp interface{} = server
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (server *RedisLinkedServer) ValidateUpdate(old runtime.Object) error {
	validations := server.updateValidations()
	var temp interface{} = server
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (server *RedisLinkedServer) createValidations() []func() error {
	return []func() error{server.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (server *RedisLinkedServer) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (server *RedisLinkedServer) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return server.validateResourceReferences()
		},
		server.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (server *RedisLinkedServer) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&server.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (server *RedisLinkedServer) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*RedisLinkedServer)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, server)
}

// AssignProperties_From_RedisLinkedServer populates our RedisLinkedServer from the provided source RedisLinkedServer
func (server *RedisLinkedServer) AssignProperties_From_RedisLinkedServer(source *alpha20201201s.RedisLinkedServer) error {

	// ObjectMeta
	server.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Redis_LinkedServer_Spec
	err := spec.AssignProperties_From_Redis_LinkedServer_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Redis_LinkedServer_Spec() to populate field Spec")
	}
	server.Spec = spec

	// Status
	var status RedisLinkedServerWithProperties_STATUS
	err = status.AssignProperties_From_RedisLinkedServerWithProperties_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RedisLinkedServerWithProperties_STATUS() to populate field Status")
	}
	server.Status = status

	// No error
	return nil
}

// AssignProperties_To_RedisLinkedServer populates the provided destination RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) AssignProperties_To_RedisLinkedServer(destination *alpha20201201s.RedisLinkedServer) error {

	// ObjectMeta
	destination.ObjectMeta = *server.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20201201s.Redis_LinkedServer_Spec
	err := server.Spec.AssignProperties_To_Redis_LinkedServer_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Redis_LinkedServer_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20201201s.RedisLinkedServerWithProperties_STATUS
	err = server.Status.AssignProperties_To_RedisLinkedServerWithProperties_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RedisLinkedServerWithProperties_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (server *RedisLinkedServer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: server.Spec.OriginalVersion(),
		Kind:    "RedisLinkedServer",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of RedisLinkedServer. Use v1beta20201201.RedisLinkedServer instead
type RedisLinkedServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisLinkedServer `json:"items"`
}

type Redis_LinkedServer_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// +kubebuilder:validation:Required
	LinkedRedisCacheReference *genruntime.ResourceReference `armReference:"LinkedRedisCacheId" json:"linkedRedisCacheReference,omitempty"`
	Location                  *string                       `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/Redis resource
	Owner *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`

	// +kubebuilder:validation:Required
	ServerRole *RedisLinkedServerCreateProperties_ServerRole `json:"serverRole,omitempty"`
	Tags       map[string]string                             `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &Redis_LinkedServer_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (server *Redis_LinkedServer_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if server == nil {
		return nil, nil
	}
	result := &Redis_LinkedServer_Spec_ARM{}

	// Set property ‘Location’:
	if server.Location != nil {
		location := *server.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if server.LinkedRedisCacheLocation != nil ||
		server.LinkedRedisCacheReference != nil ||
		server.ServerRole != nil {
		result.Properties = &RedisLinkedServerCreateProperties_ARM{}
	}
	if server.LinkedRedisCacheLocation != nil {
		linkedRedisCacheLocation := *server.LinkedRedisCacheLocation
		result.Properties.LinkedRedisCacheLocation = &linkedRedisCacheLocation
	}
	if server.LinkedRedisCacheReference != nil {
		linkedRedisCacheIdARMID, err := resolved.ResolvedReferences.Lookup(*server.LinkedRedisCacheReference)
		if err != nil {
			return nil, err
		}
		linkedRedisCacheId := linkedRedisCacheIdARMID
		result.Properties.LinkedRedisCacheId = &linkedRedisCacheId
	}
	if server.ServerRole != nil {
		serverRole := *server.ServerRole
		result.Properties.ServerRole = &serverRole
	}

	// Set property ‘Tags’:
	if server.Tags != nil {
		result.Tags = make(map[string]string, len(server.Tags))
		for key, value := range server.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (server *Redis_LinkedServer_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Redis_LinkedServer_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (server *Redis_LinkedServer_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Redis_LinkedServer_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Redis_LinkedServer_Spec_ARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	server.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘LinkedRedisCacheLocation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.LinkedRedisCacheLocation != nil {
			linkedRedisCacheLocation := *typedInput.Properties.LinkedRedisCacheLocation
			server.LinkedRedisCacheLocation = &linkedRedisCacheLocation
		}
	}

	// no assignment for property ‘LinkedRedisCacheReference’

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		server.Location = &location
	}

	// Set property ‘Owner’:
	server.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘ServerRole’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ServerRole != nil {
			serverRole := *typedInput.Properties.ServerRole
			server.ServerRole = &serverRole
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		server.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			server.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Redis_LinkedServer_Spec{}

// ConvertSpecFrom populates our Redis_LinkedServer_Spec from the provided source
func (server *Redis_LinkedServer_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20201201s.Redis_LinkedServer_Spec)
	if ok {
		// Populate our instance from source
		return server.AssignProperties_From_Redis_LinkedServer_Spec(src)
	}

	// Convert to an intermediate form
	src = &alpha20201201s.Redis_LinkedServer_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = server.AssignProperties_From_Redis_LinkedServer_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Redis_LinkedServer_Spec
func (server *Redis_LinkedServer_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20201201s.Redis_LinkedServer_Spec)
	if ok {
		// Populate destination from our instance
		return server.AssignProperties_To_Redis_LinkedServer_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20201201s.Redis_LinkedServer_Spec{}
	err := server.AssignProperties_To_Redis_LinkedServer_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_Redis_LinkedServer_Spec populates our Redis_LinkedServer_Spec from the provided source Redis_LinkedServer_Spec
func (server *Redis_LinkedServer_Spec) AssignProperties_From_Redis_LinkedServer_Spec(source *alpha20201201s.Redis_LinkedServer_Spec) error {

	// AzureName
	server.AzureName = source.AzureName

	// LinkedRedisCacheLocation
	server.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	if source.LinkedRedisCacheReference != nil {
		linkedRedisCacheReference := source.LinkedRedisCacheReference.Copy()
		server.LinkedRedisCacheReference = &linkedRedisCacheReference
	} else {
		server.LinkedRedisCacheReference = nil
	}

	// Location
	server.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		server.Owner = &owner
	} else {
		server.Owner = nil
	}

	// ServerRole
	if source.ServerRole != nil {
		serverRole := RedisLinkedServerCreateProperties_ServerRole(*source.ServerRole)
		server.ServerRole = &serverRole
	} else {
		server.ServerRole = nil
	}

	// Tags
	server.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_Redis_LinkedServer_Spec populates the provided destination Redis_LinkedServer_Spec from our Redis_LinkedServer_Spec
func (server *Redis_LinkedServer_Spec) AssignProperties_To_Redis_LinkedServer_Spec(destination *alpha20201201s.Redis_LinkedServer_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = server.AzureName

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(server.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	if server.LinkedRedisCacheReference != nil {
		linkedRedisCacheReference := server.LinkedRedisCacheReference.Copy()
		destination.LinkedRedisCacheReference = &linkedRedisCacheReference
	} else {
		destination.LinkedRedisCacheReference = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(server.Location)

	// OriginalVersion
	destination.OriginalVersion = server.OriginalVersion()

	// Owner
	if server.Owner != nil {
		owner := server.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// ServerRole
	if server.ServerRole != nil {
		serverRole := string(*server.ServerRole)
		destination.ServerRole = &serverRole
	} else {
		destination.ServerRole = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(server.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (server *Redis_LinkedServer_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (server *Redis_LinkedServer_Spec) SetAzureName(azureName string) { server.AzureName = azureName }

// Deprecated version of RedisLinkedServerWithProperties_STATUS. Use v1beta20201201.RedisLinkedServerWithProperties_STATUS instead
type RedisLinkedServerWithProperties_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions               []conditions.Condition                         `json:"conditions,omitempty"`
	Id                       *string                                        `json:"id,omitempty"`
	LinkedRedisCacheId       *string                                        `json:"linkedRedisCacheId,omitempty"`
	LinkedRedisCacheLocation *string                                        `json:"linkedRedisCacheLocation,omitempty"`
	Name                     *string                                        `json:"name,omitempty"`
	ProvisioningState        *string                                        `json:"provisioningState,omitempty"`
	ServerRole               *RedisLinkedServerProperties_ServerRole_STATUS `json:"serverRole,omitempty"`
	Type                     *string                                        `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisLinkedServerWithProperties_STATUS{}

// ConvertStatusFrom populates our RedisLinkedServerWithProperties_STATUS from the provided source
func (properties *RedisLinkedServerWithProperties_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20201201s.RedisLinkedServerWithProperties_STATUS)
	if ok {
		// Populate our instance from source
		return properties.AssignProperties_From_RedisLinkedServerWithProperties_STATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20201201s.RedisLinkedServerWithProperties_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = properties.AssignProperties_From_RedisLinkedServerWithProperties_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisLinkedServerWithProperties_STATUS
func (properties *RedisLinkedServerWithProperties_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20201201s.RedisLinkedServerWithProperties_STATUS)
	if ok {
		// Populate destination from our instance
		return properties.AssignProperties_To_RedisLinkedServerWithProperties_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20201201s.RedisLinkedServerWithProperties_STATUS{}
	err := properties.AssignProperties_To_RedisLinkedServerWithProperties_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &RedisLinkedServerWithProperties_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (properties *RedisLinkedServerWithProperties_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisLinkedServerWithProperties_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (properties *RedisLinkedServerWithProperties_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisLinkedServerWithProperties_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisLinkedServerWithProperties_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		properties.Id = &id
	}

	// Set property ‘LinkedRedisCacheId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.LinkedRedisCacheId != nil {
			linkedRedisCacheId := *typedInput.Properties.LinkedRedisCacheId
			properties.LinkedRedisCacheId = &linkedRedisCacheId
		}
	}

	// Set property ‘LinkedRedisCacheLocation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.LinkedRedisCacheLocation != nil {
			linkedRedisCacheLocation := *typedInput.Properties.LinkedRedisCacheLocation
			properties.LinkedRedisCacheLocation = &linkedRedisCacheLocation
		}
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		properties.Name = &name
	}

	// Set property ‘ProvisioningState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			properties.ProvisioningState = &provisioningState
		}
	}

	// Set property ‘ServerRole’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ServerRole != nil {
			serverRole := *typedInput.Properties.ServerRole
			properties.ServerRole = &serverRole
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		properties.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_RedisLinkedServerWithProperties_STATUS populates our RedisLinkedServerWithProperties_STATUS from the provided source RedisLinkedServerWithProperties_STATUS
func (properties *RedisLinkedServerWithProperties_STATUS) AssignProperties_From_RedisLinkedServerWithProperties_STATUS(source *alpha20201201s.RedisLinkedServerWithProperties_STATUS) error {

	// Conditions
	properties.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	properties.Id = genruntime.ClonePointerToString(source.Id)

	// LinkedRedisCacheId
	properties.LinkedRedisCacheId = genruntime.ClonePointerToString(source.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	properties.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// Name
	properties.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	properties.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// ServerRole
	if source.ServerRole != nil {
		serverRole := RedisLinkedServerProperties_ServerRole_STATUS(*source.ServerRole)
		properties.ServerRole = &serverRole
	} else {
		properties.ServerRole = nil
	}

	// Type
	properties.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_RedisLinkedServerWithProperties_STATUS populates the provided destination RedisLinkedServerWithProperties_STATUS from our RedisLinkedServerWithProperties_STATUS
func (properties *RedisLinkedServerWithProperties_STATUS) AssignProperties_To_RedisLinkedServerWithProperties_STATUS(destination *alpha20201201s.RedisLinkedServerWithProperties_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(properties.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(properties.Id)

	// LinkedRedisCacheId
	destination.LinkedRedisCacheId = genruntime.ClonePointerToString(properties.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(properties.LinkedRedisCacheLocation)

	// Name
	destination.Name = genruntime.ClonePointerToString(properties.Name)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(properties.ProvisioningState)

	// ServerRole
	if properties.ServerRole != nil {
		serverRole := string(*properties.ServerRole)
		destination.ServerRole = &serverRole
	} else {
		destination.ServerRole = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(properties.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Deprecated version of RedisLinkedServerCreateProperties_ServerRole. Use
// v1beta20201201.RedisLinkedServerCreateProperties_ServerRole instead
// +kubebuilder:validation:Enum={"Primary","Secondary"}
type RedisLinkedServerCreateProperties_ServerRole string

const (
	RedisLinkedServerCreateProperties_ServerRole_Primary   = RedisLinkedServerCreateProperties_ServerRole("Primary")
	RedisLinkedServerCreateProperties_ServerRole_Secondary = RedisLinkedServerCreateProperties_ServerRole("Secondary")
)

// Deprecated version of RedisLinkedServerProperties_ServerRole_STATUS. Use
// v1beta20201201.RedisLinkedServerProperties_ServerRole_STATUS instead
type RedisLinkedServerProperties_ServerRole_STATUS string

const (
	RedisLinkedServerProperties_ServerRole_STATUS_Primary   = RedisLinkedServerProperties_ServerRole_STATUS("Primary")
	RedisLinkedServerProperties_ServerRole_STATUS_Secondary = RedisLinkedServerProperties_ServerRole_STATUS("Secondary")
)

func init() {
	SchemeBuilder.Register(&RedisLinkedServer{}, &RedisLinkedServerList{})
}
