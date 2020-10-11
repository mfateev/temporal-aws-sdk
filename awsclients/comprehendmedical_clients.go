// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"go.temporal.io/sdk/workflow"
)

type ComprehendMedicalClient interface {
	DescribeEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error)
	DescribeEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) *ComprehendMedicalDescribeEntitiesDetectionV2JobFuture

	DescribeICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error)
	DescribeICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) *ComprehendMedicalDescribeICD10CMInferenceJobFuture

	DescribePHIDetectionJob(ctx workflow.Context, input *comprehendmedical.DescribePHIDetectionJobInput) (*comprehendmedical.DescribePHIDetectionJobOutput, error)
	DescribePHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.DescribePHIDetectionJobInput) *ComprehendMedicalDescribePHIDetectionJobFuture

	DescribeRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error)
	DescribeRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) *ComprehendMedicalDescribeRxNormInferenceJobFuture

	DetectEntities(ctx workflow.Context, input *comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error)
	DetectEntitiesAsync(ctx workflow.Context, input *comprehendmedical.DetectEntitiesInput) *ComprehendMedicalDetectEntitiesFuture

	DetectEntitiesV2(ctx workflow.Context, input *comprehendmedical.DetectEntitiesV2Input) (*comprehendmedical.DetectEntitiesV2Output, error)
	DetectEntitiesV2Async(ctx workflow.Context, input *comprehendmedical.DetectEntitiesV2Input) *ComprehendMedicalDetectEntitiesV2Future

	DetectPHI(ctx workflow.Context, input *comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error)
	DetectPHIAsync(ctx workflow.Context, input *comprehendmedical.DetectPHIInput) *ComprehendMedicalDetectPHIFuture

	InferICD10CM(ctx workflow.Context, input *comprehendmedical.InferICD10CMInput) (*comprehendmedical.InferICD10CMOutput, error)
	InferICD10CMAsync(ctx workflow.Context, input *comprehendmedical.InferICD10CMInput) *ComprehendMedicalInferICD10CMFuture

	InferRxNorm(ctx workflow.Context, input *comprehendmedical.InferRxNormInput) (*comprehendmedical.InferRxNormOutput, error)
	InferRxNormAsync(ctx workflow.Context, input *comprehendmedical.InferRxNormInput) *ComprehendMedicalInferRxNormFuture

	ListEntitiesDetectionV2Jobs(ctx workflow.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error)
	ListEntitiesDetectionV2JobsAsync(ctx workflow.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) *ComprehendMedicalListEntitiesDetectionV2JobsFuture

	ListICD10CMInferenceJobs(ctx workflow.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error)
	ListICD10CMInferenceJobsAsync(ctx workflow.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) *ComprehendMedicalListICD10CMInferenceJobsFuture

	ListPHIDetectionJobs(ctx workflow.Context, input *comprehendmedical.ListPHIDetectionJobsInput) (*comprehendmedical.ListPHIDetectionJobsOutput, error)
	ListPHIDetectionJobsAsync(ctx workflow.Context, input *comprehendmedical.ListPHIDetectionJobsInput) *ComprehendMedicalListPHIDetectionJobsFuture

	ListRxNormInferenceJobs(ctx workflow.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) (*comprehendmedical.ListRxNormInferenceJobsOutput, error)
	ListRxNormInferenceJobsAsync(ctx workflow.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) *ComprehendMedicalListRxNormInferenceJobsFuture

	StartEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error)
	StartEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) *ComprehendMedicalStartEntitiesDetectionV2JobFuture

	StartICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) (*comprehendmedical.StartICD10CMInferenceJobOutput, error)
	StartICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) *ComprehendMedicalStartICD10CMInferenceJobFuture

	StartPHIDetectionJob(ctx workflow.Context, input *comprehendmedical.StartPHIDetectionJobInput) (*comprehendmedical.StartPHIDetectionJobOutput, error)
	StartPHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.StartPHIDetectionJobInput) *ComprehendMedicalStartPHIDetectionJobFuture

	StartRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.StartRxNormInferenceJobInput) (*comprehendmedical.StartRxNormInferenceJobOutput, error)
	StartRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StartRxNormInferenceJobInput) *ComprehendMedicalStartRxNormInferenceJobFuture

	StopEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error)
	StopEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) *ComprehendMedicalStopEntitiesDetectionV2JobFuture

	StopICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) (*comprehendmedical.StopICD10CMInferenceJobOutput, error)
	StopICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) *ComprehendMedicalStopICD10CMInferenceJobFuture

	StopPHIDetectionJob(ctx workflow.Context, input *comprehendmedical.StopPHIDetectionJobInput) (*comprehendmedical.StopPHIDetectionJobOutput, error)
	StopPHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.StopPHIDetectionJobInput) *ComprehendMedicalStopPHIDetectionJobFuture

	StopRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.StopRxNormInferenceJobInput) (*comprehendmedical.StopRxNormInferenceJobOutput, error)
	StopRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StopRxNormInferenceJobInput) *ComprehendMedicalStopRxNormInferenceJobFuture
}

type ComprehendMedicalStub struct{}

func NewComprehendMedicalStub() ComprehendMedicalClient {
	return &ComprehendMedicalStub{}
}

type ComprehendMedicalDescribeEntitiesDetectionV2JobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalDescribeEntitiesDetectionV2JobFuture) Get(ctx workflow.Context) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.DescribeEntitiesDetectionV2JobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalDescribeICD10CMInferenceJobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalDescribeICD10CMInferenceJobFuture) Get(ctx workflow.Context) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.DescribeICD10CMInferenceJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalDescribePHIDetectionJobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalDescribePHIDetectionJobFuture) Get(ctx workflow.Context) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	var output comprehendmedical.DescribePHIDetectionJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalDescribeRxNormInferenceJobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalDescribeRxNormInferenceJobFuture) Get(ctx workflow.Context) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error) {
	var output comprehendmedical.DescribeRxNormInferenceJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalDetectEntitiesFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalDetectEntitiesFuture) Get(ctx workflow.Context) (*comprehendmedical.DetectEntitiesOutput, error) {
	var output comprehendmedical.DetectEntitiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalDetectEntitiesV2Future struct {
	Future workflow.Future
}

func (r *ComprehendMedicalDetectEntitiesV2Future) Get(ctx workflow.Context) (*comprehendmedical.DetectEntitiesV2Output, error) {
	var output comprehendmedical.DetectEntitiesV2Output
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalDetectPHIFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalDetectPHIFuture) Get(ctx workflow.Context) (*comprehendmedical.DetectPHIOutput, error) {
	var output comprehendmedical.DetectPHIOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalInferICD10CMFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalInferICD10CMFuture) Get(ctx workflow.Context) (*comprehendmedical.InferICD10CMOutput, error) {
	var output comprehendmedical.InferICD10CMOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalInferRxNormFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalInferRxNormFuture) Get(ctx workflow.Context) (*comprehendmedical.InferRxNormOutput, error) {
	var output comprehendmedical.InferRxNormOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalListEntitiesDetectionV2JobsFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalListEntitiesDetectionV2JobsFuture) Get(ctx workflow.Context) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	var output comprehendmedical.ListEntitiesDetectionV2JobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalListICD10CMInferenceJobsFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalListICD10CMInferenceJobsFuture) Get(ctx workflow.Context) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error) {
	var output comprehendmedical.ListICD10CMInferenceJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalListPHIDetectionJobsFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalListPHIDetectionJobsFuture) Get(ctx workflow.Context) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	var output comprehendmedical.ListPHIDetectionJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalListRxNormInferenceJobsFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalListRxNormInferenceJobsFuture) Get(ctx workflow.Context) (*comprehendmedical.ListRxNormInferenceJobsOutput, error) {
	var output comprehendmedical.ListRxNormInferenceJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalStartEntitiesDetectionV2JobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalStartEntitiesDetectionV2JobFuture) Get(ctx workflow.Context) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.StartEntitiesDetectionV2JobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalStartICD10CMInferenceJobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalStartICD10CMInferenceJobFuture) Get(ctx workflow.Context) (*comprehendmedical.StartICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.StartICD10CMInferenceJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalStartPHIDetectionJobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalStartPHIDetectionJobFuture) Get(ctx workflow.Context) (*comprehendmedical.StartPHIDetectionJobOutput, error) {
	var output comprehendmedical.StartPHIDetectionJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalStartRxNormInferenceJobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalStartRxNormInferenceJobFuture) Get(ctx workflow.Context) (*comprehendmedical.StartRxNormInferenceJobOutput, error) {
	var output comprehendmedical.StartRxNormInferenceJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalStopEntitiesDetectionV2JobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalStopEntitiesDetectionV2JobFuture) Get(ctx workflow.Context) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.StopEntitiesDetectionV2JobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalStopICD10CMInferenceJobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalStopICD10CMInferenceJobFuture) Get(ctx workflow.Context) (*comprehendmedical.StopICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.StopICD10CMInferenceJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalStopPHIDetectionJobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalStopPHIDetectionJobFuture) Get(ctx workflow.Context) (*comprehendmedical.StopPHIDetectionJobOutput, error) {
	var output comprehendmedical.StopPHIDetectionJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ComprehendMedicalStopRxNormInferenceJobFuture struct {
	Future workflow.Future
}

func (r *ComprehendMedicalStopRxNormInferenceJobFuture) Get(ctx workflow.Context) (*comprehendmedical.StopRxNormInferenceJobOutput, error) {
	var output comprehendmedical.StopRxNormInferenceJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DescribeEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.DescribeEntitiesDetectionV2JobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DescribeEntitiesDetectionV2Job", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DescribeEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) *ComprehendMedicalDescribeEntitiesDetectionV2JobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DescribeEntitiesDetectionV2Job", input)
	return &ComprehendMedicalDescribeEntitiesDetectionV2JobFuture{Future: future}
}

func (a *ComprehendMedicalStub) DescribeICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.DescribeICD10CMInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DescribeICD10CMInferenceJob", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DescribeICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) *ComprehendMedicalDescribeICD10CMInferenceJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DescribeICD10CMInferenceJob", input)
	return &ComprehendMedicalDescribeICD10CMInferenceJobFuture{Future: future}
}

func (a *ComprehendMedicalStub) DescribePHIDetectionJob(ctx workflow.Context, input *comprehendmedical.DescribePHIDetectionJobInput) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	var output comprehendmedical.DescribePHIDetectionJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DescribePHIDetectionJob", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DescribePHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.DescribePHIDetectionJobInput) *ComprehendMedicalDescribePHIDetectionJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DescribePHIDetectionJob", input)
	return &ComprehendMedicalDescribePHIDetectionJobFuture{Future: future}
}

func (a *ComprehendMedicalStub) DescribeRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error) {
	var output comprehendmedical.DescribeRxNormInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DescribeRxNormInferenceJob", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DescribeRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) *ComprehendMedicalDescribeRxNormInferenceJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DescribeRxNormInferenceJob", input)
	return &ComprehendMedicalDescribeRxNormInferenceJobFuture{Future: future}
}

func (a *ComprehendMedicalStub) DetectEntities(ctx workflow.Context, input *comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error) {
	var output comprehendmedical.DetectEntitiesOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DetectEntities", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DetectEntitiesAsync(ctx workflow.Context, input *comprehendmedical.DetectEntitiesInput) *ComprehendMedicalDetectEntitiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DetectEntities", input)
	return &ComprehendMedicalDetectEntitiesFuture{Future: future}
}

func (a *ComprehendMedicalStub) DetectEntitiesV2(ctx workflow.Context, input *comprehendmedical.DetectEntitiesV2Input) (*comprehendmedical.DetectEntitiesV2Output, error) {
	var output comprehendmedical.DetectEntitiesV2Output
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DetectEntitiesV2", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DetectEntitiesV2Async(ctx workflow.Context, input *comprehendmedical.DetectEntitiesV2Input) *ComprehendMedicalDetectEntitiesV2Future {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DetectEntitiesV2", input)
	return &ComprehendMedicalDetectEntitiesV2Future{Future: future}
}

func (a *ComprehendMedicalStub) DetectPHI(ctx workflow.Context, input *comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error) {
	var output comprehendmedical.DetectPHIOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DetectPHI", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) DetectPHIAsync(ctx workflow.Context, input *comprehendmedical.DetectPHIInput) *ComprehendMedicalDetectPHIFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.DetectPHI", input)
	return &ComprehendMedicalDetectPHIFuture{Future: future}
}

func (a *ComprehendMedicalStub) InferICD10CM(ctx workflow.Context, input *comprehendmedical.InferICD10CMInput) (*comprehendmedical.InferICD10CMOutput, error) {
	var output comprehendmedical.InferICD10CMOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.InferICD10CM", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) InferICD10CMAsync(ctx workflow.Context, input *comprehendmedical.InferICD10CMInput) *ComprehendMedicalInferICD10CMFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.InferICD10CM", input)
	return &ComprehendMedicalInferICD10CMFuture{Future: future}
}

func (a *ComprehendMedicalStub) InferRxNorm(ctx workflow.Context, input *comprehendmedical.InferRxNormInput) (*comprehendmedical.InferRxNormOutput, error) {
	var output comprehendmedical.InferRxNormOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.InferRxNorm", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) InferRxNormAsync(ctx workflow.Context, input *comprehendmedical.InferRxNormInput) *ComprehendMedicalInferRxNormFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.InferRxNorm", input)
	return &ComprehendMedicalInferRxNormFuture{Future: future}
}

func (a *ComprehendMedicalStub) ListEntitiesDetectionV2Jobs(ctx workflow.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	var output comprehendmedical.ListEntitiesDetectionV2JobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.ListEntitiesDetectionV2Jobs", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) ListEntitiesDetectionV2JobsAsync(ctx workflow.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) *ComprehendMedicalListEntitiesDetectionV2JobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.ListEntitiesDetectionV2Jobs", input)
	return &ComprehendMedicalListEntitiesDetectionV2JobsFuture{Future: future}
}

func (a *ComprehendMedicalStub) ListICD10CMInferenceJobs(ctx workflow.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error) {
	var output comprehendmedical.ListICD10CMInferenceJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.ListICD10CMInferenceJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) ListICD10CMInferenceJobsAsync(ctx workflow.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) *ComprehendMedicalListICD10CMInferenceJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.ListICD10CMInferenceJobs", input)
	return &ComprehendMedicalListICD10CMInferenceJobsFuture{Future: future}
}

func (a *ComprehendMedicalStub) ListPHIDetectionJobs(ctx workflow.Context, input *comprehendmedical.ListPHIDetectionJobsInput) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	var output comprehendmedical.ListPHIDetectionJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.ListPHIDetectionJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) ListPHIDetectionJobsAsync(ctx workflow.Context, input *comprehendmedical.ListPHIDetectionJobsInput) *ComprehendMedicalListPHIDetectionJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.ListPHIDetectionJobs", input)
	return &ComprehendMedicalListPHIDetectionJobsFuture{Future: future}
}

func (a *ComprehendMedicalStub) ListRxNormInferenceJobs(ctx workflow.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) (*comprehendmedical.ListRxNormInferenceJobsOutput, error) {
	var output comprehendmedical.ListRxNormInferenceJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.ListRxNormInferenceJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) ListRxNormInferenceJobsAsync(ctx workflow.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) *ComprehendMedicalListRxNormInferenceJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.ListRxNormInferenceJobs", input)
	return &ComprehendMedicalListRxNormInferenceJobsFuture{Future: future}
}

func (a *ComprehendMedicalStub) StartEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.StartEntitiesDetectionV2JobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StartEntitiesDetectionV2Job", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StartEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) *ComprehendMedicalStartEntitiesDetectionV2JobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StartEntitiesDetectionV2Job", input)
	return &ComprehendMedicalStartEntitiesDetectionV2JobFuture{Future: future}
}

func (a *ComprehendMedicalStub) StartICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) (*comprehendmedical.StartICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.StartICD10CMInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StartICD10CMInferenceJob", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StartICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) *ComprehendMedicalStartICD10CMInferenceJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StartICD10CMInferenceJob", input)
	return &ComprehendMedicalStartICD10CMInferenceJobFuture{Future: future}
}

func (a *ComprehendMedicalStub) StartPHIDetectionJob(ctx workflow.Context, input *comprehendmedical.StartPHIDetectionJobInput) (*comprehendmedical.StartPHIDetectionJobOutput, error) {
	var output comprehendmedical.StartPHIDetectionJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StartPHIDetectionJob", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StartPHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.StartPHIDetectionJobInput) *ComprehendMedicalStartPHIDetectionJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StartPHIDetectionJob", input)
	return &ComprehendMedicalStartPHIDetectionJobFuture{Future: future}
}

func (a *ComprehendMedicalStub) StartRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.StartRxNormInferenceJobInput) (*comprehendmedical.StartRxNormInferenceJobOutput, error) {
	var output comprehendmedical.StartRxNormInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StartRxNormInferenceJob", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StartRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StartRxNormInferenceJobInput) *ComprehendMedicalStartRxNormInferenceJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StartRxNormInferenceJob", input)
	return &ComprehendMedicalStartRxNormInferenceJobFuture{Future: future}
}

func (a *ComprehendMedicalStub) StopEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error) {
	var output comprehendmedical.StopEntitiesDetectionV2JobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StopEntitiesDetectionV2Job", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StopEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) *ComprehendMedicalStopEntitiesDetectionV2JobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StopEntitiesDetectionV2Job", input)
	return &ComprehendMedicalStopEntitiesDetectionV2JobFuture{Future: future}
}

func (a *ComprehendMedicalStub) StopICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) (*comprehendmedical.StopICD10CMInferenceJobOutput, error) {
	var output comprehendmedical.StopICD10CMInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StopICD10CMInferenceJob", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StopICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) *ComprehendMedicalStopICD10CMInferenceJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StopICD10CMInferenceJob", input)
	return &ComprehendMedicalStopICD10CMInferenceJobFuture{Future: future}
}

func (a *ComprehendMedicalStub) StopPHIDetectionJob(ctx workflow.Context, input *comprehendmedical.StopPHIDetectionJobInput) (*comprehendmedical.StopPHIDetectionJobOutput, error) {
	var output comprehendmedical.StopPHIDetectionJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StopPHIDetectionJob", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StopPHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.StopPHIDetectionJobInput) *ComprehendMedicalStopPHIDetectionJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StopPHIDetectionJob", input)
	return &ComprehendMedicalStopPHIDetectionJobFuture{Future: future}
}

func (a *ComprehendMedicalStub) StopRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.StopRxNormInferenceJobInput) (*comprehendmedical.StopRxNormInferenceJobOutput, error) {
	var output comprehendmedical.StopRxNormInferenceJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StopRxNormInferenceJob", input).Get(ctx, &output)
	return &output, err
}

func (a *ComprehendMedicalStub) StopRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StopRxNormInferenceJobInput) *ComprehendMedicalStopRxNormInferenceJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.comprehendmedical.StopRxNormInferenceJob", input)
	return &ComprehendMedicalStopRxNormInferenceJobFuture{Future: future}
}
