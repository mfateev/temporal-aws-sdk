package awsclients

import (
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	"go.temporal.io/sdk/workflow"
)

type SSOAdminClient interface {
       AttachManagedPolicyToPermissionSet(ctx workflow.Context, input *ssoadmin.AttachManagedPolicyToPermissionSetInput) (*ssoadmin.AttachManagedPolicyToPermissionSetOutput, error)
       AttachManagedPolicyToPermissionSetAsync(ctx workflow.Context, input *ssoadmin.AttachManagedPolicyToPermissionSetInput) *SsoadminAttachManagedPolicyToPermissionSetResult

       CreateAccountAssignment(ctx workflow.Context, input *ssoadmin.CreateAccountAssignmentInput) (*ssoadmin.CreateAccountAssignmentOutput, error)
       CreateAccountAssignmentAsync(ctx workflow.Context, input *ssoadmin.CreateAccountAssignmentInput) *SsoadminCreateAccountAssignmentResult

       CreatePermissionSet(ctx workflow.Context, input *ssoadmin.CreatePermissionSetInput) (*ssoadmin.CreatePermissionSetOutput, error)
       CreatePermissionSetAsync(ctx workflow.Context, input *ssoadmin.CreatePermissionSetInput) *SsoadminCreatePermissionSetResult

       DeleteAccountAssignment(ctx workflow.Context, input *ssoadmin.DeleteAccountAssignmentInput) (*ssoadmin.DeleteAccountAssignmentOutput, error)
       DeleteAccountAssignmentAsync(ctx workflow.Context, input *ssoadmin.DeleteAccountAssignmentInput) *SsoadminDeleteAccountAssignmentResult

       DeleteInlinePolicyFromPermissionSet(ctx workflow.Context, input *ssoadmin.DeleteInlinePolicyFromPermissionSetInput) (*ssoadmin.DeleteInlinePolicyFromPermissionSetOutput, error)
       DeleteInlinePolicyFromPermissionSetAsync(ctx workflow.Context, input *ssoadmin.DeleteInlinePolicyFromPermissionSetInput) *SsoadminDeleteInlinePolicyFromPermissionSetResult

       DeletePermissionSet(ctx workflow.Context, input *ssoadmin.DeletePermissionSetInput) (*ssoadmin.DeletePermissionSetOutput, error)
       DeletePermissionSetAsync(ctx workflow.Context, input *ssoadmin.DeletePermissionSetInput) *SsoadminDeletePermissionSetResult

       DescribeAccountAssignmentCreationStatus(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentCreationStatusInput) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error)
       DescribeAccountAssignmentCreationStatusAsync(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentCreationStatusInput) *SsoadminDescribeAccountAssignmentCreationStatusResult

       DescribeAccountAssignmentDeletionStatus(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentDeletionStatusInput) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error)
       DescribeAccountAssignmentDeletionStatusAsync(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentDeletionStatusInput) *SsoadminDescribeAccountAssignmentDeletionStatusResult

       DescribePermissionSet(ctx workflow.Context, input *ssoadmin.DescribePermissionSetInput) (*ssoadmin.DescribePermissionSetOutput, error)
       DescribePermissionSetAsync(ctx workflow.Context, input *ssoadmin.DescribePermissionSetInput) *SsoadminDescribePermissionSetResult

       DescribePermissionSetProvisioningStatus(ctx workflow.Context, input *ssoadmin.DescribePermissionSetProvisioningStatusInput) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error)
       DescribePermissionSetProvisioningStatusAsync(ctx workflow.Context, input *ssoadmin.DescribePermissionSetProvisioningStatusInput) *SsoadminDescribePermissionSetProvisioningStatusResult

       DetachManagedPolicyFromPermissionSet(ctx workflow.Context, input *ssoadmin.DetachManagedPolicyFromPermissionSetInput) (*ssoadmin.DetachManagedPolicyFromPermissionSetOutput, error)
       DetachManagedPolicyFromPermissionSetAsync(ctx workflow.Context, input *ssoadmin.DetachManagedPolicyFromPermissionSetInput) *SsoadminDetachManagedPolicyFromPermissionSetResult

       GetInlinePolicyForPermissionSet(ctx workflow.Context, input *ssoadmin.GetInlinePolicyForPermissionSetInput) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error)
       GetInlinePolicyForPermissionSetAsync(ctx workflow.Context, input *ssoadmin.GetInlinePolicyForPermissionSetInput) *SsoadminGetInlinePolicyForPermissionSetResult

       ListAccountAssignmentCreationStatus(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentCreationStatusInput) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error)
       ListAccountAssignmentCreationStatusAsync(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentCreationStatusInput) *SsoadminListAccountAssignmentCreationStatusResult

       ListAccountAssignmentDeletionStatus(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentDeletionStatusInput) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error)
       ListAccountAssignmentDeletionStatusAsync(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentDeletionStatusInput) *SsoadminListAccountAssignmentDeletionStatusResult

       ListAccountAssignments(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentsInput) (*ssoadmin.ListAccountAssignmentsOutput, error)
       ListAccountAssignmentsAsync(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentsInput) *SsoadminListAccountAssignmentsResult

       ListAccountsForProvisionedPermissionSet(ctx workflow.Context, input *ssoadmin.ListAccountsForProvisionedPermissionSetInput) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error)
       ListAccountsForProvisionedPermissionSetAsync(ctx workflow.Context, input *ssoadmin.ListAccountsForProvisionedPermissionSetInput) *SsoadminListAccountsForProvisionedPermissionSetResult

       ListInstances(ctx workflow.Context, input *ssoadmin.ListInstancesInput) (*ssoadmin.ListInstancesOutput, error)
       ListInstancesAsync(ctx workflow.Context, input *ssoadmin.ListInstancesInput) *SsoadminListInstancesResult

       ListManagedPoliciesInPermissionSet(ctx workflow.Context, input *ssoadmin.ListManagedPoliciesInPermissionSetInput) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error)
       ListManagedPoliciesInPermissionSetAsync(ctx workflow.Context, input *ssoadmin.ListManagedPoliciesInPermissionSetInput) *SsoadminListManagedPoliciesInPermissionSetResult

       ListPermissionSetProvisioningStatus(ctx workflow.Context, input *ssoadmin.ListPermissionSetProvisioningStatusInput) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error)
       ListPermissionSetProvisioningStatusAsync(ctx workflow.Context, input *ssoadmin.ListPermissionSetProvisioningStatusInput) *SsoadminListPermissionSetProvisioningStatusResult

       ListPermissionSets(ctx workflow.Context, input *ssoadmin.ListPermissionSetsInput) (*ssoadmin.ListPermissionSetsOutput, error)
       ListPermissionSetsAsync(ctx workflow.Context, input *ssoadmin.ListPermissionSetsInput) *SsoadminListPermissionSetsResult

       ListPermissionSetsProvisionedToAccount(ctx workflow.Context, input *ssoadmin.ListPermissionSetsProvisionedToAccountInput) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error)
       ListPermissionSetsProvisionedToAccountAsync(ctx workflow.Context, input *ssoadmin.ListPermissionSetsProvisionedToAccountInput) *SsoadminListPermissionSetsProvisionedToAccountResult

       ListTagsForResource(ctx workflow.Context, input *ssoadmin.ListTagsForResourceInput) (*ssoadmin.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *ssoadmin.ListTagsForResourceInput) *SsoadminListTagsForResourceResult

       ProvisionPermissionSet(ctx workflow.Context, input *ssoadmin.ProvisionPermissionSetInput) (*ssoadmin.ProvisionPermissionSetOutput, error)
       ProvisionPermissionSetAsync(ctx workflow.Context, input *ssoadmin.ProvisionPermissionSetInput) *SsoadminProvisionPermissionSetResult

       PutInlinePolicyToPermissionSet(ctx workflow.Context, input *ssoadmin.PutInlinePolicyToPermissionSetInput) (*ssoadmin.PutInlinePolicyToPermissionSetOutput, error)
       PutInlinePolicyToPermissionSetAsync(ctx workflow.Context, input *ssoadmin.PutInlinePolicyToPermissionSetInput) *SsoadminPutInlinePolicyToPermissionSetResult

       TagResource(ctx workflow.Context, input *ssoadmin.TagResourceInput) (*ssoadmin.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *ssoadmin.TagResourceInput) *SsoadminTagResourceResult

       UntagResource(ctx workflow.Context, input *ssoadmin.UntagResourceInput) (*ssoadmin.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *ssoadmin.UntagResourceInput) *SsoadminUntagResourceResult

       UpdatePermissionSet(ctx workflow.Context, input *ssoadmin.UpdatePermissionSetInput) (*ssoadmin.UpdatePermissionSetOutput, error)
       UpdatePermissionSetAsync(ctx workflow.Context, input *ssoadmin.UpdatePermissionSetInput) *SsoadminUpdatePermissionSetResult
}

type SSOAdminStub struct {
}

func NewSSOAdminStub() SSOAdminClient {
    return &SSOAdminStub{}
}

type SsoadminAttachManagedPolicyToPermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminAttachManagedPolicyToPermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.AttachManagedPolicyToPermissionSetOutput, error) {
    var output ssoadmin.AttachManagedPolicyToPermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminCreateAccountAssignmentResult struct {
	Result workflow.Future
}

func (r *SsoadminCreateAccountAssignmentResult) Get(ctx workflow.Context) (*ssoadmin.CreateAccountAssignmentOutput, error) {
    var output ssoadmin.CreateAccountAssignmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminCreatePermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminCreatePermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.CreatePermissionSetOutput, error) {
    var output ssoadmin.CreatePermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminDeleteAccountAssignmentResult struct {
	Result workflow.Future
}

func (r *SsoadminDeleteAccountAssignmentResult) Get(ctx workflow.Context) (*ssoadmin.DeleteAccountAssignmentOutput, error) {
    var output ssoadmin.DeleteAccountAssignmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminDeleteInlinePolicyFromPermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminDeleteInlinePolicyFromPermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.DeleteInlinePolicyFromPermissionSetOutput, error) {
    var output ssoadmin.DeleteInlinePolicyFromPermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminDeletePermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminDeletePermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.DeletePermissionSetOutput, error) {
    var output ssoadmin.DeletePermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminDescribeAccountAssignmentCreationStatusResult struct {
	Result workflow.Future
}

func (r *SsoadminDescribeAccountAssignmentCreationStatusResult) Get(ctx workflow.Context) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error) {
    var output ssoadmin.DescribeAccountAssignmentCreationStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminDescribeAccountAssignmentDeletionStatusResult struct {
	Result workflow.Future
}

func (r *SsoadminDescribeAccountAssignmentDeletionStatusResult) Get(ctx workflow.Context) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error) {
    var output ssoadmin.DescribeAccountAssignmentDeletionStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminDescribePermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminDescribePermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.DescribePermissionSetOutput, error) {
    var output ssoadmin.DescribePermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminDescribePermissionSetProvisioningStatusResult struct {
	Result workflow.Future
}

func (r *SsoadminDescribePermissionSetProvisioningStatusResult) Get(ctx workflow.Context) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error) {
    var output ssoadmin.DescribePermissionSetProvisioningStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminDetachManagedPolicyFromPermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminDetachManagedPolicyFromPermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.DetachManagedPolicyFromPermissionSetOutput, error) {
    var output ssoadmin.DetachManagedPolicyFromPermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminGetInlinePolicyForPermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminGetInlinePolicyForPermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error) {
    var output ssoadmin.GetInlinePolicyForPermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminListAccountAssignmentCreationStatusResult struct {
	Result workflow.Future
}

func (r *SsoadminListAccountAssignmentCreationStatusResult) Get(ctx workflow.Context) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error) {
    var output ssoadmin.ListAccountAssignmentCreationStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminListAccountAssignmentDeletionStatusResult struct {
	Result workflow.Future
}

func (r *SsoadminListAccountAssignmentDeletionStatusResult) Get(ctx workflow.Context) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error) {
    var output ssoadmin.ListAccountAssignmentDeletionStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminListAccountAssignmentsResult struct {
	Result workflow.Future
}

func (r *SsoadminListAccountAssignmentsResult) Get(ctx workflow.Context) (*ssoadmin.ListAccountAssignmentsOutput, error) {
    var output ssoadmin.ListAccountAssignmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminListAccountsForProvisionedPermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminListAccountsForProvisionedPermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error) {
    var output ssoadmin.ListAccountsForProvisionedPermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminListInstancesResult struct {
	Result workflow.Future
}

func (r *SsoadminListInstancesResult) Get(ctx workflow.Context) (*ssoadmin.ListInstancesOutput, error) {
    var output ssoadmin.ListInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminListManagedPoliciesInPermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminListManagedPoliciesInPermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error) {
    var output ssoadmin.ListManagedPoliciesInPermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminListPermissionSetProvisioningStatusResult struct {
	Result workflow.Future
}

func (r *SsoadminListPermissionSetProvisioningStatusResult) Get(ctx workflow.Context) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error) {
    var output ssoadmin.ListPermissionSetProvisioningStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminListPermissionSetsResult struct {
	Result workflow.Future
}

func (r *SsoadminListPermissionSetsResult) Get(ctx workflow.Context) (*ssoadmin.ListPermissionSetsOutput, error) {
    var output ssoadmin.ListPermissionSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminListPermissionSetsProvisionedToAccountResult struct {
	Result workflow.Future
}

func (r *SsoadminListPermissionSetsProvisionedToAccountResult) Get(ctx workflow.Context) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error) {
    var output ssoadmin.ListPermissionSetsProvisionedToAccountOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *SsoadminListTagsForResourceResult) Get(ctx workflow.Context) (*ssoadmin.ListTagsForResourceOutput, error) {
    var output ssoadmin.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminProvisionPermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminProvisionPermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.ProvisionPermissionSetOutput, error) {
    var output ssoadmin.ProvisionPermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminPutInlinePolicyToPermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminPutInlinePolicyToPermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.PutInlinePolicyToPermissionSetOutput, error) {
    var output ssoadmin.PutInlinePolicyToPermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminTagResourceResult struct {
	Result workflow.Future
}

func (r *SsoadminTagResourceResult) Get(ctx workflow.Context) (*ssoadmin.TagResourceOutput, error) {
    var output ssoadmin.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminUntagResourceResult struct {
	Result workflow.Future
}

func (r *SsoadminUntagResourceResult) Get(ctx workflow.Context) (*ssoadmin.UntagResourceOutput, error) {
    var output ssoadmin.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type SsoadminUpdatePermissionSetResult struct {
	Result workflow.Future
}

func (r *SsoadminUpdatePermissionSetResult) Get(ctx workflow.Context) (*ssoadmin.UpdatePermissionSetOutput, error) {
    var output ssoadmin.UpdatePermissionSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) AttachManagedPolicyToPermissionSet(ctx workflow.Context, input *ssoadmin.AttachManagedPolicyToPermissionSetInput) (*ssoadmin.AttachManagedPolicyToPermissionSetOutput, error) {
    var output ssoadmin.AttachManagedPolicyToPermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.AttachManagedPolicyToPermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) AttachManagedPolicyToPermissionSetAsync(ctx workflow.Context, input *ssoadmin.AttachManagedPolicyToPermissionSetInput) *SsoadminAttachManagedPolicyToPermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.AttachManagedPolicyToPermissionSet", input)
    return &SsoadminAttachManagedPolicyToPermissionSetResult{Result: future}
}

func (a *SSOAdminStub) CreateAccountAssignment(ctx workflow.Context, input *ssoadmin.CreateAccountAssignmentInput) (*ssoadmin.CreateAccountAssignmentOutput, error) {
    var output ssoadmin.CreateAccountAssignmentOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.CreateAccountAssignment", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) CreateAccountAssignmentAsync(ctx workflow.Context, input *ssoadmin.CreateAccountAssignmentInput) *SsoadminCreateAccountAssignmentResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.CreateAccountAssignment", input)
    return &SsoadminCreateAccountAssignmentResult{Result: future}
}

func (a *SSOAdminStub) CreatePermissionSet(ctx workflow.Context, input *ssoadmin.CreatePermissionSetInput) (*ssoadmin.CreatePermissionSetOutput, error) {
    var output ssoadmin.CreatePermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.CreatePermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) CreatePermissionSetAsync(ctx workflow.Context, input *ssoadmin.CreatePermissionSetInput) *SsoadminCreatePermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.CreatePermissionSet", input)
    return &SsoadminCreatePermissionSetResult{Result: future}
}

func (a *SSOAdminStub) DeleteAccountAssignment(ctx workflow.Context, input *ssoadmin.DeleteAccountAssignmentInput) (*ssoadmin.DeleteAccountAssignmentOutput, error) {
    var output ssoadmin.DeleteAccountAssignmentOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.DeleteAccountAssignment", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) DeleteAccountAssignmentAsync(ctx workflow.Context, input *ssoadmin.DeleteAccountAssignmentInput) *SsoadminDeleteAccountAssignmentResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.DeleteAccountAssignment", input)
    return &SsoadminDeleteAccountAssignmentResult{Result: future}
}

func (a *SSOAdminStub) DeleteInlinePolicyFromPermissionSet(ctx workflow.Context, input *ssoadmin.DeleteInlinePolicyFromPermissionSetInput) (*ssoadmin.DeleteInlinePolicyFromPermissionSetOutput, error) {
    var output ssoadmin.DeleteInlinePolicyFromPermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.DeleteInlinePolicyFromPermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) DeleteInlinePolicyFromPermissionSetAsync(ctx workflow.Context, input *ssoadmin.DeleteInlinePolicyFromPermissionSetInput) *SsoadminDeleteInlinePolicyFromPermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.DeleteInlinePolicyFromPermissionSet", input)
    return &SsoadminDeleteInlinePolicyFromPermissionSetResult{Result: future}
}

func (a *SSOAdminStub) DeletePermissionSet(ctx workflow.Context, input *ssoadmin.DeletePermissionSetInput) (*ssoadmin.DeletePermissionSetOutput, error) {
    var output ssoadmin.DeletePermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.DeletePermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) DeletePermissionSetAsync(ctx workflow.Context, input *ssoadmin.DeletePermissionSetInput) *SsoadminDeletePermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.DeletePermissionSet", input)
    return &SsoadminDeletePermissionSetResult{Result: future}
}

func (a *SSOAdminStub) DescribeAccountAssignmentCreationStatus(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentCreationStatusInput) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error) {
    var output ssoadmin.DescribeAccountAssignmentCreationStatusOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.DescribeAccountAssignmentCreationStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) DescribeAccountAssignmentCreationStatusAsync(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentCreationStatusInput) *SsoadminDescribeAccountAssignmentCreationStatusResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.DescribeAccountAssignmentCreationStatus", input)
    return &SsoadminDescribeAccountAssignmentCreationStatusResult{Result: future}
}

func (a *SSOAdminStub) DescribeAccountAssignmentDeletionStatus(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentDeletionStatusInput) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error) {
    var output ssoadmin.DescribeAccountAssignmentDeletionStatusOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.DescribeAccountAssignmentDeletionStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) DescribeAccountAssignmentDeletionStatusAsync(ctx workflow.Context, input *ssoadmin.DescribeAccountAssignmentDeletionStatusInput) *SsoadminDescribeAccountAssignmentDeletionStatusResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.DescribeAccountAssignmentDeletionStatus", input)
    return &SsoadminDescribeAccountAssignmentDeletionStatusResult{Result: future}
}

func (a *SSOAdminStub) DescribePermissionSet(ctx workflow.Context, input *ssoadmin.DescribePermissionSetInput) (*ssoadmin.DescribePermissionSetOutput, error) {
    var output ssoadmin.DescribePermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.DescribePermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) DescribePermissionSetAsync(ctx workflow.Context, input *ssoadmin.DescribePermissionSetInput) *SsoadminDescribePermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.DescribePermissionSet", input)
    return &SsoadminDescribePermissionSetResult{Result: future}
}

func (a *SSOAdminStub) DescribePermissionSetProvisioningStatus(ctx workflow.Context, input *ssoadmin.DescribePermissionSetProvisioningStatusInput) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error) {
    var output ssoadmin.DescribePermissionSetProvisioningStatusOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.DescribePermissionSetProvisioningStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) DescribePermissionSetProvisioningStatusAsync(ctx workflow.Context, input *ssoadmin.DescribePermissionSetProvisioningStatusInput) *SsoadminDescribePermissionSetProvisioningStatusResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.DescribePermissionSetProvisioningStatus", input)
    return &SsoadminDescribePermissionSetProvisioningStatusResult{Result: future}
}

func (a *SSOAdminStub) DetachManagedPolicyFromPermissionSet(ctx workflow.Context, input *ssoadmin.DetachManagedPolicyFromPermissionSetInput) (*ssoadmin.DetachManagedPolicyFromPermissionSetOutput, error) {
    var output ssoadmin.DetachManagedPolicyFromPermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.DetachManagedPolicyFromPermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) DetachManagedPolicyFromPermissionSetAsync(ctx workflow.Context, input *ssoadmin.DetachManagedPolicyFromPermissionSetInput) *SsoadminDetachManagedPolicyFromPermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.DetachManagedPolicyFromPermissionSet", input)
    return &SsoadminDetachManagedPolicyFromPermissionSetResult{Result: future}
}

func (a *SSOAdminStub) GetInlinePolicyForPermissionSet(ctx workflow.Context, input *ssoadmin.GetInlinePolicyForPermissionSetInput) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error) {
    var output ssoadmin.GetInlinePolicyForPermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.GetInlinePolicyForPermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) GetInlinePolicyForPermissionSetAsync(ctx workflow.Context, input *ssoadmin.GetInlinePolicyForPermissionSetInput) *SsoadminGetInlinePolicyForPermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.GetInlinePolicyForPermissionSet", input)
    return &SsoadminGetInlinePolicyForPermissionSetResult{Result: future}
}

func (a *SSOAdminStub) ListAccountAssignmentCreationStatus(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentCreationStatusInput) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error) {
    var output ssoadmin.ListAccountAssignmentCreationStatusOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.ListAccountAssignmentCreationStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) ListAccountAssignmentCreationStatusAsync(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentCreationStatusInput) *SsoadminListAccountAssignmentCreationStatusResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.ListAccountAssignmentCreationStatus", input)
    return &SsoadminListAccountAssignmentCreationStatusResult{Result: future}
}

func (a *SSOAdminStub) ListAccountAssignmentDeletionStatus(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentDeletionStatusInput) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error) {
    var output ssoadmin.ListAccountAssignmentDeletionStatusOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.ListAccountAssignmentDeletionStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) ListAccountAssignmentDeletionStatusAsync(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentDeletionStatusInput) *SsoadminListAccountAssignmentDeletionStatusResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.ListAccountAssignmentDeletionStatus", input)
    return &SsoadminListAccountAssignmentDeletionStatusResult{Result: future}
}

func (a *SSOAdminStub) ListAccountAssignments(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentsInput) (*ssoadmin.ListAccountAssignmentsOutput, error) {
    var output ssoadmin.ListAccountAssignmentsOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.ListAccountAssignments", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) ListAccountAssignmentsAsync(ctx workflow.Context, input *ssoadmin.ListAccountAssignmentsInput) *SsoadminListAccountAssignmentsResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.ListAccountAssignments", input)
    return &SsoadminListAccountAssignmentsResult{Result: future}
}

func (a *SSOAdminStub) ListAccountsForProvisionedPermissionSet(ctx workflow.Context, input *ssoadmin.ListAccountsForProvisionedPermissionSetInput) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error) {
    var output ssoadmin.ListAccountsForProvisionedPermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.ListAccountsForProvisionedPermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) ListAccountsForProvisionedPermissionSetAsync(ctx workflow.Context, input *ssoadmin.ListAccountsForProvisionedPermissionSetInput) *SsoadminListAccountsForProvisionedPermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.ListAccountsForProvisionedPermissionSet", input)
    return &SsoadminListAccountsForProvisionedPermissionSetResult{Result: future}
}

func (a *SSOAdminStub) ListInstances(ctx workflow.Context, input *ssoadmin.ListInstancesInput) (*ssoadmin.ListInstancesOutput, error) {
    var output ssoadmin.ListInstancesOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.ListInstances", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) ListInstancesAsync(ctx workflow.Context, input *ssoadmin.ListInstancesInput) *SsoadminListInstancesResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.ListInstances", input)
    return &SsoadminListInstancesResult{Result: future}
}

func (a *SSOAdminStub) ListManagedPoliciesInPermissionSet(ctx workflow.Context, input *ssoadmin.ListManagedPoliciesInPermissionSetInput) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error) {
    var output ssoadmin.ListManagedPoliciesInPermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.ListManagedPoliciesInPermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) ListManagedPoliciesInPermissionSetAsync(ctx workflow.Context, input *ssoadmin.ListManagedPoliciesInPermissionSetInput) *SsoadminListManagedPoliciesInPermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.ListManagedPoliciesInPermissionSet", input)
    return &SsoadminListManagedPoliciesInPermissionSetResult{Result: future}
}

func (a *SSOAdminStub) ListPermissionSetProvisioningStatus(ctx workflow.Context, input *ssoadmin.ListPermissionSetProvisioningStatusInput) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error) {
    var output ssoadmin.ListPermissionSetProvisioningStatusOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.ListPermissionSetProvisioningStatus", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) ListPermissionSetProvisioningStatusAsync(ctx workflow.Context, input *ssoadmin.ListPermissionSetProvisioningStatusInput) *SsoadminListPermissionSetProvisioningStatusResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.ListPermissionSetProvisioningStatus", input)
    return &SsoadminListPermissionSetProvisioningStatusResult{Result: future}
}

func (a *SSOAdminStub) ListPermissionSets(ctx workflow.Context, input *ssoadmin.ListPermissionSetsInput) (*ssoadmin.ListPermissionSetsOutput, error) {
    var output ssoadmin.ListPermissionSetsOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.ListPermissionSets", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) ListPermissionSetsAsync(ctx workflow.Context, input *ssoadmin.ListPermissionSetsInput) *SsoadminListPermissionSetsResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.ListPermissionSets", input)
    return &SsoadminListPermissionSetsResult{Result: future}
}

func (a *SSOAdminStub) ListPermissionSetsProvisionedToAccount(ctx workflow.Context, input *ssoadmin.ListPermissionSetsProvisionedToAccountInput) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error) {
    var output ssoadmin.ListPermissionSetsProvisionedToAccountOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.ListPermissionSetsProvisionedToAccount", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) ListPermissionSetsProvisionedToAccountAsync(ctx workflow.Context, input *ssoadmin.ListPermissionSetsProvisionedToAccountInput) *SsoadminListPermissionSetsProvisionedToAccountResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.ListPermissionSetsProvisionedToAccount", input)
    return &SsoadminListPermissionSetsProvisionedToAccountResult{Result: future}
}

func (a *SSOAdminStub) ListTagsForResource(ctx workflow.Context, input *ssoadmin.ListTagsForResourceInput) (*ssoadmin.ListTagsForResourceOutput, error) {
    var output ssoadmin.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.ListTagsForResource", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) ListTagsForResourceAsync(ctx workflow.Context, input *ssoadmin.ListTagsForResourceInput) *SsoadminListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.ListTagsForResource", input)
    return &SsoadminListTagsForResourceResult{Result: future}
}

func (a *SSOAdminStub) ProvisionPermissionSet(ctx workflow.Context, input *ssoadmin.ProvisionPermissionSetInput) (*ssoadmin.ProvisionPermissionSetOutput, error) {
    var output ssoadmin.ProvisionPermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.ProvisionPermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) ProvisionPermissionSetAsync(ctx workflow.Context, input *ssoadmin.ProvisionPermissionSetInput) *SsoadminProvisionPermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.ProvisionPermissionSet", input)
    return &SsoadminProvisionPermissionSetResult{Result: future}
}

func (a *SSOAdminStub) PutInlinePolicyToPermissionSet(ctx workflow.Context, input *ssoadmin.PutInlinePolicyToPermissionSetInput) (*ssoadmin.PutInlinePolicyToPermissionSetOutput, error) {
    var output ssoadmin.PutInlinePolicyToPermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.PutInlinePolicyToPermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) PutInlinePolicyToPermissionSetAsync(ctx workflow.Context, input *ssoadmin.PutInlinePolicyToPermissionSetInput) *SsoadminPutInlinePolicyToPermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.PutInlinePolicyToPermissionSet", input)
    return &SsoadminPutInlinePolicyToPermissionSetResult{Result: future}
}

func (a *SSOAdminStub) TagResource(ctx workflow.Context, input *ssoadmin.TagResourceInput) (*ssoadmin.TagResourceOutput, error) {
    var output ssoadmin.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.TagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) TagResourceAsync(ctx workflow.Context, input *ssoadmin.TagResourceInput) *SsoadminTagResourceResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.TagResource", input)
    return &SsoadminTagResourceResult{Result: future}
}

func (a *SSOAdminStub) UntagResource(ctx workflow.Context, input *ssoadmin.UntagResourceInput) (*ssoadmin.UntagResourceOutput, error) {
    var output ssoadmin.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.UntagResource", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) UntagResourceAsync(ctx workflow.Context, input *ssoadmin.UntagResourceInput) *SsoadminUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.UntagResource", input)
    return &SsoadminUntagResourceResult{Result: future}
}

func (a *SSOAdminStub) UpdatePermissionSet(ctx workflow.Context, input *ssoadmin.UpdatePermissionSetInput) (*ssoadmin.UpdatePermissionSetOutput, error) {
    var output ssoadmin.UpdatePermissionSetOutput
    err := workflow.ExecuteActivity(ctx, "SSOAdmin.UpdatePermissionSet", input).Get(ctx, &output)
    return &output, err
}

func (a *SSOAdminStub) UpdatePermissionSetAsync(ctx workflow.Context, input *ssoadmin.UpdatePermissionSetInput) *SsoadminUpdatePermissionSetResult {
    future := workflow.ExecuteActivity(ctx, "SSOAdmin.UpdatePermissionSet", input)
    return &SsoadminUpdatePermissionSetResult{Result: future}
}