package awsclients

import (
	"github.com/aws/aws-sdk-go/service/organizations"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type OrganizationsClient interface {
       AcceptHandshake(ctx workflow.Context, input *organizations.AcceptHandshakeInput) (*organizations.AcceptHandshakeOutput, error)
       AcceptHandshakeAsync(ctx workflow.Context, input *organizations.AcceptHandshakeInput) *OrganizationsAcceptHandshakeResult

       AttachPolicy(ctx workflow.Context, input *organizations.AttachPolicyInput) (*organizations.AttachPolicyOutput, error)
       AttachPolicyAsync(ctx workflow.Context, input *organizations.AttachPolicyInput) *OrganizationsAttachPolicyResult

       CancelHandshake(ctx workflow.Context, input *organizations.CancelHandshakeInput) (*organizations.CancelHandshakeOutput, error)
       CancelHandshakeAsync(ctx workflow.Context, input *organizations.CancelHandshakeInput) *OrganizationsCancelHandshakeResult

       CreateAccount(ctx workflow.Context, input *organizations.CreateAccountInput) (*organizations.CreateAccountOutput, error)
       CreateAccountAsync(ctx workflow.Context, input *organizations.CreateAccountInput) *OrganizationsCreateAccountResult

       CreateGovCloudAccount(ctx workflow.Context, input *organizations.CreateGovCloudAccountInput) (*organizations.CreateGovCloudAccountOutput, error)
       CreateGovCloudAccountAsync(ctx workflow.Context, input *organizations.CreateGovCloudAccountInput) *OrganizationsCreateGovCloudAccountResult

       CreateOrganization(ctx workflow.Context, input *organizations.CreateOrganizationInput) (*organizations.CreateOrganizationOutput, error)
       CreateOrganizationAsync(ctx workflow.Context, input *organizations.CreateOrganizationInput) *OrganizationsCreateOrganizationResult

       CreateOrganizationalUnit(ctx workflow.Context, input *organizations.CreateOrganizationalUnitInput) (*organizations.CreateOrganizationalUnitOutput, error)
       CreateOrganizationalUnitAsync(ctx workflow.Context, input *organizations.CreateOrganizationalUnitInput) *OrganizationsCreateOrganizationalUnitResult

       CreatePolicy(ctx workflow.Context, input *organizations.CreatePolicyInput) (*organizations.CreatePolicyOutput, error)
       CreatePolicyAsync(ctx workflow.Context, input *organizations.CreatePolicyInput) *OrganizationsCreatePolicyResult

       DeclineHandshake(ctx workflow.Context, input *organizations.DeclineHandshakeInput) (*organizations.DeclineHandshakeOutput, error)
       DeclineHandshakeAsync(ctx workflow.Context, input *organizations.DeclineHandshakeInput) *OrganizationsDeclineHandshakeResult

       DeleteOrganization(ctx workflow.Context, input *organizations.DeleteOrganizationInput) (*organizations.DeleteOrganizationOutput, error)
       DeleteOrganizationAsync(ctx workflow.Context, input *organizations.DeleteOrganizationInput) *OrganizationsDeleteOrganizationResult

       DeleteOrganizationalUnit(ctx workflow.Context, input *organizations.DeleteOrganizationalUnitInput) (*organizations.DeleteOrganizationalUnitOutput, error)
       DeleteOrganizationalUnitAsync(ctx workflow.Context, input *organizations.DeleteOrganizationalUnitInput) *OrganizationsDeleteOrganizationalUnitResult

       DeletePolicy(ctx workflow.Context, input *organizations.DeletePolicyInput) (*organizations.DeletePolicyOutput, error)
       DeletePolicyAsync(ctx workflow.Context, input *organizations.DeletePolicyInput) *OrganizationsDeletePolicyResult

       DeregisterDelegatedAdministrator(ctx workflow.Context, input *organizations.DeregisterDelegatedAdministratorInput) (*organizations.DeregisterDelegatedAdministratorOutput, error)
       DeregisterDelegatedAdministratorAsync(ctx workflow.Context, input *organizations.DeregisterDelegatedAdministratorInput) *OrganizationsDeregisterDelegatedAdministratorResult

       DescribeAccount(ctx workflow.Context, input *organizations.DescribeAccountInput) (*organizations.DescribeAccountOutput, error)
       DescribeAccountAsync(ctx workflow.Context, input *organizations.DescribeAccountInput) *OrganizationsDescribeAccountResult

       DescribeCreateAccountStatus(ctx workflow.Context, input *organizations.DescribeCreateAccountStatusInput) (*organizations.DescribeCreateAccountStatusOutput, error)
       DescribeCreateAccountStatusAsync(ctx workflow.Context, input *organizations.DescribeCreateAccountStatusInput) *OrganizationsDescribeCreateAccountStatusResult

       DescribeEffectivePolicy(ctx workflow.Context, input *organizations.DescribeEffectivePolicyInput) (*organizations.DescribeEffectivePolicyOutput, error)
       DescribeEffectivePolicyAsync(ctx workflow.Context, input *organizations.DescribeEffectivePolicyInput) *OrganizationsDescribeEffectivePolicyResult

       DescribeHandshake(ctx workflow.Context, input *organizations.DescribeHandshakeInput) (*organizations.DescribeHandshakeOutput, error)
       DescribeHandshakeAsync(ctx workflow.Context, input *organizations.DescribeHandshakeInput) *OrganizationsDescribeHandshakeResult

       DescribeOrganization(ctx workflow.Context, input *organizations.DescribeOrganizationInput) (*organizations.DescribeOrganizationOutput, error)
       DescribeOrganizationAsync(ctx workflow.Context, input *organizations.DescribeOrganizationInput) *OrganizationsDescribeOrganizationResult

       DescribeOrganizationalUnit(ctx workflow.Context, input *organizations.DescribeOrganizationalUnitInput) (*organizations.DescribeOrganizationalUnitOutput, error)
       DescribeOrganizationalUnitAsync(ctx workflow.Context, input *organizations.DescribeOrganizationalUnitInput) *OrganizationsDescribeOrganizationalUnitResult

       DescribePolicy(ctx workflow.Context, input *organizations.DescribePolicyInput) (*organizations.DescribePolicyOutput, error)
       DescribePolicyAsync(ctx workflow.Context, input *organizations.DescribePolicyInput) *OrganizationsDescribePolicyResult

       DetachPolicy(ctx workflow.Context, input *organizations.DetachPolicyInput) (*organizations.DetachPolicyOutput, error)
       DetachPolicyAsync(ctx workflow.Context, input *organizations.DetachPolicyInput) *OrganizationsDetachPolicyResult

       DisableAWSServiceAccess(ctx workflow.Context, input *organizations.DisableAWSServiceAccessInput) (*organizations.DisableAWSServiceAccessOutput, error)
       DisableAWSServiceAccessAsync(ctx workflow.Context, input *organizations.DisableAWSServiceAccessInput) *OrganizationsDisableAWSServiceAccessResult

       DisablePolicyType(ctx workflow.Context, input *organizations.DisablePolicyTypeInput) (*organizations.DisablePolicyTypeOutput, error)
       DisablePolicyTypeAsync(ctx workflow.Context, input *organizations.DisablePolicyTypeInput) *OrganizationsDisablePolicyTypeResult

       EnableAWSServiceAccess(ctx workflow.Context, input *organizations.EnableAWSServiceAccessInput) (*organizations.EnableAWSServiceAccessOutput, error)
       EnableAWSServiceAccessAsync(ctx workflow.Context, input *organizations.EnableAWSServiceAccessInput) *OrganizationsEnableAWSServiceAccessResult

       EnableAllFeatures(ctx workflow.Context, input *organizations.EnableAllFeaturesInput) (*organizations.EnableAllFeaturesOutput, error)
       EnableAllFeaturesAsync(ctx workflow.Context, input *organizations.EnableAllFeaturesInput) *OrganizationsEnableAllFeaturesResult

       EnablePolicyType(ctx workflow.Context, input *organizations.EnablePolicyTypeInput) (*organizations.EnablePolicyTypeOutput, error)
       EnablePolicyTypeAsync(ctx workflow.Context, input *organizations.EnablePolicyTypeInput) *OrganizationsEnablePolicyTypeResult

       InviteAccountToOrganization(ctx workflow.Context, input *organizations.InviteAccountToOrganizationInput) (*organizations.InviteAccountToOrganizationOutput, error)
       InviteAccountToOrganizationAsync(ctx workflow.Context, input *organizations.InviteAccountToOrganizationInput) *OrganizationsInviteAccountToOrganizationResult

       LeaveOrganization(ctx workflow.Context, input *organizations.LeaveOrganizationInput) (*organizations.LeaveOrganizationOutput, error)
       LeaveOrganizationAsync(ctx workflow.Context, input *organizations.LeaveOrganizationInput) *OrganizationsLeaveOrganizationResult

       ListAWSServiceAccessForOrganization(ctx workflow.Context, input *organizations.ListAWSServiceAccessForOrganizationInput) (*organizations.ListAWSServiceAccessForOrganizationOutput, error)
       ListAWSServiceAccessForOrganizationAsync(ctx workflow.Context, input *organizations.ListAWSServiceAccessForOrganizationInput) *OrganizationsListAWSServiceAccessForOrganizationResult

       ListAccounts(ctx workflow.Context, input *organizations.ListAccountsInput) (*organizations.ListAccountsOutput, error)
       ListAccountsAsync(ctx workflow.Context, input *organizations.ListAccountsInput) *OrganizationsListAccountsResult

       ListAccountsForParent(ctx workflow.Context, input *organizations.ListAccountsForParentInput) (*organizations.ListAccountsForParentOutput, error)
       ListAccountsForParentAsync(ctx workflow.Context, input *organizations.ListAccountsForParentInput) *OrganizationsListAccountsForParentResult

       ListChildren(ctx workflow.Context, input *organizations.ListChildrenInput) (*organizations.ListChildrenOutput, error)
       ListChildrenAsync(ctx workflow.Context, input *organizations.ListChildrenInput) *OrganizationsListChildrenResult

       ListCreateAccountStatus(ctx workflow.Context, input *organizations.ListCreateAccountStatusInput) (*organizations.ListCreateAccountStatusOutput, error)
       ListCreateAccountStatusAsync(ctx workflow.Context, input *organizations.ListCreateAccountStatusInput) *OrganizationsListCreateAccountStatusResult

       ListDelegatedAdministrators(ctx workflow.Context, input *organizations.ListDelegatedAdministratorsInput) (*organizations.ListDelegatedAdministratorsOutput, error)
       ListDelegatedAdministratorsAsync(ctx workflow.Context, input *organizations.ListDelegatedAdministratorsInput) *OrganizationsListDelegatedAdministratorsResult

       ListDelegatedServicesForAccount(ctx workflow.Context, input *organizations.ListDelegatedServicesForAccountInput) (*organizations.ListDelegatedServicesForAccountOutput, error)
       ListDelegatedServicesForAccountAsync(ctx workflow.Context, input *organizations.ListDelegatedServicesForAccountInput) *OrganizationsListDelegatedServicesForAccountResult

       ListHandshakesForAccount(ctx workflow.Context, input *organizations.ListHandshakesForAccountInput) (*organizations.ListHandshakesForAccountOutput, error)
       ListHandshakesForAccountAsync(ctx workflow.Context, input *organizations.ListHandshakesForAccountInput) *OrganizationsListHandshakesForAccountResult

       ListHandshakesForOrganization(ctx workflow.Context, input *organizations.ListHandshakesForOrganizationInput) (*organizations.ListHandshakesForOrganizationOutput, error)
       ListHandshakesForOrganizationAsync(ctx workflow.Context, input *organizations.ListHandshakesForOrganizationInput) *OrganizationsListHandshakesForOrganizationResult

       ListOrganizationalUnitsForParent(ctx workflow.Context, input *organizations.ListOrganizationalUnitsForParentInput) (*organizations.ListOrganizationalUnitsForParentOutput, error)
       ListOrganizationalUnitsForParentAsync(ctx workflow.Context, input *organizations.ListOrganizationalUnitsForParentInput) *OrganizationsListOrganizationalUnitsForParentResult

       ListParents(ctx workflow.Context, input *organizations.ListParentsInput) (*organizations.ListParentsOutput, error)
       ListParentsAsync(ctx workflow.Context, input *organizations.ListParentsInput) *OrganizationsListParentsResult

       ListPolicies(ctx workflow.Context, input *organizations.ListPoliciesInput) (*organizations.ListPoliciesOutput, error)
       ListPoliciesAsync(ctx workflow.Context, input *organizations.ListPoliciesInput) *OrganizationsListPoliciesResult

       ListPoliciesForTarget(ctx workflow.Context, input *organizations.ListPoliciesForTargetInput) (*organizations.ListPoliciesForTargetOutput, error)
       ListPoliciesForTargetAsync(ctx workflow.Context, input *organizations.ListPoliciesForTargetInput) *OrganizationsListPoliciesForTargetResult

       ListRoots(ctx workflow.Context, input *organizations.ListRootsInput) (*organizations.ListRootsOutput, error)
       ListRootsAsync(ctx workflow.Context, input *organizations.ListRootsInput) *OrganizationsListRootsResult

       ListTagsForResource(ctx workflow.Context, input *organizations.ListTagsForResourceInput) (*organizations.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *organizations.ListTagsForResourceInput) *OrganizationsListTagsForResourceResult

       ListTargetsForPolicy(ctx workflow.Context, input *organizations.ListTargetsForPolicyInput) (*organizations.ListTargetsForPolicyOutput, error)
       ListTargetsForPolicyAsync(ctx workflow.Context, input *organizations.ListTargetsForPolicyInput) *OrganizationsListTargetsForPolicyResult

       MoveAccount(ctx workflow.Context, input *organizations.MoveAccountInput) (*organizations.MoveAccountOutput, error)
       MoveAccountAsync(ctx workflow.Context, input *organizations.MoveAccountInput) *OrganizationsMoveAccountResult

       RegisterDelegatedAdministrator(ctx workflow.Context, input *organizations.RegisterDelegatedAdministratorInput) (*organizations.RegisterDelegatedAdministratorOutput, error)
       RegisterDelegatedAdministratorAsync(ctx workflow.Context, input *organizations.RegisterDelegatedAdministratorInput) *OrganizationsRegisterDelegatedAdministratorResult

       RemoveAccountFromOrganization(ctx workflow.Context, input *organizations.RemoveAccountFromOrganizationInput) (*organizations.RemoveAccountFromOrganizationOutput, error)
       RemoveAccountFromOrganizationAsync(ctx workflow.Context, input *organizations.RemoveAccountFromOrganizationInput) *OrganizationsRemoveAccountFromOrganizationResult

       TagResource(ctx workflow.Context, input *organizations.TagResourceInput) (*organizations.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *organizations.TagResourceInput) *OrganizationsTagResourceResult

       UntagResource(ctx workflow.Context, input *organizations.UntagResourceInput) (*organizations.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *organizations.UntagResourceInput) *OrganizationsUntagResourceResult

       UpdateOrganizationalUnit(ctx workflow.Context, input *organizations.UpdateOrganizationalUnitInput) (*organizations.UpdateOrganizationalUnitOutput, error)
       UpdateOrganizationalUnitAsync(ctx workflow.Context, input *organizations.UpdateOrganizationalUnitInput) *OrganizationsUpdateOrganizationalUnitResult

       UpdatePolicy(ctx workflow.Context, input *organizations.UpdatePolicyInput) (*organizations.UpdatePolicyOutput, error)
       UpdatePolicyAsync(ctx workflow.Context, input *organizations.UpdatePolicyInput) *OrganizationsUpdatePolicyResult
}

type OrganizationsAcceptHandshakeResult struct {
	Result workflow.Future
}

func (r *OrganizationsAcceptHandshakeResult) Get(ctx workflow.Context) (*organizations.AcceptHandshakeOutput, error) {
    var output organizations.AcceptHandshakeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsAttachPolicyResult struct {
	Result workflow.Future
}

func (r *OrganizationsAttachPolicyResult) Get(ctx workflow.Context) (*organizations.AttachPolicyOutput, error) {
    var output organizations.AttachPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsCancelHandshakeResult struct {
	Result workflow.Future
}

func (r *OrganizationsCancelHandshakeResult) Get(ctx workflow.Context) (*organizations.CancelHandshakeOutput, error) {
    var output organizations.CancelHandshakeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsCreateAccountResult struct {
	Result workflow.Future
}

func (r *OrganizationsCreateAccountResult) Get(ctx workflow.Context) (*organizations.CreateAccountOutput, error) {
    var output organizations.CreateAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsCreateGovCloudAccountResult struct {
	Result workflow.Future
}

func (r *OrganizationsCreateGovCloudAccountResult) Get(ctx workflow.Context) (*organizations.CreateGovCloudAccountOutput, error) {
    var output organizations.CreateGovCloudAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsCreateOrganizationResult struct {
	Result workflow.Future
}

func (r *OrganizationsCreateOrganizationResult) Get(ctx workflow.Context) (*organizations.CreateOrganizationOutput, error) {
    var output organizations.CreateOrganizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsCreateOrganizationalUnitResult struct {
	Result workflow.Future
}

func (r *OrganizationsCreateOrganizationalUnitResult) Get(ctx workflow.Context) (*organizations.CreateOrganizationalUnitOutput, error) {
    var output organizations.CreateOrganizationalUnitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsCreatePolicyResult struct {
	Result workflow.Future
}

func (r *OrganizationsCreatePolicyResult) Get(ctx workflow.Context) (*organizations.CreatePolicyOutput, error) {
    var output organizations.CreatePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDeclineHandshakeResult struct {
	Result workflow.Future
}

func (r *OrganizationsDeclineHandshakeResult) Get(ctx workflow.Context) (*organizations.DeclineHandshakeOutput, error) {
    var output organizations.DeclineHandshakeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDeleteOrganizationResult struct {
	Result workflow.Future
}

func (r *OrganizationsDeleteOrganizationResult) Get(ctx workflow.Context) (*organizations.DeleteOrganizationOutput, error) {
    var output organizations.DeleteOrganizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDeleteOrganizationalUnitResult struct {
	Result workflow.Future
}

func (r *OrganizationsDeleteOrganizationalUnitResult) Get(ctx workflow.Context) (*organizations.DeleteOrganizationalUnitOutput, error) {
    var output organizations.DeleteOrganizationalUnitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDeletePolicyResult struct {
	Result workflow.Future
}

func (r *OrganizationsDeletePolicyResult) Get(ctx workflow.Context) (*organizations.DeletePolicyOutput, error) {
    var output organizations.DeletePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDeregisterDelegatedAdministratorResult struct {
	Result workflow.Future
}

func (r *OrganizationsDeregisterDelegatedAdministratorResult) Get(ctx workflow.Context) (*organizations.DeregisterDelegatedAdministratorOutput, error) {
    var output organizations.DeregisterDelegatedAdministratorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDescribeAccountResult struct {
	Result workflow.Future
}

func (r *OrganizationsDescribeAccountResult) Get(ctx workflow.Context) (*organizations.DescribeAccountOutput, error) {
    var output organizations.DescribeAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDescribeCreateAccountStatusResult struct {
	Result workflow.Future
}

func (r *OrganizationsDescribeCreateAccountStatusResult) Get(ctx workflow.Context) (*organizations.DescribeCreateAccountStatusOutput, error) {
    var output organizations.DescribeCreateAccountStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDescribeEffectivePolicyResult struct {
	Result workflow.Future
}

func (r *OrganizationsDescribeEffectivePolicyResult) Get(ctx workflow.Context) (*organizations.DescribeEffectivePolicyOutput, error) {
    var output organizations.DescribeEffectivePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDescribeHandshakeResult struct {
	Result workflow.Future
}

func (r *OrganizationsDescribeHandshakeResult) Get(ctx workflow.Context) (*organizations.DescribeHandshakeOutput, error) {
    var output organizations.DescribeHandshakeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDescribeOrganizationResult struct {
	Result workflow.Future
}

func (r *OrganizationsDescribeOrganizationResult) Get(ctx workflow.Context) (*organizations.DescribeOrganizationOutput, error) {
    var output organizations.DescribeOrganizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDescribeOrganizationalUnitResult struct {
	Result workflow.Future
}

func (r *OrganizationsDescribeOrganizationalUnitResult) Get(ctx workflow.Context) (*organizations.DescribeOrganizationalUnitOutput, error) {
    var output organizations.DescribeOrganizationalUnitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDescribePolicyResult struct {
	Result workflow.Future
}

func (r *OrganizationsDescribePolicyResult) Get(ctx workflow.Context) (*organizations.DescribePolicyOutput, error) {
    var output organizations.DescribePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDetachPolicyResult struct {
	Result workflow.Future
}

func (r *OrganizationsDetachPolicyResult) Get(ctx workflow.Context) (*organizations.DetachPolicyOutput, error) {
    var output organizations.DetachPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDisableAWSServiceAccessResult struct {
	Result workflow.Future
}

func (r *OrganizationsDisableAWSServiceAccessResult) Get(ctx workflow.Context) (*organizations.DisableAWSServiceAccessOutput, error) {
    var output organizations.DisableAWSServiceAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsDisablePolicyTypeResult struct {
	Result workflow.Future
}

func (r *OrganizationsDisablePolicyTypeResult) Get(ctx workflow.Context) (*organizations.DisablePolicyTypeOutput, error) {
    var output organizations.DisablePolicyTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsEnableAWSServiceAccessResult struct {
	Result workflow.Future
}

func (r *OrganizationsEnableAWSServiceAccessResult) Get(ctx workflow.Context) (*organizations.EnableAWSServiceAccessOutput, error) {
    var output organizations.EnableAWSServiceAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsEnableAllFeaturesResult struct {
	Result workflow.Future
}

func (r *OrganizationsEnableAllFeaturesResult) Get(ctx workflow.Context) (*organizations.EnableAllFeaturesOutput, error) {
    var output organizations.EnableAllFeaturesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsEnablePolicyTypeResult struct {
	Result workflow.Future
}

func (r *OrganizationsEnablePolicyTypeResult) Get(ctx workflow.Context) (*organizations.EnablePolicyTypeOutput, error) {
    var output organizations.EnablePolicyTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsInviteAccountToOrganizationResult struct {
	Result workflow.Future
}

func (r *OrganizationsInviteAccountToOrganizationResult) Get(ctx workflow.Context) (*organizations.InviteAccountToOrganizationOutput, error) {
    var output organizations.InviteAccountToOrganizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsLeaveOrganizationResult struct {
	Result workflow.Future
}

func (r *OrganizationsLeaveOrganizationResult) Get(ctx workflow.Context) (*organizations.LeaveOrganizationOutput, error) {
    var output organizations.LeaveOrganizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListAWSServiceAccessForOrganizationResult struct {
	Result workflow.Future
}

func (r *OrganizationsListAWSServiceAccessForOrganizationResult) Get(ctx workflow.Context) (*organizations.ListAWSServiceAccessForOrganizationOutput, error) {
    var output organizations.ListAWSServiceAccessForOrganizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListAccountsResult struct {
	Result workflow.Future
}

func (r *OrganizationsListAccountsResult) Get(ctx workflow.Context) (*organizations.ListAccountsOutput, error) {
    var output organizations.ListAccountsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListAccountsForParentResult struct {
	Result workflow.Future
}

func (r *OrganizationsListAccountsForParentResult) Get(ctx workflow.Context) (*organizations.ListAccountsForParentOutput, error) {
    var output organizations.ListAccountsForParentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListChildrenResult struct {
	Result workflow.Future
}

func (r *OrganizationsListChildrenResult) Get(ctx workflow.Context) (*organizations.ListChildrenOutput, error) {
    var output organizations.ListChildrenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListCreateAccountStatusResult struct {
	Result workflow.Future
}

func (r *OrganizationsListCreateAccountStatusResult) Get(ctx workflow.Context) (*organizations.ListCreateAccountStatusOutput, error) {
    var output organizations.ListCreateAccountStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListDelegatedAdministratorsResult struct {
	Result workflow.Future
}

func (r *OrganizationsListDelegatedAdministratorsResult) Get(ctx workflow.Context) (*organizations.ListDelegatedAdministratorsOutput, error) {
    var output organizations.ListDelegatedAdministratorsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListDelegatedServicesForAccountResult struct {
	Result workflow.Future
}

func (r *OrganizationsListDelegatedServicesForAccountResult) Get(ctx workflow.Context) (*organizations.ListDelegatedServicesForAccountOutput, error) {
    var output organizations.ListDelegatedServicesForAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListHandshakesForAccountResult struct {
	Result workflow.Future
}

func (r *OrganizationsListHandshakesForAccountResult) Get(ctx workflow.Context) (*organizations.ListHandshakesForAccountOutput, error) {
    var output organizations.ListHandshakesForAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListHandshakesForOrganizationResult struct {
	Result workflow.Future
}

func (r *OrganizationsListHandshakesForOrganizationResult) Get(ctx workflow.Context) (*organizations.ListHandshakesForOrganizationOutput, error) {
    var output organizations.ListHandshakesForOrganizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListOrganizationalUnitsForParentResult struct {
	Result workflow.Future
}

func (r *OrganizationsListOrganizationalUnitsForParentResult) Get(ctx workflow.Context) (*organizations.ListOrganizationalUnitsForParentOutput, error) {
    var output organizations.ListOrganizationalUnitsForParentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListParentsResult struct {
	Result workflow.Future
}

func (r *OrganizationsListParentsResult) Get(ctx workflow.Context) (*organizations.ListParentsOutput, error) {
    var output organizations.ListParentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListPoliciesResult struct {
	Result workflow.Future
}

func (r *OrganizationsListPoliciesResult) Get(ctx workflow.Context) (*organizations.ListPoliciesOutput, error) {
    var output organizations.ListPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListPoliciesForTargetResult struct {
	Result workflow.Future
}

func (r *OrganizationsListPoliciesForTargetResult) Get(ctx workflow.Context) (*organizations.ListPoliciesForTargetOutput, error) {
    var output organizations.ListPoliciesForTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListRootsResult struct {
	Result workflow.Future
}

func (r *OrganizationsListRootsResult) Get(ctx workflow.Context) (*organizations.ListRootsOutput, error) {
    var output organizations.ListRootsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *OrganizationsListTagsForResourceResult) Get(ctx workflow.Context) (*organizations.ListTagsForResourceOutput, error) {
    var output organizations.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsListTargetsForPolicyResult struct {
	Result workflow.Future
}

func (r *OrganizationsListTargetsForPolicyResult) Get(ctx workflow.Context) (*organizations.ListTargetsForPolicyOutput, error) {
    var output organizations.ListTargetsForPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsMoveAccountResult struct {
	Result workflow.Future
}

func (r *OrganizationsMoveAccountResult) Get(ctx workflow.Context) (*organizations.MoveAccountOutput, error) {
    var output organizations.MoveAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsRegisterDelegatedAdministratorResult struct {
	Result workflow.Future
}

func (r *OrganizationsRegisterDelegatedAdministratorResult) Get(ctx workflow.Context) (*organizations.RegisterDelegatedAdministratorOutput, error) {
    var output organizations.RegisterDelegatedAdministratorOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsRemoveAccountFromOrganizationResult struct {
	Result workflow.Future
}

func (r *OrganizationsRemoveAccountFromOrganizationResult) Get(ctx workflow.Context) (*organizations.RemoveAccountFromOrganizationOutput, error) {
    var output organizations.RemoveAccountFromOrganizationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsTagResourceResult struct {
	Result workflow.Future
}

func (r *OrganizationsTagResourceResult) Get(ctx workflow.Context) (*organizations.TagResourceOutput, error) {
    var output organizations.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsUntagResourceResult struct {
	Result workflow.Future
}

func (r *OrganizationsUntagResourceResult) Get(ctx workflow.Context) (*organizations.UntagResourceOutput, error) {
    var output organizations.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsUpdateOrganizationalUnitResult struct {
	Result workflow.Future
}

func (r *OrganizationsUpdateOrganizationalUnitResult) Get(ctx workflow.Context) (*organizations.UpdateOrganizationalUnitOutput, error) {
    var output organizations.UpdateOrganizationalUnitOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsUpdatePolicyResult struct {
	Result workflow.Future
}

func (r *OrganizationsUpdatePolicyResult) Get(ctx workflow.Context) (*organizations.UpdatePolicyOutput, error) {
    var output organizations.UpdatePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type OrganizationsStub struct {
    activities awsactivities.OrganizationsActivities
}

func NewOrganizationsStub() OrganizationsClient {
    return &OrganizationsStub{}
}

func (a *OrganizationsStub) AcceptHandshake(ctx workflow.Context, input *organizations.AcceptHandshakeInput) (*organizations.AcceptHandshakeOutput, error) {
    var output organizations.AcceptHandshakeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptHandshake, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) AcceptHandshakeAsync(ctx workflow.Context, input *organizations.AcceptHandshakeInput) *OrganizationsAcceptHandshakeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptHandshake, input)
    return &OrganizationsAcceptHandshakeResult{Result: future}
}

func (a *OrganizationsStub) AttachPolicy(ctx workflow.Context, input *organizations.AttachPolicyInput) (*organizations.AttachPolicyOutput, error) {
    var output organizations.AttachPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AttachPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) AttachPolicyAsync(ctx workflow.Context, input *organizations.AttachPolicyInput) *OrganizationsAttachPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AttachPolicy, input)
    return &OrganizationsAttachPolicyResult{Result: future}
}

func (a *OrganizationsStub) CancelHandshake(ctx workflow.Context, input *organizations.CancelHandshakeInput) (*organizations.CancelHandshakeOutput, error) {
    var output organizations.CancelHandshakeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelHandshake, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) CancelHandshakeAsync(ctx workflow.Context, input *organizations.CancelHandshakeInput) *OrganizationsCancelHandshakeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelHandshake, input)
    return &OrganizationsCancelHandshakeResult{Result: future}
}

func (a *OrganizationsStub) CreateAccount(ctx workflow.Context, input *organizations.CreateAccountInput) (*organizations.CreateAccountOutput, error) {
    var output organizations.CreateAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) CreateAccountAsync(ctx workflow.Context, input *organizations.CreateAccountInput) *OrganizationsCreateAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAccount, input)
    return &OrganizationsCreateAccountResult{Result: future}
}

func (a *OrganizationsStub) CreateGovCloudAccount(ctx workflow.Context, input *organizations.CreateGovCloudAccountInput) (*organizations.CreateGovCloudAccountOutput, error) {
    var output organizations.CreateGovCloudAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGovCloudAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) CreateGovCloudAccountAsync(ctx workflow.Context, input *organizations.CreateGovCloudAccountInput) *OrganizationsCreateGovCloudAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateGovCloudAccount, input)
    return &OrganizationsCreateGovCloudAccountResult{Result: future}
}

func (a *OrganizationsStub) CreateOrganization(ctx workflow.Context, input *organizations.CreateOrganizationInput) (*organizations.CreateOrganizationOutput, error) {
    var output organizations.CreateOrganizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateOrganization, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) CreateOrganizationAsync(ctx workflow.Context, input *organizations.CreateOrganizationInput) *OrganizationsCreateOrganizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateOrganization, input)
    return &OrganizationsCreateOrganizationResult{Result: future}
}

func (a *OrganizationsStub) CreateOrganizationalUnit(ctx workflow.Context, input *organizations.CreateOrganizationalUnitInput) (*organizations.CreateOrganizationalUnitOutput, error) {
    var output organizations.CreateOrganizationalUnitOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateOrganizationalUnit, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) CreateOrganizationalUnitAsync(ctx workflow.Context, input *organizations.CreateOrganizationalUnitInput) *OrganizationsCreateOrganizationalUnitResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateOrganizationalUnit, input)
    return &OrganizationsCreateOrganizationalUnitResult{Result: future}
}

func (a *OrganizationsStub) CreatePolicy(ctx workflow.Context, input *organizations.CreatePolicyInput) (*organizations.CreatePolicyOutput, error) {
    var output organizations.CreatePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) CreatePolicyAsync(ctx workflow.Context, input *organizations.CreatePolicyInput) *OrganizationsCreatePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePolicy, input)
    return &OrganizationsCreatePolicyResult{Result: future}
}

func (a *OrganizationsStub) DeclineHandshake(ctx workflow.Context, input *organizations.DeclineHandshakeInput) (*organizations.DeclineHandshakeOutput, error) {
    var output organizations.DeclineHandshakeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeclineHandshake, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DeclineHandshakeAsync(ctx workflow.Context, input *organizations.DeclineHandshakeInput) *OrganizationsDeclineHandshakeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeclineHandshake, input)
    return &OrganizationsDeclineHandshakeResult{Result: future}
}

func (a *OrganizationsStub) DeleteOrganization(ctx workflow.Context, input *organizations.DeleteOrganizationInput) (*organizations.DeleteOrganizationOutput, error) {
    var output organizations.DeleteOrganizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteOrganization, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DeleteOrganizationAsync(ctx workflow.Context, input *organizations.DeleteOrganizationInput) *OrganizationsDeleteOrganizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteOrganization, input)
    return &OrganizationsDeleteOrganizationResult{Result: future}
}

func (a *OrganizationsStub) DeleteOrganizationalUnit(ctx workflow.Context, input *organizations.DeleteOrganizationalUnitInput) (*organizations.DeleteOrganizationalUnitOutput, error) {
    var output organizations.DeleteOrganizationalUnitOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteOrganizationalUnit, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DeleteOrganizationalUnitAsync(ctx workflow.Context, input *organizations.DeleteOrganizationalUnitInput) *OrganizationsDeleteOrganizationalUnitResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteOrganizationalUnit, input)
    return &OrganizationsDeleteOrganizationalUnitResult{Result: future}
}

func (a *OrganizationsStub) DeletePolicy(ctx workflow.Context, input *organizations.DeletePolicyInput) (*organizations.DeletePolicyOutput, error) {
    var output organizations.DeletePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DeletePolicyAsync(ctx workflow.Context, input *organizations.DeletePolicyInput) *OrganizationsDeletePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePolicy, input)
    return &OrganizationsDeletePolicyResult{Result: future}
}

func (a *OrganizationsStub) DeregisterDelegatedAdministrator(ctx workflow.Context, input *organizations.DeregisterDelegatedAdministratorInput) (*organizations.DeregisterDelegatedAdministratorOutput, error) {
    var output organizations.DeregisterDelegatedAdministratorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterDelegatedAdministrator, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DeregisterDelegatedAdministratorAsync(ctx workflow.Context, input *organizations.DeregisterDelegatedAdministratorInput) *OrganizationsDeregisterDelegatedAdministratorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterDelegatedAdministrator, input)
    return &OrganizationsDeregisterDelegatedAdministratorResult{Result: future}
}

func (a *OrganizationsStub) DescribeAccount(ctx workflow.Context, input *organizations.DescribeAccountInput) (*organizations.DescribeAccountOutput, error) {
    var output organizations.DescribeAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DescribeAccountAsync(ctx workflow.Context, input *organizations.DescribeAccountInput) *OrganizationsDescribeAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAccount, input)
    return &OrganizationsDescribeAccountResult{Result: future}
}

func (a *OrganizationsStub) DescribeCreateAccountStatus(ctx workflow.Context, input *organizations.DescribeCreateAccountStatusInput) (*organizations.DescribeCreateAccountStatusOutput, error) {
    var output organizations.DescribeCreateAccountStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCreateAccountStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DescribeCreateAccountStatusAsync(ctx workflow.Context, input *organizations.DescribeCreateAccountStatusInput) *OrganizationsDescribeCreateAccountStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCreateAccountStatus, input)
    return &OrganizationsDescribeCreateAccountStatusResult{Result: future}
}

func (a *OrganizationsStub) DescribeEffectivePolicy(ctx workflow.Context, input *organizations.DescribeEffectivePolicyInput) (*organizations.DescribeEffectivePolicyOutput, error) {
    var output organizations.DescribeEffectivePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEffectivePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DescribeEffectivePolicyAsync(ctx workflow.Context, input *organizations.DescribeEffectivePolicyInput) *OrganizationsDescribeEffectivePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEffectivePolicy, input)
    return &OrganizationsDescribeEffectivePolicyResult{Result: future}
}

func (a *OrganizationsStub) DescribeHandshake(ctx workflow.Context, input *organizations.DescribeHandshakeInput) (*organizations.DescribeHandshakeOutput, error) {
    var output organizations.DescribeHandshakeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHandshake, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DescribeHandshakeAsync(ctx workflow.Context, input *organizations.DescribeHandshakeInput) *OrganizationsDescribeHandshakeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeHandshake, input)
    return &OrganizationsDescribeHandshakeResult{Result: future}
}

func (a *OrganizationsStub) DescribeOrganization(ctx workflow.Context, input *organizations.DescribeOrganizationInput) (*organizations.DescribeOrganizationOutput, error) {
    var output organizations.DescribeOrganizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganization, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DescribeOrganizationAsync(ctx workflow.Context, input *organizations.DescribeOrganizationInput) *OrganizationsDescribeOrganizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganization, input)
    return &OrganizationsDescribeOrganizationResult{Result: future}
}

func (a *OrganizationsStub) DescribeOrganizationalUnit(ctx workflow.Context, input *organizations.DescribeOrganizationalUnitInput) (*organizations.DescribeOrganizationalUnitOutput, error) {
    var output organizations.DescribeOrganizationalUnitOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationalUnit, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DescribeOrganizationalUnitAsync(ctx workflow.Context, input *organizations.DescribeOrganizationalUnitInput) *OrganizationsDescribeOrganizationalUnitResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOrganizationalUnit, input)
    return &OrganizationsDescribeOrganizationalUnitResult{Result: future}
}

func (a *OrganizationsStub) DescribePolicy(ctx workflow.Context, input *organizations.DescribePolicyInput) (*organizations.DescribePolicyOutput, error) {
    var output organizations.DescribePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DescribePolicyAsync(ctx workflow.Context, input *organizations.DescribePolicyInput) *OrganizationsDescribePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePolicy, input)
    return &OrganizationsDescribePolicyResult{Result: future}
}

func (a *OrganizationsStub) DetachPolicy(ctx workflow.Context, input *organizations.DetachPolicyInput) (*organizations.DetachPolicyOutput, error) {
    var output organizations.DetachPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DetachPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DetachPolicyAsync(ctx workflow.Context, input *organizations.DetachPolicyInput) *OrganizationsDetachPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DetachPolicy, input)
    return &OrganizationsDetachPolicyResult{Result: future}
}

func (a *OrganizationsStub) DisableAWSServiceAccess(ctx workflow.Context, input *organizations.DisableAWSServiceAccessInput) (*organizations.DisableAWSServiceAccessOutput, error) {
    var output organizations.DisableAWSServiceAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableAWSServiceAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DisableAWSServiceAccessAsync(ctx workflow.Context, input *organizations.DisableAWSServiceAccessInput) *OrganizationsDisableAWSServiceAccessResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableAWSServiceAccess, input)
    return &OrganizationsDisableAWSServiceAccessResult{Result: future}
}

func (a *OrganizationsStub) DisablePolicyType(ctx workflow.Context, input *organizations.DisablePolicyTypeInput) (*organizations.DisablePolicyTypeOutput, error) {
    var output organizations.DisablePolicyTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisablePolicyType, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) DisablePolicyTypeAsync(ctx workflow.Context, input *organizations.DisablePolicyTypeInput) *OrganizationsDisablePolicyTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisablePolicyType, input)
    return &OrganizationsDisablePolicyTypeResult{Result: future}
}

func (a *OrganizationsStub) EnableAWSServiceAccess(ctx workflow.Context, input *organizations.EnableAWSServiceAccessInput) (*organizations.EnableAWSServiceAccessOutput, error) {
    var output organizations.EnableAWSServiceAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableAWSServiceAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) EnableAWSServiceAccessAsync(ctx workflow.Context, input *organizations.EnableAWSServiceAccessInput) *OrganizationsEnableAWSServiceAccessResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableAWSServiceAccess, input)
    return &OrganizationsEnableAWSServiceAccessResult{Result: future}
}

func (a *OrganizationsStub) EnableAllFeatures(ctx workflow.Context, input *organizations.EnableAllFeaturesInput) (*organizations.EnableAllFeaturesOutput, error) {
    var output organizations.EnableAllFeaturesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableAllFeatures, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) EnableAllFeaturesAsync(ctx workflow.Context, input *organizations.EnableAllFeaturesInput) *OrganizationsEnableAllFeaturesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableAllFeatures, input)
    return &OrganizationsEnableAllFeaturesResult{Result: future}
}

func (a *OrganizationsStub) EnablePolicyType(ctx workflow.Context, input *organizations.EnablePolicyTypeInput) (*organizations.EnablePolicyTypeOutput, error) {
    var output organizations.EnablePolicyTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnablePolicyType, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) EnablePolicyTypeAsync(ctx workflow.Context, input *organizations.EnablePolicyTypeInput) *OrganizationsEnablePolicyTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnablePolicyType, input)
    return &OrganizationsEnablePolicyTypeResult{Result: future}
}

func (a *OrganizationsStub) InviteAccountToOrganization(ctx workflow.Context, input *organizations.InviteAccountToOrganizationInput) (*organizations.InviteAccountToOrganizationOutput, error) {
    var output organizations.InviteAccountToOrganizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.InviteAccountToOrganization, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) InviteAccountToOrganizationAsync(ctx workflow.Context, input *organizations.InviteAccountToOrganizationInput) *OrganizationsInviteAccountToOrganizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.InviteAccountToOrganization, input)
    return &OrganizationsInviteAccountToOrganizationResult{Result: future}
}

func (a *OrganizationsStub) LeaveOrganization(ctx workflow.Context, input *organizations.LeaveOrganizationInput) (*organizations.LeaveOrganizationOutput, error) {
    var output organizations.LeaveOrganizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.LeaveOrganization, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) LeaveOrganizationAsync(ctx workflow.Context, input *organizations.LeaveOrganizationInput) *OrganizationsLeaveOrganizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.LeaveOrganization, input)
    return &OrganizationsLeaveOrganizationResult{Result: future}
}

func (a *OrganizationsStub) ListAWSServiceAccessForOrganization(ctx workflow.Context, input *organizations.ListAWSServiceAccessForOrganizationInput) (*organizations.ListAWSServiceAccessForOrganizationOutput, error) {
    var output organizations.ListAWSServiceAccessForOrganizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAWSServiceAccessForOrganization, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListAWSServiceAccessForOrganizationAsync(ctx workflow.Context, input *organizations.ListAWSServiceAccessForOrganizationInput) *OrganizationsListAWSServiceAccessForOrganizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAWSServiceAccessForOrganization, input)
    return &OrganizationsListAWSServiceAccessForOrganizationResult{Result: future}
}

func (a *OrganizationsStub) ListAccounts(ctx workflow.Context, input *organizations.ListAccountsInput) (*organizations.ListAccountsOutput, error) {
    var output organizations.ListAccountsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAccounts, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListAccountsAsync(ctx workflow.Context, input *organizations.ListAccountsInput) *OrganizationsListAccountsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAccounts, input)
    return &OrganizationsListAccountsResult{Result: future}
}

func (a *OrganizationsStub) ListAccountsForParent(ctx workflow.Context, input *organizations.ListAccountsForParentInput) (*organizations.ListAccountsForParentOutput, error) {
    var output organizations.ListAccountsForParentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAccountsForParent, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListAccountsForParentAsync(ctx workflow.Context, input *organizations.ListAccountsForParentInput) *OrganizationsListAccountsForParentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAccountsForParent, input)
    return &OrganizationsListAccountsForParentResult{Result: future}
}

func (a *OrganizationsStub) ListChildren(ctx workflow.Context, input *organizations.ListChildrenInput) (*organizations.ListChildrenOutput, error) {
    var output organizations.ListChildrenOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListChildren, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListChildrenAsync(ctx workflow.Context, input *organizations.ListChildrenInput) *OrganizationsListChildrenResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListChildren, input)
    return &OrganizationsListChildrenResult{Result: future}
}

func (a *OrganizationsStub) ListCreateAccountStatus(ctx workflow.Context, input *organizations.ListCreateAccountStatusInput) (*organizations.ListCreateAccountStatusOutput, error) {
    var output organizations.ListCreateAccountStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCreateAccountStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListCreateAccountStatusAsync(ctx workflow.Context, input *organizations.ListCreateAccountStatusInput) *OrganizationsListCreateAccountStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListCreateAccountStatus, input)
    return &OrganizationsListCreateAccountStatusResult{Result: future}
}

func (a *OrganizationsStub) ListDelegatedAdministrators(ctx workflow.Context, input *organizations.ListDelegatedAdministratorsInput) (*organizations.ListDelegatedAdministratorsOutput, error) {
    var output organizations.ListDelegatedAdministratorsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDelegatedAdministrators, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListDelegatedAdministratorsAsync(ctx workflow.Context, input *organizations.ListDelegatedAdministratorsInput) *OrganizationsListDelegatedAdministratorsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDelegatedAdministrators, input)
    return &OrganizationsListDelegatedAdministratorsResult{Result: future}
}

func (a *OrganizationsStub) ListDelegatedServicesForAccount(ctx workflow.Context, input *organizations.ListDelegatedServicesForAccountInput) (*organizations.ListDelegatedServicesForAccountOutput, error) {
    var output organizations.ListDelegatedServicesForAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDelegatedServicesForAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListDelegatedServicesForAccountAsync(ctx workflow.Context, input *organizations.ListDelegatedServicesForAccountInput) *OrganizationsListDelegatedServicesForAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDelegatedServicesForAccount, input)
    return &OrganizationsListDelegatedServicesForAccountResult{Result: future}
}

func (a *OrganizationsStub) ListHandshakesForAccount(ctx workflow.Context, input *organizations.ListHandshakesForAccountInput) (*organizations.ListHandshakesForAccountOutput, error) {
    var output organizations.ListHandshakesForAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHandshakesForAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListHandshakesForAccountAsync(ctx workflow.Context, input *organizations.ListHandshakesForAccountInput) *OrganizationsListHandshakesForAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHandshakesForAccount, input)
    return &OrganizationsListHandshakesForAccountResult{Result: future}
}

func (a *OrganizationsStub) ListHandshakesForOrganization(ctx workflow.Context, input *organizations.ListHandshakesForOrganizationInput) (*organizations.ListHandshakesForOrganizationOutput, error) {
    var output organizations.ListHandshakesForOrganizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHandshakesForOrganization, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListHandshakesForOrganizationAsync(ctx workflow.Context, input *organizations.ListHandshakesForOrganizationInput) *OrganizationsListHandshakesForOrganizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHandshakesForOrganization, input)
    return &OrganizationsListHandshakesForOrganizationResult{Result: future}
}

func (a *OrganizationsStub) ListOrganizationalUnitsForParent(ctx workflow.Context, input *organizations.ListOrganizationalUnitsForParentInput) (*organizations.ListOrganizationalUnitsForParentOutput, error) {
    var output organizations.ListOrganizationalUnitsForParentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOrganizationalUnitsForParent, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListOrganizationalUnitsForParentAsync(ctx workflow.Context, input *organizations.ListOrganizationalUnitsForParentInput) *OrganizationsListOrganizationalUnitsForParentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOrganizationalUnitsForParent, input)
    return &OrganizationsListOrganizationalUnitsForParentResult{Result: future}
}

func (a *OrganizationsStub) ListParents(ctx workflow.Context, input *organizations.ListParentsInput) (*organizations.ListParentsOutput, error) {
    var output organizations.ListParentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListParents, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListParentsAsync(ctx workflow.Context, input *organizations.ListParentsInput) *OrganizationsListParentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListParents, input)
    return &OrganizationsListParentsResult{Result: future}
}

func (a *OrganizationsStub) ListPolicies(ctx workflow.Context, input *organizations.ListPoliciesInput) (*organizations.ListPoliciesOutput, error) {
    var output organizations.ListPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListPoliciesAsync(ctx workflow.Context, input *organizations.ListPoliciesInput) *OrganizationsListPoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPolicies, input)
    return &OrganizationsListPoliciesResult{Result: future}
}

func (a *OrganizationsStub) ListPoliciesForTarget(ctx workflow.Context, input *organizations.ListPoliciesForTargetInput) (*organizations.ListPoliciesForTargetOutput, error) {
    var output organizations.ListPoliciesForTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPoliciesForTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListPoliciesForTargetAsync(ctx workflow.Context, input *organizations.ListPoliciesForTargetInput) *OrganizationsListPoliciesForTargetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPoliciesForTarget, input)
    return &OrganizationsListPoliciesForTargetResult{Result: future}
}

func (a *OrganizationsStub) ListRoots(ctx workflow.Context, input *organizations.ListRootsInput) (*organizations.ListRootsOutput, error) {
    var output organizations.ListRootsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRoots, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListRootsAsync(ctx workflow.Context, input *organizations.ListRootsInput) *OrganizationsListRootsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRoots, input)
    return &OrganizationsListRootsResult{Result: future}
}

func (a *OrganizationsStub) ListTagsForResource(ctx workflow.Context, input *organizations.ListTagsForResourceInput) (*organizations.ListTagsForResourceOutput, error) {
    var output organizations.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListTagsForResourceAsync(ctx workflow.Context, input *organizations.ListTagsForResourceInput) *OrganizationsListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &OrganizationsListTagsForResourceResult{Result: future}
}

func (a *OrganizationsStub) ListTargetsForPolicy(ctx workflow.Context, input *organizations.ListTargetsForPolicyInput) (*organizations.ListTargetsForPolicyOutput, error) {
    var output organizations.ListTargetsForPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTargetsForPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) ListTargetsForPolicyAsync(ctx workflow.Context, input *organizations.ListTargetsForPolicyInput) *OrganizationsListTargetsForPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTargetsForPolicy, input)
    return &OrganizationsListTargetsForPolicyResult{Result: future}
}

func (a *OrganizationsStub) MoveAccount(ctx workflow.Context, input *organizations.MoveAccountInput) (*organizations.MoveAccountOutput, error) {
    var output organizations.MoveAccountOutput
    err := workflow.ExecuteActivity(ctx, a.activities.MoveAccount, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) MoveAccountAsync(ctx workflow.Context, input *organizations.MoveAccountInput) *OrganizationsMoveAccountResult {
    future := workflow.ExecuteActivity(ctx, a.activities.MoveAccount, input)
    return &OrganizationsMoveAccountResult{Result: future}
}

func (a *OrganizationsStub) RegisterDelegatedAdministrator(ctx workflow.Context, input *organizations.RegisterDelegatedAdministratorInput) (*organizations.RegisterDelegatedAdministratorOutput, error) {
    var output organizations.RegisterDelegatedAdministratorOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterDelegatedAdministrator, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) RegisterDelegatedAdministratorAsync(ctx workflow.Context, input *organizations.RegisterDelegatedAdministratorInput) *OrganizationsRegisterDelegatedAdministratorResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterDelegatedAdministrator, input)
    return &OrganizationsRegisterDelegatedAdministratorResult{Result: future}
}

func (a *OrganizationsStub) RemoveAccountFromOrganization(ctx workflow.Context, input *organizations.RemoveAccountFromOrganizationInput) (*organizations.RemoveAccountFromOrganizationOutput, error) {
    var output organizations.RemoveAccountFromOrganizationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveAccountFromOrganization, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) RemoveAccountFromOrganizationAsync(ctx workflow.Context, input *organizations.RemoveAccountFromOrganizationInput) *OrganizationsRemoveAccountFromOrganizationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveAccountFromOrganization, input)
    return &OrganizationsRemoveAccountFromOrganizationResult{Result: future}
}

func (a *OrganizationsStub) TagResource(ctx workflow.Context, input *organizations.TagResourceInput) (*organizations.TagResourceOutput, error) {
    var output organizations.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) TagResourceAsync(ctx workflow.Context, input *organizations.TagResourceInput) *OrganizationsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &OrganizationsTagResourceResult{Result: future}
}

func (a *OrganizationsStub) UntagResource(ctx workflow.Context, input *organizations.UntagResourceInput) (*organizations.UntagResourceOutput, error) {
    var output organizations.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) UntagResourceAsync(ctx workflow.Context, input *organizations.UntagResourceInput) *OrganizationsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &OrganizationsUntagResourceResult{Result: future}
}

func (a *OrganizationsStub) UpdateOrganizationalUnit(ctx workflow.Context, input *organizations.UpdateOrganizationalUnitInput) (*organizations.UpdateOrganizationalUnitOutput, error) {
    var output organizations.UpdateOrganizationalUnitOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateOrganizationalUnit, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) UpdateOrganizationalUnitAsync(ctx workflow.Context, input *organizations.UpdateOrganizationalUnitInput) *OrganizationsUpdateOrganizationalUnitResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateOrganizationalUnit, input)
    return &OrganizationsUpdateOrganizationalUnitResult{Result: future}
}

func (a *OrganizationsStub) UpdatePolicy(ctx workflow.Context, input *organizations.UpdatePolicyInput) (*organizations.UpdatePolicyOutput, error) {
    var output organizations.UpdatePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *OrganizationsStub) UpdatePolicyAsync(ctx workflow.Context, input *organizations.UpdatePolicyInput) *OrganizationsUpdatePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdatePolicy, input)
    return &OrganizationsUpdatePolicyResult{Result: future}
}
