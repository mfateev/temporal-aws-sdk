package awsclients

import (
	"github.com/aws/aws-sdk-go/service/support"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type SupportClient interface {
       AddAttachmentsToSet(ctx workflow.Context, input *support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error)
       AddAttachmentsToSetAsync(ctx workflow.Context, input *support.AddAttachmentsToSetInput) *SupportAddAttachmentsToSetResult

       AddCommunicationToCase(ctx workflow.Context, input *support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error)
       AddCommunicationToCaseAsync(ctx workflow.Context, input *support.AddCommunicationToCaseInput) *SupportAddCommunicationToCaseResult

       CreateCase(ctx workflow.Context, input *support.CreateCaseInput) (*support.CreateCaseOutput, error)
       CreateCaseAsync(ctx workflow.Context, input *support.CreateCaseInput) *SupportCreateCaseResult

       DescribeAttachment(ctx workflow.Context, input *support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error)
       DescribeAttachmentAsync(ctx workflow.Context, input *support.DescribeAttachmentInput) *SupportDescribeAttachmentResult

       DescribeCases(ctx workflow.Context, input *support.DescribeCasesInput) (*support.DescribeCasesOutput, error)
       DescribeCasesAsync(ctx workflow.Context, input *support.DescribeCasesInput) *SupportDescribeCasesResult

       DescribeCommunications(ctx workflow.Context, input *support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error)
       DescribeCommunicationsAsync(ctx workflow.Context, input *support.DescribeCommunicationsInput) *SupportDescribeCommunicationsResult

       DescribeServices(ctx workflow.Context, input *support.DescribeServicesInput) (*support.DescribeServicesOutput, error)
       DescribeServicesAsync(ctx workflow.Context, input *support.DescribeServicesInput) *SupportDescribeServicesResult

       DescribeSeverityLevels(ctx workflow.Context, input *support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error)
       DescribeSeverityLevelsAsync(ctx workflow.Context, input *support.DescribeSeverityLevelsInput) *SupportDescribeSeverityLevelsResult

       DescribeTrustedAdvisorCheckRefreshStatuses(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error)
       DescribeTrustedAdvisorCheckRefreshStatusesAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) *SupportDescribeTrustedAdvisorCheckRefreshStatusesResult

       DescribeTrustedAdvisorCheckResult(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error)
       DescribeTrustedAdvisorCheckResultAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckResultInput) *SupportDescribeTrustedAdvisorCheckResultResult

       DescribeTrustedAdvisorCheckSummaries(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error)
       DescribeTrustedAdvisorCheckSummariesAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput) *SupportDescribeTrustedAdvisorCheckSummariesResult

       DescribeTrustedAdvisorChecks(ctx workflow.Context, input *support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error)
       DescribeTrustedAdvisorChecksAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorChecksInput) *SupportDescribeTrustedAdvisorChecksResult

       RefreshTrustedAdvisorCheck(ctx workflow.Context, input *support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error)
       RefreshTrustedAdvisorCheckAsync(ctx workflow.Context, input *support.RefreshTrustedAdvisorCheckInput) *SupportRefreshTrustedAdvisorCheckResult

       ResolveCase(ctx workflow.Context, input *support.ResolveCaseInput) (*support.ResolveCaseOutput, error)
       ResolveCaseAsync(ctx workflow.Context, input *support.ResolveCaseInput) *SupportResolveCaseResult
}

type SupportAddAttachmentsToSetResult struct {
	Result workflow.Future
}

func (r *SupportAddAttachmentsToSetResult) Get(ctx workflow.Context) (*support.AddAttachmentsToSetOutput, error) {
    var output support.AddAttachmentsToSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportAddCommunicationToCaseResult struct {
	Result workflow.Future
}

func (r *SupportAddCommunicationToCaseResult) Get(ctx workflow.Context) (*support.AddCommunicationToCaseOutput, error) {
    var output support.AddCommunicationToCaseOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportCreateCaseResult struct {
	Result workflow.Future
}

func (r *SupportCreateCaseResult) Get(ctx workflow.Context) (*support.CreateCaseOutput, error) {
    var output support.CreateCaseOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportDescribeAttachmentResult struct {
	Result workflow.Future
}

func (r *SupportDescribeAttachmentResult) Get(ctx workflow.Context) (*support.DescribeAttachmentOutput, error) {
    var output support.DescribeAttachmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportDescribeCasesResult struct {
	Result workflow.Future
}

func (r *SupportDescribeCasesResult) Get(ctx workflow.Context) (*support.DescribeCasesOutput, error) {
    var output support.DescribeCasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportDescribeCommunicationsResult struct {
	Result workflow.Future
}

func (r *SupportDescribeCommunicationsResult) Get(ctx workflow.Context) (*support.DescribeCommunicationsOutput, error) {
    var output support.DescribeCommunicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportDescribeServicesResult struct {
	Result workflow.Future
}

func (r *SupportDescribeServicesResult) Get(ctx workflow.Context) (*support.DescribeServicesOutput, error) {
    var output support.DescribeServicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportDescribeSeverityLevelsResult struct {
	Result workflow.Future
}

func (r *SupportDescribeSeverityLevelsResult) Get(ctx workflow.Context) (*support.DescribeSeverityLevelsOutput, error) {
    var output support.DescribeSeverityLevelsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportDescribeTrustedAdvisorCheckRefreshStatusesResult struct {
	Result workflow.Future
}

func (r *SupportDescribeTrustedAdvisorCheckRefreshStatusesResult) Get(ctx workflow.Context) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
    var output support.DescribeTrustedAdvisorCheckRefreshStatusesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportDescribeTrustedAdvisorCheckResultResult struct {
	Result workflow.Future
}

func (r *SupportDescribeTrustedAdvisorCheckResultResult) Get(ctx workflow.Context) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
    var output support.DescribeTrustedAdvisorCheckResultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportDescribeTrustedAdvisorCheckSummariesResult struct {
	Result workflow.Future
}

func (r *SupportDescribeTrustedAdvisorCheckSummariesResult) Get(ctx workflow.Context) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
    var output support.DescribeTrustedAdvisorCheckSummariesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportDescribeTrustedAdvisorChecksResult struct {
	Result workflow.Future
}

func (r *SupportDescribeTrustedAdvisorChecksResult) Get(ctx workflow.Context) (*support.DescribeTrustedAdvisorChecksOutput, error) {
    var output support.DescribeTrustedAdvisorChecksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportRefreshTrustedAdvisorCheckResult struct {
	Result workflow.Future
}

func (r *SupportRefreshTrustedAdvisorCheckResult) Get(ctx workflow.Context) (*support.RefreshTrustedAdvisorCheckOutput, error) {
    var output support.RefreshTrustedAdvisorCheckOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportResolveCaseResult struct {
	Result workflow.Future
}

func (r *SupportResolveCaseResult) Get(ctx workflow.Context) (*support.ResolveCaseOutput, error) {
    var output support.ResolveCaseOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SupportStub struct {
    activities awsactivities.SupportActivities
}

func NewSupportStub() SupportClient {
    return &SupportStub{}
}

func (a *SupportStub) AddAttachmentsToSet(ctx workflow.Context, input *support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error) {
    var output support.AddAttachmentsToSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddAttachmentsToSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) AddAttachmentsToSetAsync(ctx workflow.Context, input *support.AddAttachmentsToSetInput) *SupportAddAttachmentsToSetResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddAttachmentsToSet, input)
    return &SupportAddAttachmentsToSetResult{Result: future}
}

func (a *SupportStub) AddCommunicationToCase(ctx workflow.Context, input *support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error) {
    var output support.AddCommunicationToCaseOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddCommunicationToCase, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) AddCommunicationToCaseAsync(ctx workflow.Context, input *support.AddCommunicationToCaseInput) *SupportAddCommunicationToCaseResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddCommunicationToCase, input)
    return &SupportAddCommunicationToCaseResult{Result: future}
}

func (a *SupportStub) CreateCase(ctx workflow.Context, input *support.CreateCaseInput) (*support.CreateCaseOutput, error) {
    var output support.CreateCaseOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCase, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) CreateCaseAsync(ctx workflow.Context, input *support.CreateCaseInput) *SupportCreateCaseResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCase, input)
    return &SupportCreateCaseResult{Result: future}
}

func (a *SupportStub) DescribeAttachment(ctx workflow.Context, input *support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error) {
    var output support.DescribeAttachmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAttachment, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) DescribeAttachmentAsync(ctx workflow.Context, input *support.DescribeAttachmentInput) *SupportDescribeAttachmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAttachment, input)
    return &SupportDescribeAttachmentResult{Result: future}
}

func (a *SupportStub) DescribeCases(ctx workflow.Context, input *support.DescribeCasesInput) (*support.DescribeCasesOutput, error) {
    var output support.DescribeCasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCases, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) DescribeCasesAsync(ctx workflow.Context, input *support.DescribeCasesInput) *SupportDescribeCasesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCases, input)
    return &SupportDescribeCasesResult{Result: future}
}

func (a *SupportStub) DescribeCommunications(ctx workflow.Context, input *support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error) {
    var output support.DescribeCommunicationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCommunications, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) DescribeCommunicationsAsync(ctx workflow.Context, input *support.DescribeCommunicationsInput) *SupportDescribeCommunicationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCommunications, input)
    return &SupportDescribeCommunicationsResult{Result: future}
}

func (a *SupportStub) DescribeServices(ctx workflow.Context, input *support.DescribeServicesInput) (*support.DescribeServicesOutput, error) {
    var output support.DescribeServicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeServices, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) DescribeServicesAsync(ctx workflow.Context, input *support.DescribeServicesInput) *SupportDescribeServicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeServices, input)
    return &SupportDescribeServicesResult{Result: future}
}

func (a *SupportStub) DescribeSeverityLevels(ctx workflow.Context, input *support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error) {
    var output support.DescribeSeverityLevelsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSeverityLevels, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) DescribeSeverityLevelsAsync(ctx workflow.Context, input *support.DescribeSeverityLevelsInput) *SupportDescribeSeverityLevelsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSeverityLevels, input)
    return &SupportDescribeSeverityLevelsResult{Result: future}
}

func (a *SupportStub) DescribeTrustedAdvisorCheckRefreshStatuses(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
    var output support.DescribeTrustedAdvisorCheckRefreshStatusesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTrustedAdvisorCheckRefreshStatuses, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) DescribeTrustedAdvisorCheckRefreshStatusesAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) *SupportDescribeTrustedAdvisorCheckRefreshStatusesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTrustedAdvisorCheckRefreshStatuses, input)
    return &SupportDescribeTrustedAdvisorCheckRefreshStatusesResult{Result: future}
}

func (a *SupportStub) DescribeTrustedAdvisorCheckResult(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
    var output support.DescribeTrustedAdvisorCheckResultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTrustedAdvisorCheckResult, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) DescribeTrustedAdvisorCheckResultAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckResultInput) *SupportDescribeTrustedAdvisorCheckResultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTrustedAdvisorCheckResult, input)
    return &SupportDescribeTrustedAdvisorCheckResultResult{Result: future}
}

func (a *SupportStub) DescribeTrustedAdvisorCheckSummaries(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
    var output support.DescribeTrustedAdvisorCheckSummariesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTrustedAdvisorCheckSummaries, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) DescribeTrustedAdvisorCheckSummariesAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput) *SupportDescribeTrustedAdvisorCheckSummariesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTrustedAdvisorCheckSummaries, input)
    return &SupportDescribeTrustedAdvisorCheckSummariesResult{Result: future}
}

func (a *SupportStub) DescribeTrustedAdvisorChecks(ctx workflow.Context, input *support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error) {
    var output support.DescribeTrustedAdvisorChecksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTrustedAdvisorChecks, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) DescribeTrustedAdvisorChecksAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorChecksInput) *SupportDescribeTrustedAdvisorChecksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTrustedAdvisorChecks, input)
    return &SupportDescribeTrustedAdvisorChecksResult{Result: future}
}

func (a *SupportStub) RefreshTrustedAdvisorCheck(ctx workflow.Context, input *support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error) {
    var output support.RefreshTrustedAdvisorCheckOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RefreshTrustedAdvisorCheck, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) RefreshTrustedAdvisorCheckAsync(ctx workflow.Context, input *support.RefreshTrustedAdvisorCheckInput) *SupportRefreshTrustedAdvisorCheckResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RefreshTrustedAdvisorCheck, input)
    return &SupportRefreshTrustedAdvisorCheckResult{Result: future}
}

func (a *SupportStub) ResolveCase(ctx workflow.Context, input *support.ResolveCaseInput) (*support.ResolveCaseOutput, error) {
    var output support.ResolveCaseOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResolveCase, input).Get(ctx, &output)
    return &output, err
}

func (a *SupportStub) ResolveCaseAsync(ctx workflow.Context, input *support.ResolveCaseInput) *SupportResolveCaseResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResolveCase, input)
    return &SupportResolveCaseResult{Result: future}
}
