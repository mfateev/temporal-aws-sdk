package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/robomaker"
	"go.temporal.io/sdk/workflow"
)

type RoboMakerClient interface {
    BatchDeleteWorlds(ctx workflow.Context, input *robomaker.BatchDeleteWorldsInput) (*robomaker.BatchDeleteWorldsOutput, error)
    BatchDeleteWorldsAsync(ctx workflow.Context, input *robomaker.BatchDeleteWorldsInput) *RobomakerBatchDeleteWorldsResult

    BatchDescribeSimulationJob(ctx workflow.Context, input *robomaker.BatchDescribeSimulationJobInput) (*robomaker.BatchDescribeSimulationJobOutput, error)
    BatchDescribeSimulationJobAsync(ctx workflow.Context, input *robomaker.BatchDescribeSimulationJobInput) *RobomakerBatchDescribeSimulationJobResult

    CancelDeploymentJob(ctx workflow.Context, input *robomaker.CancelDeploymentJobInput) (*robomaker.CancelDeploymentJobOutput, error)
    CancelDeploymentJobAsync(ctx workflow.Context, input *robomaker.CancelDeploymentJobInput) *RobomakerCancelDeploymentJobResult

    CancelSimulationJob(ctx workflow.Context, input *robomaker.CancelSimulationJobInput) (*robomaker.CancelSimulationJobOutput, error)
    CancelSimulationJobAsync(ctx workflow.Context, input *robomaker.CancelSimulationJobInput) *RobomakerCancelSimulationJobResult

    CancelSimulationJobBatch(ctx workflow.Context, input *robomaker.CancelSimulationJobBatchInput) (*robomaker.CancelSimulationJobBatchOutput, error)
    CancelSimulationJobBatchAsync(ctx workflow.Context, input *robomaker.CancelSimulationJobBatchInput) *RobomakerCancelSimulationJobBatchResult

    CancelWorldExportJob(ctx workflow.Context, input *robomaker.CancelWorldExportJobInput) (*robomaker.CancelWorldExportJobOutput, error)
    CancelWorldExportJobAsync(ctx workflow.Context, input *robomaker.CancelWorldExportJobInput) *RobomakerCancelWorldExportJobResult

    CancelWorldGenerationJob(ctx workflow.Context, input *robomaker.CancelWorldGenerationJobInput) (*robomaker.CancelWorldGenerationJobOutput, error)
    CancelWorldGenerationJobAsync(ctx workflow.Context, input *robomaker.CancelWorldGenerationJobInput) *RobomakerCancelWorldGenerationJobResult

    CreateDeploymentJob(ctx workflow.Context, input *robomaker.CreateDeploymentJobInput) (*robomaker.CreateDeploymentJobOutput, error)
    CreateDeploymentJobAsync(ctx workflow.Context, input *robomaker.CreateDeploymentJobInput) *RobomakerCreateDeploymentJobResult

    CreateFleet(ctx workflow.Context, input *robomaker.CreateFleetInput) (*robomaker.CreateFleetOutput, error)
    CreateFleetAsync(ctx workflow.Context, input *robomaker.CreateFleetInput) *RobomakerCreateFleetResult

    CreateRobot(ctx workflow.Context, input *robomaker.CreateRobotInput) (*robomaker.CreateRobotOutput, error)
    CreateRobotAsync(ctx workflow.Context, input *robomaker.CreateRobotInput) *RobomakerCreateRobotResult

    CreateRobotApplication(ctx workflow.Context, input *robomaker.CreateRobotApplicationInput) (*robomaker.CreateRobotApplicationOutput, error)
    CreateRobotApplicationAsync(ctx workflow.Context, input *robomaker.CreateRobotApplicationInput) *RobomakerCreateRobotApplicationResult

    CreateRobotApplicationVersion(ctx workflow.Context, input *robomaker.CreateRobotApplicationVersionInput) (*robomaker.CreateRobotApplicationVersionOutput, error)
    CreateRobotApplicationVersionAsync(ctx workflow.Context, input *robomaker.CreateRobotApplicationVersionInput) *RobomakerCreateRobotApplicationVersionResult

    CreateSimulationApplication(ctx workflow.Context, input *robomaker.CreateSimulationApplicationInput) (*robomaker.CreateSimulationApplicationOutput, error)
    CreateSimulationApplicationAsync(ctx workflow.Context, input *robomaker.CreateSimulationApplicationInput) *RobomakerCreateSimulationApplicationResult

    CreateSimulationApplicationVersion(ctx workflow.Context, input *robomaker.CreateSimulationApplicationVersionInput) (*robomaker.CreateSimulationApplicationVersionOutput, error)
    CreateSimulationApplicationVersionAsync(ctx workflow.Context, input *robomaker.CreateSimulationApplicationVersionInput) *RobomakerCreateSimulationApplicationVersionResult

    CreateSimulationJob(ctx workflow.Context, input *robomaker.CreateSimulationJobInput) (*robomaker.CreateSimulationJobOutput, error)
    CreateSimulationJobAsync(ctx workflow.Context, input *robomaker.CreateSimulationJobInput) *RobomakerCreateSimulationJobResult

    CreateWorldExportJob(ctx workflow.Context, input *robomaker.CreateWorldExportJobInput) (*robomaker.CreateWorldExportJobOutput, error)
    CreateWorldExportJobAsync(ctx workflow.Context, input *robomaker.CreateWorldExportJobInput) *RobomakerCreateWorldExportJobResult

    CreateWorldGenerationJob(ctx workflow.Context, input *robomaker.CreateWorldGenerationJobInput) (*robomaker.CreateWorldGenerationJobOutput, error)
    CreateWorldGenerationJobAsync(ctx workflow.Context, input *robomaker.CreateWorldGenerationJobInput) *RobomakerCreateWorldGenerationJobResult

    CreateWorldTemplate(ctx workflow.Context, input *robomaker.CreateWorldTemplateInput) (*robomaker.CreateWorldTemplateOutput, error)
    CreateWorldTemplateAsync(ctx workflow.Context, input *robomaker.CreateWorldTemplateInput) *RobomakerCreateWorldTemplateResult

    DeleteFleet(ctx workflow.Context, input *robomaker.DeleteFleetInput) (*robomaker.DeleteFleetOutput, error)
    DeleteFleetAsync(ctx workflow.Context, input *robomaker.DeleteFleetInput) *RobomakerDeleteFleetResult

    DeleteRobot(ctx workflow.Context, input *robomaker.DeleteRobotInput) (*robomaker.DeleteRobotOutput, error)
    DeleteRobotAsync(ctx workflow.Context, input *robomaker.DeleteRobotInput) *RobomakerDeleteRobotResult

    DeleteRobotApplication(ctx workflow.Context, input *robomaker.DeleteRobotApplicationInput) (*robomaker.DeleteRobotApplicationOutput, error)
    DeleteRobotApplicationAsync(ctx workflow.Context, input *robomaker.DeleteRobotApplicationInput) *RobomakerDeleteRobotApplicationResult

    DeleteSimulationApplication(ctx workflow.Context, input *robomaker.DeleteSimulationApplicationInput) (*robomaker.DeleteSimulationApplicationOutput, error)
    DeleteSimulationApplicationAsync(ctx workflow.Context, input *robomaker.DeleteSimulationApplicationInput) *RobomakerDeleteSimulationApplicationResult

    DeleteWorldTemplate(ctx workflow.Context, input *robomaker.DeleteWorldTemplateInput) (*robomaker.DeleteWorldTemplateOutput, error)
    DeleteWorldTemplateAsync(ctx workflow.Context, input *robomaker.DeleteWorldTemplateInput) *RobomakerDeleteWorldTemplateResult

    DeregisterRobot(ctx workflow.Context, input *robomaker.DeregisterRobotInput) (*robomaker.DeregisterRobotOutput, error)
    DeregisterRobotAsync(ctx workflow.Context, input *robomaker.DeregisterRobotInput) *RobomakerDeregisterRobotResult

    DescribeDeploymentJob(ctx workflow.Context, input *robomaker.DescribeDeploymentJobInput) (*robomaker.DescribeDeploymentJobOutput, error)
    DescribeDeploymentJobAsync(ctx workflow.Context, input *robomaker.DescribeDeploymentJobInput) *RobomakerDescribeDeploymentJobResult

    DescribeFleet(ctx workflow.Context, input *robomaker.DescribeFleetInput) (*robomaker.DescribeFleetOutput, error)
    DescribeFleetAsync(ctx workflow.Context, input *robomaker.DescribeFleetInput) *RobomakerDescribeFleetResult

    DescribeRobot(ctx workflow.Context, input *robomaker.DescribeRobotInput) (*robomaker.DescribeRobotOutput, error)
    DescribeRobotAsync(ctx workflow.Context, input *robomaker.DescribeRobotInput) *RobomakerDescribeRobotResult

    DescribeRobotApplication(ctx workflow.Context, input *robomaker.DescribeRobotApplicationInput) (*robomaker.DescribeRobotApplicationOutput, error)
    DescribeRobotApplicationAsync(ctx workflow.Context, input *robomaker.DescribeRobotApplicationInput) *RobomakerDescribeRobotApplicationResult

    DescribeSimulationApplication(ctx workflow.Context, input *robomaker.DescribeSimulationApplicationInput) (*robomaker.DescribeSimulationApplicationOutput, error)
    DescribeSimulationApplicationAsync(ctx workflow.Context, input *robomaker.DescribeSimulationApplicationInput) *RobomakerDescribeSimulationApplicationResult

    DescribeSimulationJob(ctx workflow.Context, input *robomaker.DescribeSimulationJobInput) (*robomaker.DescribeSimulationJobOutput, error)
    DescribeSimulationJobAsync(ctx workflow.Context, input *robomaker.DescribeSimulationJobInput) *RobomakerDescribeSimulationJobResult

    DescribeSimulationJobBatch(ctx workflow.Context, input *robomaker.DescribeSimulationJobBatchInput) (*robomaker.DescribeSimulationJobBatchOutput, error)
    DescribeSimulationJobBatchAsync(ctx workflow.Context, input *robomaker.DescribeSimulationJobBatchInput) *RobomakerDescribeSimulationJobBatchResult

    DescribeWorld(ctx workflow.Context, input *robomaker.DescribeWorldInput) (*robomaker.DescribeWorldOutput, error)
    DescribeWorldAsync(ctx workflow.Context, input *robomaker.DescribeWorldInput) *RobomakerDescribeWorldResult

    DescribeWorldExportJob(ctx workflow.Context, input *robomaker.DescribeWorldExportJobInput) (*robomaker.DescribeWorldExportJobOutput, error)
    DescribeWorldExportJobAsync(ctx workflow.Context, input *robomaker.DescribeWorldExportJobInput) *RobomakerDescribeWorldExportJobResult

    DescribeWorldGenerationJob(ctx workflow.Context, input *robomaker.DescribeWorldGenerationJobInput) (*robomaker.DescribeWorldGenerationJobOutput, error)
    DescribeWorldGenerationJobAsync(ctx workflow.Context, input *robomaker.DescribeWorldGenerationJobInput) *RobomakerDescribeWorldGenerationJobResult

    DescribeWorldTemplate(ctx workflow.Context, input *robomaker.DescribeWorldTemplateInput) (*robomaker.DescribeWorldTemplateOutput, error)
    DescribeWorldTemplateAsync(ctx workflow.Context, input *robomaker.DescribeWorldTemplateInput) *RobomakerDescribeWorldTemplateResult

    GetWorldTemplateBody(ctx workflow.Context, input *robomaker.GetWorldTemplateBodyInput) (*robomaker.GetWorldTemplateBodyOutput, error)
    GetWorldTemplateBodyAsync(ctx workflow.Context, input *robomaker.GetWorldTemplateBodyInput) *RobomakerGetWorldTemplateBodyResult

    ListDeploymentJobs(ctx workflow.Context, input *robomaker.ListDeploymentJobsInput) (*robomaker.ListDeploymentJobsOutput, error)
    ListDeploymentJobsAsync(ctx workflow.Context, input *robomaker.ListDeploymentJobsInput) *RobomakerListDeploymentJobsResult

    ListFleets(ctx workflow.Context, input *robomaker.ListFleetsInput) (*robomaker.ListFleetsOutput, error)
    ListFleetsAsync(ctx workflow.Context, input *robomaker.ListFleetsInput) *RobomakerListFleetsResult

    ListRobotApplications(ctx workflow.Context, input *robomaker.ListRobotApplicationsInput) (*robomaker.ListRobotApplicationsOutput, error)
    ListRobotApplicationsAsync(ctx workflow.Context, input *robomaker.ListRobotApplicationsInput) *RobomakerListRobotApplicationsResult

    ListRobots(ctx workflow.Context, input *robomaker.ListRobotsInput) (*robomaker.ListRobotsOutput, error)
    ListRobotsAsync(ctx workflow.Context, input *robomaker.ListRobotsInput) *RobomakerListRobotsResult

    ListSimulationApplications(ctx workflow.Context, input *robomaker.ListSimulationApplicationsInput) (*robomaker.ListSimulationApplicationsOutput, error)
    ListSimulationApplicationsAsync(ctx workflow.Context, input *robomaker.ListSimulationApplicationsInput) *RobomakerListSimulationApplicationsResult

    ListSimulationJobBatches(ctx workflow.Context, input *robomaker.ListSimulationJobBatchesInput) (*robomaker.ListSimulationJobBatchesOutput, error)
    ListSimulationJobBatchesAsync(ctx workflow.Context, input *robomaker.ListSimulationJobBatchesInput) *RobomakerListSimulationJobBatchesResult

    ListSimulationJobs(ctx workflow.Context, input *robomaker.ListSimulationJobsInput) (*robomaker.ListSimulationJobsOutput, error)
    ListSimulationJobsAsync(ctx workflow.Context, input *robomaker.ListSimulationJobsInput) *RobomakerListSimulationJobsResult

    ListTagsForResource(ctx workflow.Context, input *robomaker.ListTagsForResourceInput) (*robomaker.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *robomaker.ListTagsForResourceInput) *RobomakerListTagsForResourceResult

    ListWorldExportJobs(ctx workflow.Context, input *robomaker.ListWorldExportJobsInput) (*robomaker.ListWorldExportJobsOutput, error)
    ListWorldExportJobsAsync(ctx workflow.Context, input *robomaker.ListWorldExportJobsInput) *RobomakerListWorldExportJobsResult

    ListWorldGenerationJobs(ctx workflow.Context, input *robomaker.ListWorldGenerationJobsInput) (*robomaker.ListWorldGenerationJobsOutput, error)
    ListWorldGenerationJobsAsync(ctx workflow.Context, input *robomaker.ListWorldGenerationJobsInput) *RobomakerListWorldGenerationJobsResult

    ListWorldTemplates(ctx workflow.Context, input *robomaker.ListWorldTemplatesInput) (*robomaker.ListWorldTemplatesOutput, error)
    ListWorldTemplatesAsync(ctx workflow.Context, input *robomaker.ListWorldTemplatesInput) *RobomakerListWorldTemplatesResult

    ListWorlds(ctx workflow.Context, input *robomaker.ListWorldsInput) (*robomaker.ListWorldsOutput, error)
    ListWorldsAsync(ctx workflow.Context, input *robomaker.ListWorldsInput) *RobomakerListWorldsResult

    RegisterRobot(ctx workflow.Context, input *robomaker.RegisterRobotInput) (*robomaker.RegisterRobotOutput, error)
    RegisterRobotAsync(ctx workflow.Context, input *robomaker.RegisterRobotInput) *RobomakerRegisterRobotResult

    RestartSimulationJob(ctx workflow.Context, input *robomaker.RestartSimulationJobInput) (*robomaker.RestartSimulationJobOutput, error)
    RestartSimulationJobAsync(ctx workflow.Context, input *robomaker.RestartSimulationJobInput) *RobomakerRestartSimulationJobResult

    StartSimulationJobBatch(ctx workflow.Context, input *robomaker.StartSimulationJobBatchInput) (*robomaker.StartSimulationJobBatchOutput, error)
    StartSimulationJobBatchAsync(ctx workflow.Context, input *robomaker.StartSimulationJobBatchInput) *RobomakerStartSimulationJobBatchResult

    SyncDeploymentJob(ctx workflow.Context, input *robomaker.SyncDeploymentJobInput) (*robomaker.SyncDeploymentJobOutput, error)
    SyncDeploymentJobAsync(ctx workflow.Context, input *robomaker.SyncDeploymentJobInput) *RobomakerSyncDeploymentJobResult

    TagResource(ctx workflow.Context, input *robomaker.TagResourceInput) (*robomaker.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *robomaker.TagResourceInput) *RobomakerTagResourceResult

    UntagResource(ctx workflow.Context, input *robomaker.UntagResourceInput) (*robomaker.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *robomaker.UntagResourceInput) *RobomakerUntagResourceResult

    UpdateRobotApplication(ctx workflow.Context, input *robomaker.UpdateRobotApplicationInput) (*robomaker.UpdateRobotApplicationOutput, error)
    UpdateRobotApplicationAsync(ctx workflow.Context, input *robomaker.UpdateRobotApplicationInput) *RobomakerUpdateRobotApplicationResult

    UpdateSimulationApplication(ctx workflow.Context, input *robomaker.UpdateSimulationApplicationInput) (*robomaker.UpdateSimulationApplicationOutput, error)
    UpdateSimulationApplicationAsync(ctx workflow.Context, input *robomaker.UpdateSimulationApplicationInput) *RobomakerUpdateSimulationApplicationResult

    UpdateWorldTemplate(ctx workflow.Context, input *robomaker.UpdateWorldTemplateInput) (*robomaker.UpdateWorldTemplateOutput, error)
    UpdateWorldTemplateAsync(ctx workflow.Context, input *robomaker.UpdateWorldTemplateInput) *RobomakerUpdateWorldTemplateResult
}
type RobomakerBatchDeleteWorldsResult struct {
	Result workflow.Future
}

func (r *RobomakerBatchDeleteWorldsResult) Get(ctx workflow.Context) (*robomaker.BatchDeleteWorldsOutput, error) {
    var output robomaker.BatchDeleteWorldsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerBatchDescribeSimulationJobResult struct {
	Result workflow.Future
}

func (r *RobomakerBatchDescribeSimulationJobResult) Get(ctx workflow.Context) (*robomaker.BatchDescribeSimulationJobOutput, error) {
    var output robomaker.BatchDescribeSimulationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCancelDeploymentJobResult struct {
	Result workflow.Future
}

func (r *RobomakerCancelDeploymentJobResult) Get(ctx workflow.Context) (*robomaker.CancelDeploymentJobOutput, error) {
    var output robomaker.CancelDeploymentJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCancelSimulationJobResult struct {
	Result workflow.Future
}

func (r *RobomakerCancelSimulationJobResult) Get(ctx workflow.Context) (*robomaker.CancelSimulationJobOutput, error) {
    var output robomaker.CancelSimulationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCancelSimulationJobBatchResult struct {
	Result workflow.Future
}

func (r *RobomakerCancelSimulationJobBatchResult) Get(ctx workflow.Context) (*robomaker.CancelSimulationJobBatchOutput, error) {
    var output robomaker.CancelSimulationJobBatchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCancelWorldExportJobResult struct {
	Result workflow.Future
}

func (r *RobomakerCancelWorldExportJobResult) Get(ctx workflow.Context) (*robomaker.CancelWorldExportJobOutput, error) {
    var output robomaker.CancelWorldExportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCancelWorldGenerationJobResult struct {
	Result workflow.Future
}

func (r *RobomakerCancelWorldGenerationJobResult) Get(ctx workflow.Context) (*robomaker.CancelWorldGenerationJobOutput, error) {
    var output robomaker.CancelWorldGenerationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCreateDeploymentJobResult struct {
	Result workflow.Future
}

func (r *RobomakerCreateDeploymentJobResult) Get(ctx workflow.Context) (*robomaker.CreateDeploymentJobOutput, error) {
    var output robomaker.CreateDeploymentJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCreateFleetResult struct {
	Result workflow.Future
}

func (r *RobomakerCreateFleetResult) Get(ctx workflow.Context) (*robomaker.CreateFleetOutput, error) {
    var output robomaker.CreateFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCreateRobotResult struct {
	Result workflow.Future
}

func (r *RobomakerCreateRobotResult) Get(ctx workflow.Context) (*robomaker.CreateRobotOutput, error) {
    var output robomaker.CreateRobotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCreateRobotApplicationResult struct {
	Result workflow.Future
}

func (r *RobomakerCreateRobotApplicationResult) Get(ctx workflow.Context) (*robomaker.CreateRobotApplicationOutput, error) {
    var output robomaker.CreateRobotApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCreateRobotApplicationVersionResult struct {
	Result workflow.Future
}

func (r *RobomakerCreateRobotApplicationVersionResult) Get(ctx workflow.Context) (*robomaker.CreateRobotApplicationVersionOutput, error) {
    var output robomaker.CreateRobotApplicationVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCreateSimulationApplicationResult struct {
	Result workflow.Future
}

func (r *RobomakerCreateSimulationApplicationResult) Get(ctx workflow.Context) (*robomaker.CreateSimulationApplicationOutput, error) {
    var output robomaker.CreateSimulationApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCreateSimulationApplicationVersionResult struct {
	Result workflow.Future
}

func (r *RobomakerCreateSimulationApplicationVersionResult) Get(ctx workflow.Context) (*robomaker.CreateSimulationApplicationVersionOutput, error) {
    var output robomaker.CreateSimulationApplicationVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCreateSimulationJobResult struct {
	Result workflow.Future
}

func (r *RobomakerCreateSimulationJobResult) Get(ctx workflow.Context) (*robomaker.CreateSimulationJobOutput, error) {
    var output robomaker.CreateSimulationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCreateWorldExportJobResult struct {
	Result workflow.Future
}

func (r *RobomakerCreateWorldExportJobResult) Get(ctx workflow.Context) (*robomaker.CreateWorldExportJobOutput, error) {
    var output robomaker.CreateWorldExportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCreateWorldGenerationJobResult struct {
	Result workflow.Future
}

func (r *RobomakerCreateWorldGenerationJobResult) Get(ctx workflow.Context) (*robomaker.CreateWorldGenerationJobOutput, error) {
    var output robomaker.CreateWorldGenerationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerCreateWorldTemplateResult struct {
	Result workflow.Future
}

func (r *RobomakerCreateWorldTemplateResult) Get(ctx workflow.Context) (*robomaker.CreateWorldTemplateOutput, error) {
    var output robomaker.CreateWorldTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDeleteFleetResult struct {
	Result workflow.Future
}

func (r *RobomakerDeleteFleetResult) Get(ctx workflow.Context) (*robomaker.DeleteFleetOutput, error) {
    var output robomaker.DeleteFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDeleteRobotResult struct {
	Result workflow.Future
}

func (r *RobomakerDeleteRobotResult) Get(ctx workflow.Context) (*robomaker.DeleteRobotOutput, error) {
    var output robomaker.DeleteRobotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDeleteRobotApplicationResult struct {
	Result workflow.Future
}

func (r *RobomakerDeleteRobotApplicationResult) Get(ctx workflow.Context) (*robomaker.DeleteRobotApplicationOutput, error) {
    var output robomaker.DeleteRobotApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDeleteSimulationApplicationResult struct {
	Result workflow.Future
}

func (r *RobomakerDeleteSimulationApplicationResult) Get(ctx workflow.Context) (*robomaker.DeleteSimulationApplicationOutput, error) {
    var output robomaker.DeleteSimulationApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDeleteWorldTemplateResult struct {
	Result workflow.Future
}

func (r *RobomakerDeleteWorldTemplateResult) Get(ctx workflow.Context) (*robomaker.DeleteWorldTemplateOutput, error) {
    var output robomaker.DeleteWorldTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDeregisterRobotResult struct {
	Result workflow.Future
}

func (r *RobomakerDeregisterRobotResult) Get(ctx workflow.Context) (*robomaker.DeregisterRobotOutput, error) {
    var output robomaker.DeregisterRobotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDescribeDeploymentJobResult struct {
	Result workflow.Future
}

func (r *RobomakerDescribeDeploymentJobResult) Get(ctx workflow.Context) (*robomaker.DescribeDeploymentJobOutput, error) {
    var output robomaker.DescribeDeploymentJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDescribeFleetResult struct {
	Result workflow.Future
}

func (r *RobomakerDescribeFleetResult) Get(ctx workflow.Context) (*robomaker.DescribeFleetOutput, error) {
    var output robomaker.DescribeFleetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDescribeRobotResult struct {
	Result workflow.Future
}

func (r *RobomakerDescribeRobotResult) Get(ctx workflow.Context) (*robomaker.DescribeRobotOutput, error) {
    var output robomaker.DescribeRobotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDescribeRobotApplicationResult struct {
	Result workflow.Future
}

func (r *RobomakerDescribeRobotApplicationResult) Get(ctx workflow.Context) (*robomaker.DescribeRobotApplicationOutput, error) {
    var output robomaker.DescribeRobotApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDescribeSimulationApplicationResult struct {
	Result workflow.Future
}

func (r *RobomakerDescribeSimulationApplicationResult) Get(ctx workflow.Context) (*robomaker.DescribeSimulationApplicationOutput, error) {
    var output robomaker.DescribeSimulationApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDescribeSimulationJobResult struct {
	Result workflow.Future
}

func (r *RobomakerDescribeSimulationJobResult) Get(ctx workflow.Context) (*robomaker.DescribeSimulationJobOutput, error) {
    var output robomaker.DescribeSimulationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDescribeSimulationJobBatchResult struct {
	Result workflow.Future
}

func (r *RobomakerDescribeSimulationJobBatchResult) Get(ctx workflow.Context) (*robomaker.DescribeSimulationJobBatchOutput, error) {
    var output robomaker.DescribeSimulationJobBatchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDescribeWorldResult struct {
	Result workflow.Future
}

func (r *RobomakerDescribeWorldResult) Get(ctx workflow.Context) (*robomaker.DescribeWorldOutput, error) {
    var output robomaker.DescribeWorldOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDescribeWorldExportJobResult struct {
	Result workflow.Future
}

func (r *RobomakerDescribeWorldExportJobResult) Get(ctx workflow.Context) (*robomaker.DescribeWorldExportJobOutput, error) {
    var output robomaker.DescribeWorldExportJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDescribeWorldGenerationJobResult struct {
	Result workflow.Future
}

func (r *RobomakerDescribeWorldGenerationJobResult) Get(ctx workflow.Context) (*robomaker.DescribeWorldGenerationJobOutput, error) {
    var output robomaker.DescribeWorldGenerationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerDescribeWorldTemplateResult struct {
	Result workflow.Future
}

func (r *RobomakerDescribeWorldTemplateResult) Get(ctx workflow.Context) (*robomaker.DescribeWorldTemplateOutput, error) {
    var output robomaker.DescribeWorldTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerGetWorldTemplateBodyResult struct {
	Result workflow.Future
}

func (r *RobomakerGetWorldTemplateBodyResult) Get(ctx workflow.Context) (*robomaker.GetWorldTemplateBodyOutput, error) {
    var output robomaker.GetWorldTemplateBodyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListDeploymentJobsResult struct {
	Result workflow.Future
}

func (r *RobomakerListDeploymentJobsResult) Get(ctx workflow.Context) (*robomaker.ListDeploymentJobsOutput, error) {
    var output robomaker.ListDeploymentJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListFleetsResult struct {
	Result workflow.Future
}

func (r *RobomakerListFleetsResult) Get(ctx workflow.Context) (*robomaker.ListFleetsOutput, error) {
    var output robomaker.ListFleetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListRobotApplicationsResult struct {
	Result workflow.Future
}

func (r *RobomakerListRobotApplicationsResult) Get(ctx workflow.Context) (*robomaker.ListRobotApplicationsOutput, error) {
    var output robomaker.ListRobotApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListRobotsResult struct {
	Result workflow.Future
}

func (r *RobomakerListRobotsResult) Get(ctx workflow.Context) (*robomaker.ListRobotsOutput, error) {
    var output robomaker.ListRobotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListSimulationApplicationsResult struct {
	Result workflow.Future
}

func (r *RobomakerListSimulationApplicationsResult) Get(ctx workflow.Context) (*robomaker.ListSimulationApplicationsOutput, error) {
    var output robomaker.ListSimulationApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListSimulationJobBatchesResult struct {
	Result workflow.Future
}

func (r *RobomakerListSimulationJobBatchesResult) Get(ctx workflow.Context) (*robomaker.ListSimulationJobBatchesOutput, error) {
    var output robomaker.ListSimulationJobBatchesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListSimulationJobsResult struct {
	Result workflow.Future
}

func (r *RobomakerListSimulationJobsResult) Get(ctx workflow.Context) (*robomaker.ListSimulationJobsOutput, error) {
    var output robomaker.ListSimulationJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *RobomakerListTagsForResourceResult) Get(ctx workflow.Context) (*robomaker.ListTagsForResourceOutput, error) {
    var output robomaker.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListWorldExportJobsResult struct {
	Result workflow.Future
}

func (r *RobomakerListWorldExportJobsResult) Get(ctx workflow.Context) (*robomaker.ListWorldExportJobsOutput, error) {
    var output robomaker.ListWorldExportJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListWorldGenerationJobsResult struct {
	Result workflow.Future
}

func (r *RobomakerListWorldGenerationJobsResult) Get(ctx workflow.Context) (*robomaker.ListWorldGenerationJobsOutput, error) {
    var output robomaker.ListWorldGenerationJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListWorldTemplatesResult struct {
	Result workflow.Future
}

func (r *RobomakerListWorldTemplatesResult) Get(ctx workflow.Context) (*robomaker.ListWorldTemplatesOutput, error) {
    var output robomaker.ListWorldTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerListWorldsResult struct {
	Result workflow.Future
}

func (r *RobomakerListWorldsResult) Get(ctx workflow.Context) (*robomaker.ListWorldsOutput, error) {
    var output robomaker.ListWorldsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerRegisterRobotResult struct {
	Result workflow.Future
}

func (r *RobomakerRegisterRobotResult) Get(ctx workflow.Context) (*robomaker.RegisterRobotOutput, error) {
    var output robomaker.RegisterRobotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerRestartSimulationJobResult struct {
	Result workflow.Future
}

func (r *RobomakerRestartSimulationJobResult) Get(ctx workflow.Context) (*robomaker.RestartSimulationJobOutput, error) {
    var output robomaker.RestartSimulationJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerStartSimulationJobBatchResult struct {
	Result workflow.Future
}

func (r *RobomakerStartSimulationJobBatchResult) Get(ctx workflow.Context) (*robomaker.StartSimulationJobBatchOutput, error) {
    var output robomaker.StartSimulationJobBatchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerSyncDeploymentJobResult struct {
	Result workflow.Future
}

func (r *RobomakerSyncDeploymentJobResult) Get(ctx workflow.Context) (*robomaker.SyncDeploymentJobOutput, error) {
    var output robomaker.SyncDeploymentJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerTagResourceResult struct {
	Result workflow.Future
}

func (r *RobomakerTagResourceResult) Get(ctx workflow.Context) (*robomaker.TagResourceOutput, error) {
    var output robomaker.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerUntagResourceResult struct {
	Result workflow.Future
}

func (r *RobomakerUntagResourceResult) Get(ctx workflow.Context) (*robomaker.UntagResourceOutput, error) {
    var output robomaker.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerUpdateRobotApplicationResult struct {
	Result workflow.Future
}

func (r *RobomakerUpdateRobotApplicationResult) Get(ctx workflow.Context) (*robomaker.UpdateRobotApplicationOutput, error) {
    var output robomaker.UpdateRobotApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerUpdateSimulationApplicationResult struct {
	Result workflow.Future
}

func (r *RobomakerUpdateSimulationApplicationResult) Get(ctx workflow.Context) (*robomaker.UpdateSimulationApplicationOutput, error) {
    var output robomaker.UpdateSimulationApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RobomakerUpdateWorldTemplateResult struct {
	Result workflow.Future
}

func (r *RobomakerUpdateWorldTemplateResult) Get(ctx workflow.Context) (*robomaker.UpdateWorldTemplateOutput, error) {
    var output robomaker.UpdateWorldTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type RoboMakerStub struct {
    activities RoboMakerClient
}

func NewRoboMakerStub() RoboMakerClient {
    return &RoboMakerStub{}
}

func (a *RoboMakerStub) BatchDeleteWorlds(ctx workflow.Context, input *robomaker.BatchDeleteWorldsInput) (*robomaker.BatchDeleteWorldsOutput, error) {
    var output robomaker.BatchDeleteWorldsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteWorlds, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) BatchDescribeSimulationJob(ctx workflow.Context, input *robomaker.BatchDescribeSimulationJobInput) (*robomaker.BatchDescribeSimulationJobOutput, error) {
    var output robomaker.BatchDescribeSimulationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDescribeSimulationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CancelDeploymentJob(ctx workflow.Context, input *robomaker.CancelDeploymentJobInput) (*robomaker.CancelDeploymentJobOutput, error) {
    var output robomaker.CancelDeploymentJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelDeploymentJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CancelSimulationJob(ctx workflow.Context, input *robomaker.CancelSimulationJobInput) (*robomaker.CancelSimulationJobOutput, error) {
    var output robomaker.CancelSimulationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelSimulationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CancelSimulationJobBatch(ctx workflow.Context, input *robomaker.CancelSimulationJobBatchInput) (*robomaker.CancelSimulationJobBatchOutput, error) {
    var output robomaker.CancelSimulationJobBatchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelSimulationJobBatch, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CancelWorldExportJob(ctx workflow.Context, input *robomaker.CancelWorldExportJobInput) (*robomaker.CancelWorldExportJobOutput, error) {
    var output robomaker.CancelWorldExportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelWorldExportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CancelWorldGenerationJob(ctx workflow.Context, input *robomaker.CancelWorldGenerationJobInput) (*robomaker.CancelWorldGenerationJobOutput, error) {
    var output robomaker.CancelWorldGenerationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelWorldGenerationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CreateDeploymentJob(ctx workflow.Context, input *robomaker.CreateDeploymentJobInput) (*robomaker.CreateDeploymentJobOutput, error) {
    var output robomaker.CreateDeploymentJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeploymentJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CreateFleet(ctx workflow.Context, input *robomaker.CreateFleetInput) (*robomaker.CreateFleetOutput, error) {
    var output robomaker.CreateFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CreateRobot(ctx workflow.Context, input *robomaker.CreateRobotInput) (*robomaker.CreateRobotOutput, error) {
    var output robomaker.CreateRobotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRobot, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CreateRobotApplication(ctx workflow.Context, input *robomaker.CreateRobotApplicationInput) (*robomaker.CreateRobotApplicationOutput, error) {
    var output robomaker.CreateRobotApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRobotApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CreateRobotApplicationVersion(ctx workflow.Context, input *robomaker.CreateRobotApplicationVersionInput) (*robomaker.CreateRobotApplicationVersionOutput, error) {
    var output robomaker.CreateRobotApplicationVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRobotApplicationVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CreateSimulationApplication(ctx workflow.Context, input *robomaker.CreateSimulationApplicationInput) (*robomaker.CreateSimulationApplicationOutput, error) {
    var output robomaker.CreateSimulationApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSimulationApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CreateSimulationApplicationVersion(ctx workflow.Context, input *robomaker.CreateSimulationApplicationVersionInput) (*robomaker.CreateSimulationApplicationVersionOutput, error) {
    var output robomaker.CreateSimulationApplicationVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSimulationApplicationVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CreateSimulationJob(ctx workflow.Context, input *robomaker.CreateSimulationJobInput) (*robomaker.CreateSimulationJobOutput, error) {
    var output robomaker.CreateSimulationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSimulationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CreateWorldExportJob(ctx workflow.Context, input *robomaker.CreateWorldExportJobInput) (*robomaker.CreateWorldExportJobOutput, error) {
    var output robomaker.CreateWorldExportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWorldExportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CreateWorldGenerationJob(ctx workflow.Context, input *robomaker.CreateWorldGenerationJobInput) (*robomaker.CreateWorldGenerationJobOutput, error) {
    var output robomaker.CreateWorldGenerationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWorldGenerationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) CreateWorldTemplate(ctx workflow.Context, input *robomaker.CreateWorldTemplateInput) (*robomaker.CreateWorldTemplateOutput, error) {
    var output robomaker.CreateWorldTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWorldTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DeleteFleet(ctx workflow.Context, input *robomaker.DeleteFleetInput) (*robomaker.DeleteFleetOutput, error) {
    var output robomaker.DeleteFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DeleteRobot(ctx workflow.Context, input *robomaker.DeleteRobotInput) (*robomaker.DeleteRobotOutput, error) {
    var output robomaker.DeleteRobotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRobot, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DeleteRobotApplication(ctx workflow.Context, input *robomaker.DeleteRobotApplicationInput) (*robomaker.DeleteRobotApplicationOutput, error) {
    var output robomaker.DeleteRobotApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRobotApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DeleteSimulationApplication(ctx workflow.Context, input *robomaker.DeleteSimulationApplicationInput) (*robomaker.DeleteSimulationApplicationOutput, error) {
    var output robomaker.DeleteSimulationApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSimulationApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DeleteWorldTemplate(ctx workflow.Context, input *robomaker.DeleteWorldTemplateInput) (*robomaker.DeleteWorldTemplateOutput, error) {
    var output robomaker.DeleteWorldTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteWorldTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DeregisterRobot(ctx workflow.Context, input *robomaker.DeregisterRobotInput) (*robomaker.DeregisterRobotOutput, error) {
    var output robomaker.DeregisterRobotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterRobot, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DescribeDeploymentJob(ctx workflow.Context, input *robomaker.DescribeDeploymentJobInput) (*robomaker.DescribeDeploymentJobOutput, error) {
    var output robomaker.DescribeDeploymentJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDeploymentJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DescribeFleet(ctx workflow.Context, input *robomaker.DescribeFleetInput) (*robomaker.DescribeFleetOutput, error) {
    var output robomaker.DescribeFleetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeFleet, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DescribeRobot(ctx workflow.Context, input *robomaker.DescribeRobotInput) (*robomaker.DescribeRobotOutput, error) {
    var output robomaker.DescribeRobotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRobot, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DescribeRobotApplication(ctx workflow.Context, input *robomaker.DescribeRobotApplicationInput) (*robomaker.DescribeRobotApplicationOutput, error) {
    var output robomaker.DescribeRobotApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRobotApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DescribeSimulationApplication(ctx workflow.Context, input *robomaker.DescribeSimulationApplicationInput) (*robomaker.DescribeSimulationApplicationOutput, error) {
    var output robomaker.DescribeSimulationApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSimulationApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DescribeSimulationJob(ctx workflow.Context, input *robomaker.DescribeSimulationJobInput) (*robomaker.DescribeSimulationJobOutput, error) {
    var output robomaker.DescribeSimulationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSimulationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DescribeSimulationJobBatch(ctx workflow.Context, input *robomaker.DescribeSimulationJobBatchInput) (*robomaker.DescribeSimulationJobBatchOutput, error) {
    var output robomaker.DescribeSimulationJobBatchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSimulationJobBatch, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DescribeWorld(ctx workflow.Context, input *robomaker.DescribeWorldInput) (*robomaker.DescribeWorldOutput, error) {
    var output robomaker.DescribeWorldOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorld, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DescribeWorldExportJob(ctx workflow.Context, input *robomaker.DescribeWorldExportJobInput) (*robomaker.DescribeWorldExportJobOutput, error) {
    var output robomaker.DescribeWorldExportJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorldExportJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DescribeWorldGenerationJob(ctx workflow.Context, input *robomaker.DescribeWorldGenerationJobInput) (*robomaker.DescribeWorldGenerationJobOutput, error) {
    var output robomaker.DescribeWorldGenerationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorldGenerationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) DescribeWorldTemplate(ctx workflow.Context, input *robomaker.DescribeWorldTemplateInput) (*robomaker.DescribeWorldTemplateOutput, error) {
    var output robomaker.DescribeWorldTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorldTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) GetWorldTemplateBody(ctx workflow.Context, input *robomaker.GetWorldTemplateBodyInput) (*robomaker.GetWorldTemplateBodyOutput, error) {
    var output robomaker.GetWorldTemplateBodyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWorldTemplateBody, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListDeploymentJobs(ctx workflow.Context, input *robomaker.ListDeploymentJobsInput) (*robomaker.ListDeploymentJobsOutput, error) {
    var output robomaker.ListDeploymentJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeploymentJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListFleets(ctx workflow.Context, input *robomaker.ListFleetsInput) (*robomaker.ListFleetsOutput, error) {
    var output robomaker.ListFleetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListFleets, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListRobotApplications(ctx workflow.Context, input *robomaker.ListRobotApplicationsInput) (*robomaker.ListRobotApplicationsOutput, error) {
    var output robomaker.ListRobotApplicationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRobotApplications, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListRobots(ctx workflow.Context, input *robomaker.ListRobotsInput) (*robomaker.ListRobotsOutput, error) {
    var output robomaker.ListRobotsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRobots, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListSimulationApplications(ctx workflow.Context, input *robomaker.ListSimulationApplicationsInput) (*robomaker.ListSimulationApplicationsOutput, error) {
    var output robomaker.ListSimulationApplicationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSimulationApplications, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListSimulationJobBatches(ctx workflow.Context, input *robomaker.ListSimulationJobBatchesInput) (*robomaker.ListSimulationJobBatchesOutput, error) {
    var output robomaker.ListSimulationJobBatchesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSimulationJobBatches, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListSimulationJobs(ctx workflow.Context, input *robomaker.ListSimulationJobsInput) (*robomaker.ListSimulationJobsOutput, error) {
    var output robomaker.ListSimulationJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSimulationJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListTagsForResource(ctx workflow.Context, input *robomaker.ListTagsForResourceInput) (*robomaker.ListTagsForResourceOutput, error) {
    var output robomaker.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListWorldExportJobs(ctx workflow.Context, input *robomaker.ListWorldExportJobsInput) (*robomaker.ListWorldExportJobsOutput, error) {
    var output robomaker.ListWorldExportJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWorldExportJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListWorldGenerationJobs(ctx workflow.Context, input *robomaker.ListWorldGenerationJobsInput) (*robomaker.ListWorldGenerationJobsOutput, error) {
    var output robomaker.ListWorldGenerationJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWorldGenerationJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListWorldTemplates(ctx workflow.Context, input *robomaker.ListWorldTemplatesInput) (*robomaker.ListWorldTemplatesOutput, error) {
    var output robomaker.ListWorldTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWorldTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) ListWorlds(ctx workflow.Context, input *robomaker.ListWorldsInput) (*robomaker.ListWorldsOutput, error) {
    var output robomaker.ListWorldsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWorlds, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) RegisterRobot(ctx workflow.Context, input *robomaker.RegisterRobotInput) (*robomaker.RegisterRobotOutput, error) {
    var output robomaker.RegisterRobotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterRobot, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) RestartSimulationJob(ctx workflow.Context, input *robomaker.RestartSimulationJobInput) (*robomaker.RestartSimulationJobOutput, error) {
    var output robomaker.RestartSimulationJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestartSimulationJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) StartSimulationJobBatch(ctx workflow.Context, input *robomaker.StartSimulationJobBatchInput) (*robomaker.StartSimulationJobBatchOutput, error) {
    var output robomaker.StartSimulationJobBatchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartSimulationJobBatch, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) SyncDeploymentJob(ctx workflow.Context, input *robomaker.SyncDeploymentJobInput) (*robomaker.SyncDeploymentJobOutput, error) {
    var output robomaker.SyncDeploymentJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SyncDeploymentJob, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) TagResource(ctx workflow.Context, input *robomaker.TagResourceInput) (*robomaker.TagResourceOutput, error) {
    var output robomaker.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) UntagResource(ctx workflow.Context, input *robomaker.UntagResourceInput) (*robomaker.UntagResourceOutput, error) {
    var output robomaker.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) UpdateRobotApplication(ctx workflow.Context, input *robomaker.UpdateRobotApplicationInput) (*robomaker.UpdateRobotApplicationOutput, error) {
    var output robomaker.UpdateRobotApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRobotApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) UpdateSimulationApplication(ctx workflow.Context, input *robomaker.UpdateSimulationApplicationInput) (*robomaker.UpdateSimulationApplicationOutput, error) {
    var output robomaker.UpdateSimulationApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSimulationApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *RoboMakerStub) UpdateWorldTemplate(ctx workflow.Context, input *robomaker.UpdateWorldTemplateInput) (*robomaker.UpdateWorldTemplateOutput, error) {
    var output robomaker.UpdateWorldTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateWorldTemplate, input).Get(ctx, &output)
    return &output, err
}
