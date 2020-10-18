// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package storagegateway

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/storagegateway"
	"github.com/aws/aws-sdk-go/service/storagegateway/storagegatewayiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client storagegatewayiface.StorageGatewayAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := storagegateway.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (storagegatewayiface.StorageGatewayAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return storagegateway.New(sess), nil
}

func (a *Activities) ActivateGateway(ctx context.Context, input *storagegateway.ActivateGatewayInput) (*storagegateway.ActivateGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ActivateGatewayWithContext(ctx, input)
}

func (a *Activities) AddCache(ctx context.Context, input *storagegateway.AddCacheInput) (*storagegateway.AddCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddCacheWithContext(ctx, input)
}

func (a *Activities) AddTagsToResource(ctx context.Context, input *storagegateway.AddTagsToResourceInput) (*storagegateway.AddTagsToResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddTagsToResourceWithContext(ctx, input)
}

func (a *Activities) AddUploadBuffer(ctx context.Context, input *storagegateway.AddUploadBufferInput) (*storagegateway.AddUploadBufferOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddUploadBufferWithContext(ctx, input)
}

func (a *Activities) AddWorkingStorage(ctx context.Context, input *storagegateway.AddWorkingStorageInput) (*storagegateway.AddWorkingStorageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddWorkingStorageWithContext(ctx, input)
}

func (a *Activities) AssignTapePool(ctx context.Context, input *storagegateway.AssignTapePoolInput) (*storagegateway.AssignTapePoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssignTapePoolWithContext(ctx, input)
}

func (a *Activities) AttachVolume(ctx context.Context, input *storagegateway.AttachVolumeInput) (*storagegateway.AttachVolumeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AttachVolumeWithContext(ctx, input)
}

func (a *Activities) CancelArchival(ctx context.Context, input *storagegateway.CancelArchivalInput) (*storagegateway.CancelArchivalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelArchivalWithContext(ctx, input)
}

func (a *Activities) CancelRetrieval(ctx context.Context, input *storagegateway.CancelRetrievalInput) (*storagegateway.CancelRetrievalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CancelRetrievalWithContext(ctx, input)
}

func (a *Activities) CreateCachediSCSIVolume(ctx context.Context, input *storagegateway.CreateCachediSCSIVolumeInput) (*storagegateway.CreateCachediSCSIVolumeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateCachediSCSIVolumeWithContext(ctx, input)
}

func (a *Activities) CreateNFSFileShare(ctx context.Context, input *storagegateway.CreateNFSFileShareInput) (*storagegateway.CreateNFSFileShareOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateNFSFileShareWithContext(ctx, input)
}

func (a *Activities) CreateSMBFileShare(ctx context.Context, input *storagegateway.CreateSMBFileShareInput) (*storagegateway.CreateSMBFileShareOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateSMBFileShareWithContext(ctx, input)
}

func (a *Activities) CreateSnapshot(ctx context.Context, input *storagegateway.CreateSnapshotInput) (*storagegateway.CreateSnapshotOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSnapshotWithContext(ctx, input)
}

func (a *Activities) CreateSnapshotFromVolumeRecoveryPoint(ctx context.Context, input *storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSnapshotFromVolumeRecoveryPointWithContext(ctx, input)
}

func (a *Activities) CreateStorediSCSIVolume(ctx context.Context, input *storagegateway.CreateStorediSCSIVolumeInput) (*storagegateway.CreateStorediSCSIVolumeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateStorediSCSIVolumeWithContext(ctx, input)
}

func (a *Activities) CreateTapePool(ctx context.Context, input *storagegateway.CreateTapePoolInput) (*storagegateway.CreateTapePoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTapePoolWithContext(ctx, input)
}

func (a *Activities) CreateTapeWithBarcode(ctx context.Context, input *storagegateway.CreateTapeWithBarcodeInput) (*storagegateway.CreateTapeWithBarcodeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTapeWithBarcodeWithContext(ctx, input)
}

func (a *Activities) CreateTapes(ctx context.Context, input *storagegateway.CreateTapesInput) (*storagegateway.CreateTapesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.CreateTapesWithContext(ctx, input)
}

func (a *Activities) DeleteAutomaticTapeCreationPolicy(ctx context.Context, input *storagegateway.DeleteAutomaticTapeCreationPolicyInput) (*storagegateway.DeleteAutomaticTapeCreationPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAutomaticTapeCreationPolicyWithContext(ctx, input)
}

func (a *Activities) DeleteBandwidthRateLimit(ctx context.Context, input *storagegateway.DeleteBandwidthRateLimitInput) (*storagegateway.DeleteBandwidthRateLimitOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBandwidthRateLimitWithContext(ctx, input)
}

func (a *Activities) DeleteChapCredentials(ctx context.Context, input *storagegateway.DeleteChapCredentialsInput) (*storagegateway.DeleteChapCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteChapCredentialsWithContext(ctx, input)
}

func (a *Activities) DeleteFileShare(ctx context.Context, input *storagegateway.DeleteFileShareInput) (*storagegateway.DeleteFileShareOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFileShareWithContext(ctx, input)
}

func (a *Activities) DeleteGateway(ctx context.Context, input *storagegateway.DeleteGatewayInput) (*storagegateway.DeleteGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteGatewayWithContext(ctx, input)
}

func (a *Activities) DeleteSnapshotSchedule(ctx context.Context, input *storagegateway.DeleteSnapshotScheduleInput) (*storagegateway.DeleteSnapshotScheduleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSnapshotScheduleWithContext(ctx, input)
}

func (a *Activities) DeleteTape(ctx context.Context, input *storagegateway.DeleteTapeInput) (*storagegateway.DeleteTapeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTapeWithContext(ctx, input)
}

func (a *Activities) DeleteTapeArchive(ctx context.Context, input *storagegateway.DeleteTapeArchiveInput) (*storagegateway.DeleteTapeArchiveOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTapeArchiveWithContext(ctx, input)
}

func (a *Activities) DeleteTapePool(ctx context.Context, input *storagegateway.DeleteTapePoolInput) (*storagegateway.DeleteTapePoolOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTapePoolWithContext(ctx, input)
}

func (a *Activities) DeleteVolume(ctx context.Context, input *storagegateway.DeleteVolumeInput) (*storagegateway.DeleteVolumeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVolumeWithContext(ctx, input)
}

func (a *Activities) DescribeAvailabilityMonitorTest(ctx context.Context, input *storagegateway.DescribeAvailabilityMonitorTestInput) (*storagegateway.DescribeAvailabilityMonitorTestOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAvailabilityMonitorTestWithContext(ctx, input)
}

func (a *Activities) DescribeBandwidthRateLimit(ctx context.Context, input *storagegateway.DescribeBandwidthRateLimitInput) (*storagegateway.DescribeBandwidthRateLimitOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeBandwidthRateLimitWithContext(ctx, input)
}

func (a *Activities) DescribeCache(ctx context.Context, input *storagegateway.DescribeCacheInput) (*storagegateway.DescribeCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCacheWithContext(ctx, input)
}

func (a *Activities) DescribeCachediSCSIVolumes(ctx context.Context, input *storagegateway.DescribeCachediSCSIVolumesInput) (*storagegateway.DescribeCachediSCSIVolumesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeCachediSCSIVolumesWithContext(ctx, input)
}

func (a *Activities) DescribeChapCredentials(ctx context.Context, input *storagegateway.DescribeChapCredentialsInput) (*storagegateway.DescribeChapCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeChapCredentialsWithContext(ctx, input)
}

func (a *Activities) DescribeGatewayInformation(ctx context.Context, input *storagegateway.DescribeGatewayInformationInput) (*storagegateway.DescribeGatewayInformationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeGatewayInformationWithContext(ctx, input)
}

func (a *Activities) DescribeMaintenanceStartTime(ctx context.Context, input *storagegateway.DescribeMaintenanceStartTimeInput) (*storagegateway.DescribeMaintenanceStartTimeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeMaintenanceStartTimeWithContext(ctx, input)
}

func (a *Activities) DescribeNFSFileShares(ctx context.Context, input *storagegateway.DescribeNFSFileSharesInput) (*storagegateway.DescribeNFSFileSharesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeNFSFileSharesWithContext(ctx, input)
}

func (a *Activities) DescribeSMBFileShares(ctx context.Context, input *storagegateway.DescribeSMBFileSharesInput) (*storagegateway.DescribeSMBFileSharesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSMBFileSharesWithContext(ctx, input)
}

func (a *Activities) DescribeSMBSettings(ctx context.Context, input *storagegateway.DescribeSMBSettingsInput) (*storagegateway.DescribeSMBSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSMBSettingsWithContext(ctx, input)
}

func (a *Activities) DescribeSnapshotSchedule(ctx context.Context, input *storagegateway.DescribeSnapshotScheduleInput) (*storagegateway.DescribeSnapshotScheduleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSnapshotScheduleWithContext(ctx, input)
}

func (a *Activities) DescribeStorediSCSIVolumes(ctx context.Context, input *storagegateway.DescribeStorediSCSIVolumesInput) (*storagegateway.DescribeStorediSCSIVolumesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeStorediSCSIVolumesWithContext(ctx, input)
}

func (a *Activities) DescribeTapeArchives(ctx context.Context, input *storagegateway.DescribeTapeArchivesInput) (*storagegateway.DescribeTapeArchivesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTapeArchivesWithContext(ctx, input)
}

func (a *Activities) DescribeTapeRecoveryPoints(ctx context.Context, input *storagegateway.DescribeTapeRecoveryPointsInput) (*storagegateway.DescribeTapeRecoveryPointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTapeRecoveryPointsWithContext(ctx, input)
}

func (a *Activities) DescribeTapes(ctx context.Context, input *storagegateway.DescribeTapesInput) (*storagegateway.DescribeTapesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTapesWithContext(ctx, input)
}

func (a *Activities) DescribeUploadBuffer(ctx context.Context, input *storagegateway.DescribeUploadBufferInput) (*storagegateway.DescribeUploadBufferOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUploadBufferWithContext(ctx, input)
}

func (a *Activities) DescribeVTLDevices(ctx context.Context, input *storagegateway.DescribeVTLDevicesInput) (*storagegateway.DescribeVTLDevicesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeVTLDevicesWithContext(ctx, input)
}

func (a *Activities) DescribeWorkingStorage(ctx context.Context, input *storagegateway.DescribeWorkingStorageInput) (*storagegateway.DescribeWorkingStorageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeWorkingStorageWithContext(ctx, input)
}

func (a *Activities) DetachVolume(ctx context.Context, input *storagegateway.DetachVolumeInput) (*storagegateway.DetachVolumeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetachVolumeWithContext(ctx, input)
}

func (a *Activities) DisableGateway(ctx context.Context, input *storagegateway.DisableGatewayInput) (*storagegateway.DisableGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisableGatewayWithContext(ctx, input)
}

func (a *Activities) JoinDomain(ctx context.Context, input *storagegateway.JoinDomainInput) (*storagegateway.JoinDomainOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.JoinDomainWithContext(ctx, input)
}

func (a *Activities) ListAutomaticTapeCreationPolicies(ctx context.Context, input *storagegateway.ListAutomaticTapeCreationPoliciesInput) (*storagegateway.ListAutomaticTapeCreationPoliciesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAutomaticTapeCreationPoliciesWithContext(ctx, input)
}

func (a *Activities) ListFileShares(ctx context.Context, input *storagegateway.ListFileSharesInput) (*storagegateway.ListFileSharesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFileSharesWithContext(ctx, input)
}

func (a *Activities) ListGateways(ctx context.Context, input *storagegateway.ListGatewaysInput) (*storagegateway.ListGatewaysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListGatewaysWithContext(ctx, input)
}

func (a *Activities) ListLocalDisks(ctx context.Context, input *storagegateway.ListLocalDisksInput) (*storagegateway.ListLocalDisksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLocalDisksWithContext(ctx, input)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *storagegateway.ListTagsForResourceInput) (*storagegateway.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Activities) ListTapePools(ctx context.Context, input *storagegateway.ListTapePoolsInput) (*storagegateway.ListTapePoolsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTapePoolsWithContext(ctx, input)
}

func (a *Activities) ListTapes(ctx context.Context, input *storagegateway.ListTapesInput) (*storagegateway.ListTapesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTapesWithContext(ctx, input)
}

func (a *Activities) ListVolumeInitiators(ctx context.Context, input *storagegateway.ListVolumeInitiatorsInput) (*storagegateway.ListVolumeInitiatorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVolumeInitiatorsWithContext(ctx, input)
}

func (a *Activities) ListVolumeRecoveryPoints(ctx context.Context, input *storagegateway.ListVolumeRecoveryPointsInput) (*storagegateway.ListVolumeRecoveryPointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVolumeRecoveryPointsWithContext(ctx, input)
}

func (a *Activities) ListVolumes(ctx context.Context, input *storagegateway.ListVolumesInput) (*storagegateway.ListVolumesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVolumesWithContext(ctx, input)
}

func (a *Activities) NotifyWhenUploaded(ctx context.Context, input *storagegateway.NotifyWhenUploadedInput) (*storagegateway.NotifyWhenUploadedOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.NotifyWhenUploadedWithContext(ctx, input)
}

func (a *Activities) RefreshCache(ctx context.Context, input *storagegateway.RefreshCacheInput) (*storagegateway.RefreshCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RefreshCacheWithContext(ctx, input)
}

func (a *Activities) RemoveTagsFromResource(ctx context.Context, input *storagegateway.RemoveTagsFromResourceInput) (*storagegateway.RemoveTagsFromResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *Activities) ResetCache(ctx context.Context, input *storagegateway.ResetCacheInput) (*storagegateway.ResetCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ResetCacheWithContext(ctx, input)
}

func (a *Activities) RetrieveTapeArchive(ctx context.Context, input *storagegateway.RetrieveTapeArchiveInput) (*storagegateway.RetrieveTapeArchiveOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RetrieveTapeArchiveWithContext(ctx, input)
}

func (a *Activities) RetrieveTapeRecoveryPoint(ctx context.Context, input *storagegateway.RetrieveTapeRecoveryPointInput) (*storagegateway.RetrieveTapeRecoveryPointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RetrieveTapeRecoveryPointWithContext(ctx, input)
}

func (a *Activities) SetLocalConsolePassword(ctx context.Context, input *storagegateway.SetLocalConsolePasswordInput) (*storagegateway.SetLocalConsolePasswordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetLocalConsolePasswordWithContext(ctx, input)
}

func (a *Activities) SetSMBGuestPassword(ctx context.Context, input *storagegateway.SetSMBGuestPasswordInput) (*storagegateway.SetSMBGuestPasswordOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetSMBGuestPasswordWithContext(ctx, input)
}

func (a *Activities) ShutdownGateway(ctx context.Context, input *storagegateway.ShutdownGatewayInput) (*storagegateway.ShutdownGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ShutdownGatewayWithContext(ctx, input)
}

func (a *Activities) StartAvailabilityMonitorTest(ctx context.Context, input *storagegateway.StartAvailabilityMonitorTestInput) (*storagegateway.StartAvailabilityMonitorTestOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartAvailabilityMonitorTestWithContext(ctx, input)
}

func (a *Activities) StartGateway(ctx context.Context, input *storagegateway.StartGatewayInput) (*storagegateway.StartGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartGatewayWithContext(ctx, input)
}

func (a *Activities) UpdateAutomaticTapeCreationPolicy(ctx context.Context, input *storagegateway.UpdateAutomaticTapeCreationPolicyInput) (*storagegateway.UpdateAutomaticTapeCreationPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateAutomaticTapeCreationPolicyWithContext(ctx, input)
}

func (a *Activities) UpdateBandwidthRateLimit(ctx context.Context, input *storagegateway.UpdateBandwidthRateLimitInput) (*storagegateway.UpdateBandwidthRateLimitOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateBandwidthRateLimitWithContext(ctx, input)
}

func (a *Activities) UpdateChapCredentials(ctx context.Context, input *storagegateway.UpdateChapCredentialsInput) (*storagegateway.UpdateChapCredentialsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateChapCredentialsWithContext(ctx, input)
}

func (a *Activities) UpdateGatewayInformation(ctx context.Context, input *storagegateway.UpdateGatewayInformationInput) (*storagegateway.UpdateGatewayInformationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateGatewayInformationWithContext(ctx, input)
}

func (a *Activities) UpdateGatewaySoftwareNow(ctx context.Context, input *storagegateway.UpdateGatewaySoftwareNowInput) (*storagegateway.UpdateGatewaySoftwareNowOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateGatewaySoftwareNowWithContext(ctx, input)
}

func (a *Activities) UpdateMaintenanceStartTime(ctx context.Context, input *storagegateway.UpdateMaintenanceStartTimeInput) (*storagegateway.UpdateMaintenanceStartTimeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateMaintenanceStartTimeWithContext(ctx, input)
}

func (a *Activities) UpdateNFSFileShare(ctx context.Context, input *storagegateway.UpdateNFSFileShareInput) (*storagegateway.UpdateNFSFileShareOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateNFSFileShareWithContext(ctx, input)
}

func (a *Activities) UpdateSMBFileShare(ctx context.Context, input *storagegateway.UpdateSMBFileShareInput) (*storagegateway.UpdateSMBFileShareOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSMBFileShareWithContext(ctx, input)
}

func (a *Activities) UpdateSMBSecurityStrategy(ctx context.Context, input *storagegateway.UpdateSMBSecurityStrategyInput) (*storagegateway.UpdateSMBSecurityStrategyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSMBSecurityStrategyWithContext(ctx, input)
}

func (a *Activities) UpdateSnapshotSchedule(ctx context.Context, input *storagegateway.UpdateSnapshotScheduleInput) (*storagegateway.UpdateSnapshotScheduleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSnapshotScheduleWithContext(ctx, input)
}

func (a *Activities) UpdateVTLDeviceType(ctx context.Context, input *storagegateway.UpdateVTLDeviceTypeInput) (*storagegateway.UpdateVTLDeviceTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateVTLDeviceTypeWithContext(ctx, input)
}