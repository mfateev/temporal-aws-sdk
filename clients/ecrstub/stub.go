// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ecrstub

import (
	"github.com/aws/aws-sdk-go/service/ecr"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ECRBatchCheckLayerAvailabilityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRBatchCheckLayerAvailabilityFuture) Get(ctx workflow.Context) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
	var output ecr.BatchCheckLayerAvailabilityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRBatchDeleteImageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRBatchDeleteImageFuture) Get(ctx workflow.Context) (*ecr.BatchDeleteImageOutput, error) {
	var output ecr.BatchDeleteImageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRBatchGetImageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRBatchGetImageFuture) Get(ctx workflow.Context) (*ecr.BatchGetImageOutput, error) {
	var output ecr.BatchGetImageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRCompleteLayerUploadFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRCompleteLayerUploadFuture) Get(ctx workflow.Context) (*ecr.CompleteLayerUploadOutput, error) {
	var output ecr.CompleteLayerUploadOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRCreateRepositoryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRCreateRepositoryFuture) Get(ctx workflow.Context) (*ecr.CreateRepositoryOutput, error) {
	var output ecr.CreateRepositoryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRDeleteLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRDeleteLifecyclePolicyFuture) Get(ctx workflow.Context) (*ecr.DeleteLifecyclePolicyOutput, error) {
	var output ecr.DeleteLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRDeleteRepositoryFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRDeleteRepositoryFuture) Get(ctx workflow.Context) (*ecr.DeleteRepositoryOutput, error) {
	var output ecr.DeleteRepositoryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRDeleteRepositoryPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRDeleteRepositoryPolicyFuture) Get(ctx workflow.Context) (*ecr.DeleteRepositoryPolicyOutput, error) {
	var output ecr.DeleteRepositoryPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRDescribeImageScanFindingsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRDescribeImageScanFindingsFuture) Get(ctx workflow.Context) (*ecr.DescribeImageScanFindingsOutput, error) {
	var output ecr.DescribeImageScanFindingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRDescribeImagesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRDescribeImagesFuture) Get(ctx workflow.Context) (*ecr.DescribeImagesOutput, error) {
	var output ecr.DescribeImagesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRDescribeRepositoriesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRDescribeRepositoriesFuture) Get(ctx workflow.Context) (*ecr.DescribeRepositoriesOutput, error) {
	var output ecr.DescribeRepositoriesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRGetAuthorizationTokenFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRGetAuthorizationTokenFuture) Get(ctx workflow.Context) (*ecr.GetAuthorizationTokenOutput, error) {
	var output ecr.GetAuthorizationTokenOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRGetDownloadUrlForLayerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRGetDownloadUrlForLayerFuture) Get(ctx workflow.Context) (*ecr.GetDownloadUrlForLayerOutput, error) {
	var output ecr.GetDownloadUrlForLayerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRGetLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRGetLifecyclePolicyFuture) Get(ctx workflow.Context) (*ecr.GetLifecyclePolicyOutput, error) {
	var output ecr.GetLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRGetLifecyclePolicyPreviewFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRGetLifecyclePolicyPreviewFuture) Get(ctx workflow.Context) (*ecr.GetLifecyclePolicyPreviewOutput, error) {
	var output ecr.GetLifecyclePolicyPreviewOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRGetRepositoryPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRGetRepositoryPolicyFuture) Get(ctx workflow.Context) (*ecr.GetRepositoryPolicyOutput, error) {
	var output ecr.GetRepositoryPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRInitiateLayerUploadFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRInitiateLayerUploadFuture) Get(ctx workflow.Context) (*ecr.InitiateLayerUploadOutput, error) {
	var output ecr.InitiateLayerUploadOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRListImagesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRListImagesFuture) Get(ctx workflow.Context) (*ecr.ListImagesOutput, error) {
	var output ecr.ListImagesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRListTagsForResourceFuture) Get(ctx workflow.Context) (*ecr.ListTagsForResourceOutput, error) {
	var output ecr.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRPutImageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRPutImageFuture) Get(ctx workflow.Context) (*ecr.PutImageOutput, error) {
	var output ecr.PutImageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRPutImageScanningConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRPutImageScanningConfigurationFuture) Get(ctx workflow.Context) (*ecr.PutImageScanningConfigurationOutput, error) {
	var output ecr.PutImageScanningConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRPutImageTagMutabilityFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRPutImageTagMutabilityFuture) Get(ctx workflow.Context) (*ecr.PutImageTagMutabilityOutput, error) {
	var output ecr.PutImageTagMutabilityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRPutLifecyclePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRPutLifecyclePolicyFuture) Get(ctx workflow.Context) (*ecr.PutLifecyclePolicyOutput, error) {
	var output ecr.PutLifecyclePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRSetRepositoryPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRSetRepositoryPolicyFuture) Get(ctx workflow.Context) (*ecr.SetRepositoryPolicyOutput, error) {
	var output ecr.SetRepositoryPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRStartImageScanFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRStartImageScanFuture) Get(ctx workflow.Context) (*ecr.StartImageScanOutput, error) {
	var output ecr.StartImageScanOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRStartLifecyclePolicyPreviewFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRStartLifecyclePolicyPreviewFuture) Get(ctx workflow.Context) (*ecr.StartLifecyclePolicyPreviewOutput, error) {
	var output ecr.StartLifecyclePolicyPreviewOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRTagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRTagResourceFuture) Get(ctx workflow.Context) (*ecr.TagResourceOutput, error) {
	var output ecr.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRUntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRUntagResourceFuture) Get(ctx workflow.Context) (*ecr.UntagResourceOutput, error) {
	var output ecr.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ECRUploadLayerPartFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ECRUploadLayerPartFuture) Get(ctx workflow.Context) (*ecr.UploadLayerPartOutput, error) {
	var output ecr.UploadLayerPartOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchCheckLayerAvailability(ctx workflow.Context, input *ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
	var output ecr.BatchCheckLayerAvailabilityOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.BatchCheckLayerAvailability", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchCheckLayerAvailabilityAsync(ctx workflow.Context, input *ecr.BatchCheckLayerAvailabilityInput) *ECRBatchCheckLayerAvailabilityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.BatchCheckLayerAvailability", input)
	return &ECRBatchCheckLayerAvailabilityFuture{Future: future}
}

func (a *stub) BatchDeleteImage(ctx workflow.Context, input *ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error) {
	var output ecr.BatchDeleteImageOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.BatchDeleteImage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchDeleteImageAsync(ctx workflow.Context, input *ecr.BatchDeleteImageInput) *ECRBatchDeleteImageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.BatchDeleteImage", input)
	return &ECRBatchDeleteImageFuture{Future: future}
}

func (a *stub) BatchGetImage(ctx workflow.Context, input *ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error) {
	var output ecr.BatchGetImageOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.BatchGetImage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BatchGetImageAsync(ctx workflow.Context, input *ecr.BatchGetImageInput) *ECRBatchGetImageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.BatchGetImage", input)
	return &ECRBatchGetImageFuture{Future: future}
}

func (a *stub) CompleteLayerUpload(ctx workflow.Context, input *ecr.CompleteLayerUploadInput) (*ecr.CompleteLayerUploadOutput, error) {
	var output ecr.CompleteLayerUploadOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.CompleteLayerUpload", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CompleteLayerUploadAsync(ctx workflow.Context, input *ecr.CompleteLayerUploadInput) *ECRCompleteLayerUploadFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.CompleteLayerUpload", input)
	return &ECRCompleteLayerUploadFuture{Future: future}
}

func (a *stub) CreateRepository(ctx workflow.Context, input *ecr.CreateRepositoryInput) (*ecr.CreateRepositoryOutput, error) {
	var output ecr.CreateRepositoryOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.CreateRepository", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateRepositoryAsync(ctx workflow.Context, input *ecr.CreateRepositoryInput) *ECRCreateRepositoryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.CreateRepository", input)
	return &ECRCreateRepositoryFuture{Future: future}
}

func (a *stub) DeleteLifecyclePolicy(ctx workflow.Context, input *ecr.DeleteLifecyclePolicyInput) (*ecr.DeleteLifecyclePolicyOutput, error) {
	var output ecr.DeleteLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.DeleteLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLifecyclePolicyAsync(ctx workflow.Context, input *ecr.DeleteLifecyclePolicyInput) *ECRDeleteLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.DeleteLifecyclePolicy", input)
	return &ECRDeleteLifecyclePolicyFuture{Future: future}
}

func (a *stub) DeleteRepository(ctx workflow.Context, input *ecr.DeleteRepositoryInput) (*ecr.DeleteRepositoryOutput, error) {
	var output ecr.DeleteRepositoryOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.DeleteRepository", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRepositoryAsync(ctx workflow.Context, input *ecr.DeleteRepositoryInput) *ECRDeleteRepositoryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.DeleteRepository", input)
	return &ECRDeleteRepositoryFuture{Future: future}
}

func (a *stub) DeleteRepositoryPolicy(ctx workflow.Context, input *ecr.DeleteRepositoryPolicyInput) (*ecr.DeleteRepositoryPolicyOutput, error) {
	var output ecr.DeleteRepositoryPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.DeleteRepositoryPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteRepositoryPolicyAsync(ctx workflow.Context, input *ecr.DeleteRepositoryPolicyInput) *ECRDeleteRepositoryPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.DeleteRepositoryPolicy", input)
	return &ECRDeleteRepositoryPolicyFuture{Future: future}
}

func (a *stub) DescribeImageScanFindings(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) (*ecr.DescribeImageScanFindingsOutput, error) {
	var output ecr.DescribeImageScanFindingsOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.DescribeImageScanFindings", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeImageScanFindingsAsync(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) *ECRDescribeImageScanFindingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.DescribeImageScanFindings", input)
	return &ECRDescribeImageScanFindingsFuture{Future: future}
}

func (a *stub) DescribeImages(ctx workflow.Context, input *ecr.DescribeImagesInput) (*ecr.DescribeImagesOutput, error) {
	var output ecr.DescribeImagesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.DescribeImages", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeImagesAsync(ctx workflow.Context, input *ecr.DescribeImagesInput) *ECRDescribeImagesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.DescribeImages", input)
	return &ECRDescribeImagesFuture{Future: future}
}

func (a *stub) DescribeRepositories(ctx workflow.Context, input *ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error) {
	var output ecr.DescribeRepositoriesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.DescribeRepositories", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeRepositoriesAsync(ctx workflow.Context, input *ecr.DescribeRepositoriesInput) *ECRDescribeRepositoriesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.DescribeRepositories", input)
	return &ECRDescribeRepositoriesFuture{Future: future}
}

func (a *stub) GetAuthorizationToken(ctx workflow.Context, input *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error) {
	var output ecr.GetAuthorizationTokenOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.GetAuthorizationToken", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAuthorizationTokenAsync(ctx workflow.Context, input *ecr.GetAuthorizationTokenInput) *ECRGetAuthorizationTokenFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.GetAuthorizationToken", input)
	return &ECRGetAuthorizationTokenFuture{Future: future}
}

func (a *stub) GetDownloadUrlForLayer(ctx workflow.Context, input *ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error) {
	var output ecr.GetDownloadUrlForLayerOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.GetDownloadUrlForLayer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetDownloadUrlForLayerAsync(ctx workflow.Context, input *ecr.GetDownloadUrlForLayerInput) *ECRGetDownloadUrlForLayerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.GetDownloadUrlForLayer", input)
	return &ECRGetDownloadUrlForLayerFuture{Future: future}
}

func (a *stub) GetLifecyclePolicy(ctx workflow.Context, input *ecr.GetLifecyclePolicyInput) (*ecr.GetLifecyclePolicyOutput, error) {
	var output ecr.GetLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.GetLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLifecyclePolicyAsync(ctx workflow.Context, input *ecr.GetLifecyclePolicyInput) *ECRGetLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.GetLifecyclePolicy", input)
	return &ECRGetLifecyclePolicyFuture{Future: future}
}

func (a *stub) GetLifecyclePolicyPreview(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) (*ecr.GetLifecyclePolicyPreviewOutput, error) {
	var output ecr.GetLifecyclePolicyPreviewOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.GetLifecyclePolicyPreview", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetLifecyclePolicyPreviewAsync(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) *ECRGetLifecyclePolicyPreviewFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.GetLifecyclePolicyPreview", input)
	return &ECRGetLifecyclePolicyPreviewFuture{Future: future}
}

func (a *stub) GetRepositoryPolicy(ctx workflow.Context, input *ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error) {
	var output ecr.GetRepositoryPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.GetRepositoryPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetRepositoryPolicyAsync(ctx workflow.Context, input *ecr.GetRepositoryPolicyInput) *ECRGetRepositoryPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.GetRepositoryPolicy", input)
	return &ECRGetRepositoryPolicyFuture{Future: future}
}

func (a *stub) InitiateLayerUpload(ctx workflow.Context, input *ecr.InitiateLayerUploadInput) (*ecr.InitiateLayerUploadOutput, error) {
	var output ecr.InitiateLayerUploadOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.InitiateLayerUpload", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) InitiateLayerUploadAsync(ctx workflow.Context, input *ecr.InitiateLayerUploadInput) *ECRInitiateLayerUploadFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.InitiateLayerUpload", input)
	return &ECRInitiateLayerUploadFuture{Future: future}
}

func (a *stub) ListImages(ctx workflow.Context, input *ecr.ListImagesInput) (*ecr.ListImagesOutput, error) {
	var output ecr.ListImagesOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.ListImages", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListImagesAsync(ctx workflow.Context, input *ecr.ListImagesInput) *ECRListImagesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.ListImages", input)
	return &ECRListImagesFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *ecr.ListTagsForResourceInput) (*ecr.ListTagsForResourceOutput, error) {
	var output ecr.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *ecr.ListTagsForResourceInput) *ECRListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.ListTagsForResource", input)
	return &ECRListTagsForResourceFuture{Future: future}
}

func (a *stub) PutImage(ctx workflow.Context, input *ecr.PutImageInput) (*ecr.PutImageOutput, error) {
	var output ecr.PutImageOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.PutImage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutImageAsync(ctx workflow.Context, input *ecr.PutImageInput) *ECRPutImageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.PutImage", input)
	return &ECRPutImageFuture{Future: future}
}

func (a *stub) PutImageScanningConfiguration(ctx workflow.Context, input *ecr.PutImageScanningConfigurationInput) (*ecr.PutImageScanningConfigurationOutput, error) {
	var output ecr.PutImageScanningConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.PutImageScanningConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutImageScanningConfigurationAsync(ctx workflow.Context, input *ecr.PutImageScanningConfigurationInput) *ECRPutImageScanningConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.PutImageScanningConfiguration", input)
	return &ECRPutImageScanningConfigurationFuture{Future: future}
}

func (a *stub) PutImageTagMutability(ctx workflow.Context, input *ecr.PutImageTagMutabilityInput) (*ecr.PutImageTagMutabilityOutput, error) {
	var output ecr.PutImageTagMutabilityOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.PutImageTagMutability", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutImageTagMutabilityAsync(ctx workflow.Context, input *ecr.PutImageTagMutabilityInput) *ECRPutImageTagMutabilityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.PutImageTagMutability", input)
	return &ECRPutImageTagMutabilityFuture{Future: future}
}

func (a *stub) PutLifecyclePolicy(ctx workflow.Context, input *ecr.PutLifecyclePolicyInput) (*ecr.PutLifecyclePolicyOutput, error) {
	var output ecr.PutLifecyclePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.PutLifecyclePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutLifecyclePolicyAsync(ctx workflow.Context, input *ecr.PutLifecyclePolicyInput) *ECRPutLifecyclePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.PutLifecyclePolicy", input)
	return &ECRPutLifecyclePolicyFuture{Future: future}
}

func (a *stub) SetRepositoryPolicy(ctx workflow.Context, input *ecr.SetRepositoryPolicyInput) (*ecr.SetRepositoryPolicyOutput, error) {
	var output ecr.SetRepositoryPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.SetRepositoryPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetRepositoryPolicyAsync(ctx workflow.Context, input *ecr.SetRepositoryPolicyInput) *ECRSetRepositoryPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.SetRepositoryPolicy", input)
	return &ECRSetRepositoryPolicyFuture{Future: future}
}

func (a *stub) StartImageScan(ctx workflow.Context, input *ecr.StartImageScanInput) (*ecr.StartImageScanOutput, error) {
	var output ecr.StartImageScanOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.StartImageScan", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartImageScanAsync(ctx workflow.Context, input *ecr.StartImageScanInput) *ECRStartImageScanFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.StartImageScan", input)
	return &ECRStartImageScanFuture{Future: future}
}

func (a *stub) StartLifecyclePolicyPreview(ctx workflow.Context, input *ecr.StartLifecyclePolicyPreviewInput) (*ecr.StartLifecyclePolicyPreviewOutput, error) {
	var output ecr.StartLifecyclePolicyPreviewOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.StartLifecyclePolicyPreview", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartLifecyclePolicyPreviewAsync(ctx workflow.Context, input *ecr.StartLifecyclePolicyPreviewInput) *ECRStartLifecyclePolicyPreviewFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.StartLifecyclePolicyPreview", input)
	return &ECRStartLifecyclePolicyPreviewFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *ecr.TagResourceInput) (*ecr.TagResourceOutput, error) {
	var output ecr.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *ecr.TagResourceInput) *ECRTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.TagResource", input)
	return &ECRTagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *ecr.UntagResourceInput) (*ecr.UntagResourceOutput, error) {
	var output ecr.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *ecr.UntagResourceInput) *ECRUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.UntagResource", input)
	return &ECRUntagResourceFuture{Future: future}
}

func (a *stub) UploadLayerPart(ctx workflow.Context, input *ecr.UploadLayerPartInput) (*ecr.UploadLayerPartOutput, error) {
	var output ecr.UploadLayerPartOutput
	err := workflow.ExecuteActivity(ctx, "aws.ecr.UploadLayerPart", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UploadLayerPartAsync(ctx workflow.Context, input *ecr.UploadLayerPartInput) *ECRUploadLayerPartFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.UploadLayerPart", input)
	return &ECRUploadLayerPartFuture{Future: future}
}

func (a *stub) WaitUntilImageScanComplete(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) error {
	return workflow.ExecuteActivity(ctx, "aws.ecr.WaitUntilImageScanComplete", input).Get(ctx, nil)
}

func (a *stub) WaitUntilImageScanCompleteAsync(ctx workflow.Context, input *ecr.DescribeImageScanFindingsInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.WaitUntilImageScanComplete", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilLifecyclePolicyPreviewComplete(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) error {
	return workflow.ExecuteActivity(ctx, "aws.ecr.WaitUntilLifecyclePolicyPreviewComplete", input).Get(ctx, nil)
}

func (a *stub) WaitUntilLifecyclePolicyPreviewCompleteAsync(ctx workflow.Context, input *ecr.GetLifecyclePolicyPreviewInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ecr.WaitUntilLifecyclePolicyPreviewComplete", input)
	return clients.NewVoidFuture(future)
}
