package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/signer"
	"go.temporal.io/sdk/workflow"
)

type SignerClient interface {
    CancelSigningProfile(ctx workflow.Context, input *signer.CancelSigningProfileInput) (*signer.CancelSigningProfileOutput, error)
    CancelSigningProfileAsync(ctx workflow.Context, input *signer.CancelSigningProfileInput) *SignerCancelSigningProfileResult

    DescribeSigningJob(ctx workflow.Context, input *signer.DescribeSigningJobInput) (*signer.DescribeSigningJobOutput, error)
    DescribeSigningJobAsync(ctx workflow.Context, input *signer.DescribeSigningJobInput) *SignerDescribeSigningJobResult

    GetSigningPlatform(ctx workflow.Context, input *signer.GetSigningPlatformInput) (*signer.GetSigningPlatformOutput, error)
    GetSigningPlatformAsync(ctx workflow.Context, input *signer.GetSigningPlatformInput) *SignerGetSigningPlatformResult

    GetSigningProfile(ctx workflow.Context, input *signer.GetSigningProfileInput) (*signer.GetSigningProfileOutput, error)
    GetSigningProfileAsync(ctx workflow.Context, input *signer.GetSigningProfileInput) *SignerGetSigningProfileResult

    ListSigningJobs(ctx workflow.Context, input *signer.ListSigningJobsInput) (*signer.ListSigningJobsOutput, error)
    ListSigningJobsAsync(ctx workflow.Context, input *signer.ListSigningJobsInput) *SignerListSigningJobsResult

    ListSigningPlatforms(ctx workflow.Context, input *signer.ListSigningPlatformsInput) (*signer.ListSigningPlatformsOutput, error)
    ListSigningPlatformsAsync(ctx workflow.Context, input *signer.ListSigningPlatformsInput) *SignerListSigningPlatformsResult

    ListSigningProfiles(ctx workflow.Context, input *signer.ListSigningProfilesInput) (*signer.ListSigningProfilesOutput, error)
    ListSigningProfilesAsync(ctx workflow.Context, input *signer.ListSigningProfilesInput) *SignerListSigningProfilesResult

    ListTagsForResource(ctx workflow.Context, input *signer.ListTagsForResourceInput) (*signer.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *signer.ListTagsForResourceInput) *SignerListTagsForResourceResult

    PutSigningProfile(ctx workflow.Context, input *signer.PutSigningProfileInput) (*signer.PutSigningProfileOutput, error)
    PutSigningProfileAsync(ctx workflow.Context, input *signer.PutSigningProfileInput) *SignerPutSigningProfileResult

    StartSigningJob(ctx workflow.Context, input *signer.StartSigningJobInput) (*signer.StartSigningJobOutput, error)
    StartSigningJobAsync(ctx workflow.Context, input *signer.StartSigningJobInput) *SignerStartSigningJobResult

    TagResource(ctx workflow.Context, input *signer.TagResourceInput) (*signer.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *signer.TagResourceInput) *SignerTagResourceResult

    UntagResource(ctx workflow.Context, input *signer.UntagResourceInput) (*signer.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *signer.UntagResourceInput) *SignerUntagResourceResult

    WaitUntilSuccessfulSigningJob(ctx workflow.Context, input *signer.DescribeSigningJobInput) error}
type SignerCancelSigningProfileResult struct {
	Result workflow.Future
}

func (r *SignerCancelSigningProfileResult) Get(ctx workflow.Context) (*signer.CancelSigningProfileOutput, error) {
    var output signer.CancelSigningProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SignerDescribeSigningJobResult struct {
	Result workflow.Future
}

func (r *SignerDescribeSigningJobResult) Get(ctx workflow.Context) (*signer.DescribeSigningJobOutput, error) {
    var output signer.DescribeSigningJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SignerGetSigningPlatformResult struct {
	Result workflow.Future
}

func (r *SignerGetSigningPlatformResult) Get(ctx workflow.Context) (*signer.GetSigningPlatformOutput, error) {
    var output signer.GetSigningPlatformOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SignerGetSigningProfileResult struct {
	Result workflow.Future
}

func (r *SignerGetSigningProfileResult) Get(ctx workflow.Context) (*signer.GetSigningProfileOutput, error) {
    var output signer.GetSigningProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SignerListSigningJobsResult struct {
	Result workflow.Future
}

func (r *SignerListSigningJobsResult) Get(ctx workflow.Context) (*signer.ListSigningJobsOutput, error) {
    var output signer.ListSigningJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SignerListSigningPlatformsResult struct {
	Result workflow.Future
}

func (r *SignerListSigningPlatformsResult) Get(ctx workflow.Context) (*signer.ListSigningPlatformsOutput, error) {
    var output signer.ListSigningPlatformsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SignerListSigningProfilesResult struct {
	Result workflow.Future
}

func (r *SignerListSigningProfilesResult) Get(ctx workflow.Context) (*signer.ListSigningProfilesOutput, error) {
    var output signer.ListSigningProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SignerListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *SignerListTagsForResourceResult) Get(ctx workflow.Context) (*signer.ListTagsForResourceOutput, error) {
    var output signer.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SignerPutSigningProfileResult struct {
	Result workflow.Future
}

func (r *SignerPutSigningProfileResult) Get(ctx workflow.Context) (*signer.PutSigningProfileOutput, error) {
    var output signer.PutSigningProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SignerStartSigningJobResult struct {
	Result workflow.Future
}

func (r *SignerStartSigningJobResult) Get(ctx workflow.Context) (*signer.StartSigningJobOutput, error) {
    var output signer.StartSigningJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SignerTagResourceResult struct {
	Result workflow.Future
}

func (r *SignerTagResourceResult) Get(ctx workflow.Context) (*signer.TagResourceOutput, error) {
    var output signer.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SignerUntagResourceResult struct {
	Result workflow.Future
}

func (r *SignerUntagResourceResult) Get(ctx workflow.Context) (*signer.UntagResourceOutput, error) {
    var output signer.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type SignerStub struct {
    activities SignerClient
}

func NewSignerStub() SignerClient {
    return &SignerStub{}
}

func (a *SignerStub) CancelSigningProfile(ctx workflow.Context, input *signer.CancelSigningProfileInput) (*signer.CancelSigningProfileOutput, error) {
    var output signer.CancelSigningProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelSigningProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *SignerStub) DescribeSigningJob(ctx workflow.Context, input *signer.DescribeSigningJobInput) (*signer.DescribeSigningJobOutput, error) {
    var output signer.DescribeSigningJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSigningJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SignerStub) GetSigningPlatform(ctx workflow.Context, input *signer.GetSigningPlatformInput) (*signer.GetSigningPlatformOutput, error) {
    var output signer.GetSigningPlatformOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSigningPlatform, input).Get(ctx, &output)
    return &output, err
}

func (a *SignerStub) GetSigningProfile(ctx workflow.Context, input *signer.GetSigningProfileInput) (*signer.GetSigningProfileOutput, error) {
    var output signer.GetSigningProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSigningProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *SignerStub) ListSigningJobs(ctx workflow.Context, input *signer.ListSigningJobsInput) (*signer.ListSigningJobsOutput, error) {
    var output signer.ListSigningJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSigningJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *SignerStub) ListSigningPlatforms(ctx workflow.Context, input *signer.ListSigningPlatformsInput) (*signer.ListSigningPlatformsOutput, error) {
    var output signer.ListSigningPlatformsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSigningPlatforms, input).Get(ctx, &output)
    return &output, err
}

func (a *SignerStub) ListSigningProfiles(ctx workflow.Context, input *signer.ListSigningProfilesInput) (*signer.ListSigningProfilesOutput, error) {
    var output signer.ListSigningProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSigningProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *SignerStub) ListTagsForResource(ctx workflow.Context, input *signer.ListTagsForResourceInput) (*signer.ListTagsForResourceOutput, error) {
    var output signer.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SignerStub) PutSigningProfile(ctx workflow.Context, input *signer.PutSigningProfileInput) (*signer.PutSigningProfileOutput, error) {
    var output signer.PutSigningProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutSigningProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *SignerStub) StartSigningJob(ctx workflow.Context, input *signer.StartSigningJobInput) (*signer.StartSigningJobOutput, error) {
    var output signer.StartSigningJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartSigningJob, input).Get(ctx, &output)
    return &output, err
}

func (a *SignerStub) TagResource(ctx workflow.Context, input *signer.TagResourceInput) (*signer.TagResourceOutput, error) {
    var output signer.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SignerStub) UntagResource(ctx workflow.Context, input *signer.UntagResourceInput) (*signer.UntagResourceOutput, error) {
    var output signer.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}


func (a *SignerStub) WaitUntilSuccessfulSigningJob(ctx workflow.Context, input *signer.DescribeSigningJobInput) error {
    return a.client.WaitUntilSuccessfulSigningJob(input)
}
