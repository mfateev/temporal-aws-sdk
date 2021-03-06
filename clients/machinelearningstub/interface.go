// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package machinelearningstub

import (
	"github.com/aws/aws-sdk-go/service/machinelearning"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddTags(ctx workflow.Context, input *machinelearning.AddTagsInput) (*machinelearning.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *machinelearning.AddTagsInput) *MachineLearningAddTagsFuture

	CreateBatchPrediction(ctx workflow.Context, input *machinelearning.CreateBatchPredictionInput) (*machinelearning.CreateBatchPredictionOutput, error)
	CreateBatchPredictionAsync(ctx workflow.Context, input *machinelearning.CreateBatchPredictionInput) *MachineLearningCreateBatchPredictionFuture

	CreateDataSourceFromRDS(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRDSInput) (*machinelearning.CreateDataSourceFromRDSOutput, error)
	CreateDataSourceFromRDSAsync(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRDSInput) *MachineLearningCreateDataSourceFromRDSFuture

	CreateDataSourceFromRedshift(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRedshiftInput) (*machinelearning.CreateDataSourceFromRedshiftOutput, error)
	CreateDataSourceFromRedshiftAsync(ctx workflow.Context, input *machinelearning.CreateDataSourceFromRedshiftInput) *MachineLearningCreateDataSourceFromRedshiftFuture

	CreateDataSourceFromS3(ctx workflow.Context, input *machinelearning.CreateDataSourceFromS3Input) (*machinelearning.CreateDataSourceFromS3Output, error)
	CreateDataSourceFromS3Async(ctx workflow.Context, input *machinelearning.CreateDataSourceFromS3Input) *MachineLearningCreateDataSourceFromS3Future

	CreateEvaluation(ctx workflow.Context, input *machinelearning.CreateEvaluationInput) (*machinelearning.CreateEvaluationOutput, error)
	CreateEvaluationAsync(ctx workflow.Context, input *machinelearning.CreateEvaluationInput) *MachineLearningCreateEvaluationFuture

	CreateMLModel(ctx workflow.Context, input *machinelearning.CreateMLModelInput) (*machinelearning.CreateMLModelOutput, error)
	CreateMLModelAsync(ctx workflow.Context, input *machinelearning.CreateMLModelInput) *MachineLearningCreateMLModelFuture

	CreateRealtimeEndpoint(ctx workflow.Context, input *machinelearning.CreateRealtimeEndpointInput) (*machinelearning.CreateRealtimeEndpointOutput, error)
	CreateRealtimeEndpointAsync(ctx workflow.Context, input *machinelearning.CreateRealtimeEndpointInput) *MachineLearningCreateRealtimeEndpointFuture

	DeleteBatchPrediction(ctx workflow.Context, input *machinelearning.DeleteBatchPredictionInput) (*machinelearning.DeleteBatchPredictionOutput, error)
	DeleteBatchPredictionAsync(ctx workflow.Context, input *machinelearning.DeleteBatchPredictionInput) *MachineLearningDeleteBatchPredictionFuture

	DeleteDataSource(ctx workflow.Context, input *machinelearning.DeleteDataSourceInput) (*machinelearning.DeleteDataSourceOutput, error)
	DeleteDataSourceAsync(ctx workflow.Context, input *machinelearning.DeleteDataSourceInput) *MachineLearningDeleteDataSourceFuture

	DeleteEvaluation(ctx workflow.Context, input *machinelearning.DeleteEvaluationInput) (*machinelearning.DeleteEvaluationOutput, error)
	DeleteEvaluationAsync(ctx workflow.Context, input *machinelearning.DeleteEvaluationInput) *MachineLearningDeleteEvaluationFuture

	DeleteMLModel(ctx workflow.Context, input *machinelearning.DeleteMLModelInput) (*machinelearning.DeleteMLModelOutput, error)
	DeleteMLModelAsync(ctx workflow.Context, input *machinelearning.DeleteMLModelInput) *MachineLearningDeleteMLModelFuture

	DeleteRealtimeEndpoint(ctx workflow.Context, input *machinelearning.DeleteRealtimeEndpointInput) (*machinelearning.DeleteRealtimeEndpointOutput, error)
	DeleteRealtimeEndpointAsync(ctx workflow.Context, input *machinelearning.DeleteRealtimeEndpointInput) *MachineLearningDeleteRealtimeEndpointFuture

	DeleteTags(ctx workflow.Context, input *machinelearning.DeleteTagsInput) (*machinelearning.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *machinelearning.DeleteTagsInput) *MachineLearningDeleteTagsFuture

	DescribeBatchPredictions(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error)
	DescribeBatchPredictionsAsync(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) *MachineLearningDescribeBatchPredictionsFuture

	DescribeDataSources(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error)
	DescribeDataSourcesAsync(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) *MachineLearningDescribeDataSourcesFuture

	DescribeEvaluations(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error)
	DescribeEvaluationsAsync(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) *MachineLearningDescribeEvaluationsFuture

	DescribeMLModels(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error)
	DescribeMLModelsAsync(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) *MachineLearningDescribeMLModelsFuture

	DescribeTags(ctx workflow.Context, input *machinelearning.DescribeTagsInput) (*machinelearning.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *machinelearning.DescribeTagsInput) *MachineLearningDescribeTagsFuture

	GetBatchPrediction(ctx workflow.Context, input *machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error)
	GetBatchPredictionAsync(ctx workflow.Context, input *machinelearning.GetBatchPredictionInput) *MachineLearningGetBatchPredictionFuture

	GetDataSource(ctx workflow.Context, input *machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error)
	GetDataSourceAsync(ctx workflow.Context, input *machinelearning.GetDataSourceInput) *MachineLearningGetDataSourceFuture

	GetEvaluation(ctx workflow.Context, input *machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error)
	GetEvaluationAsync(ctx workflow.Context, input *machinelearning.GetEvaluationInput) *MachineLearningGetEvaluationFuture

	GetMLModel(ctx workflow.Context, input *machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error)
	GetMLModelAsync(ctx workflow.Context, input *machinelearning.GetMLModelInput) *MachineLearningGetMLModelFuture

	Predict(ctx workflow.Context, input *machinelearning.PredictInput) (*machinelearning.PredictOutput, error)
	PredictAsync(ctx workflow.Context, input *machinelearning.PredictInput) *MachineLearningPredictFuture

	UpdateBatchPrediction(ctx workflow.Context, input *machinelearning.UpdateBatchPredictionInput) (*machinelearning.UpdateBatchPredictionOutput, error)
	UpdateBatchPredictionAsync(ctx workflow.Context, input *machinelearning.UpdateBatchPredictionInput) *MachineLearningUpdateBatchPredictionFuture

	UpdateDataSource(ctx workflow.Context, input *machinelearning.UpdateDataSourceInput) (*machinelearning.UpdateDataSourceOutput, error)
	UpdateDataSourceAsync(ctx workflow.Context, input *machinelearning.UpdateDataSourceInput) *MachineLearningUpdateDataSourceFuture

	UpdateEvaluation(ctx workflow.Context, input *machinelearning.UpdateEvaluationInput) (*machinelearning.UpdateEvaluationOutput, error)
	UpdateEvaluationAsync(ctx workflow.Context, input *machinelearning.UpdateEvaluationInput) *MachineLearningUpdateEvaluationFuture

	UpdateMLModel(ctx workflow.Context, input *machinelearning.UpdateMLModelInput) (*machinelearning.UpdateMLModelOutput, error)
	UpdateMLModelAsync(ctx workflow.Context, input *machinelearning.UpdateMLModelInput) *MachineLearningUpdateMLModelFuture

	WaitUntilBatchPredictionAvailable(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) error
	WaitUntilBatchPredictionAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeBatchPredictionsInput) *clients.VoidFuture

	WaitUntilDataSourceAvailable(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) error
	WaitUntilDataSourceAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeDataSourcesInput) *clients.VoidFuture

	WaitUntilEvaluationAvailable(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) error
	WaitUntilEvaluationAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeEvaluationsInput) *clients.VoidFuture

	WaitUntilMLModelAvailable(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) error
	WaitUntilMLModelAvailableAsync(ctx workflow.Context, input *machinelearning.DescribeMLModelsInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
