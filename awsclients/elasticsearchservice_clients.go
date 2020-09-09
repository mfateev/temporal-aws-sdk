package awsclients

import (
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ElasticsearchServiceClient interface {
       AcceptInboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.AcceptInboundCrossClusterSearchConnectionOutput, error)
       AcceptInboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput) *ElasticsearchserviceAcceptInboundCrossClusterSearchConnectionResult

       AddTags(ctx workflow.Context, input *elasticsearchservice.AddTagsInput) (*elasticsearchservice.AddTagsOutput, error)
       AddTagsAsync(ctx workflow.Context, input *elasticsearchservice.AddTagsInput) *ElasticsearchserviceAddTagsResult

       AssociatePackage(ctx workflow.Context, input *elasticsearchservice.AssociatePackageInput) (*elasticsearchservice.AssociatePackageOutput, error)
       AssociatePackageAsync(ctx workflow.Context, input *elasticsearchservice.AssociatePackageInput) *ElasticsearchserviceAssociatePackageResult

       CancelElasticsearchServiceSoftwareUpdate(ctx workflow.Context, input *elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateOutput, error)
       CancelElasticsearchServiceSoftwareUpdateAsync(ctx workflow.Context, input *elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput) *ElasticsearchserviceCancelElasticsearchServiceSoftwareUpdateResult

       CreateElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.CreateElasticsearchDomainInput) (*elasticsearchservice.CreateElasticsearchDomainOutput, error)
       CreateElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.CreateElasticsearchDomainInput) *ElasticsearchserviceCreateElasticsearchDomainResult

       CreateOutboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.CreateOutboundCrossClusterSearchConnectionOutput, error)
       CreateOutboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput) *ElasticsearchserviceCreateOutboundCrossClusterSearchConnectionResult

       CreatePackage(ctx workflow.Context, input *elasticsearchservice.CreatePackageInput) (*elasticsearchservice.CreatePackageOutput, error)
       CreatePackageAsync(ctx workflow.Context, input *elasticsearchservice.CreatePackageInput) *ElasticsearchserviceCreatePackageResult

       DeleteElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchDomainInput) (*elasticsearchservice.DeleteElasticsearchDomainOutput, error)
       DeleteElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchDomainInput) *ElasticsearchserviceDeleteElasticsearchDomainResult

       DeleteElasticsearchServiceRole(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchServiceRoleInput) (*elasticsearchservice.DeleteElasticsearchServiceRoleOutput, error)
       DeleteElasticsearchServiceRoleAsync(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchServiceRoleInput) *ElasticsearchserviceDeleteElasticsearchServiceRoleResult

       DeleteInboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteInboundCrossClusterSearchConnectionOutput, error)
       DeleteInboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput) *ElasticsearchserviceDeleteInboundCrossClusterSearchConnectionResult

       DeleteOutboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionOutput, error)
       DeleteOutboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput) *ElasticsearchserviceDeleteOutboundCrossClusterSearchConnectionResult

       DeletePackage(ctx workflow.Context, input *elasticsearchservice.DeletePackageInput) (*elasticsearchservice.DeletePackageOutput, error)
       DeletePackageAsync(ctx workflow.Context, input *elasticsearchservice.DeletePackageInput) *ElasticsearchserviceDeletePackageResult

       DescribeElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainInput) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error)
       DescribeElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainInput) *ElasticsearchserviceDescribeElasticsearchDomainResult

       DescribeElasticsearchDomainConfig(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainConfigInput) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error)
       DescribeElasticsearchDomainConfigAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainConfigInput) *ElasticsearchserviceDescribeElasticsearchDomainConfigResult

       DescribeElasticsearchDomains(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainsInput) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error)
       DescribeElasticsearchDomainsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainsInput) *ElasticsearchserviceDescribeElasticsearchDomainsResult

       DescribeElasticsearchInstanceTypeLimits(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput) (*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput, error)
       DescribeElasticsearchInstanceTypeLimitsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput) *ElasticsearchserviceDescribeElasticsearchInstanceTypeLimitsResult

       DescribeInboundCrossClusterSearchConnections(ctx workflow.Context, input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, error)
       DescribeInboundCrossClusterSearchConnectionsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput) *ElasticsearchserviceDescribeInboundCrossClusterSearchConnectionsResult

       DescribeOutboundCrossClusterSearchConnections(ctx workflow.Context, input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, error)
       DescribeOutboundCrossClusterSearchConnectionsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput) *ElasticsearchserviceDescribeOutboundCrossClusterSearchConnectionsResult

       DescribePackages(ctx workflow.Context, input *elasticsearchservice.DescribePackagesInput) (*elasticsearchservice.DescribePackagesOutput, error)
       DescribePackagesAsync(ctx workflow.Context, input *elasticsearchservice.DescribePackagesInput) *ElasticsearchserviceDescribePackagesResult

       DescribeReservedElasticsearchInstanceOfferings(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput) (*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, error)
       DescribeReservedElasticsearchInstanceOfferingsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput) *ElasticsearchserviceDescribeReservedElasticsearchInstanceOfferingsResult

       DescribeReservedElasticsearchInstances(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput) (*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error)
       DescribeReservedElasticsearchInstancesAsync(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput) *ElasticsearchserviceDescribeReservedElasticsearchInstancesResult

       DissociatePackage(ctx workflow.Context, input *elasticsearchservice.DissociatePackageInput) (*elasticsearchservice.DissociatePackageOutput, error)
       DissociatePackageAsync(ctx workflow.Context, input *elasticsearchservice.DissociatePackageInput) *ElasticsearchserviceDissociatePackageResult

       GetCompatibleElasticsearchVersions(ctx workflow.Context, input *elasticsearchservice.GetCompatibleElasticsearchVersionsInput) (*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput, error)
       GetCompatibleElasticsearchVersionsAsync(ctx workflow.Context, input *elasticsearchservice.GetCompatibleElasticsearchVersionsInput) *ElasticsearchserviceGetCompatibleElasticsearchVersionsResult

       GetUpgradeHistory(ctx workflow.Context, input *elasticsearchservice.GetUpgradeHistoryInput) (*elasticsearchservice.GetUpgradeHistoryOutput, error)
       GetUpgradeHistoryAsync(ctx workflow.Context, input *elasticsearchservice.GetUpgradeHistoryInput) *ElasticsearchserviceGetUpgradeHistoryResult

       GetUpgradeStatus(ctx workflow.Context, input *elasticsearchservice.GetUpgradeStatusInput) (*elasticsearchservice.GetUpgradeStatusOutput, error)
       GetUpgradeStatusAsync(ctx workflow.Context, input *elasticsearchservice.GetUpgradeStatusInput) *ElasticsearchserviceGetUpgradeStatusResult

       ListDomainNames(ctx workflow.Context, input *elasticsearchservice.ListDomainNamesInput) (*elasticsearchservice.ListDomainNamesOutput, error)
       ListDomainNamesAsync(ctx workflow.Context, input *elasticsearchservice.ListDomainNamesInput) *ElasticsearchserviceListDomainNamesResult

       ListDomainsForPackage(ctx workflow.Context, input *elasticsearchservice.ListDomainsForPackageInput) (*elasticsearchservice.ListDomainsForPackageOutput, error)
       ListDomainsForPackageAsync(ctx workflow.Context, input *elasticsearchservice.ListDomainsForPackageInput) *ElasticsearchserviceListDomainsForPackageResult

       ListElasticsearchInstanceTypes(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchInstanceTypesInput) (*elasticsearchservice.ListElasticsearchInstanceTypesOutput, error)
       ListElasticsearchInstanceTypesAsync(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchInstanceTypesInput) *ElasticsearchserviceListElasticsearchInstanceTypesResult

       ListElasticsearchVersions(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchVersionsInput) (*elasticsearchservice.ListElasticsearchVersionsOutput, error)
       ListElasticsearchVersionsAsync(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchVersionsInput) *ElasticsearchserviceListElasticsearchVersionsResult

       ListPackagesForDomain(ctx workflow.Context, input *elasticsearchservice.ListPackagesForDomainInput) (*elasticsearchservice.ListPackagesForDomainOutput, error)
       ListPackagesForDomainAsync(ctx workflow.Context, input *elasticsearchservice.ListPackagesForDomainInput) *ElasticsearchserviceListPackagesForDomainResult

       ListTags(ctx workflow.Context, input *elasticsearchservice.ListTagsInput) (*elasticsearchservice.ListTagsOutput, error)
       ListTagsAsync(ctx workflow.Context, input *elasticsearchservice.ListTagsInput) *ElasticsearchserviceListTagsResult

       PurchaseReservedElasticsearchInstanceOffering(ctx workflow.Context, input *elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput) (*elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingOutput, error)
       PurchaseReservedElasticsearchInstanceOfferingAsync(ctx workflow.Context, input *elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput) *ElasticsearchservicePurchaseReservedElasticsearchInstanceOfferingResult

       RejectInboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.RejectInboundCrossClusterSearchConnectionOutput, error)
       RejectInboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput) *ElasticsearchserviceRejectInboundCrossClusterSearchConnectionResult

       RemoveTags(ctx workflow.Context, input *elasticsearchservice.RemoveTagsInput) (*elasticsearchservice.RemoveTagsOutput, error)
       RemoveTagsAsync(ctx workflow.Context, input *elasticsearchservice.RemoveTagsInput) *ElasticsearchserviceRemoveTagsResult

       StartElasticsearchServiceSoftwareUpdate(ctx workflow.Context, input *elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.StartElasticsearchServiceSoftwareUpdateOutput, error)
       StartElasticsearchServiceSoftwareUpdateAsync(ctx workflow.Context, input *elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput) *ElasticsearchserviceStartElasticsearchServiceSoftwareUpdateResult

       UpdateElasticsearchDomainConfig(ctx workflow.Context, input *elasticsearchservice.UpdateElasticsearchDomainConfigInput) (*elasticsearchservice.UpdateElasticsearchDomainConfigOutput, error)
       UpdateElasticsearchDomainConfigAsync(ctx workflow.Context, input *elasticsearchservice.UpdateElasticsearchDomainConfigInput) *ElasticsearchserviceUpdateElasticsearchDomainConfigResult

       UpgradeElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.UpgradeElasticsearchDomainInput) (*elasticsearchservice.UpgradeElasticsearchDomainOutput, error)
       UpgradeElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.UpgradeElasticsearchDomainInput) *ElasticsearchserviceUpgradeElasticsearchDomainResult
}

type ElasticsearchserviceAcceptInboundCrossClusterSearchConnectionResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceAcceptInboundCrossClusterSearchConnectionResult) Get(ctx workflow.Context) (*elasticsearchservice.AcceptInboundCrossClusterSearchConnectionOutput, error) {
    var output elasticsearchservice.AcceptInboundCrossClusterSearchConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceAddTagsResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceAddTagsResult) Get(ctx workflow.Context) (*elasticsearchservice.AddTagsOutput, error) {
    var output elasticsearchservice.AddTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceAssociatePackageResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceAssociatePackageResult) Get(ctx workflow.Context) (*elasticsearchservice.AssociatePackageOutput, error) {
    var output elasticsearchservice.AssociatePackageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceCancelElasticsearchServiceSoftwareUpdateResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceCancelElasticsearchServiceSoftwareUpdateResult) Get(ctx workflow.Context) (*elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateOutput, error) {
    var output elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceCreateElasticsearchDomainResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceCreateElasticsearchDomainResult) Get(ctx workflow.Context) (*elasticsearchservice.CreateElasticsearchDomainOutput, error) {
    var output elasticsearchservice.CreateElasticsearchDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceCreateOutboundCrossClusterSearchConnectionResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceCreateOutboundCrossClusterSearchConnectionResult) Get(ctx workflow.Context) (*elasticsearchservice.CreateOutboundCrossClusterSearchConnectionOutput, error) {
    var output elasticsearchservice.CreateOutboundCrossClusterSearchConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceCreatePackageResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceCreatePackageResult) Get(ctx workflow.Context) (*elasticsearchservice.CreatePackageOutput, error) {
    var output elasticsearchservice.CreatePackageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDeleteElasticsearchDomainResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDeleteElasticsearchDomainResult) Get(ctx workflow.Context) (*elasticsearchservice.DeleteElasticsearchDomainOutput, error) {
    var output elasticsearchservice.DeleteElasticsearchDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDeleteElasticsearchServiceRoleResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDeleteElasticsearchServiceRoleResult) Get(ctx workflow.Context) (*elasticsearchservice.DeleteElasticsearchServiceRoleOutput, error) {
    var output elasticsearchservice.DeleteElasticsearchServiceRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDeleteInboundCrossClusterSearchConnectionResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDeleteInboundCrossClusterSearchConnectionResult) Get(ctx workflow.Context) (*elasticsearchservice.DeleteInboundCrossClusterSearchConnectionOutput, error) {
    var output elasticsearchservice.DeleteInboundCrossClusterSearchConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDeleteOutboundCrossClusterSearchConnectionResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDeleteOutboundCrossClusterSearchConnectionResult) Get(ctx workflow.Context) (*elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionOutput, error) {
    var output elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDeletePackageResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDeletePackageResult) Get(ctx workflow.Context) (*elasticsearchservice.DeletePackageOutput, error) {
    var output elasticsearchservice.DeletePackageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDescribeElasticsearchDomainResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDescribeElasticsearchDomainResult) Get(ctx workflow.Context) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error) {
    var output elasticsearchservice.DescribeElasticsearchDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDescribeElasticsearchDomainConfigResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDescribeElasticsearchDomainConfigResult) Get(ctx workflow.Context) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error) {
    var output elasticsearchservice.DescribeElasticsearchDomainConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDescribeElasticsearchDomainsResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDescribeElasticsearchDomainsResult) Get(ctx workflow.Context) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error) {
    var output elasticsearchservice.DescribeElasticsearchDomainsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDescribeElasticsearchInstanceTypeLimitsResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDescribeElasticsearchInstanceTypeLimitsResult) Get(ctx workflow.Context) (*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput, error) {
    var output elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDescribeInboundCrossClusterSearchConnectionsResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDescribeInboundCrossClusterSearchConnectionsResult) Get(ctx workflow.Context) (*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, error) {
    var output elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDescribeOutboundCrossClusterSearchConnectionsResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDescribeOutboundCrossClusterSearchConnectionsResult) Get(ctx workflow.Context) (*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
    var output elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDescribePackagesResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDescribePackagesResult) Get(ctx workflow.Context) (*elasticsearchservice.DescribePackagesOutput, error) {
    var output elasticsearchservice.DescribePackagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDescribeReservedElasticsearchInstanceOfferingsResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDescribeReservedElasticsearchInstanceOfferingsResult) Get(ctx workflow.Context) (*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, error) {
    var output elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDescribeReservedElasticsearchInstancesResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDescribeReservedElasticsearchInstancesResult) Get(ctx workflow.Context) (*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error) {
    var output elasticsearchservice.DescribeReservedElasticsearchInstancesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceDissociatePackageResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceDissociatePackageResult) Get(ctx workflow.Context) (*elasticsearchservice.DissociatePackageOutput, error) {
    var output elasticsearchservice.DissociatePackageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceGetCompatibleElasticsearchVersionsResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceGetCompatibleElasticsearchVersionsResult) Get(ctx workflow.Context) (*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput, error) {
    var output elasticsearchservice.GetCompatibleElasticsearchVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceGetUpgradeHistoryResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceGetUpgradeHistoryResult) Get(ctx workflow.Context) (*elasticsearchservice.GetUpgradeHistoryOutput, error) {
    var output elasticsearchservice.GetUpgradeHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceGetUpgradeStatusResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceGetUpgradeStatusResult) Get(ctx workflow.Context) (*elasticsearchservice.GetUpgradeStatusOutput, error) {
    var output elasticsearchservice.GetUpgradeStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceListDomainNamesResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceListDomainNamesResult) Get(ctx workflow.Context) (*elasticsearchservice.ListDomainNamesOutput, error) {
    var output elasticsearchservice.ListDomainNamesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceListDomainsForPackageResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceListDomainsForPackageResult) Get(ctx workflow.Context) (*elasticsearchservice.ListDomainsForPackageOutput, error) {
    var output elasticsearchservice.ListDomainsForPackageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceListElasticsearchInstanceTypesResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceListElasticsearchInstanceTypesResult) Get(ctx workflow.Context) (*elasticsearchservice.ListElasticsearchInstanceTypesOutput, error) {
    var output elasticsearchservice.ListElasticsearchInstanceTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceListElasticsearchVersionsResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceListElasticsearchVersionsResult) Get(ctx workflow.Context) (*elasticsearchservice.ListElasticsearchVersionsOutput, error) {
    var output elasticsearchservice.ListElasticsearchVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceListPackagesForDomainResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceListPackagesForDomainResult) Get(ctx workflow.Context) (*elasticsearchservice.ListPackagesForDomainOutput, error) {
    var output elasticsearchservice.ListPackagesForDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceListTagsResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceListTagsResult) Get(ctx workflow.Context) (*elasticsearchservice.ListTagsOutput, error) {
    var output elasticsearchservice.ListTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchservicePurchaseReservedElasticsearchInstanceOfferingResult struct {
	Result workflow.Future
}

func (r *ElasticsearchservicePurchaseReservedElasticsearchInstanceOfferingResult) Get(ctx workflow.Context) (*elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingOutput, error) {
    var output elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceRejectInboundCrossClusterSearchConnectionResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceRejectInboundCrossClusterSearchConnectionResult) Get(ctx workflow.Context) (*elasticsearchservice.RejectInboundCrossClusterSearchConnectionOutput, error) {
    var output elasticsearchservice.RejectInboundCrossClusterSearchConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceRemoveTagsResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceRemoveTagsResult) Get(ctx workflow.Context) (*elasticsearchservice.RemoveTagsOutput, error) {
    var output elasticsearchservice.RemoveTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceStartElasticsearchServiceSoftwareUpdateResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceStartElasticsearchServiceSoftwareUpdateResult) Get(ctx workflow.Context) (*elasticsearchservice.StartElasticsearchServiceSoftwareUpdateOutput, error) {
    var output elasticsearchservice.StartElasticsearchServiceSoftwareUpdateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceUpdateElasticsearchDomainConfigResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceUpdateElasticsearchDomainConfigResult) Get(ctx workflow.Context) (*elasticsearchservice.UpdateElasticsearchDomainConfigOutput, error) {
    var output elasticsearchservice.UpdateElasticsearchDomainConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchserviceUpgradeElasticsearchDomainResult struct {
	Result workflow.Future
}

func (r *ElasticsearchserviceUpgradeElasticsearchDomainResult) Get(ctx workflow.Context) (*elasticsearchservice.UpgradeElasticsearchDomainOutput, error) {
    var output elasticsearchservice.UpgradeElasticsearchDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticsearchServiceStub struct {
    activities awsactivities.ElasticsearchServiceActivities
}

func NewElasticsearchServiceStub() ElasticsearchServiceClient {
    return &ElasticsearchServiceStub{}
}

func (a *ElasticsearchServiceStub) AcceptInboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.AcceptInboundCrossClusterSearchConnectionOutput, error) {
    var output elasticsearchservice.AcceptInboundCrossClusterSearchConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptInboundCrossClusterSearchConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) AcceptInboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput) *ElasticsearchserviceAcceptInboundCrossClusterSearchConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptInboundCrossClusterSearchConnection, input)
    return &ElasticsearchserviceAcceptInboundCrossClusterSearchConnectionResult{Result: future}
}

func (a *ElasticsearchServiceStub) AddTags(ctx workflow.Context, input *elasticsearchservice.AddTagsInput) (*elasticsearchservice.AddTagsOutput, error) {
    var output elasticsearchservice.AddTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddTags, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) AddTagsAsync(ctx workflow.Context, input *elasticsearchservice.AddTagsInput) *ElasticsearchserviceAddTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddTags, input)
    return &ElasticsearchserviceAddTagsResult{Result: future}
}

func (a *ElasticsearchServiceStub) AssociatePackage(ctx workflow.Context, input *elasticsearchservice.AssociatePackageInput) (*elasticsearchservice.AssociatePackageOutput, error) {
    var output elasticsearchservice.AssociatePackageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociatePackage, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) AssociatePackageAsync(ctx workflow.Context, input *elasticsearchservice.AssociatePackageInput) *ElasticsearchserviceAssociatePackageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociatePackage, input)
    return &ElasticsearchserviceAssociatePackageResult{Result: future}
}

func (a *ElasticsearchServiceStub) CancelElasticsearchServiceSoftwareUpdate(ctx workflow.Context, input *elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateOutput, error) {
    var output elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelElasticsearchServiceSoftwareUpdate, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) CancelElasticsearchServiceSoftwareUpdateAsync(ctx workflow.Context, input *elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput) *ElasticsearchserviceCancelElasticsearchServiceSoftwareUpdateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CancelElasticsearchServiceSoftwareUpdate, input)
    return &ElasticsearchserviceCancelElasticsearchServiceSoftwareUpdateResult{Result: future}
}

func (a *ElasticsearchServiceStub) CreateElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.CreateElasticsearchDomainInput) (*elasticsearchservice.CreateElasticsearchDomainOutput, error) {
    var output elasticsearchservice.CreateElasticsearchDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateElasticsearchDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) CreateElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.CreateElasticsearchDomainInput) *ElasticsearchserviceCreateElasticsearchDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateElasticsearchDomain, input)
    return &ElasticsearchserviceCreateElasticsearchDomainResult{Result: future}
}

func (a *ElasticsearchServiceStub) CreateOutboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.CreateOutboundCrossClusterSearchConnectionOutput, error) {
    var output elasticsearchservice.CreateOutboundCrossClusterSearchConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateOutboundCrossClusterSearchConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) CreateOutboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput) *ElasticsearchserviceCreateOutboundCrossClusterSearchConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateOutboundCrossClusterSearchConnection, input)
    return &ElasticsearchserviceCreateOutboundCrossClusterSearchConnectionResult{Result: future}
}

func (a *ElasticsearchServiceStub) CreatePackage(ctx workflow.Context, input *elasticsearchservice.CreatePackageInput) (*elasticsearchservice.CreatePackageOutput, error) {
    var output elasticsearchservice.CreatePackageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePackage, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) CreatePackageAsync(ctx workflow.Context, input *elasticsearchservice.CreatePackageInput) *ElasticsearchserviceCreatePackageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePackage, input)
    return &ElasticsearchserviceCreatePackageResult{Result: future}
}

func (a *ElasticsearchServiceStub) DeleteElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchDomainInput) (*elasticsearchservice.DeleteElasticsearchDomainOutput, error) {
    var output elasticsearchservice.DeleteElasticsearchDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteElasticsearchDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DeleteElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchDomainInput) *ElasticsearchserviceDeleteElasticsearchDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteElasticsearchDomain, input)
    return &ElasticsearchserviceDeleteElasticsearchDomainResult{Result: future}
}

func (a *ElasticsearchServiceStub) DeleteElasticsearchServiceRole(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchServiceRoleInput) (*elasticsearchservice.DeleteElasticsearchServiceRoleOutput, error) {
    var output elasticsearchservice.DeleteElasticsearchServiceRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteElasticsearchServiceRole, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DeleteElasticsearchServiceRoleAsync(ctx workflow.Context, input *elasticsearchservice.DeleteElasticsearchServiceRoleInput) *ElasticsearchserviceDeleteElasticsearchServiceRoleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteElasticsearchServiceRole, input)
    return &ElasticsearchserviceDeleteElasticsearchServiceRoleResult{Result: future}
}

func (a *ElasticsearchServiceStub) DeleteInboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteInboundCrossClusterSearchConnectionOutput, error) {
    var output elasticsearchservice.DeleteInboundCrossClusterSearchConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInboundCrossClusterSearchConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DeleteInboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput) *ElasticsearchserviceDeleteInboundCrossClusterSearchConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteInboundCrossClusterSearchConnection, input)
    return &ElasticsearchserviceDeleteInboundCrossClusterSearchConnectionResult{Result: future}
}

func (a *ElasticsearchServiceStub) DeleteOutboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput) (*elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionOutput, error) {
    var output elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteOutboundCrossClusterSearchConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DeleteOutboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput) *ElasticsearchserviceDeleteOutboundCrossClusterSearchConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteOutboundCrossClusterSearchConnection, input)
    return &ElasticsearchserviceDeleteOutboundCrossClusterSearchConnectionResult{Result: future}
}

func (a *ElasticsearchServiceStub) DeletePackage(ctx workflow.Context, input *elasticsearchservice.DeletePackageInput) (*elasticsearchservice.DeletePackageOutput, error) {
    var output elasticsearchservice.DeletePackageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePackage, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DeletePackageAsync(ctx workflow.Context, input *elasticsearchservice.DeletePackageInput) *ElasticsearchserviceDeletePackageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePackage, input)
    return &ElasticsearchserviceDeletePackageResult{Result: future}
}

func (a *ElasticsearchServiceStub) DescribeElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainInput) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error) {
    var output elasticsearchservice.DescribeElasticsearchDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticsearchDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DescribeElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainInput) *ElasticsearchserviceDescribeElasticsearchDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticsearchDomain, input)
    return &ElasticsearchserviceDescribeElasticsearchDomainResult{Result: future}
}

func (a *ElasticsearchServiceStub) DescribeElasticsearchDomainConfig(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainConfigInput) (*elasticsearchservice.DescribeElasticsearchDomainConfigOutput, error) {
    var output elasticsearchservice.DescribeElasticsearchDomainConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticsearchDomainConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DescribeElasticsearchDomainConfigAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainConfigInput) *ElasticsearchserviceDescribeElasticsearchDomainConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticsearchDomainConfig, input)
    return &ElasticsearchserviceDescribeElasticsearchDomainConfigResult{Result: future}
}

func (a *ElasticsearchServiceStub) DescribeElasticsearchDomains(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainsInput) (*elasticsearchservice.DescribeElasticsearchDomainsOutput, error) {
    var output elasticsearchservice.DescribeElasticsearchDomainsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticsearchDomains, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DescribeElasticsearchDomainsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchDomainsInput) *ElasticsearchserviceDescribeElasticsearchDomainsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticsearchDomains, input)
    return &ElasticsearchserviceDescribeElasticsearchDomainsResult{Result: future}
}

func (a *ElasticsearchServiceStub) DescribeElasticsearchInstanceTypeLimits(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput) (*elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput, error) {
    var output elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticsearchInstanceTypeLimits, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DescribeElasticsearchInstanceTypeLimitsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput) *ElasticsearchserviceDescribeElasticsearchInstanceTypeLimitsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeElasticsearchInstanceTypeLimits, input)
    return &ElasticsearchserviceDescribeElasticsearchInstanceTypeLimitsResult{Result: future}
}

func (a *ElasticsearchServiceStub) DescribeInboundCrossClusterSearchConnections(ctx workflow.Context, input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput, error) {
    var output elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInboundCrossClusterSearchConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DescribeInboundCrossClusterSearchConnectionsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput) *ElasticsearchserviceDescribeInboundCrossClusterSearchConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInboundCrossClusterSearchConnections, input)
    return &ElasticsearchserviceDescribeInboundCrossClusterSearchConnectionsResult{Result: future}
}

func (a *ElasticsearchServiceStub) DescribeOutboundCrossClusterSearchConnections(ctx workflow.Context, input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput) (*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
    var output elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeOutboundCrossClusterSearchConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DescribeOutboundCrossClusterSearchConnectionsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput) *ElasticsearchserviceDescribeOutboundCrossClusterSearchConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeOutboundCrossClusterSearchConnections, input)
    return &ElasticsearchserviceDescribeOutboundCrossClusterSearchConnectionsResult{Result: future}
}

func (a *ElasticsearchServiceStub) DescribePackages(ctx workflow.Context, input *elasticsearchservice.DescribePackagesInput) (*elasticsearchservice.DescribePackagesOutput, error) {
    var output elasticsearchservice.DescribePackagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePackages, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DescribePackagesAsync(ctx workflow.Context, input *elasticsearchservice.DescribePackagesInput) *ElasticsearchserviceDescribePackagesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePackages, input)
    return &ElasticsearchserviceDescribePackagesResult{Result: future}
}

func (a *ElasticsearchServiceStub) DescribeReservedElasticsearchInstanceOfferings(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput) (*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput, error) {
    var output elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedElasticsearchInstanceOfferings, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DescribeReservedElasticsearchInstanceOfferingsAsync(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput) *ElasticsearchserviceDescribeReservedElasticsearchInstanceOfferingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedElasticsearchInstanceOfferings, input)
    return &ElasticsearchserviceDescribeReservedElasticsearchInstanceOfferingsResult{Result: future}
}

func (a *ElasticsearchServiceStub) DescribeReservedElasticsearchInstances(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput) (*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput, error) {
    var output elasticsearchservice.DescribeReservedElasticsearchInstancesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedElasticsearchInstances, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DescribeReservedElasticsearchInstancesAsync(ctx workflow.Context, input *elasticsearchservice.DescribeReservedElasticsearchInstancesInput) *ElasticsearchserviceDescribeReservedElasticsearchInstancesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeReservedElasticsearchInstances, input)
    return &ElasticsearchserviceDescribeReservedElasticsearchInstancesResult{Result: future}
}

func (a *ElasticsearchServiceStub) DissociatePackage(ctx workflow.Context, input *elasticsearchservice.DissociatePackageInput) (*elasticsearchservice.DissociatePackageOutput, error) {
    var output elasticsearchservice.DissociatePackageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DissociatePackage, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) DissociatePackageAsync(ctx workflow.Context, input *elasticsearchservice.DissociatePackageInput) *ElasticsearchserviceDissociatePackageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DissociatePackage, input)
    return &ElasticsearchserviceDissociatePackageResult{Result: future}
}

func (a *ElasticsearchServiceStub) GetCompatibleElasticsearchVersions(ctx workflow.Context, input *elasticsearchservice.GetCompatibleElasticsearchVersionsInput) (*elasticsearchservice.GetCompatibleElasticsearchVersionsOutput, error) {
    var output elasticsearchservice.GetCompatibleElasticsearchVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCompatibleElasticsearchVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) GetCompatibleElasticsearchVersionsAsync(ctx workflow.Context, input *elasticsearchservice.GetCompatibleElasticsearchVersionsInput) *ElasticsearchserviceGetCompatibleElasticsearchVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCompatibleElasticsearchVersions, input)
    return &ElasticsearchserviceGetCompatibleElasticsearchVersionsResult{Result: future}
}

func (a *ElasticsearchServiceStub) GetUpgradeHistory(ctx workflow.Context, input *elasticsearchservice.GetUpgradeHistoryInput) (*elasticsearchservice.GetUpgradeHistoryOutput, error) {
    var output elasticsearchservice.GetUpgradeHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUpgradeHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) GetUpgradeHistoryAsync(ctx workflow.Context, input *elasticsearchservice.GetUpgradeHistoryInput) *ElasticsearchserviceGetUpgradeHistoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetUpgradeHistory, input)
    return &ElasticsearchserviceGetUpgradeHistoryResult{Result: future}
}

func (a *ElasticsearchServiceStub) GetUpgradeStatus(ctx workflow.Context, input *elasticsearchservice.GetUpgradeStatusInput) (*elasticsearchservice.GetUpgradeStatusOutput, error) {
    var output elasticsearchservice.GetUpgradeStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetUpgradeStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) GetUpgradeStatusAsync(ctx workflow.Context, input *elasticsearchservice.GetUpgradeStatusInput) *ElasticsearchserviceGetUpgradeStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetUpgradeStatus, input)
    return &ElasticsearchserviceGetUpgradeStatusResult{Result: future}
}

func (a *ElasticsearchServiceStub) ListDomainNames(ctx workflow.Context, input *elasticsearchservice.ListDomainNamesInput) (*elasticsearchservice.ListDomainNamesOutput, error) {
    var output elasticsearchservice.ListDomainNamesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomainNames, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) ListDomainNamesAsync(ctx workflow.Context, input *elasticsearchservice.ListDomainNamesInput) *ElasticsearchserviceListDomainNamesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDomainNames, input)
    return &ElasticsearchserviceListDomainNamesResult{Result: future}
}

func (a *ElasticsearchServiceStub) ListDomainsForPackage(ctx workflow.Context, input *elasticsearchservice.ListDomainsForPackageInput) (*elasticsearchservice.ListDomainsForPackageOutput, error) {
    var output elasticsearchservice.ListDomainsForPackageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomainsForPackage, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) ListDomainsForPackageAsync(ctx workflow.Context, input *elasticsearchservice.ListDomainsForPackageInput) *ElasticsearchserviceListDomainsForPackageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDomainsForPackage, input)
    return &ElasticsearchserviceListDomainsForPackageResult{Result: future}
}

func (a *ElasticsearchServiceStub) ListElasticsearchInstanceTypes(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchInstanceTypesInput) (*elasticsearchservice.ListElasticsearchInstanceTypesOutput, error) {
    var output elasticsearchservice.ListElasticsearchInstanceTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListElasticsearchInstanceTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) ListElasticsearchInstanceTypesAsync(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchInstanceTypesInput) *ElasticsearchserviceListElasticsearchInstanceTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListElasticsearchInstanceTypes, input)
    return &ElasticsearchserviceListElasticsearchInstanceTypesResult{Result: future}
}

func (a *ElasticsearchServiceStub) ListElasticsearchVersions(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchVersionsInput) (*elasticsearchservice.ListElasticsearchVersionsOutput, error) {
    var output elasticsearchservice.ListElasticsearchVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListElasticsearchVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) ListElasticsearchVersionsAsync(ctx workflow.Context, input *elasticsearchservice.ListElasticsearchVersionsInput) *ElasticsearchserviceListElasticsearchVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListElasticsearchVersions, input)
    return &ElasticsearchserviceListElasticsearchVersionsResult{Result: future}
}

func (a *ElasticsearchServiceStub) ListPackagesForDomain(ctx workflow.Context, input *elasticsearchservice.ListPackagesForDomainInput) (*elasticsearchservice.ListPackagesForDomainOutput, error) {
    var output elasticsearchservice.ListPackagesForDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPackagesForDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) ListPackagesForDomainAsync(ctx workflow.Context, input *elasticsearchservice.ListPackagesForDomainInput) *ElasticsearchserviceListPackagesForDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPackagesForDomain, input)
    return &ElasticsearchserviceListPackagesForDomainResult{Result: future}
}

func (a *ElasticsearchServiceStub) ListTags(ctx workflow.Context, input *elasticsearchservice.ListTagsInput) (*elasticsearchservice.ListTagsOutput, error) {
    var output elasticsearchservice.ListTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTags, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) ListTagsAsync(ctx workflow.Context, input *elasticsearchservice.ListTagsInput) *ElasticsearchserviceListTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTags, input)
    return &ElasticsearchserviceListTagsResult{Result: future}
}

func (a *ElasticsearchServiceStub) PurchaseReservedElasticsearchInstanceOffering(ctx workflow.Context, input *elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput) (*elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingOutput, error) {
    var output elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PurchaseReservedElasticsearchInstanceOffering, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) PurchaseReservedElasticsearchInstanceOfferingAsync(ctx workflow.Context, input *elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput) *ElasticsearchservicePurchaseReservedElasticsearchInstanceOfferingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PurchaseReservedElasticsearchInstanceOffering, input)
    return &ElasticsearchservicePurchaseReservedElasticsearchInstanceOfferingResult{Result: future}
}

func (a *ElasticsearchServiceStub) RejectInboundCrossClusterSearchConnection(ctx workflow.Context, input *elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput) (*elasticsearchservice.RejectInboundCrossClusterSearchConnectionOutput, error) {
    var output elasticsearchservice.RejectInboundCrossClusterSearchConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectInboundCrossClusterSearchConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) RejectInboundCrossClusterSearchConnectionAsync(ctx workflow.Context, input *elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput) *ElasticsearchserviceRejectInboundCrossClusterSearchConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RejectInboundCrossClusterSearchConnection, input)
    return &ElasticsearchserviceRejectInboundCrossClusterSearchConnectionResult{Result: future}
}

func (a *ElasticsearchServiceStub) RemoveTags(ctx workflow.Context, input *elasticsearchservice.RemoveTagsInput) (*elasticsearchservice.RemoveTagsOutput, error) {
    var output elasticsearchservice.RemoveTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTags, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) RemoveTagsAsync(ctx workflow.Context, input *elasticsearchservice.RemoveTagsInput) *ElasticsearchserviceRemoveTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemoveTags, input)
    return &ElasticsearchserviceRemoveTagsResult{Result: future}
}

func (a *ElasticsearchServiceStub) StartElasticsearchServiceSoftwareUpdate(ctx workflow.Context, input *elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput) (*elasticsearchservice.StartElasticsearchServiceSoftwareUpdateOutput, error) {
    var output elasticsearchservice.StartElasticsearchServiceSoftwareUpdateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartElasticsearchServiceSoftwareUpdate, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) StartElasticsearchServiceSoftwareUpdateAsync(ctx workflow.Context, input *elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput) *ElasticsearchserviceStartElasticsearchServiceSoftwareUpdateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartElasticsearchServiceSoftwareUpdate, input)
    return &ElasticsearchserviceStartElasticsearchServiceSoftwareUpdateResult{Result: future}
}

func (a *ElasticsearchServiceStub) UpdateElasticsearchDomainConfig(ctx workflow.Context, input *elasticsearchservice.UpdateElasticsearchDomainConfigInput) (*elasticsearchservice.UpdateElasticsearchDomainConfigOutput, error) {
    var output elasticsearchservice.UpdateElasticsearchDomainConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateElasticsearchDomainConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) UpdateElasticsearchDomainConfigAsync(ctx workflow.Context, input *elasticsearchservice.UpdateElasticsearchDomainConfigInput) *ElasticsearchserviceUpdateElasticsearchDomainConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateElasticsearchDomainConfig, input)
    return &ElasticsearchserviceUpdateElasticsearchDomainConfigResult{Result: future}
}

func (a *ElasticsearchServiceStub) UpgradeElasticsearchDomain(ctx workflow.Context, input *elasticsearchservice.UpgradeElasticsearchDomainInput) (*elasticsearchservice.UpgradeElasticsearchDomainOutput, error) {
    var output elasticsearchservice.UpgradeElasticsearchDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpgradeElasticsearchDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticsearchServiceStub) UpgradeElasticsearchDomainAsync(ctx workflow.Context, input *elasticsearchservice.UpgradeElasticsearchDomainInput) *ElasticsearchserviceUpgradeElasticsearchDomainResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpgradeElasticsearchDomain, input)
    return &ElasticsearchserviceUpgradeElasticsearchDomainResult{Result: future}
}
