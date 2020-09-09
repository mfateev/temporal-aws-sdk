package awsclients

import (
	"github.com/aws/aws-sdk-go/service/polly"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type PollyClient interface {
       DeleteLexicon(ctx workflow.Context, input *polly.DeleteLexiconInput) (*polly.DeleteLexiconOutput, error)
       DeleteLexiconAsync(ctx workflow.Context, input *polly.DeleteLexiconInput) *PollyDeleteLexiconResult

       DescribeVoices(ctx workflow.Context, input *polly.DescribeVoicesInput) (*polly.DescribeVoicesOutput, error)
       DescribeVoicesAsync(ctx workflow.Context, input *polly.DescribeVoicesInput) *PollyDescribeVoicesResult

       GetLexicon(ctx workflow.Context, input *polly.GetLexiconInput) (*polly.GetLexiconOutput, error)
       GetLexiconAsync(ctx workflow.Context, input *polly.GetLexiconInput) *PollyGetLexiconResult

       GetSpeechSynthesisTask(ctx workflow.Context, input *polly.GetSpeechSynthesisTaskInput) (*polly.GetSpeechSynthesisTaskOutput, error)
       GetSpeechSynthesisTaskAsync(ctx workflow.Context, input *polly.GetSpeechSynthesisTaskInput) *PollyGetSpeechSynthesisTaskResult

       ListLexicons(ctx workflow.Context, input *polly.ListLexiconsInput) (*polly.ListLexiconsOutput, error)
       ListLexiconsAsync(ctx workflow.Context, input *polly.ListLexiconsInput) *PollyListLexiconsResult

       ListSpeechSynthesisTasks(ctx workflow.Context, input *polly.ListSpeechSynthesisTasksInput) (*polly.ListSpeechSynthesisTasksOutput, error)
       ListSpeechSynthesisTasksAsync(ctx workflow.Context, input *polly.ListSpeechSynthesisTasksInput) *PollyListSpeechSynthesisTasksResult

       PutLexicon(ctx workflow.Context, input *polly.PutLexiconInput) (*polly.PutLexiconOutput, error)
       PutLexiconAsync(ctx workflow.Context, input *polly.PutLexiconInput) *PollyPutLexiconResult

       StartSpeechSynthesisTask(ctx workflow.Context, input *polly.StartSpeechSynthesisTaskInput) (*polly.StartSpeechSynthesisTaskOutput, error)
       StartSpeechSynthesisTaskAsync(ctx workflow.Context, input *polly.StartSpeechSynthesisTaskInput) *PollyStartSpeechSynthesisTaskResult

       SynthesizeSpeech(ctx workflow.Context, input *polly.SynthesizeSpeechInput) (*polly.SynthesizeSpeechOutput, error)
       SynthesizeSpeechAsync(ctx workflow.Context, input *polly.SynthesizeSpeechInput) *PollySynthesizeSpeechResult
}

type PollyDeleteLexiconResult struct {
	Result workflow.Future
}

func (r *PollyDeleteLexiconResult) Get(ctx workflow.Context) (*polly.DeleteLexiconOutput, error) {
    var output polly.DeleteLexiconOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PollyDescribeVoicesResult struct {
	Result workflow.Future
}

func (r *PollyDescribeVoicesResult) Get(ctx workflow.Context) (*polly.DescribeVoicesOutput, error) {
    var output polly.DescribeVoicesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PollyGetLexiconResult struct {
	Result workflow.Future
}

func (r *PollyGetLexiconResult) Get(ctx workflow.Context) (*polly.GetLexiconOutput, error) {
    var output polly.GetLexiconOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PollyGetSpeechSynthesisTaskResult struct {
	Result workflow.Future
}

func (r *PollyGetSpeechSynthesisTaskResult) Get(ctx workflow.Context) (*polly.GetSpeechSynthesisTaskOutput, error) {
    var output polly.GetSpeechSynthesisTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PollyListLexiconsResult struct {
	Result workflow.Future
}

func (r *PollyListLexiconsResult) Get(ctx workflow.Context) (*polly.ListLexiconsOutput, error) {
    var output polly.ListLexiconsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PollyListSpeechSynthesisTasksResult struct {
	Result workflow.Future
}

func (r *PollyListSpeechSynthesisTasksResult) Get(ctx workflow.Context) (*polly.ListSpeechSynthesisTasksOutput, error) {
    var output polly.ListSpeechSynthesisTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PollyPutLexiconResult struct {
	Result workflow.Future
}

func (r *PollyPutLexiconResult) Get(ctx workflow.Context) (*polly.PutLexiconOutput, error) {
    var output polly.PutLexiconOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PollyStartSpeechSynthesisTaskResult struct {
	Result workflow.Future
}

func (r *PollyStartSpeechSynthesisTaskResult) Get(ctx workflow.Context) (*polly.StartSpeechSynthesisTaskOutput, error) {
    var output polly.StartSpeechSynthesisTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PollySynthesizeSpeechResult struct {
	Result workflow.Future
}

func (r *PollySynthesizeSpeechResult) Get(ctx workflow.Context) (*polly.SynthesizeSpeechOutput, error) {
    var output polly.SynthesizeSpeechOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type PollyStub struct {
    activities awsactivities.PollyActivities
}

func NewPollyStub() PollyClient {
    return &PollyStub{}
}

func (a *PollyStub) DeleteLexicon(ctx workflow.Context, input *polly.DeleteLexiconInput) (*polly.DeleteLexiconOutput, error) {
    var output polly.DeleteLexiconOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLexicon, input).Get(ctx, &output)
    return &output, err
}

func (a *PollyStub) DeleteLexiconAsync(ctx workflow.Context, input *polly.DeleteLexiconInput) *PollyDeleteLexiconResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLexicon, input)
    return &PollyDeleteLexiconResult{Result: future}
}

func (a *PollyStub) DescribeVoices(ctx workflow.Context, input *polly.DescribeVoicesInput) (*polly.DescribeVoicesOutput, error) {
    var output polly.DescribeVoicesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVoices, input).Get(ctx, &output)
    return &output, err
}

func (a *PollyStub) DescribeVoicesAsync(ctx workflow.Context, input *polly.DescribeVoicesInput) *PollyDescribeVoicesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVoices, input)
    return &PollyDescribeVoicesResult{Result: future}
}

func (a *PollyStub) GetLexicon(ctx workflow.Context, input *polly.GetLexiconInput) (*polly.GetLexiconOutput, error) {
    var output polly.GetLexiconOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLexicon, input).Get(ctx, &output)
    return &output, err
}

func (a *PollyStub) GetLexiconAsync(ctx workflow.Context, input *polly.GetLexiconInput) *PollyGetLexiconResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetLexicon, input)
    return &PollyGetLexiconResult{Result: future}
}

func (a *PollyStub) GetSpeechSynthesisTask(ctx workflow.Context, input *polly.GetSpeechSynthesisTaskInput) (*polly.GetSpeechSynthesisTaskOutput, error) {
    var output polly.GetSpeechSynthesisTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSpeechSynthesisTask, input).Get(ctx, &output)
    return &output, err
}

func (a *PollyStub) GetSpeechSynthesisTaskAsync(ctx workflow.Context, input *polly.GetSpeechSynthesisTaskInput) *PollyGetSpeechSynthesisTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSpeechSynthesisTask, input)
    return &PollyGetSpeechSynthesisTaskResult{Result: future}
}

func (a *PollyStub) ListLexicons(ctx workflow.Context, input *polly.ListLexiconsInput) (*polly.ListLexiconsOutput, error) {
    var output polly.ListLexiconsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLexicons, input).Get(ctx, &output)
    return &output, err
}

func (a *PollyStub) ListLexiconsAsync(ctx workflow.Context, input *polly.ListLexiconsInput) *PollyListLexiconsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListLexicons, input)
    return &PollyListLexiconsResult{Result: future}
}

func (a *PollyStub) ListSpeechSynthesisTasks(ctx workflow.Context, input *polly.ListSpeechSynthesisTasksInput) (*polly.ListSpeechSynthesisTasksOutput, error) {
    var output polly.ListSpeechSynthesisTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSpeechSynthesisTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *PollyStub) ListSpeechSynthesisTasksAsync(ctx workflow.Context, input *polly.ListSpeechSynthesisTasksInput) *PollyListSpeechSynthesisTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSpeechSynthesisTasks, input)
    return &PollyListSpeechSynthesisTasksResult{Result: future}
}

func (a *PollyStub) PutLexicon(ctx workflow.Context, input *polly.PutLexiconInput) (*polly.PutLexiconOutput, error) {
    var output polly.PutLexiconOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutLexicon, input).Get(ctx, &output)
    return &output, err
}

func (a *PollyStub) PutLexiconAsync(ctx workflow.Context, input *polly.PutLexiconInput) *PollyPutLexiconResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutLexicon, input)
    return &PollyPutLexiconResult{Result: future}
}

func (a *PollyStub) StartSpeechSynthesisTask(ctx workflow.Context, input *polly.StartSpeechSynthesisTaskInput) (*polly.StartSpeechSynthesisTaskOutput, error) {
    var output polly.StartSpeechSynthesisTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartSpeechSynthesisTask, input).Get(ctx, &output)
    return &output, err
}

func (a *PollyStub) StartSpeechSynthesisTaskAsync(ctx workflow.Context, input *polly.StartSpeechSynthesisTaskInput) *PollyStartSpeechSynthesisTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartSpeechSynthesisTask, input)
    return &PollyStartSpeechSynthesisTaskResult{Result: future}
}

func (a *PollyStub) SynthesizeSpeech(ctx workflow.Context, input *polly.SynthesizeSpeechInput) (*polly.SynthesizeSpeechOutput, error) {
    var output polly.SynthesizeSpeechOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SynthesizeSpeech, input).Get(ctx, &output)
    return &output, err
}

func (a *PollyStub) SynthesizeSpeechAsync(ctx workflow.Context, input *polly.SynthesizeSpeechInput) *PollySynthesizeSpeechResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SynthesizeSpeech, input)
    return &PollySynthesizeSpeechResult{Result: future}
}