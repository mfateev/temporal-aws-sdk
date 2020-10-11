// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"go.temporal.io/sdk/workflow"
)

type CloudHSMClient interface {
	AddTagsToResource(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error)
	AddTagsToResourceAsync(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) *CloudHSMAddTagsToResourceFuture

	CreateHapg(ctx workflow.Context, input *cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error)
	CreateHapgAsync(ctx workflow.Context, input *cloudhsm.CreateHapgInput) *CloudHSMCreateHapgFuture

	CreateHsm(ctx workflow.Context, input *cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error)
	CreateHsmAsync(ctx workflow.Context, input *cloudhsm.CreateHsmInput) *CloudHSMCreateHsmFuture

	CreateLunaClient(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error)
	CreateLunaClientAsync(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) *CloudHSMCreateLunaClientFuture

	DeleteHapg(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error)
	DeleteHapgAsync(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) *CloudHSMDeleteHapgFuture

	DeleteHsm(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error)
	DeleteHsmAsync(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) *CloudHSMDeleteHsmFuture

	DeleteLunaClient(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error)
	DeleteLunaClientAsync(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) *CloudHSMDeleteLunaClientFuture

	DescribeHapg(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error)
	DescribeHapgAsync(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) *CloudHSMDescribeHapgFuture

	DescribeHsm(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error)
	DescribeHsmAsync(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) *CloudHSMDescribeHsmFuture

	DescribeLunaClient(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error)
	DescribeLunaClientAsync(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) *CloudHSMDescribeLunaClientFuture

	GetConfig(ctx workflow.Context, input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error)
	GetConfigAsync(ctx workflow.Context, input *cloudhsm.GetConfigInput) *CloudHSMGetConfigFuture

	ListAvailableZones(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error)
	ListAvailableZonesAsync(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) *CloudHSMListAvailableZonesFuture

	ListHapgs(ctx workflow.Context, input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error)
	ListHapgsAsync(ctx workflow.Context, input *cloudhsm.ListHapgsInput) *CloudHSMListHapgsFuture

	ListHsms(ctx workflow.Context, input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error)
	ListHsmsAsync(ctx workflow.Context, input *cloudhsm.ListHsmsInput) *CloudHSMListHsmsFuture

	ListLunaClients(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error)
	ListLunaClientsAsync(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) *CloudHSMListLunaClientsFuture

	ListTagsForResource(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) *CloudHSMListTagsForResourceFuture

	ModifyHapg(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error)
	ModifyHapgAsync(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) *CloudHSMModifyHapgFuture

	ModifyHsm(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error)
	ModifyHsmAsync(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) *CloudHSMModifyHsmFuture

	ModifyLunaClient(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error)
	ModifyLunaClientAsync(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) *CloudHSMModifyLunaClientFuture

	RemoveTagsFromResource(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceAsync(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) *CloudHSMRemoveTagsFromResourceFuture
}

type CloudHSMStub struct{}

func NewCloudHSMStub() CloudHSMClient {
	return &CloudHSMStub{}
}

type CloudHSMAddTagsToResourceFuture struct {
	Future workflow.Future
}

func (r *CloudHSMAddTagsToResourceFuture) Get(ctx workflow.Context) (*cloudhsm.AddTagsToResourceOutput, error) {
	var output cloudhsm.AddTagsToResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMCreateHapgFuture struct {
	Future workflow.Future
}

func (r *CloudHSMCreateHapgFuture) Get(ctx workflow.Context) (*cloudhsm.CreateHapgOutput, error) {
	var output cloudhsm.CreateHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMCreateHsmFuture struct {
	Future workflow.Future
}

func (r *CloudHSMCreateHsmFuture) Get(ctx workflow.Context) (*cloudhsm.CreateHsmOutput, error) {
	var output cloudhsm.CreateHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMCreateLunaClientFuture struct {
	Future workflow.Future
}

func (r *CloudHSMCreateLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.CreateLunaClientOutput, error) {
	var output cloudhsm.CreateLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMDeleteHapgFuture struct {
	Future workflow.Future
}

func (r *CloudHSMDeleteHapgFuture) Get(ctx workflow.Context) (*cloudhsm.DeleteHapgOutput, error) {
	var output cloudhsm.DeleteHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMDeleteHsmFuture struct {
	Future workflow.Future
}

func (r *CloudHSMDeleteHsmFuture) Get(ctx workflow.Context) (*cloudhsm.DeleteHsmOutput, error) {
	var output cloudhsm.DeleteHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMDeleteLunaClientFuture struct {
	Future workflow.Future
}

func (r *CloudHSMDeleteLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.DeleteLunaClientOutput, error) {
	var output cloudhsm.DeleteLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMDescribeHapgFuture struct {
	Future workflow.Future
}

func (r *CloudHSMDescribeHapgFuture) Get(ctx workflow.Context) (*cloudhsm.DescribeHapgOutput, error) {
	var output cloudhsm.DescribeHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMDescribeHsmFuture struct {
	Future workflow.Future
}

func (r *CloudHSMDescribeHsmFuture) Get(ctx workflow.Context) (*cloudhsm.DescribeHsmOutput, error) {
	var output cloudhsm.DescribeHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMDescribeLunaClientFuture struct {
	Future workflow.Future
}

func (r *CloudHSMDescribeLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.DescribeLunaClientOutput, error) {
	var output cloudhsm.DescribeLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMGetConfigFuture struct {
	Future workflow.Future
}

func (r *CloudHSMGetConfigFuture) Get(ctx workflow.Context) (*cloudhsm.GetConfigOutput, error) {
	var output cloudhsm.GetConfigOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMListAvailableZonesFuture struct {
	Future workflow.Future
}

func (r *CloudHSMListAvailableZonesFuture) Get(ctx workflow.Context) (*cloudhsm.ListAvailableZonesOutput, error) {
	var output cloudhsm.ListAvailableZonesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMListHapgsFuture struct {
	Future workflow.Future
}

func (r *CloudHSMListHapgsFuture) Get(ctx workflow.Context) (*cloudhsm.ListHapgsOutput, error) {
	var output cloudhsm.ListHapgsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMListHsmsFuture struct {
	Future workflow.Future
}

func (r *CloudHSMListHsmsFuture) Get(ctx workflow.Context) (*cloudhsm.ListHsmsOutput, error) {
	var output cloudhsm.ListHsmsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMListLunaClientsFuture struct {
	Future workflow.Future
}

func (r *CloudHSMListLunaClientsFuture) Get(ctx workflow.Context) (*cloudhsm.ListLunaClientsOutput, error) {
	var output cloudhsm.ListLunaClientsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *CloudHSMListTagsForResourceFuture) Get(ctx workflow.Context) (*cloudhsm.ListTagsForResourceOutput, error) {
	var output cloudhsm.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMModifyHapgFuture struct {
	Future workflow.Future
}

func (r *CloudHSMModifyHapgFuture) Get(ctx workflow.Context) (*cloudhsm.ModifyHapgOutput, error) {
	var output cloudhsm.ModifyHapgOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMModifyHsmFuture struct {
	Future workflow.Future
}

func (r *CloudHSMModifyHsmFuture) Get(ctx workflow.Context) (*cloudhsm.ModifyHsmOutput, error) {
	var output cloudhsm.ModifyHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMModifyLunaClientFuture struct {
	Future workflow.Future
}

func (r *CloudHSMModifyLunaClientFuture) Get(ctx workflow.Context) (*cloudhsm.ModifyLunaClientOutput, error) {
	var output cloudhsm.ModifyLunaClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMRemoveTagsFromResourceFuture struct {
	Future workflow.Future
}

func (r *CloudHSMRemoveTagsFromResourceFuture) Get(ctx workflow.Context) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
	var output cloudhsm.RemoveTagsFromResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) AddTagsToResource(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) (*cloudhsm.AddTagsToResourceOutput, error) {
	var output cloudhsm.AddTagsToResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.AddTagsToResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) AddTagsToResourceAsync(ctx workflow.Context, input *cloudhsm.AddTagsToResourceInput) *CloudHSMAddTagsToResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.AddTagsToResource", input)
	return &CloudHSMAddTagsToResourceFuture{Future: future}
}

func (a *CloudHSMStub) CreateHapg(ctx workflow.Context, input *cloudhsm.CreateHapgInput) (*cloudhsm.CreateHapgOutput, error) {
	var output cloudhsm.CreateHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) CreateHapgAsync(ctx workflow.Context, input *cloudhsm.CreateHapgInput) *CloudHSMCreateHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateHapg", input)
	return &CloudHSMCreateHapgFuture{Future: future}
}

func (a *CloudHSMStub) CreateHsm(ctx workflow.Context, input *cloudhsm.CreateHsmInput) (*cloudhsm.CreateHsmOutput, error) {
	var output cloudhsm.CreateHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) CreateHsmAsync(ctx workflow.Context, input *cloudhsm.CreateHsmInput) *CloudHSMCreateHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateHsm", input)
	return &CloudHSMCreateHsmFuture{Future: future}
}

func (a *CloudHSMStub) CreateLunaClient(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) (*cloudhsm.CreateLunaClientOutput, error) {
	var output cloudhsm.CreateLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) CreateLunaClientAsync(ctx workflow.Context, input *cloudhsm.CreateLunaClientInput) *CloudHSMCreateLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.CreateLunaClient", input)
	return &CloudHSMCreateLunaClientFuture{Future: future}
}

func (a *CloudHSMStub) DeleteHapg(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) (*cloudhsm.DeleteHapgOutput, error) {
	var output cloudhsm.DeleteHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DeleteHapgAsync(ctx workflow.Context, input *cloudhsm.DeleteHapgInput) *CloudHSMDeleteHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteHapg", input)
	return &CloudHSMDeleteHapgFuture{Future: future}
}

func (a *CloudHSMStub) DeleteHsm(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) (*cloudhsm.DeleteHsmOutput, error) {
	var output cloudhsm.DeleteHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DeleteHsmAsync(ctx workflow.Context, input *cloudhsm.DeleteHsmInput) *CloudHSMDeleteHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteHsm", input)
	return &CloudHSMDeleteHsmFuture{Future: future}
}

func (a *CloudHSMStub) DeleteLunaClient(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) (*cloudhsm.DeleteLunaClientOutput, error) {
	var output cloudhsm.DeleteLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DeleteLunaClientAsync(ctx workflow.Context, input *cloudhsm.DeleteLunaClientInput) *CloudHSMDeleteLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DeleteLunaClient", input)
	return &CloudHSMDeleteLunaClientFuture{Future: future}
}

func (a *CloudHSMStub) DescribeHapg(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) (*cloudhsm.DescribeHapgOutput, error) {
	var output cloudhsm.DescribeHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DescribeHapgAsync(ctx workflow.Context, input *cloudhsm.DescribeHapgInput) *CloudHSMDescribeHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeHapg", input)
	return &CloudHSMDescribeHapgFuture{Future: future}
}

func (a *CloudHSMStub) DescribeHsm(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) (*cloudhsm.DescribeHsmOutput, error) {
	var output cloudhsm.DescribeHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DescribeHsmAsync(ctx workflow.Context, input *cloudhsm.DescribeHsmInput) *CloudHSMDescribeHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeHsm", input)
	return &CloudHSMDescribeHsmFuture{Future: future}
}

func (a *CloudHSMStub) DescribeLunaClient(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) (*cloudhsm.DescribeLunaClientOutput, error) {
	var output cloudhsm.DescribeLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) DescribeLunaClientAsync(ctx workflow.Context, input *cloudhsm.DescribeLunaClientInput) *CloudHSMDescribeLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.DescribeLunaClient", input)
	return &CloudHSMDescribeLunaClientFuture{Future: future}
}

func (a *CloudHSMStub) GetConfig(ctx workflow.Context, input *cloudhsm.GetConfigInput) (*cloudhsm.GetConfigOutput, error) {
	var output cloudhsm.GetConfigOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.GetConfig", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) GetConfigAsync(ctx workflow.Context, input *cloudhsm.GetConfigInput) *CloudHSMGetConfigFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.GetConfig", input)
	return &CloudHSMGetConfigFuture{Future: future}
}

func (a *CloudHSMStub) ListAvailableZones(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) (*cloudhsm.ListAvailableZonesOutput, error) {
	var output cloudhsm.ListAvailableZonesOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListAvailableZones", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ListAvailableZonesAsync(ctx workflow.Context, input *cloudhsm.ListAvailableZonesInput) *CloudHSMListAvailableZonesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListAvailableZones", input)
	return &CloudHSMListAvailableZonesFuture{Future: future}
}

func (a *CloudHSMStub) ListHapgs(ctx workflow.Context, input *cloudhsm.ListHapgsInput) (*cloudhsm.ListHapgsOutput, error) {
	var output cloudhsm.ListHapgsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListHapgs", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ListHapgsAsync(ctx workflow.Context, input *cloudhsm.ListHapgsInput) *CloudHSMListHapgsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListHapgs", input)
	return &CloudHSMListHapgsFuture{Future: future}
}

func (a *CloudHSMStub) ListHsms(ctx workflow.Context, input *cloudhsm.ListHsmsInput) (*cloudhsm.ListHsmsOutput, error) {
	var output cloudhsm.ListHsmsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListHsms", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ListHsmsAsync(ctx workflow.Context, input *cloudhsm.ListHsmsInput) *CloudHSMListHsmsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListHsms", input)
	return &CloudHSMListHsmsFuture{Future: future}
}

func (a *CloudHSMStub) ListLunaClients(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) (*cloudhsm.ListLunaClientsOutput, error) {
	var output cloudhsm.ListLunaClientsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListLunaClients", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ListLunaClientsAsync(ctx workflow.Context, input *cloudhsm.ListLunaClientsInput) *CloudHSMListLunaClientsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListLunaClients", input)
	return &CloudHSMListLunaClientsFuture{Future: future}
}

func (a *CloudHSMStub) ListTagsForResource(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) (*cloudhsm.ListTagsForResourceOutput, error) {
	var output cloudhsm.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ListTagsForResourceAsync(ctx workflow.Context, input *cloudhsm.ListTagsForResourceInput) *CloudHSMListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ListTagsForResource", input)
	return &CloudHSMListTagsForResourceFuture{Future: future}
}

func (a *CloudHSMStub) ModifyHapg(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) (*cloudhsm.ModifyHapgOutput, error) {
	var output cloudhsm.ModifyHapgOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyHapg", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ModifyHapgAsync(ctx workflow.Context, input *cloudhsm.ModifyHapgInput) *CloudHSMModifyHapgFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyHapg", input)
	return &CloudHSMModifyHapgFuture{Future: future}
}

func (a *CloudHSMStub) ModifyHsm(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) (*cloudhsm.ModifyHsmOutput, error) {
	var output cloudhsm.ModifyHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ModifyHsmAsync(ctx workflow.Context, input *cloudhsm.ModifyHsmInput) *CloudHSMModifyHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyHsm", input)
	return &CloudHSMModifyHsmFuture{Future: future}
}

func (a *CloudHSMStub) ModifyLunaClient(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) (*cloudhsm.ModifyLunaClientOutput, error) {
	var output cloudhsm.ModifyLunaClientOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyLunaClient", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) ModifyLunaClientAsync(ctx workflow.Context, input *cloudhsm.ModifyLunaClientInput) *CloudHSMModifyLunaClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.ModifyLunaClient", input)
	return &CloudHSMModifyLunaClientFuture{Future: future}
}

func (a *CloudHSMStub) RemoveTagsFromResource(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) (*cloudhsm.RemoveTagsFromResourceOutput, error) {
	var output cloudhsm.RemoveTagsFromResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsm.RemoveTagsFromResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMStub) RemoveTagsFromResourceAsync(ctx workflow.Context, input *cloudhsm.RemoveTagsFromResourceInput) *CloudHSMRemoveTagsFromResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsm.RemoveTagsFromResource", input)
	return &CloudHSMRemoveTagsFromResourceFuture{Future: future}
}
