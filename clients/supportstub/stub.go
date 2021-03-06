// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package supportstub

import (
	"github.com/aws/aws-sdk-go/service/support"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type SupportAddAttachmentsToSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportAddAttachmentsToSetFuture) Get(ctx workflow.Context) (*support.AddAttachmentsToSetOutput, error) {
	var output support.AddAttachmentsToSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportAddCommunicationToCaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportAddCommunicationToCaseFuture) Get(ctx workflow.Context) (*support.AddCommunicationToCaseOutput, error) {
	var output support.AddCommunicationToCaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportCreateCaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportCreateCaseFuture) Get(ctx workflow.Context) (*support.CreateCaseOutput, error) {
	var output support.CreateCaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportDescribeAttachmentFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportDescribeAttachmentFuture) Get(ctx workflow.Context) (*support.DescribeAttachmentOutput, error) {
	var output support.DescribeAttachmentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportDescribeCasesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportDescribeCasesFuture) Get(ctx workflow.Context) (*support.DescribeCasesOutput, error) {
	var output support.DescribeCasesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportDescribeCommunicationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportDescribeCommunicationsFuture) Get(ctx workflow.Context) (*support.DescribeCommunicationsOutput, error) {
	var output support.DescribeCommunicationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportDescribeServicesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportDescribeServicesFuture) Get(ctx workflow.Context) (*support.DescribeServicesOutput, error) {
	var output support.DescribeServicesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportDescribeSeverityLevelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportDescribeSeverityLevelsFuture) Get(ctx workflow.Context) (*support.DescribeSeverityLevelsOutput, error) {
	var output support.DescribeSeverityLevelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportDescribeTrustedAdvisorCheckRefreshStatusesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportDescribeTrustedAdvisorCheckRefreshStatusesFuture) Get(ctx workflow.Context) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
	var output support.DescribeTrustedAdvisorCheckRefreshStatusesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportDescribeTrustedAdvisorCheckResultFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportDescribeTrustedAdvisorCheckResultFuture) Get(ctx workflow.Context) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
	var output support.DescribeTrustedAdvisorCheckResultOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportDescribeTrustedAdvisorCheckSummariesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportDescribeTrustedAdvisorCheckSummariesFuture) Get(ctx workflow.Context) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
	var output support.DescribeTrustedAdvisorCheckSummariesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportDescribeTrustedAdvisorChecksFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportDescribeTrustedAdvisorChecksFuture) Get(ctx workflow.Context) (*support.DescribeTrustedAdvisorChecksOutput, error) {
	var output support.DescribeTrustedAdvisorChecksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportRefreshTrustedAdvisorCheckFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportRefreshTrustedAdvisorCheckFuture) Get(ctx workflow.Context) (*support.RefreshTrustedAdvisorCheckOutput, error) {
	var output support.RefreshTrustedAdvisorCheckOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SupportResolveCaseFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SupportResolveCaseFuture) Get(ctx workflow.Context) (*support.ResolveCaseOutput, error) {
	var output support.ResolveCaseOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddAttachmentsToSet(ctx workflow.Context, input *support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error) {
	var output support.AddAttachmentsToSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.AddAttachmentsToSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddAttachmentsToSetAsync(ctx workflow.Context, input *support.AddAttachmentsToSetInput) *SupportAddAttachmentsToSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.AddAttachmentsToSet", input)
	return &SupportAddAttachmentsToSetFuture{Future: future}
}

func (a *stub) AddCommunicationToCase(ctx workflow.Context, input *support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error) {
	var output support.AddCommunicationToCaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.AddCommunicationToCase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddCommunicationToCaseAsync(ctx workflow.Context, input *support.AddCommunicationToCaseInput) *SupportAddCommunicationToCaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.AddCommunicationToCase", input)
	return &SupportAddCommunicationToCaseFuture{Future: future}
}

func (a *stub) CreateCase(ctx workflow.Context, input *support.CreateCaseInput) (*support.CreateCaseOutput, error) {
	var output support.CreateCaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.CreateCase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateCaseAsync(ctx workflow.Context, input *support.CreateCaseInput) *SupportCreateCaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.CreateCase", input)
	return &SupportCreateCaseFuture{Future: future}
}

func (a *stub) DescribeAttachment(ctx workflow.Context, input *support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error) {
	var output support.DescribeAttachmentOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.DescribeAttachment", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAttachmentAsync(ctx workflow.Context, input *support.DescribeAttachmentInput) *SupportDescribeAttachmentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.DescribeAttachment", input)
	return &SupportDescribeAttachmentFuture{Future: future}
}

func (a *stub) DescribeCases(ctx workflow.Context, input *support.DescribeCasesInput) (*support.DescribeCasesOutput, error) {
	var output support.DescribeCasesOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.DescribeCases", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeCasesAsync(ctx workflow.Context, input *support.DescribeCasesInput) *SupportDescribeCasesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.DescribeCases", input)
	return &SupportDescribeCasesFuture{Future: future}
}

func (a *stub) DescribeCommunications(ctx workflow.Context, input *support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error) {
	var output support.DescribeCommunicationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.DescribeCommunications", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeCommunicationsAsync(ctx workflow.Context, input *support.DescribeCommunicationsInput) *SupportDescribeCommunicationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.DescribeCommunications", input)
	return &SupportDescribeCommunicationsFuture{Future: future}
}

func (a *stub) DescribeServices(ctx workflow.Context, input *support.DescribeServicesInput) (*support.DescribeServicesOutput, error) {
	var output support.DescribeServicesOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.DescribeServices", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeServicesAsync(ctx workflow.Context, input *support.DescribeServicesInput) *SupportDescribeServicesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.DescribeServices", input)
	return &SupportDescribeServicesFuture{Future: future}
}

func (a *stub) DescribeSeverityLevels(ctx workflow.Context, input *support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error) {
	var output support.DescribeSeverityLevelsOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.DescribeSeverityLevels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeSeverityLevelsAsync(ctx workflow.Context, input *support.DescribeSeverityLevelsInput) *SupportDescribeSeverityLevelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.DescribeSeverityLevels", input)
	return &SupportDescribeSeverityLevelsFuture{Future: future}
}

func (a *stub) DescribeTrustedAdvisorCheckRefreshStatuses(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
	var output support.DescribeTrustedAdvisorCheckRefreshStatusesOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.DescribeTrustedAdvisorCheckRefreshStatuses", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTrustedAdvisorCheckRefreshStatusesAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) *SupportDescribeTrustedAdvisorCheckRefreshStatusesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.DescribeTrustedAdvisorCheckRefreshStatuses", input)
	return &SupportDescribeTrustedAdvisorCheckRefreshStatusesFuture{Future: future}
}

func (a *stub) DescribeTrustedAdvisorCheckResult(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
	var output support.DescribeTrustedAdvisorCheckResultOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.DescribeTrustedAdvisorCheckResult", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTrustedAdvisorCheckResultAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckResultInput) *SupportDescribeTrustedAdvisorCheckResultFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.DescribeTrustedAdvisorCheckResult", input)
	return &SupportDescribeTrustedAdvisorCheckResultFuture{Future: future}
}

func (a *stub) DescribeTrustedAdvisorCheckSummaries(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
	var output support.DescribeTrustedAdvisorCheckSummariesOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.DescribeTrustedAdvisorCheckSummaries", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTrustedAdvisorCheckSummariesAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput) *SupportDescribeTrustedAdvisorCheckSummariesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.DescribeTrustedAdvisorCheckSummaries", input)
	return &SupportDescribeTrustedAdvisorCheckSummariesFuture{Future: future}
}

func (a *stub) DescribeTrustedAdvisorChecks(ctx workflow.Context, input *support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error) {
	var output support.DescribeTrustedAdvisorChecksOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.DescribeTrustedAdvisorChecks", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTrustedAdvisorChecksAsync(ctx workflow.Context, input *support.DescribeTrustedAdvisorChecksInput) *SupportDescribeTrustedAdvisorChecksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.DescribeTrustedAdvisorChecks", input)
	return &SupportDescribeTrustedAdvisorChecksFuture{Future: future}
}

func (a *stub) RefreshTrustedAdvisorCheck(ctx workflow.Context, input *support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error) {
	var output support.RefreshTrustedAdvisorCheckOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.RefreshTrustedAdvisorCheck", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RefreshTrustedAdvisorCheckAsync(ctx workflow.Context, input *support.RefreshTrustedAdvisorCheckInput) *SupportRefreshTrustedAdvisorCheckFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.RefreshTrustedAdvisorCheck", input)
	return &SupportRefreshTrustedAdvisorCheckFuture{Future: future}
}

func (a *stub) ResolveCase(ctx workflow.Context, input *support.ResolveCaseInput) (*support.ResolveCaseOutput, error) {
	var output support.ResolveCaseOutput
	err := workflow.ExecuteActivity(ctx, "aws.support.ResolveCase", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ResolveCaseAsync(ctx workflow.Context, input *support.ResolveCaseInput) *SupportResolveCaseFuture {
	future := workflow.ExecuteActivity(ctx, "aws.support.ResolveCase", input)
	return &SupportResolveCaseFuture{Future: future}
}
