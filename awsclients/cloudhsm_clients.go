package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CloudHSMClient interface {
       AddTagsToResource(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error)
       AddTagsToResourceAsync(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) *CloudhsmAddTagsToResourceResult

       CreateHapg(ctx workflow.Context, input *cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error)
       CreateHapgAsync(ctx workflow.Context, input *cloudhsm.CreateHapgInput) *CloudhsmCreateHapgResult

       CreateHsm(ctx workflow.Context, input *cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error)
       CreateHsmAsync(ctx workflow.Context, input *cloudhsm.CreateHsmInput) *CloudhsmCreateHsmResult

       CreateLunaClient(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error)
       CreateLunaClientAsync(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) *CloudhsmCreateLunaClientResult

       DeleteHapg(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error)
       DeleteHapgAsync(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) *CloudhsmDeleteHapgResult

       DeleteHsm(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error)
       DeleteHsmAsync(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) *CloudhsmDeleteHsmResult

       DeleteLunaClient(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error)
       DeleteLunaClientAsync(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) *CloudhsmDeleteLunaClientResult

       DescribeHapg(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error)
       DescribeHapgAsync(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) *CloudhsmDescribeHapgResult

       DescribeHsm(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error)
       DescribeHsmAsync(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) *CloudhsmDescribeHsmResult

       DescribeLunaClient(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error)
       DescribeLunaClientAsync(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) *CloudhsmDescribeLunaClientResult

       GetConfig(ctx workflow.Context, input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error)
       GetConfigAsync(ctx workflow.Context, input *cloudhsm.GetConfigInput) *CloudhsmGetConfigResult

       ListAvailableZones(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error)
       ListAvailableZonesAsync(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) *CloudhsmListAvailableZonesResult

       ListHapgs(ctx workflow.Context, input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error)
       ListHapgsAsync(ctx workflow.Context, input *cloudhsm.ListHapgsInput) *CloudhsmListHapgsResult

       ListHsms(ctx workflow.Context, input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error)
       ListHsmsAsync(ctx workflow.Context, input *cloudhsm.ListHsmsInput) *CloudhsmListHsmsResult

       ListLunaClients(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error)
       ListLunaClientsAsync(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) *CloudhsmListLunaClientsResult

       ListTagsForResource(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) *CloudhsmListTagsForResourceResult

       ModifyHapg(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error)
       ModifyHapgAsync(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) *CloudhsmModifyHapgResult

       ModifyHsm(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error)
       ModifyHsmAsync(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) *CloudhsmModifyHsmResult

       ModifyLunaClient(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error)
       ModifyLunaClientAsync(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) *CloudhsmModifyLunaClientResult

       RemoveTagsFromResource(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error)
       RemoveTagsFromResourceAsync(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) *CloudhsmRemoveTagsFromResourceResult
}

type CloudhsmAddTagsToResourceResult struct {
	Result workflow.Future
}

func (r *CloudhsmAddTagsToResourceResult) Get(ctx workflow.Context) (*cloudhsm.AddTagsToResourceOutput, error) {
    var output cloudhsm.AddTagsToResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmCreateHapgResult struct {
	Result workflow.Future
}

func (r *CloudhsmCreateHapgResult) Get(ctx workflow.Context) (*cloudhsm.CreateHapgOutput, error) {
    var output cloudhsm.CreateHapgOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmCreateHsmResult struct {
	Result workflow.Future
}

func (r *CloudhsmCreateHsmResult) Get(ctx workflow.Context) (*cloudhsm.CreateHsmOutput, error) {
    var output cloudhsm.CreateHsmOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmCreateLunaClientResult struct {
	Result workflow.Future
}

func (r *CloudhsmCreateLunaClientResult) Get(ctx workflow.Context) (*cloudhsm.CreateLunaClientOutput, error) {
    var output cloudhsm.CreateLunaClientOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmDeleteHapgResult struct {
	Result workflow.Future
}

func (r *CloudhsmDeleteHapgResult) Get(ctx workflow.Context) (*cloudhsm.DeleteHapgOutput, error) {
    var output cloudhsm.DeleteHapgOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmDeleteHsmResult struct {
	Result workflow.Future
}

func (r *CloudhsmDeleteHsmResult) Get(ctx workflow.Context) (*cloudhsm.DeleteHsmOutput, error) {
    var output cloudhsm.DeleteHsmOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmDeleteLunaClientResult struct {
	Result workflow.Future
}

func (r *CloudhsmDeleteLunaClientResult) Get(ctx workflow.Context) (*cloudhsm.DeleteLunaClientOutput, error) {
    var output cloudhsm.DeleteLunaClientOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmDescribeHapgResult struct {
	Result workflow.Future
}

func (r *CloudhsmDescribeHapgResult) Get(ctx workflow.Context) (*cloudhsm.DescribeHapgOutput, error) {
    var output cloudhsm.DescribeHapgOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmDescribeHsmResult struct {
	Result workflow.Future
}

func (r *CloudhsmDescribeHsmResult) Get(ctx workflow.Context) (*cloudhsm.DescribeHsmOutput, error) {
    var output cloudhsm.DescribeHsmOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmDescribeLunaClientResult struct {
	Result workflow.Future
}

func (r *CloudhsmDescribeLunaClientResult) Get(ctx workflow.Context) (*cloudhsm.DescribeLunaClientOutput, error) {
    var output cloudhsm.DescribeLunaClientOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmGetConfigResult struct {
	Result workflow.Future
}

func (r *CloudhsmGetConfigResult) Get(ctx workflow.Context) (*cloudhsm.GetConfigOutput, error) {
    var output cloudhsm.GetConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmListAvailableZonesResult struct {
	Result workflow.Future
}

func (r *CloudhsmListAvailableZonesResult) Get(ctx workflow.Context) (*cloudhsm.ListAvailableZonesOutput, error) {
    var output cloudhsm.ListAvailableZonesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmListHapgsResult struct {
	Result workflow.Future
}

func (r *CloudhsmListHapgsResult) Get(ctx workflow.Context) (*cloudhsm.ListHapgsOutput, error) {
    var output cloudhsm.ListHapgsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmListHsmsResult struct {
	Result workflow.Future
}

func (r *CloudhsmListHsmsResult) Get(ctx workflow.Context) (*cloudhsm.ListHsmsOutput, error) {
    var output cloudhsm.ListHsmsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmListLunaClientsResult struct {
	Result workflow.Future
}

func (r *CloudhsmListLunaClientsResult) Get(ctx workflow.Context) (*cloudhsm.ListLunaClientsOutput, error) {
    var output cloudhsm.ListLunaClientsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CloudhsmListTagsForResourceResult) Get(ctx workflow.Context) (*cloudhsm.ListTagsForResourceOutput, error) {
    var output cloudhsm.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmModifyHapgResult struct {
	Result workflow.Future
}

func (r *CloudhsmModifyHapgResult) Get(ctx workflow.Context) (*cloudhsm.ModifyHapgOutput, error) {
    var output cloudhsm.ModifyHapgOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmModifyHsmResult struct {
	Result workflow.Future
}

func (r *CloudhsmModifyHsmResult) Get(ctx workflow.Context) (*cloudhsm.ModifyHsmOutput, error) {
    var output cloudhsm.ModifyHsmOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmModifyLunaClientResult struct {
	Result workflow.Future
}

func (r *CloudhsmModifyLunaClientResult) Get(ctx workflow.Context) (*cloudhsm.ModifyLunaClientOutput, error) {
    var output cloudhsm.ModifyLunaClientOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudhsmRemoveTagsFromResourceResult struct {
	Result workflow.Future
}

func (r *CloudhsmRemoveTagsFromResourceResult) Get(ctx workflow.Context) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
    var output cloudhsm.RemoveTagsFromResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CloudHSMStub struct {
    activities awsactivities.CloudHSMActivities
}

func NewCloudHSMStub() CloudHSMClient {
    return &CloudHSMStub{}
}

func (a *CloudHSMStub) AddTagsToResource(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error) {
    var output cloudhsm.AddTagsToResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTagsToResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) AddTagsToResourceAsync(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) *CloudhsmAddTagsToResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddTagsToResource, input)
    return &CloudhsmAddTagsToResourceResult{Result: future}
}

func (a *CloudHSMStub) CreateHapg(ctx workflow.Context, input *cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error) {
    var output cloudhsm.CreateHapgOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHapg, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) CreateHapgAsync(ctx workflow.Context, input *cloudhsm.CreateHapgInput) *CloudhsmCreateHapgResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateHapg, input)
    return &CloudhsmCreateHapgResult{Result: future}
}

func (a *CloudHSMStub) CreateHsm(ctx workflow.Context, input *cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error) {
    var output cloudhsm.CreateHsmOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHsm, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) CreateHsmAsync(ctx workflow.Context, input *cloudhsm.CreateHsmInput) *CloudhsmCreateHsmResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateHsm, input)
    return &CloudhsmCreateHsmResult{Result: future}
}

func (a *CloudHSMStub) CreateLunaClient(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error) {
    var output cloudhsm.CreateLunaClientOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLunaClient, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) CreateLunaClientAsync(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) *CloudhsmCreateLunaClientResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLunaClient, input)
    return &CloudhsmCreateLunaClientResult{Result: future}
}

func (a *CloudHSMStub) DeleteHapg(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error) {
    var output cloudhsm.DeleteHapgOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteHapg, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) DeleteHapgAsync(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) *CloudhsmDeleteHapgResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteHapg, input)
    return &CloudhsmDeleteHapgResult{Result: future}
}

func (a *CloudHSMStub) DeleteHsm(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error) {
    var output cloudhsm.DeleteHsmOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteHsm, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) DeleteHsmAsync(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) *CloudhsmDeleteHsmResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteHsm, input)
    return &CloudhsmDeleteHsmResult{Result: future}
}

func (a *CloudHSMStub) DeleteLunaClient(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error) {
    var output cloudhsm.DeleteLunaClientOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLunaClient, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) DeleteLunaClientAsync(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) *CloudhsmDeleteLunaClientResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLunaClient, input)
    return &CloudhsmDeleteLunaClientResult{Result: future}
}

func (a *CloudHSMStub) DescribeHapg(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error) {
    var output cloudhsm.DescribeHapgOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHapg, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) DescribeHapgAsync(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) *CloudhsmDescribeHapgResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeHapg, input)
    return &CloudhsmDescribeHapgResult{Result: future}
}

func (a *CloudHSMStub) DescribeHsm(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error) {
    var output cloudhsm.DescribeHsmOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHsm, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) DescribeHsmAsync(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) *CloudhsmDescribeHsmResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeHsm, input)
    return &CloudhsmDescribeHsmResult{Result: future}
}

func (a *CloudHSMStub) DescribeLunaClient(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error) {
    var output cloudhsm.DescribeLunaClientOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLunaClient, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) DescribeLunaClientAsync(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) *CloudhsmDescribeLunaClientResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLunaClient, input)
    return &CloudhsmDescribeLunaClientResult{Result: future}
}

func (a *CloudHSMStub) GetConfig(ctx workflow.Context, input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error) {
    var output cloudhsm.GetConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) GetConfigAsync(ctx workflow.Context, input *cloudhsm.GetConfigInput) *CloudhsmGetConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetConfig, input)
    return &CloudhsmGetConfigResult{Result: future}
}

func (a *CloudHSMStub) ListAvailableZones(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error) {
    var output cloudhsm.ListAvailableZonesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAvailableZones, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) ListAvailableZonesAsync(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) *CloudhsmListAvailableZonesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAvailableZones, input)
    return &CloudhsmListAvailableZonesResult{Result: future}
}

func (a *CloudHSMStub) ListHapgs(ctx workflow.Context, input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error) {
    var output cloudhsm.ListHapgsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHapgs, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) ListHapgsAsync(ctx workflow.Context, input *cloudhsm.ListHapgsInput) *CloudhsmListHapgsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHapgs, input)
    return &CloudhsmListHapgsResult{Result: future}
}

func (a *CloudHSMStub) ListHsms(ctx workflow.Context, input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error) {
    var output cloudhsm.ListHsmsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHsms, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) ListHsmsAsync(ctx workflow.Context, input *cloudhsm.ListHsmsInput) *CloudhsmListHsmsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHsms, input)
    return &CloudhsmListHsmsResult{Result: future}
}

func (a *CloudHSMStub) ListLunaClients(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error) {
    var output cloudhsm.ListLunaClientsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLunaClients, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) ListLunaClientsAsync(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) *CloudhsmListLunaClientsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListLunaClients, input)
    return &CloudhsmListLunaClientsResult{Result: future}
}

func (a *CloudHSMStub) ListTagsForResource(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error) {
    var output cloudhsm.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) ListTagsForResourceAsync(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) *CloudhsmListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &CloudhsmListTagsForResourceResult{Result: future}
}

func (a *CloudHSMStub) ModifyHapg(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error) {
    var output cloudhsm.ModifyHapgOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyHapg, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) ModifyHapgAsync(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) *CloudhsmModifyHapgResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyHapg, input)
    return &CloudhsmModifyHapgResult{Result: future}
}

func (a *CloudHSMStub) ModifyHsm(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error) {
    var output cloudhsm.ModifyHsmOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyHsm, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) ModifyHsmAsync(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) *CloudhsmModifyHsmResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyHsm, input)
    return &CloudhsmModifyHsmResult{Result: future}
}

func (a *CloudHSMStub) ModifyLunaClient(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error) {
    var output cloudhsm.ModifyLunaClientOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyLunaClient, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) ModifyLunaClientAsync(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) *CloudhsmModifyLunaClientResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyLunaClient, input)
    return &CloudhsmModifyLunaClientResult{Result: future}
}

func (a *CloudHSMStub) RemoveTagsFromResource(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
    var output cloudhsm.RemoveTagsFromResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CloudHSMStub) RemoveTagsFromResourceAsync(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) *CloudhsmRemoveTagsFromResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromResource, input)
    return &CloudhsmRemoveTagsFromResourceResult{Result: future}
}