package awsclients

import (
	"github.com/aws/aws-sdk-go/service/dataexchange"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type DataExchangeClient interface {
       CancelJob(ctx workflow.Context, input *dataexchange.CancelJobInput) (*dataexchange.CancelJobOutput, error)
       CancelJobAsync(ctx workflow.Context, input *dataexchange.CancelJobInput) *DataexchangeCancelJobResult

       CreateDataSet(ctx workflow.Context, input *dataexchange.CreateDataSetInput) (*dataexchange.CreateDataSetOutput, error)
       CreateDataSetAsync(ctx workflow.Context, input *dataexchange.CreateDataSetInput) *DataexchangeCreateDataSetResult

       CreateJob(ctx workflow.Context, input *dataexchange.CreateJobInput) (*dataexchange.CreateJobOutput, error)
       CreateJobAsync(ctx workflow.Context, input *dataexchange.CreateJobInput) *DataexchangeCreateJobResult

       CreateRevision(ctx workflow.Context, input *dataexchange.CreateRevisionInput) (*dataexchange.CreateRevisionOutput, error)
       CreateRevisionAsync(ctx workflow.Context, input *dataexchange.CreateRevisionInput) *DataexchangeCreateRevisionResult

       DeleteAsset(ctx workflow.Context, input *dataexchange.DeleteAssetInput) (*dataexchange.DeleteAssetOutput, error)
       DeleteAssetAsync(ctx workflow.Context, input *dataexchange.DeleteAssetInput) *DataexchangeDeleteAssetResult

       DeleteDataSet(ctx workflow.Context, input *dataexchange.DeleteDataSetInput) (*dataexchange.DeleteDataSetOutput, error)
       DeleteDataSetAsync(ctx workflow.Context, input *dataexchange.DeleteDataSetInput) *DataexchangeDeleteDataSetResult

       DeleteRevision(ctx workflow.Context, input *dataexchange.DeleteRevisionInput) (*dataexchange.DeleteRevisionOutput, error)
       DeleteRevisionAsync(ctx workflow.Context, input *dataexchange.DeleteRevisionInput) *DataexchangeDeleteRevisionResult

       GetAsset(ctx workflow.Context, input *dataexchange.GetAssetInput) (*dataexchange.GetAssetOutput, error)
       GetAssetAsync(ctx workflow.Context, input *dataexchange.GetAssetInput) *DataexchangeGetAssetResult

       GetDataSet(ctx workflow.Context, input *dataexchange.GetDataSetInput) (*dataexchange.GetDataSetOutput, error)
       GetDataSetAsync(ctx workflow.Context, input *dataexchange.GetDataSetInput) *DataexchangeGetDataSetResult

       GetJob(ctx workflow.Context, input *dataexchange.GetJobInput) (*dataexchange.GetJobOutput, error)
       GetJobAsync(ctx workflow.Context, input *dataexchange.GetJobInput) *DataexchangeGetJobResult

       GetRevision(ctx workflow.Context, input *dataexchange.GetRevisionInput) (*dataexchange.GetRevisionOutput, error)
       GetRevisionAsync(ctx workflow.Context, input *dataexchange.GetRevisionInput) *DataexchangeGetRevisionResult

       ListDataSetRevisions(ctx workflow.Context, input *dataexchange.ListDataSetRevisionsInput) (*dataexchange.ListDataSetRevisionsOutput, error)
       ListDataSetRevisionsAsync(ctx workflow.Context, input *dataexchange.ListDataSetRevisionsInput) *DataexchangeListDataSetRevisionsResult

       ListDataSets(ctx workflow.Context, input *dataexchange.ListDataSetsInput) (*dataexchange.ListDataSetsOutput, error)
       ListDataSetsAsync(ctx workflow.Context, input *dataexchange.ListDataSetsInput) *DataexchangeListDataSetsResult

       ListJobs(ctx workflow.Context, input *dataexchange.ListJobsInput) (*dataexchange.ListJobsOutput, error)
       ListJobsAsync(ctx workflow.Context, input *dataexchange.ListJobsInput) *DataexchangeListJobsResult

       ListRevisionAssets(ctx workflow.Context, input *dataexchange.ListRevisionAssetsInput) (*dataexchange.ListRevisionAssetsOutput, error)
       ListRevisionAssetsAsync(ctx workflow.Context, input *dataexchange.ListRevisionAssetsInput) *DataexchangeListRevisionAssetsResult

       ListTagsForResource(ctx workflow.Context, input *dataexchange.ListTagsForResourceInput) (*dataexchange.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *dataexchange.ListTagsForResourceInput) *DataexchangeListTagsForResourceResult

       StartJob(ctx workflow.Context, input *dataexchange.StartJobInput) (*dataexchange.StartJobOutput, error)
       StartJobAsync(ctx workflow.Context, input *dataexchange.StartJobInput) *DataexchangeStartJobResult

       TagResource(ctx workflow.Context, input *dataexchange.TagResourceInput) (*dataexchange.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *dataexchange.TagResourceInput) *DataexchangeTagResourceResult

       UntagResource(ctx workflow.Context, input *dataexchange.UntagResourceInput) (*dataexchange.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *dataexchange.UntagResourceInput) *DataexchangeUntagResourceResult

       UpdateAsset(ctx workflow.Context, input *dataexchange.UpdateAssetInput) (*dataexchange.UpdateAssetOutput, error)
       UpdateAssetAsync(ctx workflow.Context, input *dataexchange.UpdateAssetInput) *DataexchangeUpdateAssetResult

       UpdateDataSet(ctx workflow.Context, input *dataexchange.UpdateDataSetInput) (*dataexchange.UpdateDataSetOutput, error)
       UpdateDataSetAsync(ctx workflow.Context, input *dataexchange.UpdateDataSetInput) *DataexchangeUpdateDataSetResult

       UpdateRevision(ctx workflow.Context, input *dataexchange.UpdateRevisionInput) (*dataexchange.UpdateRevisionOutput, error)
       UpdateRevisionAsync(ctx workflow.Context, input *dataexchange.UpdateRevisionInput) *DataexchangeUpdateRevisionResult
}

type DataexchangeCancelJobResult struct {
	Result workflow.Future
}

func (r *DataexchangeCancelJobResult) Get(ctx workflow.Context) (*dataexchange.CancelJobOutput, error) {
    var output dataexchange.CancelJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeCreateDataSetResult struct {
	Result workflow.Future
}

func (r *DataexchangeCreateDataSetResult) Get(ctx workflow.Context) (*dataexchange.CreateDataSetOutput, error) {
    var output dataexchange.CreateDataSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeCreateJobResult struct {
	Result workflow.Future
}

func (r *DataexchangeCreateJobResult) Get(ctx workflow.Context) (*dataexchange.CreateJobOutput, error) {
    var output dataexchange.CreateJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeCreateRevisionResult struct {
	Result workflow.Future
}

func (r *DataexchangeCreateRevisionResult) Get(ctx workflow.Context) (*dataexchange.CreateRevisionOutput, error) {
    var output dataexchange.CreateRevisionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeDeleteAssetResult struct {
	Result workflow.Future
}

func (r *DataexchangeDeleteAssetResult) Get(ctx workflow.Context) (*dataexchange.DeleteAssetOutput, error) {
    var output dataexchange.DeleteAssetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeDeleteDataSetResult struct {
	Result workflow.Future
}

func (r *DataexchangeDeleteDataSetResult) Get(ctx workflow.Context) (*dataexchange.DeleteDataSetOutput, error) {
    var output dataexchange.DeleteDataSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeDeleteRevisionResult struct {
	Result workflow.Future
}

func (r *DataexchangeDeleteRevisionResult) Get(ctx workflow.Context) (*dataexchange.DeleteRevisionOutput, error) {
    var output dataexchange.DeleteRevisionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeGetAssetResult struct {
	Result workflow.Future
}

func (r *DataexchangeGetAssetResult) Get(ctx workflow.Context) (*dataexchange.GetAssetOutput, error) {
    var output dataexchange.GetAssetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeGetDataSetResult struct {
	Result workflow.Future
}

func (r *DataexchangeGetDataSetResult) Get(ctx workflow.Context) (*dataexchange.GetDataSetOutput, error) {
    var output dataexchange.GetDataSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeGetJobResult struct {
	Result workflow.Future
}

func (r *DataexchangeGetJobResult) Get(ctx workflow.Context) (*dataexchange.GetJobOutput, error) {
    var output dataexchange.GetJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeGetRevisionResult struct {
	Result workflow.Future
}

func (r *DataexchangeGetRevisionResult) Get(ctx workflow.Context) (*dataexchange.GetRevisionOutput, error) {
    var output dataexchange.GetRevisionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeListDataSetRevisionsResult struct {
	Result workflow.Future
}

func (r *DataexchangeListDataSetRevisionsResult) Get(ctx workflow.Context) (*dataexchange.ListDataSetRevisionsOutput, error) {
    var output dataexchange.ListDataSetRevisionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeListDataSetsResult struct {
	Result workflow.Future
}

func (r *DataexchangeListDataSetsResult) Get(ctx workflow.Context) (*dataexchange.ListDataSetsOutput, error) {
    var output dataexchange.ListDataSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeListJobsResult struct {
	Result workflow.Future
}

func (r *DataexchangeListJobsResult) Get(ctx workflow.Context) (*dataexchange.ListJobsOutput, error) {
    var output dataexchange.ListJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeListRevisionAssetsResult struct {
	Result workflow.Future
}

func (r *DataexchangeListRevisionAssetsResult) Get(ctx workflow.Context) (*dataexchange.ListRevisionAssetsOutput, error) {
    var output dataexchange.ListRevisionAssetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *DataexchangeListTagsForResourceResult) Get(ctx workflow.Context) (*dataexchange.ListTagsForResourceOutput, error) {
    var output dataexchange.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeStartJobResult struct {
	Result workflow.Future
}

func (r *DataexchangeStartJobResult) Get(ctx workflow.Context) (*dataexchange.StartJobOutput, error) {
    var output dataexchange.StartJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeTagResourceResult struct {
	Result workflow.Future
}

func (r *DataexchangeTagResourceResult) Get(ctx workflow.Context) (*dataexchange.TagResourceOutput, error) {
    var output dataexchange.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeUntagResourceResult struct {
	Result workflow.Future
}

func (r *DataexchangeUntagResourceResult) Get(ctx workflow.Context) (*dataexchange.UntagResourceOutput, error) {
    var output dataexchange.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeUpdateAssetResult struct {
	Result workflow.Future
}

func (r *DataexchangeUpdateAssetResult) Get(ctx workflow.Context) (*dataexchange.UpdateAssetOutput, error) {
    var output dataexchange.UpdateAssetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeUpdateDataSetResult struct {
	Result workflow.Future
}

func (r *DataexchangeUpdateDataSetResult) Get(ctx workflow.Context) (*dataexchange.UpdateDataSetOutput, error) {
    var output dataexchange.UpdateDataSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataexchangeUpdateRevisionResult struct {
	Result workflow.Future
}

func (r *DataexchangeUpdateRevisionResult) Get(ctx workflow.Context) (*dataexchange.UpdateRevisionOutput, error) {
    var output dataexchange.UpdateRevisionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DataExchangeStub struct {
    activities awsactivities.DataExchangeActivities
}

func NewDataExchangeStub() DataExchangeClient {
    return &DataExchangeStub{}
}

func (a *DataExchangeStub) CancelJob(ctx workflow.Context, input *dataexchange.CancelJobInput) (*dataexchange.CancelJobOutput, error) {
    var output dataexchange.CancelJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) CancelJobAsync(ctx workflow.Context, input *dataexchange.CancelJobInput) *DataexchangeCancelJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelJob, input)
    return &DataexchangeCancelJobResult{Result: future}
}

func (a *DataExchangeStub) CreateDataSet(ctx workflow.Context, input *dataexchange.CreateDataSetInput) (*dataexchange.CreateDataSetOutput, error) {
    var output dataexchange.CreateDataSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDataSet, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) CreateDataSetAsync(ctx workflow.Context, input *dataexchange.CreateDataSetInput) *DataexchangeCreateDataSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDataSet, input)
    return &DataexchangeCreateDataSetResult{Result: future}
}

func (a *DataExchangeStub) CreateJob(ctx workflow.Context, input *dataexchange.CreateJobInput) (*dataexchange.CreateJobOutput, error) {
    var output dataexchange.CreateJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) CreateJobAsync(ctx workflow.Context, input *dataexchange.CreateJobInput) *DataexchangeCreateJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input)
    return &DataexchangeCreateJobResult{Result: future}
}

func (a *DataExchangeStub) CreateRevision(ctx workflow.Context, input *dataexchange.CreateRevisionInput) (*dataexchange.CreateRevisionOutput, error) {
    var output dataexchange.CreateRevisionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRevision, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) CreateRevisionAsync(ctx workflow.Context, input *dataexchange.CreateRevisionInput) *DataexchangeCreateRevisionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateRevision, input)
    return &DataexchangeCreateRevisionResult{Result: future}
}

func (a *DataExchangeStub) DeleteAsset(ctx workflow.Context, input *dataexchange.DeleteAssetInput) (*dataexchange.DeleteAssetOutput, error) {
    var output dataexchange.DeleteAssetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteAsset, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) DeleteAssetAsync(ctx workflow.Context, input *dataexchange.DeleteAssetInput) *DataexchangeDeleteAssetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteAsset, input)
    return &DataexchangeDeleteAssetResult{Result: future}
}

func (a *DataExchangeStub) DeleteDataSet(ctx workflow.Context, input *dataexchange.DeleteDataSetInput) (*dataexchange.DeleteDataSetOutput, error) {
    var output dataexchange.DeleteDataSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDataSet, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) DeleteDataSetAsync(ctx workflow.Context, input *dataexchange.DeleteDataSetInput) *DataexchangeDeleteDataSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDataSet, input)
    return &DataexchangeDeleteDataSetResult{Result: future}
}

func (a *DataExchangeStub) DeleteRevision(ctx workflow.Context, input *dataexchange.DeleteRevisionInput) (*dataexchange.DeleteRevisionOutput, error) {
    var output dataexchange.DeleteRevisionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRevision, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) DeleteRevisionAsync(ctx workflow.Context, input *dataexchange.DeleteRevisionInput) *DataexchangeDeleteRevisionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRevision, input)
    return &DataexchangeDeleteRevisionResult{Result: future}
}

func (a *DataExchangeStub) GetAsset(ctx workflow.Context, input *dataexchange.GetAssetInput) (*dataexchange.GetAssetOutput, error) {
    var output dataexchange.GetAssetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAsset, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) GetAssetAsync(ctx workflow.Context, input *dataexchange.GetAssetInput) *DataexchangeGetAssetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAsset, input)
    return &DataexchangeGetAssetResult{Result: future}
}

func (a *DataExchangeStub) GetDataSet(ctx workflow.Context, input *dataexchange.GetDataSetInput) (*dataexchange.GetDataSetOutput, error) {
    var output dataexchange.GetDataSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDataSet, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) GetDataSetAsync(ctx workflow.Context, input *dataexchange.GetDataSetInput) *DataexchangeGetDataSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDataSet, input)
    return &DataexchangeGetDataSetResult{Result: future}
}

func (a *DataExchangeStub) GetJob(ctx workflow.Context, input *dataexchange.GetJobInput) (*dataexchange.GetJobOutput, error) {
    var output dataexchange.GetJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJob, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) GetJobAsync(ctx workflow.Context, input *dataexchange.GetJobInput) *DataexchangeGetJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJob, input)
    return &DataexchangeGetJobResult{Result: future}
}

func (a *DataExchangeStub) GetRevision(ctx workflow.Context, input *dataexchange.GetRevisionInput) (*dataexchange.GetRevisionOutput, error) {
    var output dataexchange.GetRevisionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRevision, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) GetRevisionAsync(ctx workflow.Context, input *dataexchange.GetRevisionInput) *DataexchangeGetRevisionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRevision, input)
    return &DataexchangeGetRevisionResult{Result: future}
}

func (a *DataExchangeStub) ListDataSetRevisions(ctx workflow.Context, input *dataexchange.ListDataSetRevisionsInput) (*dataexchange.ListDataSetRevisionsOutput, error) {
    var output dataexchange.ListDataSetRevisionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDataSetRevisions, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) ListDataSetRevisionsAsync(ctx workflow.Context, input *dataexchange.ListDataSetRevisionsInput) *DataexchangeListDataSetRevisionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDataSetRevisions, input)
    return &DataexchangeListDataSetRevisionsResult{Result: future}
}

func (a *DataExchangeStub) ListDataSets(ctx workflow.Context, input *dataexchange.ListDataSetsInput) (*dataexchange.ListDataSetsOutput, error) {
    var output dataexchange.ListDataSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDataSets, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) ListDataSetsAsync(ctx workflow.Context, input *dataexchange.ListDataSetsInput) *DataexchangeListDataSetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDataSets, input)
    return &DataexchangeListDataSetsResult{Result: future}
}

func (a *DataExchangeStub) ListJobs(ctx workflow.Context, input *dataexchange.ListJobsInput) (*dataexchange.ListJobsOutput, error) {
    var output dataexchange.ListJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) ListJobsAsync(ctx workflow.Context, input *dataexchange.ListJobsInput) *DataexchangeListJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input)
    return &DataexchangeListJobsResult{Result: future}
}

func (a *DataExchangeStub) ListRevisionAssets(ctx workflow.Context, input *dataexchange.ListRevisionAssetsInput) (*dataexchange.ListRevisionAssetsOutput, error) {
    var output dataexchange.ListRevisionAssetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRevisionAssets, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) ListRevisionAssetsAsync(ctx workflow.Context, input *dataexchange.ListRevisionAssetsInput) *DataexchangeListRevisionAssetsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRevisionAssets, input)
    return &DataexchangeListRevisionAssetsResult{Result: future}
}

func (a *DataExchangeStub) ListTagsForResource(ctx workflow.Context, input *dataexchange.ListTagsForResourceInput) (*dataexchange.ListTagsForResourceOutput, error) {
    var output dataexchange.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) ListTagsForResourceAsync(ctx workflow.Context, input *dataexchange.ListTagsForResourceInput) *DataexchangeListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &DataexchangeListTagsForResourceResult{Result: future}
}

func (a *DataExchangeStub) StartJob(ctx workflow.Context, input *dataexchange.StartJobInput) (*dataexchange.StartJobOutput, error) {
    var output dataexchange.StartJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartJob, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) StartJobAsync(ctx workflow.Context, input *dataexchange.StartJobInput) *DataexchangeStartJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartJob, input)
    return &DataexchangeStartJobResult{Result: future}
}

func (a *DataExchangeStub) TagResource(ctx workflow.Context, input *dataexchange.TagResourceInput) (*dataexchange.TagResourceOutput, error) {
    var output dataexchange.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) TagResourceAsync(ctx workflow.Context, input *dataexchange.TagResourceInput) *DataexchangeTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &DataexchangeTagResourceResult{Result: future}
}

func (a *DataExchangeStub) UntagResource(ctx workflow.Context, input *dataexchange.UntagResourceInput) (*dataexchange.UntagResourceOutput, error) {
    var output dataexchange.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) UntagResourceAsync(ctx workflow.Context, input *dataexchange.UntagResourceInput) *DataexchangeUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &DataexchangeUntagResourceResult{Result: future}
}

func (a *DataExchangeStub) UpdateAsset(ctx workflow.Context, input *dataexchange.UpdateAssetInput) (*dataexchange.UpdateAssetOutput, error) {
    var output dataexchange.UpdateAssetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAsset, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) UpdateAssetAsync(ctx workflow.Context, input *dataexchange.UpdateAssetInput) *DataexchangeUpdateAssetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateAsset, input)
    return &DataexchangeUpdateAssetResult{Result: future}
}

func (a *DataExchangeStub) UpdateDataSet(ctx workflow.Context, input *dataexchange.UpdateDataSetInput) (*dataexchange.UpdateDataSetOutput, error) {
    var output dataexchange.UpdateDataSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDataSet, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) UpdateDataSetAsync(ctx workflow.Context, input *dataexchange.UpdateDataSetInput) *DataexchangeUpdateDataSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDataSet, input)
    return &DataexchangeUpdateDataSetResult{Result: future}
}

func (a *DataExchangeStub) UpdateRevision(ctx workflow.Context, input *dataexchange.UpdateRevisionInput) (*dataexchange.UpdateRevisionOutput, error) {
    var output dataexchange.UpdateRevisionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRevision, input).Get(ctx, &output)
    return &output, err
}

func (a *DataExchangeStub) UpdateRevisionAsync(ctx workflow.Context, input *dataexchange.UpdateRevisionInput) *DataexchangeUpdateRevisionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateRevision, input)
    return &DataexchangeUpdateRevisionResult{Result: future}
}