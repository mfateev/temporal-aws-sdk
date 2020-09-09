package awsclients

import (
	"github.com/aws/aws-sdk-go/service/ram"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type RAMClient interface {
       AcceptResourceShareInvitation(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error)
       AcceptResourceShareInvitationAsync(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) *RamAcceptResourceShareInvitationResult

       AssociateResourceShare(ctx workflow.Context, input *ram.AssociateResourceShareInput) (*ram.AssociateResourceShareOutput, error)
       AssociateResourceShareAsync(ctx workflow.Context, input *ram.AssociateResourceShareInput) *RamAssociateResourceShareResult

       AssociateResourceSharePermission(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) (*ram.AssociateResourceSharePermissionOutput, error)
       AssociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) *RamAssociateResourceSharePermissionResult

       CreateResourceShare(ctx workflow.Context, input *ram.CreateResourceShareInput) (*ram.CreateResourceShareOutput, error)
       CreateResourceShareAsync(ctx workflow.Context, input *ram.CreateResourceShareInput) *RamCreateResourceShareResult

       DeleteResourceShare(ctx workflow.Context, input *ram.DeleteResourceShareInput) (*ram.DeleteResourceShareOutput, error)
       DeleteResourceShareAsync(ctx workflow.Context, input *ram.DeleteResourceShareInput) *RamDeleteResourceShareResult

       DisassociateResourceShare(ctx workflow.Context, input *ram.DisassociateResourceShareInput) (*ram.DisassociateResourceShareOutput, error)
       DisassociateResourceShareAsync(ctx workflow.Context, input *ram.DisassociateResourceShareInput) *RamDisassociateResourceShareResult

       DisassociateResourceSharePermission(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) (*ram.DisassociateResourceSharePermissionOutput, error)
       DisassociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) *RamDisassociateResourceSharePermissionResult

       EnableSharingWithAwsOrganization(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) (*ram.EnableSharingWithAwsOrganizationOutput, error)
       EnableSharingWithAwsOrganizationAsync(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) *RamEnableSharingWithAwsOrganizationResult

       GetPermission(ctx workflow.Context, input *ram.GetPermissionInput) (*ram.GetPermissionOutput, error)
       GetPermissionAsync(ctx workflow.Context, input *ram.GetPermissionInput) *RamGetPermissionResult

       GetResourcePolicies(ctx workflow.Context, input *ram.GetResourcePoliciesInput) (*ram.GetResourcePoliciesOutput, error)
       GetResourcePoliciesAsync(ctx workflow.Context, input *ram.GetResourcePoliciesInput) *RamGetResourcePoliciesResult

       GetResourceShareAssociations(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) (*ram.GetResourceShareAssociationsOutput, error)
       GetResourceShareAssociationsAsync(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) *RamGetResourceShareAssociationsResult

       GetResourceShareInvitations(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) (*ram.GetResourceShareInvitationsOutput, error)
       GetResourceShareInvitationsAsync(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) *RamGetResourceShareInvitationsResult

       GetResourceShares(ctx workflow.Context, input *ram.GetResourceSharesInput) (*ram.GetResourceSharesOutput, error)
       GetResourceSharesAsync(ctx workflow.Context, input *ram.GetResourceSharesInput) *RamGetResourceSharesResult

       ListPendingInvitationResources(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) (*ram.ListPendingInvitationResourcesOutput, error)
       ListPendingInvitationResourcesAsync(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) *RamListPendingInvitationResourcesResult

       ListPermissions(ctx workflow.Context, input *ram.ListPermissionsInput) (*ram.ListPermissionsOutput, error)
       ListPermissionsAsync(ctx workflow.Context, input *ram.ListPermissionsInput) *RamListPermissionsResult

       ListPrincipals(ctx workflow.Context, input *ram.ListPrincipalsInput) (*ram.ListPrincipalsOutput, error)
       ListPrincipalsAsync(ctx workflow.Context, input *ram.ListPrincipalsInput) *RamListPrincipalsResult

       ListResourceSharePermissions(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) (*ram.ListResourceSharePermissionsOutput, error)
       ListResourceSharePermissionsAsync(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) *RamListResourceSharePermissionsResult

       ListResourceTypes(ctx workflow.Context, input *ram.ListResourceTypesInput) (*ram.ListResourceTypesOutput, error)
       ListResourceTypesAsync(ctx workflow.Context, input *ram.ListResourceTypesInput) *RamListResourceTypesResult

       ListResources(ctx workflow.Context, input *ram.ListResourcesInput) (*ram.ListResourcesOutput, error)
       ListResourcesAsync(ctx workflow.Context, input *ram.ListResourcesInput) *RamListResourcesResult

       PromoteResourceShareCreatedFromPolicy(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error)
       PromoteResourceShareCreatedFromPolicyAsync(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) *RamPromoteResourceShareCreatedFromPolicyResult

       RejectResourceShareInvitation(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) (*ram.RejectResourceShareInvitationOutput, error)
       RejectResourceShareInvitationAsync(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) *RamRejectResourceShareInvitationResult

       TagResource(ctx workflow.Context, input *ram.TagResourceInput) (*ram.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *ram.TagResourceInput) *RamTagResourceResult

       UntagResource(ctx workflow.Context, input *ram.UntagResourceInput) (*ram.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *ram.UntagResourceInput) *RamUntagResourceResult

       UpdateResourceShare(ctx workflow.Context, input *ram.UpdateResourceShareInput) (*ram.UpdateResourceShareOutput, error)
       UpdateResourceShareAsync(ctx workflow.Context, input *ram.UpdateResourceShareInput) *RamUpdateResourceShareResult
}

type RamAcceptResourceShareInvitationResult struct {
	Result workflow.Future
}

func (r *RamAcceptResourceShareInvitationResult) Get(ctx workflow.Context) (*ram.AcceptResourceShareInvitationOutput, error) {
    var output ram.AcceptResourceShareInvitationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamAssociateResourceShareResult struct {
	Result workflow.Future
}

func (r *RamAssociateResourceShareResult) Get(ctx workflow.Context) (*ram.AssociateResourceShareOutput, error) {
    var output ram.AssociateResourceShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamAssociateResourceSharePermissionResult struct {
	Result workflow.Future
}

func (r *RamAssociateResourceSharePermissionResult) Get(ctx workflow.Context) (*ram.AssociateResourceSharePermissionOutput, error) {
    var output ram.AssociateResourceSharePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamCreateResourceShareResult struct {
	Result workflow.Future
}

func (r *RamCreateResourceShareResult) Get(ctx workflow.Context) (*ram.CreateResourceShareOutput, error) {
    var output ram.CreateResourceShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamDeleteResourceShareResult struct {
	Result workflow.Future
}

func (r *RamDeleteResourceShareResult) Get(ctx workflow.Context) (*ram.DeleteResourceShareOutput, error) {
    var output ram.DeleteResourceShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamDisassociateResourceShareResult struct {
	Result workflow.Future
}

func (r *RamDisassociateResourceShareResult) Get(ctx workflow.Context) (*ram.DisassociateResourceShareOutput, error) {
    var output ram.DisassociateResourceShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamDisassociateResourceSharePermissionResult struct {
	Result workflow.Future
}

func (r *RamDisassociateResourceSharePermissionResult) Get(ctx workflow.Context) (*ram.DisassociateResourceSharePermissionOutput, error) {
    var output ram.DisassociateResourceSharePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamEnableSharingWithAwsOrganizationResult struct {
	Result workflow.Future
}

func (r *RamEnableSharingWithAwsOrganizationResult) Get(ctx workflow.Context) (*ram.EnableSharingWithAwsOrganizationOutput, error) {
    var output ram.EnableSharingWithAwsOrganizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamGetPermissionResult struct {
	Result workflow.Future
}

func (r *RamGetPermissionResult) Get(ctx workflow.Context) (*ram.GetPermissionOutput, error) {
    var output ram.GetPermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamGetResourcePoliciesResult struct {
	Result workflow.Future
}

func (r *RamGetResourcePoliciesResult) Get(ctx workflow.Context) (*ram.GetResourcePoliciesOutput, error) {
    var output ram.GetResourcePoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamGetResourceShareAssociationsResult struct {
	Result workflow.Future
}

func (r *RamGetResourceShareAssociationsResult) Get(ctx workflow.Context) (*ram.GetResourceShareAssociationsOutput, error) {
    var output ram.GetResourceShareAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamGetResourceShareInvitationsResult struct {
	Result workflow.Future
}

func (r *RamGetResourceShareInvitationsResult) Get(ctx workflow.Context) (*ram.GetResourceShareInvitationsOutput, error) {
    var output ram.GetResourceShareInvitationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamGetResourceSharesResult struct {
	Result workflow.Future
}

func (r *RamGetResourceSharesResult) Get(ctx workflow.Context) (*ram.GetResourceSharesOutput, error) {
    var output ram.GetResourceSharesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamListPendingInvitationResourcesResult struct {
	Result workflow.Future
}

func (r *RamListPendingInvitationResourcesResult) Get(ctx workflow.Context) (*ram.ListPendingInvitationResourcesOutput, error) {
    var output ram.ListPendingInvitationResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamListPermissionsResult struct {
	Result workflow.Future
}

func (r *RamListPermissionsResult) Get(ctx workflow.Context) (*ram.ListPermissionsOutput, error) {
    var output ram.ListPermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamListPrincipalsResult struct {
	Result workflow.Future
}

func (r *RamListPrincipalsResult) Get(ctx workflow.Context) (*ram.ListPrincipalsOutput, error) {
    var output ram.ListPrincipalsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamListResourceSharePermissionsResult struct {
	Result workflow.Future
}

func (r *RamListResourceSharePermissionsResult) Get(ctx workflow.Context) (*ram.ListResourceSharePermissionsOutput, error) {
    var output ram.ListResourceSharePermissionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamListResourceTypesResult struct {
	Result workflow.Future
}

func (r *RamListResourceTypesResult) Get(ctx workflow.Context) (*ram.ListResourceTypesOutput, error) {
    var output ram.ListResourceTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamListResourcesResult struct {
	Result workflow.Future
}

func (r *RamListResourcesResult) Get(ctx workflow.Context) (*ram.ListResourcesOutput, error) {
    var output ram.ListResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamPromoteResourceShareCreatedFromPolicyResult struct {
	Result workflow.Future
}

func (r *RamPromoteResourceShareCreatedFromPolicyResult) Get(ctx workflow.Context) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error) {
    var output ram.PromoteResourceShareCreatedFromPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamRejectResourceShareInvitationResult struct {
	Result workflow.Future
}

func (r *RamRejectResourceShareInvitationResult) Get(ctx workflow.Context) (*ram.RejectResourceShareInvitationOutput, error) {
    var output ram.RejectResourceShareInvitationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamTagResourceResult struct {
	Result workflow.Future
}

func (r *RamTagResourceResult) Get(ctx workflow.Context) (*ram.TagResourceOutput, error) {
    var output ram.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamUntagResourceResult struct {
	Result workflow.Future
}

func (r *RamUntagResourceResult) Get(ctx workflow.Context) (*ram.UntagResourceOutput, error) {
    var output ram.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RamUpdateResourceShareResult struct {
	Result workflow.Future
}

func (r *RamUpdateResourceShareResult) Get(ctx workflow.Context) (*ram.UpdateResourceShareOutput, error) {
    var output ram.UpdateResourceShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RAMStub struct {
    activities awsactivities.RAMActivities
}

func NewRAMStub() RAMClient {
    return &RAMStub{}
}

func (a *RAMStub) AcceptResourceShareInvitation(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error) {
    var output ram.AcceptResourceShareInvitationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptResourceShareInvitation, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) AcceptResourceShareInvitationAsync(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) *RamAcceptResourceShareInvitationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptResourceShareInvitation, input)
    return &RamAcceptResourceShareInvitationResult{Result: future}
}

func (a *RAMStub) AssociateResourceShare(ctx workflow.Context, input *ram.AssociateResourceShareInput) (*ram.AssociateResourceShareOutput, error) {
    var output ram.AssociateResourceShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateResourceShare, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) AssociateResourceShareAsync(ctx workflow.Context, input *ram.AssociateResourceShareInput) *RamAssociateResourceShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateResourceShare, input)
    return &RamAssociateResourceShareResult{Result: future}
}

func (a *RAMStub) AssociateResourceSharePermission(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) (*ram.AssociateResourceSharePermissionOutput, error) {
    var output ram.AssociateResourceSharePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateResourceSharePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) AssociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) *RamAssociateResourceSharePermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateResourceSharePermission, input)
    return &RamAssociateResourceSharePermissionResult{Result: future}
}

func (a *RAMStub) CreateResourceShare(ctx workflow.Context, input *ram.CreateResourceShareInput) (*ram.CreateResourceShareOutput, error) {
    var output ram.CreateResourceShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateResourceShare, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) CreateResourceShareAsync(ctx workflow.Context, input *ram.CreateResourceShareInput) *RamCreateResourceShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateResourceShare, input)
    return &RamCreateResourceShareResult{Result: future}
}

func (a *RAMStub) DeleteResourceShare(ctx workflow.Context, input *ram.DeleteResourceShareInput) (*ram.DeleteResourceShareOutput, error) {
    var output ram.DeleteResourceShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResourceShare, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) DeleteResourceShareAsync(ctx workflow.Context, input *ram.DeleteResourceShareInput) *RamDeleteResourceShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteResourceShare, input)
    return &RamDeleteResourceShareResult{Result: future}
}

func (a *RAMStub) DisassociateResourceShare(ctx workflow.Context, input *ram.DisassociateResourceShareInput) (*ram.DisassociateResourceShareOutput, error) {
    var output ram.DisassociateResourceShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateResourceShare, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) DisassociateResourceShareAsync(ctx workflow.Context, input *ram.DisassociateResourceShareInput) *RamDisassociateResourceShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateResourceShare, input)
    return &RamDisassociateResourceShareResult{Result: future}
}

func (a *RAMStub) DisassociateResourceSharePermission(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) (*ram.DisassociateResourceSharePermissionOutput, error) {
    var output ram.DisassociateResourceSharePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateResourceSharePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) DisassociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) *RamDisassociateResourceSharePermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateResourceSharePermission, input)
    return &RamDisassociateResourceSharePermissionResult{Result: future}
}

func (a *RAMStub) EnableSharingWithAwsOrganization(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) (*ram.EnableSharingWithAwsOrganizationOutput, error) {
    var output ram.EnableSharingWithAwsOrganizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableSharingWithAwsOrganization, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) EnableSharingWithAwsOrganizationAsync(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) *RamEnableSharingWithAwsOrganizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableSharingWithAwsOrganization, input)
    return &RamEnableSharingWithAwsOrganizationResult{Result: future}
}

func (a *RAMStub) GetPermission(ctx workflow.Context, input *ram.GetPermissionInput) (*ram.GetPermissionOutput, error) {
    var output ram.GetPermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPermission, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) GetPermissionAsync(ctx workflow.Context, input *ram.GetPermissionInput) *RamGetPermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPermission, input)
    return &RamGetPermissionResult{Result: future}
}

func (a *RAMStub) GetResourcePolicies(ctx workflow.Context, input *ram.GetResourcePoliciesInput) (*ram.GetResourcePoliciesOutput, error) {
    var output ram.GetResourcePoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResourcePolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) GetResourcePoliciesAsync(ctx workflow.Context, input *ram.GetResourcePoliciesInput) *RamGetResourcePoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResourcePolicies, input)
    return &RamGetResourcePoliciesResult{Result: future}
}

func (a *RAMStub) GetResourceShareAssociations(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) (*ram.GetResourceShareAssociationsOutput, error) {
    var output ram.GetResourceShareAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResourceShareAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) GetResourceShareAssociationsAsync(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) *RamGetResourceShareAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResourceShareAssociations, input)
    return &RamGetResourceShareAssociationsResult{Result: future}
}

func (a *RAMStub) GetResourceShareInvitations(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) (*ram.GetResourceShareInvitationsOutput, error) {
    var output ram.GetResourceShareInvitationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResourceShareInvitations, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) GetResourceShareInvitationsAsync(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) *RamGetResourceShareInvitationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResourceShareInvitations, input)
    return &RamGetResourceShareInvitationsResult{Result: future}
}

func (a *RAMStub) GetResourceShares(ctx workflow.Context, input *ram.GetResourceSharesInput) (*ram.GetResourceSharesOutput, error) {
    var output ram.GetResourceSharesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResourceShares, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) GetResourceSharesAsync(ctx workflow.Context, input *ram.GetResourceSharesInput) *RamGetResourceSharesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResourceShares, input)
    return &RamGetResourceSharesResult{Result: future}
}

func (a *RAMStub) ListPendingInvitationResources(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) (*ram.ListPendingInvitationResourcesOutput, error) {
    var output ram.ListPendingInvitationResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPendingInvitationResources, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) ListPendingInvitationResourcesAsync(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) *RamListPendingInvitationResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPendingInvitationResources, input)
    return &RamListPendingInvitationResourcesResult{Result: future}
}

func (a *RAMStub) ListPermissions(ctx workflow.Context, input *ram.ListPermissionsInput) (*ram.ListPermissionsOutput, error) {
    var output ram.ListPermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) ListPermissionsAsync(ctx workflow.Context, input *ram.ListPermissionsInput) *RamListPermissionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPermissions, input)
    return &RamListPermissionsResult{Result: future}
}

func (a *RAMStub) ListPrincipals(ctx workflow.Context, input *ram.ListPrincipalsInput) (*ram.ListPrincipalsOutput, error) {
    var output ram.ListPrincipalsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPrincipals, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) ListPrincipalsAsync(ctx workflow.Context, input *ram.ListPrincipalsInput) *RamListPrincipalsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPrincipals, input)
    return &RamListPrincipalsResult{Result: future}
}

func (a *RAMStub) ListResourceSharePermissions(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) (*ram.ListResourceSharePermissionsOutput, error) {
    var output ram.ListResourceSharePermissionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourceSharePermissions, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) ListResourceSharePermissionsAsync(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) *RamListResourceSharePermissionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResourceSharePermissions, input)
    return &RamListResourceSharePermissionsResult{Result: future}
}

func (a *RAMStub) ListResourceTypes(ctx workflow.Context, input *ram.ListResourceTypesInput) (*ram.ListResourceTypesOutput, error) {
    var output ram.ListResourceTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourceTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) ListResourceTypesAsync(ctx workflow.Context, input *ram.ListResourceTypesInput) *RamListResourceTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResourceTypes, input)
    return &RamListResourceTypesResult{Result: future}
}

func (a *RAMStub) ListResources(ctx workflow.Context, input *ram.ListResourcesInput) (*ram.ListResourcesOutput, error) {
    var output ram.ListResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResources, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) ListResourcesAsync(ctx workflow.Context, input *ram.ListResourcesInput) *RamListResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResources, input)
    return &RamListResourcesResult{Result: future}
}

func (a *RAMStub) PromoteResourceShareCreatedFromPolicy(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error) {
    var output ram.PromoteResourceShareCreatedFromPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PromoteResourceShareCreatedFromPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) PromoteResourceShareCreatedFromPolicyAsync(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) *RamPromoteResourceShareCreatedFromPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PromoteResourceShareCreatedFromPolicy, input)
    return &RamPromoteResourceShareCreatedFromPolicyResult{Result: future}
}

func (a *RAMStub) RejectResourceShareInvitation(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) (*ram.RejectResourceShareInvitationOutput, error) {
    var output ram.RejectResourceShareInvitationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectResourceShareInvitation, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) RejectResourceShareInvitationAsync(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) *RamRejectResourceShareInvitationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RejectResourceShareInvitation, input)
    return &RamRejectResourceShareInvitationResult{Result: future}
}

func (a *RAMStub) TagResource(ctx workflow.Context, input *ram.TagResourceInput) (*ram.TagResourceOutput, error) {
    var output ram.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) TagResourceAsync(ctx workflow.Context, input *ram.TagResourceInput) *RamTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &RamTagResourceResult{Result: future}
}

func (a *RAMStub) UntagResource(ctx workflow.Context, input *ram.UntagResourceInput) (*ram.UntagResourceOutput, error) {
    var output ram.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) UntagResourceAsync(ctx workflow.Context, input *ram.UntagResourceInput) *RamUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &RamUntagResourceResult{Result: future}
}

func (a *RAMStub) UpdateResourceShare(ctx workflow.Context, input *ram.UpdateResourceShareInput) (*ram.UpdateResourceShareOutput, error) {
    var output ram.UpdateResourceShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateResourceShare, input).Get(ctx, &output)
    return &output, err
}

func (a *RAMStub) UpdateResourceShareAsync(ctx workflow.Context, input *ram.UpdateResourceShareInput) *RamUpdateResourceShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateResourceShare, input)
    return &RamUpdateResourceShareResult{Result: future}
}
