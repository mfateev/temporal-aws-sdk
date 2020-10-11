// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"go.temporal.io/sdk/workflow"
)

type ManagedBlockchainClient interface {
	CreateMember(ctx workflow.Context, input *managedblockchain.CreateMemberInput) (*managedblockchain.CreateMemberOutput, error)
	CreateMemberAsync(ctx workflow.Context, input *managedblockchain.CreateMemberInput) *ManagedBlockchainCreateMemberFuture

	CreateNetwork(ctx workflow.Context, input *managedblockchain.CreateNetworkInput) (*managedblockchain.CreateNetworkOutput, error)
	CreateNetworkAsync(ctx workflow.Context, input *managedblockchain.CreateNetworkInput) *ManagedBlockchainCreateNetworkFuture

	CreateNode(ctx workflow.Context, input *managedblockchain.CreateNodeInput) (*managedblockchain.CreateNodeOutput, error)
	CreateNodeAsync(ctx workflow.Context, input *managedblockchain.CreateNodeInput) *ManagedBlockchainCreateNodeFuture

	CreateProposal(ctx workflow.Context, input *managedblockchain.CreateProposalInput) (*managedblockchain.CreateProposalOutput, error)
	CreateProposalAsync(ctx workflow.Context, input *managedblockchain.CreateProposalInput) *ManagedBlockchainCreateProposalFuture

	DeleteMember(ctx workflow.Context, input *managedblockchain.DeleteMemberInput) (*managedblockchain.DeleteMemberOutput, error)
	DeleteMemberAsync(ctx workflow.Context, input *managedblockchain.DeleteMemberInput) *ManagedBlockchainDeleteMemberFuture

	DeleteNode(ctx workflow.Context, input *managedblockchain.DeleteNodeInput) (*managedblockchain.DeleteNodeOutput, error)
	DeleteNodeAsync(ctx workflow.Context, input *managedblockchain.DeleteNodeInput) *ManagedBlockchainDeleteNodeFuture

	GetMember(ctx workflow.Context, input *managedblockchain.GetMemberInput) (*managedblockchain.GetMemberOutput, error)
	GetMemberAsync(ctx workflow.Context, input *managedblockchain.GetMemberInput) *ManagedBlockchainGetMemberFuture

	GetNetwork(ctx workflow.Context, input *managedblockchain.GetNetworkInput) (*managedblockchain.GetNetworkOutput, error)
	GetNetworkAsync(ctx workflow.Context, input *managedblockchain.GetNetworkInput) *ManagedBlockchainGetNetworkFuture

	GetNode(ctx workflow.Context, input *managedblockchain.GetNodeInput) (*managedblockchain.GetNodeOutput, error)
	GetNodeAsync(ctx workflow.Context, input *managedblockchain.GetNodeInput) *ManagedBlockchainGetNodeFuture

	GetProposal(ctx workflow.Context, input *managedblockchain.GetProposalInput) (*managedblockchain.GetProposalOutput, error)
	GetProposalAsync(ctx workflow.Context, input *managedblockchain.GetProposalInput) *ManagedBlockchainGetProposalFuture

	ListInvitations(ctx workflow.Context, input *managedblockchain.ListInvitationsInput) (*managedblockchain.ListInvitationsOutput, error)
	ListInvitationsAsync(ctx workflow.Context, input *managedblockchain.ListInvitationsInput) *ManagedBlockchainListInvitationsFuture

	ListMembers(ctx workflow.Context, input *managedblockchain.ListMembersInput) (*managedblockchain.ListMembersOutput, error)
	ListMembersAsync(ctx workflow.Context, input *managedblockchain.ListMembersInput) *ManagedBlockchainListMembersFuture

	ListNetworks(ctx workflow.Context, input *managedblockchain.ListNetworksInput) (*managedblockchain.ListNetworksOutput, error)
	ListNetworksAsync(ctx workflow.Context, input *managedblockchain.ListNetworksInput) *ManagedBlockchainListNetworksFuture

	ListNodes(ctx workflow.Context, input *managedblockchain.ListNodesInput) (*managedblockchain.ListNodesOutput, error)
	ListNodesAsync(ctx workflow.Context, input *managedblockchain.ListNodesInput) *ManagedBlockchainListNodesFuture

	ListProposalVotes(ctx workflow.Context, input *managedblockchain.ListProposalVotesInput) (*managedblockchain.ListProposalVotesOutput, error)
	ListProposalVotesAsync(ctx workflow.Context, input *managedblockchain.ListProposalVotesInput) *ManagedBlockchainListProposalVotesFuture

	ListProposals(ctx workflow.Context, input *managedblockchain.ListProposalsInput) (*managedblockchain.ListProposalsOutput, error)
	ListProposalsAsync(ctx workflow.Context, input *managedblockchain.ListProposalsInput) *ManagedBlockchainListProposalsFuture

	RejectInvitation(ctx workflow.Context, input *managedblockchain.RejectInvitationInput) (*managedblockchain.RejectInvitationOutput, error)
	RejectInvitationAsync(ctx workflow.Context, input *managedblockchain.RejectInvitationInput) *ManagedBlockchainRejectInvitationFuture

	UpdateMember(ctx workflow.Context, input *managedblockchain.UpdateMemberInput) (*managedblockchain.UpdateMemberOutput, error)
	UpdateMemberAsync(ctx workflow.Context, input *managedblockchain.UpdateMemberInput) *ManagedBlockchainUpdateMemberFuture

	UpdateNode(ctx workflow.Context, input *managedblockchain.UpdateNodeInput) (*managedblockchain.UpdateNodeOutput, error)
	UpdateNodeAsync(ctx workflow.Context, input *managedblockchain.UpdateNodeInput) *ManagedBlockchainUpdateNodeFuture

	VoteOnProposal(ctx workflow.Context, input *managedblockchain.VoteOnProposalInput) (*managedblockchain.VoteOnProposalOutput, error)
	VoteOnProposalAsync(ctx workflow.Context, input *managedblockchain.VoteOnProposalInput) *ManagedBlockchainVoteOnProposalFuture
}

type ManagedBlockchainStub struct{}

func NewManagedBlockchainStub() ManagedBlockchainClient {
	return &ManagedBlockchainStub{}
}

type ManagedBlockchainCreateMemberFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainCreateMemberFuture) Get(ctx workflow.Context) (*managedblockchain.CreateMemberOutput, error) {
	var output managedblockchain.CreateMemberOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainCreateNetworkFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainCreateNetworkFuture) Get(ctx workflow.Context) (*managedblockchain.CreateNetworkOutput, error) {
	var output managedblockchain.CreateNetworkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainCreateNodeFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainCreateNodeFuture) Get(ctx workflow.Context) (*managedblockchain.CreateNodeOutput, error) {
	var output managedblockchain.CreateNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainCreateProposalFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainCreateProposalFuture) Get(ctx workflow.Context) (*managedblockchain.CreateProposalOutput, error) {
	var output managedblockchain.CreateProposalOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainDeleteMemberFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainDeleteMemberFuture) Get(ctx workflow.Context) (*managedblockchain.DeleteMemberOutput, error) {
	var output managedblockchain.DeleteMemberOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainDeleteNodeFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainDeleteNodeFuture) Get(ctx workflow.Context) (*managedblockchain.DeleteNodeOutput, error) {
	var output managedblockchain.DeleteNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainGetMemberFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainGetMemberFuture) Get(ctx workflow.Context) (*managedblockchain.GetMemberOutput, error) {
	var output managedblockchain.GetMemberOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainGetNetworkFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainGetNetworkFuture) Get(ctx workflow.Context) (*managedblockchain.GetNetworkOutput, error) {
	var output managedblockchain.GetNetworkOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainGetNodeFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainGetNodeFuture) Get(ctx workflow.Context) (*managedblockchain.GetNodeOutput, error) {
	var output managedblockchain.GetNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainGetProposalFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainGetProposalFuture) Get(ctx workflow.Context) (*managedblockchain.GetProposalOutput, error) {
	var output managedblockchain.GetProposalOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListInvitationsFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainListInvitationsFuture) Get(ctx workflow.Context) (*managedblockchain.ListInvitationsOutput, error) {
	var output managedblockchain.ListInvitationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListMembersFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainListMembersFuture) Get(ctx workflow.Context) (*managedblockchain.ListMembersOutput, error) {
	var output managedblockchain.ListMembersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListNetworksFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainListNetworksFuture) Get(ctx workflow.Context) (*managedblockchain.ListNetworksOutput, error) {
	var output managedblockchain.ListNetworksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListNodesFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainListNodesFuture) Get(ctx workflow.Context) (*managedblockchain.ListNodesOutput, error) {
	var output managedblockchain.ListNodesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListProposalVotesFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainListProposalVotesFuture) Get(ctx workflow.Context) (*managedblockchain.ListProposalVotesOutput, error) {
	var output managedblockchain.ListProposalVotesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainListProposalsFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainListProposalsFuture) Get(ctx workflow.Context) (*managedblockchain.ListProposalsOutput, error) {
	var output managedblockchain.ListProposalsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainRejectInvitationFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainRejectInvitationFuture) Get(ctx workflow.Context) (*managedblockchain.RejectInvitationOutput, error) {
	var output managedblockchain.RejectInvitationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainUpdateMemberFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainUpdateMemberFuture) Get(ctx workflow.Context) (*managedblockchain.UpdateMemberOutput, error) {
	var output managedblockchain.UpdateMemberOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainUpdateNodeFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainUpdateNodeFuture) Get(ctx workflow.Context) (*managedblockchain.UpdateNodeOutput, error) {
	var output managedblockchain.UpdateNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ManagedBlockchainVoteOnProposalFuture struct {
	Future workflow.Future
}

func (r *ManagedBlockchainVoteOnProposalFuture) Get(ctx workflow.Context) (*managedblockchain.VoteOnProposalOutput, error) {
	var output managedblockchain.VoteOnProposalOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) CreateMember(ctx workflow.Context, input *managedblockchain.CreateMemberInput) (*managedblockchain.CreateMemberOutput, error) {
	var output managedblockchain.CreateMemberOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateMember", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) CreateMemberAsync(ctx workflow.Context, input *managedblockchain.CreateMemberInput) *ManagedBlockchainCreateMemberFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateMember", input)
	return &ManagedBlockchainCreateMemberFuture{Future: future}
}

func (a *ManagedBlockchainStub) CreateNetwork(ctx workflow.Context, input *managedblockchain.CreateNetworkInput) (*managedblockchain.CreateNetworkOutput, error) {
	var output managedblockchain.CreateNetworkOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateNetwork", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) CreateNetworkAsync(ctx workflow.Context, input *managedblockchain.CreateNetworkInput) *ManagedBlockchainCreateNetworkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateNetwork", input)
	return &ManagedBlockchainCreateNetworkFuture{Future: future}
}

func (a *ManagedBlockchainStub) CreateNode(ctx workflow.Context, input *managedblockchain.CreateNodeInput) (*managedblockchain.CreateNodeOutput, error) {
	var output managedblockchain.CreateNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateNode", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) CreateNodeAsync(ctx workflow.Context, input *managedblockchain.CreateNodeInput) *ManagedBlockchainCreateNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateNode", input)
	return &ManagedBlockchainCreateNodeFuture{Future: future}
}

func (a *ManagedBlockchainStub) CreateProposal(ctx workflow.Context, input *managedblockchain.CreateProposalInput) (*managedblockchain.CreateProposalOutput, error) {
	var output managedblockchain.CreateProposalOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateProposal", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) CreateProposalAsync(ctx workflow.Context, input *managedblockchain.CreateProposalInput) *ManagedBlockchainCreateProposalFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.CreateProposal", input)
	return &ManagedBlockchainCreateProposalFuture{Future: future}
}

func (a *ManagedBlockchainStub) DeleteMember(ctx workflow.Context, input *managedblockchain.DeleteMemberInput) (*managedblockchain.DeleteMemberOutput, error) {
	var output managedblockchain.DeleteMemberOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.DeleteMember", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) DeleteMemberAsync(ctx workflow.Context, input *managedblockchain.DeleteMemberInput) *ManagedBlockchainDeleteMemberFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.DeleteMember", input)
	return &ManagedBlockchainDeleteMemberFuture{Future: future}
}

func (a *ManagedBlockchainStub) DeleteNode(ctx workflow.Context, input *managedblockchain.DeleteNodeInput) (*managedblockchain.DeleteNodeOutput, error) {
	var output managedblockchain.DeleteNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.DeleteNode", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) DeleteNodeAsync(ctx workflow.Context, input *managedblockchain.DeleteNodeInput) *ManagedBlockchainDeleteNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.DeleteNode", input)
	return &ManagedBlockchainDeleteNodeFuture{Future: future}
}

func (a *ManagedBlockchainStub) GetMember(ctx workflow.Context, input *managedblockchain.GetMemberInput) (*managedblockchain.GetMemberOutput, error) {
	var output managedblockchain.GetMemberOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetMember", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) GetMemberAsync(ctx workflow.Context, input *managedblockchain.GetMemberInput) *ManagedBlockchainGetMemberFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetMember", input)
	return &ManagedBlockchainGetMemberFuture{Future: future}
}

func (a *ManagedBlockchainStub) GetNetwork(ctx workflow.Context, input *managedblockchain.GetNetworkInput) (*managedblockchain.GetNetworkOutput, error) {
	var output managedblockchain.GetNetworkOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetNetwork", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) GetNetworkAsync(ctx workflow.Context, input *managedblockchain.GetNetworkInput) *ManagedBlockchainGetNetworkFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetNetwork", input)
	return &ManagedBlockchainGetNetworkFuture{Future: future}
}

func (a *ManagedBlockchainStub) GetNode(ctx workflow.Context, input *managedblockchain.GetNodeInput) (*managedblockchain.GetNodeOutput, error) {
	var output managedblockchain.GetNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetNode", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) GetNodeAsync(ctx workflow.Context, input *managedblockchain.GetNodeInput) *ManagedBlockchainGetNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetNode", input)
	return &ManagedBlockchainGetNodeFuture{Future: future}
}

func (a *ManagedBlockchainStub) GetProposal(ctx workflow.Context, input *managedblockchain.GetProposalInput) (*managedblockchain.GetProposalOutput, error) {
	var output managedblockchain.GetProposalOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetProposal", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) GetProposalAsync(ctx workflow.Context, input *managedblockchain.GetProposalInput) *ManagedBlockchainGetProposalFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.GetProposal", input)
	return &ManagedBlockchainGetProposalFuture{Future: future}
}

func (a *ManagedBlockchainStub) ListInvitations(ctx workflow.Context, input *managedblockchain.ListInvitationsInput) (*managedblockchain.ListInvitationsOutput, error) {
	var output managedblockchain.ListInvitationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListInvitations", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) ListInvitationsAsync(ctx workflow.Context, input *managedblockchain.ListInvitationsInput) *ManagedBlockchainListInvitationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListInvitations", input)
	return &ManagedBlockchainListInvitationsFuture{Future: future}
}

func (a *ManagedBlockchainStub) ListMembers(ctx workflow.Context, input *managedblockchain.ListMembersInput) (*managedblockchain.ListMembersOutput, error) {
	var output managedblockchain.ListMembersOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListMembers", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) ListMembersAsync(ctx workflow.Context, input *managedblockchain.ListMembersInput) *ManagedBlockchainListMembersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListMembers", input)
	return &ManagedBlockchainListMembersFuture{Future: future}
}

func (a *ManagedBlockchainStub) ListNetworks(ctx workflow.Context, input *managedblockchain.ListNetworksInput) (*managedblockchain.ListNetworksOutput, error) {
	var output managedblockchain.ListNetworksOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListNetworks", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) ListNetworksAsync(ctx workflow.Context, input *managedblockchain.ListNetworksInput) *ManagedBlockchainListNetworksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListNetworks", input)
	return &ManagedBlockchainListNetworksFuture{Future: future}
}

func (a *ManagedBlockchainStub) ListNodes(ctx workflow.Context, input *managedblockchain.ListNodesInput) (*managedblockchain.ListNodesOutput, error) {
	var output managedblockchain.ListNodesOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListNodes", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) ListNodesAsync(ctx workflow.Context, input *managedblockchain.ListNodesInput) *ManagedBlockchainListNodesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListNodes", input)
	return &ManagedBlockchainListNodesFuture{Future: future}
}

func (a *ManagedBlockchainStub) ListProposalVotes(ctx workflow.Context, input *managedblockchain.ListProposalVotesInput) (*managedblockchain.ListProposalVotesOutput, error) {
	var output managedblockchain.ListProposalVotesOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListProposalVotes", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) ListProposalVotesAsync(ctx workflow.Context, input *managedblockchain.ListProposalVotesInput) *ManagedBlockchainListProposalVotesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListProposalVotes", input)
	return &ManagedBlockchainListProposalVotesFuture{Future: future}
}

func (a *ManagedBlockchainStub) ListProposals(ctx workflow.Context, input *managedblockchain.ListProposalsInput) (*managedblockchain.ListProposalsOutput, error) {
	var output managedblockchain.ListProposalsOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListProposals", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) ListProposalsAsync(ctx workflow.Context, input *managedblockchain.ListProposalsInput) *ManagedBlockchainListProposalsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.ListProposals", input)
	return &ManagedBlockchainListProposalsFuture{Future: future}
}

func (a *ManagedBlockchainStub) RejectInvitation(ctx workflow.Context, input *managedblockchain.RejectInvitationInput) (*managedblockchain.RejectInvitationOutput, error) {
	var output managedblockchain.RejectInvitationOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.RejectInvitation", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) RejectInvitationAsync(ctx workflow.Context, input *managedblockchain.RejectInvitationInput) *ManagedBlockchainRejectInvitationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.RejectInvitation", input)
	return &ManagedBlockchainRejectInvitationFuture{Future: future}
}

func (a *ManagedBlockchainStub) UpdateMember(ctx workflow.Context, input *managedblockchain.UpdateMemberInput) (*managedblockchain.UpdateMemberOutput, error) {
	var output managedblockchain.UpdateMemberOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.UpdateMember", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) UpdateMemberAsync(ctx workflow.Context, input *managedblockchain.UpdateMemberInput) *ManagedBlockchainUpdateMemberFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.UpdateMember", input)
	return &ManagedBlockchainUpdateMemberFuture{Future: future}
}

func (a *ManagedBlockchainStub) UpdateNode(ctx workflow.Context, input *managedblockchain.UpdateNodeInput) (*managedblockchain.UpdateNodeOutput, error) {
	var output managedblockchain.UpdateNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.UpdateNode", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) UpdateNodeAsync(ctx workflow.Context, input *managedblockchain.UpdateNodeInput) *ManagedBlockchainUpdateNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.UpdateNode", input)
	return &ManagedBlockchainUpdateNodeFuture{Future: future}
}

func (a *ManagedBlockchainStub) VoteOnProposal(ctx workflow.Context, input *managedblockchain.VoteOnProposalInput) (*managedblockchain.VoteOnProposalOutput, error) {
	var output managedblockchain.VoteOnProposalOutput
	err := workflow.ExecuteActivity(ctx, "aws.managedblockchain.VoteOnProposal", input).Get(ctx, &output)
	return &output, err
}

func (a *ManagedBlockchainStub) VoteOnProposalAsync(ctx workflow.Context, input *managedblockchain.VoteOnProposalInput) *ManagedBlockchainVoteOnProposalFuture {
	future := workflow.ExecuteActivity(ctx, "aws.managedblockchain.VoteOnProposal", input)
	return &ManagedBlockchainVoteOnProposalFuture{Future: future}
}
