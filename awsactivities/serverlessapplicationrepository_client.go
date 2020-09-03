package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"go.temporal.io/sdk/workflow"
)

type ServerlessApplicationRepositoryClient interface {
    CreateApplication(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationRequest) (*serverlessapplicationrepository.CreateApplicationOutput, error)
    CreateApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationRequest) *ServerlessapplicationrepositoryCreateApplicationResult

    CreateApplicationVersion(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error)
    CreateApplicationVersionAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) *ServerlessapplicationrepositoryCreateApplicationVersionResult

    CreateCloudFormationChangeSet(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error)
    CreateCloudFormationChangeSetAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) *ServerlessapplicationrepositoryCreateCloudFormationChangeSetResult

    CreateCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error)
    CreateCloudFormationTemplateAsync(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) *ServerlessapplicationrepositoryCreateCloudFormationTemplateResult

    DeleteApplication(ctx workflow.Context, input *serverlessapplicationrepository.DeleteApplicationInput) (*serverlessapplicationrepository.DeleteApplicationOutput, error)
    DeleteApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.DeleteApplicationInput) *ServerlessapplicationrepositoryDeleteApplicationResult

    GetApplication(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationInput) (*serverlessapplicationrepository.GetApplicationOutput, error)
    GetApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationInput) *ServerlessapplicationrepositoryGetApplicationResult

    GetApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error)
    GetApplicationPolicyAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) *ServerlessapplicationrepositoryGetApplicationPolicyResult

    GetCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error)
    GetCloudFormationTemplateAsync(ctx workflow.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) *ServerlessapplicationrepositoryGetCloudFormationTemplateResult

    ListApplicationDependencies(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error)
    ListApplicationDependenciesAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) *ServerlessapplicationrepositoryListApplicationDependenciesResult

    ListApplicationVersions(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error)
    ListApplicationVersionsAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) *ServerlessapplicationrepositoryListApplicationVersionsResult

    ListApplications(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationsInput) (*serverlessapplicationrepository.ListApplicationsOutput, error)
    ListApplicationsAsync(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationsInput) *ServerlessapplicationrepositoryListApplicationsResult

    PutApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error)
    PutApplicationPolicyAsync(ctx workflow.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) *ServerlessapplicationrepositoryPutApplicationPolicyResult

    UnshareApplication(ctx workflow.Context, input *serverlessapplicationrepository.UnshareApplicationInput) (*serverlessapplicationrepository.UnshareApplicationOutput, error)
    UnshareApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.UnshareApplicationInput) *ServerlessapplicationrepositoryUnshareApplicationResult

    UpdateApplication(ctx workflow.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) (*serverlessapplicationrepository.UpdateApplicationOutput, error)
    UpdateApplicationAsync(ctx workflow.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) *ServerlessapplicationrepositoryUpdateApplicationResult
}
type ServerlessapplicationrepositoryCreateApplicationResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryCreateApplicationResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
    var output serverlessapplicationrepository.CreateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryCreateApplicationVersionResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryCreateApplicationVersionResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error) {
    var output serverlessapplicationrepository.CreateApplicationVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryCreateCloudFormationChangeSetResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryCreateCloudFormationChangeSetResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error) {
    var output serverlessapplicationrepository.CreateCloudFormationChangeSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryCreateCloudFormationTemplateResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryCreateCloudFormationTemplateResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error) {
    var output serverlessapplicationrepository.CreateCloudFormationTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryDeleteApplicationResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryDeleteApplicationResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.DeleteApplicationOutput, error) {
    var output serverlessapplicationrepository.DeleteApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryGetApplicationResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryGetApplicationResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.GetApplicationOutput, error) {
    var output serverlessapplicationrepository.GetApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryGetApplicationPolicyResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryGetApplicationPolicyResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
    var output serverlessapplicationrepository.GetApplicationPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryGetCloudFormationTemplateResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryGetCloudFormationTemplateResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
    var output serverlessapplicationrepository.GetCloudFormationTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryListApplicationDependenciesResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryListApplicationDependenciesResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
    var output serverlessapplicationrepository.ListApplicationDependenciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryListApplicationVersionsResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryListApplicationVersionsResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
    var output serverlessapplicationrepository.ListApplicationVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryListApplicationsResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryListApplicationsResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
    var output serverlessapplicationrepository.ListApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryPutApplicationPolicyResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryPutApplicationPolicyResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error) {
    var output serverlessapplicationrepository.PutApplicationPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryUnshareApplicationResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryUnshareApplicationResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.UnshareApplicationOutput, error) {
    var output serverlessapplicationrepository.UnshareApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServerlessapplicationrepositoryUpdateApplicationResult struct {
	Result workflow.Future
}

func (r *ServerlessapplicationrepositoryUpdateApplicationResult) Get(ctx workflow.Context) (*serverlessapplicationrepository.UpdateApplicationOutput, error) {
    var output serverlessapplicationrepository.UpdateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type ServerlessApplicationRepositoryStub struct {
    activities ServerlessApplicationRepositoryClient
}

func NewServerlessApplicationRepositoryStub() ServerlessApplicationRepositoryClient {
    return &ServerlessApplicationRepositoryStub{}
}

func (a *ServerlessApplicationRepositoryStub) CreateApplication(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationRequest) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
    var output serverlessapplicationrepository.CreateApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) CreateApplicationVersion(ctx workflow.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error) {
    var output serverlessapplicationrepository.CreateApplicationVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApplicationVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) CreateCloudFormationChangeSet(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error) {
    var output serverlessapplicationrepository.CreateCloudFormationChangeSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCloudFormationChangeSet, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) CreateCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error) {
    var output serverlessapplicationrepository.CreateCloudFormationTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCloudFormationTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) DeleteApplication(ctx workflow.Context, input *serverlessapplicationrepository.DeleteApplicationInput) (*serverlessapplicationrepository.DeleteApplicationOutput, error) {
    var output serverlessapplicationrepository.DeleteApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) GetApplication(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationInput) (*serverlessapplicationrepository.GetApplicationOutput, error) {
    var output serverlessapplicationrepository.GetApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) GetApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
    var output serverlessapplicationrepository.GetApplicationPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApplicationPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) GetCloudFormationTemplate(ctx workflow.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
    var output serverlessapplicationrepository.GetCloudFormationTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCloudFormationTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) ListApplicationDependencies(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
    var output serverlessapplicationrepository.ListApplicationDependenciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListApplicationDependencies, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) ListApplicationVersions(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
    var output serverlessapplicationrepository.ListApplicationVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListApplicationVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) ListApplications(ctx workflow.Context, input *serverlessapplicationrepository.ListApplicationsInput) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
    var output serverlessapplicationrepository.ListApplicationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListApplications, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) PutApplicationPolicy(ctx workflow.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error) {
    var output serverlessapplicationrepository.PutApplicationPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutApplicationPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) UnshareApplication(ctx workflow.Context, input *serverlessapplicationrepository.UnshareApplicationInput) (*serverlessapplicationrepository.UnshareApplicationOutput, error) {
    var output serverlessapplicationrepository.UnshareApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UnshareApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *ServerlessApplicationRepositoryStub) UpdateApplication(ctx workflow.Context, input *serverlessapplicationrepository.UpdateApplicationRequest) (*serverlessapplicationrepository.UpdateApplicationOutput, error) {
    var output serverlessapplicationrepository.UpdateApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApplication, input).Get(ctx, &output)
    return &output, err
}
