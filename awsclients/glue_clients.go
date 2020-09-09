package awsclients

import (
	"github.com/aws/aws-sdk-go/service/glue"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type GlueClient interface {
       BatchCreatePartition(ctx workflow.Context, input *glue.BatchCreatePartitionInput) (*glue.BatchCreatePartitionOutput, error)
       BatchCreatePartitionAsync(ctx workflow.Context, input *glue.BatchCreatePartitionInput) *GlueBatchCreatePartitionResult

       BatchDeleteConnection(ctx workflow.Context, input *glue.BatchDeleteConnectionInput) (*glue.BatchDeleteConnectionOutput, error)
       BatchDeleteConnectionAsync(ctx workflow.Context, input *glue.BatchDeleteConnectionInput) *GlueBatchDeleteConnectionResult

       BatchDeletePartition(ctx workflow.Context, input *glue.BatchDeletePartitionInput) (*glue.BatchDeletePartitionOutput, error)
       BatchDeletePartitionAsync(ctx workflow.Context, input *glue.BatchDeletePartitionInput) *GlueBatchDeletePartitionResult

       BatchDeleteTable(ctx workflow.Context, input *glue.BatchDeleteTableInput) (*glue.BatchDeleteTableOutput, error)
       BatchDeleteTableAsync(ctx workflow.Context, input *glue.BatchDeleteTableInput) *GlueBatchDeleteTableResult

       BatchDeleteTableVersion(ctx workflow.Context, input *glue.BatchDeleteTableVersionInput) (*glue.BatchDeleteTableVersionOutput, error)
       BatchDeleteTableVersionAsync(ctx workflow.Context, input *glue.BatchDeleteTableVersionInput) *GlueBatchDeleteTableVersionResult

       BatchGetCrawlers(ctx workflow.Context, input *glue.BatchGetCrawlersInput) (*glue.BatchGetCrawlersOutput, error)
       BatchGetCrawlersAsync(ctx workflow.Context, input *glue.BatchGetCrawlersInput) *GlueBatchGetCrawlersResult

       BatchGetDevEndpoints(ctx workflow.Context, input *glue.BatchGetDevEndpointsInput) (*glue.BatchGetDevEndpointsOutput, error)
       BatchGetDevEndpointsAsync(ctx workflow.Context, input *glue.BatchGetDevEndpointsInput) *GlueBatchGetDevEndpointsResult

       BatchGetJobs(ctx workflow.Context, input *glue.BatchGetJobsInput) (*glue.BatchGetJobsOutput, error)
       BatchGetJobsAsync(ctx workflow.Context, input *glue.BatchGetJobsInput) *GlueBatchGetJobsResult

       BatchGetPartition(ctx workflow.Context, input *glue.BatchGetPartitionInput) (*glue.BatchGetPartitionOutput, error)
       BatchGetPartitionAsync(ctx workflow.Context, input *glue.BatchGetPartitionInput) *GlueBatchGetPartitionResult

       BatchGetTriggers(ctx workflow.Context, input *glue.BatchGetTriggersInput) (*glue.BatchGetTriggersOutput, error)
       BatchGetTriggersAsync(ctx workflow.Context, input *glue.BatchGetTriggersInput) *GlueBatchGetTriggersResult

       BatchGetWorkflows(ctx workflow.Context, input *glue.BatchGetWorkflowsInput) (*glue.BatchGetWorkflowsOutput, error)
       BatchGetWorkflowsAsync(ctx workflow.Context, input *glue.BatchGetWorkflowsInput) *GlueBatchGetWorkflowsResult

       BatchStopJobRun(ctx workflow.Context, input *glue.BatchStopJobRunInput) (*glue.BatchStopJobRunOutput, error)
       BatchStopJobRunAsync(ctx workflow.Context, input *glue.BatchStopJobRunInput) *GlueBatchStopJobRunResult

       CancelMLTaskRun(ctx workflow.Context, input *glue.CancelMLTaskRunInput) (*glue.CancelMLTaskRunOutput, error)
       CancelMLTaskRunAsync(ctx workflow.Context, input *glue.CancelMLTaskRunInput) *GlueCancelMLTaskRunResult

       CreateClassifier(ctx workflow.Context, input *glue.CreateClassifierInput) (*glue.CreateClassifierOutput, error)
       CreateClassifierAsync(ctx workflow.Context, input *glue.CreateClassifierInput) *GlueCreateClassifierResult

       CreateConnection(ctx workflow.Context, input *glue.CreateConnectionInput) (*glue.CreateConnectionOutput, error)
       CreateConnectionAsync(ctx workflow.Context, input *glue.CreateConnectionInput) *GlueCreateConnectionResult

       CreateCrawler(ctx workflow.Context, input *glue.CreateCrawlerInput) (*glue.CreateCrawlerOutput, error)
       CreateCrawlerAsync(ctx workflow.Context, input *glue.CreateCrawlerInput) *GlueCreateCrawlerResult

       CreateDatabase(ctx workflow.Context, input *glue.CreateDatabaseInput) (*glue.CreateDatabaseOutput, error)
       CreateDatabaseAsync(ctx workflow.Context, input *glue.CreateDatabaseInput) *GlueCreateDatabaseResult

       CreateDevEndpoint(ctx workflow.Context, input *glue.CreateDevEndpointInput) (*glue.CreateDevEndpointOutput, error)
       CreateDevEndpointAsync(ctx workflow.Context, input *glue.CreateDevEndpointInput) *GlueCreateDevEndpointResult

       CreateJob(ctx workflow.Context, input *glue.CreateJobInput) (*glue.CreateJobOutput, error)
       CreateJobAsync(ctx workflow.Context, input *glue.CreateJobInput) *GlueCreateJobResult

       CreateMLTransform(ctx workflow.Context, input *glue.CreateMLTransformInput) (*glue.CreateMLTransformOutput, error)
       CreateMLTransformAsync(ctx workflow.Context, input *glue.CreateMLTransformInput) *GlueCreateMLTransformResult

       CreatePartition(ctx workflow.Context, input *glue.CreatePartitionInput) (*glue.CreatePartitionOutput, error)
       CreatePartitionAsync(ctx workflow.Context, input *glue.CreatePartitionInput) *GlueCreatePartitionResult

       CreateScript(ctx workflow.Context, input *glue.CreateScriptInput) (*glue.CreateScriptOutput, error)
       CreateScriptAsync(ctx workflow.Context, input *glue.CreateScriptInput) *GlueCreateScriptResult

       CreateSecurityConfiguration(ctx workflow.Context, input *glue.CreateSecurityConfigurationInput) (*glue.CreateSecurityConfigurationOutput, error)
       CreateSecurityConfigurationAsync(ctx workflow.Context, input *glue.CreateSecurityConfigurationInput) *GlueCreateSecurityConfigurationResult

       CreateTable(ctx workflow.Context, input *glue.CreateTableInput) (*glue.CreateTableOutput, error)
       CreateTableAsync(ctx workflow.Context, input *glue.CreateTableInput) *GlueCreateTableResult

       CreateTrigger(ctx workflow.Context, input *glue.CreateTriggerInput) (*glue.CreateTriggerOutput, error)
       CreateTriggerAsync(ctx workflow.Context, input *glue.CreateTriggerInput) *GlueCreateTriggerResult

       CreateUserDefinedFunction(ctx workflow.Context, input *glue.CreateUserDefinedFunctionInput) (*glue.CreateUserDefinedFunctionOutput, error)
       CreateUserDefinedFunctionAsync(ctx workflow.Context, input *glue.CreateUserDefinedFunctionInput) *GlueCreateUserDefinedFunctionResult

       CreateWorkflow(ctx workflow.Context, input *glue.CreateWorkflowInput) (*glue.CreateWorkflowOutput, error)
       CreateWorkflowAsync(ctx workflow.Context, input *glue.CreateWorkflowInput) *GlueCreateWorkflowResult

       DeleteClassifier(ctx workflow.Context, input *glue.DeleteClassifierInput) (*glue.DeleteClassifierOutput, error)
       DeleteClassifierAsync(ctx workflow.Context, input *glue.DeleteClassifierInput) *GlueDeleteClassifierResult

       DeleteColumnStatisticsForPartition(ctx workflow.Context, input *glue.DeleteColumnStatisticsForPartitionInput) (*glue.DeleteColumnStatisticsForPartitionOutput, error)
       DeleteColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.DeleteColumnStatisticsForPartitionInput) *GlueDeleteColumnStatisticsForPartitionResult

       DeleteColumnStatisticsForTable(ctx workflow.Context, input *glue.DeleteColumnStatisticsForTableInput) (*glue.DeleteColumnStatisticsForTableOutput, error)
       DeleteColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.DeleteColumnStatisticsForTableInput) *GlueDeleteColumnStatisticsForTableResult

       DeleteConnection(ctx workflow.Context, input *glue.DeleteConnectionInput) (*glue.DeleteConnectionOutput, error)
       DeleteConnectionAsync(ctx workflow.Context, input *glue.DeleteConnectionInput) *GlueDeleteConnectionResult

       DeleteCrawler(ctx workflow.Context, input *glue.DeleteCrawlerInput) (*glue.DeleteCrawlerOutput, error)
       DeleteCrawlerAsync(ctx workflow.Context, input *glue.DeleteCrawlerInput) *GlueDeleteCrawlerResult

       DeleteDatabase(ctx workflow.Context, input *glue.DeleteDatabaseInput) (*glue.DeleteDatabaseOutput, error)
       DeleteDatabaseAsync(ctx workflow.Context, input *glue.DeleteDatabaseInput) *GlueDeleteDatabaseResult

       DeleteDevEndpoint(ctx workflow.Context, input *glue.DeleteDevEndpointInput) (*glue.DeleteDevEndpointOutput, error)
       DeleteDevEndpointAsync(ctx workflow.Context, input *glue.DeleteDevEndpointInput) *GlueDeleteDevEndpointResult

       DeleteJob(ctx workflow.Context, input *glue.DeleteJobInput) (*glue.DeleteJobOutput, error)
       DeleteJobAsync(ctx workflow.Context, input *glue.DeleteJobInput) *GlueDeleteJobResult

       DeleteMLTransform(ctx workflow.Context, input *glue.DeleteMLTransformInput) (*glue.DeleteMLTransformOutput, error)
       DeleteMLTransformAsync(ctx workflow.Context, input *glue.DeleteMLTransformInput) *GlueDeleteMLTransformResult

       DeletePartition(ctx workflow.Context, input *glue.DeletePartitionInput) (*glue.DeletePartitionOutput, error)
       DeletePartitionAsync(ctx workflow.Context, input *glue.DeletePartitionInput) *GlueDeletePartitionResult

       DeleteResourcePolicy(ctx workflow.Context, input *glue.DeleteResourcePolicyInput) (*glue.DeleteResourcePolicyOutput, error)
       DeleteResourcePolicyAsync(ctx workflow.Context, input *glue.DeleteResourcePolicyInput) *GlueDeleteResourcePolicyResult

       DeleteSecurityConfiguration(ctx workflow.Context, input *glue.DeleteSecurityConfigurationInput) (*glue.DeleteSecurityConfigurationOutput, error)
       DeleteSecurityConfigurationAsync(ctx workflow.Context, input *glue.DeleteSecurityConfigurationInput) *GlueDeleteSecurityConfigurationResult

       DeleteTable(ctx workflow.Context, input *glue.DeleteTableInput) (*glue.DeleteTableOutput, error)
       DeleteTableAsync(ctx workflow.Context, input *glue.DeleteTableInput) *GlueDeleteTableResult

       DeleteTableVersion(ctx workflow.Context, input *glue.DeleteTableVersionInput) (*glue.DeleteTableVersionOutput, error)
       DeleteTableVersionAsync(ctx workflow.Context, input *glue.DeleteTableVersionInput) *GlueDeleteTableVersionResult

       DeleteTrigger(ctx workflow.Context, input *glue.DeleteTriggerInput) (*glue.DeleteTriggerOutput, error)
       DeleteTriggerAsync(ctx workflow.Context, input *glue.DeleteTriggerInput) *GlueDeleteTriggerResult

       DeleteUserDefinedFunction(ctx workflow.Context, input *glue.DeleteUserDefinedFunctionInput) (*glue.DeleteUserDefinedFunctionOutput, error)
       DeleteUserDefinedFunctionAsync(ctx workflow.Context, input *glue.DeleteUserDefinedFunctionInput) *GlueDeleteUserDefinedFunctionResult

       DeleteWorkflow(ctx workflow.Context, input *glue.DeleteWorkflowInput) (*glue.DeleteWorkflowOutput, error)
       DeleteWorkflowAsync(ctx workflow.Context, input *glue.DeleteWorkflowInput) *GlueDeleteWorkflowResult

       GetCatalogImportStatus(ctx workflow.Context, input *glue.GetCatalogImportStatusInput) (*glue.GetCatalogImportStatusOutput, error)
       GetCatalogImportStatusAsync(ctx workflow.Context, input *glue.GetCatalogImportStatusInput) *GlueGetCatalogImportStatusResult

       GetClassifier(ctx workflow.Context, input *glue.GetClassifierInput) (*glue.GetClassifierOutput, error)
       GetClassifierAsync(ctx workflow.Context, input *glue.GetClassifierInput) *GlueGetClassifierResult

       GetClassifiers(ctx workflow.Context, input *glue.GetClassifiersInput) (*glue.GetClassifiersOutput, error)
       GetClassifiersAsync(ctx workflow.Context, input *glue.GetClassifiersInput) *GlueGetClassifiersResult

       GetColumnStatisticsForPartition(ctx workflow.Context, input *glue.GetColumnStatisticsForPartitionInput) (*glue.GetColumnStatisticsForPartitionOutput, error)
       GetColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.GetColumnStatisticsForPartitionInput) *GlueGetColumnStatisticsForPartitionResult

       GetColumnStatisticsForTable(ctx workflow.Context, input *glue.GetColumnStatisticsForTableInput) (*glue.GetColumnStatisticsForTableOutput, error)
       GetColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.GetColumnStatisticsForTableInput) *GlueGetColumnStatisticsForTableResult

       GetConnection(ctx workflow.Context, input *glue.GetConnectionInput) (*glue.GetConnectionOutput, error)
       GetConnectionAsync(ctx workflow.Context, input *glue.GetConnectionInput) *GlueGetConnectionResult

       GetConnections(ctx workflow.Context, input *glue.GetConnectionsInput) (*glue.GetConnectionsOutput, error)
       GetConnectionsAsync(ctx workflow.Context, input *glue.GetConnectionsInput) *GlueGetConnectionsResult

       GetCrawler(ctx workflow.Context, input *glue.GetCrawlerInput) (*glue.GetCrawlerOutput, error)
       GetCrawlerAsync(ctx workflow.Context, input *glue.GetCrawlerInput) *GlueGetCrawlerResult

       GetCrawlerMetrics(ctx workflow.Context, input *glue.GetCrawlerMetricsInput) (*glue.GetCrawlerMetricsOutput, error)
       GetCrawlerMetricsAsync(ctx workflow.Context, input *glue.GetCrawlerMetricsInput) *GlueGetCrawlerMetricsResult

       GetCrawlers(ctx workflow.Context, input *glue.GetCrawlersInput) (*glue.GetCrawlersOutput, error)
       GetCrawlersAsync(ctx workflow.Context, input *glue.GetCrawlersInput) *GlueGetCrawlersResult

       GetDataCatalogEncryptionSettings(ctx workflow.Context, input *glue.GetDataCatalogEncryptionSettingsInput) (*glue.GetDataCatalogEncryptionSettingsOutput, error)
       GetDataCatalogEncryptionSettingsAsync(ctx workflow.Context, input *glue.GetDataCatalogEncryptionSettingsInput) *GlueGetDataCatalogEncryptionSettingsResult

       GetDatabase(ctx workflow.Context, input *glue.GetDatabaseInput) (*glue.GetDatabaseOutput, error)
       GetDatabaseAsync(ctx workflow.Context, input *glue.GetDatabaseInput) *GlueGetDatabaseResult

       GetDatabases(ctx workflow.Context, input *glue.GetDatabasesInput) (*glue.GetDatabasesOutput, error)
       GetDatabasesAsync(ctx workflow.Context, input *glue.GetDatabasesInput) *GlueGetDatabasesResult

       GetDataflowGraph(ctx workflow.Context, input *glue.GetDataflowGraphInput) (*glue.GetDataflowGraphOutput, error)
       GetDataflowGraphAsync(ctx workflow.Context, input *glue.GetDataflowGraphInput) *GlueGetDataflowGraphResult

       GetDevEndpoint(ctx workflow.Context, input *glue.GetDevEndpointInput) (*glue.GetDevEndpointOutput, error)
       GetDevEndpointAsync(ctx workflow.Context, input *glue.GetDevEndpointInput) *GlueGetDevEndpointResult

       GetDevEndpoints(ctx workflow.Context, input *glue.GetDevEndpointsInput) (*glue.GetDevEndpointsOutput, error)
       GetDevEndpointsAsync(ctx workflow.Context, input *glue.GetDevEndpointsInput) *GlueGetDevEndpointsResult

       GetJob(ctx workflow.Context, input *glue.GetJobInput) (*glue.GetJobOutput, error)
       GetJobAsync(ctx workflow.Context, input *glue.GetJobInput) *GlueGetJobResult

       GetJobBookmark(ctx workflow.Context, input *glue.GetJobBookmarkInput) (*glue.GetJobBookmarkOutput, error)
       GetJobBookmarkAsync(ctx workflow.Context, input *glue.GetJobBookmarkInput) *GlueGetJobBookmarkResult

       GetJobRun(ctx workflow.Context, input *glue.GetJobRunInput) (*glue.GetJobRunOutput, error)
       GetJobRunAsync(ctx workflow.Context, input *glue.GetJobRunInput) *GlueGetJobRunResult

       GetJobRuns(ctx workflow.Context, input *glue.GetJobRunsInput) (*glue.GetJobRunsOutput, error)
       GetJobRunsAsync(ctx workflow.Context, input *glue.GetJobRunsInput) *GlueGetJobRunsResult

       GetJobs(ctx workflow.Context, input *glue.GetJobsInput) (*glue.GetJobsOutput, error)
       GetJobsAsync(ctx workflow.Context, input *glue.GetJobsInput) *GlueGetJobsResult

       GetMLTaskRun(ctx workflow.Context, input *glue.GetMLTaskRunInput) (*glue.GetMLTaskRunOutput, error)
       GetMLTaskRunAsync(ctx workflow.Context, input *glue.GetMLTaskRunInput) *GlueGetMLTaskRunResult

       GetMLTaskRuns(ctx workflow.Context, input *glue.GetMLTaskRunsInput) (*glue.GetMLTaskRunsOutput, error)
       GetMLTaskRunsAsync(ctx workflow.Context, input *glue.GetMLTaskRunsInput) *GlueGetMLTaskRunsResult

       GetMLTransform(ctx workflow.Context, input *glue.GetMLTransformInput) (*glue.GetMLTransformOutput, error)
       GetMLTransformAsync(ctx workflow.Context, input *glue.GetMLTransformInput) *GlueGetMLTransformResult

       GetMLTransforms(ctx workflow.Context, input *glue.GetMLTransformsInput) (*glue.GetMLTransformsOutput, error)
       GetMLTransformsAsync(ctx workflow.Context, input *glue.GetMLTransformsInput) *GlueGetMLTransformsResult

       GetMapping(ctx workflow.Context, input *glue.GetMappingInput) (*glue.GetMappingOutput, error)
       GetMappingAsync(ctx workflow.Context, input *glue.GetMappingInput) *GlueGetMappingResult

       GetPartition(ctx workflow.Context, input *glue.GetPartitionInput) (*glue.GetPartitionOutput, error)
       GetPartitionAsync(ctx workflow.Context, input *glue.GetPartitionInput) *GlueGetPartitionResult

       GetPartitions(ctx workflow.Context, input *glue.GetPartitionsInput) (*glue.GetPartitionsOutput, error)
       GetPartitionsAsync(ctx workflow.Context, input *glue.GetPartitionsInput) *GlueGetPartitionsResult

       GetPlan(ctx workflow.Context, input *glue.GetPlanInput) (*glue.GetPlanOutput, error)
       GetPlanAsync(ctx workflow.Context, input *glue.GetPlanInput) *GlueGetPlanResult

       GetResourcePolicies(ctx workflow.Context, input *glue.GetResourcePoliciesInput) (*glue.GetResourcePoliciesOutput, error)
       GetResourcePoliciesAsync(ctx workflow.Context, input *glue.GetResourcePoliciesInput) *GlueGetResourcePoliciesResult

       GetResourcePolicy(ctx workflow.Context, input *glue.GetResourcePolicyInput) (*glue.GetResourcePolicyOutput, error)
       GetResourcePolicyAsync(ctx workflow.Context, input *glue.GetResourcePolicyInput) *GlueGetResourcePolicyResult

       GetSecurityConfiguration(ctx workflow.Context, input *glue.GetSecurityConfigurationInput) (*glue.GetSecurityConfigurationOutput, error)
       GetSecurityConfigurationAsync(ctx workflow.Context, input *glue.GetSecurityConfigurationInput) *GlueGetSecurityConfigurationResult

       GetSecurityConfigurations(ctx workflow.Context, input *glue.GetSecurityConfigurationsInput) (*glue.GetSecurityConfigurationsOutput, error)
       GetSecurityConfigurationsAsync(ctx workflow.Context, input *glue.GetSecurityConfigurationsInput) *GlueGetSecurityConfigurationsResult

       GetTable(ctx workflow.Context, input *glue.GetTableInput) (*glue.GetTableOutput, error)
       GetTableAsync(ctx workflow.Context, input *glue.GetTableInput) *GlueGetTableResult

       GetTableVersion(ctx workflow.Context, input *glue.GetTableVersionInput) (*glue.GetTableVersionOutput, error)
       GetTableVersionAsync(ctx workflow.Context, input *glue.GetTableVersionInput) *GlueGetTableVersionResult

       GetTableVersions(ctx workflow.Context, input *glue.GetTableVersionsInput) (*glue.GetTableVersionsOutput, error)
       GetTableVersionsAsync(ctx workflow.Context, input *glue.GetTableVersionsInput) *GlueGetTableVersionsResult

       GetTables(ctx workflow.Context, input *glue.GetTablesInput) (*glue.GetTablesOutput, error)
       GetTablesAsync(ctx workflow.Context, input *glue.GetTablesInput) *GlueGetTablesResult

       GetTags(ctx workflow.Context, input *glue.GetTagsInput) (*glue.GetTagsOutput, error)
       GetTagsAsync(ctx workflow.Context, input *glue.GetTagsInput) *GlueGetTagsResult

       GetTrigger(ctx workflow.Context, input *glue.GetTriggerInput) (*glue.GetTriggerOutput, error)
       GetTriggerAsync(ctx workflow.Context, input *glue.GetTriggerInput) *GlueGetTriggerResult

       GetTriggers(ctx workflow.Context, input *glue.GetTriggersInput) (*glue.GetTriggersOutput, error)
       GetTriggersAsync(ctx workflow.Context, input *glue.GetTriggersInput) *GlueGetTriggersResult

       GetUserDefinedFunction(ctx workflow.Context, input *glue.GetUserDefinedFunctionInput) (*glue.GetUserDefinedFunctionOutput, error)
       GetUserDefinedFunctionAsync(ctx workflow.Context, input *glue.GetUserDefinedFunctionInput) *GlueGetUserDefinedFunctionResult

       GetUserDefinedFunctions(ctx workflow.Context, input *glue.GetUserDefinedFunctionsInput) (*glue.GetUserDefinedFunctionsOutput, error)
       GetUserDefinedFunctionsAsync(ctx workflow.Context, input *glue.GetUserDefinedFunctionsInput) *GlueGetUserDefinedFunctionsResult

       GetWorkflow(ctx workflow.Context, input *glue.GetWorkflowInput) (*glue.GetWorkflowOutput, error)
       GetWorkflowAsync(ctx workflow.Context, input *glue.GetWorkflowInput) *GlueGetWorkflowResult

       GetWorkflowRun(ctx workflow.Context, input *glue.GetWorkflowRunInput) (*glue.GetWorkflowRunOutput, error)
       GetWorkflowRunAsync(ctx workflow.Context, input *glue.GetWorkflowRunInput) *GlueGetWorkflowRunResult

       GetWorkflowRunProperties(ctx workflow.Context, input *glue.GetWorkflowRunPropertiesInput) (*glue.GetWorkflowRunPropertiesOutput, error)
       GetWorkflowRunPropertiesAsync(ctx workflow.Context, input *glue.GetWorkflowRunPropertiesInput) *GlueGetWorkflowRunPropertiesResult

       GetWorkflowRuns(ctx workflow.Context, input *glue.GetWorkflowRunsInput) (*glue.GetWorkflowRunsOutput, error)
       GetWorkflowRunsAsync(ctx workflow.Context, input *glue.GetWorkflowRunsInput) *GlueGetWorkflowRunsResult

       ImportCatalogToGlue(ctx workflow.Context, input *glue.ImportCatalogToGlueInput) (*glue.ImportCatalogToGlueOutput, error)
       ImportCatalogToGlueAsync(ctx workflow.Context, input *glue.ImportCatalogToGlueInput) *GlueImportCatalogToGlueResult

       ListCrawlers(ctx workflow.Context, input *glue.ListCrawlersInput) (*glue.ListCrawlersOutput, error)
       ListCrawlersAsync(ctx workflow.Context, input *glue.ListCrawlersInput) *GlueListCrawlersResult

       ListDevEndpoints(ctx workflow.Context, input *glue.ListDevEndpointsInput) (*glue.ListDevEndpointsOutput, error)
       ListDevEndpointsAsync(ctx workflow.Context, input *glue.ListDevEndpointsInput) *GlueListDevEndpointsResult

       ListJobs(ctx workflow.Context, input *glue.ListJobsInput) (*glue.ListJobsOutput, error)
       ListJobsAsync(ctx workflow.Context, input *glue.ListJobsInput) *GlueListJobsResult

       ListMLTransforms(ctx workflow.Context, input *glue.ListMLTransformsInput) (*glue.ListMLTransformsOutput, error)
       ListMLTransformsAsync(ctx workflow.Context, input *glue.ListMLTransformsInput) *GlueListMLTransformsResult

       ListTriggers(ctx workflow.Context, input *glue.ListTriggersInput) (*glue.ListTriggersOutput, error)
       ListTriggersAsync(ctx workflow.Context, input *glue.ListTriggersInput) *GlueListTriggersResult

       ListWorkflows(ctx workflow.Context, input *glue.ListWorkflowsInput) (*glue.ListWorkflowsOutput, error)
       ListWorkflowsAsync(ctx workflow.Context, input *glue.ListWorkflowsInput) *GlueListWorkflowsResult

       PutDataCatalogEncryptionSettings(ctx workflow.Context, input *glue.PutDataCatalogEncryptionSettingsInput) (*glue.PutDataCatalogEncryptionSettingsOutput, error)
       PutDataCatalogEncryptionSettingsAsync(ctx workflow.Context, input *glue.PutDataCatalogEncryptionSettingsInput) *GluePutDataCatalogEncryptionSettingsResult

       PutResourcePolicy(ctx workflow.Context, input *glue.PutResourcePolicyInput) (*glue.PutResourcePolicyOutput, error)
       PutResourcePolicyAsync(ctx workflow.Context, input *glue.PutResourcePolicyInput) *GluePutResourcePolicyResult

       PutWorkflowRunProperties(ctx workflow.Context, input *glue.PutWorkflowRunPropertiesInput) (*glue.PutWorkflowRunPropertiesOutput, error)
       PutWorkflowRunPropertiesAsync(ctx workflow.Context, input *glue.PutWorkflowRunPropertiesInput) *GluePutWorkflowRunPropertiesResult

       ResetJobBookmark(ctx workflow.Context, input *glue.ResetJobBookmarkInput) (*glue.ResetJobBookmarkOutput, error)
       ResetJobBookmarkAsync(ctx workflow.Context, input *glue.ResetJobBookmarkInput) *GlueResetJobBookmarkResult

       ResumeWorkflowRun(ctx workflow.Context, input *glue.ResumeWorkflowRunInput) (*glue.ResumeWorkflowRunOutput, error)
       ResumeWorkflowRunAsync(ctx workflow.Context, input *glue.ResumeWorkflowRunInput) *GlueResumeWorkflowRunResult

       SearchTables(ctx workflow.Context, input *glue.SearchTablesInput) (*glue.SearchTablesOutput, error)
       SearchTablesAsync(ctx workflow.Context, input *glue.SearchTablesInput) *GlueSearchTablesResult

       StartCrawler(ctx workflow.Context, input *glue.StartCrawlerInput) (*glue.StartCrawlerOutput, error)
       StartCrawlerAsync(ctx workflow.Context, input *glue.StartCrawlerInput) *GlueStartCrawlerResult

       StartCrawlerSchedule(ctx workflow.Context, input *glue.StartCrawlerScheduleInput) (*glue.StartCrawlerScheduleOutput, error)
       StartCrawlerScheduleAsync(ctx workflow.Context, input *glue.StartCrawlerScheduleInput) *GlueStartCrawlerScheduleResult

       StartExportLabelsTaskRun(ctx workflow.Context, input *glue.StartExportLabelsTaskRunInput) (*glue.StartExportLabelsTaskRunOutput, error)
       StartExportLabelsTaskRunAsync(ctx workflow.Context, input *glue.StartExportLabelsTaskRunInput) *GlueStartExportLabelsTaskRunResult

       StartImportLabelsTaskRun(ctx workflow.Context, input *glue.StartImportLabelsTaskRunInput) (*glue.StartImportLabelsTaskRunOutput, error)
       StartImportLabelsTaskRunAsync(ctx workflow.Context, input *glue.StartImportLabelsTaskRunInput) *GlueStartImportLabelsTaskRunResult

       StartJobRun(ctx workflow.Context, input *glue.StartJobRunInput) (*glue.StartJobRunOutput, error)
       StartJobRunAsync(ctx workflow.Context, input *glue.StartJobRunInput) *GlueStartJobRunResult

       StartMLEvaluationTaskRun(ctx workflow.Context, input *glue.StartMLEvaluationTaskRunInput) (*glue.StartMLEvaluationTaskRunOutput, error)
       StartMLEvaluationTaskRunAsync(ctx workflow.Context, input *glue.StartMLEvaluationTaskRunInput) *GlueStartMLEvaluationTaskRunResult

       StartMLLabelingSetGenerationTaskRun(ctx workflow.Context, input *glue.StartMLLabelingSetGenerationTaskRunInput) (*glue.StartMLLabelingSetGenerationTaskRunOutput, error)
       StartMLLabelingSetGenerationTaskRunAsync(ctx workflow.Context, input *glue.StartMLLabelingSetGenerationTaskRunInput) *GlueStartMLLabelingSetGenerationTaskRunResult

       StartTrigger(ctx workflow.Context, input *glue.StartTriggerInput) (*glue.StartTriggerOutput, error)
       StartTriggerAsync(ctx workflow.Context, input *glue.StartTriggerInput) *GlueStartTriggerResult

       StartWorkflowRun(ctx workflow.Context, input *glue.StartWorkflowRunInput) (*glue.StartWorkflowRunOutput, error)
       StartWorkflowRunAsync(ctx workflow.Context, input *glue.StartWorkflowRunInput) *GlueStartWorkflowRunResult

       StopCrawler(ctx workflow.Context, input *glue.StopCrawlerInput) (*glue.StopCrawlerOutput, error)
       StopCrawlerAsync(ctx workflow.Context, input *glue.StopCrawlerInput) *GlueStopCrawlerResult

       StopCrawlerSchedule(ctx workflow.Context, input *glue.StopCrawlerScheduleInput) (*glue.StopCrawlerScheduleOutput, error)
       StopCrawlerScheduleAsync(ctx workflow.Context, input *glue.StopCrawlerScheduleInput) *GlueStopCrawlerScheduleResult

       StopTrigger(ctx workflow.Context, input *glue.StopTriggerInput) (*glue.StopTriggerOutput, error)
       StopTriggerAsync(ctx workflow.Context, input *glue.StopTriggerInput) *GlueStopTriggerResult

       StopWorkflowRun(ctx workflow.Context, input *glue.StopWorkflowRunInput) (*glue.StopWorkflowRunOutput, error)
       StopWorkflowRunAsync(ctx workflow.Context, input *glue.StopWorkflowRunInput) *GlueStopWorkflowRunResult

       TagResource(ctx workflow.Context, input *glue.TagResourceInput) (*glue.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *glue.TagResourceInput) *GlueTagResourceResult

       UntagResource(ctx workflow.Context, input *glue.UntagResourceInput) (*glue.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *glue.UntagResourceInput) *GlueUntagResourceResult

       UpdateClassifier(ctx workflow.Context, input *glue.UpdateClassifierInput) (*glue.UpdateClassifierOutput, error)
       UpdateClassifierAsync(ctx workflow.Context, input *glue.UpdateClassifierInput) *GlueUpdateClassifierResult

       UpdateColumnStatisticsForPartition(ctx workflow.Context, input *glue.UpdateColumnStatisticsForPartitionInput) (*glue.UpdateColumnStatisticsForPartitionOutput, error)
       UpdateColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.UpdateColumnStatisticsForPartitionInput) *GlueUpdateColumnStatisticsForPartitionResult

       UpdateColumnStatisticsForTable(ctx workflow.Context, input *glue.UpdateColumnStatisticsForTableInput) (*glue.UpdateColumnStatisticsForTableOutput, error)
       UpdateColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.UpdateColumnStatisticsForTableInput) *GlueUpdateColumnStatisticsForTableResult

       UpdateConnection(ctx workflow.Context, input *glue.UpdateConnectionInput) (*glue.UpdateConnectionOutput, error)
       UpdateConnectionAsync(ctx workflow.Context, input *glue.UpdateConnectionInput) *GlueUpdateConnectionResult

       UpdateCrawler(ctx workflow.Context, input *glue.UpdateCrawlerInput) (*glue.UpdateCrawlerOutput, error)
       UpdateCrawlerAsync(ctx workflow.Context, input *glue.UpdateCrawlerInput) *GlueUpdateCrawlerResult

       UpdateCrawlerSchedule(ctx workflow.Context, input *glue.UpdateCrawlerScheduleInput) (*glue.UpdateCrawlerScheduleOutput, error)
       UpdateCrawlerScheduleAsync(ctx workflow.Context, input *glue.UpdateCrawlerScheduleInput) *GlueUpdateCrawlerScheduleResult

       UpdateDatabase(ctx workflow.Context, input *glue.UpdateDatabaseInput) (*glue.UpdateDatabaseOutput, error)
       UpdateDatabaseAsync(ctx workflow.Context, input *glue.UpdateDatabaseInput) *GlueUpdateDatabaseResult

       UpdateDevEndpoint(ctx workflow.Context, input *glue.UpdateDevEndpointInput) (*glue.UpdateDevEndpointOutput, error)
       UpdateDevEndpointAsync(ctx workflow.Context, input *glue.UpdateDevEndpointInput) *GlueUpdateDevEndpointResult

       UpdateJob(ctx workflow.Context, input *glue.UpdateJobInput) (*glue.UpdateJobOutput, error)
       UpdateJobAsync(ctx workflow.Context, input *glue.UpdateJobInput) *GlueUpdateJobResult

       UpdateMLTransform(ctx workflow.Context, input *glue.UpdateMLTransformInput) (*glue.UpdateMLTransformOutput, error)
       UpdateMLTransformAsync(ctx workflow.Context, input *glue.UpdateMLTransformInput) *GlueUpdateMLTransformResult

       UpdatePartition(ctx workflow.Context, input *glue.UpdatePartitionInput) (*glue.UpdatePartitionOutput, error)
       UpdatePartitionAsync(ctx workflow.Context, input *glue.UpdatePartitionInput) *GlueUpdatePartitionResult

       UpdateTable(ctx workflow.Context, input *glue.UpdateTableInput) (*glue.UpdateTableOutput, error)
       UpdateTableAsync(ctx workflow.Context, input *glue.UpdateTableInput) *GlueUpdateTableResult

       UpdateTrigger(ctx workflow.Context, input *glue.UpdateTriggerInput) (*glue.UpdateTriggerOutput, error)
       UpdateTriggerAsync(ctx workflow.Context, input *glue.UpdateTriggerInput) *GlueUpdateTriggerResult

       UpdateUserDefinedFunction(ctx workflow.Context, input *glue.UpdateUserDefinedFunctionInput) (*glue.UpdateUserDefinedFunctionOutput, error)
       UpdateUserDefinedFunctionAsync(ctx workflow.Context, input *glue.UpdateUserDefinedFunctionInput) *GlueUpdateUserDefinedFunctionResult

       UpdateWorkflow(ctx workflow.Context, input *glue.UpdateWorkflowInput) (*glue.UpdateWorkflowOutput, error)
       UpdateWorkflowAsync(ctx workflow.Context, input *glue.UpdateWorkflowInput) *GlueUpdateWorkflowResult
}

type GlueBatchCreatePartitionResult struct {
	Result workflow.Future
}

func (r *GlueBatchCreatePartitionResult) Get(ctx workflow.Context) (*glue.BatchCreatePartitionOutput, error) {
    var output glue.BatchCreatePartitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueBatchDeleteConnectionResult struct {
	Result workflow.Future
}

func (r *GlueBatchDeleteConnectionResult) Get(ctx workflow.Context) (*glue.BatchDeleteConnectionOutput, error) {
    var output glue.BatchDeleteConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueBatchDeletePartitionResult struct {
	Result workflow.Future
}

func (r *GlueBatchDeletePartitionResult) Get(ctx workflow.Context) (*glue.BatchDeletePartitionOutput, error) {
    var output glue.BatchDeletePartitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueBatchDeleteTableResult struct {
	Result workflow.Future
}

func (r *GlueBatchDeleteTableResult) Get(ctx workflow.Context) (*glue.BatchDeleteTableOutput, error) {
    var output glue.BatchDeleteTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueBatchDeleteTableVersionResult struct {
	Result workflow.Future
}

func (r *GlueBatchDeleteTableVersionResult) Get(ctx workflow.Context) (*glue.BatchDeleteTableVersionOutput, error) {
    var output glue.BatchDeleteTableVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueBatchGetCrawlersResult struct {
	Result workflow.Future
}

func (r *GlueBatchGetCrawlersResult) Get(ctx workflow.Context) (*glue.BatchGetCrawlersOutput, error) {
    var output glue.BatchGetCrawlersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueBatchGetDevEndpointsResult struct {
	Result workflow.Future
}

func (r *GlueBatchGetDevEndpointsResult) Get(ctx workflow.Context) (*glue.BatchGetDevEndpointsOutput, error) {
    var output glue.BatchGetDevEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueBatchGetJobsResult struct {
	Result workflow.Future
}

func (r *GlueBatchGetJobsResult) Get(ctx workflow.Context) (*glue.BatchGetJobsOutput, error) {
    var output glue.BatchGetJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueBatchGetPartitionResult struct {
	Result workflow.Future
}

func (r *GlueBatchGetPartitionResult) Get(ctx workflow.Context) (*glue.BatchGetPartitionOutput, error) {
    var output glue.BatchGetPartitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueBatchGetTriggersResult struct {
	Result workflow.Future
}

func (r *GlueBatchGetTriggersResult) Get(ctx workflow.Context) (*glue.BatchGetTriggersOutput, error) {
    var output glue.BatchGetTriggersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueBatchGetWorkflowsResult struct {
	Result workflow.Future
}

func (r *GlueBatchGetWorkflowsResult) Get(ctx workflow.Context) (*glue.BatchGetWorkflowsOutput, error) {
    var output glue.BatchGetWorkflowsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueBatchStopJobRunResult struct {
	Result workflow.Future
}

func (r *GlueBatchStopJobRunResult) Get(ctx workflow.Context) (*glue.BatchStopJobRunOutput, error) {
    var output glue.BatchStopJobRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCancelMLTaskRunResult struct {
	Result workflow.Future
}

func (r *GlueCancelMLTaskRunResult) Get(ctx workflow.Context) (*glue.CancelMLTaskRunOutput, error) {
    var output glue.CancelMLTaskRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateClassifierResult struct {
	Result workflow.Future
}

func (r *GlueCreateClassifierResult) Get(ctx workflow.Context) (*glue.CreateClassifierOutput, error) {
    var output glue.CreateClassifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateConnectionResult struct {
	Result workflow.Future
}

func (r *GlueCreateConnectionResult) Get(ctx workflow.Context) (*glue.CreateConnectionOutput, error) {
    var output glue.CreateConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateCrawlerResult struct {
	Result workflow.Future
}

func (r *GlueCreateCrawlerResult) Get(ctx workflow.Context) (*glue.CreateCrawlerOutput, error) {
    var output glue.CreateCrawlerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateDatabaseResult struct {
	Result workflow.Future
}

func (r *GlueCreateDatabaseResult) Get(ctx workflow.Context) (*glue.CreateDatabaseOutput, error) {
    var output glue.CreateDatabaseOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateDevEndpointResult struct {
	Result workflow.Future
}

func (r *GlueCreateDevEndpointResult) Get(ctx workflow.Context) (*glue.CreateDevEndpointOutput, error) {
    var output glue.CreateDevEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateJobResult struct {
	Result workflow.Future
}

func (r *GlueCreateJobResult) Get(ctx workflow.Context) (*glue.CreateJobOutput, error) {
    var output glue.CreateJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateMLTransformResult struct {
	Result workflow.Future
}

func (r *GlueCreateMLTransformResult) Get(ctx workflow.Context) (*glue.CreateMLTransformOutput, error) {
    var output glue.CreateMLTransformOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreatePartitionResult struct {
	Result workflow.Future
}

func (r *GlueCreatePartitionResult) Get(ctx workflow.Context) (*glue.CreatePartitionOutput, error) {
    var output glue.CreatePartitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateScriptResult struct {
	Result workflow.Future
}

func (r *GlueCreateScriptResult) Get(ctx workflow.Context) (*glue.CreateScriptOutput, error) {
    var output glue.CreateScriptOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateSecurityConfigurationResult struct {
	Result workflow.Future
}

func (r *GlueCreateSecurityConfigurationResult) Get(ctx workflow.Context) (*glue.CreateSecurityConfigurationOutput, error) {
    var output glue.CreateSecurityConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateTableResult struct {
	Result workflow.Future
}

func (r *GlueCreateTableResult) Get(ctx workflow.Context) (*glue.CreateTableOutput, error) {
    var output glue.CreateTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateTriggerResult struct {
	Result workflow.Future
}

func (r *GlueCreateTriggerResult) Get(ctx workflow.Context) (*glue.CreateTriggerOutput, error) {
    var output glue.CreateTriggerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateUserDefinedFunctionResult struct {
	Result workflow.Future
}

func (r *GlueCreateUserDefinedFunctionResult) Get(ctx workflow.Context) (*glue.CreateUserDefinedFunctionOutput, error) {
    var output glue.CreateUserDefinedFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueCreateWorkflowResult struct {
	Result workflow.Future
}

func (r *GlueCreateWorkflowResult) Get(ctx workflow.Context) (*glue.CreateWorkflowOutput, error) {
    var output glue.CreateWorkflowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteClassifierResult struct {
	Result workflow.Future
}

func (r *GlueDeleteClassifierResult) Get(ctx workflow.Context) (*glue.DeleteClassifierOutput, error) {
    var output glue.DeleteClassifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteColumnStatisticsForPartitionResult struct {
	Result workflow.Future
}

func (r *GlueDeleteColumnStatisticsForPartitionResult) Get(ctx workflow.Context) (*glue.DeleteColumnStatisticsForPartitionOutput, error) {
    var output glue.DeleteColumnStatisticsForPartitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteColumnStatisticsForTableResult struct {
	Result workflow.Future
}

func (r *GlueDeleteColumnStatisticsForTableResult) Get(ctx workflow.Context) (*glue.DeleteColumnStatisticsForTableOutput, error) {
    var output glue.DeleteColumnStatisticsForTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteConnectionResult struct {
	Result workflow.Future
}

func (r *GlueDeleteConnectionResult) Get(ctx workflow.Context) (*glue.DeleteConnectionOutput, error) {
    var output glue.DeleteConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteCrawlerResult struct {
	Result workflow.Future
}

func (r *GlueDeleteCrawlerResult) Get(ctx workflow.Context) (*glue.DeleteCrawlerOutput, error) {
    var output glue.DeleteCrawlerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteDatabaseResult struct {
	Result workflow.Future
}

func (r *GlueDeleteDatabaseResult) Get(ctx workflow.Context) (*glue.DeleteDatabaseOutput, error) {
    var output glue.DeleteDatabaseOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteDevEndpointResult struct {
	Result workflow.Future
}

func (r *GlueDeleteDevEndpointResult) Get(ctx workflow.Context) (*glue.DeleteDevEndpointOutput, error) {
    var output glue.DeleteDevEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteJobResult struct {
	Result workflow.Future
}

func (r *GlueDeleteJobResult) Get(ctx workflow.Context) (*glue.DeleteJobOutput, error) {
    var output glue.DeleteJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteMLTransformResult struct {
	Result workflow.Future
}

func (r *GlueDeleteMLTransformResult) Get(ctx workflow.Context) (*glue.DeleteMLTransformOutput, error) {
    var output glue.DeleteMLTransformOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeletePartitionResult struct {
	Result workflow.Future
}

func (r *GlueDeletePartitionResult) Get(ctx workflow.Context) (*glue.DeletePartitionOutput, error) {
    var output glue.DeletePartitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteResourcePolicyResult struct {
	Result workflow.Future
}

func (r *GlueDeleteResourcePolicyResult) Get(ctx workflow.Context) (*glue.DeleteResourcePolicyOutput, error) {
    var output glue.DeleteResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteSecurityConfigurationResult struct {
	Result workflow.Future
}

func (r *GlueDeleteSecurityConfigurationResult) Get(ctx workflow.Context) (*glue.DeleteSecurityConfigurationOutput, error) {
    var output glue.DeleteSecurityConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteTableResult struct {
	Result workflow.Future
}

func (r *GlueDeleteTableResult) Get(ctx workflow.Context) (*glue.DeleteTableOutput, error) {
    var output glue.DeleteTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteTableVersionResult struct {
	Result workflow.Future
}

func (r *GlueDeleteTableVersionResult) Get(ctx workflow.Context) (*glue.DeleteTableVersionOutput, error) {
    var output glue.DeleteTableVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteTriggerResult struct {
	Result workflow.Future
}

func (r *GlueDeleteTriggerResult) Get(ctx workflow.Context) (*glue.DeleteTriggerOutput, error) {
    var output glue.DeleteTriggerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteUserDefinedFunctionResult struct {
	Result workflow.Future
}

func (r *GlueDeleteUserDefinedFunctionResult) Get(ctx workflow.Context) (*glue.DeleteUserDefinedFunctionOutput, error) {
    var output glue.DeleteUserDefinedFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueDeleteWorkflowResult struct {
	Result workflow.Future
}

func (r *GlueDeleteWorkflowResult) Get(ctx workflow.Context) (*glue.DeleteWorkflowOutput, error) {
    var output glue.DeleteWorkflowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetCatalogImportStatusResult struct {
	Result workflow.Future
}

func (r *GlueGetCatalogImportStatusResult) Get(ctx workflow.Context) (*glue.GetCatalogImportStatusOutput, error) {
    var output glue.GetCatalogImportStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetClassifierResult struct {
	Result workflow.Future
}

func (r *GlueGetClassifierResult) Get(ctx workflow.Context) (*glue.GetClassifierOutput, error) {
    var output glue.GetClassifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetClassifiersResult struct {
	Result workflow.Future
}

func (r *GlueGetClassifiersResult) Get(ctx workflow.Context) (*glue.GetClassifiersOutput, error) {
    var output glue.GetClassifiersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetColumnStatisticsForPartitionResult struct {
	Result workflow.Future
}

func (r *GlueGetColumnStatisticsForPartitionResult) Get(ctx workflow.Context) (*glue.GetColumnStatisticsForPartitionOutput, error) {
    var output glue.GetColumnStatisticsForPartitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetColumnStatisticsForTableResult struct {
	Result workflow.Future
}

func (r *GlueGetColumnStatisticsForTableResult) Get(ctx workflow.Context) (*glue.GetColumnStatisticsForTableOutput, error) {
    var output glue.GetColumnStatisticsForTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetConnectionResult struct {
	Result workflow.Future
}

func (r *GlueGetConnectionResult) Get(ctx workflow.Context) (*glue.GetConnectionOutput, error) {
    var output glue.GetConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetConnectionsResult struct {
	Result workflow.Future
}

func (r *GlueGetConnectionsResult) Get(ctx workflow.Context) (*glue.GetConnectionsOutput, error) {
    var output glue.GetConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetCrawlerResult struct {
	Result workflow.Future
}

func (r *GlueGetCrawlerResult) Get(ctx workflow.Context) (*glue.GetCrawlerOutput, error) {
    var output glue.GetCrawlerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetCrawlerMetricsResult struct {
	Result workflow.Future
}

func (r *GlueGetCrawlerMetricsResult) Get(ctx workflow.Context) (*glue.GetCrawlerMetricsOutput, error) {
    var output glue.GetCrawlerMetricsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetCrawlersResult struct {
	Result workflow.Future
}

func (r *GlueGetCrawlersResult) Get(ctx workflow.Context) (*glue.GetCrawlersOutput, error) {
    var output glue.GetCrawlersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetDataCatalogEncryptionSettingsResult struct {
	Result workflow.Future
}

func (r *GlueGetDataCatalogEncryptionSettingsResult) Get(ctx workflow.Context) (*glue.GetDataCatalogEncryptionSettingsOutput, error) {
    var output glue.GetDataCatalogEncryptionSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetDatabaseResult struct {
	Result workflow.Future
}

func (r *GlueGetDatabaseResult) Get(ctx workflow.Context) (*glue.GetDatabaseOutput, error) {
    var output glue.GetDatabaseOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetDatabasesResult struct {
	Result workflow.Future
}

func (r *GlueGetDatabasesResult) Get(ctx workflow.Context) (*glue.GetDatabasesOutput, error) {
    var output glue.GetDatabasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetDataflowGraphResult struct {
	Result workflow.Future
}

func (r *GlueGetDataflowGraphResult) Get(ctx workflow.Context) (*glue.GetDataflowGraphOutput, error) {
    var output glue.GetDataflowGraphOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetDevEndpointResult struct {
	Result workflow.Future
}

func (r *GlueGetDevEndpointResult) Get(ctx workflow.Context) (*glue.GetDevEndpointOutput, error) {
    var output glue.GetDevEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetDevEndpointsResult struct {
	Result workflow.Future
}

func (r *GlueGetDevEndpointsResult) Get(ctx workflow.Context) (*glue.GetDevEndpointsOutput, error) {
    var output glue.GetDevEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetJobResult struct {
	Result workflow.Future
}

func (r *GlueGetJobResult) Get(ctx workflow.Context) (*glue.GetJobOutput, error) {
    var output glue.GetJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetJobBookmarkResult struct {
	Result workflow.Future
}

func (r *GlueGetJobBookmarkResult) Get(ctx workflow.Context) (*glue.GetJobBookmarkOutput, error) {
    var output glue.GetJobBookmarkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetJobRunResult struct {
	Result workflow.Future
}

func (r *GlueGetJobRunResult) Get(ctx workflow.Context) (*glue.GetJobRunOutput, error) {
    var output glue.GetJobRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetJobRunsResult struct {
	Result workflow.Future
}

func (r *GlueGetJobRunsResult) Get(ctx workflow.Context) (*glue.GetJobRunsOutput, error) {
    var output glue.GetJobRunsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetJobsResult struct {
	Result workflow.Future
}

func (r *GlueGetJobsResult) Get(ctx workflow.Context) (*glue.GetJobsOutput, error) {
    var output glue.GetJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetMLTaskRunResult struct {
	Result workflow.Future
}

func (r *GlueGetMLTaskRunResult) Get(ctx workflow.Context) (*glue.GetMLTaskRunOutput, error) {
    var output glue.GetMLTaskRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetMLTaskRunsResult struct {
	Result workflow.Future
}

func (r *GlueGetMLTaskRunsResult) Get(ctx workflow.Context) (*glue.GetMLTaskRunsOutput, error) {
    var output glue.GetMLTaskRunsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetMLTransformResult struct {
	Result workflow.Future
}

func (r *GlueGetMLTransformResult) Get(ctx workflow.Context) (*glue.GetMLTransformOutput, error) {
    var output glue.GetMLTransformOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetMLTransformsResult struct {
	Result workflow.Future
}

func (r *GlueGetMLTransformsResult) Get(ctx workflow.Context) (*glue.GetMLTransformsOutput, error) {
    var output glue.GetMLTransformsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetMappingResult struct {
	Result workflow.Future
}

func (r *GlueGetMappingResult) Get(ctx workflow.Context) (*glue.GetMappingOutput, error) {
    var output glue.GetMappingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetPartitionResult struct {
	Result workflow.Future
}

func (r *GlueGetPartitionResult) Get(ctx workflow.Context) (*glue.GetPartitionOutput, error) {
    var output glue.GetPartitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetPartitionsResult struct {
	Result workflow.Future
}

func (r *GlueGetPartitionsResult) Get(ctx workflow.Context) (*glue.GetPartitionsOutput, error) {
    var output glue.GetPartitionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetPlanResult struct {
	Result workflow.Future
}

func (r *GlueGetPlanResult) Get(ctx workflow.Context) (*glue.GetPlanOutput, error) {
    var output glue.GetPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetResourcePoliciesResult struct {
	Result workflow.Future
}

func (r *GlueGetResourcePoliciesResult) Get(ctx workflow.Context) (*glue.GetResourcePoliciesOutput, error) {
    var output glue.GetResourcePoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetResourcePolicyResult struct {
	Result workflow.Future
}

func (r *GlueGetResourcePolicyResult) Get(ctx workflow.Context) (*glue.GetResourcePolicyOutput, error) {
    var output glue.GetResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetSecurityConfigurationResult struct {
	Result workflow.Future
}

func (r *GlueGetSecurityConfigurationResult) Get(ctx workflow.Context) (*glue.GetSecurityConfigurationOutput, error) {
    var output glue.GetSecurityConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetSecurityConfigurationsResult struct {
	Result workflow.Future
}

func (r *GlueGetSecurityConfigurationsResult) Get(ctx workflow.Context) (*glue.GetSecurityConfigurationsOutput, error) {
    var output glue.GetSecurityConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetTableResult struct {
	Result workflow.Future
}

func (r *GlueGetTableResult) Get(ctx workflow.Context) (*glue.GetTableOutput, error) {
    var output glue.GetTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetTableVersionResult struct {
	Result workflow.Future
}

func (r *GlueGetTableVersionResult) Get(ctx workflow.Context) (*glue.GetTableVersionOutput, error) {
    var output glue.GetTableVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetTableVersionsResult struct {
	Result workflow.Future
}

func (r *GlueGetTableVersionsResult) Get(ctx workflow.Context) (*glue.GetTableVersionsOutput, error) {
    var output glue.GetTableVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetTablesResult struct {
	Result workflow.Future
}

func (r *GlueGetTablesResult) Get(ctx workflow.Context) (*glue.GetTablesOutput, error) {
    var output glue.GetTablesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetTagsResult struct {
	Result workflow.Future
}

func (r *GlueGetTagsResult) Get(ctx workflow.Context) (*glue.GetTagsOutput, error) {
    var output glue.GetTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetTriggerResult struct {
	Result workflow.Future
}

func (r *GlueGetTriggerResult) Get(ctx workflow.Context) (*glue.GetTriggerOutput, error) {
    var output glue.GetTriggerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetTriggersResult struct {
	Result workflow.Future
}

func (r *GlueGetTriggersResult) Get(ctx workflow.Context) (*glue.GetTriggersOutput, error) {
    var output glue.GetTriggersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetUserDefinedFunctionResult struct {
	Result workflow.Future
}

func (r *GlueGetUserDefinedFunctionResult) Get(ctx workflow.Context) (*glue.GetUserDefinedFunctionOutput, error) {
    var output glue.GetUserDefinedFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetUserDefinedFunctionsResult struct {
	Result workflow.Future
}

func (r *GlueGetUserDefinedFunctionsResult) Get(ctx workflow.Context) (*glue.GetUserDefinedFunctionsOutput, error) {
    var output glue.GetUserDefinedFunctionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetWorkflowResult struct {
	Result workflow.Future
}

func (r *GlueGetWorkflowResult) Get(ctx workflow.Context) (*glue.GetWorkflowOutput, error) {
    var output glue.GetWorkflowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetWorkflowRunResult struct {
	Result workflow.Future
}

func (r *GlueGetWorkflowRunResult) Get(ctx workflow.Context) (*glue.GetWorkflowRunOutput, error) {
    var output glue.GetWorkflowRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetWorkflowRunPropertiesResult struct {
	Result workflow.Future
}

func (r *GlueGetWorkflowRunPropertiesResult) Get(ctx workflow.Context) (*glue.GetWorkflowRunPropertiesOutput, error) {
    var output glue.GetWorkflowRunPropertiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueGetWorkflowRunsResult struct {
	Result workflow.Future
}

func (r *GlueGetWorkflowRunsResult) Get(ctx workflow.Context) (*glue.GetWorkflowRunsOutput, error) {
    var output glue.GetWorkflowRunsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueImportCatalogToGlueResult struct {
	Result workflow.Future
}

func (r *GlueImportCatalogToGlueResult) Get(ctx workflow.Context) (*glue.ImportCatalogToGlueOutput, error) {
    var output glue.ImportCatalogToGlueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueListCrawlersResult struct {
	Result workflow.Future
}

func (r *GlueListCrawlersResult) Get(ctx workflow.Context) (*glue.ListCrawlersOutput, error) {
    var output glue.ListCrawlersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueListDevEndpointsResult struct {
	Result workflow.Future
}

func (r *GlueListDevEndpointsResult) Get(ctx workflow.Context) (*glue.ListDevEndpointsOutput, error) {
    var output glue.ListDevEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueListJobsResult struct {
	Result workflow.Future
}

func (r *GlueListJobsResult) Get(ctx workflow.Context) (*glue.ListJobsOutput, error) {
    var output glue.ListJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueListMLTransformsResult struct {
	Result workflow.Future
}

func (r *GlueListMLTransformsResult) Get(ctx workflow.Context) (*glue.ListMLTransformsOutput, error) {
    var output glue.ListMLTransformsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueListTriggersResult struct {
	Result workflow.Future
}

func (r *GlueListTriggersResult) Get(ctx workflow.Context) (*glue.ListTriggersOutput, error) {
    var output glue.ListTriggersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueListWorkflowsResult struct {
	Result workflow.Future
}

func (r *GlueListWorkflowsResult) Get(ctx workflow.Context) (*glue.ListWorkflowsOutput, error) {
    var output glue.ListWorkflowsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GluePutDataCatalogEncryptionSettingsResult struct {
	Result workflow.Future
}

func (r *GluePutDataCatalogEncryptionSettingsResult) Get(ctx workflow.Context) (*glue.PutDataCatalogEncryptionSettingsOutput, error) {
    var output glue.PutDataCatalogEncryptionSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GluePutResourcePolicyResult struct {
	Result workflow.Future
}

func (r *GluePutResourcePolicyResult) Get(ctx workflow.Context) (*glue.PutResourcePolicyOutput, error) {
    var output glue.PutResourcePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GluePutWorkflowRunPropertiesResult struct {
	Result workflow.Future
}

func (r *GluePutWorkflowRunPropertiesResult) Get(ctx workflow.Context) (*glue.PutWorkflowRunPropertiesOutput, error) {
    var output glue.PutWorkflowRunPropertiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueResetJobBookmarkResult struct {
	Result workflow.Future
}

func (r *GlueResetJobBookmarkResult) Get(ctx workflow.Context) (*glue.ResetJobBookmarkOutput, error) {
    var output glue.ResetJobBookmarkOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueResumeWorkflowRunResult struct {
	Result workflow.Future
}

func (r *GlueResumeWorkflowRunResult) Get(ctx workflow.Context) (*glue.ResumeWorkflowRunOutput, error) {
    var output glue.ResumeWorkflowRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueSearchTablesResult struct {
	Result workflow.Future
}

func (r *GlueSearchTablesResult) Get(ctx workflow.Context) (*glue.SearchTablesOutput, error) {
    var output glue.SearchTablesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStartCrawlerResult struct {
	Result workflow.Future
}

func (r *GlueStartCrawlerResult) Get(ctx workflow.Context) (*glue.StartCrawlerOutput, error) {
    var output glue.StartCrawlerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStartCrawlerScheduleResult struct {
	Result workflow.Future
}

func (r *GlueStartCrawlerScheduleResult) Get(ctx workflow.Context) (*glue.StartCrawlerScheduleOutput, error) {
    var output glue.StartCrawlerScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStartExportLabelsTaskRunResult struct {
	Result workflow.Future
}

func (r *GlueStartExportLabelsTaskRunResult) Get(ctx workflow.Context) (*glue.StartExportLabelsTaskRunOutput, error) {
    var output glue.StartExportLabelsTaskRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStartImportLabelsTaskRunResult struct {
	Result workflow.Future
}

func (r *GlueStartImportLabelsTaskRunResult) Get(ctx workflow.Context) (*glue.StartImportLabelsTaskRunOutput, error) {
    var output glue.StartImportLabelsTaskRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStartJobRunResult struct {
	Result workflow.Future
}

func (r *GlueStartJobRunResult) Get(ctx workflow.Context) (*glue.StartJobRunOutput, error) {
    var output glue.StartJobRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStartMLEvaluationTaskRunResult struct {
	Result workflow.Future
}

func (r *GlueStartMLEvaluationTaskRunResult) Get(ctx workflow.Context) (*glue.StartMLEvaluationTaskRunOutput, error) {
    var output glue.StartMLEvaluationTaskRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStartMLLabelingSetGenerationTaskRunResult struct {
	Result workflow.Future
}

func (r *GlueStartMLLabelingSetGenerationTaskRunResult) Get(ctx workflow.Context) (*glue.StartMLLabelingSetGenerationTaskRunOutput, error) {
    var output glue.StartMLLabelingSetGenerationTaskRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStartTriggerResult struct {
	Result workflow.Future
}

func (r *GlueStartTriggerResult) Get(ctx workflow.Context) (*glue.StartTriggerOutput, error) {
    var output glue.StartTriggerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStartWorkflowRunResult struct {
	Result workflow.Future
}

func (r *GlueStartWorkflowRunResult) Get(ctx workflow.Context) (*glue.StartWorkflowRunOutput, error) {
    var output glue.StartWorkflowRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStopCrawlerResult struct {
	Result workflow.Future
}

func (r *GlueStopCrawlerResult) Get(ctx workflow.Context) (*glue.StopCrawlerOutput, error) {
    var output glue.StopCrawlerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStopCrawlerScheduleResult struct {
	Result workflow.Future
}

func (r *GlueStopCrawlerScheduleResult) Get(ctx workflow.Context) (*glue.StopCrawlerScheduleOutput, error) {
    var output glue.StopCrawlerScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStopTriggerResult struct {
	Result workflow.Future
}

func (r *GlueStopTriggerResult) Get(ctx workflow.Context) (*glue.StopTriggerOutput, error) {
    var output glue.StopTriggerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStopWorkflowRunResult struct {
	Result workflow.Future
}

func (r *GlueStopWorkflowRunResult) Get(ctx workflow.Context) (*glue.StopWorkflowRunOutput, error) {
    var output glue.StopWorkflowRunOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueTagResourceResult struct {
	Result workflow.Future
}

func (r *GlueTagResourceResult) Get(ctx workflow.Context) (*glue.TagResourceOutput, error) {
    var output glue.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUntagResourceResult struct {
	Result workflow.Future
}

func (r *GlueUntagResourceResult) Get(ctx workflow.Context) (*glue.UntagResourceOutput, error) {
    var output glue.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateClassifierResult struct {
	Result workflow.Future
}

func (r *GlueUpdateClassifierResult) Get(ctx workflow.Context) (*glue.UpdateClassifierOutput, error) {
    var output glue.UpdateClassifierOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateColumnStatisticsForPartitionResult struct {
	Result workflow.Future
}

func (r *GlueUpdateColumnStatisticsForPartitionResult) Get(ctx workflow.Context) (*glue.UpdateColumnStatisticsForPartitionOutput, error) {
    var output glue.UpdateColumnStatisticsForPartitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateColumnStatisticsForTableResult struct {
	Result workflow.Future
}

func (r *GlueUpdateColumnStatisticsForTableResult) Get(ctx workflow.Context) (*glue.UpdateColumnStatisticsForTableOutput, error) {
    var output glue.UpdateColumnStatisticsForTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateConnectionResult struct {
	Result workflow.Future
}

func (r *GlueUpdateConnectionResult) Get(ctx workflow.Context) (*glue.UpdateConnectionOutput, error) {
    var output glue.UpdateConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateCrawlerResult struct {
	Result workflow.Future
}

func (r *GlueUpdateCrawlerResult) Get(ctx workflow.Context) (*glue.UpdateCrawlerOutput, error) {
    var output glue.UpdateCrawlerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateCrawlerScheduleResult struct {
	Result workflow.Future
}

func (r *GlueUpdateCrawlerScheduleResult) Get(ctx workflow.Context) (*glue.UpdateCrawlerScheduleOutput, error) {
    var output glue.UpdateCrawlerScheduleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateDatabaseResult struct {
	Result workflow.Future
}

func (r *GlueUpdateDatabaseResult) Get(ctx workflow.Context) (*glue.UpdateDatabaseOutput, error) {
    var output glue.UpdateDatabaseOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateDevEndpointResult struct {
	Result workflow.Future
}

func (r *GlueUpdateDevEndpointResult) Get(ctx workflow.Context) (*glue.UpdateDevEndpointOutput, error) {
    var output glue.UpdateDevEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateJobResult struct {
	Result workflow.Future
}

func (r *GlueUpdateJobResult) Get(ctx workflow.Context) (*glue.UpdateJobOutput, error) {
    var output glue.UpdateJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateMLTransformResult struct {
	Result workflow.Future
}

func (r *GlueUpdateMLTransformResult) Get(ctx workflow.Context) (*glue.UpdateMLTransformOutput, error) {
    var output glue.UpdateMLTransformOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdatePartitionResult struct {
	Result workflow.Future
}

func (r *GlueUpdatePartitionResult) Get(ctx workflow.Context) (*glue.UpdatePartitionOutput, error) {
    var output glue.UpdatePartitionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateTableResult struct {
	Result workflow.Future
}

func (r *GlueUpdateTableResult) Get(ctx workflow.Context) (*glue.UpdateTableOutput, error) {
    var output glue.UpdateTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateTriggerResult struct {
	Result workflow.Future
}

func (r *GlueUpdateTriggerResult) Get(ctx workflow.Context) (*glue.UpdateTriggerOutput, error) {
    var output glue.UpdateTriggerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateUserDefinedFunctionResult struct {
	Result workflow.Future
}

func (r *GlueUpdateUserDefinedFunctionResult) Get(ctx workflow.Context) (*glue.UpdateUserDefinedFunctionOutput, error) {
    var output glue.UpdateUserDefinedFunctionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueUpdateWorkflowResult struct {
	Result workflow.Future
}

func (r *GlueUpdateWorkflowResult) Get(ctx workflow.Context) (*glue.UpdateWorkflowOutput, error) {
    var output glue.UpdateWorkflowOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GlueStub struct {
    activities awsactivities.GlueActivities
}

func NewGlueStub() GlueClient {
    return &GlueStub{}
}

func (a *GlueStub) BatchCreatePartition(ctx workflow.Context, input *glue.BatchCreatePartitionInput) (*glue.BatchCreatePartitionOutput, error) {
    var output glue.BatchCreatePartitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchCreatePartition, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchCreatePartitionAsync(ctx workflow.Context, input *glue.BatchCreatePartitionInput) *GlueBatchCreatePartitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchCreatePartition, input)
    return &GlueBatchCreatePartitionResult{Result: future}
}

func (a *GlueStub) BatchDeleteConnection(ctx workflow.Context, input *glue.BatchDeleteConnectionInput) (*glue.BatchDeleteConnectionOutput, error) {
    var output glue.BatchDeleteConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchDeleteConnectionAsync(ctx workflow.Context, input *glue.BatchDeleteConnectionInput) *GlueBatchDeleteConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteConnection, input)
    return &GlueBatchDeleteConnectionResult{Result: future}
}

func (a *GlueStub) BatchDeletePartition(ctx workflow.Context, input *glue.BatchDeletePartitionInput) (*glue.BatchDeletePartitionOutput, error) {
    var output glue.BatchDeletePartitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDeletePartition, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchDeletePartitionAsync(ctx workflow.Context, input *glue.BatchDeletePartitionInput) *GlueBatchDeletePartitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDeletePartition, input)
    return &GlueBatchDeletePartitionResult{Result: future}
}

func (a *GlueStub) BatchDeleteTable(ctx workflow.Context, input *glue.BatchDeleteTableInput) (*glue.BatchDeleteTableOutput, error) {
    var output glue.BatchDeleteTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteTable, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchDeleteTableAsync(ctx workflow.Context, input *glue.BatchDeleteTableInput) *GlueBatchDeleteTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteTable, input)
    return &GlueBatchDeleteTableResult{Result: future}
}

func (a *GlueStub) BatchDeleteTableVersion(ctx workflow.Context, input *glue.BatchDeleteTableVersionInput) (*glue.BatchDeleteTableVersionOutput, error) {
    var output glue.BatchDeleteTableVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteTableVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchDeleteTableVersionAsync(ctx workflow.Context, input *glue.BatchDeleteTableVersionInput) *GlueBatchDeleteTableVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteTableVersion, input)
    return &GlueBatchDeleteTableVersionResult{Result: future}
}

func (a *GlueStub) BatchGetCrawlers(ctx workflow.Context, input *glue.BatchGetCrawlersInput) (*glue.BatchGetCrawlersOutput, error) {
    var output glue.BatchGetCrawlersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetCrawlers, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchGetCrawlersAsync(ctx workflow.Context, input *glue.BatchGetCrawlersInput) *GlueBatchGetCrawlersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetCrawlers, input)
    return &GlueBatchGetCrawlersResult{Result: future}
}

func (a *GlueStub) BatchGetDevEndpoints(ctx workflow.Context, input *glue.BatchGetDevEndpointsInput) (*glue.BatchGetDevEndpointsOutput, error) {
    var output glue.BatchGetDevEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetDevEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchGetDevEndpointsAsync(ctx workflow.Context, input *glue.BatchGetDevEndpointsInput) *GlueBatchGetDevEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetDevEndpoints, input)
    return &GlueBatchGetDevEndpointsResult{Result: future}
}

func (a *GlueStub) BatchGetJobs(ctx workflow.Context, input *glue.BatchGetJobsInput) (*glue.BatchGetJobsOutput, error) {
    var output glue.BatchGetJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchGetJobsAsync(ctx workflow.Context, input *glue.BatchGetJobsInput) *GlueBatchGetJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetJobs, input)
    return &GlueBatchGetJobsResult{Result: future}
}

func (a *GlueStub) BatchGetPartition(ctx workflow.Context, input *glue.BatchGetPartitionInput) (*glue.BatchGetPartitionOutput, error) {
    var output glue.BatchGetPartitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetPartition, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchGetPartitionAsync(ctx workflow.Context, input *glue.BatchGetPartitionInput) *GlueBatchGetPartitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetPartition, input)
    return &GlueBatchGetPartitionResult{Result: future}
}

func (a *GlueStub) BatchGetTriggers(ctx workflow.Context, input *glue.BatchGetTriggersInput) (*glue.BatchGetTriggersOutput, error) {
    var output glue.BatchGetTriggersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetTriggers, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchGetTriggersAsync(ctx workflow.Context, input *glue.BatchGetTriggersInput) *GlueBatchGetTriggersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetTriggers, input)
    return &GlueBatchGetTriggersResult{Result: future}
}

func (a *GlueStub) BatchGetWorkflows(ctx workflow.Context, input *glue.BatchGetWorkflowsInput) (*glue.BatchGetWorkflowsOutput, error) {
    var output glue.BatchGetWorkflowsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchGetWorkflows, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchGetWorkflowsAsync(ctx workflow.Context, input *glue.BatchGetWorkflowsInput) *GlueBatchGetWorkflowsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchGetWorkflows, input)
    return &GlueBatchGetWorkflowsResult{Result: future}
}

func (a *GlueStub) BatchStopJobRun(ctx workflow.Context, input *glue.BatchStopJobRunInput) (*glue.BatchStopJobRunOutput, error) {
    var output glue.BatchStopJobRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchStopJobRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) BatchStopJobRunAsync(ctx workflow.Context, input *glue.BatchStopJobRunInput) *GlueBatchStopJobRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchStopJobRun, input)
    return &GlueBatchStopJobRunResult{Result: future}
}

func (a *GlueStub) CancelMLTaskRun(ctx workflow.Context, input *glue.CancelMLTaskRunInput) (*glue.CancelMLTaskRunOutput, error) {
    var output glue.CancelMLTaskRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelMLTaskRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CancelMLTaskRunAsync(ctx workflow.Context, input *glue.CancelMLTaskRunInput) *GlueCancelMLTaskRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelMLTaskRun, input)
    return &GlueCancelMLTaskRunResult{Result: future}
}

func (a *GlueStub) CreateClassifier(ctx workflow.Context, input *glue.CreateClassifierInput) (*glue.CreateClassifierOutput, error) {
    var output glue.CreateClassifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateClassifier, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateClassifierAsync(ctx workflow.Context, input *glue.CreateClassifierInput) *GlueCreateClassifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateClassifier, input)
    return &GlueCreateClassifierResult{Result: future}
}

func (a *GlueStub) CreateConnection(ctx workflow.Context, input *glue.CreateConnectionInput) (*glue.CreateConnectionOutput, error) {
    var output glue.CreateConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateConnectionAsync(ctx workflow.Context, input *glue.CreateConnectionInput) *GlueCreateConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateConnection, input)
    return &GlueCreateConnectionResult{Result: future}
}

func (a *GlueStub) CreateCrawler(ctx workflow.Context, input *glue.CreateCrawlerInput) (*glue.CreateCrawlerOutput, error) {
    var output glue.CreateCrawlerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCrawler, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateCrawlerAsync(ctx workflow.Context, input *glue.CreateCrawlerInput) *GlueCreateCrawlerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateCrawler, input)
    return &GlueCreateCrawlerResult{Result: future}
}

func (a *GlueStub) CreateDatabase(ctx workflow.Context, input *glue.CreateDatabaseInput) (*glue.CreateDatabaseOutput, error) {
    var output glue.CreateDatabaseOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDatabase, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateDatabaseAsync(ctx workflow.Context, input *glue.CreateDatabaseInput) *GlueCreateDatabaseResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDatabase, input)
    return &GlueCreateDatabaseResult{Result: future}
}

func (a *GlueStub) CreateDevEndpoint(ctx workflow.Context, input *glue.CreateDevEndpointInput) (*glue.CreateDevEndpointOutput, error) {
    var output glue.CreateDevEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDevEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateDevEndpointAsync(ctx workflow.Context, input *glue.CreateDevEndpointInput) *GlueCreateDevEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDevEndpoint, input)
    return &GlueCreateDevEndpointResult{Result: future}
}

func (a *GlueStub) CreateJob(ctx workflow.Context, input *glue.CreateJobInput) (*glue.CreateJobOutput, error) {
    var output glue.CreateJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateJobAsync(ctx workflow.Context, input *glue.CreateJobInput) *GlueCreateJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateJob, input)
    return &GlueCreateJobResult{Result: future}
}

func (a *GlueStub) CreateMLTransform(ctx workflow.Context, input *glue.CreateMLTransformInput) (*glue.CreateMLTransformOutput, error) {
    var output glue.CreateMLTransformOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMLTransform, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateMLTransformAsync(ctx workflow.Context, input *glue.CreateMLTransformInput) *GlueCreateMLTransformResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateMLTransform, input)
    return &GlueCreateMLTransformResult{Result: future}
}

func (a *GlueStub) CreatePartition(ctx workflow.Context, input *glue.CreatePartitionInput) (*glue.CreatePartitionOutput, error) {
    var output glue.CreatePartitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePartition, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreatePartitionAsync(ctx workflow.Context, input *glue.CreatePartitionInput) *GlueCreatePartitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePartition, input)
    return &GlueCreatePartitionResult{Result: future}
}

func (a *GlueStub) CreateScript(ctx workflow.Context, input *glue.CreateScriptInput) (*glue.CreateScriptOutput, error) {
    var output glue.CreateScriptOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateScript, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateScriptAsync(ctx workflow.Context, input *glue.CreateScriptInput) *GlueCreateScriptResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateScript, input)
    return &GlueCreateScriptResult{Result: future}
}

func (a *GlueStub) CreateSecurityConfiguration(ctx workflow.Context, input *glue.CreateSecurityConfigurationInput) (*glue.CreateSecurityConfigurationOutput, error) {
    var output glue.CreateSecurityConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSecurityConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateSecurityConfigurationAsync(ctx workflow.Context, input *glue.CreateSecurityConfigurationInput) *GlueCreateSecurityConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSecurityConfiguration, input)
    return &GlueCreateSecurityConfigurationResult{Result: future}
}

func (a *GlueStub) CreateTable(ctx workflow.Context, input *glue.CreateTableInput) (*glue.CreateTableOutput, error) {
    var output glue.CreateTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTable, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateTableAsync(ctx workflow.Context, input *glue.CreateTableInput) *GlueCreateTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTable, input)
    return &GlueCreateTableResult{Result: future}
}

func (a *GlueStub) CreateTrigger(ctx workflow.Context, input *glue.CreateTriggerInput) (*glue.CreateTriggerOutput, error) {
    var output glue.CreateTriggerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTrigger, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateTriggerAsync(ctx workflow.Context, input *glue.CreateTriggerInput) *GlueCreateTriggerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTrigger, input)
    return &GlueCreateTriggerResult{Result: future}
}

func (a *GlueStub) CreateUserDefinedFunction(ctx workflow.Context, input *glue.CreateUserDefinedFunctionInput) (*glue.CreateUserDefinedFunctionOutput, error) {
    var output glue.CreateUserDefinedFunctionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUserDefinedFunction, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateUserDefinedFunctionAsync(ctx workflow.Context, input *glue.CreateUserDefinedFunctionInput) *GlueCreateUserDefinedFunctionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateUserDefinedFunction, input)
    return &GlueCreateUserDefinedFunctionResult{Result: future}
}

func (a *GlueStub) CreateWorkflow(ctx workflow.Context, input *glue.CreateWorkflowInput) (*glue.CreateWorkflowOutput, error) {
    var output glue.CreateWorkflowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWorkflow, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) CreateWorkflowAsync(ctx workflow.Context, input *glue.CreateWorkflowInput) *GlueCreateWorkflowResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateWorkflow, input)
    return &GlueCreateWorkflowResult{Result: future}
}

func (a *GlueStub) DeleteClassifier(ctx workflow.Context, input *glue.DeleteClassifierInput) (*glue.DeleteClassifierOutput, error) {
    var output glue.DeleteClassifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteClassifier, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteClassifierAsync(ctx workflow.Context, input *glue.DeleteClassifierInput) *GlueDeleteClassifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteClassifier, input)
    return &GlueDeleteClassifierResult{Result: future}
}

func (a *GlueStub) DeleteColumnStatisticsForPartition(ctx workflow.Context, input *glue.DeleteColumnStatisticsForPartitionInput) (*glue.DeleteColumnStatisticsForPartitionOutput, error) {
    var output glue.DeleteColumnStatisticsForPartitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteColumnStatisticsForPartition, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.DeleteColumnStatisticsForPartitionInput) *GlueDeleteColumnStatisticsForPartitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteColumnStatisticsForPartition, input)
    return &GlueDeleteColumnStatisticsForPartitionResult{Result: future}
}

func (a *GlueStub) DeleteColumnStatisticsForTable(ctx workflow.Context, input *glue.DeleteColumnStatisticsForTableInput) (*glue.DeleteColumnStatisticsForTableOutput, error) {
    var output glue.DeleteColumnStatisticsForTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteColumnStatisticsForTable, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.DeleteColumnStatisticsForTableInput) *GlueDeleteColumnStatisticsForTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteColumnStatisticsForTable, input)
    return &GlueDeleteColumnStatisticsForTableResult{Result: future}
}

func (a *GlueStub) DeleteConnection(ctx workflow.Context, input *glue.DeleteConnectionInput) (*glue.DeleteConnectionOutput, error) {
    var output glue.DeleteConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteConnectionAsync(ctx workflow.Context, input *glue.DeleteConnectionInput) *GlueDeleteConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConnection, input)
    return &GlueDeleteConnectionResult{Result: future}
}

func (a *GlueStub) DeleteCrawler(ctx workflow.Context, input *glue.DeleteCrawlerInput) (*glue.DeleteCrawlerOutput, error) {
    var output glue.DeleteCrawlerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCrawler, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteCrawlerAsync(ctx workflow.Context, input *glue.DeleteCrawlerInput) *GlueDeleteCrawlerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteCrawler, input)
    return &GlueDeleteCrawlerResult{Result: future}
}

func (a *GlueStub) DeleteDatabase(ctx workflow.Context, input *glue.DeleteDatabaseInput) (*glue.DeleteDatabaseOutput, error) {
    var output glue.DeleteDatabaseOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDatabase, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteDatabaseAsync(ctx workflow.Context, input *glue.DeleteDatabaseInput) *GlueDeleteDatabaseResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDatabase, input)
    return &GlueDeleteDatabaseResult{Result: future}
}

func (a *GlueStub) DeleteDevEndpoint(ctx workflow.Context, input *glue.DeleteDevEndpointInput) (*glue.DeleteDevEndpointOutput, error) {
    var output glue.DeleteDevEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDevEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteDevEndpointAsync(ctx workflow.Context, input *glue.DeleteDevEndpointInput) *GlueDeleteDevEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDevEndpoint, input)
    return &GlueDeleteDevEndpointResult{Result: future}
}

func (a *GlueStub) DeleteJob(ctx workflow.Context, input *glue.DeleteJobInput) (*glue.DeleteJobOutput, error) {
    var output glue.DeleteJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteJob, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteJobAsync(ctx workflow.Context, input *glue.DeleteJobInput) *GlueDeleteJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteJob, input)
    return &GlueDeleteJobResult{Result: future}
}

func (a *GlueStub) DeleteMLTransform(ctx workflow.Context, input *glue.DeleteMLTransformInput) (*glue.DeleteMLTransformOutput, error) {
    var output glue.DeleteMLTransformOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMLTransform, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteMLTransformAsync(ctx workflow.Context, input *glue.DeleteMLTransformInput) *GlueDeleteMLTransformResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMLTransform, input)
    return &GlueDeleteMLTransformResult{Result: future}
}

func (a *GlueStub) DeletePartition(ctx workflow.Context, input *glue.DeletePartitionInput) (*glue.DeletePartitionOutput, error) {
    var output glue.DeletePartitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePartition, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeletePartitionAsync(ctx workflow.Context, input *glue.DeletePartitionInput) *GlueDeletePartitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePartition, input)
    return &GlueDeletePartitionResult{Result: future}
}

func (a *GlueStub) DeleteResourcePolicy(ctx workflow.Context, input *glue.DeleteResourcePolicyInput) (*glue.DeleteResourcePolicyOutput, error) {
    var output glue.DeleteResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteResourcePolicyAsync(ctx workflow.Context, input *glue.DeleteResourcePolicyInput) *GlueDeleteResourcePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteResourcePolicy, input)
    return &GlueDeleteResourcePolicyResult{Result: future}
}

func (a *GlueStub) DeleteSecurityConfiguration(ctx workflow.Context, input *glue.DeleteSecurityConfigurationInput) (*glue.DeleteSecurityConfigurationOutput, error) {
    var output glue.DeleteSecurityConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSecurityConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteSecurityConfigurationAsync(ctx workflow.Context, input *glue.DeleteSecurityConfigurationInput) *GlueDeleteSecurityConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteSecurityConfiguration, input)
    return &GlueDeleteSecurityConfigurationResult{Result: future}
}

func (a *GlueStub) DeleteTable(ctx workflow.Context, input *glue.DeleteTableInput) (*glue.DeleteTableOutput, error) {
    var output glue.DeleteTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTable, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteTableAsync(ctx workflow.Context, input *glue.DeleteTableInput) *GlueDeleteTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTable, input)
    return &GlueDeleteTableResult{Result: future}
}

func (a *GlueStub) DeleteTableVersion(ctx workflow.Context, input *glue.DeleteTableVersionInput) (*glue.DeleteTableVersionOutput, error) {
    var output glue.DeleteTableVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTableVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteTableVersionAsync(ctx workflow.Context, input *glue.DeleteTableVersionInput) *GlueDeleteTableVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTableVersion, input)
    return &GlueDeleteTableVersionResult{Result: future}
}

func (a *GlueStub) DeleteTrigger(ctx workflow.Context, input *glue.DeleteTriggerInput) (*glue.DeleteTriggerOutput, error) {
    var output glue.DeleteTriggerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTrigger, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteTriggerAsync(ctx workflow.Context, input *glue.DeleteTriggerInput) *GlueDeleteTriggerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTrigger, input)
    return &GlueDeleteTriggerResult{Result: future}
}

func (a *GlueStub) DeleteUserDefinedFunction(ctx workflow.Context, input *glue.DeleteUserDefinedFunctionInput) (*glue.DeleteUserDefinedFunctionOutput, error) {
    var output glue.DeleteUserDefinedFunctionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUserDefinedFunction, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteUserDefinedFunctionAsync(ctx workflow.Context, input *glue.DeleteUserDefinedFunctionInput) *GlueDeleteUserDefinedFunctionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteUserDefinedFunction, input)
    return &GlueDeleteUserDefinedFunctionResult{Result: future}
}

func (a *GlueStub) DeleteWorkflow(ctx workflow.Context, input *glue.DeleteWorkflowInput) (*glue.DeleteWorkflowOutput, error) {
    var output glue.DeleteWorkflowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteWorkflow, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) DeleteWorkflowAsync(ctx workflow.Context, input *glue.DeleteWorkflowInput) *GlueDeleteWorkflowResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteWorkflow, input)
    return &GlueDeleteWorkflowResult{Result: future}
}

func (a *GlueStub) GetCatalogImportStatus(ctx workflow.Context, input *glue.GetCatalogImportStatusInput) (*glue.GetCatalogImportStatusOutput, error) {
    var output glue.GetCatalogImportStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCatalogImportStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetCatalogImportStatusAsync(ctx workflow.Context, input *glue.GetCatalogImportStatusInput) *GlueGetCatalogImportStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCatalogImportStatus, input)
    return &GlueGetCatalogImportStatusResult{Result: future}
}

func (a *GlueStub) GetClassifier(ctx workflow.Context, input *glue.GetClassifierInput) (*glue.GetClassifierOutput, error) {
    var output glue.GetClassifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetClassifier, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetClassifierAsync(ctx workflow.Context, input *glue.GetClassifierInput) *GlueGetClassifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetClassifier, input)
    return &GlueGetClassifierResult{Result: future}
}

func (a *GlueStub) GetClassifiers(ctx workflow.Context, input *glue.GetClassifiersInput) (*glue.GetClassifiersOutput, error) {
    var output glue.GetClassifiersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetClassifiers, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetClassifiersAsync(ctx workflow.Context, input *glue.GetClassifiersInput) *GlueGetClassifiersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetClassifiers, input)
    return &GlueGetClassifiersResult{Result: future}
}

func (a *GlueStub) GetColumnStatisticsForPartition(ctx workflow.Context, input *glue.GetColumnStatisticsForPartitionInput) (*glue.GetColumnStatisticsForPartitionOutput, error) {
    var output glue.GetColumnStatisticsForPartitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetColumnStatisticsForPartition, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.GetColumnStatisticsForPartitionInput) *GlueGetColumnStatisticsForPartitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetColumnStatisticsForPartition, input)
    return &GlueGetColumnStatisticsForPartitionResult{Result: future}
}

func (a *GlueStub) GetColumnStatisticsForTable(ctx workflow.Context, input *glue.GetColumnStatisticsForTableInput) (*glue.GetColumnStatisticsForTableOutput, error) {
    var output glue.GetColumnStatisticsForTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetColumnStatisticsForTable, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.GetColumnStatisticsForTableInput) *GlueGetColumnStatisticsForTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetColumnStatisticsForTable, input)
    return &GlueGetColumnStatisticsForTableResult{Result: future}
}

func (a *GlueStub) GetConnection(ctx workflow.Context, input *glue.GetConnectionInput) (*glue.GetConnectionOutput, error) {
    var output glue.GetConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetConnectionAsync(ctx workflow.Context, input *glue.GetConnectionInput) *GlueGetConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetConnection, input)
    return &GlueGetConnectionResult{Result: future}
}

func (a *GlueStub) GetConnections(ctx workflow.Context, input *glue.GetConnectionsInput) (*glue.GetConnectionsOutput, error) {
    var output glue.GetConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetConnectionsAsync(ctx workflow.Context, input *glue.GetConnectionsInput) *GlueGetConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetConnections, input)
    return &GlueGetConnectionsResult{Result: future}
}

func (a *GlueStub) GetCrawler(ctx workflow.Context, input *glue.GetCrawlerInput) (*glue.GetCrawlerOutput, error) {
    var output glue.GetCrawlerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCrawler, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetCrawlerAsync(ctx workflow.Context, input *glue.GetCrawlerInput) *GlueGetCrawlerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCrawler, input)
    return &GlueGetCrawlerResult{Result: future}
}

func (a *GlueStub) GetCrawlerMetrics(ctx workflow.Context, input *glue.GetCrawlerMetricsInput) (*glue.GetCrawlerMetricsOutput, error) {
    var output glue.GetCrawlerMetricsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCrawlerMetrics, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetCrawlerMetricsAsync(ctx workflow.Context, input *glue.GetCrawlerMetricsInput) *GlueGetCrawlerMetricsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCrawlerMetrics, input)
    return &GlueGetCrawlerMetricsResult{Result: future}
}

func (a *GlueStub) GetCrawlers(ctx workflow.Context, input *glue.GetCrawlersInput) (*glue.GetCrawlersOutput, error) {
    var output glue.GetCrawlersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCrawlers, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetCrawlersAsync(ctx workflow.Context, input *glue.GetCrawlersInput) *GlueGetCrawlersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCrawlers, input)
    return &GlueGetCrawlersResult{Result: future}
}

func (a *GlueStub) GetDataCatalogEncryptionSettings(ctx workflow.Context, input *glue.GetDataCatalogEncryptionSettingsInput) (*glue.GetDataCatalogEncryptionSettingsOutput, error) {
    var output glue.GetDataCatalogEncryptionSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDataCatalogEncryptionSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetDataCatalogEncryptionSettingsAsync(ctx workflow.Context, input *glue.GetDataCatalogEncryptionSettingsInput) *GlueGetDataCatalogEncryptionSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDataCatalogEncryptionSettings, input)
    return &GlueGetDataCatalogEncryptionSettingsResult{Result: future}
}

func (a *GlueStub) GetDatabase(ctx workflow.Context, input *glue.GetDatabaseInput) (*glue.GetDatabaseOutput, error) {
    var output glue.GetDatabaseOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDatabase, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetDatabaseAsync(ctx workflow.Context, input *glue.GetDatabaseInput) *GlueGetDatabaseResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDatabase, input)
    return &GlueGetDatabaseResult{Result: future}
}

func (a *GlueStub) GetDatabases(ctx workflow.Context, input *glue.GetDatabasesInput) (*glue.GetDatabasesOutput, error) {
    var output glue.GetDatabasesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDatabases, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetDatabasesAsync(ctx workflow.Context, input *glue.GetDatabasesInput) *GlueGetDatabasesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDatabases, input)
    return &GlueGetDatabasesResult{Result: future}
}

func (a *GlueStub) GetDataflowGraph(ctx workflow.Context, input *glue.GetDataflowGraphInput) (*glue.GetDataflowGraphOutput, error) {
    var output glue.GetDataflowGraphOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDataflowGraph, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetDataflowGraphAsync(ctx workflow.Context, input *glue.GetDataflowGraphInput) *GlueGetDataflowGraphResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDataflowGraph, input)
    return &GlueGetDataflowGraphResult{Result: future}
}

func (a *GlueStub) GetDevEndpoint(ctx workflow.Context, input *glue.GetDevEndpointInput) (*glue.GetDevEndpointOutput, error) {
    var output glue.GetDevEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDevEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetDevEndpointAsync(ctx workflow.Context, input *glue.GetDevEndpointInput) *GlueGetDevEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDevEndpoint, input)
    return &GlueGetDevEndpointResult{Result: future}
}

func (a *GlueStub) GetDevEndpoints(ctx workflow.Context, input *glue.GetDevEndpointsInput) (*glue.GetDevEndpointsOutput, error) {
    var output glue.GetDevEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDevEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetDevEndpointsAsync(ctx workflow.Context, input *glue.GetDevEndpointsInput) *GlueGetDevEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDevEndpoints, input)
    return &GlueGetDevEndpointsResult{Result: future}
}

func (a *GlueStub) GetJob(ctx workflow.Context, input *glue.GetJobInput) (*glue.GetJobOutput, error) {
    var output glue.GetJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJob, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetJobAsync(ctx workflow.Context, input *glue.GetJobInput) *GlueGetJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJob, input)
    return &GlueGetJobResult{Result: future}
}

func (a *GlueStub) GetJobBookmark(ctx workflow.Context, input *glue.GetJobBookmarkInput) (*glue.GetJobBookmarkOutput, error) {
    var output glue.GetJobBookmarkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJobBookmark, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetJobBookmarkAsync(ctx workflow.Context, input *glue.GetJobBookmarkInput) *GlueGetJobBookmarkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJobBookmark, input)
    return &GlueGetJobBookmarkResult{Result: future}
}

func (a *GlueStub) GetJobRun(ctx workflow.Context, input *glue.GetJobRunInput) (*glue.GetJobRunOutput, error) {
    var output glue.GetJobRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJobRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetJobRunAsync(ctx workflow.Context, input *glue.GetJobRunInput) *GlueGetJobRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJobRun, input)
    return &GlueGetJobRunResult{Result: future}
}

func (a *GlueStub) GetJobRuns(ctx workflow.Context, input *glue.GetJobRunsInput) (*glue.GetJobRunsOutput, error) {
    var output glue.GetJobRunsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJobRuns, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetJobRunsAsync(ctx workflow.Context, input *glue.GetJobRunsInput) *GlueGetJobRunsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJobRuns, input)
    return &GlueGetJobRunsResult{Result: future}
}

func (a *GlueStub) GetJobs(ctx workflow.Context, input *glue.GetJobsInput) (*glue.GetJobsOutput, error) {
    var output glue.GetJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetJobsAsync(ctx workflow.Context, input *glue.GetJobsInput) *GlueGetJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJobs, input)
    return &GlueGetJobsResult{Result: future}
}

func (a *GlueStub) GetMLTaskRun(ctx workflow.Context, input *glue.GetMLTaskRunInput) (*glue.GetMLTaskRunOutput, error) {
    var output glue.GetMLTaskRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMLTaskRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetMLTaskRunAsync(ctx workflow.Context, input *glue.GetMLTaskRunInput) *GlueGetMLTaskRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMLTaskRun, input)
    return &GlueGetMLTaskRunResult{Result: future}
}

func (a *GlueStub) GetMLTaskRuns(ctx workflow.Context, input *glue.GetMLTaskRunsInput) (*glue.GetMLTaskRunsOutput, error) {
    var output glue.GetMLTaskRunsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMLTaskRuns, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetMLTaskRunsAsync(ctx workflow.Context, input *glue.GetMLTaskRunsInput) *GlueGetMLTaskRunsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMLTaskRuns, input)
    return &GlueGetMLTaskRunsResult{Result: future}
}

func (a *GlueStub) GetMLTransform(ctx workflow.Context, input *glue.GetMLTransformInput) (*glue.GetMLTransformOutput, error) {
    var output glue.GetMLTransformOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMLTransform, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetMLTransformAsync(ctx workflow.Context, input *glue.GetMLTransformInput) *GlueGetMLTransformResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMLTransform, input)
    return &GlueGetMLTransformResult{Result: future}
}

func (a *GlueStub) GetMLTransforms(ctx workflow.Context, input *glue.GetMLTransformsInput) (*glue.GetMLTransformsOutput, error) {
    var output glue.GetMLTransformsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMLTransforms, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetMLTransformsAsync(ctx workflow.Context, input *glue.GetMLTransformsInput) *GlueGetMLTransformsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMLTransforms, input)
    return &GlueGetMLTransformsResult{Result: future}
}

func (a *GlueStub) GetMapping(ctx workflow.Context, input *glue.GetMappingInput) (*glue.GetMappingOutput, error) {
    var output glue.GetMappingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMapping, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetMappingAsync(ctx workflow.Context, input *glue.GetMappingInput) *GlueGetMappingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMapping, input)
    return &GlueGetMappingResult{Result: future}
}

func (a *GlueStub) GetPartition(ctx workflow.Context, input *glue.GetPartitionInput) (*glue.GetPartitionOutput, error) {
    var output glue.GetPartitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPartition, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetPartitionAsync(ctx workflow.Context, input *glue.GetPartitionInput) *GlueGetPartitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPartition, input)
    return &GlueGetPartitionResult{Result: future}
}

func (a *GlueStub) GetPartitions(ctx workflow.Context, input *glue.GetPartitionsInput) (*glue.GetPartitionsOutput, error) {
    var output glue.GetPartitionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPartitions, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetPartitionsAsync(ctx workflow.Context, input *glue.GetPartitionsInput) *GlueGetPartitionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPartitions, input)
    return &GlueGetPartitionsResult{Result: future}
}

func (a *GlueStub) GetPlan(ctx workflow.Context, input *glue.GetPlanInput) (*glue.GetPlanOutput, error) {
    var output glue.GetPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetPlanAsync(ctx workflow.Context, input *glue.GetPlanInput) *GlueGetPlanResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetPlan, input)
    return &GlueGetPlanResult{Result: future}
}

func (a *GlueStub) GetResourcePolicies(ctx workflow.Context, input *glue.GetResourcePoliciesInput) (*glue.GetResourcePoliciesOutput, error) {
    var output glue.GetResourcePoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResourcePolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetResourcePoliciesAsync(ctx workflow.Context, input *glue.GetResourcePoliciesInput) *GlueGetResourcePoliciesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResourcePolicies, input)
    return &GlueGetResourcePoliciesResult{Result: future}
}

func (a *GlueStub) GetResourcePolicy(ctx workflow.Context, input *glue.GetResourcePolicyInput) (*glue.GetResourcePolicyOutput, error) {
    var output glue.GetResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetResourcePolicyAsync(ctx workflow.Context, input *glue.GetResourcePolicyInput) *GlueGetResourcePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResourcePolicy, input)
    return &GlueGetResourcePolicyResult{Result: future}
}

func (a *GlueStub) GetSecurityConfiguration(ctx workflow.Context, input *glue.GetSecurityConfigurationInput) (*glue.GetSecurityConfigurationOutput, error) {
    var output glue.GetSecurityConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSecurityConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetSecurityConfigurationAsync(ctx workflow.Context, input *glue.GetSecurityConfigurationInput) *GlueGetSecurityConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSecurityConfiguration, input)
    return &GlueGetSecurityConfigurationResult{Result: future}
}

func (a *GlueStub) GetSecurityConfigurations(ctx workflow.Context, input *glue.GetSecurityConfigurationsInput) (*glue.GetSecurityConfigurationsOutput, error) {
    var output glue.GetSecurityConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSecurityConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetSecurityConfigurationsAsync(ctx workflow.Context, input *glue.GetSecurityConfigurationsInput) *GlueGetSecurityConfigurationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSecurityConfigurations, input)
    return &GlueGetSecurityConfigurationsResult{Result: future}
}

func (a *GlueStub) GetTable(ctx workflow.Context, input *glue.GetTableInput) (*glue.GetTableOutput, error) {
    var output glue.GetTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTable, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetTableAsync(ctx workflow.Context, input *glue.GetTableInput) *GlueGetTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTable, input)
    return &GlueGetTableResult{Result: future}
}

func (a *GlueStub) GetTableVersion(ctx workflow.Context, input *glue.GetTableVersionInput) (*glue.GetTableVersionOutput, error) {
    var output glue.GetTableVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTableVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetTableVersionAsync(ctx workflow.Context, input *glue.GetTableVersionInput) *GlueGetTableVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTableVersion, input)
    return &GlueGetTableVersionResult{Result: future}
}

func (a *GlueStub) GetTableVersions(ctx workflow.Context, input *glue.GetTableVersionsInput) (*glue.GetTableVersionsOutput, error) {
    var output glue.GetTableVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTableVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetTableVersionsAsync(ctx workflow.Context, input *glue.GetTableVersionsInput) *GlueGetTableVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTableVersions, input)
    return &GlueGetTableVersionsResult{Result: future}
}

func (a *GlueStub) GetTables(ctx workflow.Context, input *glue.GetTablesInput) (*glue.GetTablesOutput, error) {
    var output glue.GetTablesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTables, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetTablesAsync(ctx workflow.Context, input *glue.GetTablesInput) *GlueGetTablesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTables, input)
    return &GlueGetTablesResult{Result: future}
}

func (a *GlueStub) GetTags(ctx workflow.Context, input *glue.GetTagsInput) (*glue.GetTagsOutput, error) {
    var output glue.GetTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTags, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetTagsAsync(ctx workflow.Context, input *glue.GetTagsInput) *GlueGetTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTags, input)
    return &GlueGetTagsResult{Result: future}
}

func (a *GlueStub) GetTrigger(ctx workflow.Context, input *glue.GetTriggerInput) (*glue.GetTriggerOutput, error) {
    var output glue.GetTriggerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTrigger, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetTriggerAsync(ctx workflow.Context, input *glue.GetTriggerInput) *GlueGetTriggerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTrigger, input)
    return &GlueGetTriggerResult{Result: future}
}

func (a *GlueStub) GetTriggers(ctx workflow.Context, input *glue.GetTriggersInput) (*glue.GetTriggersOutput, error) {
    var output glue.GetTriggersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTriggers, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetTriggersAsync(ctx workflow.Context, input *glue.GetTriggersInput) *GlueGetTriggersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetTriggers, input)
    return &GlueGetTriggersResult{Result: future}
}

func (a *GlueStub) GetUserDefinedFunction(ctx workflow.Context, input *glue.GetUserDefinedFunctionInput) (*glue.GetUserDefinedFunctionOutput, error) {
    var output glue.GetUserDefinedFunctionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUserDefinedFunction, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetUserDefinedFunctionAsync(ctx workflow.Context, input *glue.GetUserDefinedFunctionInput) *GlueGetUserDefinedFunctionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetUserDefinedFunction, input)
    return &GlueGetUserDefinedFunctionResult{Result: future}
}

func (a *GlueStub) GetUserDefinedFunctions(ctx workflow.Context, input *glue.GetUserDefinedFunctionsInput) (*glue.GetUserDefinedFunctionsOutput, error) {
    var output glue.GetUserDefinedFunctionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUserDefinedFunctions, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetUserDefinedFunctionsAsync(ctx workflow.Context, input *glue.GetUserDefinedFunctionsInput) *GlueGetUserDefinedFunctionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetUserDefinedFunctions, input)
    return &GlueGetUserDefinedFunctionsResult{Result: future}
}

func (a *GlueStub) GetWorkflow(ctx workflow.Context, input *glue.GetWorkflowInput) (*glue.GetWorkflowOutput, error) {
    var output glue.GetWorkflowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWorkflow, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetWorkflowAsync(ctx workflow.Context, input *glue.GetWorkflowInput) *GlueGetWorkflowResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetWorkflow, input)
    return &GlueGetWorkflowResult{Result: future}
}

func (a *GlueStub) GetWorkflowRun(ctx workflow.Context, input *glue.GetWorkflowRunInput) (*glue.GetWorkflowRunOutput, error) {
    var output glue.GetWorkflowRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWorkflowRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetWorkflowRunAsync(ctx workflow.Context, input *glue.GetWorkflowRunInput) *GlueGetWorkflowRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetWorkflowRun, input)
    return &GlueGetWorkflowRunResult{Result: future}
}

func (a *GlueStub) GetWorkflowRunProperties(ctx workflow.Context, input *glue.GetWorkflowRunPropertiesInput) (*glue.GetWorkflowRunPropertiesOutput, error) {
    var output glue.GetWorkflowRunPropertiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWorkflowRunProperties, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetWorkflowRunPropertiesAsync(ctx workflow.Context, input *glue.GetWorkflowRunPropertiesInput) *GlueGetWorkflowRunPropertiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetWorkflowRunProperties, input)
    return &GlueGetWorkflowRunPropertiesResult{Result: future}
}

func (a *GlueStub) GetWorkflowRuns(ctx workflow.Context, input *glue.GetWorkflowRunsInput) (*glue.GetWorkflowRunsOutput, error) {
    var output glue.GetWorkflowRunsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWorkflowRuns, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) GetWorkflowRunsAsync(ctx workflow.Context, input *glue.GetWorkflowRunsInput) *GlueGetWorkflowRunsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetWorkflowRuns, input)
    return &GlueGetWorkflowRunsResult{Result: future}
}

func (a *GlueStub) ImportCatalogToGlue(ctx workflow.Context, input *glue.ImportCatalogToGlueInput) (*glue.ImportCatalogToGlueOutput, error) {
    var output glue.ImportCatalogToGlueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportCatalogToGlue, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) ImportCatalogToGlueAsync(ctx workflow.Context, input *glue.ImportCatalogToGlueInput) *GlueImportCatalogToGlueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ImportCatalogToGlue, input)
    return &GlueImportCatalogToGlueResult{Result: future}
}

func (a *GlueStub) ListCrawlers(ctx workflow.Context, input *glue.ListCrawlersInput) (*glue.ListCrawlersOutput, error) {
    var output glue.ListCrawlersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCrawlers, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) ListCrawlersAsync(ctx workflow.Context, input *glue.ListCrawlersInput) *GlueListCrawlersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListCrawlers, input)
    return &GlueListCrawlersResult{Result: future}
}

func (a *GlueStub) ListDevEndpoints(ctx workflow.Context, input *glue.ListDevEndpointsInput) (*glue.ListDevEndpointsOutput, error) {
    var output glue.ListDevEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDevEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) ListDevEndpointsAsync(ctx workflow.Context, input *glue.ListDevEndpointsInput) *GlueListDevEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDevEndpoints, input)
    return &GlueListDevEndpointsResult{Result: future}
}

func (a *GlueStub) ListJobs(ctx workflow.Context, input *glue.ListJobsInput) (*glue.ListJobsOutput, error) {
    var output glue.ListJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) ListJobsAsync(ctx workflow.Context, input *glue.ListJobsInput) *GlueListJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input)
    return &GlueListJobsResult{Result: future}
}

func (a *GlueStub) ListMLTransforms(ctx workflow.Context, input *glue.ListMLTransformsInput) (*glue.ListMLTransformsOutput, error) {
    var output glue.ListMLTransformsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMLTransforms, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) ListMLTransformsAsync(ctx workflow.Context, input *glue.ListMLTransformsInput) *GlueListMLTransformsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListMLTransforms, input)
    return &GlueListMLTransformsResult{Result: future}
}

func (a *GlueStub) ListTriggers(ctx workflow.Context, input *glue.ListTriggersInput) (*glue.ListTriggersOutput, error) {
    var output glue.ListTriggersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTriggers, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) ListTriggersAsync(ctx workflow.Context, input *glue.ListTriggersInput) *GlueListTriggersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTriggers, input)
    return &GlueListTriggersResult{Result: future}
}

func (a *GlueStub) ListWorkflows(ctx workflow.Context, input *glue.ListWorkflowsInput) (*glue.ListWorkflowsOutput, error) {
    var output glue.ListWorkflowsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWorkflows, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) ListWorkflowsAsync(ctx workflow.Context, input *glue.ListWorkflowsInput) *GlueListWorkflowsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListWorkflows, input)
    return &GlueListWorkflowsResult{Result: future}
}

func (a *GlueStub) PutDataCatalogEncryptionSettings(ctx workflow.Context, input *glue.PutDataCatalogEncryptionSettingsInput) (*glue.PutDataCatalogEncryptionSettingsOutput, error) {
    var output glue.PutDataCatalogEncryptionSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutDataCatalogEncryptionSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) PutDataCatalogEncryptionSettingsAsync(ctx workflow.Context, input *glue.PutDataCatalogEncryptionSettingsInput) *GluePutDataCatalogEncryptionSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutDataCatalogEncryptionSettings, input)
    return &GluePutDataCatalogEncryptionSettingsResult{Result: future}
}

func (a *GlueStub) PutResourcePolicy(ctx workflow.Context, input *glue.PutResourcePolicyInput) (*glue.PutResourcePolicyOutput, error) {
    var output glue.PutResourcePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutResourcePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) PutResourcePolicyAsync(ctx workflow.Context, input *glue.PutResourcePolicyInput) *GluePutResourcePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutResourcePolicy, input)
    return &GluePutResourcePolicyResult{Result: future}
}

func (a *GlueStub) PutWorkflowRunProperties(ctx workflow.Context, input *glue.PutWorkflowRunPropertiesInput) (*glue.PutWorkflowRunPropertiesOutput, error) {
    var output glue.PutWorkflowRunPropertiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutWorkflowRunProperties, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) PutWorkflowRunPropertiesAsync(ctx workflow.Context, input *glue.PutWorkflowRunPropertiesInput) *GluePutWorkflowRunPropertiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutWorkflowRunProperties, input)
    return &GluePutWorkflowRunPropertiesResult{Result: future}
}

func (a *GlueStub) ResetJobBookmark(ctx workflow.Context, input *glue.ResetJobBookmarkInput) (*glue.ResetJobBookmarkOutput, error) {
    var output glue.ResetJobBookmarkOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResetJobBookmark, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) ResetJobBookmarkAsync(ctx workflow.Context, input *glue.ResetJobBookmarkInput) *GlueResetJobBookmarkResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResetJobBookmark, input)
    return &GlueResetJobBookmarkResult{Result: future}
}

func (a *GlueStub) ResumeWorkflowRun(ctx workflow.Context, input *glue.ResumeWorkflowRunInput) (*glue.ResumeWorkflowRunOutput, error) {
    var output glue.ResumeWorkflowRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResumeWorkflowRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) ResumeWorkflowRunAsync(ctx workflow.Context, input *glue.ResumeWorkflowRunInput) *GlueResumeWorkflowRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResumeWorkflowRun, input)
    return &GlueResumeWorkflowRunResult{Result: future}
}

func (a *GlueStub) SearchTables(ctx workflow.Context, input *glue.SearchTablesInput) (*glue.SearchTablesOutput, error) {
    var output glue.SearchTablesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchTables, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) SearchTablesAsync(ctx workflow.Context, input *glue.SearchTablesInput) *GlueSearchTablesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SearchTables, input)
    return &GlueSearchTablesResult{Result: future}
}

func (a *GlueStub) StartCrawler(ctx workflow.Context, input *glue.StartCrawlerInput) (*glue.StartCrawlerOutput, error) {
    var output glue.StartCrawlerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartCrawler, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StartCrawlerAsync(ctx workflow.Context, input *glue.StartCrawlerInput) *GlueStartCrawlerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartCrawler, input)
    return &GlueStartCrawlerResult{Result: future}
}

func (a *GlueStub) StartCrawlerSchedule(ctx workflow.Context, input *glue.StartCrawlerScheduleInput) (*glue.StartCrawlerScheduleOutput, error) {
    var output glue.StartCrawlerScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartCrawlerSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StartCrawlerScheduleAsync(ctx workflow.Context, input *glue.StartCrawlerScheduleInput) *GlueStartCrawlerScheduleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartCrawlerSchedule, input)
    return &GlueStartCrawlerScheduleResult{Result: future}
}

func (a *GlueStub) StartExportLabelsTaskRun(ctx workflow.Context, input *glue.StartExportLabelsTaskRunInput) (*glue.StartExportLabelsTaskRunOutput, error) {
    var output glue.StartExportLabelsTaskRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartExportLabelsTaskRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StartExportLabelsTaskRunAsync(ctx workflow.Context, input *glue.StartExportLabelsTaskRunInput) *GlueStartExportLabelsTaskRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartExportLabelsTaskRun, input)
    return &GlueStartExportLabelsTaskRunResult{Result: future}
}

func (a *GlueStub) StartImportLabelsTaskRun(ctx workflow.Context, input *glue.StartImportLabelsTaskRunInput) (*glue.StartImportLabelsTaskRunOutput, error) {
    var output glue.StartImportLabelsTaskRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartImportLabelsTaskRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StartImportLabelsTaskRunAsync(ctx workflow.Context, input *glue.StartImportLabelsTaskRunInput) *GlueStartImportLabelsTaskRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartImportLabelsTaskRun, input)
    return &GlueStartImportLabelsTaskRunResult{Result: future}
}

func (a *GlueStub) StartJobRun(ctx workflow.Context, input *glue.StartJobRunInput) (*glue.StartJobRunOutput, error) {
    var output glue.StartJobRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartJobRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StartJobRunAsync(ctx workflow.Context, input *glue.StartJobRunInput) *GlueStartJobRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartJobRun, input)
    return &GlueStartJobRunResult{Result: future}
}

func (a *GlueStub) StartMLEvaluationTaskRun(ctx workflow.Context, input *glue.StartMLEvaluationTaskRunInput) (*glue.StartMLEvaluationTaskRunOutput, error) {
    var output glue.StartMLEvaluationTaskRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartMLEvaluationTaskRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StartMLEvaluationTaskRunAsync(ctx workflow.Context, input *glue.StartMLEvaluationTaskRunInput) *GlueStartMLEvaluationTaskRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartMLEvaluationTaskRun, input)
    return &GlueStartMLEvaluationTaskRunResult{Result: future}
}

func (a *GlueStub) StartMLLabelingSetGenerationTaskRun(ctx workflow.Context, input *glue.StartMLLabelingSetGenerationTaskRunInput) (*glue.StartMLLabelingSetGenerationTaskRunOutput, error) {
    var output glue.StartMLLabelingSetGenerationTaskRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartMLLabelingSetGenerationTaskRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StartMLLabelingSetGenerationTaskRunAsync(ctx workflow.Context, input *glue.StartMLLabelingSetGenerationTaskRunInput) *GlueStartMLLabelingSetGenerationTaskRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartMLLabelingSetGenerationTaskRun, input)
    return &GlueStartMLLabelingSetGenerationTaskRunResult{Result: future}
}

func (a *GlueStub) StartTrigger(ctx workflow.Context, input *glue.StartTriggerInput) (*glue.StartTriggerOutput, error) {
    var output glue.StartTriggerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartTrigger, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StartTriggerAsync(ctx workflow.Context, input *glue.StartTriggerInput) *GlueStartTriggerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartTrigger, input)
    return &GlueStartTriggerResult{Result: future}
}

func (a *GlueStub) StartWorkflowRun(ctx workflow.Context, input *glue.StartWorkflowRunInput) (*glue.StartWorkflowRunOutput, error) {
    var output glue.StartWorkflowRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartWorkflowRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StartWorkflowRunAsync(ctx workflow.Context, input *glue.StartWorkflowRunInput) *GlueStartWorkflowRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartWorkflowRun, input)
    return &GlueStartWorkflowRunResult{Result: future}
}

func (a *GlueStub) StopCrawler(ctx workflow.Context, input *glue.StopCrawlerInput) (*glue.StopCrawlerOutput, error) {
    var output glue.StopCrawlerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopCrawler, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StopCrawlerAsync(ctx workflow.Context, input *glue.StopCrawlerInput) *GlueStopCrawlerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopCrawler, input)
    return &GlueStopCrawlerResult{Result: future}
}

func (a *GlueStub) StopCrawlerSchedule(ctx workflow.Context, input *glue.StopCrawlerScheduleInput) (*glue.StopCrawlerScheduleOutput, error) {
    var output glue.StopCrawlerScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopCrawlerSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StopCrawlerScheduleAsync(ctx workflow.Context, input *glue.StopCrawlerScheduleInput) *GlueStopCrawlerScheduleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopCrawlerSchedule, input)
    return &GlueStopCrawlerScheduleResult{Result: future}
}

func (a *GlueStub) StopTrigger(ctx workflow.Context, input *glue.StopTriggerInput) (*glue.StopTriggerOutput, error) {
    var output glue.StopTriggerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopTrigger, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StopTriggerAsync(ctx workflow.Context, input *glue.StopTriggerInput) *GlueStopTriggerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopTrigger, input)
    return &GlueStopTriggerResult{Result: future}
}

func (a *GlueStub) StopWorkflowRun(ctx workflow.Context, input *glue.StopWorkflowRunInput) (*glue.StopWorkflowRunOutput, error) {
    var output glue.StopWorkflowRunOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopWorkflowRun, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) StopWorkflowRunAsync(ctx workflow.Context, input *glue.StopWorkflowRunInput) *GlueStopWorkflowRunResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopWorkflowRun, input)
    return &GlueStopWorkflowRunResult{Result: future}
}

func (a *GlueStub) TagResource(ctx workflow.Context, input *glue.TagResourceInput) (*glue.TagResourceOutput, error) {
    var output glue.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) TagResourceAsync(ctx workflow.Context, input *glue.TagResourceInput) *GlueTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &GlueTagResourceResult{Result: future}
}

func (a *GlueStub) UntagResource(ctx workflow.Context, input *glue.UntagResourceInput) (*glue.UntagResourceOutput, error) {
    var output glue.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UntagResourceAsync(ctx workflow.Context, input *glue.UntagResourceInput) *GlueUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &GlueUntagResourceResult{Result: future}
}

func (a *GlueStub) UpdateClassifier(ctx workflow.Context, input *glue.UpdateClassifierInput) (*glue.UpdateClassifierOutput, error) {
    var output glue.UpdateClassifierOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateClassifier, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateClassifierAsync(ctx workflow.Context, input *glue.UpdateClassifierInput) *GlueUpdateClassifierResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateClassifier, input)
    return &GlueUpdateClassifierResult{Result: future}
}

func (a *GlueStub) UpdateColumnStatisticsForPartition(ctx workflow.Context, input *glue.UpdateColumnStatisticsForPartitionInput) (*glue.UpdateColumnStatisticsForPartitionOutput, error) {
    var output glue.UpdateColumnStatisticsForPartitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateColumnStatisticsForPartition, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateColumnStatisticsForPartitionAsync(ctx workflow.Context, input *glue.UpdateColumnStatisticsForPartitionInput) *GlueUpdateColumnStatisticsForPartitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateColumnStatisticsForPartition, input)
    return &GlueUpdateColumnStatisticsForPartitionResult{Result: future}
}

func (a *GlueStub) UpdateColumnStatisticsForTable(ctx workflow.Context, input *glue.UpdateColumnStatisticsForTableInput) (*glue.UpdateColumnStatisticsForTableOutput, error) {
    var output glue.UpdateColumnStatisticsForTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateColumnStatisticsForTable, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateColumnStatisticsForTableAsync(ctx workflow.Context, input *glue.UpdateColumnStatisticsForTableInput) *GlueUpdateColumnStatisticsForTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateColumnStatisticsForTable, input)
    return &GlueUpdateColumnStatisticsForTableResult{Result: future}
}

func (a *GlueStub) UpdateConnection(ctx workflow.Context, input *glue.UpdateConnectionInput) (*glue.UpdateConnectionOutput, error) {
    var output glue.UpdateConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateConnectionAsync(ctx workflow.Context, input *glue.UpdateConnectionInput) *GlueUpdateConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateConnection, input)
    return &GlueUpdateConnectionResult{Result: future}
}

func (a *GlueStub) UpdateCrawler(ctx workflow.Context, input *glue.UpdateCrawlerInput) (*glue.UpdateCrawlerOutput, error) {
    var output glue.UpdateCrawlerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCrawler, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateCrawlerAsync(ctx workflow.Context, input *glue.UpdateCrawlerInput) *GlueUpdateCrawlerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateCrawler, input)
    return &GlueUpdateCrawlerResult{Result: future}
}

func (a *GlueStub) UpdateCrawlerSchedule(ctx workflow.Context, input *glue.UpdateCrawlerScheduleInput) (*glue.UpdateCrawlerScheduleOutput, error) {
    var output glue.UpdateCrawlerScheduleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCrawlerSchedule, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateCrawlerScheduleAsync(ctx workflow.Context, input *glue.UpdateCrawlerScheduleInput) *GlueUpdateCrawlerScheduleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateCrawlerSchedule, input)
    return &GlueUpdateCrawlerScheduleResult{Result: future}
}

func (a *GlueStub) UpdateDatabase(ctx workflow.Context, input *glue.UpdateDatabaseInput) (*glue.UpdateDatabaseOutput, error) {
    var output glue.UpdateDatabaseOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDatabase, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateDatabaseAsync(ctx workflow.Context, input *glue.UpdateDatabaseInput) *GlueUpdateDatabaseResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDatabase, input)
    return &GlueUpdateDatabaseResult{Result: future}
}

func (a *GlueStub) UpdateDevEndpoint(ctx workflow.Context, input *glue.UpdateDevEndpointInput) (*glue.UpdateDevEndpointOutput, error) {
    var output glue.UpdateDevEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDevEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateDevEndpointAsync(ctx workflow.Context, input *glue.UpdateDevEndpointInput) *GlueUpdateDevEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDevEndpoint, input)
    return &GlueUpdateDevEndpointResult{Result: future}
}

func (a *GlueStub) UpdateJob(ctx workflow.Context, input *glue.UpdateJobInput) (*glue.UpdateJobOutput, error) {
    var output glue.UpdateJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateJob, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateJobAsync(ctx workflow.Context, input *glue.UpdateJobInput) *GlueUpdateJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateJob, input)
    return &GlueUpdateJobResult{Result: future}
}

func (a *GlueStub) UpdateMLTransform(ctx workflow.Context, input *glue.UpdateMLTransformInput) (*glue.UpdateMLTransformOutput, error) {
    var output glue.UpdateMLTransformOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMLTransform, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateMLTransformAsync(ctx workflow.Context, input *glue.UpdateMLTransformInput) *GlueUpdateMLTransformResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateMLTransform, input)
    return &GlueUpdateMLTransformResult{Result: future}
}

func (a *GlueStub) UpdatePartition(ctx workflow.Context, input *glue.UpdatePartitionInput) (*glue.UpdatePartitionOutput, error) {
    var output glue.UpdatePartitionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePartition, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdatePartitionAsync(ctx workflow.Context, input *glue.UpdatePartitionInput) *GlueUpdatePartitionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdatePartition, input)
    return &GlueUpdatePartitionResult{Result: future}
}

func (a *GlueStub) UpdateTable(ctx workflow.Context, input *glue.UpdateTableInput) (*glue.UpdateTableOutput, error) {
    var output glue.UpdateTableOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTable, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateTableAsync(ctx workflow.Context, input *glue.UpdateTableInput) *GlueUpdateTableResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTable, input)
    return &GlueUpdateTableResult{Result: future}
}

func (a *GlueStub) UpdateTrigger(ctx workflow.Context, input *glue.UpdateTriggerInput) (*glue.UpdateTriggerOutput, error) {
    var output glue.UpdateTriggerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTrigger, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateTriggerAsync(ctx workflow.Context, input *glue.UpdateTriggerInput) *GlueUpdateTriggerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTrigger, input)
    return &GlueUpdateTriggerResult{Result: future}
}

func (a *GlueStub) UpdateUserDefinedFunction(ctx workflow.Context, input *glue.UpdateUserDefinedFunctionInput) (*glue.UpdateUserDefinedFunctionOutput, error) {
    var output glue.UpdateUserDefinedFunctionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserDefinedFunction, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateUserDefinedFunctionAsync(ctx workflow.Context, input *glue.UpdateUserDefinedFunctionInput) *GlueUpdateUserDefinedFunctionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateUserDefinedFunction, input)
    return &GlueUpdateUserDefinedFunctionResult{Result: future}
}

func (a *GlueStub) UpdateWorkflow(ctx workflow.Context, input *glue.UpdateWorkflowInput) (*glue.UpdateWorkflowOutput, error) {
    var output glue.UpdateWorkflowOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateWorkflow, input).Get(ctx, &output)
    return &output, err
}

func (a *GlueStub) UpdateWorkflowAsync(ctx workflow.Context, input *glue.UpdateWorkflowInput) *GlueUpdateWorkflowResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateWorkflow, input)
    return &GlueUpdateWorkflowResult{Result: future}
}
