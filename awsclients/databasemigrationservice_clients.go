package awsclients

import (
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type DatabaseMigrationServiceClient interface {
       AddTagsToResource(ctx workflow.Context, input *databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error)
       AddTagsToResourceAsync(ctx workflow.Context, input *databasemigrationservice.AddTagsToResourceInput) *DatabasemigrationserviceAddTagsToResourceResult

       ApplyPendingMaintenanceAction(ctx workflow.Context, input *databasemigrationservice.ApplyPendingMaintenanceActionInput) (*databasemigrationservice.ApplyPendingMaintenanceActionOutput, error)
       ApplyPendingMaintenanceActionAsync(ctx workflow.Context, input *databasemigrationservice.ApplyPendingMaintenanceActionInput) *DatabasemigrationserviceApplyPendingMaintenanceActionResult

       CancelReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) (*databasemigrationservice.CancelReplicationTaskAssessmentRunOutput, error)
       CancelReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) *DatabasemigrationserviceCancelReplicationTaskAssessmentRunResult

       CreateEndpoint(ctx workflow.Context, input *databasemigrationservice.CreateEndpointInput) (*databasemigrationservice.CreateEndpointOutput, error)
       CreateEndpointAsync(ctx workflow.Context, input *databasemigrationservice.CreateEndpointInput) *DatabasemigrationserviceCreateEndpointResult

       CreateEventSubscription(ctx workflow.Context, input *databasemigrationservice.CreateEventSubscriptionInput) (*databasemigrationservice.CreateEventSubscriptionOutput, error)
       CreateEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.CreateEventSubscriptionInput) *DatabasemigrationserviceCreateEventSubscriptionResult

       CreateReplicationInstance(ctx workflow.Context, input *databasemigrationservice.CreateReplicationInstanceInput) (*databasemigrationservice.CreateReplicationInstanceOutput, error)
       CreateReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationInstanceInput) *DatabasemigrationserviceCreateReplicationInstanceResult

       CreateReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.CreateReplicationSubnetGroupInput) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error)
       CreateReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationSubnetGroupInput) *DatabasemigrationserviceCreateReplicationSubnetGroupResult

       CreateReplicationTask(ctx workflow.Context, input *databasemigrationservice.CreateReplicationTaskInput) (*databasemigrationservice.CreateReplicationTaskOutput, error)
       CreateReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationTaskInput) *DatabasemigrationserviceCreateReplicationTaskResult

       DeleteCertificate(ctx workflow.Context, input *databasemigrationservice.DeleteCertificateInput) (*databasemigrationservice.DeleteCertificateOutput, error)
       DeleteCertificateAsync(ctx workflow.Context, input *databasemigrationservice.DeleteCertificateInput) *DatabasemigrationserviceDeleteCertificateResult

       DeleteConnection(ctx workflow.Context, input *databasemigrationservice.DeleteConnectionInput) (*databasemigrationservice.DeleteConnectionOutput, error)
       DeleteConnectionAsync(ctx workflow.Context, input *databasemigrationservice.DeleteConnectionInput) *DatabasemigrationserviceDeleteConnectionResult

       DeleteEndpoint(ctx workflow.Context, input *databasemigrationservice.DeleteEndpointInput) (*databasemigrationservice.DeleteEndpointOutput, error)
       DeleteEndpointAsync(ctx workflow.Context, input *databasemigrationservice.DeleteEndpointInput) *DatabasemigrationserviceDeleteEndpointResult

       DeleteEventSubscription(ctx workflow.Context, input *databasemigrationservice.DeleteEventSubscriptionInput) (*databasemigrationservice.DeleteEventSubscriptionOutput, error)
       DeleteEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.DeleteEventSubscriptionInput) *DatabasemigrationserviceDeleteEventSubscriptionResult

       DeleteReplicationInstance(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationInstanceInput) (*databasemigrationservice.DeleteReplicationInstanceOutput, error)
       DeleteReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationInstanceInput) *DatabasemigrationserviceDeleteReplicationInstanceResult

       DeleteReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationSubnetGroupInput) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error)
       DeleteReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationSubnetGroupInput) *DatabasemigrationserviceDeleteReplicationSubnetGroupResult

       DeleteReplicationTask(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskInput) (*databasemigrationservice.DeleteReplicationTaskOutput, error)
       DeleteReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskInput) *DatabasemigrationserviceDeleteReplicationTaskResult

       DeleteReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) (*databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput, error)
       DeleteReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) *DatabasemigrationserviceDeleteReplicationTaskAssessmentRunResult

       DescribeAccountAttributes(ctx workflow.Context, input *databasemigrationservice.DescribeAccountAttributesInput) (*databasemigrationservice.DescribeAccountAttributesOutput, error)
       DescribeAccountAttributesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeAccountAttributesInput) *DatabasemigrationserviceDescribeAccountAttributesResult

       DescribeApplicableIndividualAssessments(ctx workflow.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error)
       DescribeApplicableIndividualAssessmentsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) *DatabasemigrationserviceDescribeApplicableIndividualAssessmentsResult

       DescribeCertificates(ctx workflow.Context, input *databasemigrationservice.DescribeCertificatesInput) (*databasemigrationservice.DescribeCertificatesOutput, error)
       DescribeCertificatesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeCertificatesInput) *DatabasemigrationserviceDescribeCertificatesResult

       DescribeConnections(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) (*databasemigrationservice.DescribeConnectionsOutput, error)
       DescribeConnectionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) *DatabasemigrationserviceDescribeConnectionsResult

       DescribeEndpointTypes(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointTypesInput) (*databasemigrationservice.DescribeEndpointTypesOutput, error)
       DescribeEndpointTypesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointTypesInput) *DatabasemigrationserviceDescribeEndpointTypesResult

       DescribeEndpoints(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) (*databasemigrationservice.DescribeEndpointsOutput, error)
       DescribeEndpointsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) *DatabasemigrationserviceDescribeEndpointsResult

       DescribeEventCategories(ctx workflow.Context, input *databasemigrationservice.DescribeEventCategoriesInput) (*databasemigrationservice.DescribeEventCategoriesOutput, error)
       DescribeEventCategoriesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventCategoriesInput) *DatabasemigrationserviceDescribeEventCategoriesResult

       DescribeEventSubscriptions(ctx workflow.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error)
       DescribeEventSubscriptionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput) *DatabasemigrationserviceDescribeEventSubscriptionsResult

       DescribeEvents(ctx workflow.Context, input *databasemigrationservice.DescribeEventsInput) (*databasemigrationservice.DescribeEventsOutput, error)
       DescribeEventsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventsInput) *DatabasemigrationserviceDescribeEventsResult

       DescribeOrderableReplicationInstances(ctx workflow.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error)
       DescribeOrderableReplicationInstancesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) *DatabasemigrationserviceDescribeOrderableReplicationInstancesResult

       DescribePendingMaintenanceActions(ctx workflow.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error)
       DescribePendingMaintenanceActionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput) *DatabasemigrationserviceDescribePendingMaintenanceActionsResult

       DescribeRefreshSchemasStatus(ctx workflow.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error)
       DescribeRefreshSchemasStatusAsync(ctx workflow.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput) *DatabasemigrationserviceDescribeRefreshSchemasStatusResult

       DescribeReplicationInstanceTaskLogs(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error)
       DescribeReplicationInstanceTaskLogsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) *DatabasemigrationserviceDescribeReplicationInstanceTaskLogsResult

       DescribeReplicationInstances(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error)
       DescribeReplicationInstancesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) *DatabasemigrationserviceDescribeReplicationInstancesResult

       DescribeReplicationSubnetGroups(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error)
       DescribeReplicationSubnetGroupsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) *DatabasemigrationserviceDescribeReplicationSubnetGroupsResult

       DescribeReplicationTaskAssessmentResults(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error)
       DescribeReplicationTaskAssessmentResultsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) *DatabasemigrationserviceDescribeReplicationTaskAssessmentResultsResult

       DescribeReplicationTaskAssessmentRuns(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error)
       DescribeReplicationTaskAssessmentRunsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) *DatabasemigrationserviceDescribeReplicationTaskAssessmentRunsResult

       DescribeReplicationTaskIndividualAssessments(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error)
       DescribeReplicationTaskIndividualAssessmentsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) *DatabasemigrationserviceDescribeReplicationTaskIndividualAssessmentsResult

       DescribeReplicationTasks(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) (*databasemigrationservice.DescribeReplicationTasksOutput, error)
       DescribeReplicationTasksAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *DatabasemigrationserviceDescribeReplicationTasksResult

       DescribeSchemas(ctx workflow.Context, input *databasemigrationservice.DescribeSchemasInput) (*databasemigrationservice.DescribeSchemasOutput, error)
       DescribeSchemasAsync(ctx workflow.Context, input *databasemigrationservice.DescribeSchemasInput) *DatabasemigrationserviceDescribeSchemasResult

       DescribeTableStatistics(ctx workflow.Context, input *databasemigrationservice.DescribeTableStatisticsInput) (*databasemigrationservice.DescribeTableStatisticsOutput, error)
       DescribeTableStatisticsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeTableStatisticsInput) *DatabasemigrationserviceDescribeTableStatisticsResult

       ImportCertificate(ctx workflow.Context, input *databasemigrationservice.ImportCertificateInput) (*databasemigrationservice.ImportCertificateOutput, error)
       ImportCertificateAsync(ctx workflow.Context, input *databasemigrationservice.ImportCertificateInput) *DatabasemigrationserviceImportCertificateResult

       ListTagsForResource(ctx workflow.Context, input *databasemigrationservice.ListTagsForResourceInput) (*databasemigrationservice.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *databasemigrationservice.ListTagsForResourceInput) *DatabasemigrationserviceListTagsForResourceResult

       ModifyEndpoint(ctx workflow.Context, input *databasemigrationservice.ModifyEndpointInput) (*databasemigrationservice.ModifyEndpointOutput, error)
       ModifyEndpointAsync(ctx workflow.Context, input *databasemigrationservice.ModifyEndpointInput) *DatabasemigrationserviceModifyEndpointResult

       ModifyEventSubscription(ctx workflow.Context, input *databasemigrationservice.ModifyEventSubscriptionInput) (*databasemigrationservice.ModifyEventSubscriptionOutput, error)
       ModifyEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.ModifyEventSubscriptionInput) *DatabasemigrationserviceModifyEventSubscriptionResult

       ModifyReplicationInstance(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationInstanceInput) (*databasemigrationservice.ModifyReplicationInstanceOutput, error)
       ModifyReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationInstanceInput) *DatabasemigrationserviceModifyReplicationInstanceResult

       ModifyReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationSubnetGroupInput) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error)
       ModifyReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationSubnetGroupInput) *DatabasemigrationserviceModifyReplicationSubnetGroupResult

       ModifyReplicationTask(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationTaskInput) (*databasemigrationservice.ModifyReplicationTaskOutput, error)
       ModifyReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationTaskInput) *DatabasemigrationserviceModifyReplicationTaskResult

       RebootReplicationInstance(ctx workflow.Context, input *databasemigrationservice.RebootReplicationInstanceInput) (*databasemigrationservice.RebootReplicationInstanceOutput, error)
       RebootReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.RebootReplicationInstanceInput) *DatabasemigrationserviceRebootReplicationInstanceResult

       RefreshSchemas(ctx workflow.Context, input *databasemigrationservice.RefreshSchemasInput) (*databasemigrationservice.RefreshSchemasOutput, error)
       RefreshSchemasAsync(ctx workflow.Context, input *databasemigrationservice.RefreshSchemasInput) *DatabasemigrationserviceRefreshSchemasResult

       ReloadTables(ctx workflow.Context, input *databasemigrationservice.ReloadTablesInput) (*databasemigrationservice.ReloadTablesOutput, error)
       ReloadTablesAsync(ctx workflow.Context, input *databasemigrationservice.ReloadTablesInput) *DatabasemigrationserviceReloadTablesResult

       RemoveTagsFromResource(ctx workflow.Context, input *databasemigrationservice.RemoveTagsFromResourceInput) (*databasemigrationservice.RemoveTagsFromResourceOutput, error)
       RemoveTagsFromResourceAsync(ctx workflow.Context, input *databasemigrationservice.RemoveTagsFromResourceInput) *DatabasemigrationserviceRemoveTagsFromResourceResult

       StartReplicationTask(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskInput) (*databasemigrationservice.StartReplicationTaskOutput, error)
       StartReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskInput) *DatabasemigrationserviceStartReplicationTaskResult

       StartReplicationTaskAssessment(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentInput) (*databasemigrationservice.StartReplicationTaskAssessmentOutput, error)
       StartReplicationTaskAssessmentAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentInput) *DatabasemigrationserviceStartReplicationTaskAssessmentResult

       StartReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) (*databasemigrationservice.StartReplicationTaskAssessmentRunOutput, error)
       StartReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) *DatabasemigrationserviceStartReplicationTaskAssessmentRunResult

       StopReplicationTask(ctx workflow.Context, input *databasemigrationservice.StopReplicationTaskInput) (*databasemigrationservice.StopReplicationTaskOutput, error)
       StopReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.StopReplicationTaskInput) *DatabasemigrationserviceStopReplicationTaskResult

       TestConnection(ctx workflow.Context, input *databasemigrationservice.TestConnectionInput) (*databasemigrationservice.TestConnectionOutput, error)
       TestConnectionAsync(ctx workflow.Context, input *databasemigrationservice.TestConnectionInput) *DatabasemigrationserviceTestConnectionResult

       WaitUntilEndpointDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) error
       WaitUntilReplicationInstanceAvailable(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error
       WaitUntilReplicationInstanceDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error
       WaitUntilReplicationTaskDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
       WaitUntilReplicationTaskReady(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
       WaitUntilReplicationTaskRunning(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
       WaitUntilReplicationTaskStopped(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error
       WaitUntilTestConnectionSucceeds(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) error}

type DatabasemigrationserviceAddTagsToResourceResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceAddTagsToResourceResult) Get(ctx workflow.Context) (*databasemigrationservice.AddTagsToResourceOutput, error) {
    var output databasemigrationservice.AddTagsToResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceApplyPendingMaintenanceActionResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceApplyPendingMaintenanceActionResult) Get(ctx workflow.Context) (*databasemigrationservice.ApplyPendingMaintenanceActionOutput, error) {
    var output databasemigrationservice.ApplyPendingMaintenanceActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceCancelReplicationTaskAssessmentRunResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceCancelReplicationTaskAssessmentRunResult) Get(ctx workflow.Context) (*databasemigrationservice.CancelReplicationTaskAssessmentRunOutput, error) {
    var output databasemigrationservice.CancelReplicationTaskAssessmentRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceCreateEndpointResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceCreateEndpointResult) Get(ctx workflow.Context) (*databasemigrationservice.CreateEndpointOutput, error) {
    var output databasemigrationservice.CreateEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceCreateEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceCreateEventSubscriptionResult) Get(ctx workflow.Context) (*databasemigrationservice.CreateEventSubscriptionOutput, error) {
    var output databasemigrationservice.CreateEventSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceCreateReplicationInstanceResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceCreateReplicationInstanceResult) Get(ctx workflow.Context) (*databasemigrationservice.CreateReplicationInstanceOutput, error) {
    var output databasemigrationservice.CreateReplicationInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceCreateReplicationSubnetGroupResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceCreateReplicationSubnetGroupResult) Get(ctx workflow.Context) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error) {
    var output databasemigrationservice.CreateReplicationSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceCreateReplicationTaskResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceCreateReplicationTaskResult) Get(ctx workflow.Context) (*databasemigrationservice.CreateReplicationTaskOutput, error) {
    var output databasemigrationservice.CreateReplicationTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDeleteCertificateResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDeleteCertificateResult) Get(ctx workflow.Context) (*databasemigrationservice.DeleteCertificateOutput, error) {
    var output databasemigrationservice.DeleteCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDeleteConnectionResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDeleteConnectionResult) Get(ctx workflow.Context) (*databasemigrationservice.DeleteConnectionOutput, error) {
    var output databasemigrationservice.DeleteConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDeleteEndpointResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDeleteEndpointResult) Get(ctx workflow.Context) (*databasemigrationservice.DeleteEndpointOutput, error) {
    var output databasemigrationservice.DeleteEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDeleteEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDeleteEventSubscriptionResult) Get(ctx workflow.Context) (*databasemigrationservice.DeleteEventSubscriptionOutput, error) {
    var output databasemigrationservice.DeleteEventSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDeleteReplicationInstanceResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDeleteReplicationInstanceResult) Get(ctx workflow.Context) (*databasemigrationservice.DeleteReplicationInstanceOutput, error) {
    var output databasemigrationservice.DeleteReplicationInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDeleteReplicationSubnetGroupResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDeleteReplicationSubnetGroupResult) Get(ctx workflow.Context) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error) {
    var output databasemigrationservice.DeleteReplicationSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDeleteReplicationTaskResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDeleteReplicationTaskResult) Get(ctx workflow.Context) (*databasemigrationservice.DeleteReplicationTaskOutput, error) {
    var output databasemigrationservice.DeleteReplicationTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDeleteReplicationTaskAssessmentRunResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDeleteReplicationTaskAssessmentRunResult) Get(ctx workflow.Context) (*databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput, error) {
    var output databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeAccountAttributesResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeAccountAttributesResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeAccountAttributesOutput, error) {
    var output databasemigrationservice.DescribeAccountAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeApplicableIndividualAssessmentsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeApplicableIndividualAssessmentsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error) {
    var output databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeCertificatesResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeCertificatesResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeCertificatesOutput, error) {
    var output databasemigrationservice.DescribeCertificatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeConnectionsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeConnectionsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeConnectionsOutput, error) {
    var output databasemigrationservice.DescribeConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeEndpointTypesResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeEndpointTypesResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeEndpointTypesOutput, error) {
    var output databasemigrationservice.DescribeEndpointTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeEndpointsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeEndpointsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeEndpointsOutput, error) {
    var output databasemigrationservice.DescribeEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeEventCategoriesResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeEventCategoriesResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeEventCategoriesOutput, error) {
    var output databasemigrationservice.DescribeEventCategoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeEventSubscriptionsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeEventSubscriptionsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error) {
    var output databasemigrationservice.DescribeEventSubscriptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeEventsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeEventsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeEventsOutput, error) {
    var output databasemigrationservice.DescribeEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeOrderableReplicationInstancesResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeOrderableReplicationInstancesResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error) {
    var output databasemigrationservice.DescribeOrderableReplicationInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribePendingMaintenanceActionsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribePendingMaintenanceActionsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error) {
    var output databasemigrationservice.DescribePendingMaintenanceActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeRefreshSchemasStatusResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeRefreshSchemasStatusResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error) {
    var output databasemigrationservice.DescribeRefreshSchemasStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeReplicationInstanceTaskLogsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeReplicationInstanceTaskLogsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error) {
    var output databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeReplicationInstancesResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeReplicationInstancesResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
    var output databasemigrationservice.DescribeReplicationInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeReplicationSubnetGroupsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeReplicationSubnetGroupsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error) {
    var output databasemigrationservice.DescribeReplicationSubnetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeReplicationTaskAssessmentResultsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeReplicationTaskAssessmentResultsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error) {
    var output databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeReplicationTaskAssessmentRunsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeReplicationTaskAssessmentRunsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error) {
    var output databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeReplicationTaskIndividualAssessmentsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeReplicationTaskIndividualAssessmentsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error) {
    var output databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeReplicationTasksResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeReplicationTasksResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeReplicationTasksOutput, error) {
    var output databasemigrationservice.DescribeReplicationTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeSchemasResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeSchemasResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeSchemasOutput, error) {
    var output databasemigrationservice.DescribeSchemasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceDescribeTableStatisticsResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceDescribeTableStatisticsResult) Get(ctx workflow.Context) (*databasemigrationservice.DescribeTableStatisticsOutput, error) {
    var output databasemigrationservice.DescribeTableStatisticsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceImportCertificateResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceImportCertificateResult) Get(ctx workflow.Context) (*databasemigrationservice.ImportCertificateOutput, error) {
    var output databasemigrationservice.ImportCertificateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceListTagsForResourceResult) Get(ctx workflow.Context) (*databasemigrationservice.ListTagsForResourceOutput, error) {
    var output databasemigrationservice.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceModifyEndpointResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceModifyEndpointResult) Get(ctx workflow.Context) (*databasemigrationservice.ModifyEndpointOutput, error) {
    var output databasemigrationservice.ModifyEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceModifyEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceModifyEventSubscriptionResult) Get(ctx workflow.Context) (*databasemigrationservice.ModifyEventSubscriptionOutput, error) {
    var output databasemigrationservice.ModifyEventSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceModifyReplicationInstanceResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceModifyReplicationInstanceResult) Get(ctx workflow.Context) (*databasemigrationservice.ModifyReplicationInstanceOutput, error) {
    var output databasemigrationservice.ModifyReplicationInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceModifyReplicationSubnetGroupResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceModifyReplicationSubnetGroupResult) Get(ctx workflow.Context) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error) {
    var output databasemigrationservice.ModifyReplicationSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceModifyReplicationTaskResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceModifyReplicationTaskResult) Get(ctx workflow.Context) (*databasemigrationservice.ModifyReplicationTaskOutput, error) {
    var output databasemigrationservice.ModifyReplicationTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceRebootReplicationInstanceResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceRebootReplicationInstanceResult) Get(ctx workflow.Context) (*databasemigrationservice.RebootReplicationInstanceOutput, error) {
    var output databasemigrationservice.RebootReplicationInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceRefreshSchemasResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceRefreshSchemasResult) Get(ctx workflow.Context) (*databasemigrationservice.RefreshSchemasOutput, error) {
    var output databasemigrationservice.RefreshSchemasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceReloadTablesResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceReloadTablesResult) Get(ctx workflow.Context) (*databasemigrationservice.ReloadTablesOutput, error) {
    var output databasemigrationservice.ReloadTablesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceRemoveTagsFromResourceResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceRemoveTagsFromResourceResult) Get(ctx workflow.Context) (*databasemigrationservice.RemoveTagsFromResourceOutput, error) {
    var output databasemigrationservice.RemoveTagsFromResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceStartReplicationTaskResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceStartReplicationTaskResult) Get(ctx workflow.Context) (*databasemigrationservice.StartReplicationTaskOutput, error) {
    var output databasemigrationservice.StartReplicationTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceStartReplicationTaskAssessmentResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceStartReplicationTaskAssessmentResult) Get(ctx workflow.Context) (*databasemigrationservice.StartReplicationTaskAssessmentOutput, error) {
    var output databasemigrationservice.StartReplicationTaskAssessmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceStartReplicationTaskAssessmentRunResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceStartReplicationTaskAssessmentRunResult) Get(ctx workflow.Context) (*databasemigrationservice.StartReplicationTaskAssessmentRunOutput, error) {
    var output databasemigrationservice.StartReplicationTaskAssessmentRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceStopReplicationTaskResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceStopReplicationTaskResult) Get(ctx workflow.Context) (*databasemigrationservice.StopReplicationTaskOutput, error) {
    var output databasemigrationservice.StopReplicationTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabasemigrationserviceTestConnectionResult struct {
	Result workflow.Future
}

func (r *DatabasemigrationserviceTestConnectionResult) Get(ctx workflow.Context) (*databasemigrationservice.TestConnectionOutput, error) {
    var output databasemigrationservice.TestConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DatabaseMigrationServiceStub struct {
    activities awsactivities.DatabaseMigrationServiceActivities
}

func NewDatabaseMigrationServiceStub() DatabaseMigrationServiceClient {
    return &DatabaseMigrationServiceStub{}
}

func (a *DatabaseMigrationServiceStub) AddTagsToResource(ctx workflow.Context, input *databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error) {
    var output databasemigrationservice.AddTagsToResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTagsToResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) AddTagsToResourceAsync(ctx workflow.Context, input *databasemigrationservice.AddTagsToResourceInput) *DatabasemigrationserviceAddTagsToResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddTagsToResource, input)
    return &DatabasemigrationserviceAddTagsToResourceResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) ApplyPendingMaintenanceAction(ctx workflow.Context, input *databasemigrationservice.ApplyPendingMaintenanceActionInput) (*databasemigrationservice.ApplyPendingMaintenanceActionOutput, error) {
    var output databasemigrationservice.ApplyPendingMaintenanceActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ApplyPendingMaintenanceAction, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) ApplyPendingMaintenanceActionAsync(ctx workflow.Context, input *databasemigrationservice.ApplyPendingMaintenanceActionInput) *DatabasemigrationserviceApplyPendingMaintenanceActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ApplyPendingMaintenanceAction, input)
    return &DatabasemigrationserviceApplyPendingMaintenanceActionResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) CancelReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) (*databasemigrationservice.CancelReplicationTaskAssessmentRunOutput, error) {
    var output databasemigrationservice.CancelReplicationTaskAssessmentRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelReplicationTaskAssessmentRun, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) CancelReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) *DatabasemigrationserviceCancelReplicationTaskAssessmentRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelReplicationTaskAssessmentRun, input)
    return &DatabasemigrationserviceCancelReplicationTaskAssessmentRunResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) CreateEndpoint(ctx workflow.Context, input *databasemigrationservice.CreateEndpointInput) (*databasemigrationservice.CreateEndpointOutput, error) {
    var output databasemigrationservice.CreateEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) CreateEndpointAsync(ctx workflow.Context, input *databasemigrationservice.CreateEndpointInput) *DatabasemigrationserviceCreateEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEndpoint, input)
    return &DatabasemigrationserviceCreateEndpointResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) CreateEventSubscription(ctx workflow.Context, input *databasemigrationservice.CreateEventSubscriptionInput) (*databasemigrationservice.CreateEventSubscriptionOutput, error) {
    var output databasemigrationservice.CreateEventSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEventSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) CreateEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.CreateEventSubscriptionInput) *DatabasemigrationserviceCreateEventSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEventSubscription, input)
    return &DatabasemigrationserviceCreateEventSubscriptionResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) CreateReplicationInstance(ctx workflow.Context, input *databasemigrationservice.CreateReplicationInstanceInput) (*databasemigrationservice.CreateReplicationInstanceOutput, error) {
    var output databasemigrationservice.CreateReplicationInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateReplicationInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) CreateReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationInstanceInput) *DatabasemigrationserviceCreateReplicationInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateReplicationInstance, input)
    return &DatabasemigrationserviceCreateReplicationInstanceResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) CreateReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.CreateReplicationSubnetGroupInput) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error) {
    var output databasemigrationservice.CreateReplicationSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateReplicationSubnetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) CreateReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationSubnetGroupInput) *DatabasemigrationserviceCreateReplicationSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateReplicationSubnetGroup, input)
    return &DatabasemigrationserviceCreateReplicationSubnetGroupResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) CreateReplicationTask(ctx workflow.Context, input *databasemigrationservice.CreateReplicationTaskInput) (*databasemigrationservice.CreateReplicationTaskOutput, error) {
    var output databasemigrationservice.CreateReplicationTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateReplicationTask, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) CreateReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.CreateReplicationTaskInput) *DatabasemigrationserviceCreateReplicationTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateReplicationTask, input)
    return &DatabasemigrationserviceCreateReplicationTaskResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DeleteCertificate(ctx workflow.Context, input *databasemigrationservice.DeleteCertificateInput) (*databasemigrationservice.DeleteCertificateOutput, error) {
    var output databasemigrationservice.DeleteCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DeleteCertificateAsync(ctx workflow.Context, input *databasemigrationservice.DeleteCertificateInput) *DatabasemigrationserviceDeleteCertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCertificate, input)
    return &DatabasemigrationserviceDeleteCertificateResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DeleteConnection(ctx workflow.Context, input *databasemigrationservice.DeleteConnectionInput) (*databasemigrationservice.DeleteConnectionOutput, error) {
    var output databasemigrationservice.DeleteConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DeleteConnectionAsync(ctx workflow.Context, input *databasemigrationservice.DeleteConnectionInput) *DatabasemigrationserviceDeleteConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConnection, input)
    return &DatabasemigrationserviceDeleteConnectionResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DeleteEndpoint(ctx workflow.Context, input *databasemigrationservice.DeleteEndpointInput) (*databasemigrationservice.DeleteEndpointOutput, error) {
    var output databasemigrationservice.DeleteEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DeleteEndpointAsync(ctx workflow.Context, input *databasemigrationservice.DeleteEndpointInput) *DatabasemigrationserviceDeleteEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEndpoint, input)
    return &DatabasemigrationserviceDeleteEndpointResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DeleteEventSubscription(ctx workflow.Context, input *databasemigrationservice.DeleteEventSubscriptionInput) (*databasemigrationservice.DeleteEventSubscriptionOutput, error) {
    var output databasemigrationservice.DeleteEventSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEventSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DeleteEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.DeleteEventSubscriptionInput) *DatabasemigrationserviceDeleteEventSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEventSubscription, input)
    return &DatabasemigrationserviceDeleteEventSubscriptionResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DeleteReplicationInstance(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationInstanceInput) (*databasemigrationservice.DeleteReplicationInstanceOutput, error) {
    var output databasemigrationservice.DeleteReplicationInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteReplicationInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DeleteReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationInstanceInput) *DatabasemigrationserviceDeleteReplicationInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteReplicationInstance, input)
    return &DatabasemigrationserviceDeleteReplicationInstanceResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DeleteReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationSubnetGroupInput) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error) {
    var output databasemigrationservice.DeleteReplicationSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteReplicationSubnetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DeleteReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationSubnetGroupInput) *DatabasemigrationserviceDeleteReplicationSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteReplicationSubnetGroup, input)
    return &DatabasemigrationserviceDeleteReplicationSubnetGroupResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DeleteReplicationTask(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskInput) (*databasemigrationservice.DeleteReplicationTaskOutput, error) {
    var output databasemigrationservice.DeleteReplicationTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteReplicationTask, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DeleteReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskInput) *DatabasemigrationserviceDeleteReplicationTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteReplicationTask, input)
    return &DatabasemigrationserviceDeleteReplicationTaskResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DeleteReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) (*databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput, error) {
    var output databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteReplicationTaskAssessmentRun, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DeleteReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) *DatabasemigrationserviceDeleteReplicationTaskAssessmentRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteReplicationTaskAssessmentRun, input)
    return &DatabasemigrationserviceDeleteReplicationTaskAssessmentRunResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeAccountAttributes(ctx workflow.Context, input *databasemigrationservice.DescribeAccountAttributesInput) (*databasemigrationservice.DescribeAccountAttributesOutput, error) {
    var output databasemigrationservice.DescribeAccountAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeAccountAttributesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeAccountAttributesInput) *DatabasemigrationserviceDescribeAccountAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountAttributes, input)
    return &DatabasemigrationserviceDescribeAccountAttributesResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeApplicableIndividualAssessments(ctx workflow.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error) {
    var output databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeApplicableIndividualAssessments, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeApplicableIndividualAssessmentsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) *DatabasemigrationserviceDescribeApplicableIndividualAssessmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeApplicableIndividualAssessments, input)
    return &DatabasemigrationserviceDescribeApplicableIndividualAssessmentsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeCertificates(ctx workflow.Context, input *databasemigrationservice.DescribeCertificatesInput) (*databasemigrationservice.DescribeCertificatesOutput, error) {
    var output databasemigrationservice.DescribeCertificatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCertificates, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeCertificatesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeCertificatesInput) *DatabasemigrationserviceDescribeCertificatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCertificates, input)
    return &DatabasemigrationserviceDescribeCertificatesResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeConnections(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) (*databasemigrationservice.DescribeConnectionsOutput, error) {
    var output databasemigrationservice.DescribeConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeConnectionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) *DatabasemigrationserviceDescribeConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConnections, input)
    return &DatabasemigrationserviceDescribeConnectionsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeEndpointTypes(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointTypesInput) (*databasemigrationservice.DescribeEndpointTypesOutput, error) {
    var output databasemigrationservice.DescribeEndpointTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpointTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeEndpointTypesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointTypesInput) *DatabasemigrationserviceDescribeEndpointTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpointTypes, input)
    return &DatabasemigrationserviceDescribeEndpointTypesResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeEndpoints(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) (*databasemigrationservice.DescribeEndpointsOutput, error) {
    var output databasemigrationservice.DescribeEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeEndpointsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) *DatabasemigrationserviceDescribeEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEndpoints, input)
    return &DatabasemigrationserviceDescribeEndpointsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeEventCategories(ctx workflow.Context, input *databasemigrationservice.DescribeEventCategoriesInput) (*databasemigrationservice.DescribeEventCategoriesOutput, error) {
    var output databasemigrationservice.DescribeEventCategoriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventCategories, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeEventCategoriesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventCategoriesInput) *DatabasemigrationserviceDescribeEventCategoriesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventCategories, input)
    return &DatabasemigrationserviceDescribeEventCategoriesResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeEventSubscriptions(ctx workflow.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error) {
    var output databasemigrationservice.DescribeEventSubscriptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventSubscriptions, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeEventSubscriptionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput) *DatabasemigrationserviceDescribeEventSubscriptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEventSubscriptions, input)
    return &DatabasemigrationserviceDescribeEventSubscriptionsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeEvents(ctx workflow.Context, input *databasemigrationservice.DescribeEventsInput) (*databasemigrationservice.DescribeEventsOutput, error) {
    var output databasemigrationservice.DescribeEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeEventsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEventsInput) *DatabasemigrationserviceDescribeEventsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input)
    return &DatabasemigrationserviceDescribeEventsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeOrderableReplicationInstances(ctx workflow.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error) {
    var output databasemigrationservice.DescribeOrderableReplicationInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrderableReplicationInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeOrderableReplicationInstancesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) *DatabasemigrationserviceDescribeOrderableReplicationInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOrderableReplicationInstances, input)
    return &DatabasemigrationserviceDescribeOrderableReplicationInstancesResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribePendingMaintenanceActions(ctx workflow.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error) {
    var output databasemigrationservice.DescribePendingMaintenanceActionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePendingMaintenanceActions, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribePendingMaintenanceActionsAsync(ctx workflow.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput) *DatabasemigrationserviceDescribePendingMaintenanceActionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePendingMaintenanceActions, input)
    return &DatabasemigrationserviceDescribePendingMaintenanceActionsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeRefreshSchemasStatus(ctx workflow.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error) {
    var output databasemigrationservice.DescribeRefreshSchemasStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRefreshSchemasStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeRefreshSchemasStatusAsync(ctx workflow.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput) *DatabasemigrationserviceDescribeRefreshSchemasStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRefreshSchemasStatus, input)
    return &DatabasemigrationserviceDescribeRefreshSchemasStatusResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationInstanceTaskLogs(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error) {
    var output databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationInstanceTaskLogs, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationInstanceTaskLogsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) *DatabasemigrationserviceDescribeReplicationInstanceTaskLogsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationInstanceTaskLogs, input)
    return &DatabasemigrationserviceDescribeReplicationInstanceTaskLogsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationInstances(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
    var output databasemigrationservice.DescribeReplicationInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationInstancesAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) *DatabasemigrationserviceDescribeReplicationInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationInstances, input)
    return &DatabasemigrationserviceDescribeReplicationInstancesResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationSubnetGroups(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error) {
    var output databasemigrationservice.DescribeReplicationSubnetGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationSubnetGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationSubnetGroupsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) *DatabasemigrationserviceDescribeReplicationSubnetGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationSubnetGroups, input)
    return &DatabasemigrationserviceDescribeReplicationSubnetGroupsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationTaskAssessmentResults(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error) {
    var output databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationTaskAssessmentResults, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationTaskAssessmentResultsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) *DatabasemigrationserviceDescribeReplicationTaskAssessmentResultsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationTaskAssessmentResults, input)
    return &DatabasemigrationserviceDescribeReplicationTaskAssessmentResultsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationTaskAssessmentRuns(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error) {
    var output databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationTaskAssessmentRuns, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationTaskAssessmentRunsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) *DatabasemigrationserviceDescribeReplicationTaskAssessmentRunsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationTaskAssessmentRuns, input)
    return &DatabasemigrationserviceDescribeReplicationTaskAssessmentRunsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationTaskIndividualAssessments(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error) {
    var output databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationTaskIndividualAssessments, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationTaskIndividualAssessmentsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) *DatabasemigrationserviceDescribeReplicationTaskIndividualAssessmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationTaskIndividualAssessments, input)
    return &DatabasemigrationserviceDescribeReplicationTaskIndividualAssessmentsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationTasks(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) (*databasemigrationservice.DescribeReplicationTasksOutput, error) {
    var output databasemigrationservice.DescribeReplicationTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeReplicationTasksAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) *DatabasemigrationserviceDescribeReplicationTasksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReplicationTasks, input)
    return &DatabasemigrationserviceDescribeReplicationTasksResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeSchemas(ctx workflow.Context, input *databasemigrationservice.DescribeSchemasInput) (*databasemigrationservice.DescribeSchemasOutput, error) {
    var output databasemigrationservice.DescribeSchemasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSchemas, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeSchemasAsync(ctx workflow.Context, input *databasemigrationservice.DescribeSchemasInput) *DatabasemigrationserviceDescribeSchemasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSchemas, input)
    return &DatabasemigrationserviceDescribeSchemasResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) DescribeTableStatistics(ctx workflow.Context, input *databasemigrationservice.DescribeTableStatisticsInput) (*databasemigrationservice.DescribeTableStatisticsOutput, error) {
    var output databasemigrationservice.DescribeTableStatisticsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTableStatistics, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) DescribeTableStatisticsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeTableStatisticsInput) *DatabasemigrationserviceDescribeTableStatisticsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTableStatistics, input)
    return &DatabasemigrationserviceDescribeTableStatisticsResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) ImportCertificate(ctx workflow.Context, input *databasemigrationservice.ImportCertificateInput) (*databasemigrationservice.ImportCertificateOutput, error) {
    var output databasemigrationservice.ImportCertificateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportCertificate, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) ImportCertificateAsync(ctx workflow.Context, input *databasemigrationservice.ImportCertificateInput) *DatabasemigrationserviceImportCertificateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportCertificate, input)
    return &DatabasemigrationserviceImportCertificateResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) ListTagsForResource(ctx workflow.Context, input *databasemigrationservice.ListTagsForResourceInput) (*databasemigrationservice.ListTagsForResourceOutput, error) {
    var output databasemigrationservice.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) ListTagsForResourceAsync(ctx workflow.Context, input *databasemigrationservice.ListTagsForResourceInput) *DatabasemigrationserviceListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &DatabasemigrationserviceListTagsForResourceResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) ModifyEndpoint(ctx workflow.Context, input *databasemigrationservice.ModifyEndpointInput) (*databasemigrationservice.ModifyEndpointOutput, error) {
    var output databasemigrationservice.ModifyEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) ModifyEndpointAsync(ctx workflow.Context, input *databasemigrationservice.ModifyEndpointInput) *DatabasemigrationserviceModifyEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyEndpoint, input)
    return &DatabasemigrationserviceModifyEndpointResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) ModifyEventSubscription(ctx workflow.Context, input *databasemigrationservice.ModifyEventSubscriptionInput) (*databasemigrationservice.ModifyEventSubscriptionOutput, error) {
    var output databasemigrationservice.ModifyEventSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyEventSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) ModifyEventSubscriptionAsync(ctx workflow.Context, input *databasemigrationservice.ModifyEventSubscriptionInput) *DatabasemigrationserviceModifyEventSubscriptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyEventSubscription, input)
    return &DatabasemigrationserviceModifyEventSubscriptionResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) ModifyReplicationInstance(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationInstanceInput) (*databasemigrationservice.ModifyReplicationInstanceOutput, error) {
    var output databasemigrationservice.ModifyReplicationInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyReplicationInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) ModifyReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationInstanceInput) *DatabasemigrationserviceModifyReplicationInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyReplicationInstance, input)
    return &DatabasemigrationserviceModifyReplicationInstanceResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) ModifyReplicationSubnetGroup(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationSubnetGroupInput) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error) {
    var output databasemigrationservice.ModifyReplicationSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyReplicationSubnetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) ModifyReplicationSubnetGroupAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationSubnetGroupInput) *DatabasemigrationserviceModifyReplicationSubnetGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyReplicationSubnetGroup, input)
    return &DatabasemigrationserviceModifyReplicationSubnetGroupResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) ModifyReplicationTask(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationTaskInput) (*databasemigrationservice.ModifyReplicationTaskOutput, error) {
    var output databasemigrationservice.ModifyReplicationTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyReplicationTask, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) ModifyReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.ModifyReplicationTaskInput) *DatabasemigrationserviceModifyReplicationTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ModifyReplicationTask, input)
    return &DatabasemigrationserviceModifyReplicationTaskResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) RebootReplicationInstance(ctx workflow.Context, input *databasemigrationservice.RebootReplicationInstanceInput) (*databasemigrationservice.RebootReplicationInstanceOutput, error) {
    var output databasemigrationservice.RebootReplicationInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RebootReplicationInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) RebootReplicationInstanceAsync(ctx workflow.Context, input *databasemigrationservice.RebootReplicationInstanceInput) *DatabasemigrationserviceRebootReplicationInstanceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RebootReplicationInstance, input)
    return &DatabasemigrationserviceRebootReplicationInstanceResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) RefreshSchemas(ctx workflow.Context, input *databasemigrationservice.RefreshSchemasInput) (*databasemigrationservice.RefreshSchemasOutput, error) {
    var output databasemigrationservice.RefreshSchemasOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RefreshSchemas, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) RefreshSchemasAsync(ctx workflow.Context, input *databasemigrationservice.RefreshSchemasInput) *DatabasemigrationserviceRefreshSchemasResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RefreshSchemas, input)
    return &DatabasemigrationserviceRefreshSchemasResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) ReloadTables(ctx workflow.Context, input *databasemigrationservice.ReloadTablesInput) (*databasemigrationservice.ReloadTablesOutput, error) {
    var output databasemigrationservice.ReloadTablesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReloadTables, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) ReloadTablesAsync(ctx workflow.Context, input *databasemigrationservice.ReloadTablesInput) *DatabasemigrationserviceReloadTablesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReloadTables, input)
    return &DatabasemigrationserviceReloadTablesResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) RemoveTagsFromResource(ctx workflow.Context, input *databasemigrationservice.RemoveTagsFromResourceInput) (*databasemigrationservice.RemoveTagsFromResourceOutput, error) {
    var output databasemigrationservice.RemoveTagsFromResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) RemoveTagsFromResourceAsync(ctx workflow.Context, input *databasemigrationservice.RemoveTagsFromResourceInput) *DatabasemigrationserviceRemoveTagsFromResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromResource, input)
    return &DatabasemigrationserviceRemoveTagsFromResourceResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) StartReplicationTask(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskInput) (*databasemigrationservice.StartReplicationTaskOutput, error) {
    var output databasemigrationservice.StartReplicationTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartReplicationTask, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) StartReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskInput) *DatabasemigrationserviceStartReplicationTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartReplicationTask, input)
    return &DatabasemigrationserviceStartReplicationTaskResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) StartReplicationTaskAssessment(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentInput) (*databasemigrationservice.StartReplicationTaskAssessmentOutput, error) {
    var output databasemigrationservice.StartReplicationTaskAssessmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartReplicationTaskAssessment, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) StartReplicationTaskAssessmentAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentInput) *DatabasemigrationserviceStartReplicationTaskAssessmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartReplicationTaskAssessment, input)
    return &DatabasemigrationserviceStartReplicationTaskAssessmentResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) StartReplicationTaskAssessmentRun(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) (*databasemigrationservice.StartReplicationTaskAssessmentRunOutput, error) {
    var output databasemigrationservice.StartReplicationTaskAssessmentRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartReplicationTaskAssessmentRun, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) StartReplicationTaskAssessmentRunAsync(ctx workflow.Context, input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) *DatabasemigrationserviceStartReplicationTaskAssessmentRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartReplicationTaskAssessmentRun, input)
    return &DatabasemigrationserviceStartReplicationTaskAssessmentRunResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) StopReplicationTask(ctx workflow.Context, input *databasemigrationservice.StopReplicationTaskInput) (*databasemigrationservice.StopReplicationTaskOutput, error) {
    var output databasemigrationservice.StopReplicationTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopReplicationTask, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) StopReplicationTaskAsync(ctx workflow.Context, input *databasemigrationservice.StopReplicationTaskInput) *DatabasemigrationserviceStopReplicationTaskResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopReplicationTask, input)
    return &DatabasemigrationserviceStopReplicationTaskResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) TestConnection(ctx workflow.Context, input *databasemigrationservice.TestConnectionInput) (*databasemigrationservice.TestConnectionOutput, error) {
    var output databasemigrationservice.TestConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TestConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *DatabaseMigrationServiceStub) TestConnectionAsync(ctx workflow.Context, input *databasemigrationservice.TestConnectionInput) *DatabasemigrationserviceTestConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TestConnection, input)
    return &DatabasemigrationserviceTestConnectionResult{Result: future}
}

func (a *DatabaseMigrationServiceStub) WaitUntilEndpointDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilEndpointDeleted, input).Get(ctx, nil)
}

func (a *DatabaseMigrationServiceStub) WaitUntilEndpointDeletedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeEndpointsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilEndpointDeleted, input)
}


func (a *DatabaseMigrationServiceStub) WaitUntilReplicationInstanceAvailable(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationInstanceAvailable, input).Get(ctx, nil)
}

func (a *DatabaseMigrationServiceStub) WaitUntilReplicationInstanceAvailableAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationInstanceAvailable, input)
}


func (a *DatabaseMigrationServiceStub) WaitUntilReplicationInstanceDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationInstanceDeleted, input).Get(ctx, nil)
}

func (a *DatabaseMigrationServiceStub) WaitUntilReplicationInstanceDeletedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationInstanceDeleted, input)
}


func (a *DatabaseMigrationServiceStub) WaitUntilReplicationTaskDeleted(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationTaskDeleted, input).Get(ctx, nil)
}

func (a *DatabaseMigrationServiceStub) WaitUntilReplicationTaskDeletedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationTaskDeleted, input)
}


func (a *DatabaseMigrationServiceStub) WaitUntilReplicationTaskReady(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationTaskReady, input).Get(ctx, nil)
}

func (a *DatabaseMigrationServiceStub) WaitUntilReplicationTaskReadyAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationTaskReady, input)
}


func (a *DatabaseMigrationServiceStub) WaitUntilReplicationTaskRunning(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationTaskRunning, input).Get(ctx, nil)
}

func (a *DatabaseMigrationServiceStub) WaitUntilReplicationTaskRunningAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationTaskRunning, input)
}


func (a *DatabaseMigrationServiceStub) WaitUntilReplicationTaskStopped(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationTaskStopped, input).Get(ctx, nil)
}

func (a *DatabaseMigrationServiceStub) WaitUntilReplicationTaskStoppedAsync(ctx workflow.Context, input *databasemigrationservice.DescribeReplicationTasksInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilReplicationTaskStopped, input)
}


func (a *DatabaseMigrationServiceStub) WaitUntilTestConnectionSucceeds(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilTestConnectionSucceeds, input).Get(ctx, nil)
}

func (a *DatabaseMigrationServiceStub) WaitUntilTestConnectionSucceedsAsync(ctx workflow.Context, input *databasemigrationservice.DescribeConnectionsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilTestConnectionSucceeds, input)
}

