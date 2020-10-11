// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"go.temporal.io/sdk/workflow"
)

type CostExplorerClient interface {
	CreateAnomalyMonitor(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) (*costexplorer.CreateAnomalyMonitorOutput, error)
	CreateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) *CostExplorerCreateAnomalyMonitorFuture

	CreateAnomalySubscription(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) (*costexplorer.CreateAnomalySubscriptionOutput, error)
	CreateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) *CostExplorerCreateAnomalySubscriptionFuture

	CreateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) (*costexplorer.CreateCostCategoryDefinitionOutput, error)
	CreateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) *CostExplorerCreateCostCategoryDefinitionFuture

	DeleteAnomalyMonitor(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) (*costexplorer.DeleteAnomalyMonitorOutput, error)
	DeleteAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) *CostExplorerDeleteAnomalyMonitorFuture

	DeleteAnomalySubscription(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) (*costexplorer.DeleteAnomalySubscriptionOutput, error)
	DeleteAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) *CostExplorerDeleteAnomalySubscriptionFuture

	DeleteCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) (*costexplorer.DeleteCostCategoryDefinitionOutput, error)
	DeleteCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) *CostExplorerDeleteCostCategoryDefinitionFuture

	DescribeCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) (*costexplorer.DescribeCostCategoryDefinitionOutput, error)
	DescribeCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) *CostExplorerDescribeCostCategoryDefinitionFuture

	GetAnomalies(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) (*costexplorer.GetAnomaliesOutput, error)
	GetAnomaliesAsync(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) *CostExplorerGetAnomaliesFuture

	GetAnomalyMonitors(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) (*costexplorer.GetAnomalyMonitorsOutput, error)
	GetAnomalyMonitorsAsync(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) *CostExplorerGetAnomalyMonitorsFuture

	GetAnomalySubscriptions(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) (*costexplorer.GetAnomalySubscriptionsOutput, error)
	GetAnomalySubscriptionsAsync(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) *CostExplorerGetAnomalySubscriptionsFuture

	GetCostAndUsage(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error)
	GetCostAndUsageAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) *CostExplorerGetCostAndUsageFuture

	GetCostAndUsageWithResources(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) (*costexplorer.GetCostAndUsageWithResourcesOutput, error)
	GetCostAndUsageWithResourcesAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) *CostExplorerGetCostAndUsageWithResourcesFuture

	GetCostForecast(ctx workflow.Context, input *costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error)
	GetCostForecastAsync(ctx workflow.Context, input *costexplorer.GetCostForecastInput) *CostExplorerGetCostForecastFuture

	GetDimensionValues(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error)
	GetDimensionValuesAsync(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) *CostExplorerGetDimensionValuesFuture

	GetReservationCoverage(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error)
	GetReservationCoverageAsync(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) *CostExplorerGetReservationCoverageFuture

	GetReservationPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error)
	GetReservationPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) *CostExplorerGetReservationPurchaseRecommendationFuture

	GetReservationUtilization(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error)
	GetReservationUtilizationAsync(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) *CostExplorerGetReservationUtilizationFuture

	GetRightsizingRecommendation(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) (*costexplorer.GetRightsizingRecommendationOutput, error)
	GetRightsizingRecommendationAsync(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) *CostExplorerGetRightsizingRecommendationFuture

	GetSavingsPlansCoverage(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) (*costexplorer.GetSavingsPlansCoverageOutput, error)
	GetSavingsPlansCoverageAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) *CostExplorerGetSavingsPlansCoverageFuture

	GetSavingsPlansPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error)
	GetSavingsPlansPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) *CostExplorerGetSavingsPlansPurchaseRecommendationFuture

	GetSavingsPlansUtilization(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) (*costexplorer.GetSavingsPlansUtilizationOutput, error)
	GetSavingsPlansUtilizationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) *CostExplorerGetSavingsPlansUtilizationFuture

	GetSavingsPlansUtilizationDetails(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error)
	GetSavingsPlansUtilizationDetailsAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) *CostExplorerGetSavingsPlansUtilizationDetailsFuture

	GetTags(ctx workflow.Context, input *costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *costexplorer.GetTagsInput) *CostExplorerGetTagsFuture

	GetUsageForecast(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error)
	GetUsageForecastAsync(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) *CostExplorerGetUsageForecastFuture

	ListCostCategoryDefinitions(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) (*costexplorer.ListCostCategoryDefinitionsOutput, error)
	ListCostCategoryDefinitionsAsync(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) *CostExplorerListCostCategoryDefinitionsFuture

	ProvideAnomalyFeedback(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) (*costexplorer.ProvideAnomalyFeedbackOutput, error)
	ProvideAnomalyFeedbackAsync(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) *CostExplorerProvideAnomalyFeedbackFuture

	UpdateAnomalyMonitor(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) (*costexplorer.UpdateAnomalyMonitorOutput, error)
	UpdateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) *CostExplorerUpdateAnomalyMonitorFuture

	UpdateAnomalySubscription(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) (*costexplorer.UpdateAnomalySubscriptionOutput, error)
	UpdateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) *CostExplorerUpdateAnomalySubscriptionFuture

	UpdateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) (*costexplorer.UpdateCostCategoryDefinitionOutput, error)
	UpdateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) *CostExplorerUpdateCostCategoryDefinitionFuture
}

type CostExplorerStub struct{}

func NewCostExplorerStub() CostExplorerClient {
	return &CostExplorerStub{}
}

type CostExplorerCreateAnomalyMonitorFuture struct {
	Future workflow.Future
}

func (r *CostExplorerCreateAnomalyMonitorFuture) Get(ctx workflow.Context) (*costexplorer.CreateAnomalyMonitorOutput, error) {
	var output costexplorer.CreateAnomalyMonitorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerCreateAnomalySubscriptionFuture struct {
	Future workflow.Future
}

func (r *CostExplorerCreateAnomalySubscriptionFuture) Get(ctx workflow.Context) (*costexplorer.CreateAnomalySubscriptionOutput, error) {
	var output costexplorer.CreateAnomalySubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerCreateCostCategoryDefinitionFuture struct {
	Future workflow.Future
}

func (r *CostExplorerCreateCostCategoryDefinitionFuture) Get(ctx workflow.Context) (*costexplorer.CreateCostCategoryDefinitionOutput, error) {
	var output costexplorer.CreateCostCategoryDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerDeleteAnomalyMonitorFuture struct {
	Future workflow.Future
}

func (r *CostExplorerDeleteAnomalyMonitorFuture) Get(ctx workflow.Context) (*costexplorer.DeleteAnomalyMonitorOutput, error) {
	var output costexplorer.DeleteAnomalyMonitorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerDeleteAnomalySubscriptionFuture struct {
	Future workflow.Future
}

func (r *CostExplorerDeleteAnomalySubscriptionFuture) Get(ctx workflow.Context) (*costexplorer.DeleteAnomalySubscriptionOutput, error) {
	var output costexplorer.DeleteAnomalySubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerDeleteCostCategoryDefinitionFuture struct {
	Future workflow.Future
}

func (r *CostExplorerDeleteCostCategoryDefinitionFuture) Get(ctx workflow.Context) (*costexplorer.DeleteCostCategoryDefinitionOutput, error) {
	var output costexplorer.DeleteCostCategoryDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerDescribeCostCategoryDefinitionFuture struct {
	Future workflow.Future
}

func (r *CostExplorerDescribeCostCategoryDefinitionFuture) Get(ctx workflow.Context) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {
	var output costexplorer.DescribeCostCategoryDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetAnomaliesFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetAnomaliesFuture) Get(ctx workflow.Context) (*costexplorer.GetAnomaliesOutput, error) {
	var output costexplorer.GetAnomaliesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetAnomalyMonitorsFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetAnomalyMonitorsFuture) Get(ctx workflow.Context) (*costexplorer.GetAnomalyMonitorsOutput, error) {
	var output costexplorer.GetAnomalyMonitorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetAnomalySubscriptionsFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetAnomalySubscriptionsFuture) Get(ctx workflow.Context) (*costexplorer.GetAnomalySubscriptionsOutput, error) {
	var output costexplorer.GetAnomalySubscriptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetCostAndUsageFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetCostAndUsageFuture) Get(ctx workflow.Context) (*costexplorer.GetCostAndUsageOutput, error) {
	var output costexplorer.GetCostAndUsageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetCostAndUsageWithResourcesFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetCostAndUsageWithResourcesFuture) Get(ctx workflow.Context) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {
	var output costexplorer.GetCostAndUsageWithResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetCostForecastFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetCostForecastFuture) Get(ctx workflow.Context) (*costexplorer.GetCostForecastOutput, error) {
	var output costexplorer.GetCostForecastOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetDimensionValuesFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetDimensionValuesFuture) Get(ctx workflow.Context) (*costexplorer.GetDimensionValuesOutput, error) {
	var output costexplorer.GetDimensionValuesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetReservationCoverageFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetReservationCoverageFuture) Get(ctx workflow.Context) (*costexplorer.GetReservationCoverageOutput, error) {
	var output costexplorer.GetReservationCoverageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetReservationPurchaseRecommendationFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetReservationPurchaseRecommendationFuture) Get(ctx workflow.Context) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
	var output costexplorer.GetReservationPurchaseRecommendationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetReservationUtilizationFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetReservationUtilizationFuture) Get(ctx workflow.Context) (*costexplorer.GetReservationUtilizationOutput, error) {
	var output costexplorer.GetReservationUtilizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetRightsizingRecommendationFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetRightsizingRecommendationFuture) Get(ctx workflow.Context) (*costexplorer.GetRightsizingRecommendationOutput, error) {
	var output costexplorer.GetRightsizingRecommendationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetSavingsPlansCoverageFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetSavingsPlansCoverageFuture) Get(ctx workflow.Context) (*costexplorer.GetSavingsPlansCoverageOutput, error) {
	var output costexplorer.GetSavingsPlansCoverageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetSavingsPlansPurchaseRecommendationFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetSavingsPlansPurchaseRecommendationFuture) Get(ctx workflow.Context) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {
	var output costexplorer.GetSavingsPlansPurchaseRecommendationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetSavingsPlansUtilizationFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetSavingsPlansUtilizationFuture) Get(ctx workflow.Context) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {
	var output costexplorer.GetSavingsPlansUtilizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetSavingsPlansUtilizationDetailsFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetSavingsPlansUtilizationDetailsFuture) Get(ctx workflow.Context) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {
	var output costexplorer.GetSavingsPlansUtilizationDetailsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetTagsFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetTagsFuture) Get(ctx workflow.Context) (*costexplorer.GetTagsOutput, error) {
	var output costexplorer.GetTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerGetUsageForecastFuture struct {
	Future workflow.Future
}

func (r *CostExplorerGetUsageForecastFuture) Get(ctx workflow.Context) (*costexplorer.GetUsageForecastOutput, error) {
	var output costexplorer.GetUsageForecastOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerListCostCategoryDefinitionsFuture struct {
	Future workflow.Future
}

func (r *CostExplorerListCostCategoryDefinitionsFuture) Get(ctx workflow.Context) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {
	var output costexplorer.ListCostCategoryDefinitionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerProvideAnomalyFeedbackFuture struct {
	Future workflow.Future
}

func (r *CostExplorerProvideAnomalyFeedbackFuture) Get(ctx workflow.Context) (*costexplorer.ProvideAnomalyFeedbackOutput, error) {
	var output costexplorer.ProvideAnomalyFeedbackOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerUpdateAnomalyMonitorFuture struct {
	Future workflow.Future
}

func (r *CostExplorerUpdateAnomalyMonitorFuture) Get(ctx workflow.Context) (*costexplorer.UpdateAnomalyMonitorOutput, error) {
	var output costexplorer.UpdateAnomalyMonitorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerUpdateAnomalySubscriptionFuture struct {
	Future workflow.Future
}

func (r *CostExplorerUpdateAnomalySubscriptionFuture) Get(ctx workflow.Context) (*costexplorer.UpdateAnomalySubscriptionOutput, error) {
	var output costexplorer.UpdateAnomalySubscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostExplorerUpdateCostCategoryDefinitionFuture struct {
	Future workflow.Future
}

func (r *CostExplorerUpdateCostCategoryDefinitionFuture) Get(ctx workflow.Context) (*costexplorer.UpdateCostCategoryDefinitionOutput, error) {
	var output costexplorer.UpdateCostCategoryDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) CreateAnomalyMonitor(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) (*costexplorer.CreateAnomalyMonitorOutput, error) {
	var output costexplorer.CreateAnomalyMonitorOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateAnomalyMonitor", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) CreateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.CreateAnomalyMonitorInput) *CostExplorerCreateAnomalyMonitorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateAnomalyMonitor", input)
	return &CostExplorerCreateAnomalyMonitorFuture{Future: future}
}

func (a *CostExplorerStub) CreateAnomalySubscription(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) (*costexplorer.CreateAnomalySubscriptionOutput, error) {
	var output costexplorer.CreateAnomalySubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateAnomalySubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) CreateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.CreateAnomalySubscriptionInput) *CostExplorerCreateAnomalySubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateAnomalySubscription", input)
	return &CostExplorerCreateAnomalySubscriptionFuture{Future: future}
}

func (a *CostExplorerStub) CreateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) (*costexplorer.CreateCostCategoryDefinitionOutput, error) {
	var output costexplorer.CreateCostCategoryDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateCostCategoryDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) CreateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.CreateCostCategoryDefinitionInput) *CostExplorerCreateCostCategoryDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.CreateCostCategoryDefinition", input)
	return &CostExplorerCreateCostCategoryDefinitionFuture{Future: future}
}

func (a *CostExplorerStub) DeleteAnomalyMonitor(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) (*costexplorer.DeleteAnomalyMonitorOutput, error) {
	var output costexplorer.DeleteAnomalyMonitorOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteAnomalyMonitor", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) DeleteAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalyMonitorInput) *CostExplorerDeleteAnomalyMonitorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteAnomalyMonitor", input)
	return &CostExplorerDeleteAnomalyMonitorFuture{Future: future}
}

func (a *CostExplorerStub) DeleteAnomalySubscription(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) (*costexplorer.DeleteAnomalySubscriptionOutput, error) {
	var output costexplorer.DeleteAnomalySubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteAnomalySubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) DeleteAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.DeleteAnomalySubscriptionInput) *CostExplorerDeleteAnomalySubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteAnomalySubscription", input)
	return &CostExplorerDeleteAnomalySubscriptionFuture{Future: future}
}

func (a *CostExplorerStub) DeleteCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) (*costexplorer.DeleteCostCategoryDefinitionOutput, error) {
	var output costexplorer.DeleteCostCategoryDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteCostCategoryDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) DeleteCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DeleteCostCategoryDefinitionInput) *CostExplorerDeleteCostCategoryDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.DeleteCostCategoryDefinition", input)
	return &CostExplorerDeleteCostCategoryDefinitionFuture{Future: future}
}

func (a *CostExplorerStub) DescribeCostCategoryDefinition(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {
	var output costexplorer.DescribeCostCategoryDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.DescribeCostCategoryDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) DescribeCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.DescribeCostCategoryDefinitionInput) *CostExplorerDescribeCostCategoryDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.DescribeCostCategoryDefinition", input)
	return &CostExplorerDescribeCostCategoryDefinitionFuture{Future: future}
}

func (a *CostExplorerStub) GetAnomalies(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) (*costexplorer.GetAnomaliesOutput, error) {
	var output costexplorer.GetAnomaliesOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalies", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetAnomaliesAsync(ctx workflow.Context, input *costexplorer.GetAnomaliesInput) *CostExplorerGetAnomaliesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalies", input)
	return &CostExplorerGetAnomaliesFuture{Future: future}
}

func (a *CostExplorerStub) GetAnomalyMonitors(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) (*costexplorer.GetAnomalyMonitorsOutput, error) {
	var output costexplorer.GetAnomalyMonitorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalyMonitors", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetAnomalyMonitorsAsync(ctx workflow.Context, input *costexplorer.GetAnomalyMonitorsInput) *CostExplorerGetAnomalyMonitorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalyMonitors", input)
	return &CostExplorerGetAnomalyMonitorsFuture{Future: future}
}

func (a *CostExplorerStub) GetAnomalySubscriptions(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) (*costexplorer.GetAnomalySubscriptionsOutput, error) {
	var output costexplorer.GetAnomalySubscriptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalySubscriptions", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetAnomalySubscriptionsAsync(ctx workflow.Context, input *costexplorer.GetAnomalySubscriptionsInput) *CostExplorerGetAnomalySubscriptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetAnomalySubscriptions", input)
	return &CostExplorerGetAnomalySubscriptionsFuture{Future: future}
}

func (a *CostExplorerStub) GetCostAndUsage(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error) {
	var output costexplorer.GetCostAndUsageOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostAndUsage", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetCostAndUsageAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageInput) *CostExplorerGetCostAndUsageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostAndUsage", input)
	return &CostExplorerGetCostAndUsageFuture{Future: future}
}

func (a *CostExplorerStub) GetCostAndUsageWithResources(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {
	var output costexplorer.GetCostAndUsageWithResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostAndUsageWithResources", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetCostAndUsageWithResourcesAsync(ctx workflow.Context, input *costexplorer.GetCostAndUsageWithResourcesInput) *CostExplorerGetCostAndUsageWithResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostAndUsageWithResources", input)
	return &CostExplorerGetCostAndUsageWithResourcesFuture{Future: future}
}

func (a *CostExplorerStub) GetCostForecast(ctx workflow.Context, input *costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error) {
	var output costexplorer.GetCostForecastOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostForecast", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetCostForecastAsync(ctx workflow.Context, input *costexplorer.GetCostForecastInput) *CostExplorerGetCostForecastFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetCostForecast", input)
	return &CostExplorerGetCostForecastFuture{Future: future}
}

func (a *CostExplorerStub) GetDimensionValues(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error) {
	var output costexplorer.GetDimensionValuesOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetDimensionValues", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetDimensionValuesAsync(ctx workflow.Context, input *costexplorer.GetDimensionValuesInput) *CostExplorerGetDimensionValuesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetDimensionValues", input)
	return &CostExplorerGetDimensionValuesFuture{Future: future}
}

func (a *CostExplorerStub) GetReservationCoverage(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error) {
	var output costexplorer.GetReservationCoverageOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationCoverage", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetReservationCoverageAsync(ctx workflow.Context, input *costexplorer.GetReservationCoverageInput) *CostExplorerGetReservationCoverageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationCoverage", input)
	return &CostExplorerGetReservationCoverageFuture{Future: future}
}

func (a *CostExplorerStub) GetReservationPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
	var output costexplorer.GetReservationPurchaseRecommendationOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationPurchaseRecommendation", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetReservationPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetReservationPurchaseRecommendationInput) *CostExplorerGetReservationPurchaseRecommendationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationPurchaseRecommendation", input)
	return &CostExplorerGetReservationPurchaseRecommendationFuture{Future: future}
}

func (a *CostExplorerStub) GetReservationUtilization(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error) {
	var output costexplorer.GetReservationUtilizationOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationUtilization", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetReservationUtilizationAsync(ctx workflow.Context, input *costexplorer.GetReservationUtilizationInput) *CostExplorerGetReservationUtilizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetReservationUtilization", input)
	return &CostExplorerGetReservationUtilizationFuture{Future: future}
}

func (a *CostExplorerStub) GetRightsizingRecommendation(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) (*costexplorer.GetRightsizingRecommendationOutput, error) {
	var output costexplorer.GetRightsizingRecommendationOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetRightsizingRecommendation", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetRightsizingRecommendationAsync(ctx workflow.Context, input *costexplorer.GetRightsizingRecommendationInput) *CostExplorerGetRightsizingRecommendationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetRightsizingRecommendation", input)
	return &CostExplorerGetRightsizingRecommendationFuture{Future: future}
}

func (a *CostExplorerStub) GetSavingsPlansCoverage(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) (*costexplorer.GetSavingsPlansCoverageOutput, error) {
	var output costexplorer.GetSavingsPlansCoverageOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansCoverage", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetSavingsPlansCoverageAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansCoverageInput) *CostExplorerGetSavingsPlansCoverageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansCoverage", input)
	return &CostExplorerGetSavingsPlansCoverageFuture{Future: future}
}

func (a *CostExplorerStub) GetSavingsPlansPurchaseRecommendation(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {
	var output costexplorer.GetSavingsPlansPurchaseRecommendationOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansPurchaseRecommendation", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetSavingsPlansPurchaseRecommendationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) *CostExplorerGetSavingsPlansPurchaseRecommendationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansPurchaseRecommendation", input)
	return &CostExplorerGetSavingsPlansPurchaseRecommendationFuture{Future: future}
}

func (a *CostExplorerStub) GetSavingsPlansUtilization(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {
	var output costexplorer.GetSavingsPlansUtilizationOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansUtilization", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetSavingsPlansUtilizationAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationInput) *CostExplorerGetSavingsPlansUtilizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansUtilization", input)
	return &CostExplorerGetSavingsPlansUtilizationFuture{Future: future}
}

func (a *CostExplorerStub) GetSavingsPlansUtilizationDetails(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {
	var output costexplorer.GetSavingsPlansUtilizationDetailsOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansUtilizationDetails", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetSavingsPlansUtilizationDetailsAsync(ctx workflow.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput) *CostExplorerGetSavingsPlansUtilizationDetailsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetSavingsPlansUtilizationDetails", input)
	return &CostExplorerGetSavingsPlansUtilizationDetailsFuture{Future: future}
}

func (a *CostExplorerStub) GetTags(ctx workflow.Context, input *costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error) {
	var output costexplorer.GetTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetTags", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetTagsAsync(ctx workflow.Context, input *costexplorer.GetTagsInput) *CostExplorerGetTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetTags", input)
	return &CostExplorerGetTagsFuture{Future: future}
}

func (a *CostExplorerStub) GetUsageForecast(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error) {
	var output costexplorer.GetUsageForecastOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetUsageForecast", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) GetUsageForecastAsync(ctx workflow.Context, input *costexplorer.GetUsageForecastInput) *CostExplorerGetUsageForecastFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.GetUsageForecast", input)
	return &CostExplorerGetUsageForecastFuture{Future: future}
}

func (a *CostExplorerStub) ListCostCategoryDefinitions(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {
	var output costexplorer.ListCostCategoryDefinitionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.ListCostCategoryDefinitions", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) ListCostCategoryDefinitionsAsync(ctx workflow.Context, input *costexplorer.ListCostCategoryDefinitionsInput) *CostExplorerListCostCategoryDefinitionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.ListCostCategoryDefinitions", input)
	return &CostExplorerListCostCategoryDefinitionsFuture{Future: future}
}

func (a *CostExplorerStub) ProvideAnomalyFeedback(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) (*costexplorer.ProvideAnomalyFeedbackOutput, error) {
	var output costexplorer.ProvideAnomalyFeedbackOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.ProvideAnomalyFeedback", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) ProvideAnomalyFeedbackAsync(ctx workflow.Context, input *costexplorer.ProvideAnomalyFeedbackInput) *CostExplorerProvideAnomalyFeedbackFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.ProvideAnomalyFeedback", input)
	return &CostExplorerProvideAnomalyFeedbackFuture{Future: future}
}

func (a *CostExplorerStub) UpdateAnomalyMonitor(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) (*costexplorer.UpdateAnomalyMonitorOutput, error) {
	var output costexplorer.UpdateAnomalyMonitorOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateAnomalyMonitor", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) UpdateAnomalyMonitorAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalyMonitorInput) *CostExplorerUpdateAnomalyMonitorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateAnomalyMonitor", input)
	return &CostExplorerUpdateAnomalyMonitorFuture{Future: future}
}

func (a *CostExplorerStub) UpdateAnomalySubscription(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) (*costexplorer.UpdateAnomalySubscriptionOutput, error) {
	var output costexplorer.UpdateAnomalySubscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateAnomalySubscription", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) UpdateAnomalySubscriptionAsync(ctx workflow.Context, input *costexplorer.UpdateAnomalySubscriptionInput) *CostExplorerUpdateAnomalySubscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateAnomalySubscription", input)
	return &CostExplorerUpdateAnomalySubscriptionFuture{Future: future}
}

func (a *CostExplorerStub) UpdateCostCategoryDefinition(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) (*costexplorer.UpdateCostCategoryDefinitionOutput, error) {
	var output costexplorer.UpdateCostCategoryDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateCostCategoryDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *CostExplorerStub) UpdateCostCategoryDefinitionAsync(ctx workflow.Context, input *costexplorer.UpdateCostCategoryDefinitionInput) *CostExplorerUpdateCostCategoryDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costexplorer.UpdateCostCategoryDefinition", input)
	return &CostExplorerUpdateCostCategoryDefinitionFuture{Future: future}
}
