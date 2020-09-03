package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/rds"
	"go.temporal.io/sdk/workflow"
)

type RDSClient interface {
    AddRoleToDBCluster(ctx workflow.Context, input *rds.AddRoleToDBClusterInput) (*rds.AddRoleToDBClusterOutput, error)
    AddRoleToDBClusterAsync(ctx workflow.Context, input *rds.AddRoleToDBClusterInput) *RdsAddRoleToDBClusterResult

    AddRoleToDBInstance(ctx workflow.Context, input *rds.AddRoleToDBInstanceInput) (*rds.AddRoleToDBInstanceOutput, error)
    AddRoleToDBInstanceAsync(ctx workflow.Context, input *rds.AddRoleToDBInstanceInput) *RdsAddRoleToDBInstanceResult

    AddSourceIdentifierToSubscription(ctx workflow.Context, input *rds.AddSourceIdentifierToSubscriptionInput) (*rds.AddSourceIdentifierToSubscriptionOutput, error)
    AddSourceIdentifierToSubscriptionAsync(ctx workflow.Context, input *rds.AddSourceIdentifierToSubscriptionInput) *RdsAddSourceIdentifierToSubscriptionResult

    AddTagsToResource(ctx workflow.Context, input *rds.AddTagsToResourceInput) (*rds.AddTagsToResourceOutput, error)
    AddTagsToResourceAsync(ctx workflow.Context, input *rds.AddTagsToResourceInput) *RdsAddTagsToResourceResult

    ApplyPendingMaintenanceAction(ctx workflow.Context, input *rds.ApplyPendingMaintenanceActionInput) (*rds.ApplyPendingMaintenanceActionOutput, error)
    ApplyPendingMaintenanceActionAsync(ctx workflow.Context, input *rds.ApplyPendingMaintenanceActionInput) *RdsApplyPendingMaintenanceActionResult

    AuthorizeDBSecurityGroupIngress(ctx workflow.Context, input *rds.AuthorizeDBSecurityGroupIngressInput) (*rds.AuthorizeDBSecurityGroupIngressOutput, error)
    AuthorizeDBSecurityGroupIngressAsync(ctx workflow.Context, input *rds.AuthorizeDBSecurityGroupIngressInput) *RdsAuthorizeDBSecurityGroupIngressResult

    BacktrackDBCluster(ctx workflow.Context, input *rds.BacktrackDBClusterInput) (*rds.BacktrackDBClusterOutput, error)
    BacktrackDBClusterAsync(ctx workflow.Context, input *rds.BacktrackDBClusterInput) *RdsBacktrackDBClusterResult

    CancelExportTask(ctx workflow.Context, input *rds.CancelExportTaskInput) (*rds.CancelExportTaskOutput, error)
    CancelExportTaskAsync(ctx workflow.Context, input *rds.CancelExportTaskInput) *RdsCancelExportTaskResult

    CopyDBClusterParameterGroup(ctx workflow.Context, input *rds.CopyDBClusterParameterGroupInput) (*rds.CopyDBClusterParameterGroupOutput, error)
    CopyDBClusterParameterGroupAsync(ctx workflow.Context, input *rds.CopyDBClusterParameterGroupInput) *RdsCopyDBClusterParameterGroupResult

    CopyDBClusterSnapshot(ctx workflow.Context, input *rds.CopyDBClusterSnapshotInput) (*rds.CopyDBClusterSnapshotOutput, error)
    CopyDBClusterSnapshotAsync(ctx workflow.Context, input *rds.CopyDBClusterSnapshotInput) *RdsCopyDBClusterSnapshotResult

    CopyDBParameterGroup(ctx workflow.Context, input *rds.CopyDBParameterGroupInput) (*rds.CopyDBParameterGroupOutput, error)
    CopyDBParameterGroupAsync(ctx workflow.Context, input *rds.CopyDBParameterGroupInput) *RdsCopyDBParameterGroupResult

    CopyDBSnapshot(ctx workflow.Context, input *rds.CopyDBSnapshotInput) (*rds.CopyDBSnapshotOutput, error)
    CopyDBSnapshotAsync(ctx workflow.Context, input *rds.CopyDBSnapshotInput) *RdsCopyDBSnapshotResult

    CopyOptionGroup(ctx workflow.Context, input *rds.CopyOptionGroupInput) (*rds.CopyOptionGroupOutput, error)
    CopyOptionGroupAsync(ctx workflow.Context, input *rds.CopyOptionGroupInput) *RdsCopyOptionGroupResult

    CreateCustomAvailabilityZone(ctx workflow.Context, input *rds.CreateCustomAvailabilityZoneInput) (*rds.CreateCustomAvailabilityZoneOutput, error)
    CreateCustomAvailabilityZoneAsync(ctx workflow.Context, input *rds.CreateCustomAvailabilityZoneInput) *RdsCreateCustomAvailabilityZoneResult

    CreateDBCluster(ctx workflow.Context, input *rds.CreateDBClusterInput) (*rds.CreateDBClusterOutput, error)
    CreateDBClusterAsync(ctx workflow.Context, input *rds.CreateDBClusterInput) *RdsCreateDBClusterResult

    CreateDBClusterEndpoint(ctx workflow.Context, input *rds.CreateDBClusterEndpointInput) (*rds.CreateDBClusterEndpointOutput, error)
    CreateDBClusterEndpointAsync(ctx workflow.Context, input *rds.CreateDBClusterEndpointInput) *RdsCreateDBClusterEndpointResult

    CreateDBClusterParameterGroup(ctx workflow.Context, input *rds.CreateDBClusterParameterGroupInput) (*rds.CreateDBClusterParameterGroupOutput, error)
    CreateDBClusterParameterGroupAsync(ctx workflow.Context, input *rds.CreateDBClusterParameterGroupInput) *RdsCreateDBClusterParameterGroupResult

    CreateDBClusterSnapshot(ctx workflow.Context, input *rds.CreateDBClusterSnapshotInput) (*rds.CreateDBClusterSnapshotOutput, error)
    CreateDBClusterSnapshotAsync(ctx workflow.Context, input *rds.CreateDBClusterSnapshotInput) *RdsCreateDBClusterSnapshotResult

    CreateDBInstance(ctx workflow.Context, input *rds.CreateDBInstanceInput) (*rds.CreateDBInstanceOutput, error)
    CreateDBInstanceAsync(ctx workflow.Context, input *rds.CreateDBInstanceInput) *RdsCreateDBInstanceResult

    CreateDBInstanceReadReplica(ctx workflow.Context, input *rds.CreateDBInstanceReadReplicaInput) (*rds.CreateDBInstanceReadReplicaOutput, error)
    CreateDBInstanceReadReplicaAsync(ctx workflow.Context, input *rds.CreateDBInstanceReadReplicaInput) *RdsCreateDBInstanceReadReplicaResult

    CreateDBParameterGroup(ctx workflow.Context, input *rds.CreateDBParameterGroupInput) (*rds.CreateDBParameterGroupOutput, error)
    CreateDBParameterGroupAsync(ctx workflow.Context, input *rds.CreateDBParameterGroupInput) *RdsCreateDBParameterGroupResult

    CreateDBProxy(ctx workflow.Context, input *rds.CreateDBProxyInput) (*rds.CreateDBProxyOutput, error)
    CreateDBProxyAsync(ctx workflow.Context, input *rds.CreateDBProxyInput) *RdsCreateDBProxyResult

    CreateDBSecurityGroup(ctx workflow.Context, input *rds.CreateDBSecurityGroupInput) (*rds.CreateDBSecurityGroupOutput, error)
    CreateDBSecurityGroupAsync(ctx workflow.Context, input *rds.CreateDBSecurityGroupInput) *RdsCreateDBSecurityGroupResult

    CreateDBSnapshot(ctx workflow.Context, input *rds.CreateDBSnapshotInput) (*rds.CreateDBSnapshotOutput, error)
    CreateDBSnapshotAsync(ctx workflow.Context, input *rds.CreateDBSnapshotInput) *RdsCreateDBSnapshotResult

    CreateDBSubnetGroup(ctx workflow.Context, input *rds.CreateDBSubnetGroupInput) (*rds.CreateDBSubnetGroupOutput, error)
    CreateDBSubnetGroupAsync(ctx workflow.Context, input *rds.CreateDBSubnetGroupInput) *RdsCreateDBSubnetGroupResult

    CreateEventSubscription(ctx workflow.Context, input *rds.CreateEventSubscriptionInput) (*rds.CreateEventSubscriptionOutput, error)
    CreateEventSubscriptionAsync(ctx workflow.Context, input *rds.CreateEventSubscriptionInput) *RdsCreateEventSubscriptionResult

    CreateGlobalCluster(ctx workflow.Context, input *rds.CreateGlobalClusterInput) (*rds.CreateGlobalClusterOutput, error)
    CreateGlobalClusterAsync(ctx workflow.Context, input *rds.CreateGlobalClusterInput) *RdsCreateGlobalClusterResult

    CreateOptionGroup(ctx workflow.Context, input *rds.CreateOptionGroupInput) (*rds.CreateOptionGroupOutput, error)
    CreateOptionGroupAsync(ctx workflow.Context, input *rds.CreateOptionGroupInput) *RdsCreateOptionGroupResult

    DeleteCustomAvailabilityZone(ctx workflow.Context, input *rds.DeleteCustomAvailabilityZoneInput) (*rds.DeleteCustomAvailabilityZoneOutput, error)
    DeleteCustomAvailabilityZoneAsync(ctx workflow.Context, input *rds.DeleteCustomAvailabilityZoneInput) *RdsDeleteCustomAvailabilityZoneResult

    DeleteDBCluster(ctx workflow.Context, input *rds.DeleteDBClusterInput) (*rds.DeleteDBClusterOutput, error)
    DeleteDBClusterAsync(ctx workflow.Context, input *rds.DeleteDBClusterInput) *RdsDeleteDBClusterResult

    DeleteDBClusterEndpoint(ctx workflow.Context, input *rds.DeleteDBClusterEndpointInput) (*rds.DeleteDBClusterEndpointOutput, error)
    DeleteDBClusterEndpointAsync(ctx workflow.Context, input *rds.DeleteDBClusterEndpointInput) *RdsDeleteDBClusterEndpointResult

    DeleteDBClusterParameterGroup(ctx workflow.Context, input *rds.DeleteDBClusterParameterGroupInput) (*rds.DeleteDBClusterParameterGroupOutput, error)
    DeleteDBClusterParameterGroupAsync(ctx workflow.Context, input *rds.DeleteDBClusterParameterGroupInput) *RdsDeleteDBClusterParameterGroupResult

    DeleteDBClusterSnapshot(ctx workflow.Context, input *rds.DeleteDBClusterSnapshotInput) (*rds.DeleteDBClusterSnapshotOutput, error)
    DeleteDBClusterSnapshotAsync(ctx workflow.Context, input *rds.DeleteDBClusterSnapshotInput) *RdsDeleteDBClusterSnapshotResult

    DeleteDBInstance(ctx workflow.Context, input *rds.DeleteDBInstanceInput) (*rds.DeleteDBInstanceOutput, error)
    DeleteDBInstanceAsync(ctx workflow.Context, input *rds.DeleteDBInstanceInput) *RdsDeleteDBInstanceResult

    DeleteDBInstanceAutomatedBackup(ctx workflow.Context, input *rds.DeleteDBInstanceAutomatedBackupInput) (*rds.DeleteDBInstanceAutomatedBackupOutput, error)
    DeleteDBInstanceAutomatedBackupAsync(ctx workflow.Context, input *rds.DeleteDBInstanceAutomatedBackupInput) *RdsDeleteDBInstanceAutomatedBackupResult

    DeleteDBParameterGroup(ctx workflow.Context, input *rds.DeleteDBParameterGroupInput) (*rds.DeleteDBParameterGroupOutput, error)
    DeleteDBParameterGroupAsync(ctx workflow.Context, input *rds.DeleteDBParameterGroupInput) *RdsDeleteDBParameterGroupResult

    DeleteDBProxy(ctx workflow.Context, input *rds.DeleteDBProxyInput) (*rds.DeleteDBProxyOutput, error)
    DeleteDBProxyAsync(ctx workflow.Context, input *rds.DeleteDBProxyInput) *RdsDeleteDBProxyResult

    DeleteDBSecurityGroup(ctx workflow.Context, input *rds.DeleteDBSecurityGroupInput) (*rds.DeleteDBSecurityGroupOutput, error)
    DeleteDBSecurityGroupAsync(ctx workflow.Context, input *rds.DeleteDBSecurityGroupInput) *RdsDeleteDBSecurityGroupResult

    DeleteDBSnapshot(ctx workflow.Context, input *rds.DeleteDBSnapshotInput) (*rds.DeleteDBSnapshotOutput, error)
    DeleteDBSnapshotAsync(ctx workflow.Context, input *rds.DeleteDBSnapshotInput) *RdsDeleteDBSnapshotResult

    DeleteDBSubnetGroup(ctx workflow.Context, input *rds.DeleteDBSubnetGroupInput) (*rds.DeleteDBSubnetGroupOutput, error)
    DeleteDBSubnetGroupAsync(ctx workflow.Context, input *rds.DeleteDBSubnetGroupInput) *RdsDeleteDBSubnetGroupResult

    DeleteEventSubscription(ctx workflow.Context, input *rds.DeleteEventSubscriptionInput) (*rds.DeleteEventSubscriptionOutput, error)
    DeleteEventSubscriptionAsync(ctx workflow.Context, input *rds.DeleteEventSubscriptionInput) *RdsDeleteEventSubscriptionResult

    DeleteGlobalCluster(ctx workflow.Context, input *rds.DeleteGlobalClusterInput) (*rds.DeleteGlobalClusterOutput, error)
    DeleteGlobalClusterAsync(ctx workflow.Context, input *rds.DeleteGlobalClusterInput) *RdsDeleteGlobalClusterResult

    DeleteInstallationMedia(ctx workflow.Context, input *rds.DeleteInstallationMediaInput) (*rds.DeleteInstallationMediaOutput, error)
    DeleteInstallationMediaAsync(ctx workflow.Context, input *rds.DeleteInstallationMediaInput) *RdsDeleteInstallationMediaResult

    DeleteOptionGroup(ctx workflow.Context, input *rds.DeleteOptionGroupInput) (*rds.DeleteOptionGroupOutput, error)
    DeleteOptionGroupAsync(ctx workflow.Context, input *rds.DeleteOptionGroupInput) *RdsDeleteOptionGroupResult

    DeregisterDBProxyTargets(ctx workflow.Context, input *rds.DeregisterDBProxyTargetsInput) (*rds.DeregisterDBProxyTargetsOutput, error)
    DeregisterDBProxyTargetsAsync(ctx workflow.Context, input *rds.DeregisterDBProxyTargetsInput) *RdsDeregisterDBProxyTargetsResult

    DescribeAccountAttributes(ctx workflow.Context, input *rds.DescribeAccountAttributesInput) (*rds.DescribeAccountAttributesOutput, error)
    DescribeAccountAttributesAsync(ctx workflow.Context, input *rds.DescribeAccountAttributesInput) *RdsDescribeAccountAttributesResult

    DescribeCertificates(ctx workflow.Context, input *rds.DescribeCertificatesInput) (*rds.DescribeCertificatesOutput, error)
    DescribeCertificatesAsync(ctx workflow.Context, input *rds.DescribeCertificatesInput) *RdsDescribeCertificatesResult

    DescribeCustomAvailabilityZones(ctx workflow.Context, input *rds.DescribeCustomAvailabilityZonesInput) (*rds.DescribeCustomAvailabilityZonesOutput, error)
    DescribeCustomAvailabilityZonesAsync(ctx workflow.Context, input *rds.DescribeCustomAvailabilityZonesInput) *RdsDescribeCustomAvailabilityZonesResult

    DescribeDBClusterBacktracks(ctx workflow.Context, input *rds.DescribeDBClusterBacktracksInput) (*rds.DescribeDBClusterBacktracksOutput, error)
    DescribeDBClusterBacktracksAsync(ctx workflow.Context, input *rds.DescribeDBClusterBacktracksInput) *RdsDescribeDBClusterBacktracksResult

    DescribeDBClusterEndpoints(ctx workflow.Context, input *rds.DescribeDBClusterEndpointsInput) (*rds.DescribeDBClusterEndpointsOutput, error)
    DescribeDBClusterEndpointsAsync(ctx workflow.Context, input *rds.DescribeDBClusterEndpointsInput) *RdsDescribeDBClusterEndpointsResult

    DescribeDBClusterParameterGroups(ctx workflow.Context, input *rds.DescribeDBClusterParameterGroupsInput) (*rds.DescribeDBClusterParameterGroupsOutput, error)
    DescribeDBClusterParameterGroupsAsync(ctx workflow.Context, input *rds.DescribeDBClusterParameterGroupsInput) *RdsDescribeDBClusterParameterGroupsResult

    DescribeDBClusterParameters(ctx workflow.Context, input *rds.DescribeDBClusterParametersInput) (*rds.DescribeDBClusterParametersOutput, error)
    DescribeDBClusterParametersAsync(ctx workflow.Context, input *rds.DescribeDBClusterParametersInput) *RdsDescribeDBClusterParametersResult

    DescribeDBClusterSnapshotAttributes(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotAttributesInput) (*rds.DescribeDBClusterSnapshotAttributesOutput, error)
    DescribeDBClusterSnapshotAttributesAsync(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotAttributesInput) *RdsDescribeDBClusterSnapshotAttributesResult

    DescribeDBClusterSnapshots(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) (*rds.DescribeDBClusterSnapshotsOutput, error)
    DescribeDBClusterSnapshotsAsync(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) *RdsDescribeDBClusterSnapshotsResult

    DescribeDBClusters(ctx workflow.Context, input *rds.DescribeDBClustersInput) (*rds.DescribeDBClustersOutput, error)
    DescribeDBClustersAsync(ctx workflow.Context, input *rds.DescribeDBClustersInput) *RdsDescribeDBClustersResult

    DescribeDBEngineVersions(ctx workflow.Context, input *rds.DescribeDBEngineVersionsInput) (*rds.DescribeDBEngineVersionsOutput, error)
    DescribeDBEngineVersionsAsync(ctx workflow.Context, input *rds.DescribeDBEngineVersionsInput) *RdsDescribeDBEngineVersionsResult

    DescribeDBInstanceAutomatedBackups(ctx workflow.Context, input *rds.DescribeDBInstanceAutomatedBackupsInput) (*rds.DescribeDBInstanceAutomatedBackupsOutput, error)
    DescribeDBInstanceAutomatedBackupsAsync(ctx workflow.Context, input *rds.DescribeDBInstanceAutomatedBackupsInput) *RdsDescribeDBInstanceAutomatedBackupsResult

    DescribeDBInstances(ctx workflow.Context, input *rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error)
    DescribeDBInstancesAsync(ctx workflow.Context, input *rds.DescribeDBInstancesInput) *RdsDescribeDBInstancesResult

    DescribeDBLogFiles(ctx workflow.Context, input *rds.DescribeDBLogFilesInput) (*rds.DescribeDBLogFilesOutput, error)
    DescribeDBLogFilesAsync(ctx workflow.Context, input *rds.DescribeDBLogFilesInput) *RdsDescribeDBLogFilesResult

    DescribeDBParameterGroups(ctx workflow.Context, input *rds.DescribeDBParameterGroupsInput) (*rds.DescribeDBParameterGroupsOutput, error)
    DescribeDBParameterGroupsAsync(ctx workflow.Context, input *rds.DescribeDBParameterGroupsInput) *RdsDescribeDBParameterGroupsResult

    DescribeDBParameters(ctx workflow.Context, input *rds.DescribeDBParametersInput) (*rds.DescribeDBParametersOutput, error)
    DescribeDBParametersAsync(ctx workflow.Context, input *rds.DescribeDBParametersInput) *RdsDescribeDBParametersResult

    DescribeDBProxies(ctx workflow.Context, input *rds.DescribeDBProxiesInput) (*rds.DescribeDBProxiesOutput, error)
    DescribeDBProxiesAsync(ctx workflow.Context, input *rds.DescribeDBProxiesInput) *RdsDescribeDBProxiesResult

    DescribeDBProxyTargetGroups(ctx workflow.Context, input *rds.DescribeDBProxyTargetGroupsInput) (*rds.DescribeDBProxyTargetGroupsOutput, error)
    DescribeDBProxyTargetGroupsAsync(ctx workflow.Context, input *rds.DescribeDBProxyTargetGroupsInput) *RdsDescribeDBProxyTargetGroupsResult

    DescribeDBProxyTargets(ctx workflow.Context, input *rds.DescribeDBProxyTargetsInput) (*rds.DescribeDBProxyTargetsOutput, error)
    DescribeDBProxyTargetsAsync(ctx workflow.Context, input *rds.DescribeDBProxyTargetsInput) *RdsDescribeDBProxyTargetsResult

    DescribeDBSecurityGroups(ctx workflow.Context, input *rds.DescribeDBSecurityGroupsInput) (*rds.DescribeDBSecurityGroupsOutput, error)
    DescribeDBSecurityGroupsAsync(ctx workflow.Context, input *rds.DescribeDBSecurityGroupsInput) *RdsDescribeDBSecurityGroupsResult

    DescribeDBSnapshotAttributes(ctx workflow.Context, input *rds.DescribeDBSnapshotAttributesInput) (*rds.DescribeDBSnapshotAttributesOutput, error)
    DescribeDBSnapshotAttributesAsync(ctx workflow.Context, input *rds.DescribeDBSnapshotAttributesInput) *RdsDescribeDBSnapshotAttributesResult

    DescribeDBSnapshots(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) (*rds.DescribeDBSnapshotsOutput, error)
    DescribeDBSnapshotsAsync(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) *RdsDescribeDBSnapshotsResult

    DescribeDBSubnetGroups(ctx workflow.Context, input *rds.DescribeDBSubnetGroupsInput) (*rds.DescribeDBSubnetGroupsOutput, error)
    DescribeDBSubnetGroupsAsync(ctx workflow.Context, input *rds.DescribeDBSubnetGroupsInput) *RdsDescribeDBSubnetGroupsResult

    DescribeEngineDefaultClusterParameters(ctx workflow.Context, input *rds.DescribeEngineDefaultClusterParametersInput) (*rds.DescribeEngineDefaultClusterParametersOutput, error)
    DescribeEngineDefaultClusterParametersAsync(ctx workflow.Context, input *rds.DescribeEngineDefaultClusterParametersInput) *RdsDescribeEngineDefaultClusterParametersResult

    DescribeEngineDefaultParameters(ctx workflow.Context, input *rds.DescribeEngineDefaultParametersInput) (*rds.DescribeEngineDefaultParametersOutput, error)
    DescribeEngineDefaultParametersAsync(ctx workflow.Context, input *rds.DescribeEngineDefaultParametersInput) *RdsDescribeEngineDefaultParametersResult

    DescribeEventCategories(ctx workflow.Context, input *rds.DescribeEventCategoriesInput) (*rds.DescribeEventCategoriesOutput, error)
    DescribeEventCategoriesAsync(ctx workflow.Context, input *rds.DescribeEventCategoriesInput) *RdsDescribeEventCategoriesResult

    DescribeEventSubscriptions(ctx workflow.Context, input *rds.DescribeEventSubscriptionsInput) (*rds.DescribeEventSubscriptionsOutput, error)
    DescribeEventSubscriptionsAsync(ctx workflow.Context, input *rds.DescribeEventSubscriptionsInput) *RdsDescribeEventSubscriptionsResult

    DescribeEvents(ctx workflow.Context, input *rds.DescribeEventsInput) (*rds.DescribeEventsOutput, error)
    DescribeEventsAsync(ctx workflow.Context, input *rds.DescribeEventsInput) *RdsDescribeEventsResult

    DescribeExportTasks(ctx workflow.Context, input *rds.DescribeExportTasksInput) (*rds.DescribeExportTasksOutput, error)
    DescribeExportTasksAsync(ctx workflow.Context, input *rds.DescribeExportTasksInput) *RdsDescribeExportTasksResult

    DescribeGlobalClusters(ctx workflow.Context, input *rds.DescribeGlobalClustersInput) (*rds.DescribeGlobalClustersOutput, error)
    DescribeGlobalClustersAsync(ctx workflow.Context, input *rds.DescribeGlobalClustersInput) *RdsDescribeGlobalClustersResult

    DescribeInstallationMedia(ctx workflow.Context, input *rds.DescribeInstallationMediaInput) (*rds.DescribeInstallationMediaOutput, error)
    DescribeInstallationMediaAsync(ctx workflow.Context, input *rds.DescribeInstallationMediaInput) *RdsDescribeInstallationMediaResult

    DescribeOptionGroupOptions(ctx workflow.Context, input *rds.DescribeOptionGroupOptionsInput) (*rds.DescribeOptionGroupOptionsOutput, error)
    DescribeOptionGroupOptionsAsync(ctx workflow.Context, input *rds.DescribeOptionGroupOptionsInput) *RdsDescribeOptionGroupOptionsResult

    DescribeOptionGroups(ctx workflow.Context, input *rds.DescribeOptionGroupsInput) (*rds.DescribeOptionGroupsOutput, error)
    DescribeOptionGroupsAsync(ctx workflow.Context, input *rds.DescribeOptionGroupsInput) *RdsDescribeOptionGroupsResult

    DescribeOrderableDBInstanceOptions(ctx workflow.Context, input *rds.DescribeOrderableDBInstanceOptionsInput) (*rds.DescribeOrderableDBInstanceOptionsOutput, error)
    DescribeOrderableDBInstanceOptionsAsync(ctx workflow.Context, input *rds.DescribeOrderableDBInstanceOptionsInput) *RdsDescribeOrderableDBInstanceOptionsResult

    DescribePendingMaintenanceActions(ctx workflow.Context, input *rds.DescribePendingMaintenanceActionsInput) (*rds.DescribePendingMaintenanceActionsOutput, error)
    DescribePendingMaintenanceActionsAsync(ctx workflow.Context, input *rds.DescribePendingMaintenanceActionsInput) *RdsDescribePendingMaintenanceActionsResult

    DescribeReservedDBInstances(ctx workflow.Context, input *rds.DescribeReservedDBInstancesInput) (*rds.DescribeReservedDBInstancesOutput, error)
    DescribeReservedDBInstancesAsync(ctx workflow.Context, input *rds.DescribeReservedDBInstancesInput) *RdsDescribeReservedDBInstancesResult

    DescribeReservedDBInstancesOfferings(ctx workflow.Context, input *rds.DescribeReservedDBInstancesOfferingsInput) (*rds.DescribeReservedDBInstancesOfferingsOutput, error)
    DescribeReservedDBInstancesOfferingsAsync(ctx workflow.Context, input *rds.DescribeReservedDBInstancesOfferingsInput) *RdsDescribeReservedDBInstancesOfferingsResult

    DescribeSourceRegions(ctx workflow.Context, input *rds.DescribeSourceRegionsInput) (*rds.DescribeSourceRegionsOutput, error)
    DescribeSourceRegionsAsync(ctx workflow.Context, input *rds.DescribeSourceRegionsInput) *RdsDescribeSourceRegionsResult

    DescribeValidDBInstanceModifications(ctx workflow.Context, input *rds.DescribeValidDBInstanceModificationsInput) (*rds.DescribeValidDBInstanceModificationsOutput, error)
    DescribeValidDBInstanceModificationsAsync(ctx workflow.Context, input *rds.DescribeValidDBInstanceModificationsInput) *RdsDescribeValidDBInstanceModificationsResult

    DownloadDBLogFilePortion(ctx workflow.Context, input *rds.DownloadDBLogFilePortionInput) (*rds.DownloadDBLogFilePortionOutput, error)
    DownloadDBLogFilePortionAsync(ctx workflow.Context, input *rds.DownloadDBLogFilePortionInput) *RdsDownloadDBLogFilePortionResult

    FailoverDBCluster(ctx workflow.Context, input *rds.FailoverDBClusterInput) (*rds.FailoverDBClusterOutput, error)
    FailoverDBClusterAsync(ctx workflow.Context, input *rds.FailoverDBClusterInput) *RdsFailoverDBClusterResult

    ImportInstallationMedia(ctx workflow.Context, input *rds.ImportInstallationMediaInput) (*rds.ImportInstallationMediaOutput, error)
    ImportInstallationMediaAsync(ctx workflow.Context, input *rds.ImportInstallationMediaInput) *RdsImportInstallationMediaResult

    ListTagsForResource(ctx workflow.Context, input *rds.ListTagsForResourceInput) (*rds.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *rds.ListTagsForResourceInput) *RdsListTagsForResourceResult

    ModifyCertificates(ctx workflow.Context, input *rds.ModifyCertificatesInput) (*rds.ModifyCertificatesOutput, error)
    ModifyCertificatesAsync(ctx workflow.Context, input *rds.ModifyCertificatesInput) *RdsModifyCertificatesResult

    ModifyCurrentDBClusterCapacity(ctx workflow.Context, input *rds.ModifyCurrentDBClusterCapacityInput) (*rds.ModifyCurrentDBClusterCapacityOutput, error)
    ModifyCurrentDBClusterCapacityAsync(ctx workflow.Context, input *rds.ModifyCurrentDBClusterCapacityInput) *RdsModifyCurrentDBClusterCapacityResult

    ModifyDBCluster(ctx workflow.Context, input *rds.ModifyDBClusterInput) (*rds.ModifyDBClusterOutput, error)
    ModifyDBClusterAsync(ctx workflow.Context, input *rds.ModifyDBClusterInput) *RdsModifyDBClusterResult

    ModifyDBClusterEndpoint(ctx workflow.Context, input *rds.ModifyDBClusterEndpointInput) (*rds.ModifyDBClusterEndpointOutput, error)
    ModifyDBClusterEndpointAsync(ctx workflow.Context, input *rds.ModifyDBClusterEndpointInput) *RdsModifyDBClusterEndpointResult

    ModifyDBClusterParameterGroup(ctx workflow.Context, input *rds.ModifyDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error)
    ModifyDBClusterParameterGroupAsync(ctx workflow.Context, input *rds.ModifyDBClusterParameterGroupInput) *RdsModifyDBClusterParameterGroupResult

    ModifyDBClusterSnapshotAttribute(ctx workflow.Context, input *rds.ModifyDBClusterSnapshotAttributeInput) (*rds.ModifyDBClusterSnapshotAttributeOutput, error)
    ModifyDBClusterSnapshotAttributeAsync(ctx workflow.Context, input *rds.ModifyDBClusterSnapshotAttributeInput) *RdsModifyDBClusterSnapshotAttributeResult

    ModifyDBInstance(ctx workflow.Context, input *rds.ModifyDBInstanceInput) (*rds.ModifyDBInstanceOutput, error)
    ModifyDBInstanceAsync(ctx workflow.Context, input *rds.ModifyDBInstanceInput) *RdsModifyDBInstanceResult

    ModifyDBParameterGroup(ctx workflow.Context, input *rds.ModifyDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error)
    ModifyDBParameterGroupAsync(ctx workflow.Context, input *rds.ModifyDBParameterGroupInput) *RdsModifyDBParameterGroupResult

    ModifyDBProxy(ctx workflow.Context, input *rds.ModifyDBProxyInput) (*rds.ModifyDBProxyOutput, error)
    ModifyDBProxyAsync(ctx workflow.Context, input *rds.ModifyDBProxyInput) *RdsModifyDBProxyResult

    ModifyDBProxyTargetGroup(ctx workflow.Context, input *rds.ModifyDBProxyTargetGroupInput) (*rds.ModifyDBProxyTargetGroupOutput, error)
    ModifyDBProxyTargetGroupAsync(ctx workflow.Context, input *rds.ModifyDBProxyTargetGroupInput) *RdsModifyDBProxyTargetGroupResult

    ModifyDBSnapshot(ctx workflow.Context, input *rds.ModifyDBSnapshotInput) (*rds.ModifyDBSnapshotOutput, error)
    ModifyDBSnapshotAsync(ctx workflow.Context, input *rds.ModifyDBSnapshotInput) *RdsModifyDBSnapshotResult

    ModifyDBSnapshotAttribute(ctx workflow.Context, input *rds.ModifyDBSnapshotAttributeInput) (*rds.ModifyDBSnapshotAttributeOutput, error)
    ModifyDBSnapshotAttributeAsync(ctx workflow.Context, input *rds.ModifyDBSnapshotAttributeInput) *RdsModifyDBSnapshotAttributeResult

    ModifyDBSubnetGroup(ctx workflow.Context, input *rds.ModifyDBSubnetGroupInput) (*rds.ModifyDBSubnetGroupOutput, error)
    ModifyDBSubnetGroupAsync(ctx workflow.Context, input *rds.ModifyDBSubnetGroupInput) *RdsModifyDBSubnetGroupResult

    ModifyEventSubscription(ctx workflow.Context, input *rds.ModifyEventSubscriptionInput) (*rds.ModifyEventSubscriptionOutput, error)
    ModifyEventSubscriptionAsync(ctx workflow.Context, input *rds.ModifyEventSubscriptionInput) *RdsModifyEventSubscriptionResult

    ModifyGlobalCluster(ctx workflow.Context, input *rds.ModifyGlobalClusterInput) (*rds.ModifyGlobalClusterOutput, error)
    ModifyGlobalClusterAsync(ctx workflow.Context, input *rds.ModifyGlobalClusterInput) *RdsModifyGlobalClusterResult

    ModifyOptionGroup(ctx workflow.Context, input *rds.ModifyOptionGroupInput) (*rds.ModifyOptionGroupOutput, error)
    ModifyOptionGroupAsync(ctx workflow.Context, input *rds.ModifyOptionGroupInput) *RdsModifyOptionGroupResult

    PromoteReadReplica(ctx workflow.Context, input *rds.PromoteReadReplicaInput) (*rds.PromoteReadReplicaOutput, error)
    PromoteReadReplicaAsync(ctx workflow.Context, input *rds.PromoteReadReplicaInput) *RdsPromoteReadReplicaResult

    PromoteReadReplicaDBCluster(ctx workflow.Context, input *rds.PromoteReadReplicaDBClusterInput) (*rds.PromoteReadReplicaDBClusterOutput, error)
    PromoteReadReplicaDBClusterAsync(ctx workflow.Context, input *rds.PromoteReadReplicaDBClusterInput) *RdsPromoteReadReplicaDBClusterResult

    PurchaseReservedDBInstancesOffering(ctx workflow.Context, input *rds.PurchaseReservedDBInstancesOfferingInput) (*rds.PurchaseReservedDBInstancesOfferingOutput, error)
    PurchaseReservedDBInstancesOfferingAsync(ctx workflow.Context, input *rds.PurchaseReservedDBInstancesOfferingInput) *RdsPurchaseReservedDBInstancesOfferingResult

    RebootDBInstance(ctx workflow.Context, input *rds.RebootDBInstanceInput) (*rds.RebootDBInstanceOutput, error)
    RebootDBInstanceAsync(ctx workflow.Context, input *rds.RebootDBInstanceInput) *RdsRebootDBInstanceResult

    RegisterDBProxyTargets(ctx workflow.Context, input *rds.RegisterDBProxyTargetsInput) (*rds.RegisterDBProxyTargetsOutput, error)
    RegisterDBProxyTargetsAsync(ctx workflow.Context, input *rds.RegisterDBProxyTargetsInput) *RdsRegisterDBProxyTargetsResult

    RemoveFromGlobalCluster(ctx workflow.Context, input *rds.RemoveFromGlobalClusterInput) (*rds.RemoveFromGlobalClusterOutput, error)
    RemoveFromGlobalClusterAsync(ctx workflow.Context, input *rds.RemoveFromGlobalClusterInput) *RdsRemoveFromGlobalClusterResult

    RemoveRoleFromDBCluster(ctx workflow.Context, input *rds.RemoveRoleFromDBClusterInput) (*rds.RemoveRoleFromDBClusterOutput, error)
    RemoveRoleFromDBClusterAsync(ctx workflow.Context, input *rds.RemoveRoleFromDBClusterInput) *RdsRemoveRoleFromDBClusterResult

    RemoveRoleFromDBInstance(ctx workflow.Context, input *rds.RemoveRoleFromDBInstanceInput) (*rds.RemoveRoleFromDBInstanceOutput, error)
    RemoveRoleFromDBInstanceAsync(ctx workflow.Context, input *rds.RemoveRoleFromDBInstanceInput) *RdsRemoveRoleFromDBInstanceResult

    RemoveSourceIdentifierFromSubscription(ctx workflow.Context, input *rds.RemoveSourceIdentifierFromSubscriptionInput) (*rds.RemoveSourceIdentifierFromSubscriptionOutput, error)
    RemoveSourceIdentifierFromSubscriptionAsync(ctx workflow.Context, input *rds.RemoveSourceIdentifierFromSubscriptionInput) *RdsRemoveSourceIdentifierFromSubscriptionResult

    RemoveTagsFromResource(ctx workflow.Context, input *rds.RemoveTagsFromResourceInput) (*rds.RemoveTagsFromResourceOutput, error)
    RemoveTagsFromResourceAsync(ctx workflow.Context, input *rds.RemoveTagsFromResourceInput) *RdsRemoveTagsFromResourceResult

    ResetDBClusterParameterGroup(ctx workflow.Context, input *rds.ResetDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error)
    ResetDBClusterParameterGroupAsync(ctx workflow.Context, input *rds.ResetDBClusterParameterGroupInput) *RdsResetDBClusterParameterGroupResult

    ResetDBParameterGroup(ctx workflow.Context, input *rds.ResetDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error)
    ResetDBParameterGroupAsync(ctx workflow.Context, input *rds.ResetDBParameterGroupInput) *RdsResetDBParameterGroupResult

    RestoreDBClusterFromS3(ctx workflow.Context, input *rds.RestoreDBClusterFromS3Input) (*rds.RestoreDBClusterFromS3Output, error)
    RestoreDBClusterFromS3Async(ctx workflow.Context, input *rds.RestoreDBClusterFromS3Input) *RdsRestoreDBClusterFromS3Result

    RestoreDBClusterFromSnapshot(ctx workflow.Context, input *rds.RestoreDBClusterFromSnapshotInput) (*rds.RestoreDBClusterFromSnapshotOutput, error)
    RestoreDBClusterFromSnapshotAsync(ctx workflow.Context, input *rds.RestoreDBClusterFromSnapshotInput) *RdsRestoreDBClusterFromSnapshotResult

    RestoreDBClusterToPointInTime(ctx workflow.Context, input *rds.RestoreDBClusterToPointInTimeInput) (*rds.RestoreDBClusterToPointInTimeOutput, error)
    RestoreDBClusterToPointInTimeAsync(ctx workflow.Context, input *rds.RestoreDBClusterToPointInTimeInput) *RdsRestoreDBClusterToPointInTimeResult

    RestoreDBInstanceFromDBSnapshot(ctx workflow.Context, input *rds.RestoreDBInstanceFromDBSnapshotInput) (*rds.RestoreDBInstanceFromDBSnapshotOutput, error)
    RestoreDBInstanceFromDBSnapshotAsync(ctx workflow.Context, input *rds.RestoreDBInstanceFromDBSnapshotInput) *RdsRestoreDBInstanceFromDBSnapshotResult

    RestoreDBInstanceFromS3(ctx workflow.Context, input *rds.RestoreDBInstanceFromS3Input) (*rds.RestoreDBInstanceFromS3Output, error)
    RestoreDBInstanceFromS3Async(ctx workflow.Context, input *rds.RestoreDBInstanceFromS3Input) *RdsRestoreDBInstanceFromS3Result

    RestoreDBInstanceToPointInTime(ctx workflow.Context, input *rds.RestoreDBInstanceToPointInTimeInput) (*rds.RestoreDBInstanceToPointInTimeOutput, error)
    RestoreDBInstanceToPointInTimeAsync(ctx workflow.Context, input *rds.RestoreDBInstanceToPointInTimeInput) *RdsRestoreDBInstanceToPointInTimeResult

    RevokeDBSecurityGroupIngress(ctx workflow.Context, input *rds.RevokeDBSecurityGroupIngressInput) (*rds.RevokeDBSecurityGroupIngressOutput, error)
    RevokeDBSecurityGroupIngressAsync(ctx workflow.Context, input *rds.RevokeDBSecurityGroupIngressInput) *RdsRevokeDBSecurityGroupIngressResult

    StartActivityStream(ctx workflow.Context, input *rds.StartActivityStreamInput) (*rds.StartActivityStreamOutput, error)
    StartActivityStreamAsync(ctx workflow.Context, input *rds.StartActivityStreamInput) *RdsStartActivityStreamResult

    StartDBCluster(ctx workflow.Context, input *rds.StartDBClusterInput) (*rds.StartDBClusterOutput, error)
    StartDBClusterAsync(ctx workflow.Context, input *rds.StartDBClusterInput) *RdsStartDBClusterResult

    StartDBInstance(ctx workflow.Context, input *rds.StartDBInstanceInput) (*rds.StartDBInstanceOutput, error)
    StartDBInstanceAsync(ctx workflow.Context, input *rds.StartDBInstanceInput) *RdsStartDBInstanceResult

    StartExportTask(ctx workflow.Context, input *rds.StartExportTaskInput) (*rds.StartExportTaskOutput, error)
    StartExportTaskAsync(ctx workflow.Context, input *rds.StartExportTaskInput) *RdsStartExportTaskResult

    StopActivityStream(ctx workflow.Context, input *rds.StopActivityStreamInput) (*rds.StopActivityStreamOutput, error)
    StopActivityStreamAsync(ctx workflow.Context, input *rds.StopActivityStreamInput) *RdsStopActivityStreamResult

    StopDBCluster(ctx workflow.Context, input *rds.StopDBClusterInput) (*rds.StopDBClusterOutput, error)
    StopDBClusterAsync(ctx workflow.Context, input *rds.StopDBClusterInput) *RdsStopDBClusterResult

    StopDBInstance(ctx workflow.Context, input *rds.StopDBInstanceInput) (*rds.StopDBInstanceOutput, error)
    StopDBInstanceAsync(ctx workflow.Context, input *rds.StopDBInstanceInput) *RdsStopDBInstanceResult

    WaitUntilDBClusterSnapshotAvailable(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) error
    WaitUntilDBClusterSnapshotDeleted(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) error
    WaitUntilDBInstanceAvailable(ctx workflow.Context, input *rds.DescribeDBInstancesInput) error
    WaitUntilDBInstanceDeleted(ctx workflow.Context, input *rds.DescribeDBInstancesInput) error
    WaitUntilDBSnapshotAvailable(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) error
    WaitUntilDBSnapshotDeleted(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) error}
type RdsAddRoleToDBClusterResult struct {
	Result workflow.Future
}

func (r *RdsAddRoleToDBClusterResult) Get(ctx workflow.Context) (*rds.AddRoleToDBClusterOutput, error) {
    var output rds.AddRoleToDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsAddRoleToDBInstanceResult struct {
	Result workflow.Future
}

func (r *RdsAddRoleToDBInstanceResult) Get(ctx workflow.Context) (*rds.AddRoleToDBInstanceOutput, error) {
    var output rds.AddRoleToDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsAddSourceIdentifierToSubscriptionResult struct {
	Result workflow.Future
}

func (r *RdsAddSourceIdentifierToSubscriptionResult) Get(ctx workflow.Context) (*rds.AddSourceIdentifierToSubscriptionOutput, error) {
    var output rds.AddSourceIdentifierToSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsAddTagsToResourceResult struct {
	Result workflow.Future
}

func (r *RdsAddTagsToResourceResult) Get(ctx workflow.Context) (*rds.AddTagsToResourceOutput, error) {
    var output rds.AddTagsToResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsApplyPendingMaintenanceActionResult struct {
	Result workflow.Future
}

func (r *RdsApplyPendingMaintenanceActionResult) Get(ctx workflow.Context) (*rds.ApplyPendingMaintenanceActionOutput, error) {
    var output rds.ApplyPendingMaintenanceActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsAuthorizeDBSecurityGroupIngressResult struct {
	Result workflow.Future
}

func (r *RdsAuthorizeDBSecurityGroupIngressResult) Get(ctx workflow.Context) (*rds.AuthorizeDBSecurityGroupIngressOutput, error) {
    var output rds.AuthorizeDBSecurityGroupIngressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsBacktrackDBClusterResult struct {
	Result workflow.Future
}

func (r *RdsBacktrackDBClusterResult) Get(ctx workflow.Context) (*rds.BacktrackDBClusterOutput, error) {
    var output rds.BacktrackDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCancelExportTaskResult struct {
	Result workflow.Future
}

func (r *RdsCancelExportTaskResult) Get(ctx workflow.Context) (*rds.CancelExportTaskOutput, error) {
    var output rds.CancelExportTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCopyDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *RdsCopyDBClusterParameterGroupResult) Get(ctx workflow.Context) (*rds.CopyDBClusterParameterGroupOutput, error) {
    var output rds.CopyDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCopyDBClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *RdsCopyDBClusterSnapshotResult) Get(ctx workflow.Context) (*rds.CopyDBClusterSnapshotOutput, error) {
    var output rds.CopyDBClusterSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCopyDBParameterGroupResult struct {
	Result workflow.Future
}

func (r *RdsCopyDBParameterGroupResult) Get(ctx workflow.Context) (*rds.CopyDBParameterGroupOutput, error) {
    var output rds.CopyDBParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCopyDBSnapshotResult struct {
	Result workflow.Future
}

func (r *RdsCopyDBSnapshotResult) Get(ctx workflow.Context) (*rds.CopyDBSnapshotOutput, error) {
    var output rds.CopyDBSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCopyOptionGroupResult struct {
	Result workflow.Future
}

func (r *RdsCopyOptionGroupResult) Get(ctx workflow.Context) (*rds.CopyOptionGroupOutput, error) {
    var output rds.CopyOptionGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateCustomAvailabilityZoneResult struct {
	Result workflow.Future
}

func (r *RdsCreateCustomAvailabilityZoneResult) Get(ctx workflow.Context) (*rds.CreateCustomAvailabilityZoneOutput, error) {
    var output rds.CreateCustomAvailabilityZoneOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateDBClusterResult struct {
	Result workflow.Future
}

func (r *RdsCreateDBClusterResult) Get(ctx workflow.Context) (*rds.CreateDBClusterOutput, error) {
    var output rds.CreateDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateDBClusterEndpointResult struct {
	Result workflow.Future
}

func (r *RdsCreateDBClusterEndpointResult) Get(ctx workflow.Context) (*rds.CreateDBClusterEndpointOutput, error) {
    var output rds.CreateDBClusterEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *RdsCreateDBClusterParameterGroupResult) Get(ctx workflow.Context) (*rds.CreateDBClusterParameterGroupOutput, error) {
    var output rds.CreateDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateDBClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *RdsCreateDBClusterSnapshotResult) Get(ctx workflow.Context) (*rds.CreateDBClusterSnapshotOutput, error) {
    var output rds.CreateDBClusterSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateDBInstanceResult struct {
	Result workflow.Future
}

func (r *RdsCreateDBInstanceResult) Get(ctx workflow.Context) (*rds.CreateDBInstanceOutput, error) {
    var output rds.CreateDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateDBInstanceReadReplicaResult struct {
	Result workflow.Future
}

func (r *RdsCreateDBInstanceReadReplicaResult) Get(ctx workflow.Context) (*rds.CreateDBInstanceReadReplicaOutput, error) {
    var output rds.CreateDBInstanceReadReplicaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateDBParameterGroupResult struct {
	Result workflow.Future
}

func (r *RdsCreateDBParameterGroupResult) Get(ctx workflow.Context) (*rds.CreateDBParameterGroupOutput, error) {
    var output rds.CreateDBParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateDBProxyResult struct {
	Result workflow.Future
}

func (r *RdsCreateDBProxyResult) Get(ctx workflow.Context) (*rds.CreateDBProxyOutput, error) {
    var output rds.CreateDBProxyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateDBSecurityGroupResult struct {
	Result workflow.Future
}

func (r *RdsCreateDBSecurityGroupResult) Get(ctx workflow.Context) (*rds.CreateDBSecurityGroupOutput, error) {
    var output rds.CreateDBSecurityGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateDBSnapshotResult struct {
	Result workflow.Future
}

func (r *RdsCreateDBSnapshotResult) Get(ctx workflow.Context) (*rds.CreateDBSnapshotOutput, error) {
    var output rds.CreateDBSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateDBSubnetGroupResult struct {
	Result workflow.Future
}

func (r *RdsCreateDBSubnetGroupResult) Get(ctx workflow.Context) (*rds.CreateDBSubnetGroupOutput, error) {
    var output rds.CreateDBSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *RdsCreateEventSubscriptionResult) Get(ctx workflow.Context) (*rds.CreateEventSubscriptionOutput, error) {
    var output rds.CreateEventSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateGlobalClusterResult struct {
	Result workflow.Future
}

func (r *RdsCreateGlobalClusterResult) Get(ctx workflow.Context) (*rds.CreateGlobalClusterOutput, error) {
    var output rds.CreateGlobalClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsCreateOptionGroupResult struct {
	Result workflow.Future
}

func (r *RdsCreateOptionGroupResult) Get(ctx workflow.Context) (*rds.CreateOptionGroupOutput, error) {
    var output rds.CreateOptionGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteCustomAvailabilityZoneResult struct {
	Result workflow.Future
}

func (r *RdsDeleteCustomAvailabilityZoneResult) Get(ctx workflow.Context) (*rds.DeleteCustomAvailabilityZoneOutput, error) {
    var output rds.DeleteCustomAvailabilityZoneOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteDBClusterResult struct {
	Result workflow.Future
}

func (r *RdsDeleteDBClusterResult) Get(ctx workflow.Context) (*rds.DeleteDBClusterOutput, error) {
    var output rds.DeleteDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteDBClusterEndpointResult struct {
	Result workflow.Future
}

func (r *RdsDeleteDBClusterEndpointResult) Get(ctx workflow.Context) (*rds.DeleteDBClusterEndpointOutput, error) {
    var output rds.DeleteDBClusterEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *RdsDeleteDBClusterParameterGroupResult) Get(ctx workflow.Context) (*rds.DeleteDBClusterParameterGroupOutput, error) {
    var output rds.DeleteDBClusterParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteDBClusterSnapshotResult struct {
	Result workflow.Future
}

func (r *RdsDeleteDBClusterSnapshotResult) Get(ctx workflow.Context) (*rds.DeleteDBClusterSnapshotOutput, error) {
    var output rds.DeleteDBClusterSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteDBInstanceResult struct {
	Result workflow.Future
}

func (r *RdsDeleteDBInstanceResult) Get(ctx workflow.Context) (*rds.DeleteDBInstanceOutput, error) {
    var output rds.DeleteDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteDBInstanceAutomatedBackupResult struct {
	Result workflow.Future
}

func (r *RdsDeleteDBInstanceAutomatedBackupResult) Get(ctx workflow.Context) (*rds.DeleteDBInstanceAutomatedBackupOutput, error) {
    var output rds.DeleteDBInstanceAutomatedBackupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteDBParameterGroupResult struct {
	Result workflow.Future
}

func (r *RdsDeleteDBParameterGroupResult) Get(ctx workflow.Context) (*rds.DeleteDBParameterGroupOutput, error) {
    var output rds.DeleteDBParameterGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteDBProxyResult struct {
	Result workflow.Future
}

func (r *RdsDeleteDBProxyResult) Get(ctx workflow.Context) (*rds.DeleteDBProxyOutput, error) {
    var output rds.DeleteDBProxyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteDBSecurityGroupResult struct {
	Result workflow.Future
}

func (r *RdsDeleteDBSecurityGroupResult) Get(ctx workflow.Context) (*rds.DeleteDBSecurityGroupOutput, error) {
    var output rds.DeleteDBSecurityGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteDBSnapshotResult struct {
	Result workflow.Future
}

func (r *RdsDeleteDBSnapshotResult) Get(ctx workflow.Context) (*rds.DeleteDBSnapshotOutput, error) {
    var output rds.DeleteDBSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteDBSubnetGroupResult struct {
	Result workflow.Future
}

func (r *RdsDeleteDBSubnetGroupResult) Get(ctx workflow.Context) (*rds.DeleteDBSubnetGroupOutput, error) {
    var output rds.DeleteDBSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *RdsDeleteEventSubscriptionResult) Get(ctx workflow.Context) (*rds.DeleteEventSubscriptionOutput, error) {
    var output rds.DeleteEventSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteGlobalClusterResult struct {
	Result workflow.Future
}

func (r *RdsDeleteGlobalClusterResult) Get(ctx workflow.Context) (*rds.DeleteGlobalClusterOutput, error) {
    var output rds.DeleteGlobalClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteInstallationMediaResult struct {
	Result workflow.Future
}

func (r *RdsDeleteInstallationMediaResult) Get(ctx workflow.Context) (*rds.DeleteInstallationMediaOutput, error) {
    var output rds.DeleteInstallationMediaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeleteOptionGroupResult struct {
	Result workflow.Future
}

func (r *RdsDeleteOptionGroupResult) Get(ctx workflow.Context) (*rds.DeleteOptionGroupOutput, error) {
    var output rds.DeleteOptionGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDeregisterDBProxyTargetsResult struct {
	Result workflow.Future
}

func (r *RdsDeregisterDBProxyTargetsResult) Get(ctx workflow.Context) (*rds.DeregisterDBProxyTargetsOutput, error) {
    var output rds.DeregisterDBProxyTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeAccountAttributesResult struct {
	Result workflow.Future
}

func (r *RdsDescribeAccountAttributesResult) Get(ctx workflow.Context) (*rds.DescribeAccountAttributesOutput, error) {
    var output rds.DescribeAccountAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeCertificatesResult struct {
	Result workflow.Future
}

func (r *RdsDescribeCertificatesResult) Get(ctx workflow.Context) (*rds.DescribeCertificatesOutput, error) {
    var output rds.DescribeCertificatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeCustomAvailabilityZonesResult struct {
	Result workflow.Future
}

func (r *RdsDescribeCustomAvailabilityZonesResult) Get(ctx workflow.Context) (*rds.DescribeCustomAvailabilityZonesOutput, error) {
    var output rds.DescribeCustomAvailabilityZonesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBClusterBacktracksResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBClusterBacktracksResult) Get(ctx workflow.Context) (*rds.DescribeDBClusterBacktracksOutput, error) {
    var output rds.DescribeDBClusterBacktracksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBClusterEndpointsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBClusterEndpointsResult) Get(ctx workflow.Context) (*rds.DescribeDBClusterEndpointsOutput, error) {
    var output rds.DescribeDBClusterEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBClusterParameterGroupsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBClusterParameterGroupsResult) Get(ctx workflow.Context) (*rds.DescribeDBClusterParameterGroupsOutput, error) {
    var output rds.DescribeDBClusterParameterGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBClusterParametersResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBClusterParametersResult) Get(ctx workflow.Context) (*rds.DescribeDBClusterParametersOutput, error) {
    var output rds.DescribeDBClusterParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBClusterSnapshotAttributesResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBClusterSnapshotAttributesResult) Get(ctx workflow.Context) (*rds.DescribeDBClusterSnapshotAttributesOutput, error) {
    var output rds.DescribeDBClusterSnapshotAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBClusterSnapshotsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBClusterSnapshotsResult) Get(ctx workflow.Context) (*rds.DescribeDBClusterSnapshotsOutput, error) {
    var output rds.DescribeDBClusterSnapshotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBClustersResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBClustersResult) Get(ctx workflow.Context) (*rds.DescribeDBClustersOutput, error) {
    var output rds.DescribeDBClustersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBEngineVersionsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBEngineVersionsResult) Get(ctx workflow.Context) (*rds.DescribeDBEngineVersionsOutput, error) {
    var output rds.DescribeDBEngineVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBInstanceAutomatedBackupsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBInstanceAutomatedBackupsResult) Get(ctx workflow.Context) (*rds.DescribeDBInstanceAutomatedBackupsOutput, error) {
    var output rds.DescribeDBInstanceAutomatedBackupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBInstancesResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBInstancesResult) Get(ctx workflow.Context) (*rds.DescribeDBInstancesOutput, error) {
    var output rds.DescribeDBInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBLogFilesResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBLogFilesResult) Get(ctx workflow.Context) (*rds.DescribeDBLogFilesOutput, error) {
    var output rds.DescribeDBLogFilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBParameterGroupsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBParameterGroupsResult) Get(ctx workflow.Context) (*rds.DescribeDBParameterGroupsOutput, error) {
    var output rds.DescribeDBParameterGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBParametersResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBParametersResult) Get(ctx workflow.Context) (*rds.DescribeDBParametersOutput, error) {
    var output rds.DescribeDBParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBProxiesResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBProxiesResult) Get(ctx workflow.Context) (*rds.DescribeDBProxiesOutput, error) {
    var output rds.DescribeDBProxiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBProxyTargetGroupsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBProxyTargetGroupsResult) Get(ctx workflow.Context) (*rds.DescribeDBProxyTargetGroupsOutput, error) {
    var output rds.DescribeDBProxyTargetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBProxyTargetsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBProxyTargetsResult) Get(ctx workflow.Context) (*rds.DescribeDBProxyTargetsOutput, error) {
    var output rds.DescribeDBProxyTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBSecurityGroupsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBSecurityGroupsResult) Get(ctx workflow.Context) (*rds.DescribeDBSecurityGroupsOutput, error) {
    var output rds.DescribeDBSecurityGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBSnapshotAttributesResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBSnapshotAttributesResult) Get(ctx workflow.Context) (*rds.DescribeDBSnapshotAttributesOutput, error) {
    var output rds.DescribeDBSnapshotAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBSnapshotsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBSnapshotsResult) Get(ctx workflow.Context) (*rds.DescribeDBSnapshotsOutput, error) {
    var output rds.DescribeDBSnapshotsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeDBSubnetGroupsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeDBSubnetGroupsResult) Get(ctx workflow.Context) (*rds.DescribeDBSubnetGroupsOutput, error) {
    var output rds.DescribeDBSubnetGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeEngineDefaultClusterParametersResult struct {
	Result workflow.Future
}

func (r *RdsDescribeEngineDefaultClusterParametersResult) Get(ctx workflow.Context) (*rds.DescribeEngineDefaultClusterParametersOutput, error) {
    var output rds.DescribeEngineDefaultClusterParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeEngineDefaultParametersResult struct {
	Result workflow.Future
}

func (r *RdsDescribeEngineDefaultParametersResult) Get(ctx workflow.Context) (*rds.DescribeEngineDefaultParametersOutput, error) {
    var output rds.DescribeEngineDefaultParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeEventCategoriesResult struct {
	Result workflow.Future
}

func (r *RdsDescribeEventCategoriesResult) Get(ctx workflow.Context) (*rds.DescribeEventCategoriesOutput, error) {
    var output rds.DescribeEventCategoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeEventSubscriptionsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeEventSubscriptionsResult) Get(ctx workflow.Context) (*rds.DescribeEventSubscriptionsOutput, error) {
    var output rds.DescribeEventSubscriptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeEventsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeEventsResult) Get(ctx workflow.Context) (*rds.DescribeEventsOutput, error) {
    var output rds.DescribeEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeExportTasksResult struct {
	Result workflow.Future
}

func (r *RdsDescribeExportTasksResult) Get(ctx workflow.Context) (*rds.DescribeExportTasksOutput, error) {
    var output rds.DescribeExportTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeGlobalClustersResult struct {
	Result workflow.Future
}

func (r *RdsDescribeGlobalClustersResult) Get(ctx workflow.Context) (*rds.DescribeGlobalClustersOutput, error) {
    var output rds.DescribeGlobalClustersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeInstallationMediaResult struct {
	Result workflow.Future
}

func (r *RdsDescribeInstallationMediaResult) Get(ctx workflow.Context) (*rds.DescribeInstallationMediaOutput, error) {
    var output rds.DescribeInstallationMediaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeOptionGroupOptionsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeOptionGroupOptionsResult) Get(ctx workflow.Context) (*rds.DescribeOptionGroupOptionsOutput, error) {
    var output rds.DescribeOptionGroupOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeOptionGroupsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeOptionGroupsResult) Get(ctx workflow.Context) (*rds.DescribeOptionGroupsOutput, error) {
    var output rds.DescribeOptionGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeOrderableDBInstanceOptionsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeOrderableDBInstanceOptionsResult) Get(ctx workflow.Context) (*rds.DescribeOrderableDBInstanceOptionsOutput, error) {
    var output rds.DescribeOrderableDBInstanceOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribePendingMaintenanceActionsResult struct {
	Result workflow.Future
}

func (r *RdsDescribePendingMaintenanceActionsResult) Get(ctx workflow.Context) (*rds.DescribePendingMaintenanceActionsOutput, error) {
    var output rds.DescribePendingMaintenanceActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeReservedDBInstancesResult struct {
	Result workflow.Future
}

func (r *RdsDescribeReservedDBInstancesResult) Get(ctx workflow.Context) (*rds.DescribeReservedDBInstancesOutput, error) {
    var output rds.DescribeReservedDBInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeReservedDBInstancesOfferingsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeReservedDBInstancesOfferingsResult) Get(ctx workflow.Context) (*rds.DescribeReservedDBInstancesOfferingsOutput, error) {
    var output rds.DescribeReservedDBInstancesOfferingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeSourceRegionsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeSourceRegionsResult) Get(ctx workflow.Context) (*rds.DescribeSourceRegionsOutput, error) {
    var output rds.DescribeSourceRegionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDescribeValidDBInstanceModificationsResult struct {
	Result workflow.Future
}

func (r *RdsDescribeValidDBInstanceModificationsResult) Get(ctx workflow.Context) (*rds.DescribeValidDBInstanceModificationsOutput, error) {
    var output rds.DescribeValidDBInstanceModificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsDownloadDBLogFilePortionResult struct {
	Result workflow.Future
}

func (r *RdsDownloadDBLogFilePortionResult) Get(ctx workflow.Context) (*rds.DownloadDBLogFilePortionOutput, error) {
    var output rds.DownloadDBLogFilePortionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsFailoverDBClusterResult struct {
	Result workflow.Future
}

func (r *RdsFailoverDBClusterResult) Get(ctx workflow.Context) (*rds.FailoverDBClusterOutput, error) {
    var output rds.FailoverDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsImportInstallationMediaResult struct {
	Result workflow.Future
}

func (r *RdsImportInstallationMediaResult) Get(ctx workflow.Context) (*rds.ImportInstallationMediaOutput, error) {
    var output rds.ImportInstallationMediaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *RdsListTagsForResourceResult) Get(ctx workflow.Context) (*rds.ListTagsForResourceOutput, error) {
    var output rds.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyCertificatesResult struct {
	Result workflow.Future
}

func (r *RdsModifyCertificatesResult) Get(ctx workflow.Context) (*rds.ModifyCertificatesOutput, error) {
    var output rds.ModifyCertificatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyCurrentDBClusterCapacityResult struct {
	Result workflow.Future
}

func (r *RdsModifyCurrentDBClusterCapacityResult) Get(ctx workflow.Context) (*rds.ModifyCurrentDBClusterCapacityOutput, error) {
    var output rds.ModifyCurrentDBClusterCapacityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyDBClusterResult struct {
	Result workflow.Future
}

func (r *RdsModifyDBClusterResult) Get(ctx workflow.Context) (*rds.ModifyDBClusterOutput, error) {
    var output rds.ModifyDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyDBClusterEndpointResult struct {
	Result workflow.Future
}

func (r *RdsModifyDBClusterEndpointResult) Get(ctx workflow.Context) (*rds.ModifyDBClusterEndpointOutput, error) {
    var output rds.ModifyDBClusterEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *RdsModifyDBClusterParameterGroupResult) Get(ctx workflow.Context) (*rds.DBClusterParameterGroupNameMessage, error) {
    var output rds.DBClusterParameterGroupNameMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyDBClusterSnapshotAttributeResult struct {
	Result workflow.Future
}

func (r *RdsModifyDBClusterSnapshotAttributeResult) Get(ctx workflow.Context) (*rds.ModifyDBClusterSnapshotAttributeOutput, error) {
    var output rds.ModifyDBClusterSnapshotAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyDBInstanceResult struct {
	Result workflow.Future
}

func (r *RdsModifyDBInstanceResult) Get(ctx workflow.Context) (*rds.ModifyDBInstanceOutput, error) {
    var output rds.ModifyDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyDBParameterGroupResult struct {
	Result workflow.Future
}

func (r *RdsModifyDBParameterGroupResult) Get(ctx workflow.Context) (*rds.DBParameterGroupNameMessage, error) {
    var output rds.DBParameterGroupNameMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyDBProxyResult struct {
	Result workflow.Future
}

func (r *RdsModifyDBProxyResult) Get(ctx workflow.Context) (*rds.ModifyDBProxyOutput, error) {
    var output rds.ModifyDBProxyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyDBProxyTargetGroupResult struct {
	Result workflow.Future
}

func (r *RdsModifyDBProxyTargetGroupResult) Get(ctx workflow.Context) (*rds.ModifyDBProxyTargetGroupOutput, error) {
    var output rds.ModifyDBProxyTargetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyDBSnapshotResult struct {
	Result workflow.Future
}

func (r *RdsModifyDBSnapshotResult) Get(ctx workflow.Context) (*rds.ModifyDBSnapshotOutput, error) {
    var output rds.ModifyDBSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyDBSnapshotAttributeResult struct {
	Result workflow.Future
}

func (r *RdsModifyDBSnapshotAttributeResult) Get(ctx workflow.Context) (*rds.ModifyDBSnapshotAttributeOutput, error) {
    var output rds.ModifyDBSnapshotAttributeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyDBSubnetGroupResult struct {
	Result workflow.Future
}

func (r *RdsModifyDBSubnetGroupResult) Get(ctx workflow.Context) (*rds.ModifyDBSubnetGroupOutput, error) {
    var output rds.ModifyDBSubnetGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyEventSubscriptionResult struct {
	Result workflow.Future
}

func (r *RdsModifyEventSubscriptionResult) Get(ctx workflow.Context) (*rds.ModifyEventSubscriptionOutput, error) {
    var output rds.ModifyEventSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyGlobalClusterResult struct {
	Result workflow.Future
}

func (r *RdsModifyGlobalClusterResult) Get(ctx workflow.Context) (*rds.ModifyGlobalClusterOutput, error) {
    var output rds.ModifyGlobalClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsModifyOptionGroupResult struct {
	Result workflow.Future
}

func (r *RdsModifyOptionGroupResult) Get(ctx workflow.Context) (*rds.ModifyOptionGroupOutput, error) {
    var output rds.ModifyOptionGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsPromoteReadReplicaResult struct {
	Result workflow.Future
}

func (r *RdsPromoteReadReplicaResult) Get(ctx workflow.Context) (*rds.PromoteReadReplicaOutput, error) {
    var output rds.PromoteReadReplicaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsPromoteReadReplicaDBClusterResult struct {
	Result workflow.Future
}

func (r *RdsPromoteReadReplicaDBClusterResult) Get(ctx workflow.Context) (*rds.PromoteReadReplicaDBClusterOutput, error) {
    var output rds.PromoteReadReplicaDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsPurchaseReservedDBInstancesOfferingResult struct {
	Result workflow.Future
}

func (r *RdsPurchaseReservedDBInstancesOfferingResult) Get(ctx workflow.Context) (*rds.PurchaseReservedDBInstancesOfferingOutput, error) {
    var output rds.PurchaseReservedDBInstancesOfferingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRebootDBInstanceResult struct {
	Result workflow.Future
}

func (r *RdsRebootDBInstanceResult) Get(ctx workflow.Context) (*rds.RebootDBInstanceOutput, error) {
    var output rds.RebootDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRegisterDBProxyTargetsResult struct {
	Result workflow.Future
}

func (r *RdsRegisterDBProxyTargetsResult) Get(ctx workflow.Context) (*rds.RegisterDBProxyTargetsOutput, error) {
    var output rds.RegisterDBProxyTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRemoveFromGlobalClusterResult struct {
	Result workflow.Future
}

func (r *RdsRemoveFromGlobalClusterResult) Get(ctx workflow.Context) (*rds.RemoveFromGlobalClusterOutput, error) {
    var output rds.RemoveFromGlobalClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRemoveRoleFromDBClusterResult struct {
	Result workflow.Future
}

func (r *RdsRemoveRoleFromDBClusterResult) Get(ctx workflow.Context) (*rds.RemoveRoleFromDBClusterOutput, error) {
    var output rds.RemoveRoleFromDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRemoveRoleFromDBInstanceResult struct {
	Result workflow.Future
}

func (r *RdsRemoveRoleFromDBInstanceResult) Get(ctx workflow.Context) (*rds.RemoveRoleFromDBInstanceOutput, error) {
    var output rds.RemoveRoleFromDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRemoveSourceIdentifierFromSubscriptionResult struct {
	Result workflow.Future
}

func (r *RdsRemoveSourceIdentifierFromSubscriptionResult) Get(ctx workflow.Context) (*rds.RemoveSourceIdentifierFromSubscriptionOutput, error) {
    var output rds.RemoveSourceIdentifierFromSubscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRemoveTagsFromResourceResult struct {
	Result workflow.Future
}

func (r *RdsRemoveTagsFromResourceResult) Get(ctx workflow.Context) (*rds.RemoveTagsFromResourceOutput, error) {
    var output rds.RemoveTagsFromResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsResetDBClusterParameterGroupResult struct {
	Result workflow.Future
}

func (r *RdsResetDBClusterParameterGroupResult) Get(ctx workflow.Context) (*rds.DBClusterParameterGroupNameMessage, error) {
    var output rds.DBClusterParameterGroupNameMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsResetDBParameterGroupResult struct {
	Result workflow.Future
}

func (r *RdsResetDBParameterGroupResult) Get(ctx workflow.Context) (*rds.DBParameterGroupNameMessage, error) {
    var output rds.DBParameterGroupNameMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRestoreDBClusterFromS3Result struct {
	Result workflow.Future
}

func (r *RdsRestoreDBClusterFromS3Result) Get(ctx workflow.Context) (*rds.RestoreDBClusterFromS3Output, error) {
    var output rds.RestoreDBClusterFromS3Output
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRestoreDBClusterFromSnapshotResult struct {
	Result workflow.Future
}

func (r *RdsRestoreDBClusterFromSnapshotResult) Get(ctx workflow.Context) (*rds.RestoreDBClusterFromSnapshotOutput, error) {
    var output rds.RestoreDBClusterFromSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRestoreDBClusterToPointInTimeResult struct {
	Result workflow.Future
}

func (r *RdsRestoreDBClusterToPointInTimeResult) Get(ctx workflow.Context) (*rds.RestoreDBClusterToPointInTimeOutput, error) {
    var output rds.RestoreDBClusterToPointInTimeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRestoreDBInstanceFromDBSnapshotResult struct {
	Result workflow.Future
}

func (r *RdsRestoreDBInstanceFromDBSnapshotResult) Get(ctx workflow.Context) (*rds.RestoreDBInstanceFromDBSnapshotOutput, error) {
    var output rds.RestoreDBInstanceFromDBSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRestoreDBInstanceFromS3Result struct {
	Result workflow.Future
}

func (r *RdsRestoreDBInstanceFromS3Result) Get(ctx workflow.Context) (*rds.RestoreDBInstanceFromS3Output, error) {
    var output rds.RestoreDBInstanceFromS3Output
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRestoreDBInstanceToPointInTimeResult struct {
	Result workflow.Future
}

func (r *RdsRestoreDBInstanceToPointInTimeResult) Get(ctx workflow.Context) (*rds.RestoreDBInstanceToPointInTimeOutput, error) {
    var output rds.RestoreDBInstanceToPointInTimeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsRevokeDBSecurityGroupIngressResult struct {
	Result workflow.Future
}

func (r *RdsRevokeDBSecurityGroupIngressResult) Get(ctx workflow.Context) (*rds.RevokeDBSecurityGroupIngressOutput, error) {
    var output rds.RevokeDBSecurityGroupIngressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsStartActivityStreamResult struct {
	Result workflow.Future
}

func (r *RdsStartActivityStreamResult) Get(ctx workflow.Context) (*rds.StartActivityStreamOutput, error) {
    var output rds.StartActivityStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsStartDBClusterResult struct {
	Result workflow.Future
}

func (r *RdsStartDBClusterResult) Get(ctx workflow.Context) (*rds.StartDBClusterOutput, error) {
    var output rds.StartDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsStartDBInstanceResult struct {
	Result workflow.Future
}

func (r *RdsStartDBInstanceResult) Get(ctx workflow.Context) (*rds.StartDBInstanceOutput, error) {
    var output rds.StartDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsStartExportTaskResult struct {
	Result workflow.Future
}

func (r *RdsStartExportTaskResult) Get(ctx workflow.Context) (*rds.StartExportTaskOutput, error) {
    var output rds.StartExportTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsStopActivityStreamResult struct {
	Result workflow.Future
}

func (r *RdsStopActivityStreamResult) Get(ctx workflow.Context) (*rds.StopActivityStreamOutput, error) {
    var output rds.StopActivityStreamOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsStopDBClusterResult struct {
	Result workflow.Future
}

func (r *RdsStopDBClusterResult) Get(ctx workflow.Context) (*rds.StopDBClusterOutput, error) {
    var output rds.StopDBClusterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type RdsStopDBInstanceResult struct {
	Result workflow.Future
}

func (r *RdsStopDBInstanceResult) Get(ctx workflow.Context) (*rds.StopDBInstanceOutput, error) {
    var output rds.StopDBInstanceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type RDSStub struct {
    activities RDSClient
}

func NewRDSStub() RDSClient {
    return &RDSStub{}
}

func (a *RDSStub) AddRoleToDBCluster(ctx workflow.Context, input *rds.AddRoleToDBClusterInput) (*rds.AddRoleToDBClusterOutput, error) {
    var output rds.AddRoleToDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddRoleToDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) AddRoleToDBInstance(ctx workflow.Context, input *rds.AddRoleToDBInstanceInput) (*rds.AddRoleToDBInstanceOutput, error) {
    var output rds.AddRoleToDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddRoleToDBInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) AddSourceIdentifierToSubscription(ctx workflow.Context, input *rds.AddSourceIdentifierToSubscriptionInput) (*rds.AddSourceIdentifierToSubscriptionOutput, error) {
    var output rds.AddSourceIdentifierToSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddSourceIdentifierToSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) AddTagsToResource(ctx workflow.Context, input *rds.AddTagsToResourceInput) (*rds.AddTagsToResourceOutput, error) {
    var output rds.AddTagsToResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTagsToResource, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ApplyPendingMaintenanceAction(ctx workflow.Context, input *rds.ApplyPendingMaintenanceActionInput) (*rds.ApplyPendingMaintenanceActionOutput, error) {
    var output rds.ApplyPendingMaintenanceActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ApplyPendingMaintenanceAction, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) AuthorizeDBSecurityGroupIngress(ctx workflow.Context, input *rds.AuthorizeDBSecurityGroupIngressInput) (*rds.AuthorizeDBSecurityGroupIngressOutput, error) {
    var output rds.AuthorizeDBSecurityGroupIngressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AuthorizeDBSecurityGroupIngress, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) BacktrackDBCluster(ctx workflow.Context, input *rds.BacktrackDBClusterInput) (*rds.BacktrackDBClusterOutput, error) {
    var output rds.BacktrackDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BacktrackDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CancelExportTask(ctx workflow.Context, input *rds.CancelExportTaskInput) (*rds.CancelExportTaskOutput, error) {
    var output rds.CancelExportTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelExportTask, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CopyDBClusterParameterGroup(ctx workflow.Context, input *rds.CopyDBClusterParameterGroupInput) (*rds.CopyDBClusterParameterGroupOutput, error) {
    var output rds.CopyDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyDBClusterParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CopyDBClusterSnapshot(ctx workflow.Context, input *rds.CopyDBClusterSnapshotInput) (*rds.CopyDBClusterSnapshotOutput, error) {
    var output rds.CopyDBClusterSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyDBClusterSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CopyDBParameterGroup(ctx workflow.Context, input *rds.CopyDBParameterGroupInput) (*rds.CopyDBParameterGroupOutput, error) {
    var output rds.CopyDBParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyDBParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CopyDBSnapshot(ctx workflow.Context, input *rds.CopyDBSnapshotInput) (*rds.CopyDBSnapshotOutput, error) {
    var output rds.CopyDBSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyDBSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CopyOptionGroup(ctx workflow.Context, input *rds.CopyOptionGroupInput) (*rds.CopyOptionGroupOutput, error) {
    var output rds.CopyOptionGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyOptionGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateCustomAvailabilityZone(ctx workflow.Context, input *rds.CreateCustomAvailabilityZoneInput) (*rds.CreateCustomAvailabilityZoneOutput, error) {
    var output rds.CreateCustomAvailabilityZoneOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCustomAvailabilityZone, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateDBCluster(ctx workflow.Context, input *rds.CreateDBClusterInput) (*rds.CreateDBClusterOutput, error) {
    var output rds.CreateDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateDBClusterEndpoint(ctx workflow.Context, input *rds.CreateDBClusterEndpointInput) (*rds.CreateDBClusterEndpointOutput, error) {
    var output rds.CreateDBClusterEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBClusterEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateDBClusterParameterGroup(ctx workflow.Context, input *rds.CreateDBClusterParameterGroupInput) (*rds.CreateDBClusterParameterGroupOutput, error) {
    var output rds.CreateDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBClusterParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateDBClusterSnapshot(ctx workflow.Context, input *rds.CreateDBClusterSnapshotInput) (*rds.CreateDBClusterSnapshotOutput, error) {
    var output rds.CreateDBClusterSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBClusterSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateDBInstance(ctx workflow.Context, input *rds.CreateDBInstanceInput) (*rds.CreateDBInstanceOutput, error) {
    var output rds.CreateDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateDBInstanceReadReplica(ctx workflow.Context, input *rds.CreateDBInstanceReadReplicaInput) (*rds.CreateDBInstanceReadReplicaOutput, error) {
    var output rds.CreateDBInstanceReadReplicaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBInstanceReadReplica, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateDBParameterGroup(ctx workflow.Context, input *rds.CreateDBParameterGroupInput) (*rds.CreateDBParameterGroupOutput, error) {
    var output rds.CreateDBParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateDBProxy(ctx workflow.Context, input *rds.CreateDBProxyInput) (*rds.CreateDBProxyOutput, error) {
    var output rds.CreateDBProxyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBProxy, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateDBSecurityGroup(ctx workflow.Context, input *rds.CreateDBSecurityGroupInput) (*rds.CreateDBSecurityGroupOutput, error) {
    var output rds.CreateDBSecurityGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBSecurityGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateDBSnapshot(ctx workflow.Context, input *rds.CreateDBSnapshotInput) (*rds.CreateDBSnapshotOutput, error) {
    var output rds.CreateDBSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateDBSubnetGroup(ctx workflow.Context, input *rds.CreateDBSubnetGroupInput) (*rds.CreateDBSubnetGroupOutput, error) {
    var output rds.CreateDBSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDBSubnetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateEventSubscription(ctx workflow.Context, input *rds.CreateEventSubscriptionInput) (*rds.CreateEventSubscriptionOutput, error) {
    var output rds.CreateEventSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEventSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateGlobalCluster(ctx workflow.Context, input *rds.CreateGlobalClusterInput) (*rds.CreateGlobalClusterOutput, error) {
    var output rds.CreateGlobalClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGlobalCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) CreateOptionGroup(ctx workflow.Context, input *rds.CreateOptionGroupInput) (*rds.CreateOptionGroupOutput, error) {
    var output rds.CreateOptionGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateOptionGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteCustomAvailabilityZone(ctx workflow.Context, input *rds.DeleteCustomAvailabilityZoneInput) (*rds.DeleteCustomAvailabilityZoneOutput, error) {
    var output rds.DeleteCustomAvailabilityZoneOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCustomAvailabilityZone, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteDBCluster(ctx workflow.Context, input *rds.DeleteDBClusterInput) (*rds.DeleteDBClusterOutput, error) {
    var output rds.DeleteDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteDBClusterEndpoint(ctx workflow.Context, input *rds.DeleteDBClusterEndpointInput) (*rds.DeleteDBClusterEndpointOutput, error) {
    var output rds.DeleteDBClusterEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBClusterEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteDBClusterParameterGroup(ctx workflow.Context, input *rds.DeleteDBClusterParameterGroupInput) (*rds.DeleteDBClusterParameterGroupOutput, error) {
    var output rds.DeleteDBClusterParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBClusterParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteDBClusterSnapshot(ctx workflow.Context, input *rds.DeleteDBClusterSnapshotInput) (*rds.DeleteDBClusterSnapshotOutput, error) {
    var output rds.DeleteDBClusterSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBClusterSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteDBInstance(ctx workflow.Context, input *rds.DeleteDBInstanceInput) (*rds.DeleteDBInstanceOutput, error) {
    var output rds.DeleteDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteDBInstanceAutomatedBackup(ctx workflow.Context, input *rds.DeleteDBInstanceAutomatedBackupInput) (*rds.DeleteDBInstanceAutomatedBackupOutput, error) {
    var output rds.DeleteDBInstanceAutomatedBackupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBInstanceAutomatedBackup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteDBParameterGroup(ctx workflow.Context, input *rds.DeleteDBParameterGroupInput) (*rds.DeleteDBParameterGroupOutput, error) {
    var output rds.DeleteDBParameterGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteDBProxy(ctx workflow.Context, input *rds.DeleteDBProxyInput) (*rds.DeleteDBProxyOutput, error) {
    var output rds.DeleteDBProxyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBProxy, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteDBSecurityGroup(ctx workflow.Context, input *rds.DeleteDBSecurityGroupInput) (*rds.DeleteDBSecurityGroupOutput, error) {
    var output rds.DeleteDBSecurityGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBSecurityGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteDBSnapshot(ctx workflow.Context, input *rds.DeleteDBSnapshotInput) (*rds.DeleteDBSnapshotOutput, error) {
    var output rds.DeleteDBSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteDBSubnetGroup(ctx workflow.Context, input *rds.DeleteDBSubnetGroupInput) (*rds.DeleteDBSubnetGroupOutput, error) {
    var output rds.DeleteDBSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDBSubnetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteEventSubscription(ctx workflow.Context, input *rds.DeleteEventSubscriptionInput) (*rds.DeleteEventSubscriptionOutput, error) {
    var output rds.DeleteEventSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEventSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteGlobalCluster(ctx workflow.Context, input *rds.DeleteGlobalClusterInput) (*rds.DeleteGlobalClusterOutput, error) {
    var output rds.DeleteGlobalClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGlobalCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteInstallationMedia(ctx workflow.Context, input *rds.DeleteInstallationMediaInput) (*rds.DeleteInstallationMediaOutput, error) {
    var output rds.DeleteInstallationMediaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInstallationMedia, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeleteOptionGroup(ctx workflow.Context, input *rds.DeleteOptionGroupInput) (*rds.DeleteOptionGroupOutput, error) {
    var output rds.DeleteOptionGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteOptionGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DeregisterDBProxyTargets(ctx workflow.Context, input *rds.DeregisterDBProxyTargetsInput) (*rds.DeregisterDBProxyTargetsOutput, error) {
    var output rds.DeregisterDBProxyTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeregisterDBProxyTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeAccountAttributes(ctx workflow.Context, input *rds.DescribeAccountAttributesInput) (*rds.DescribeAccountAttributesOutput, error) {
    var output rds.DescribeAccountAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeCertificates(ctx workflow.Context, input *rds.DescribeCertificatesInput) (*rds.DescribeCertificatesOutput, error) {
    var output rds.DescribeCertificatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCertificates, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeCustomAvailabilityZones(ctx workflow.Context, input *rds.DescribeCustomAvailabilityZonesInput) (*rds.DescribeCustomAvailabilityZonesOutput, error) {
    var output rds.DescribeCustomAvailabilityZonesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCustomAvailabilityZones, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBClusterBacktracks(ctx workflow.Context, input *rds.DescribeDBClusterBacktracksInput) (*rds.DescribeDBClusterBacktracksOutput, error) {
    var output rds.DescribeDBClusterBacktracksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterBacktracks, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBClusterEndpoints(ctx workflow.Context, input *rds.DescribeDBClusterEndpointsInput) (*rds.DescribeDBClusterEndpointsOutput, error) {
    var output rds.DescribeDBClusterEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBClusterParameterGroups(ctx workflow.Context, input *rds.DescribeDBClusterParameterGroupsInput) (*rds.DescribeDBClusterParameterGroupsOutput, error) {
    var output rds.DescribeDBClusterParameterGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterParameterGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBClusterParameters(ctx workflow.Context, input *rds.DescribeDBClusterParametersInput) (*rds.DescribeDBClusterParametersOutput, error) {
    var output rds.DescribeDBClusterParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBClusterSnapshotAttributes(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotAttributesInput) (*rds.DescribeDBClusterSnapshotAttributesOutput, error) {
    var output rds.DescribeDBClusterSnapshotAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterSnapshotAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBClusterSnapshots(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) (*rds.DescribeDBClusterSnapshotsOutput, error) {
    var output rds.DescribeDBClusterSnapshotsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusterSnapshots, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBClusters(ctx workflow.Context, input *rds.DescribeDBClustersInput) (*rds.DescribeDBClustersOutput, error) {
    var output rds.DescribeDBClustersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBClusters, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBEngineVersions(ctx workflow.Context, input *rds.DescribeDBEngineVersionsInput) (*rds.DescribeDBEngineVersionsOutput, error) {
    var output rds.DescribeDBEngineVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBEngineVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBInstanceAutomatedBackups(ctx workflow.Context, input *rds.DescribeDBInstanceAutomatedBackupsInput) (*rds.DescribeDBInstanceAutomatedBackupsOutput, error) {
    var output rds.DescribeDBInstanceAutomatedBackupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBInstanceAutomatedBackups, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBInstances(ctx workflow.Context, input *rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error) {
    var output rds.DescribeDBInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBLogFiles(ctx workflow.Context, input *rds.DescribeDBLogFilesInput) (*rds.DescribeDBLogFilesOutput, error) {
    var output rds.DescribeDBLogFilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBLogFiles, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBParameterGroups(ctx workflow.Context, input *rds.DescribeDBParameterGroupsInput) (*rds.DescribeDBParameterGroupsOutput, error) {
    var output rds.DescribeDBParameterGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBParameterGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBParameters(ctx workflow.Context, input *rds.DescribeDBParametersInput) (*rds.DescribeDBParametersOutput, error) {
    var output rds.DescribeDBParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBProxies(ctx workflow.Context, input *rds.DescribeDBProxiesInput) (*rds.DescribeDBProxiesOutput, error) {
    var output rds.DescribeDBProxiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBProxies, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBProxyTargetGroups(ctx workflow.Context, input *rds.DescribeDBProxyTargetGroupsInput) (*rds.DescribeDBProxyTargetGroupsOutput, error) {
    var output rds.DescribeDBProxyTargetGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBProxyTargetGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBProxyTargets(ctx workflow.Context, input *rds.DescribeDBProxyTargetsInput) (*rds.DescribeDBProxyTargetsOutput, error) {
    var output rds.DescribeDBProxyTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBProxyTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBSecurityGroups(ctx workflow.Context, input *rds.DescribeDBSecurityGroupsInput) (*rds.DescribeDBSecurityGroupsOutput, error) {
    var output rds.DescribeDBSecurityGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBSecurityGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBSnapshotAttributes(ctx workflow.Context, input *rds.DescribeDBSnapshotAttributesInput) (*rds.DescribeDBSnapshotAttributesOutput, error) {
    var output rds.DescribeDBSnapshotAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBSnapshotAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBSnapshots(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) (*rds.DescribeDBSnapshotsOutput, error) {
    var output rds.DescribeDBSnapshotsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBSnapshots, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeDBSubnetGroups(ctx workflow.Context, input *rds.DescribeDBSubnetGroupsInput) (*rds.DescribeDBSubnetGroupsOutput, error) {
    var output rds.DescribeDBSubnetGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDBSubnetGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeEngineDefaultClusterParameters(ctx workflow.Context, input *rds.DescribeEngineDefaultClusterParametersInput) (*rds.DescribeEngineDefaultClusterParametersOutput, error) {
    var output rds.DescribeEngineDefaultClusterParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEngineDefaultClusterParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeEngineDefaultParameters(ctx workflow.Context, input *rds.DescribeEngineDefaultParametersInput) (*rds.DescribeEngineDefaultParametersOutput, error) {
    var output rds.DescribeEngineDefaultParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEngineDefaultParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeEventCategories(ctx workflow.Context, input *rds.DescribeEventCategoriesInput) (*rds.DescribeEventCategoriesOutput, error) {
    var output rds.DescribeEventCategoriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventCategories, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeEventSubscriptions(ctx workflow.Context, input *rds.DescribeEventSubscriptionsInput) (*rds.DescribeEventSubscriptionsOutput, error) {
    var output rds.DescribeEventSubscriptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventSubscriptions, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeEvents(ctx workflow.Context, input *rds.DescribeEventsInput) (*rds.DescribeEventsOutput, error) {
    var output rds.DescribeEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeExportTasks(ctx workflow.Context, input *rds.DescribeExportTasksInput) (*rds.DescribeExportTasksOutput, error) {
    var output rds.DescribeExportTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeExportTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeGlobalClusters(ctx workflow.Context, input *rds.DescribeGlobalClustersInput) (*rds.DescribeGlobalClustersOutput, error) {
    var output rds.DescribeGlobalClustersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeGlobalClusters, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeInstallationMedia(ctx workflow.Context, input *rds.DescribeInstallationMediaInput) (*rds.DescribeInstallationMediaOutput, error) {
    var output rds.DescribeInstallationMediaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstallationMedia, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeOptionGroupOptions(ctx workflow.Context, input *rds.DescribeOptionGroupOptionsInput) (*rds.DescribeOptionGroupOptionsOutput, error) {
    var output rds.DescribeOptionGroupOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOptionGroupOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeOptionGroups(ctx workflow.Context, input *rds.DescribeOptionGroupsInput) (*rds.DescribeOptionGroupsOutput, error) {
    var output rds.DescribeOptionGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOptionGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeOrderableDBInstanceOptions(ctx workflow.Context, input *rds.DescribeOrderableDBInstanceOptionsInput) (*rds.DescribeOrderableDBInstanceOptionsOutput, error) {
    var output rds.DescribeOrderableDBInstanceOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOrderableDBInstanceOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribePendingMaintenanceActions(ctx workflow.Context, input *rds.DescribePendingMaintenanceActionsInput) (*rds.DescribePendingMaintenanceActionsOutput, error) {
    var output rds.DescribePendingMaintenanceActionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePendingMaintenanceActions, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeReservedDBInstances(ctx workflow.Context, input *rds.DescribeReservedDBInstancesInput) (*rds.DescribeReservedDBInstancesOutput, error) {
    var output rds.DescribeReservedDBInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedDBInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeReservedDBInstancesOfferings(ctx workflow.Context, input *rds.DescribeReservedDBInstancesOfferingsInput) (*rds.DescribeReservedDBInstancesOfferingsOutput, error) {
    var output rds.DescribeReservedDBInstancesOfferingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedDBInstancesOfferings, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeSourceRegions(ctx workflow.Context, input *rds.DescribeSourceRegionsInput) (*rds.DescribeSourceRegionsOutput, error) {
    var output rds.DescribeSourceRegionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSourceRegions, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DescribeValidDBInstanceModifications(ctx workflow.Context, input *rds.DescribeValidDBInstanceModificationsInput) (*rds.DescribeValidDBInstanceModificationsOutput, error) {
    var output rds.DescribeValidDBInstanceModificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeValidDBInstanceModifications, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) DownloadDBLogFilePortion(ctx workflow.Context, input *rds.DownloadDBLogFilePortionInput) (*rds.DownloadDBLogFilePortionOutput, error) {
    var output rds.DownloadDBLogFilePortionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DownloadDBLogFilePortion, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) FailoverDBCluster(ctx workflow.Context, input *rds.FailoverDBClusterInput) (*rds.FailoverDBClusterOutput, error) {
    var output rds.FailoverDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.FailoverDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ImportInstallationMedia(ctx workflow.Context, input *rds.ImportInstallationMediaInput) (*rds.ImportInstallationMediaOutput, error) {
    var output rds.ImportInstallationMediaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ImportInstallationMedia, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ListTagsForResource(ctx workflow.Context, input *rds.ListTagsForResourceInput) (*rds.ListTagsForResourceOutput, error) {
    var output rds.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyCertificates(ctx workflow.Context, input *rds.ModifyCertificatesInput) (*rds.ModifyCertificatesOutput, error) {
    var output rds.ModifyCertificatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyCertificates, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyCurrentDBClusterCapacity(ctx workflow.Context, input *rds.ModifyCurrentDBClusterCapacityInput) (*rds.ModifyCurrentDBClusterCapacityOutput, error) {
    var output rds.ModifyCurrentDBClusterCapacityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyCurrentDBClusterCapacity, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyDBCluster(ctx workflow.Context, input *rds.ModifyDBClusterInput) (*rds.ModifyDBClusterOutput, error) {
    var output rds.ModifyDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyDBClusterEndpoint(ctx workflow.Context, input *rds.ModifyDBClusterEndpointInput) (*rds.ModifyDBClusterEndpointOutput, error) {
    var output rds.ModifyDBClusterEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBClusterEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyDBClusterParameterGroup(ctx workflow.Context, input *rds.ModifyDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error) {
    var output rds.DBClusterParameterGroupNameMessage
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBClusterParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyDBClusterSnapshotAttribute(ctx workflow.Context, input *rds.ModifyDBClusterSnapshotAttributeInput) (*rds.ModifyDBClusterSnapshotAttributeOutput, error) {
    var output rds.ModifyDBClusterSnapshotAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBClusterSnapshotAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyDBInstance(ctx workflow.Context, input *rds.ModifyDBInstanceInput) (*rds.ModifyDBInstanceOutput, error) {
    var output rds.ModifyDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyDBParameterGroup(ctx workflow.Context, input *rds.ModifyDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error) {
    var output rds.DBParameterGroupNameMessage
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyDBProxy(ctx workflow.Context, input *rds.ModifyDBProxyInput) (*rds.ModifyDBProxyOutput, error) {
    var output rds.ModifyDBProxyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBProxy, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyDBProxyTargetGroup(ctx workflow.Context, input *rds.ModifyDBProxyTargetGroupInput) (*rds.ModifyDBProxyTargetGroupOutput, error) {
    var output rds.ModifyDBProxyTargetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBProxyTargetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyDBSnapshot(ctx workflow.Context, input *rds.ModifyDBSnapshotInput) (*rds.ModifyDBSnapshotOutput, error) {
    var output rds.ModifyDBSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyDBSnapshotAttribute(ctx workflow.Context, input *rds.ModifyDBSnapshotAttributeInput) (*rds.ModifyDBSnapshotAttributeOutput, error) {
    var output rds.ModifyDBSnapshotAttributeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBSnapshotAttribute, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyDBSubnetGroup(ctx workflow.Context, input *rds.ModifyDBSubnetGroupInput) (*rds.ModifyDBSubnetGroupOutput, error) {
    var output rds.ModifyDBSubnetGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyDBSubnetGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyEventSubscription(ctx workflow.Context, input *rds.ModifyEventSubscriptionInput) (*rds.ModifyEventSubscriptionOutput, error) {
    var output rds.ModifyEventSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyEventSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyGlobalCluster(ctx workflow.Context, input *rds.ModifyGlobalClusterInput) (*rds.ModifyGlobalClusterOutput, error) {
    var output rds.ModifyGlobalClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyGlobalCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ModifyOptionGroup(ctx workflow.Context, input *rds.ModifyOptionGroupInput) (*rds.ModifyOptionGroupOutput, error) {
    var output rds.ModifyOptionGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ModifyOptionGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) PromoteReadReplica(ctx workflow.Context, input *rds.PromoteReadReplicaInput) (*rds.PromoteReadReplicaOutput, error) {
    var output rds.PromoteReadReplicaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PromoteReadReplica, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) PromoteReadReplicaDBCluster(ctx workflow.Context, input *rds.PromoteReadReplicaDBClusterInput) (*rds.PromoteReadReplicaDBClusterOutput, error) {
    var output rds.PromoteReadReplicaDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PromoteReadReplicaDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) PurchaseReservedDBInstancesOffering(ctx workflow.Context, input *rds.PurchaseReservedDBInstancesOfferingInput) (*rds.PurchaseReservedDBInstancesOfferingOutput, error) {
    var output rds.PurchaseReservedDBInstancesOfferingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PurchaseReservedDBInstancesOffering, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RebootDBInstance(ctx workflow.Context, input *rds.RebootDBInstanceInput) (*rds.RebootDBInstanceOutput, error) {
    var output rds.RebootDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RebootDBInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RegisterDBProxyTargets(ctx workflow.Context, input *rds.RegisterDBProxyTargetsInput) (*rds.RegisterDBProxyTargetsOutput, error) {
    var output rds.RegisterDBProxyTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterDBProxyTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RemoveFromGlobalCluster(ctx workflow.Context, input *rds.RemoveFromGlobalClusterInput) (*rds.RemoveFromGlobalClusterOutput, error) {
    var output rds.RemoveFromGlobalClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveFromGlobalCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RemoveRoleFromDBCluster(ctx workflow.Context, input *rds.RemoveRoleFromDBClusterInput) (*rds.RemoveRoleFromDBClusterOutput, error) {
    var output rds.RemoveRoleFromDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveRoleFromDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RemoveRoleFromDBInstance(ctx workflow.Context, input *rds.RemoveRoleFromDBInstanceInput) (*rds.RemoveRoleFromDBInstanceOutput, error) {
    var output rds.RemoveRoleFromDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveRoleFromDBInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RemoveSourceIdentifierFromSubscription(ctx workflow.Context, input *rds.RemoveSourceIdentifierFromSubscriptionInput) (*rds.RemoveSourceIdentifierFromSubscriptionOutput, error) {
    var output rds.RemoveSourceIdentifierFromSubscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveSourceIdentifierFromSubscription, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RemoveTagsFromResource(ctx workflow.Context, input *rds.RemoveTagsFromResourceInput) (*rds.RemoveTagsFromResourceOutput, error) {
    var output rds.RemoveTagsFromResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTagsFromResource, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ResetDBClusterParameterGroup(ctx workflow.Context, input *rds.ResetDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error) {
    var output rds.DBClusterParameterGroupNameMessage
    err := workflow.ExecuteActivity(ctx, a.activities.ResetDBClusterParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) ResetDBParameterGroup(ctx workflow.Context, input *rds.ResetDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error) {
    var output rds.DBParameterGroupNameMessage
    err := workflow.ExecuteActivity(ctx, a.activities.ResetDBParameterGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RestoreDBClusterFromS3(ctx workflow.Context, input *rds.RestoreDBClusterFromS3Input) (*rds.RestoreDBClusterFromS3Output, error) {
    var output rds.RestoreDBClusterFromS3Output
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreDBClusterFromS3, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RestoreDBClusterFromSnapshot(ctx workflow.Context, input *rds.RestoreDBClusterFromSnapshotInput) (*rds.RestoreDBClusterFromSnapshotOutput, error) {
    var output rds.RestoreDBClusterFromSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreDBClusterFromSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RestoreDBClusterToPointInTime(ctx workflow.Context, input *rds.RestoreDBClusterToPointInTimeInput) (*rds.RestoreDBClusterToPointInTimeOutput, error) {
    var output rds.RestoreDBClusterToPointInTimeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreDBClusterToPointInTime, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RestoreDBInstanceFromDBSnapshot(ctx workflow.Context, input *rds.RestoreDBInstanceFromDBSnapshotInput) (*rds.RestoreDBInstanceFromDBSnapshotOutput, error) {
    var output rds.RestoreDBInstanceFromDBSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreDBInstanceFromDBSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RestoreDBInstanceFromS3(ctx workflow.Context, input *rds.RestoreDBInstanceFromS3Input) (*rds.RestoreDBInstanceFromS3Output, error) {
    var output rds.RestoreDBInstanceFromS3Output
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreDBInstanceFromS3, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RestoreDBInstanceToPointInTime(ctx workflow.Context, input *rds.RestoreDBInstanceToPointInTimeInput) (*rds.RestoreDBInstanceToPointInTimeOutput, error) {
    var output rds.RestoreDBInstanceToPointInTimeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestoreDBInstanceToPointInTime, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) RevokeDBSecurityGroupIngress(ctx workflow.Context, input *rds.RevokeDBSecurityGroupIngressInput) (*rds.RevokeDBSecurityGroupIngressOutput, error) {
    var output rds.RevokeDBSecurityGroupIngressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RevokeDBSecurityGroupIngress, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) StartActivityStream(ctx workflow.Context, input *rds.StartActivityStreamInput) (*rds.StartActivityStreamOutput, error) {
    var output rds.StartActivityStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartActivityStream, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) StartDBCluster(ctx workflow.Context, input *rds.StartDBClusterInput) (*rds.StartDBClusterOutput, error) {
    var output rds.StartDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) StartDBInstance(ctx workflow.Context, input *rds.StartDBInstanceInput) (*rds.StartDBInstanceOutput, error) {
    var output rds.StartDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartDBInstance, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) StartExportTask(ctx workflow.Context, input *rds.StartExportTaskInput) (*rds.StartExportTaskOutput, error) {
    var output rds.StartExportTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartExportTask, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) StopActivityStream(ctx workflow.Context, input *rds.StopActivityStreamInput) (*rds.StopActivityStreamOutput, error) {
    var output rds.StopActivityStreamOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopActivityStream, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) StopDBCluster(ctx workflow.Context, input *rds.StopDBClusterInput) (*rds.StopDBClusterOutput, error) {
    var output rds.StopDBClusterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopDBCluster, input).Get(ctx, &output)
    return &output, err
}

func (a *RDSStub) StopDBInstance(ctx workflow.Context, input *rds.StopDBInstanceInput) (*rds.StopDBInstanceOutput, error) {
    var output rds.StopDBInstanceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopDBInstance, input).Get(ctx, &output)
    return &output, err
}


func (a *RDSStub) WaitUntilDBClusterSnapshotAvailable(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) error {
    return a.client.WaitUntilDBClusterSnapshotAvailable(input)
}


func (a *RDSStub) WaitUntilDBClusterSnapshotDeleted(ctx workflow.Context, input *rds.DescribeDBClusterSnapshotsInput) error {
    return a.client.WaitUntilDBClusterSnapshotDeleted(input)
}


func (a *RDSStub) WaitUntilDBInstanceAvailable(ctx workflow.Context, input *rds.DescribeDBInstancesInput) error {
    return a.client.WaitUntilDBInstanceAvailable(input)
}


func (a *RDSStub) WaitUntilDBInstanceDeleted(ctx workflow.Context, input *rds.DescribeDBInstancesInput) error {
    return a.client.WaitUntilDBInstanceDeleted(input)
}


func (a *RDSStub) WaitUntilDBSnapshotAvailable(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) error {
    return a.client.WaitUntilDBSnapshotAvailable(input)
}


func (a *RDSStub) WaitUntilDBSnapshotDeleted(ctx workflow.Context, input *rds.DescribeDBSnapshotsInput) error {
    return a.client.WaitUntilDBSnapshotDeleted(input)
}
