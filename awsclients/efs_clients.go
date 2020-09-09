package awsclients

import (
	"github.com/aws/aws-sdk-go/service/efs"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type EFSClient interface {
       CreateAccessPoint(ctx workflow.Context, input *efs.CreateAccessPointInput) (*efs.CreateAccessPointOutput, error)
       CreateAccessPointAsync(ctx workflow.Context, input *efs.CreateAccessPointInput) *EfsCreateAccessPointResult

       CreateFileSystem(ctx workflow.Context, input *efs.CreateFileSystemInput) (*efs.FileSystemDescription, error)
       CreateFileSystemAsync(ctx workflow.Context, input *efs.CreateFileSystemInput) *EfsCreateFileSystemResult

       CreateMountTarget(ctx workflow.Context, input *efs.CreateMountTargetInput) (*efs.MountTargetDescription, error)
       CreateMountTargetAsync(ctx workflow.Context, input *efs.CreateMountTargetInput) *EfsCreateMountTargetResult

       CreateTags(ctx workflow.Context, input *efs.CreateTagsInput) (*efs.CreateTagsOutput, error)
       CreateTagsAsync(ctx workflow.Context, input *efs.CreateTagsInput) *EfsCreateTagsResult

       DeleteAccessPoint(ctx workflow.Context, input *efs.DeleteAccessPointInput) (*efs.DeleteAccessPointOutput, error)
       DeleteAccessPointAsync(ctx workflow.Context, input *efs.DeleteAccessPointInput) *EfsDeleteAccessPointResult

       DeleteFileSystem(ctx workflow.Context, input *efs.DeleteFileSystemInput) (*efs.DeleteFileSystemOutput, error)
       DeleteFileSystemAsync(ctx workflow.Context, input *efs.DeleteFileSystemInput) *EfsDeleteFileSystemResult

       DeleteFileSystemPolicy(ctx workflow.Context, input *efs.DeleteFileSystemPolicyInput) (*efs.DeleteFileSystemPolicyOutput, error)
       DeleteFileSystemPolicyAsync(ctx workflow.Context, input *efs.DeleteFileSystemPolicyInput) *EfsDeleteFileSystemPolicyResult

       DeleteMountTarget(ctx workflow.Context, input *efs.DeleteMountTargetInput) (*efs.DeleteMountTargetOutput, error)
       DeleteMountTargetAsync(ctx workflow.Context, input *efs.DeleteMountTargetInput) *EfsDeleteMountTargetResult

       DeleteTags(ctx workflow.Context, input *efs.DeleteTagsInput) (*efs.DeleteTagsOutput, error)
       DeleteTagsAsync(ctx workflow.Context, input *efs.DeleteTagsInput) *EfsDeleteTagsResult

       DescribeAccessPoints(ctx workflow.Context, input *efs.DescribeAccessPointsInput) (*efs.DescribeAccessPointsOutput, error)
       DescribeAccessPointsAsync(ctx workflow.Context, input *efs.DescribeAccessPointsInput) *EfsDescribeAccessPointsResult

       DescribeBackupPolicy(ctx workflow.Context, input *efs.DescribeBackupPolicyInput) (*efs.DescribeBackupPolicyOutput, error)
       DescribeBackupPolicyAsync(ctx workflow.Context, input *efs.DescribeBackupPolicyInput) *EfsDescribeBackupPolicyResult

       DescribeFileSystemPolicy(ctx workflow.Context, input *efs.DescribeFileSystemPolicyInput) (*efs.DescribeFileSystemPolicyOutput, error)
       DescribeFileSystemPolicyAsync(ctx workflow.Context, input *efs.DescribeFileSystemPolicyInput) *EfsDescribeFileSystemPolicyResult

       DescribeFileSystems(ctx workflow.Context, input *efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error)
       DescribeFileSystemsAsync(ctx workflow.Context, input *efs.DescribeFileSystemsInput) *EfsDescribeFileSystemsResult

       DescribeLifecycleConfiguration(ctx workflow.Context, input *efs.DescribeLifecycleConfigurationInput) (*efs.DescribeLifecycleConfigurationOutput, error)
       DescribeLifecycleConfigurationAsync(ctx workflow.Context, input *efs.DescribeLifecycleConfigurationInput) *EfsDescribeLifecycleConfigurationResult

       DescribeMountTargetSecurityGroups(ctx workflow.Context, input *efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error)
       DescribeMountTargetSecurityGroupsAsync(ctx workflow.Context, input *efs.DescribeMountTargetSecurityGroupsInput) *EfsDescribeMountTargetSecurityGroupsResult

       DescribeMountTargets(ctx workflow.Context, input *efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error)
       DescribeMountTargetsAsync(ctx workflow.Context, input *efs.DescribeMountTargetsInput) *EfsDescribeMountTargetsResult

       DescribeTags(ctx workflow.Context, input *efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error)
       DescribeTagsAsync(ctx workflow.Context, input *efs.DescribeTagsInput) *EfsDescribeTagsResult

       ListTagsForResource(ctx workflow.Context, input *efs.ListTagsForResourceInput) (*efs.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *efs.ListTagsForResourceInput) *EfsListTagsForResourceResult

       ModifyMountTargetSecurityGroups(ctx workflow.Context, input *efs.ModifyMountTargetSecurityGroupsInput) (*efs.ModifyMountTargetSecurityGroupsOutput, error)
       ModifyMountTargetSecurityGroupsAsync(ctx workflow.Context, input *efs.ModifyMountTargetSecurityGroupsInput) *EfsModifyMountTargetSecurityGroupsResult

       PutBackupPolicy(ctx workflow.Context, input *efs.PutBackupPolicyInput) (*efs.PutBackupPolicyOutput, error)
       PutBackupPolicyAsync(ctx workflow.Context, input *efs.PutBackupPolicyInput) *EfsPutBackupPolicyResult

       PutFileSystemPolicy(ctx workflow.Context, input *efs.PutFileSystemPolicyInput) (*efs.PutFileSystemPolicyOutput, error)
       PutFileSystemPolicyAsync(ctx workflow.Context, input *efs.PutFileSystemPolicyInput) *EfsPutFileSystemPolicyResult

       PutLifecycleConfiguration(ctx workflow.Context, input *efs.PutLifecycleConfigurationInput) (*efs.PutLifecycleConfigurationOutput, error)
       PutLifecycleConfigurationAsync(ctx workflow.Context, input *efs.PutLifecycleConfigurationInput) *EfsPutLifecycleConfigurationResult

       TagResource(ctx workflow.Context, input *efs.TagResourceInput) (*efs.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *efs.TagResourceInput) *EfsTagResourceResult

       UntagResource(ctx workflow.Context, input *efs.UntagResourceInput) (*efs.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *efs.UntagResourceInput) *EfsUntagResourceResult

       UpdateFileSystem(ctx workflow.Context, input *efs.UpdateFileSystemInput) (*efs.UpdateFileSystemOutput, error)
       UpdateFileSystemAsync(ctx workflow.Context, input *efs.UpdateFileSystemInput) *EfsUpdateFileSystemResult
}

type EfsCreateAccessPointResult struct {
	Result workflow.Future
}

func (r *EfsCreateAccessPointResult) Get(ctx workflow.Context) (*efs.CreateAccessPointOutput, error) {
    var output efs.CreateAccessPointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsCreateFileSystemResult struct {
	Result workflow.Future
}

func (r *EfsCreateFileSystemResult) Get(ctx workflow.Context) (*efs.FileSystemDescription, error) {
    var output efs.FileSystemDescription
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsCreateMountTargetResult struct {
	Result workflow.Future
}

func (r *EfsCreateMountTargetResult) Get(ctx workflow.Context) (*efs.MountTargetDescription, error) {
    var output efs.MountTargetDescription
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsCreateTagsResult struct {
	Result workflow.Future
}

func (r *EfsCreateTagsResult) Get(ctx workflow.Context) (*efs.CreateTagsOutput, error) {
    var output efs.CreateTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDeleteAccessPointResult struct {
	Result workflow.Future
}

func (r *EfsDeleteAccessPointResult) Get(ctx workflow.Context) (*efs.DeleteAccessPointOutput, error) {
    var output efs.DeleteAccessPointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDeleteFileSystemResult struct {
	Result workflow.Future
}

func (r *EfsDeleteFileSystemResult) Get(ctx workflow.Context) (*efs.DeleteFileSystemOutput, error) {
    var output efs.DeleteFileSystemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDeleteFileSystemPolicyResult struct {
	Result workflow.Future
}

func (r *EfsDeleteFileSystemPolicyResult) Get(ctx workflow.Context) (*efs.DeleteFileSystemPolicyOutput, error) {
    var output efs.DeleteFileSystemPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDeleteMountTargetResult struct {
	Result workflow.Future
}

func (r *EfsDeleteMountTargetResult) Get(ctx workflow.Context) (*efs.DeleteMountTargetOutput, error) {
    var output efs.DeleteMountTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDeleteTagsResult struct {
	Result workflow.Future
}

func (r *EfsDeleteTagsResult) Get(ctx workflow.Context) (*efs.DeleteTagsOutput, error) {
    var output efs.DeleteTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDescribeAccessPointsResult struct {
	Result workflow.Future
}

func (r *EfsDescribeAccessPointsResult) Get(ctx workflow.Context) (*efs.DescribeAccessPointsOutput, error) {
    var output efs.DescribeAccessPointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDescribeBackupPolicyResult struct {
	Result workflow.Future
}

func (r *EfsDescribeBackupPolicyResult) Get(ctx workflow.Context) (*efs.DescribeBackupPolicyOutput, error) {
    var output efs.DescribeBackupPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDescribeFileSystemPolicyResult struct {
	Result workflow.Future
}

func (r *EfsDescribeFileSystemPolicyResult) Get(ctx workflow.Context) (*efs.DescribeFileSystemPolicyOutput, error) {
    var output efs.DescribeFileSystemPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDescribeFileSystemsResult struct {
	Result workflow.Future
}

func (r *EfsDescribeFileSystemsResult) Get(ctx workflow.Context) (*efs.DescribeFileSystemsOutput, error) {
    var output efs.DescribeFileSystemsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDescribeLifecycleConfigurationResult struct {
	Result workflow.Future
}

func (r *EfsDescribeLifecycleConfigurationResult) Get(ctx workflow.Context) (*efs.DescribeLifecycleConfigurationOutput, error) {
    var output efs.DescribeLifecycleConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDescribeMountTargetSecurityGroupsResult struct {
	Result workflow.Future
}

func (r *EfsDescribeMountTargetSecurityGroupsResult) Get(ctx workflow.Context) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
    var output efs.DescribeMountTargetSecurityGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDescribeMountTargetsResult struct {
	Result workflow.Future
}

func (r *EfsDescribeMountTargetsResult) Get(ctx workflow.Context) (*efs.DescribeMountTargetsOutput, error) {
    var output efs.DescribeMountTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsDescribeTagsResult struct {
	Result workflow.Future
}

func (r *EfsDescribeTagsResult) Get(ctx workflow.Context) (*efs.DescribeTagsOutput, error) {
    var output efs.DescribeTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *EfsListTagsForResourceResult) Get(ctx workflow.Context) (*efs.ListTagsForResourceOutput, error) {
    var output efs.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsModifyMountTargetSecurityGroupsResult struct {
	Result workflow.Future
}

func (r *EfsModifyMountTargetSecurityGroupsResult) Get(ctx workflow.Context) (*efs.ModifyMountTargetSecurityGroupsOutput, error) {
    var output efs.ModifyMountTargetSecurityGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsPutBackupPolicyResult struct {
	Result workflow.Future
}

func (r *EfsPutBackupPolicyResult) Get(ctx workflow.Context) (*efs.PutBackupPolicyOutput, error) {
    var output efs.PutBackupPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsPutFileSystemPolicyResult struct {
	Result workflow.Future
}

func (r *EfsPutFileSystemPolicyResult) Get(ctx workflow.Context) (*efs.PutFileSystemPolicyOutput, error) {
    var output efs.PutFileSystemPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsPutLifecycleConfigurationResult struct {
	Result workflow.Future
}

func (r *EfsPutLifecycleConfigurationResult) Get(ctx workflow.Context) (*efs.PutLifecycleConfigurationOutput, error) {
    var output efs.PutLifecycleConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsTagResourceResult struct {
	Result workflow.Future
}

func (r *EfsTagResourceResult) Get(ctx workflow.Context) (*efs.TagResourceOutput, error) {
    var output efs.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsUntagResourceResult struct {
	Result workflow.Future
}

func (r *EfsUntagResourceResult) Get(ctx workflow.Context) (*efs.UntagResourceOutput, error) {
    var output efs.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EfsUpdateFileSystemResult struct {
	Result workflow.Future
}

func (r *EfsUpdateFileSystemResult) Get(ctx workflow.Context) (*efs.UpdateFileSystemOutput, error) {
    var output efs.UpdateFileSystemOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EFSStub struct {
    activities awsactivities.EFSActivities
}

func NewEFSStub() EFSClient {
    return &EFSStub{}
}

func (a *EFSStub) CreateAccessPoint(ctx workflow.Context, input *efs.CreateAccessPointInput) (*efs.CreateAccessPointOutput, error) {
    var output efs.CreateAccessPointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateAccessPoint, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) CreateAccessPointAsync(ctx workflow.Context, input *efs.CreateAccessPointInput) *EfsCreateAccessPointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateAccessPoint, input)
    return &EfsCreateAccessPointResult{Result: future}
}

func (a *EFSStub) CreateFileSystem(ctx workflow.Context, input *efs.CreateFileSystemInput) (*efs.FileSystemDescription, error) {
    var output efs.FileSystemDescription
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFileSystem, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) CreateFileSystemAsync(ctx workflow.Context, input *efs.CreateFileSystemInput) *EfsCreateFileSystemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateFileSystem, input)
    return &EfsCreateFileSystemResult{Result: future}
}

func (a *EFSStub) CreateMountTarget(ctx workflow.Context, input *efs.CreateMountTargetInput) (*efs.MountTargetDescription, error) {
    var output efs.MountTargetDescription
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMountTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) CreateMountTargetAsync(ctx workflow.Context, input *efs.CreateMountTargetInput) *EfsCreateMountTargetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateMountTarget, input)
    return &EfsCreateMountTargetResult{Result: future}
}

func (a *EFSStub) CreateTags(ctx workflow.Context, input *efs.CreateTagsInput) (*efs.CreateTagsOutput, error) {
    var output efs.CreateTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTags, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) CreateTagsAsync(ctx workflow.Context, input *efs.CreateTagsInput) *EfsCreateTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTags, input)
    return &EfsCreateTagsResult{Result: future}
}

func (a *EFSStub) DeleteAccessPoint(ctx workflow.Context, input *efs.DeleteAccessPointInput) (*efs.DeleteAccessPointOutput, error) {
    var output efs.DeleteAccessPointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAccessPoint, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DeleteAccessPointAsync(ctx workflow.Context, input *efs.DeleteAccessPointInput) *EfsDeleteAccessPointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAccessPoint, input)
    return &EfsDeleteAccessPointResult{Result: future}
}

func (a *EFSStub) DeleteFileSystem(ctx workflow.Context, input *efs.DeleteFileSystemInput) (*efs.DeleteFileSystemOutput, error) {
    var output efs.DeleteFileSystemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFileSystem, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DeleteFileSystemAsync(ctx workflow.Context, input *efs.DeleteFileSystemInput) *EfsDeleteFileSystemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFileSystem, input)
    return &EfsDeleteFileSystemResult{Result: future}
}

func (a *EFSStub) DeleteFileSystemPolicy(ctx workflow.Context, input *efs.DeleteFileSystemPolicyInput) (*efs.DeleteFileSystemPolicyOutput, error) {
    var output efs.DeleteFileSystemPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFileSystemPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DeleteFileSystemPolicyAsync(ctx workflow.Context, input *efs.DeleteFileSystemPolicyInput) *EfsDeleteFileSystemPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteFileSystemPolicy, input)
    return &EfsDeleteFileSystemPolicyResult{Result: future}
}

func (a *EFSStub) DeleteMountTarget(ctx workflow.Context, input *efs.DeleteMountTargetInput) (*efs.DeleteMountTargetOutput, error) {
    var output efs.DeleteMountTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMountTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DeleteMountTargetAsync(ctx workflow.Context, input *efs.DeleteMountTargetInput) *EfsDeleteMountTargetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMountTarget, input)
    return &EfsDeleteMountTargetResult{Result: future}
}

func (a *EFSStub) DeleteTags(ctx workflow.Context, input *efs.DeleteTagsInput) (*efs.DeleteTagsOutput, error) {
    var output efs.DeleteTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DeleteTagsAsync(ctx workflow.Context, input *efs.DeleteTagsInput) *EfsDeleteTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input)
    return &EfsDeleteTagsResult{Result: future}
}

func (a *EFSStub) DescribeAccessPoints(ctx workflow.Context, input *efs.DescribeAccessPointsInput) (*efs.DescribeAccessPointsOutput, error) {
    var output efs.DescribeAccessPointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccessPoints, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DescribeAccessPointsAsync(ctx workflow.Context, input *efs.DescribeAccessPointsInput) *EfsDescribeAccessPointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAccessPoints, input)
    return &EfsDescribeAccessPointsResult{Result: future}
}

func (a *EFSStub) DescribeBackupPolicy(ctx workflow.Context, input *efs.DescribeBackupPolicyInput) (*efs.DescribeBackupPolicyOutput, error) {
    var output efs.DescribeBackupPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBackupPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DescribeBackupPolicyAsync(ctx workflow.Context, input *efs.DescribeBackupPolicyInput) *EfsDescribeBackupPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBackupPolicy, input)
    return &EfsDescribeBackupPolicyResult{Result: future}
}

func (a *EFSStub) DescribeFileSystemPolicy(ctx workflow.Context, input *efs.DescribeFileSystemPolicyInput) (*efs.DescribeFileSystemPolicyOutput, error) {
    var output efs.DescribeFileSystemPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFileSystemPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DescribeFileSystemPolicyAsync(ctx workflow.Context, input *efs.DescribeFileSystemPolicyInput) *EfsDescribeFileSystemPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFileSystemPolicy, input)
    return &EfsDescribeFileSystemPolicyResult{Result: future}
}

func (a *EFSStub) DescribeFileSystems(ctx workflow.Context, input *efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error) {
    var output efs.DescribeFileSystemsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFileSystems, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DescribeFileSystemsAsync(ctx workflow.Context, input *efs.DescribeFileSystemsInput) *EfsDescribeFileSystemsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeFileSystems, input)
    return &EfsDescribeFileSystemsResult{Result: future}
}

func (a *EFSStub) DescribeLifecycleConfiguration(ctx workflow.Context, input *efs.DescribeLifecycleConfigurationInput) (*efs.DescribeLifecycleConfigurationOutput, error) {
    var output efs.DescribeLifecycleConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLifecycleConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DescribeLifecycleConfigurationAsync(ctx workflow.Context, input *efs.DescribeLifecycleConfigurationInput) *EfsDescribeLifecycleConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLifecycleConfiguration, input)
    return &EfsDescribeLifecycleConfigurationResult{Result: future}
}

func (a *EFSStub) DescribeMountTargetSecurityGroups(ctx workflow.Context, input *efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
    var output efs.DescribeMountTargetSecurityGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMountTargetSecurityGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DescribeMountTargetSecurityGroupsAsync(ctx workflow.Context, input *efs.DescribeMountTargetSecurityGroupsInput) *EfsDescribeMountTargetSecurityGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeMountTargetSecurityGroups, input)
    return &EfsDescribeMountTargetSecurityGroupsResult{Result: future}
}

func (a *EFSStub) DescribeMountTargets(ctx workflow.Context, input *efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error) {
    var output efs.DescribeMountTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeMountTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DescribeMountTargetsAsync(ctx workflow.Context, input *efs.DescribeMountTargetsInput) *EfsDescribeMountTargetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeMountTargets, input)
    return &EfsDescribeMountTargetsResult{Result: future}
}

func (a *EFSStub) DescribeTags(ctx workflow.Context, input *efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error) {
    var output efs.DescribeTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) DescribeTagsAsync(ctx workflow.Context, input *efs.DescribeTagsInput) *EfsDescribeTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input)
    return &EfsDescribeTagsResult{Result: future}
}

func (a *EFSStub) ListTagsForResource(ctx workflow.Context, input *efs.ListTagsForResourceInput) (*efs.ListTagsForResourceOutput, error) {
    var output efs.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) ListTagsForResourceAsync(ctx workflow.Context, input *efs.ListTagsForResourceInput) *EfsListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &EfsListTagsForResourceResult{Result: future}
}

func (a *EFSStub) ModifyMountTargetSecurityGroups(ctx workflow.Context, input *efs.ModifyMountTargetSecurityGroupsInput) (*efs.ModifyMountTargetSecurityGroupsOutput, error) {
    var output efs.ModifyMountTargetSecurityGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyMountTargetSecurityGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) ModifyMountTargetSecurityGroupsAsync(ctx workflow.Context, input *efs.ModifyMountTargetSecurityGroupsInput) *EfsModifyMountTargetSecurityGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyMountTargetSecurityGroups, input)
    return &EfsModifyMountTargetSecurityGroupsResult{Result: future}
}

func (a *EFSStub) PutBackupPolicy(ctx workflow.Context, input *efs.PutBackupPolicyInput) (*efs.PutBackupPolicyOutput, error) {
    var output efs.PutBackupPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutBackupPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) PutBackupPolicyAsync(ctx workflow.Context, input *efs.PutBackupPolicyInput) *EfsPutBackupPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutBackupPolicy, input)
    return &EfsPutBackupPolicyResult{Result: future}
}

func (a *EFSStub) PutFileSystemPolicy(ctx workflow.Context, input *efs.PutFileSystemPolicyInput) (*efs.PutFileSystemPolicyOutput, error) {
    var output efs.PutFileSystemPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutFileSystemPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) PutFileSystemPolicyAsync(ctx workflow.Context, input *efs.PutFileSystemPolicyInput) *EfsPutFileSystemPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutFileSystemPolicy, input)
    return &EfsPutFileSystemPolicyResult{Result: future}
}

func (a *EFSStub) PutLifecycleConfiguration(ctx workflow.Context, input *efs.PutLifecycleConfigurationInput) (*efs.PutLifecycleConfigurationOutput, error) {
    var output efs.PutLifecycleConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutLifecycleConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) PutLifecycleConfigurationAsync(ctx workflow.Context, input *efs.PutLifecycleConfigurationInput) *EfsPutLifecycleConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutLifecycleConfiguration, input)
    return &EfsPutLifecycleConfigurationResult{Result: future}
}

func (a *EFSStub) TagResource(ctx workflow.Context, input *efs.TagResourceInput) (*efs.TagResourceOutput, error) {
    var output efs.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) TagResourceAsync(ctx workflow.Context, input *efs.TagResourceInput) *EfsTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &EfsTagResourceResult{Result: future}
}

func (a *EFSStub) UntagResource(ctx workflow.Context, input *efs.UntagResourceInput) (*efs.UntagResourceOutput, error) {
    var output efs.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) UntagResourceAsync(ctx workflow.Context, input *efs.UntagResourceInput) *EfsUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &EfsUntagResourceResult{Result: future}
}

func (a *EFSStub) UpdateFileSystem(ctx workflow.Context, input *efs.UpdateFileSystemInput) (*efs.UpdateFileSystemOutput, error) {
    var output efs.UpdateFileSystemOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateFileSystem, input).Get(ctx, &output)
    return &output, err
}

func (a *EFSStub) UpdateFileSystemAsync(ctx workflow.Context, input *efs.UpdateFileSystemInput) *EfsUpdateFileSystemResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateFileSystem, input)
    return &EfsUpdateFileSystemResult{Result: future}
}
