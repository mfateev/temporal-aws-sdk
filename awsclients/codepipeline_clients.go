package awsclients

import (
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type CodePipelineClient interface {
       AcknowledgeJob(ctx workflow.Context, input *codepipeline.AcknowledgeJobInput) (*codepipeline.AcknowledgeJobOutput, error)
       AcknowledgeJobAsync(ctx workflow.Context, input *codepipeline.AcknowledgeJobInput) *CodepipelineAcknowledgeJobResult

       AcknowledgeThirdPartyJob(ctx workflow.Context, input *codepipeline.AcknowledgeThirdPartyJobInput) (*codepipeline.AcknowledgeThirdPartyJobOutput, error)
       AcknowledgeThirdPartyJobAsync(ctx workflow.Context, input *codepipeline.AcknowledgeThirdPartyJobInput) *CodepipelineAcknowledgeThirdPartyJobResult

       CreateCustomActionType(ctx workflow.Context, input *codepipeline.CreateCustomActionTypeInput) (*codepipeline.CreateCustomActionTypeOutput, error)
       CreateCustomActionTypeAsync(ctx workflow.Context, input *codepipeline.CreateCustomActionTypeInput) *CodepipelineCreateCustomActionTypeResult

       CreatePipeline(ctx workflow.Context, input *codepipeline.CreatePipelineInput) (*codepipeline.CreatePipelineOutput, error)
       CreatePipelineAsync(ctx workflow.Context, input *codepipeline.CreatePipelineInput) *CodepipelineCreatePipelineResult

       DeleteCustomActionType(ctx workflow.Context, input *codepipeline.DeleteCustomActionTypeInput) (*codepipeline.DeleteCustomActionTypeOutput, error)
       DeleteCustomActionTypeAsync(ctx workflow.Context, input *codepipeline.DeleteCustomActionTypeInput) *CodepipelineDeleteCustomActionTypeResult

       DeletePipeline(ctx workflow.Context, input *codepipeline.DeletePipelineInput) (*codepipeline.DeletePipelineOutput, error)
       DeletePipelineAsync(ctx workflow.Context, input *codepipeline.DeletePipelineInput) *CodepipelineDeletePipelineResult

       DeleteWebhook(ctx workflow.Context, input *codepipeline.DeleteWebhookInput) (*codepipeline.DeleteWebhookOutput, error)
       DeleteWebhookAsync(ctx workflow.Context, input *codepipeline.DeleteWebhookInput) *CodepipelineDeleteWebhookResult

       DeregisterWebhookWithThirdParty(ctx workflow.Context, input *codepipeline.DeregisterWebhookWithThirdPartyInput) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error)
       DeregisterWebhookWithThirdPartyAsync(ctx workflow.Context, input *codepipeline.DeregisterWebhookWithThirdPartyInput) *CodepipelineDeregisterWebhookWithThirdPartyResult

       DisableStageTransition(ctx workflow.Context, input *codepipeline.DisableStageTransitionInput) (*codepipeline.DisableStageTransitionOutput, error)
       DisableStageTransitionAsync(ctx workflow.Context, input *codepipeline.DisableStageTransitionInput) *CodepipelineDisableStageTransitionResult

       EnableStageTransition(ctx workflow.Context, input *codepipeline.EnableStageTransitionInput) (*codepipeline.EnableStageTransitionOutput, error)
       EnableStageTransitionAsync(ctx workflow.Context, input *codepipeline.EnableStageTransitionInput) *CodepipelineEnableStageTransitionResult

       GetJobDetails(ctx workflow.Context, input *codepipeline.GetJobDetailsInput) (*codepipeline.GetJobDetailsOutput, error)
       GetJobDetailsAsync(ctx workflow.Context, input *codepipeline.GetJobDetailsInput) *CodepipelineGetJobDetailsResult

       GetPipeline(ctx workflow.Context, input *codepipeline.GetPipelineInput) (*codepipeline.GetPipelineOutput, error)
       GetPipelineAsync(ctx workflow.Context, input *codepipeline.GetPipelineInput) *CodepipelineGetPipelineResult

       GetPipelineExecution(ctx workflow.Context, input *codepipeline.GetPipelineExecutionInput) (*codepipeline.GetPipelineExecutionOutput, error)
       GetPipelineExecutionAsync(ctx workflow.Context, input *codepipeline.GetPipelineExecutionInput) *CodepipelineGetPipelineExecutionResult

       GetPipelineState(ctx workflow.Context, input *codepipeline.GetPipelineStateInput) (*codepipeline.GetPipelineStateOutput, error)
       GetPipelineStateAsync(ctx workflow.Context, input *codepipeline.GetPipelineStateInput) *CodepipelineGetPipelineStateResult

       GetThirdPartyJobDetails(ctx workflow.Context, input *codepipeline.GetThirdPartyJobDetailsInput) (*codepipeline.GetThirdPartyJobDetailsOutput, error)
       GetThirdPartyJobDetailsAsync(ctx workflow.Context, input *codepipeline.GetThirdPartyJobDetailsInput) *CodepipelineGetThirdPartyJobDetailsResult

       ListActionExecutions(ctx workflow.Context, input *codepipeline.ListActionExecutionsInput) (*codepipeline.ListActionExecutionsOutput, error)
       ListActionExecutionsAsync(ctx workflow.Context, input *codepipeline.ListActionExecutionsInput) *CodepipelineListActionExecutionsResult

       ListActionTypes(ctx workflow.Context, input *codepipeline.ListActionTypesInput) (*codepipeline.ListActionTypesOutput, error)
       ListActionTypesAsync(ctx workflow.Context, input *codepipeline.ListActionTypesInput) *CodepipelineListActionTypesResult

       ListPipelineExecutions(ctx workflow.Context, input *codepipeline.ListPipelineExecutionsInput) (*codepipeline.ListPipelineExecutionsOutput, error)
       ListPipelineExecutionsAsync(ctx workflow.Context, input *codepipeline.ListPipelineExecutionsInput) *CodepipelineListPipelineExecutionsResult

       ListPipelines(ctx workflow.Context, input *codepipeline.ListPipelinesInput) (*codepipeline.ListPipelinesOutput, error)
       ListPipelinesAsync(ctx workflow.Context, input *codepipeline.ListPipelinesInput) *CodepipelineListPipelinesResult

       ListTagsForResource(ctx workflow.Context, input *codepipeline.ListTagsForResourceInput) (*codepipeline.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *codepipeline.ListTagsForResourceInput) *CodepipelineListTagsForResourceResult

       ListWebhooks(ctx workflow.Context, input *codepipeline.ListWebhooksInput) (*codepipeline.ListWebhooksOutput, error)
       ListWebhooksAsync(ctx workflow.Context, input *codepipeline.ListWebhooksInput) *CodepipelineListWebhooksResult

       PollForJobs(ctx workflow.Context, input *codepipeline.PollForJobsInput) (*codepipeline.PollForJobsOutput, error)
       PollForJobsAsync(ctx workflow.Context, input *codepipeline.PollForJobsInput) *CodepipelinePollForJobsResult

       PollForThirdPartyJobs(ctx workflow.Context, input *codepipeline.PollForThirdPartyJobsInput) (*codepipeline.PollForThirdPartyJobsOutput, error)
       PollForThirdPartyJobsAsync(ctx workflow.Context, input *codepipeline.PollForThirdPartyJobsInput) *CodepipelinePollForThirdPartyJobsResult

       PutActionRevision(ctx workflow.Context, input *codepipeline.PutActionRevisionInput) (*codepipeline.PutActionRevisionOutput, error)
       PutActionRevisionAsync(ctx workflow.Context, input *codepipeline.PutActionRevisionInput) *CodepipelinePutActionRevisionResult

       PutApprovalResult(ctx workflow.Context, input *codepipeline.PutApprovalResultInput) (*codepipeline.PutApprovalResultOutput, error)
       PutApprovalResultAsync(ctx workflow.Context, input *codepipeline.PutApprovalResultInput) *CodepipelinePutApprovalResultResult

       PutJobFailureResult(ctx workflow.Context, input *codepipeline.PutJobFailureResultInput) (*codepipeline.PutJobFailureResultOutput, error)
       PutJobFailureResultAsync(ctx workflow.Context, input *codepipeline.PutJobFailureResultInput) *CodepipelinePutJobFailureResultResult

       PutJobSuccessResult(ctx workflow.Context, input *codepipeline.PutJobSuccessResultInput) (*codepipeline.PutJobSuccessResultOutput, error)
       PutJobSuccessResultAsync(ctx workflow.Context, input *codepipeline.PutJobSuccessResultInput) *CodepipelinePutJobSuccessResultResult

       PutThirdPartyJobFailureResult(ctx workflow.Context, input *codepipeline.PutThirdPartyJobFailureResultInput) (*codepipeline.PutThirdPartyJobFailureResultOutput, error)
       PutThirdPartyJobFailureResultAsync(ctx workflow.Context, input *codepipeline.PutThirdPartyJobFailureResultInput) *CodepipelinePutThirdPartyJobFailureResultResult

       PutThirdPartyJobSuccessResult(ctx workflow.Context, input *codepipeline.PutThirdPartyJobSuccessResultInput) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error)
       PutThirdPartyJobSuccessResultAsync(ctx workflow.Context, input *codepipeline.PutThirdPartyJobSuccessResultInput) *CodepipelinePutThirdPartyJobSuccessResultResult

       PutWebhook(ctx workflow.Context, input *codepipeline.PutWebhookInput) (*codepipeline.PutWebhookOutput, error)
       PutWebhookAsync(ctx workflow.Context, input *codepipeline.PutWebhookInput) *CodepipelinePutWebhookResult

       RegisterWebhookWithThirdParty(ctx workflow.Context, input *codepipeline.RegisterWebhookWithThirdPartyInput) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error)
       RegisterWebhookWithThirdPartyAsync(ctx workflow.Context, input *codepipeline.RegisterWebhookWithThirdPartyInput) *CodepipelineRegisterWebhookWithThirdPartyResult

       RetryStageExecution(ctx workflow.Context, input *codepipeline.RetryStageExecutionInput) (*codepipeline.RetryStageExecutionOutput, error)
       RetryStageExecutionAsync(ctx workflow.Context, input *codepipeline.RetryStageExecutionInput) *CodepipelineRetryStageExecutionResult

       StartPipelineExecution(ctx workflow.Context, input *codepipeline.StartPipelineExecutionInput) (*codepipeline.StartPipelineExecutionOutput, error)
       StartPipelineExecutionAsync(ctx workflow.Context, input *codepipeline.StartPipelineExecutionInput) *CodepipelineStartPipelineExecutionResult

       StopPipelineExecution(ctx workflow.Context, input *codepipeline.StopPipelineExecutionInput) (*codepipeline.StopPipelineExecutionOutput, error)
       StopPipelineExecutionAsync(ctx workflow.Context, input *codepipeline.StopPipelineExecutionInput) *CodepipelineStopPipelineExecutionResult

       TagResource(ctx workflow.Context, input *codepipeline.TagResourceInput) (*codepipeline.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *codepipeline.TagResourceInput) *CodepipelineTagResourceResult

       UntagResource(ctx workflow.Context, input *codepipeline.UntagResourceInput) (*codepipeline.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *codepipeline.UntagResourceInput) *CodepipelineUntagResourceResult

       UpdatePipeline(ctx workflow.Context, input *codepipeline.UpdatePipelineInput) (*codepipeline.UpdatePipelineOutput, error)
       UpdatePipelineAsync(ctx workflow.Context, input *codepipeline.UpdatePipelineInput) *CodepipelineUpdatePipelineResult
}

type CodepipelineAcknowledgeJobResult struct {
	Result workflow.Future
}

func (r *CodepipelineAcknowledgeJobResult) Get(ctx workflow.Context) (*codepipeline.AcknowledgeJobOutput, error) {
    var output codepipeline.AcknowledgeJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineAcknowledgeThirdPartyJobResult struct {
	Result workflow.Future
}

func (r *CodepipelineAcknowledgeThirdPartyJobResult) Get(ctx workflow.Context) (*codepipeline.AcknowledgeThirdPartyJobOutput, error) {
    var output codepipeline.AcknowledgeThirdPartyJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineCreateCustomActionTypeResult struct {
	Result workflow.Future
}

func (r *CodepipelineCreateCustomActionTypeResult) Get(ctx workflow.Context) (*codepipeline.CreateCustomActionTypeOutput, error) {
    var output codepipeline.CreateCustomActionTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineCreatePipelineResult struct {
	Result workflow.Future
}

func (r *CodepipelineCreatePipelineResult) Get(ctx workflow.Context) (*codepipeline.CreatePipelineOutput, error) {
    var output codepipeline.CreatePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineDeleteCustomActionTypeResult struct {
	Result workflow.Future
}

func (r *CodepipelineDeleteCustomActionTypeResult) Get(ctx workflow.Context) (*codepipeline.DeleteCustomActionTypeOutput, error) {
    var output codepipeline.DeleteCustomActionTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineDeletePipelineResult struct {
	Result workflow.Future
}

func (r *CodepipelineDeletePipelineResult) Get(ctx workflow.Context) (*codepipeline.DeletePipelineOutput, error) {
    var output codepipeline.DeletePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineDeleteWebhookResult struct {
	Result workflow.Future
}

func (r *CodepipelineDeleteWebhookResult) Get(ctx workflow.Context) (*codepipeline.DeleteWebhookOutput, error) {
    var output codepipeline.DeleteWebhookOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineDeregisterWebhookWithThirdPartyResult struct {
	Result workflow.Future
}

func (r *CodepipelineDeregisterWebhookWithThirdPartyResult) Get(ctx workflow.Context) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error) {
    var output codepipeline.DeregisterWebhookWithThirdPartyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineDisableStageTransitionResult struct {
	Result workflow.Future
}

func (r *CodepipelineDisableStageTransitionResult) Get(ctx workflow.Context) (*codepipeline.DisableStageTransitionOutput, error) {
    var output codepipeline.DisableStageTransitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineEnableStageTransitionResult struct {
	Result workflow.Future
}

func (r *CodepipelineEnableStageTransitionResult) Get(ctx workflow.Context) (*codepipeline.EnableStageTransitionOutput, error) {
    var output codepipeline.EnableStageTransitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineGetJobDetailsResult struct {
	Result workflow.Future
}

func (r *CodepipelineGetJobDetailsResult) Get(ctx workflow.Context) (*codepipeline.GetJobDetailsOutput, error) {
    var output codepipeline.GetJobDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineGetPipelineResult struct {
	Result workflow.Future
}

func (r *CodepipelineGetPipelineResult) Get(ctx workflow.Context) (*codepipeline.GetPipelineOutput, error) {
    var output codepipeline.GetPipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineGetPipelineExecutionResult struct {
	Result workflow.Future
}

func (r *CodepipelineGetPipelineExecutionResult) Get(ctx workflow.Context) (*codepipeline.GetPipelineExecutionOutput, error) {
    var output codepipeline.GetPipelineExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineGetPipelineStateResult struct {
	Result workflow.Future
}

func (r *CodepipelineGetPipelineStateResult) Get(ctx workflow.Context) (*codepipeline.GetPipelineStateOutput, error) {
    var output codepipeline.GetPipelineStateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineGetThirdPartyJobDetailsResult struct {
	Result workflow.Future
}

func (r *CodepipelineGetThirdPartyJobDetailsResult) Get(ctx workflow.Context) (*codepipeline.GetThirdPartyJobDetailsOutput, error) {
    var output codepipeline.GetThirdPartyJobDetailsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineListActionExecutionsResult struct {
	Result workflow.Future
}

func (r *CodepipelineListActionExecutionsResult) Get(ctx workflow.Context) (*codepipeline.ListActionExecutionsOutput, error) {
    var output codepipeline.ListActionExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineListActionTypesResult struct {
	Result workflow.Future
}

func (r *CodepipelineListActionTypesResult) Get(ctx workflow.Context) (*codepipeline.ListActionTypesOutput, error) {
    var output codepipeline.ListActionTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineListPipelineExecutionsResult struct {
	Result workflow.Future
}

func (r *CodepipelineListPipelineExecutionsResult) Get(ctx workflow.Context) (*codepipeline.ListPipelineExecutionsOutput, error) {
    var output codepipeline.ListPipelineExecutionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineListPipelinesResult struct {
	Result workflow.Future
}

func (r *CodepipelineListPipelinesResult) Get(ctx workflow.Context) (*codepipeline.ListPipelinesOutput, error) {
    var output codepipeline.ListPipelinesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *CodepipelineListTagsForResourceResult) Get(ctx workflow.Context) (*codepipeline.ListTagsForResourceOutput, error) {
    var output codepipeline.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineListWebhooksResult struct {
	Result workflow.Future
}

func (r *CodepipelineListWebhooksResult) Get(ctx workflow.Context) (*codepipeline.ListWebhooksOutput, error) {
    var output codepipeline.ListWebhooksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelinePollForJobsResult struct {
	Result workflow.Future
}

func (r *CodepipelinePollForJobsResult) Get(ctx workflow.Context) (*codepipeline.PollForJobsOutput, error) {
    var output codepipeline.PollForJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelinePollForThirdPartyJobsResult struct {
	Result workflow.Future
}

func (r *CodepipelinePollForThirdPartyJobsResult) Get(ctx workflow.Context) (*codepipeline.PollForThirdPartyJobsOutput, error) {
    var output codepipeline.PollForThirdPartyJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelinePutActionRevisionResult struct {
	Result workflow.Future
}

func (r *CodepipelinePutActionRevisionResult) Get(ctx workflow.Context) (*codepipeline.PutActionRevisionOutput, error) {
    var output codepipeline.PutActionRevisionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelinePutApprovalResultResult struct {
	Result workflow.Future
}

func (r *CodepipelinePutApprovalResultResult) Get(ctx workflow.Context) (*codepipeline.PutApprovalResultOutput, error) {
    var output codepipeline.PutApprovalResultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelinePutJobFailureResultResult struct {
	Result workflow.Future
}

func (r *CodepipelinePutJobFailureResultResult) Get(ctx workflow.Context) (*codepipeline.PutJobFailureResultOutput, error) {
    var output codepipeline.PutJobFailureResultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelinePutJobSuccessResultResult struct {
	Result workflow.Future
}

func (r *CodepipelinePutJobSuccessResultResult) Get(ctx workflow.Context) (*codepipeline.PutJobSuccessResultOutput, error) {
    var output codepipeline.PutJobSuccessResultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelinePutThirdPartyJobFailureResultResult struct {
	Result workflow.Future
}

func (r *CodepipelinePutThirdPartyJobFailureResultResult) Get(ctx workflow.Context) (*codepipeline.PutThirdPartyJobFailureResultOutput, error) {
    var output codepipeline.PutThirdPartyJobFailureResultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelinePutThirdPartyJobSuccessResultResult struct {
	Result workflow.Future
}

func (r *CodepipelinePutThirdPartyJobSuccessResultResult) Get(ctx workflow.Context) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error) {
    var output codepipeline.PutThirdPartyJobSuccessResultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelinePutWebhookResult struct {
	Result workflow.Future
}

func (r *CodepipelinePutWebhookResult) Get(ctx workflow.Context) (*codepipeline.PutWebhookOutput, error) {
    var output codepipeline.PutWebhookOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineRegisterWebhookWithThirdPartyResult struct {
	Result workflow.Future
}

func (r *CodepipelineRegisterWebhookWithThirdPartyResult) Get(ctx workflow.Context) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error) {
    var output codepipeline.RegisterWebhookWithThirdPartyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineRetryStageExecutionResult struct {
	Result workflow.Future
}

func (r *CodepipelineRetryStageExecutionResult) Get(ctx workflow.Context) (*codepipeline.RetryStageExecutionOutput, error) {
    var output codepipeline.RetryStageExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineStartPipelineExecutionResult struct {
	Result workflow.Future
}

func (r *CodepipelineStartPipelineExecutionResult) Get(ctx workflow.Context) (*codepipeline.StartPipelineExecutionOutput, error) {
    var output codepipeline.StartPipelineExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineStopPipelineExecutionResult struct {
	Result workflow.Future
}

func (r *CodepipelineStopPipelineExecutionResult) Get(ctx workflow.Context) (*codepipeline.StopPipelineExecutionOutput, error) {
    var output codepipeline.StopPipelineExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineTagResourceResult struct {
	Result workflow.Future
}

func (r *CodepipelineTagResourceResult) Get(ctx workflow.Context) (*codepipeline.TagResourceOutput, error) {
    var output codepipeline.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineUntagResourceResult struct {
	Result workflow.Future
}

func (r *CodepipelineUntagResourceResult) Get(ctx workflow.Context) (*codepipeline.UntagResourceOutput, error) {
    var output codepipeline.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodepipelineUpdatePipelineResult struct {
	Result workflow.Future
}

func (r *CodepipelineUpdatePipelineResult) Get(ctx workflow.Context) (*codepipeline.UpdatePipelineOutput, error) {
    var output codepipeline.UpdatePipelineOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodePipelineStub struct {
    activities awsactivities.CodePipelineActivities
}

func NewCodePipelineStub() CodePipelineClient {
    return &CodePipelineStub{}
}

func (a *CodePipelineStub) AcknowledgeJob(ctx workflow.Context, input *codepipeline.AcknowledgeJobInput) (*codepipeline.AcknowledgeJobOutput, error) {
    var output codepipeline.AcknowledgeJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcknowledgeJob, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) AcknowledgeJobAsync(ctx workflow.Context, input *codepipeline.AcknowledgeJobInput) *CodepipelineAcknowledgeJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcknowledgeJob, input)
    return &CodepipelineAcknowledgeJobResult{Result: future}
}

func (a *CodePipelineStub) AcknowledgeThirdPartyJob(ctx workflow.Context, input *codepipeline.AcknowledgeThirdPartyJobInput) (*codepipeline.AcknowledgeThirdPartyJobOutput, error) {
    var output codepipeline.AcknowledgeThirdPartyJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcknowledgeThirdPartyJob, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) AcknowledgeThirdPartyJobAsync(ctx workflow.Context, input *codepipeline.AcknowledgeThirdPartyJobInput) *CodepipelineAcknowledgeThirdPartyJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcknowledgeThirdPartyJob, input)
    return &CodepipelineAcknowledgeThirdPartyJobResult{Result: future}
}

func (a *CodePipelineStub) CreateCustomActionType(ctx workflow.Context, input *codepipeline.CreateCustomActionTypeInput) (*codepipeline.CreateCustomActionTypeOutput, error) {
    var output codepipeline.CreateCustomActionTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCustomActionType, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) CreateCustomActionTypeAsync(ctx workflow.Context, input *codepipeline.CreateCustomActionTypeInput) *CodepipelineCreateCustomActionTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCustomActionType, input)
    return &CodepipelineCreateCustomActionTypeResult{Result: future}
}

func (a *CodePipelineStub) CreatePipeline(ctx workflow.Context, input *codepipeline.CreatePipelineInput) (*codepipeline.CreatePipelineOutput, error) {
    var output codepipeline.CreatePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) CreatePipelineAsync(ctx workflow.Context, input *codepipeline.CreatePipelineInput) *CodepipelineCreatePipelineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePipeline, input)
    return &CodepipelineCreatePipelineResult{Result: future}
}

func (a *CodePipelineStub) DeleteCustomActionType(ctx workflow.Context, input *codepipeline.DeleteCustomActionTypeInput) (*codepipeline.DeleteCustomActionTypeOutput, error) {
    var output codepipeline.DeleteCustomActionTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCustomActionType, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) DeleteCustomActionTypeAsync(ctx workflow.Context, input *codepipeline.DeleteCustomActionTypeInput) *CodepipelineDeleteCustomActionTypeResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCustomActionType, input)
    return &CodepipelineDeleteCustomActionTypeResult{Result: future}
}

func (a *CodePipelineStub) DeletePipeline(ctx workflow.Context, input *codepipeline.DeletePipelineInput) (*codepipeline.DeletePipelineOutput, error) {
    var output codepipeline.DeletePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) DeletePipelineAsync(ctx workflow.Context, input *codepipeline.DeletePipelineInput) *CodepipelineDeletePipelineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePipeline, input)
    return &CodepipelineDeletePipelineResult{Result: future}
}

func (a *CodePipelineStub) DeleteWebhook(ctx workflow.Context, input *codepipeline.DeleteWebhookInput) (*codepipeline.DeleteWebhookOutput, error) {
    var output codepipeline.DeleteWebhookOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteWebhook, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) DeleteWebhookAsync(ctx workflow.Context, input *codepipeline.DeleteWebhookInput) *CodepipelineDeleteWebhookResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteWebhook, input)
    return &CodepipelineDeleteWebhookResult{Result: future}
}

func (a *CodePipelineStub) DeregisterWebhookWithThirdParty(ctx workflow.Context, input *codepipeline.DeregisterWebhookWithThirdPartyInput) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error) {
    var output codepipeline.DeregisterWebhookWithThirdPartyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterWebhookWithThirdParty, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) DeregisterWebhookWithThirdPartyAsync(ctx workflow.Context, input *codepipeline.DeregisterWebhookWithThirdPartyInput) *CodepipelineDeregisterWebhookWithThirdPartyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeregisterWebhookWithThirdParty, input)
    return &CodepipelineDeregisterWebhookWithThirdPartyResult{Result: future}
}

func (a *CodePipelineStub) DisableStageTransition(ctx workflow.Context, input *codepipeline.DisableStageTransitionInput) (*codepipeline.DisableStageTransitionOutput, error) {
    var output codepipeline.DisableStageTransitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableStageTransition, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) DisableStageTransitionAsync(ctx workflow.Context, input *codepipeline.DisableStageTransitionInput) *CodepipelineDisableStageTransitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableStageTransition, input)
    return &CodepipelineDisableStageTransitionResult{Result: future}
}

func (a *CodePipelineStub) EnableStageTransition(ctx workflow.Context, input *codepipeline.EnableStageTransitionInput) (*codepipeline.EnableStageTransitionOutput, error) {
    var output codepipeline.EnableStageTransitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableStageTransition, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) EnableStageTransitionAsync(ctx workflow.Context, input *codepipeline.EnableStageTransitionInput) *CodepipelineEnableStageTransitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableStageTransition, input)
    return &CodepipelineEnableStageTransitionResult{Result: future}
}

func (a *CodePipelineStub) GetJobDetails(ctx workflow.Context, input *codepipeline.GetJobDetailsInput) (*codepipeline.GetJobDetailsOutput, error) {
    var output codepipeline.GetJobDetailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJobDetails, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) GetJobDetailsAsync(ctx workflow.Context, input *codepipeline.GetJobDetailsInput) *CodepipelineGetJobDetailsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJobDetails, input)
    return &CodepipelineGetJobDetailsResult{Result: future}
}

func (a *CodePipelineStub) GetPipeline(ctx workflow.Context, input *codepipeline.GetPipelineInput) (*codepipeline.GetPipelineOutput, error) {
    var output codepipeline.GetPipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) GetPipelineAsync(ctx workflow.Context, input *codepipeline.GetPipelineInput) *CodepipelineGetPipelineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPipeline, input)
    return &CodepipelineGetPipelineResult{Result: future}
}

func (a *CodePipelineStub) GetPipelineExecution(ctx workflow.Context, input *codepipeline.GetPipelineExecutionInput) (*codepipeline.GetPipelineExecutionOutput, error) {
    var output codepipeline.GetPipelineExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPipelineExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) GetPipelineExecutionAsync(ctx workflow.Context, input *codepipeline.GetPipelineExecutionInput) *CodepipelineGetPipelineExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPipelineExecution, input)
    return &CodepipelineGetPipelineExecutionResult{Result: future}
}

func (a *CodePipelineStub) GetPipelineState(ctx workflow.Context, input *codepipeline.GetPipelineStateInput) (*codepipeline.GetPipelineStateOutput, error) {
    var output codepipeline.GetPipelineStateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPipelineState, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) GetPipelineStateAsync(ctx workflow.Context, input *codepipeline.GetPipelineStateInput) *CodepipelineGetPipelineStateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPipelineState, input)
    return &CodepipelineGetPipelineStateResult{Result: future}
}

func (a *CodePipelineStub) GetThirdPartyJobDetails(ctx workflow.Context, input *codepipeline.GetThirdPartyJobDetailsInput) (*codepipeline.GetThirdPartyJobDetailsOutput, error) {
    var output codepipeline.GetThirdPartyJobDetailsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetThirdPartyJobDetails, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) GetThirdPartyJobDetailsAsync(ctx workflow.Context, input *codepipeline.GetThirdPartyJobDetailsInput) *CodepipelineGetThirdPartyJobDetailsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetThirdPartyJobDetails, input)
    return &CodepipelineGetThirdPartyJobDetailsResult{Result: future}
}

func (a *CodePipelineStub) ListActionExecutions(ctx workflow.Context, input *codepipeline.ListActionExecutionsInput) (*codepipeline.ListActionExecutionsOutput, error) {
    var output codepipeline.ListActionExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListActionExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) ListActionExecutionsAsync(ctx workflow.Context, input *codepipeline.ListActionExecutionsInput) *CodepipelineListActionExecutionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListActionExecutions, input)
    return &CodepipelineListActionExecutionsResult{Result: future}
}

func (a *CodePipelineStub) ListActionTypes(ctx workflow.Context, input *codepipeline.ListActionTypesInput) (*codepipeline.ListActionTypesOutput, error) {
    var output codepipeline.ListActionTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListActionTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) ListActionTypesAsync(ctx workflow.Context, input *codepipeline.ListActionTypesInput) *CodepipelineListActionTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListActionTypes, input)
    return &CodepipelineListActionTypesResult{Result: future}
}

func (a *CodePipelineStub) ListPipelineExecutions(ctx workflow.Context, input *codepipeline.ListPipelineExecutionsInput) (*codepipeline.ListPipelineExecutionsOutput, error) {
    var output codepipeline.ListPipelineExecutionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPipelineExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) ListPipelineExecutionsAsync(ctx workflow.Context, input *codepipeline.ListPipelineExecutionsInput) *CodepipelineListPipelineExecutionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPipelineExecutions, input)
    return &CodepipelineListPipelineExecutionsResult{Result: future}
}

func (a *CodePipelineStub) ListPipelines(ctx workflow.Context, input *codepipeline.ListPipelinesInput) (*codepipeline.ListPipelinesOutput, error) {
    var output codepipeline.ListPipelinesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPipelines, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) ListPipelinesAsync(ctx workflow.Context, input *codepipeline.ListPipelinesInput) *CodepipelineListPipelinesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPipelines, input)
    return &CodepipelineListPipelinesResult{Result: future}
}

func (a *CodePipelineStub) ListTagsForResource(ctx workflow.Context, input *codepipeline.ListTagsForResourceInput) (*codepipeline.ListTagsForResourceOutput, error) {
    var output codepipeline.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) ListTagsForResourceAsync(ctx workflow.Context, input *codepipeline.ListTagsForResourceInput) *CodepipelineListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &CodepipelineListTagsForResourceResult{Result: future}
}

func (a *CodePipelineStub) ListWebhooks(ctx workflow.Context, input *codepipeline.ListWebhooksInput) (*codepipeline.ListWebhooksOutput, error) {
    var output codepipeline.ListWebhooksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWebhooks, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) ListWebhooksAsync(ctx workflow.Context, input *codepipeline.ListWebhooksInput) *CodepipelineListWebhooksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListWebhooks, input)
    return &CodepipelineListWebhooksResult{Result: future}
}

func (a *CodePipelineStub) PollForJobs(ctx workflow.Context, input *codepipeline.PollForJobsInput) (*codepipeline.PollForJobsOutput, error) {
    var output codepipeline.PollForJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PollForJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) PollForJobsAsync(ctx workflow.Context, input *codepipeline.PollForJobsInput) *CodepipelinePollForJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PollForJobs, input)
    return &CodepipelinePollForJobsResult{Result: future}
}

func (a *CodePipelineStub) PollForThirdPartyJobs(ctx workflow.Context, input *codepipeline.PollForThirdPartyJobsInput) (*codepipeline.PollForThirdPartyJobsOutput, error) {
    var output codepipeline.PollForThirdPartyJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PollForThirdPartyJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) PollForThirdPartyJobsAsync(ctx workflow.Context, input *codepipeline.PollForThirdPartyJobsInput) *CodepipelinePollForThirdPartyJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PollForThirdPartyJobs, input)
    return &CodepipelinePollForThirdPartyJobsResult{Result: future}
}

func (a *CodePipelineStub) PutActionRevision(ctx workflow.Context, input *codepipeline.PutActionRevisionInput) (*codepipeline.PutActionRevisionOutput, error) {
    var output codepipeline.PutActionRevisionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutActionRevision, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) PutActionRevisionAsync(ctx workflow.Context, input *codepipeline.PutActionRevisionInput) *CodepipelinePutActionRevisionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutActionRevision, input)
    return &CodepipelinePutActionRevisionResult{Result: future}
}

func (a *CodePipelineStub) PutApprovalResult(ctx workflow.Context, input *codepipeline.PutApprovalResultInput) (*codepipeline.PutApprovalResultOutput, error) {
    var output codepipeline.PutApprovalResultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutApprovalResult, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) PutApprovalResultAsync(ctx workflow.Context, input *codepipeline.PutApprovalResultInput) *CodepipelinePutApprovalResultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutApprovalResult, input)
    return &CodepipelinePutApprovalResultResult{Result: future}
}

func (a *CodePipelineStub) PutJobFailureResult(ctx workflow.Context, input *codepipeline.PutJobFailureResultInput) (*codepipeline.PutJobFailureResultOutput, error) {
    var output codepipeline.PutJobFailureResultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutJobFailureResult, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) PutJobFailureResultAsync(ctx workflow.Context, input *codepipeline.PutJobFailureResultInput) *CodepipelinePutJobFailureResultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutJobFailureResult, input)
    return &CodepipelinePutJobFailureResultResult{Result: future}
}

func (a *CodePipelineStub) PutJobSuccessResult(ctx workflow.Context, input *codepipeline.PutJobSuccessResultInput) (*codepipeline.PutJobSuccessResultOutput, error) {
    var output codepipeline.PutJobSuccessResultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutJobSuccessResult, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) PutJobSuccessResultAsync(ctx workflow.Context, input *codepipeline.PutJobSuccessResultInput) *CodepipelinePutJobSuccessResultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutJobSuccessResult, input)
    return &CodepipelinePutJobSuccessResultResult{Result: future}
}

func (a *CodePipelineStub) PutThirdPartyJobFailureResult(ctx workflow.Context, input *codepipeline.PutThirdPartyJobFailureResultInput) (*codepipeline.PutThirdPartyJobFailureResultOutput, error) {
    var output codepipeline.PutThirdPartyJobFailureResultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutThirdPartyJobFailureResult, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) PutThirdPartyJobFailureResultAsync(ctx workflow.Context, input *codepipeline.PutThirdPartyJobFailureResultInput) *CodepipelinePutThirdPartyJobFailureResultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutThirdPartyJobFailureResult, input)
    return &CodepipelinePutThirdPartyJobFailureResultResult{Result: future}
}

func (a *CodePipelineStub) PutThirdPartyJobSuccessResult(ctx workflow.Context, input *codepipeline.PutThirdPartyJobSuccessResultInput) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error) {
    var output codepipeline.PutThirdPartyJobSuccessResultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutThirdPartyJobSuccessResult, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) PutThirdPartyJobSuccessResultAsync(ctx workflow.Context, input *codepipeline.PutThirdPartyJobSuccessResultInput) *CodepipelinePutThirdPartyJobSuccessResultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutThirdPartyJobSuccessResult, input)
    return &CodepipelinePutThirdPartyJobSuccessResultResult{Result: future}
}

func (a *CodePipelineStub) PutWebhook(ctx workflow.Context, input *codepipeline.PutWebhookInput) (*codepipeline.PutWebhookOutput, error) {
    var output codepipeline.PutWebhookOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutWebhook, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) PutWebhookAsync(ctx workflow.Context, input *codepipeline.PutWebhookInput) *CodepipelinePutWebhookResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutWebhook, input)
    return &CodepipelinePutWebhookResult{Result: future}
}

func (a *CodePipelineStub) RegisterWebhookWithThirdParty(ctx workflow.Context, input *codepipeline.RegisterWebhookWithThirdPartyInput) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error) {
    var output codepipeline.RegisterWebhookWithThirdPartyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterWebhookWithThirdParty, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) RegisterWebhookWithThirdPartyAsync(ctx workflow.Context, input *codepipeline.RegisterWebhookWithThirdPartyInput) *CodepipelineRegisterWebhookWithThirdPartyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RegisterWebhookWithThirdParty, input)
    return &CodepipelineRegisterWebhookWithThirdPartyResult{Result: future}
}

func (a *CodePipelineStub) RetryStageExecution(ctx workflow.Context, input *codepipeline.RetryStageExecutionInput) (*codepipeline.RetryStageExecutionOutput, error) {
    var output codepipeline.RetryStageExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RetryStageExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) RetryStageExecutionAsync(ctx workflow.Context, input *codepipeline.RetryStageExecutionInput) *CodepipelineRetryStageExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RetryStageExecution, input)
    return &CodepipelineRetryStageExecutionResult{Result: future}
}

func (a *CodePipelineStub) StartPipelineExecution(ctx workflow.Context, input *codepipeline.StartPipelineExecutionInput) (*codepipeline.StartPipelineExecutionOutput, error) {
    var output codepipeline.StartPipelineExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartPipelineExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) StartPipelineExecutionAsync(ctx workflow.Context, input *codepipeline.StartPipelineExecutionInput) *CodepipelineStartPipelineExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartPipelineExecution, input)
    return &CodepipelineStartPipelineExecutionResult{Result: future}
}

func (a *CodePipelineStub) StopPipelineExecution(ctx workflow.Context, input *codepipeline.StopPipelineExecutionInput) (*codepipeline.StopPipelineExecutionOutput, error) {
    var output codepipeline.StopPipelineExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopPipelineExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) StopPipelineExecutionAsync(ctx workflow.Context, input *codepipeline.StopPipelineExecutionInput) *CodepipelineStopPipelineExecutionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopPipelineExecution, input)
    return &CodepipelineStopPipelineExecutionResult{Result: future}
}

func (a *CodePipelineStub) TagResource(ctx workflow.Context, input *codepipeline.TagResourceInput) (*codepipeline.TagResourceOutput, error) {
    var output codepipeline.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) TagResourceAsync(ctx workflow.Context, input *codepipeline.TagResourceInput) *CodepipelineTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &CodepipelineTagResourceResult{Result: future}
}

func (a *CodePipelineStub) UntagResource(ctx workflow.Context, input *codepipeline.UntagResourceInput) (*codepipeline.UntagResourceOutput, error) {
    var output codepipeline.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) UntagResourceAsync(ctx workflow.Context, input *codepipeline.UntagResourceInput) *CodepipelineUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &CodepipelineUntagResourceResult{Result: future}
}

func (a *CodePipelineStub) UpdatePipeline(ctx workflow.Context, input *codepipeline.UpdatePipelineInput) (*codepipeline.UpdatePipelineOutput, error) {
    var output codepipeline.UpdatePipelineOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePipeline, input).Get(ctx, &output)
    return &output, err
}

func (a *CodePipelineStub) UpdatePipelineAsync(ctx workflow.Context, input *codepipeline.UpdatePipelineInput) *CodepipelineUpdatePipelineResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdatePipeline, input)
    return &CodepipelineUpdatePipelineResult{Result: future}
}
