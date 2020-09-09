package awsclients

import (
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ServiceCatalogClient interface {
       AcceptPortfolioShare(ctx workflow.Context, input *servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error)
       AcceptPortfolioShareAsync(ctx workflow.Context, input *servicecatalog.AcceptPortfolioShareInput) *ServicecatalogAcceptPortfolioShareResult

       AssociateBudgetWithResource(ctx workflow.Context, input *servicecatalog.AssociateBudgetWithResourceInput) (*servicecatalog.AssociateBudgetWithResourceOutput, error)
       AssociateBudgetWithResourceAsync(ctx workflow.Context, input *servicecatalog.AssociateBudgetWithResourceInput) *ServicecatalogAssociateBudgetWithResourceResult

       AssociatePrincipalWithPortfolio(ctx workflow.Context, input *servicecatalog.AssociatePrincipalWithPortfolioInput) (*servicecatalog.AssociatePrincipalWithPortfolioOutput, error)
       AssociatePrincipalWithPortfolioAsync(ctx workflow.Context, input *servicecatalog.AssociatePrincipalWithPortfolioInput) *ServicecatalogAssociatePrincipalWithPortfolioResult

       AssociateProductWithPortfolio(ctx workflow.Context, input *servicecatalog.AssociateProductWithPortfolioInput) (*servicecatalog.AssociateProductWithPortfolioOutput, error)
       AssociateProductWithPortfolioAsync(ctx workflow.Context, input *servicecatalog.AssociateProductWithPortfolioInput) *ServicecatalogAssociateProductWithPortfolioResult

       AssociateServiceActionWithProvisioningArtifact(ctx workflow.Context, input *servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.AssociateServiceActionWithProvisioningArtifactOutput, error)
       AssociateServiceActionWithProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) *ServicecatalogAssociateServiceActionWithProvisioningArtifactResult

       AssociateTagOptionWithResource(ctx workflow.Context, input *servicecatalog.AssociateTagOptionWithResourceInput) (*servicecatalog.AssociateTagOptionWithResourceOutput, error)
       AssociateTagOptionWithResourceAsync(ctx workflow.Context, input *servicecatalog.AssociateTagOptionWithResourceInput) *ServicecatalogAssociateTagOptionWithResourceResult

       BatchAssociateServiceActionWithProvisioningArtifact(ctx workflow.Context, input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput, error)
       BatchAssociateServiceActionWithProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) *ServicecatalogBatchAssociateServiceActionWithProvisioningArtifactResult

       BatchDisassociateServiceActionFromProvisioningArtifact(ctx workflow.Context, input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput, error)
       BatchDisassociateServiceActionFromProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) *ServicecatalogBatchDisassociateServiceActionFromProvisioningArtifactResult

       CopyProduct(ctx workflow.Context, input *servicecatalog.CopyProductInput) (*servicecatalog.CopyProductOutput, error)
       CopyProductAsync(ctx workflow.Context, input *servicecatalog.CopyProductInput) *ServicecatalogCopyProductResult

       CreateConstraint(ctx workflow.Context, input *servicecatalog.CreateConstraintInput) (*servicecatalog.CreateConstraintOutput, error)
       CreateConstraintAsync(ctx workflow.Context, input *servicecatalog.CreateConstraintInput) *ServicecatalogCreateConstraintResult

       CreatePortfolio(ctx workflow.Context, input *servicecatalog.CreatePortfolioInput) (*servicecatalog.CreatePortfolioOutput, error)
       CreatePortfolioAsync(ctx workflow.Context, input *servicecatalog.CreatePortfolioInput) *ServicecatalogCreatePortfolioResult

       CreatePortfolioShare(ctx workflow.Context, input *servicecatalog.CreatePortfolioShareInput) (*servicecatalog.CreatePortfolioShareOutput, error)
       CreatePortfolioShareAsync(ctx workflow.Context, input *servicecatalog.CreatePortfolioShareInput) *ServicecatalogCreatePortfolioShareResult

       CreateProduct(ctx workflow.Context, input *servicecatalog.CreateProductInput) (*servicecatalog.CreateProductOutput, error)
       CreateProductAsync(ctx workflow.Context, input *servicecatalog.CreateProductInput) *ServicecatalogCreateProductResult

       CreateProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.CreateProvisionedProductPlanInput) (*servicecatalog.CreateProvisionedProductPlanOutput, error)
       CreateProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.CreateProvisionedProductPlanInput) *ServicecatalogCreateProvisionedProductPlanResult

       CreateProvisioningArtifact(ctx workflow.Context, input *servicecatalog.CreateProvisioningArtifactInput) (*servicecatalog.CreateProvisioningArtifactOutput, error)
       CreateProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.CreateProvisioningArtifactInput) *ServicecatalogCreateProvisioningArtifactResult

       CreateServiceAction(ctx workflow.Context, input *servicecatalog.CreateServiceActionInput) (*servicecatalog.CreateServiceActionOutput, error)
       CreateServiceActionAsync(ctx workflow.Context, input *servicecatalog.CreateServiceActionInput) *ServicecatalogCreateServiceActionResult

       CreateTagOption(ctx workflow.Context, input *servicecatalog.CreateTagOptionInput) (*servicecatalog.CreateTagOptionOutput, error)
       CreateTagOptionAsync(ctx workflow.Context, input *servicecatalog.CreateTagOptionInput) *ServicecatalogCreateTagOptionResult

       DeleteConstraint(ctx workflow.Context, input *servicecatalog.DeleteConstraintInput) (*servicecatalog.DeleteConstraintOutput, error)
       DeleteConstraintAsync(ctx workflow.Context, input *servicecatalog.DeleteConstraintInput) *ServicecatalogDeleteConstraintResult

       DeletePortfolio(ctx workflow.Context, input *servicecatalog.DeletePortfolioInput) (*servicecatalog.DeletePortfolioOutput, error)
       DeletePortfolioAsync(ctx workflow.Context, input *servicecatalog.DeletePortfolioInput) *ServicecatalogDeletePortfolioResult

       DeletePortfolioShare(ctx workflow.Context, input *servicecatalog.DeletePortfolioShareInput) (*servicecatalog.DeletePortfolioShareOutput, error)
       DeletePortfolioShareAsync(ctx workflow.Context, input *servicecatalog.DeletePortfolioShareInput) *ServicecatalogDeletePortfolioShareResult

       DeleteProduct(ctx workflow.Context, input *servicecatalog.DeleteProductInput) (*servicecatalog.DeleteProductOutput, error)
       DeleteProductAsync(ctx workflow.Context, input *servicecatalog.DeleteProductInput) *ServicecatalogDeleteProductResult

       DeleteProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.DeleteProvisionedProductPlanInput) (*servicecatalog.DeleteProvisionedProductPlanOutput, error)
       DeleteProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.DeleteProvisionedProductPlanInput) *ServicecatalogDeleteProvisionedProductPlanResult

       DeleteProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DeleteProvisioningArtifactInput) (*servicecatalog.DeleteProvisioningArtifactOutput, error)
       DeleteProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DeleteProvisioningArtifactInput) *ServicecatalogDeleteProvisioningArtifactResult

       DeleteServiceAction(ctx workflow.Context, input *servicecatalog.DeleteServiceActionInput) (*servicecatalog.DeleteServiceActionOutput, error)
       DeleteServiceActionAsync(ctx workflow.Context, input *servicecatalog.DeleteServiceActionInput) *ServicecatalogDeleteServiceActionResult

       DeleteTagOption(ctx workflow.Context, input *servicecatalog.DeleteTagOptionInput) (*servicecatalog.DeleteTagOptionOutput, error)
       DeleteTagOptionAsync(ctx workflow.Context, input *servicecatalog.DeleteTagOptionInput) *ServicecatalogDeleteTagOptionResult

       DescribeConstraint(ctx workflow.Context, input *servicecatalog.DescribeConstraintInput) (*servicecatalog.DescribeConstraintOutput, error)
       DescribeConstraintAsync(ctx workflow.Context, input *servicecatalog.DescribeConstraintInput) *ServicecatalogDescribeConstraintResult

       DescribeCopyProductStatus(ctx workflow.Context, input *servicecatalog.DescribeCopyProductStatusInput) (*servicecatalog.DescribeCopyProductStatusOutput, error)
       DescribeCopyProductStatusAsync(ctx workflow.Context, input *servicecatalog.DescribeCopyProductStatusInput) *ServicecatalogDescribeCopyProductStatusResult

       DescribePortfolio(ctx workflow.Context, input *servicecatalog.DescribePortfolioInput) (*servicecatalog.DescribePortfolioOutput, error)
       DescribePortfolioAsync(ctx workflow.Context, input *servicecatalog.DescribePortfolioInput) *ServicecatalogDescribePortfolioResult

       DescribePortfolioShareStatus(ctx workflow.Context, input *servicecatalog.DescribePortfolioShareStatusInput) (*servicecatalog.DescribePortfolioShareStatusOutput, error)
       DescribePortfolioShareStatusAsync(ctx workflow.Context, input *servicecatalog.DescribePortfolioShareStatusInput) *ServicecatalogDescribePortfolioShareStatusResult

       DescribeProduct(ctx workflow.Context, input *servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error)
       DescribeProductAsync(ctx workflow.Context, input *servicecatalog.DescribeProductInput) *ServicecatalogDescribeProductResult

       DescribeProductAsAdmin(ctx workflow.Context, input *servicecatalog.DescribeProductAsAdminInput) (*servicecatalog.DescribeProductAsAdminOutput, error)
       DescribeProductAsAdminAsync(ctx workflow.Context, input *servicecatalog.DescribeProductAsAdminInput) *ServicecatalogDescribeProductAsAdminResult

       DescribeProductView(ctx workflow.Context, input *servicecatalog.DescribeProductViewInput) (*servicecatalog.DescribeProductViewOutput, error)
       DescribeProductViewAsync(ctx workflow.Context, input *servicecatalog.DescribeProductViewInput) *ServicecatalogDescribeProductViewResult

       DescribeProvisionedProduct(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductInput) (*servicecatalog.DescribeProvisionedProductOutput, error)
       DescribeProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductInput) *ServicecatalogDescribeProvisionedProductResult

       DescribeProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductPlanInput) (*servicecatalog.DescribeProvisionedProductPlanOutput, error)
       DescribeProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductPlanInput) *ServicecatalogDescribeProvisionedProductPlanResult

       DescribeProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DescribeProvisioningArtifactInput) (*servicecatalog.DescribeProvisioningArtifactOutput, error)
       DescribeProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisioningArtifactInput) *ServicecatalogDescribeProvisioningArtifactResult

       DescribeProvisioningParameters(ctx workflow.Context, input *servicecatalog.DescribeProvisioningParametersInput) (*servicecatalog.DescribeProvisioningParametersOutput, error)
       DescribeProvisioningParametersAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisioningParametersInput) *ServicecatalogDescribeProvisioningParametersResult

       DescribeRecord(ctx workflow.Context, input *servicecatalog.DescribeRecordInput) (*servicecatalog.DescribeRecordOutput, error)
       DescribeRecordAsync(ctx workflow.Context, input *servicecatalog.DescribeRecordInput) *ServicecatalogDescribeRecordResult

       DescribeServiceAction(ctx workflow.Context, input *servicecatalog.DescribeServiceActionInput) (*servicecatalog.DescribeServiceActionOutput, error)
       DescribeServiceActionAsync(ctx workflow.Context, input *servicecatalog.DescribeServiceActionInput) *ServicecatalogDescribeServiceActionResult

       DescribeServiceActionExecutionParameters(ctx workflow.Context, input *servicecatalog.DescribeServiceActionExecutionParametersInput) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error)
       DescribeServiceActionExecutionParametersAsync(ctx workflow.Context, input *servicecatalog.DescribeServiceActionExecutionParametersInput) *ServicecatalogDescribeServiceActionExecutionParametersResult

       DescribeTagOption(ctx workflow.Context, input *servicecatalog.DescribeTagOptionInput) (*servicecatalog.DescribeTagOptionOutput, error)
       DescribeTagOptionAsync(ctx workflow.Context, input *servicecatalog.DescribeTagOptionInput) *ServicecatalogDescribeTagOptionResult

       DisableAWSOrganizationsAccess(ctx workflow.Context, input *servicecatalog.DisableAWSOrganizationsAccessInput) (*servicecatalog.DisableAWSOrganizationsAccessOutput, error)
       DisableAWSOrganizationsAccessAsync(ctx workflow.Context, input *servicecatalog.DisableAWSOrganizationsAccessInput) *ServicecatalogDisableAWSOrganizationsAccessResult

       DisassociateBudgetFromResource(ctx workflow.Context, input *servicecatalog.DisassociateBudgetFromResourceInput) (*servicecatalog.DisassociateBudgetFromResourceOutput, error)
       DisassociateBudgetFromResourceAsync(ctx workflow.Context, input *servicecatalog.DisassociateBudgetFromResourceInput) *ServicecatalogDisassociateBudgetFromResourceResult

       DisassociatePrincipalFromPortfolio(ctx workflow.Context, input *servicecatalog.DisassociatePrincipalFromPortfolioInput) (*servicecatalog.DisassociatePrincipalFromPortfolioOutput, error)
       DisassociatePrincipalFromPortfolioAsync(ctx workflow.Context, input *servicecatalog.DisassociatePrincipalFromPortfolioInput) *ServicecatalogDisassociatePrincipalFromPortfolioResult

       DisassociateProductFromPortfolio(ctx workflow.Context, input *servicecatalog.DisassociateProductFromPortfolioInput) (*servicecatalog.DisassociateProductFromPortfolioOutput, error)
       DisassociateProductFromPortfolioAsync(ctx workflow.Context, input *servicecatalog.DisassociateProductFromPortfolioInput) *ServicecatalogDisassociateProductFromPortfolioResult

       DisassociateServiceActionFromProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.DisassociateServiceActionFromProvisioningArtifactOutput, error)
       DisassociateServiceActionFromProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) *ServicecatalogDisassociateServiceActionFromProvisioningArtifactResult

       DisassociateTagOptionFromResource(ctx workflow.Context, input *servicecatalog.DisassociateTagOptionFromResourceInput) (*servicecatalog.DisassociateTagOptionFromResourceOutput, error)
       DisassociateTagOptionFromResourceAsync(ctx workflow.Context, input *servicecatalog.DisassociateTagOptionFromResourceInput) *ServicecatalogDisassociateTagOptionFromResourceResult

       EnableAWSOrganizationsAccess(ctx workflow.Context, input *servicecatalog.EnableAWSOrganizationsAccessInput) (*servicecatalog.EnableAWSOrganizationsAccessOutput, error)
       EnableAWSOrganizationsAccessAsync(ctx workflow.Context, input *servicecatalog.EnableAWSOrganizationsAccessInput) *ServicecatalogEnableAWSOrganizationsAccessResult

       ExecuteProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductPlanInput) (*servicecatalog.ExecuteProvisionedProductPlanOutput, error)
       ExecuteProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductPlanInput) *ServicecatalogExecuteProvisionedProductPlanResult

       ExecuteProvisionedProductServiceAction(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductServiceActionInput) (*servicecatalog.ExecuteProvisionedProductServiceActionOutput, error)
       ExecuteProvisionedProductServiceActionAsync(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductServiceActionInput) *ServicecatalogExecuteProvisionedProductServiceActionResult

       GetAWSOrganizationsAccessStatus(ctx workflow.Context, input *servicecatalog.GetAWSOrganizationsAccessStatusInput) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error)
       GetAWSOrganizationsAccessStatusAsync(ctx workflow.Context, input *servicecatalog.GetAWSOrganizationsAccessStatusInput) *ServicecatalogGetAWSOrganizationsAccessStatusResult

       ListAcceptedPortfolioShares(ctx workflow.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error)
       ListAcceptedPortfolioSharesAsync(ctx workflow.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput) *ServicecatalogListAcceptedPortfolioSharesResult

       ListBudgetsForResource(ctx workflow.Context, input *servicecatalog.ListBudgetsForResourceInput) (*servicecatalog.ListBudgetsForResourceOutput, error)
       ListBudgetsForResourceAsync(ctx workflow.Context, input *servicecatalog.ListBudgetsForResourceInput) *ServicecatalogListBudgetsForResourceResult

       ListConstraintsForPortfolio(ctx workflow.Context, input *servicecatalog.ListConstraintsForPortfolioInput) (*servicecatalog.ListConstraintsForPortfolioOutput, error)
       ListConstraintsForPortfolioAsync(ctx workflow.Context, input *servicecatalog.ListConstraintsForPortfolioInput) *ServicecatalogListConstraintsForPortfolioResult

       ListLaunchPaths(ctx workflow.Context, input *servicecatalog.ListLaunchPathsInput) (*servicecatalog.ListLaunchPathsOutput, error)
       ListLaunchPathsAsync(ctx workflow.Context, input *servicecatalog.ListLaunchPathsInput) *ServicecatalogListLaunchPathsResult

       ListOrganizationPortfolioAccess(ctx workflow.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error)
       ListOrganizationPortfolioAccessAsync(ctx workflow.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput) *ServicecatalogListOrganizationPortfolioAccessResult

       ListPortfolioAccess(ctx workflow.Context, input *servicecatalog.ListPortfolioAccessInput) (*servicecatalog.ListPortfolioAccessOutput, error)
       ListPortfolioAccessAsync(ctx workflow.Context, input *servicecatalog.ListPortfolioAccessInput) *ServicecatalogListPortfolioAccessResult

       ListPortfolios(ctx workflow.Context, input *servicecatalog.ListPortfoliosInput) (*servicecatalog.ListPortfoliosOutput, error)
       ListPortfoliosAsync(ctx workflow.Context, input *servicecatalog.ListPortfoliosInput) *ServicecatalogListPortfoliosResult

       ListPortfoliosForProduct(ctx workflow.Context, input *servicecatalog.ListPortfoliosForProductInput) (*servicecatalog.ListPortfoliosForProductOutput, error)
       ListPortfoliosForProductAsync(ctx workflow.Context, input *servicecatalog.ListPortfoliosForProductInput) *ServicecatalogListPortfoliosForProductResult

       ListPrincipalsForPortfolio(ctx workflow.Context, input *servicecatalog.ListPrincipalsForPortfolioInput) (*servicecatalog.ListPrincipalsForPortfolioOutput, error)
       ListPrincipalsForPortfolioAsync(ctx workflow.Context, input *servicecatalog.ListPrincipalsForPortfolioInput) *ServicecatalogListPrincipalsForPortfolioResult

       ListProvisionedProductPlans(ctx workflow.Context, input *servicecatalog.ListProvisionedProductPlansInput) (*servicecatalog.ListProvisionedProductPlansOutput, error)
       ListProvisionedProductPlansAsync(ctx workflow.Context, input *servicecatalog.ListProvisionedProductPlansInput) *ServicecatalogListProvisionedProductPlansResult

       ListProvisioningArtifacts(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsInput) (*servicecatalog.ListProvisioningArtifactsOutput, error)
       ListProvisioningArtifactsAsync(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsInput) *ServicecatalogListProvisioningArtifactsResult

       ListProvisioningArtifactsForServiceAction(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error)
       ListProvisioningArtifactsForServiceActionAsync(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) *ServicecatalogListProvisioningArtifactsForServiceActionResult

       ListRecordHistory(ctx workflow.Context, input *servicecatalog.ListRecordHistoryInput) (*servicecatalog.ListRecordHistoryOutput, error)
       ListRecordHistoryAsync(ctx workflow.Context, input *servicecatalog.ListRecordHistoryInput) *ServicecatalogListRecordHistoryResult

       ListResourcesForTagOption(ctx workflow.Context, input *servicecatalog.ListResourcesForTagOptionInput) (*servicecatalog.ListResourcesForTagOptionOutput, error)
       ListResourcesForTagOptionAsync(ctx workflow.Context, input *servicecatalog.ListResourcesForTagOptionInput) *ServicecatalogListResourcesForTagOptionResult

       ListServiceActions(ctx workflow.Context, input *servicecatalog.ListServiceActionsInput) (*servicecatalog.ListServiceActionsOutput, error)
       ListServiceActionsAsync(ctx workflow.Context, input *servicecatalog.ListServiceActionsInput) *ServicecatalogListServiceActionsResult

       ListServiceActionsForProvisioningArtifact(ctx workflow.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error)
       ListServiceActionsForProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) *ServicecatalogListServiceActionsForProvisioningArtifactResult

       ListStackInstancesForProvisionedProduct(ctx workflow.Context, input *servicecatalog.ListStackInstancesForProvisionedProductInput) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error)
       ListStackInstancesForProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.ListStackInstancesForProvisionedProductInput) *ServicecatalogListStackInstancesForProvisionedProductResult

       ListTagOptions(ctx workflow.Context, input *servicecatalog.ListTagOptionsInput) (*servicecatalog.ListTagOptionsOutput, error)
       ListTagOptionsAsync(ctx workflow.Context, input *servicecatalog.ListTagOptionsInput) *ServicecatalogListTagOptionsResult

       ProvisionProduct(ctx workflow.Context, input *servicecatalog.ProvisionProductInput) (*servicecatalog.ProvisionProductOutput, error)
       ProvisionProductAsync(ctx workflow.Context, input *servicecatalog.ProvisionProductInput) *ServicecatalogProvisionProductResult

       RejectPortfolioShare(ctx workflow.Context, input *servicecatalog.RejectPortfolioShareInput) (*servicecatalog.RejectPortfolioShareOutput, error)
       RejectPortfolioShareAsync(ctx workflow.Context, input *servicecatalog.RejectPortfolioShareInput) *ServicecatalogRejectPortfolioShareResult

       ScanProvisionedProducts(ctx workflow.Context, input *servicecatalog.ScanProvisionedProductsInput) (*servicecatalog.ScanProvisionedProductsOutput, error)
       ScanProvisionedProductsAsync(ctx workflow.Context, input *servicecatalog.ScanProvisionedProductsInput) *ServicecatalogScanProvisionedProductsResult

       SearchProducts(ctx workflow.Context, input *servicecatalog.SearchProductsInput) (*servicecatalog.SearchProductsOutput, error)
       SearchProductsAsync(ctx workflow.Context, input *servicecatalog.SearchProductsInput) *ServicecatalogSearchProductsResult

       SearchProductsAsAdmin(ctx workflow.Context, input *servicecatalog.SearchProductsAsAdminInput) (*servicecatalog.SearchProductsAsAdminOutput, error)
       SearchProductsAsAdminAsync(ctx workflow.Context, input *servicecatalog.SearchProductsAsAdminInput) *ServicecatalogSearchProductsAsAdminResult

       SearchProvisionedProducts(ctx workflow.Context, input *servicecatalog.SearchProvisionedProductsInput) (*servicecatalog.SearchProvisionedProductsOutput, error)
       SearchProvisionedProductsAsync(ctx workflow.Context, input *servicecatalog.SearchProvisionedProductsInput) *ServicecatalogSearchProvisionedProductsResult

       TerminateProvisionedProduct(ctx workflow.Context, input *servicecatalog.TerminateProvisionedProductInput) (*servicecatalog.TerminateProvisionedProductOutput, error)
       TerminateProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.TerminateProvisionedProductInput) *ServicecatalogTerminateProvisionedProductResult

       UpdateConstraint(ctx workflow.Context, input *servicecatalog.UpdateConstraintInput) (*servicecatalog.UpdateConstraintOutput, error)
       UpdateConstraintAsync(ctx workflow.Context, input *servicecatalog.UpdateConstraintInput) *ServicecatalogUpdateConstraintResult

       UpdatePortfolio(ctx workflow.Context, input *servicecatalog.UpdatePortfolioInput) (*servicecatalog.UpdatePortfolioOutput, error)
       UpdatePortfolioAsync(ctx workflow.Context, input *servicecatalog.UpdatePortfolioInput) *ServicecatalogUpdatePortfolioResult

       UpdateProduct(ctx workflow.Context, input *servicecatalog.UpdateProductInput) (*servicecatalog.UpdateProductOutput, error)
       UpdateProductAsync(ctx workflow.Context, input *servicecatalog.UpdateProductInput) *ServicecatalogUpdateProductResult

       UpdateProvisionedProduct(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductInput) (*servicecatalog.UpdateProvisionedProductOutput, error)
       UpdateProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductInput) *ServicecatalogUpdateProvisionedProductResult

       UpdateProvisionedProductProperties(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductPropertiesInput) (*servicecatalog.UpdateProvisionedProductPropertiesOutput, error)
       UpdateProvisionedProductPropertiesAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductPropertiesInput) *ServicecatalogUpdateProvisionedProductPropertiesResult

       UpdateProvisioningArtifact(ctx workflow.Context, input *servicecatalog.UpdateProvisioningArtifactInput) (*servicecatalog.UpdateProvisioningArtifactOutput, error)
       UpdateProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisioningArtifactInput) *ServicecatalogUpdateProvisioningArtifactResult

       UpdateServiceAction(ctx workflow.Context, input *servicecatalog.UpdateServiceActionInput) (*servicecatalog.UpdateServiceActionOutput, error)
       UpdateServiceActionAsync(ctx workflow.Context, input *servicecatalog.UpdateServiceActionInput) *ServicecatalogUpdateServiceActionResult

       UpdateTagOption(ctx workflow.Context, input *servicecatalog.UpdateTagOptionInput) (*servicecatalog.UpdateTagOptionOutput, error)
       UpdateTagOptionAsync(ctx workflow.Context, input *servicecatalog.UpdateTagOptionInput) *ServicecatalogUpdateTagOptionResult
}

type ServicecatalogAcceptPortfolioShareResult struct {
	Result workflow.Future
}

func (r *ServicecatalogAcceptPortfolioShareResult) Get(ctx workflow.Context) (*servicecatalog.AcceptPortfolioShareOutput, error) {
    var output servicecatalog.AcceptPortfolioShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogAssociateBudgetWithResourceResult struct {
	Result workflow.Future
}

func (r *ServicecatalogAssociateBudgetWithResourceResult) Get(ctx workflow.Context) (*servicecatalog.AssociateBudgetWithResourceOutput, error) {
    var output servicecatalog.AssociateBudgetWithResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogAssociatePrincipalWithPortfolioResult struct {
	Result workflow.Future
}

func (r *ServicecatalogAssociatePrincipalWithPortfolioResult) Get(ctx workflow.Context) (*servicecatalog.AssociatePrincipalWithPortfolioOutput, error) {
    var output servicecatalog.AssociatePrincipalWithPortfolioOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogAssociateProductWithPortfolioResult struct {
	Result workflow.Future
}

func (r *ServicecatalogAssociateProductWithPortfolioResult) Get(ctx workflow.Context) (*servicecatalog.AssociateProductWithPortfolioOutput, error) {
    var output servicecatalog.AssociateProductWithPortfolioOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogAssociateServiceActionWithProvisioningArtifactResult struct {
	Result workflow.Future
}

func (r *ServicecatalogAssociateServiceActionWithProvisioningArtifactResult) Get(ctx workflow.Context) (*servicecatalog.AssociateServiceActionWithProvisioningArtifactOutput, error) {
    var output servicecatalog.AssociateServiceActionWithProvisioningArtifactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogAssociateTagOptionWithResourceResult struct {
	Result workflow.Future
}

func (r *ServicecatalogAssociateTagOptionWithResourceResult) Get(ctx workflow.Context) (*servicecatalog.AssociateTagOptionWithResourceOutput, error) {
    var output servicecatalog.AssociateTagOptionWithResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogBatchAssociateServiceActionWithProvisioningArtifactResult struct {
	Result workflow.Future
}

func (r *ServicecatalogBatchAssociateServiceActionWithProvisioningArtifactResult) Get(ctx workflow.Context) (*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput, error) {
    var output servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogBatchDisassociateServiceActionFromProvisioningArtifactResult struct {
	Result workflow.Future
}

func (r *ServicecatalogBatchDisassociateServiceActionFromProvisioningArtifactResult) Get(ctx workflow.Context) (*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput, error) {
    var output servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogCopyProductResult struct {
	Result workflow.Future
}

func (r *ServicecatalogCopyProductResult) Get(ctx workflow.Context) (*servicecatalog.CopyProductOutput, error) {
    var output servicecatalog.CopyProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogCreateConstraintResult struct {
	Result workflow.Future
}

func (r *ServicecatalogCreateConstraintResult) Get(ctx workflow.Context) (*servicecatalog.CreateConstraintOutput, error) {
    var output servicecatalog.CreateConstraintOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogCreatePortfolioResult struct {
	Result workflow.Future
}

func (r *ServicecatalogCreatePortfolioResult) Get(ctx workflow.Context) (*servicecatalog.CreatePortfolioOutput, error) {
    var output servicecatalog.CreatePortfolioOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogCreatePortfolioShareResult struct {
	Result workflow.Future
}

func (r *ServicecatalogCreatePortfolioShareResult) Get(ctx workflow.Context) (*servicecatalog.CreatePortfolioShareOutput, error) {
    var output servicecatalog.CreatePortfolioShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogCreateProductResult struct {
	Result workflow.Future
}

func (r *ServicecatalogCreateProductResult) Get(ctx workflow.Context) (*servicecatalog.CreateProductOutput, error) {
    var output servicecatalog.CreateProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogCreateProvisionedProductPlanResult struct {
	Result workflow.Future
}

func (r *ServicecatalogCreateProvisionedProductPlanResult) Get(ctx workflow.Context) (*servicecatalog.CreateProvisionedProductPlanOutput, error) {
    var output servicecatalog.CreateProvisionedProductPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogCreateProvisioningArtifactResult struct {
	Result workflow.Future
}

func (r *ServicecatalogCreateProvisioningArtifactResult) Get(ctx workflow.Context) (*servicecatalog.CreateProvisioningArtifactOutput, error) {
    var output servicecatalog.CreateProvisioningArtifactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogCreateServiceActionResult struct {
	Result workflow.Future
}

func (r *ServicecatalogCreateServiceActionResult) Get(ctx workflow.Context) (*servicecatalog.CreateServiceActionOutput, error) {
    var output servicecatalog.CreateServiceActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogCreateTagOptionResult struct {
	Result workflow.Future
}

func (r *ServicecatalogCreateTagOptionResult) Get(ctx workflow.Context) (*servicecatalog.CreateTagOptionOutput, error) {
    var output servicecatalog.CreateTagOptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDeleteConstraintResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDeleteConstraintResult) Get(ctx workflow.Context) (*servicecatalog.DeleteConstraintOutput, error) {
    var output servicecatalog.DeleteConstraintOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDeletePortfolioResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDeletePortfolioResult) Get(ctx workflow.Context) (*servicecatalog.DeletePortfolioOutput, error) {
    var output servicecatalog.DeletePortfolioOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDeletePortfolioShareResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDeletePortfolioShareResult) Get(ctx workflow.Context) (*servicecatalog.DeletePortfolioShareOutput, error) {
    var output servicecatalog.DeletePortfolioShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDeleteProductResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDeleteProductResult) Get(ctx workflow.Context) (*servicecatalog.DeleteProductOutput, error) {
    var output servicecatalog.DeleteProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDeleteProvisionedProductPlanResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDeleteProvisionedProductPlanResult) Get(ctx workflow.Context) (*servicecatalog.DeleteProvisionedProductPlanOutput, error) {
    var output servicecatalog.DeleteProvisionedProductPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDeleteProvisioningArtifactResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDeleteProvisioningArtifactResult) Get(ctx workflow.Context) (*servicecatalog.DeleteProvisioningArtifactOutput, error) {
    var output servicecatalog.DeleteProvisioningArtifactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDeleteServiceActionResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDeleteServiceActionResult) Get(ctx workflow.Context) (*servicecatalog.DeleteServiceActionOutput, error) {
    var output servicecatalog.DeleteServiceActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDeleteTagOptionResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDeleteTagOptionResult) Get(ctx workflow.Context) (*servicecatalog.DeleteTagOptionOutput, error) {
    var output servicecatalog.DeleteTagOptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeConstraintResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeConstraintResult) Get(ctx workflow.Context) (*servicecatalog.DescribeConstraintOutput, error) {
    var output servicecatalog.DescribeConstraintOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeCopyProductStatusResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeCopyProductStatusResult) Get(ctx workflow.Context) (*servicecatalog.DescribeCopyProductStatusOutput, error) {
    var output servicecatalog.DescribeCopyProductStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribePortfolioResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribePortfolioResult) Get(ctx workflow.Context) (*servicecatalog.DescribePortfolioOutput, error) {
    var output servicecatalog.DescribePortfolioOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribePortfolioShareStatusResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribePortfolioShareStatusResult) Get(ctx workflow.Context) (*servicecatalog.DescribePortfolioShareStatusOutput, error) {
    var output servicecatalog.DescribePortfolioShareStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeProductResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeProductResult) Get(ctx workflow.Context) (*servicecatalog.DescribeProductOutput, error) {
    var output servicecatalog.DescribeProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeProductAsAdminResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeProductAsAdminResult) Get(ctx workflow.Context) (*servicecatalog.DescribeProductAsAdminOutput, error) {
    var output servicecatalog.DescribeProductAsAdminOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeProductViewResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeProductViewResult) Get(ctx workflow.Context) (*servicecatalog.DescribeProductViewOutput, error) {
    var output servicecatalog.DescribeProductViewOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeProvisionedProductResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeProvisionedProductResult) Get(ctx workflow.Context) (*servicecatalog.DescribeProvisionedProductOutput, error) {
    var output servicecatalog.DescribeProvisionedProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeProvisionedProductPlanResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeProvisionedProductPlanResult) Get(ctx workflow.Context) (*servicecatalog.DescribeProvisionedProductPlanOutput, error) {
    var output servicecatalog.DescribeProvisionedProductPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeProvisioningArtifactResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeProvisioningArtifactResult) Get(ctx workflow.Context) (*servicecatalog.DescribeProvisioningArtifactOutput, error) {
    var output servicecatalog.DescribeProvisioningArtifactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeProvisioningParametersResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeProvisioningParametersResult) Get(ctx workflow.Context) (*servicecatalog.DescribeProvisioningParametersOutput, error) {
    var output servicecatalog.DescribeProvisioningParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeRecordResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeRecordResult) Get(ctx workflow.Context) (*servicecatalog.DescribeRecordOutput, error) {
    var output servicecatalog.DescribeRecordOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeServiceActionResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeServiceActionResult) Get(ctx workflow.Context) (*servicecatalog.DescribeServiceActionOutput, error) {
    var output servicecatalog.DescribeServiceActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeServiceActionExecutionParametersResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeServiceActionExecutionParametersResult) Get(ctx workflow.Context) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error) {
    var output servicecatalog.DescribeServiceActionExecutionParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDescribeTagOptionResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDescribeTagOptionResult) Get(ctx workflow.Context) (*servicecatalog.DescribeTagOptionOutput, error) {
    var output servicecatalog.DescribeTagOptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDisableAWSOrganizationsAccessResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDisableAWSOrganizationsAccessResult) Get(ctx workflow.Context) (*servicecatalog.DisableAWSOrganizationsAccessOutput, error) {
    var output servicecatalog.DisableAWSOrganizationsAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDisassociateBudgetFromResourceResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDisassociateBudgetFromResourceResult) Get(ctx workflow.Context) (*servicecatalog.DisassociateBudgetFromResourceOutput, error) {
    var output servicecatalog.DisassociateBudgetFromResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDisassociatePrincipalFromPortfolioResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDisassociatePrincipalFromPortfolioResult) Get(ctx workflow.Context) (*servicecatalog.DisassociatePrincipalFromPortfolioOutput, error) {
    var output servicecatalog.DisassociatePrincipalFromPortfolioOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDisassociateProductFromPortfolioResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDisassociateProductFromPortfolioResult) Get(ctx workflow.Context) (*servicecatalog.DisassociateProductFromPortfolioOutput, error) {
    var output servicecatalog.DisassociateProductFromPortfolioOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDisassociateServiceActionFromProvisioningArtifactResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDisassociateServiceActionFromProvisioningArtifactResult) Get(ctx workflow.Context) (*servicecatalog.DisassociateServiceActionFromProvisioningArtifactOutput, error) {
    var output servicecatalog.DisassociateServiceActionFromProvisioningArtifactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogDisassociateTagOptionFromResourceResult struct {
	Result workflow.Future
}

func (r *ServicecatalogDisassociateTagOptionFromResourceResult) Get(ctx workflow.Context) (*servicecatalog.DisassociateTagOptionFromResourceOutput, error) {
    var output servicecatalog.DisassociateTagOptionFromResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogEnableAWSOrganizationsAccessResult struct {
	Result workflow.Future
}

func (r *ServicecatalogEnableAWSOrganizationsAccessResult) Get(ctx workflow.Context) (*servicecatalog.EnableAWSOrganizationsAccessOutput, error) {
    var output servicecatalog.EnableAWSOrganizationsAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogExecuteProvisionedProductPlanResult struct {
	Result workflow.Future
}

func (r *ServicecatalogExecuteProvisionedProductPlanResult) Get(ctx workflow.Context) (*servicecatalog.ExecuteProvisionedProductPlanOutput, error) {
    var output servicecatalog.ExecuteProvisionedProductPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogExecuteProvisionedProductServiceActionResult struct {
	Result workflow.Future
}

func (r *ServicecatalogExecuteProvisionedProductServiceActionResult) Get(ctx workflow.Context) (*servicecatalog.ExecuteProvisionedProductServiceActionOutput, error) {
    var output servicecatalog.ExecuteProvisionedProductServiceActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogGetAWSOrganizationsAccessStatusResult struct {
	Result workflow.Future
}

func (r *ServicecatalogGetAWSOrganizationsAccessStatusResult) Get(ctx workflow.Context) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error) {
    var output servicecatalog.GetAWSOrganizationsAccessStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListAcceptedPortfolioSharesResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListAcceptedPortfolioSharesResult) Get(ctx workflow.Context) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error) {
    var output servicecatalog.ListAcceptedPortfolioSharesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListBudgetsForResourceResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListBudgetsForResourceResult) Get(ctx workflow.Context) (*servicecatalog.ListBudgetsForResourceOutput, error) {
    var output servicecatalog.ListBudgetsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListConstraintsForPortfolioResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListConstraintsForPortfolioResult) Get(ctx workflow.Context) (*servicecatalog.ListConstraintsForPortfolioOutput, error) {
    var output servicecatalog.ListConstraintsForPortfolioOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListLaunchPathsResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListLaunchPathsResult) Get(ctx workflow.Context) (*servicecatalog.ListLaunchPathsOutput, error) {
    var output servicecatalog.ListLaunchPathsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListOrganizationPortfolioAccessResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListOrganizationPortfolioAccessResult) Get(ctx workflow.Context) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error) {
    var output servicecatalog.ListOrganizationPortfolioAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListPortfolioAccessResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListPortfolioAccessResult) Get(ctx workflow.Context) (*servicecatalog.ListPortfolioAccessOutput, error) {
    var output servicecatalog.ListPortfolioAccessOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListPortfoliosResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListPortfoliosResult) Get(ctx workflow.Context) (*servicecatalog.ListPortfoliosOutput, error) {
    var output servicecatalog.ListPortfoliosOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListPortfoliosForProductResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListPortfoliosForProductResult) Get(ctx workflow.Context) (*servicecatalog.ListPortfoliosForProductOutput, error) {
    var output servicecatalog.ListPortfoliosForProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListPrincipalsForPortfolioResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListPrincipalsForPortfolioResult) Get(ctx workflow.Context) (*servicecatalog.ListPrincipalsForPortfolioOutput, error) {
    var output servicecatalog.ListPrincipalsForPortfolioOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListProvisionedProductPlansResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListProvisionedProductPlansResult) Get(ctx workflow.Context) (*servicecatalog.ListProvisionedProductPlansOutput, error) {
    var output servicecatalog.ListProvisionedProductPlansOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListProvisioningArtifactsResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListProvisioningArtifactsResult) Get(ctx workflow.Context) (*servicecatalog.ListProvisioningArtifactsOutput, error) {
    var output servicecatalog.ListProvisioningArtifactsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListProvisioningArtifactsForServiceActionResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListProvisioningArtifactsForServiceActionResult) Get(ctx workflow.Context) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error) {
    var output servicecatalog.ListProvisioningArtifactsForServiceActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListRecordHistoryResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListRecordHistoryResult) Get(ctx workflow.Context) (*servicecatalog.ListRecordHistoryOutput, error) {
    var output servicecatalog.ListRecordHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListResourcesForTagOptionResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListResourcesForTagOptionResult) Get(ctx workflow.Context) (*servicecatalog.ListResourcesForTagOptionOutput, error) {
    var output servicecatalog.ListResourcesForTagOptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListServiceActionsResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListServiceActionsResult) Get(ctx workflow.Context) (*servicecatalog.ListServiceActionsOutput, error) {
    var output servicecatalog.ListServiceActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListServiceActionsForProvisioningArtifactResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListServiceActionsForProvisioningArtifactResult) Get(ctx workflow.Context) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error) {
    var output servicecatalog.ListServiceActionsForProvisioningArtifactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListStackInstancesForProvisionedProductResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListStackInstancesForProvisionedProductResult) Get(ctx workflow.Context) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error) {
    var output servicecatalog.ListStackInstancesForProvisionedProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogListTagOptionsResult struct {
	Result workflow.Future
}

func (r *ServicecatalogListTagOptionsResult) Get(ctx workflow.Context) (*servicecatalog.ListTagOptionsOutput, error) {
    var output servicecatalog.ListTagOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogProvisionProductResult struct {
	Result workflow.Future
}

func (r *ServicecatalogProvisionProductResult) Get(ctx workflow.Context) (*servicecatalog.ProvisionProductOutput, error) {
    var output servicecatalog.ProvisionProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogRejectPortfolioShareResult struct {
	Result workflow.Future
}

func (r *ServicecatalogRejectPortfolioShareResult) Get(ctx workflow.Context) (*servicecatalog.RejectPortfolioShareOutput, error) {
    var output servicecatalog.RejectPortfolioShareOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogScanProvisionedProductsResult struct {
	Result workflow.Future
}

func (r *ServicecatalogScanProvisionedProductsResult) Get(ctx workflow.Context) (*servicecatalog.ScanProvisionedProductsOutput, error) {
    var output servicecatalog.ScanProvisionedProductsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogSearchProductsResult struct {
	Result workflow.Future
}

func (r *ServicecatalogSearchProductsResult) Get(ctx workflow.Context) (*servicecatalog.SearchProductsOutput, error) {
    var output servicecatalog.SearchProductsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogSearchProductsAsAdminResult struct {
	Result workflow.Future
}

func (r *ServicecatalogSearchProductsAsAdminResult) Get(ctx workflow.Context) (*servicecatalog.SearchProductsAsAdminOutput, error) {
    var output servicecatalog.SearchProductsAsAdminOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogSearchProvisionedProductsResult struct {
	Result workflow.Future
}

func (r *ServicecatalogSearchProvisionedProductsResult) Get(ctx workflow.Context) (*servicecatalog.SearchProvisionedProductsOutput, error) {
    var output servicecatalog.SearchProvisionedProductsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogTerminateProvisionedProductResult struct {
	Result workflow.Future
}

func (r *ServicecatalogTerminateProvisionedProductResult) Get(ctx workflow.Context) (*servicecatalog.TerminateProvisionedProductOutput, error) {
    var output servicecatalog.TerminateProvisionedProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogUpdateConstraintResult struct {
	Result workflow.Future
}

func (r *ServicecatalogUpdateConstraintResult) Get(ctx workflow.Context) (*servicecatalog.UpdateConstraintOutput, error) {
    var output servicecatalog.UpdateConstraintOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogUpdatePortfolioResult struct {
	Result workflow.Future
}

func (r *ServicecatalogUpdatePortfolioResult) Get(ctx workflow.Context) (*servicecatalog.UpdatePortfolioOutput, error) {
    var output servicecatalog.UpdatePortfolioOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogUpdateProductResult struct {
	Result workflow.Future
}

func (r *ServicecatalogUpdateProductResult) Get(ctx workflow.Context) (*servicecatalog.UpdateProductOutput, error) {
    var output servicecatalog.UpdateProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogUpdateProvisionedProductResult struct {
	Result workflow.Future
}

func (r *ServicecatalogUpdateProvisionedProductResult) Get(ctx workflow.Context) (*servicecatalog.UpdateProvisionedProductOutput, error) {
    var output servicecatalog.UpdateProvisionedProductOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogUpdateProvisionedProductPropertiesResult struct {
	Result workflow.Future
}

func (r *ServicecatalogUpdateProvisionedProductPropertiesResult) Get(ctx workflow.Context) (*servicecatalog.UpdateProvisionedProductPropertiesOutput, error) {
    var output servicecatalog.UpdateProvisionedProductPropertiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogUpdateProvisioningArtifactResult struct {
	Result workflow.Future
}

func (r *ServicecatalogUpdateProvisioningArtifactResult) Get(ctx workflow.Context) (*servicecatalog.UpdateProvisioningArtifactOutput, error) {
    var output servicecatalog.UpdateProvisioningArtifactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogUpdateServiceActionResult struct {
	Result workflow.Future
}

func (r *ServicecatalogUpdateServiceActionResult) Get(ctx workflow.Context) (*servicecatalog.UpdateServiceActionOutput, error) {
    var output servicecatalog.UpdateServiceActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServicecatalogUpdateTagOptionResult struct {
	Result workflow.Future
}

func (r *ServicecatalogUpdateTagOptionResult) Get(ctx workflow.Context) (*servicecatalog.UpdateTagOptionOutput, error) {
    var output servicecatalog.UpdateTagOptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ServiceCatalogStub struct {
    activities awsactivities.ServiceCatalogActivities
}

func NewServiceCatalogStub() ServiceCatalogClient {
    return &ServiceCatalogStub{}
}

func (a *ServiceCatalogStub) AcceptPortfolioShare(ctx workflow.Context, input *servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error) {
    var output servicecatalog.AcceptPortfolioShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptPortfolioShare, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) AcceptPortfolioShareAsync(ctx workflow.Context, input *servicecatalog.AcceptPortfolioShareInput) *ServicecatalogAcceptPortfolioShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptPortfolioShare, input)
    return &ServicecatalogAcceptPortfolioShareResult{Result: future}
}

func (a *ServiceCatalogStub) AssociateBudgetWithResource(ctx workflow.Context, input *servicecatalog.AssociateBudgetWithResourceInput) (*servicecatalog.AssociateBudgetWithResourceOutput, error) {
    var output servicecatalog.AssociateBudgetWithResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateBudgetWithResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) AssociateBudgetWithResourceAsync(ctx workflow.Context, input *servicecatalog.AssociateBudgetWithResourceInput) *ServicecatalogAssociateBudgetWithResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateBudgetWithResource, input)
    return &ServicecatalogAssociateBudgetWithResourceResult{Result: future}
}

func (a *ServiceCatalogStub) AssociatePrincipalWithPortfolio(ctx workflow.Context, input *servicecatalog.AssociatePrincipalWithPortfolioInput) (*servicecatalog.AssociatePrincipalWithPortfolioOutput, error) {
    var output servicecatalog.AssociatePrincipalWithPortfolioOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociatePrincipalWithPortfolio, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) AssociatePrincipalWithPortfolioAsync(ctx workflow.Context, input *servicecatalog.AssociatePrincipalWithPortfolioInput) *ServicecatalogAssociatePrincipalWithPortfolioResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociatePrincipalWithPortfolio, input)
    return &ServicecatalogAssociatePrincipalWithPortfolioResult{Result: future}
}

func (a *ServiceCatalogStub) AssociateProductWithPortfolio(ctx workflow.Context, input *servicecatalog.AssociateProductWithPortfolioInput) (*servicecatalog.AssociateProductWithPortfolioOutput, error) {
    var output servicecatalog.AssociateProductWithPortfolioOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateProductWithPortfolio, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) AssociateProductWithPortfolioAsync(ctx workflow.Context, input *servicecatalog.AssociateProductWithPortfolioInput) *ServicecatalogAssociateProductWithPortfolioResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateProductWithPortfolio, input)
    return &ServicecatalogAssociateProductWithPortfolioResult{Result: future}
}

func (a *ServiceCatalogStub) AssociateServiceActionWithProvisioningArtifact(ctx workflow.Context, input *servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.AssociateServiceActionWithProvisioningArtifactOutput, error) {
    var output servicecatalog.AssociateServiceActionWithProvisioningArtifactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateServiceActionWithProvisioningArtifact, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) AssociateServiceActionWithProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.AssociateServiceActionWithProvisioningArtifactInput) *ServicecatalogAssociateServiceActionWithProvisioningArtifactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateServiceActionWithProvisioningArtifact, input)
    return &ServicecatalogAssociateServiceActionWithProvisioningArtifactResult{Result: future}
}

func (a *ServiceCatalogStub) AssociateTagOptionWithResource(ctx workflow.Context, input *servicecatalog.AssociateTagOptionWithResourceInput) (*servicecatalog.AssociateTagOptionWithResourceOutput, error) {
    var output servicecatalog.AssociateTagOptionWithResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateTagOptionWithResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) AssociateTagOptionWithResourceAsync(ctx workflow.Context, input *servicecatalog.AssociateTagOptionWithResourceInput) *ServicecatalogAssociateTagOptionWithResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateTagOptionWithResource, input)
    return &ServicecatalogAssociateTagOptionWithResourceResult{Result: future}
}

func (a *ServiceCatalogStub) BatchAssociateServiceActionWithProvisioningArtifact(ctx workflow.Context, input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) (*servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput, error) {
    var output servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchAssociateServiceActionWithProvisioningArtifact, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) BatchAssociateServiceActionWithProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput) *ServicecatalogBatchAssociateServiceActionWithProvisioningArtifactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchAssociateServiceActionWithProvisioningArtifact, input)
    return &ServicecatalogBatchAssociateServiceActionWithProvisioningArtifactResult{Result: future}
}

func (a *ServiceCatalogStub) BatchDisassociateServiceActionFromProvisioningArtifact(ctx workflow.Context, input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput, error) {
    var output servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDisassociateServiceActionFromProvisioningArtifact, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) BatchDisassociateServiceActionFromProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput) *ServicecatalogBatchDisassociateServiceActionFromProvisioningArtifactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.BatchDisassociateServiceActionFromProvisioningArtifact, input)
    return &ServicecatalogBatchDisassociateServiceActionFromProvisioningArtifactResult{Result: future}
}

func (a *ServiceCatalogStub) CopyProduct(ctx workflow.Context, input *servicecatalog.CopyProductInput) (*servicecatalog.CopyProductOutput, error) {
    var output servicecatalog.CopyProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) CopyProductAsync(ctx workflow.Context, input *servicecatalog.CopyProductInput) *ServicecatalogCopyProductResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CopyProduct, input)
    return &ServicecatalogCopyProductResult{Result: future}
}

func (a *ServiceCatalogStub) CreateConstraint(ctx workflow.Context, input *servicecatalog.CreateConstraintInput) (*servicecatalog.CreateConstraintOutput, error) {
    var output servicecatalog.CreateConstraintOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConstraint, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) CreateConstraintAsync(ctx workflow.Context, input *servicecatalog.CreateConstraintInput) *ServicecatalogCreateConstraintResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateConstraint, input)
    return &ServicecatalogCreateConstraintResult{Result: future}
}

func (a *ServiceCatalogStub) CreatePortfolio(ctx workflow.Context, input *servicecatalog.CreatePortfolioInput) (*servicecatalog.CreatePortfolioOutput, error) {
    var output servicecatalog.CreatePortfolioOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePortfolio, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) CreatePortfolioAsync(ctx workflow.Context, input *servicecatalog.CreatePortfolioInput) *ServicecatalogCreatePortfolioResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePortfolio, input)
    return &ServicecatalogCreatePortfolioResult{Result: future}
}

func (a *ServiceCatalogStub) CreatePortfolioShare(ctx workflow.Context, input *servicecatalog.CreatePortfolioShareInput) (*servicecatalog.CreatePortfolioShareOutput, error) {
    var output servicecatalog.CreatePortfolioShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePortfolioShare, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) CreatePortfolioShareAsync(ctx workflow.Context, input *servicecatalog.CreatePortfolioShareInput) *ServicecatalogCreatePortfolioShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePortfolioShare, input)
    return &ServicecatalogCreatePortfolioShareResult{Result: future}
}

func (a *ServiceCatalogStub) CreateProduct(ctx workflow.Context, input *servicecatalog.CreateProductInput) (*servicecatalog.CreateProductOutput, error) {
    var output servicecatalog.CreateProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) CreateProductAsync(ctx workflow.Context, input *servicecatalog.CreateProductInput) *ServicecatalogCreateProductResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateProduct, input)
    return &ServicecatalogCreateProductResult{Result: future}
}

func (a *ServiceCatalogStub) CreateProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.CreateProvisionedProductPlanInput) (*servicecatalog.CreateProvisionedProductPlanOutput, error) {
    var output servicecatalog.CreateProvisionedProductPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProvisionedProductPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) CreateProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.CreateProvisionedProductPlanInput) *ServicecatalogCreateProvisionedProductPlanResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateProvisionedProductPlan, input)
    return &ServicecatalogCreateProvisionedProductPlanResult{Result: future}
}

func (a *ServiceCatalogStub) CreateProvisioningArtifact(ctx workflow.Context, input *servicecatalog.CreateProvisioningArtifactInput) (*servicecatalog.CreateProvisioningArtifactOutput, error) {
    var output servicecatalog.CreateProvisioningArtifactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProvisioningArtifact, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) CreateProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.CreateProvisioningArtifactInput) *ServicecatalogCreateProvisioningArtifactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateProvisioningArtifact, input)
    return &ServicecatalogCreateProvisioningArtifactResult{Result: future}
}

func (a *ServiceCatalogStub) CreateServiceAction(ctx workflow.Context, input *servicecatalog.CreateServiceActionInput) (*servicecatalog.CreateServiceActionOutput, error) {
    var output servicecatalog.CreateServiceActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateServiceAction, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) CreateServiceActionAsync(ctx workflow.Context, input *servicecatalog.CreateServiceActionInput) *ServicecatalogCreateServiceActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateServiceAction, input)
    return &ServicecatalogCreateServiceActionResult{Result: future}
}

func (a *ServiceCatalogStub) CreateTagOption(ctx workflow.Context, input *servicecatalog.CreateTagOptionInput) (*servicecatalog.CreateTagOptionOutput, error) {
    var output servicecatalog.CreateTagOptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTagOption, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) CreateTagOptionAsync(ctx workflow.Context, input *servicecatalog.CreateTagOptionInput) *ServicecatalogCreateTagOptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTagOption, input)
    return &ServicecatalogCreateTagOptionResult{Result: future}
}

func (a *ServiceCatalogStub) DeleteConstraint(ctx workflow.Context, input *servicecatalog.DeleteConstraintInput) (*servicecatalog.DeleteConstraintOutput, error) {
    var output servicecatalog.DeleteConstraintOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConstraint, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DeleteConstraintAsync(ctx workflow.Context, input *servicecatalog.DeleteConstraintInput) *ServicecatalogDeleteConstraintResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConstraint, input)
    return &ServicecatalogDeleteConstraintResult{Result: future}
}

func (a *ServiceCatalogStub) DeletePortfolio(ctx workflow.Context, input *servicecatalog.DeletePortfolioInput) (*servicecatalog.DeletePortfolioOutput, error) {
    var output servicecatalog.DeletePortfolioOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePortfolio, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DeletePortfolioAsync(ctx workflow.Context, input *servicecatalog.DeletePortfolioInput) *ServicecatalogDeletePortfolioResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePortfolio, input)
    return &ServicecatalogDeletePortfolioResult{Result: future}
}

func (a *ServiceCatalogStub) DeletePortfolioShare(ctx workflow.Context, input *servicecatalog.DeletePortfolioShareInput) (*servicecatalog.DeletePortfolioShareOutput, error) {
    var output servicecatalog.DeletePortfolioShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePortfolioShare, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DeletePortfolioShareAsync(ctx workflow.Context, input *servicecatalog.DeletePortfolioShareInput) *ServicecatalogDeletePortfolioShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePortfolioShare, input)
    return &ServicecatalogDeletePortfolioShareResult{Result: future}
}

func (a *ServiceCatalogStub) DeleteProduct(ctx workflow.Context, input *servicecatalog.DeleteProductInput) (*servicecatalog.DeleteProductOutput, error) {
    var output servicecatalog.DeleteProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DeleteProductAsync(ctx workflow.Context, input *servicecatalog.DeleteProductInput) *ServicecatalogDeleteProductResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProduct, input)
    return &ServicecatalogDeleteProductResult{Result: future}
}

func (a *ServiceCatalogStub) DeleteProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.DeleteProvisionedProductPlanInput) (*servicecatalog.DeleteProvisionedProductPlanOutput, error) {
    var output servicecatalog.DeleteProvisionedProductPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProvisionedProductPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DeleteProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.DeleteProvisionedProductPlanInput) *ServicecatalogDeleteProvisionedProductPlanResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProvisionedProductPlan, input)
    return &ServicecatalogDeleteProvisionedProductPlanResult{Result: future}
}

func (a *ServiceCatalogStub) DeleteProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DeleteProvisioningArtifactInput) (*servicecatalog.DeleteProvisioningArtifactOutput, error) {
    var output servicecatalog.DeleteProvisioningArtifactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProvisioningArtifact, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DeleteProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DeleteProvisioningArtifactInput) *ServicecatalogDeleteProvisioningArtifactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteProvisioningArtifact, input)
    return &ServicecatalogDeleteProvisioningArtifactResult{Result: future}
}

func (a *ServiceCatalogStub) DeleteServiceAction(ctx workflow.Context, input *servicecatalog.DeleteServiceActionInput) (*servicecatalog.DeleteServiceActionOutput, error) {
    var output servicecatalog.DeleteServiceActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteServiceAction, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DeleteServiceActionAsync(ctx workflow.Context, input *servicecatalog.DeleteServiceActionInput) *ServicecatalogDeleteServiceActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteServiceAction, input)
    return &ServicecatalogDeleteServiceActionResult{Result: future}
}

func (a *ServiceCatalogStub) DeleteTagOption(ctx workflow.Context, input *servicecatalog.DeleteTagOptionInput) (*servicecatalog.DeleteTagOptionOutput, error) {
    var output servicecatalog.DeleteTagOptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTagOption, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DeleteTagOptionAsync(ctx workflow.Context, input *servicecatalog.DeleteTagOptionInput) *ServicecatalogDeleteTagOptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteTagOption, input)
    return &ServicecatalogDeleteTagOptionResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeConstraint(ctx workflow.Context, input *servicecatalog.DescribeConstraintInput) (*servicecatalog.DescribeConstraintOutput, error) {
    var output servicecatalog.DescribeConstraintOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConstraint, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeConstraintAsync(ctx workflow.Context, input *servicecatalog.DescribeConstraintInput) *ServicecatalogDescribeConstraintResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConstraint, input)
    return &ServicecatalogDescribeConstraintResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeCopyProductStatus(ctx workflow.Context, input *servicecatalog.DescribeCopyProductStatusInput) (*servicecatalog.DescribeCopyProductStatusOutput, error) {
    var output servicecatalog.DescribeCopyProductStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCopyProductStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeCopyProductStatusAsync(ctx workflow.Context, input *servicecatalog.DescribeCopyProductStatusInput) *ServicecatalogDescribeCopyProductStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCopyProductStatus, input)
    return &ServicecatalogDescribeCopyProductStatusResult{Result: future}
}

func (a *ServiceCatalogStub) DescribePortfolio(ctx workflow.Context, input *servicecatalog.DescribePortfolioInput) (*servicecatalog.DescribePortfolioOutput, error) {
    var output servicecatalog.DescribePortfolioOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePortfolio, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribePortfolioAsync(ctx workflow.Context, input *servicecatalog.DescribePortfolioInput) *ServicecatalogDescribePortfolioResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePortfolio, input)
    return &ServicecatalogDescribePortfolioResult{Result: future}
}

func (a *ServiceCatalogStub) DescribePortfolioShareStatus(ctx workflow.Context, input *servicecatalog.DescribePortfolioShareStatusInput) (*servicecatalog.DescribePortfolioShareStatusOutput, error) {
    var output servicecatalog.DescribePortfolioShareStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePortfolioShareStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribePortfolioShareStatusAsync(ctx workflow.Context, input *servicecatalog.DescribePortfolioShareStatusInput) *ServicecatalogDescribePortfolioShareStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePortfolioShareStatus, input)
    return &ServicecatalogDescribePortfolioShareStatusResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeProduct(ctx workflow.Context, input *servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error) {
    var output servicecatalog.DescribeProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeProductAsync(ctx workflow.Context, input *servicecatalog.DescribeProductInput) *ServicecatalogDescribeProductResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProduct, input)
    return &ServicecatalogDescribeProductResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeProductAsAdmin(ctx workflow.Context, input *servicecatalog.DescribeProductAsAdminInput) (*servicecatalog.DescribeProductAsAdminOutput, error) {
    var output servicecatalog.DescribeProductAsAdminOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProductAsAdmin, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeProductAsAdminAsync(ctx workflow.Context, input *servicecatalog.DescribeProductAsAdminInput) *ServicecatalogDescribeProductAsAdminResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProductAsAdmin, input)
    return &ServicecatalogDescribeProductAsAdminResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeProductView(ctx workflow.Context, input *servicecatalog.DescribeProductViewInput) (*servicecatalog.DescribeProductViewOutput, error) {
    var output servicecatalog.DescribeProductViewOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProductView, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeProductViewAsync(ctx workflow.Context, input *servicecatalog.DescribeProductViewInput) *ServicecatalogDescribeProductViewResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProductView, input)
    return &ServicecatalogDescribeProductViewResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeProvisionedProduct(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductInput) (*servicecatalog.DescribeProvisionedProductOutput, error) {
    var output servicecatalog.DescribeProvisionedProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisionedProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductInput) *ServicecatalogDescribeProvisionedProductResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisionedProduct, input)
    return &ServicecatalogDescribeProvisionedProductResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductPlanInput) (*servicecatalog.DescribeProvisionedProductPlanOutput, error) {
    var output servicecatalog.DescribeProvisionedProductPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisionedProductPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisionedProductPlanInput) *ServicecatalogDescribeProvisionedProductPlanResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisionedProductPlan, input)
    return &ServicecatalogDescribeProvisionedProductPlanResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DescribeProvisioningArtifactInput) (*servicecatalog.DescribeProvisioningArtifactOutput, error) {
    var output servicecatalog.DescribeProvisioningArtifactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisioningArtifact, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisioningArtifactInput) *ServicecatalogDescribeProvisioningArtifactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisioningArtifact, input)
    return &ServicecatalogDescribeProvisioningArtifactResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeProvisioningParameters(ctx workflow.Context, input *servicecatalog.DescribeProvisioningParametersInput) (*servicecatalog.DescribeProvisioningParametersOutput, error) {
    var output servicecatalog.DescribeProvisioningParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisioningParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeProvisioningParametersAsync(ctx workflow.Context, input *servicecatalog.DescribeProvisioningParametersInput) *ServicecatalogDescribeProvisioningParametersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProvisioningParameters, input)
    return &ServicecatalogDescribeProvisioningParametersResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeRecord(ctx workflow.Context, input *servicecatalog.DescribeRecordInput) (*servicecatalog.DescribeRecordOutput, error) {
    var output servicecatalog.DescribeRecordOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRecord, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeRecordAsync(ctx workflow.Context, input *servicecatalog.DescribeRecordInput) *ServicecatalogDescribeRecordResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRecord, input)
    return &ServicecatalogDescribeRecordResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeServiceAction(ctx workflow.Context, input *servicecatalog.DescribeServiceActionInput) (*servicecatalog.DescribeServiceActionOutput, error) {
    var output servicecatalog.DescribeServiceActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeServiceAction, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeServiceActionAsync(ctx workflow.Context, input *servicecatalog.DescribeServiceActionInput) *ServicecatalogDescribeServiceActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeServiceAction, input)
    return &ServicecatalogDescribeServiceActionResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeServiceActionExecutionParameters(ctx workflow.Context, input *servicecatalog.DescribeServiceActionExecutionParametersInput) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error) {
    var output servicecatalog.DescribeServiceActionExecutionParametersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeServiceActionExecutionParameters, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeServiceActionExecutionParametersAsync(ctx workflow.Context, input *servicecatalog.DescribeServiceActionExecutionParametersInput) *ServicecatalogDescribeServiceActionExecutionParametersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeServiceActionExecutionParameters, input)
    return &ServicecatalogDescribeServiceActionExecutionParametersResult{Result: future}
}

func (a *ServiceCatalogStub) DescribeTagOption(ctx workflow.Context, input *servicecatalog.DescribeTagOptionInput) (*servicecatalog.DescribeTagOptionOutput, error) {
    var output servicecatalog.DescribeTagOptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTagOption, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DescribeTagOptionAsync(ctx workflow.Context, input *servicecatalog.DescribeTagOptionInput) *ServicecatalogDescribeTagOptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTagOption, input)
    return &ServicecatalogDescribeTagOptionResult{Result: future}
}

func (a *ServiceCatalogStub) DisableAWSOrganizationsAccess(ctx workflow.Context, input *servicecatalog.DisableAWSOrganizationsAccessInput) (*servicecatalog.DisableAWSOrganizationsAccessOutput, error) {
    var output servicecatalog.DisableAWSOrganizationsAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableAWSOrganizationsAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DisableAWSOrganizationsAccessAsync(ctx workflow.Context, input *servicecatalog.DisableAWSOrganizationsAccessInput) *ServicecatalogDisableAWSOrganizationsAccessResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisableAWSOrganizationsAccess, input)
    return &ServicecatalogDisableAWSOrganizationsAccessResult{Result: future}
}

func (a *ServiceCatalogStub) DisassociateBudgetFromResource(ctx workflow.Context, input *servicecatalog.DisassociateBudgetFromResourceInput) (*servicecatalog.DisassociateBudgetFromResourceOutput, error) {
    var output servicecatalog.DisassociateBudgetFromResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateBudgetFromResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DisassociateBudgetFromResourceAsync(ctx workflow.Context, input *servicecatalog.DisassociateBudgetFromResourceInput) *ServicecatalogDisassociateBudgetFromResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateBudgetFromResource, input)
    return &ServicecatalogDisassociateBudgetFromResourceResult{Result: future}
}

func (a *ServiceCatalogStub) DisassociatePrincipalFromPortfolio(ctx workflow.Context, input *servicecatalog.DisassociatePrincipalFromPortfolioInput) (*servicecatalog.DisassociatePrincipalFromPortfolioOutput, error) {
    var output servicecatalog.DisassociatePrincipalFromPortfolioOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociatePrincipalFromPortfolio, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DisassociatePrincipalFromPortfolioAsync(ctx workflow.Context, input *servicecatalog.DisassociatePrincipalFromPortfolioInput) *ServicecatalogDisassociatePrincipalFromPortfolioResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociatePrincipalFromPortfolio, input)
    return &ServicecatalogDisassociatePrincipalFromPortfolioResult{Result: future}
}

func (a *ServiceCatalogStub) DisassociateProductFromPortfolio(ctx workflow.Context, input *servicecatalog.DisassociateProductFromPortfolioInput) (*servicecatalog.DisassociateProductFromPortfolioOutput, error) {
    var output servicecatalog.DisassociateProductFromPortfolioOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateProductFromPortfolio, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DisassociateProductFromPortfolioAsync(ctx workflow.Context, input *servicecatalog.DisassociateProductFromPortfolioInput) *ServicecatalogDisassociateProductFromPortfolioResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateProductFromPortfolio, input)
    return &ServicecatalogDisassociateProductFromPortfolioResult{Result: future}
}

func (a *ServiceCatalogStub) DisassociateServiceActionFromProvisioningArtifact(ctx workflow.Context, input *servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) (*servicecatalog.DisassociateServiceActionFromProvisioningArtifactOutput, error) {
    var output servicecatalog.DisassociateServiceActionFromProvisioningArtifactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateServiceActionFromProvisioningArtifact, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DisassociateServiceActionFromProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput) *ServicecatalogDisassociateServiceActionFromProvisioningArtifactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateServiceActionFromProvisioningArtifact, input)
    return &ServicecatalogDisassociateServiceActionFromProvisioningArtifactResult{Result: future}
}

func (a *ServiceCatalogStub) DisassociateTagOptionFromResource(ctx workflow.Context, input *servicecatalog.DisassociateTagOptionFromResourceInput) (*servicecatalog.DisassociateTagOptionFromResourceOutput, error) {
    var output servicecatalog.DisassociateTagOptionFromResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateTagOptionFromResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) DisassociateTagOptionFromResourceAsync(ctx workflow.Context, input *servicecatalog.DisassociateTagOptionFromResourceInput) *ServicecatalogDisassociateTagOptionFromResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateTagOptionFromResource, input)
    return &ServicecatalogDisassociateTagOptionFromResourceResult{Result: future}
}

func (a *ServiceCatalogStub) EnableAWSOrganizationsAccess(ctx workflow.Context, input *servicecatalog.EnableAWSOrganizationsAccessInput) (*servicecatalog.EnableAWSOrganizationsAccessOutput, error) {
    var output servicecatalog.EnableAWSOrganizationsAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableAWSOrganizationsAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) EnableAWSOrganizationsAccessAsync(ctx workflow.Context, input *servicecatalog.EnableAWSOrganizationsAccessInput) *ServicecatalogEnableAWSOrganizationsAccessResult {
    future := workflow.ExecuteActivity(ctx, a.activities.EnableAWSOrganizationsAccess, input)
    return &ServicecatalogEnableAWSOrganizationsAccessResult{Result: future}
}

func (a *ServiceCatalogStub) ExecuteProvisionedProductPlan(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductPlanInput) (*servicecatalog.ExecuteProvisionedProductPlanOutput, error) {
    var output servicecatalog.ExecuteProvisionedProductPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExecuteProvisionedProductPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ExecuteProvisionedProductPlanAsync(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductPlanInput) *ServicecatalogExecuteProvisionedProductPlanResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ExecuteProvisionedProductPlan, input)
    return &ServicecatalogExecuteProvisionedProductPlanResult{Result: future}
}

func (a *ServiceCatalogStub) ExecuteProvisionedProductServiceAction(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductServiceActionInput) (*servicecatalog.ExecuteProvisionedProductServiceActionOutput, error) {
    var output servicecatalog.ExecuteProvisionedProductServiceActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExecuteProvisionedProductServiceAction, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ExecuteProvisionedProductServiceActionAsync(ctx workflow.Context, input *servicecatalog.ExecuteProvisionedProductServiceActionInput) *ServicecatalogExecuteProvisionedProductServiceActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ExecuteProvisionedProductServiceAction, input)
    return &ServicecatalogExecuteProvisionedProductServiceActionResult{Result: future}
}

func (a *ServiceCatalogStub) GetAWSOrganizationsAccessStatus(ctx workflow.Context, input *servicecatalog.GetAWSOrganizationsAccessStatusInput) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error) {
    var output servicecatalog.GetAWSOrganizationsAccessStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAWSOrganizationsAccessStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) GetAWSOrganizationsAccessStatusAsync(ctx workflow.Context, input *servicecatalog.GetAWSOrganizationsAccessStatusInput) *ServicecatalogGetAWSOrganizationsAccessStatusResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetAWSOrganizationsAccessStatus, input)
    return &ServicecatalogGetAWSOrganizationsAccessStatusResult{Result: future}
}

func (a *ServiceCatalogStub) ListAcceptedPortfolioShares(ctx workflow.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error) {
    var output servicecatalog.ListAcceptedPortfolioSharesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAcceptedPortfolioShares, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListAcceptedPortfolioSharesAsync(ctx workflow.Context, input *servicecatalog.ListAcceptedPortfolioSharesInput) *ServicecatalogListAcceptedPortfolioSharesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAcceptedPortfolioShares, input)
    return &ServicecatalogListAcceptedPortfolioSharesResult{Result: future}
}

func (a *ServiceCatalogStub) ListBudgetsForResource(ctx workflow.Context, input *servicecatalog.ListBudgetsForResourceInput) (*servicecatalog.ListBudgetsForResourceOutput, error) {
    var output servicecatalog.ListBudgetsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBudgetsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListBudgetsForResourceAsync(ctx workflow.Context, input *servicecatalog.ListBudgetsForResourceInput) *ServicecatalogListBudgetsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBudgetsForResource, input)
    return &ServicecatalogListBudgetsForResourceResult{Result: future}
}

func (a *ServiceCatalogStub) ListConstraintsForPortfolio(ctx workflow.Context, input *servicecatalog.ListConstraintsForPortfolioInput) (*servicecatalog.ListConstraintsForPortfolioOutput, error) {
    var output servicecatalog.ListConstraintsForPortfolioOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConstraintsForPortfolio, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListConstraintsForPortfolioAsync(ctx workflow.Context, input *servicecatalog.ListConstraintsForPortfolioInput) *ServicecatalogListConstraintsForPortfolioResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListConstraintsForPortfolio, input)
    return &ServicecatalogListConstraintsForPortfolioResult{Result: future}
}

func (a *ServiceCatalogStub) ListLaunchPaths(ctx workflow.Context, input *servicecatalog.ListLaunchPathsInput) (*servicecatalog.ListLaunchPathsOutput, error) {
    var output servicecatalog.ListLaunchPathsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLaunchPaths, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListLaunchPathsAsync(ctx workflow.Context, input *servicecatalog.ListLaunchPathsInput) *ServicecatalogListLaunchPathsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListLaunchPaths, input)
    return &ServicecatalogListLaunchPathsResult{Result: future}
}

func (a *ServiceCatalogStub) ListOrganizationPortfolioAccess(ctx workflow.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error) {
    var output servicecatalog.ListOrganizationPortfolioAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListOrganizationPortfolioAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListOrganizationPortfolioAccessAsync(ctx workflow.Context, input *servicecatalog.ListOrganizationPortfolioAccessInput) *ServicecatalogListOrganizationPortfolioAccessResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListOrganizationPortfolioAccess, input)
    return &ServicecatalogListOrganizationPortfolioAccessResult{Result: future}
}

func (a *ServiceCatalogStub) ListPortfolioAccess(ctx workflow.Context, input *servicecatalog.ListPortfolioAccessInput) (*servicecatalog.ListPortfolioAccessOutput, error) {
    var output servicecatalog.ListPortfolioAccessOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPortfolioAccess, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListPortfolioAccessAsync(ctx workflow.Context, input *servicecatalog.ListPortfolioAccessInput) *ServicecatalogListPortfolioAccessResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPortfolioAccess, input)
    return &ServicecatalogListPortfolioAccessResult{Result: future}
}

func (a *ServiceCatalogStub) ListPortfolios(ctx workflow.Context, input *servicecatalog.ListPortfoliosInput) (*servicecatalog.ListPortfoliosOutput, error) {
    var output servicecatalog.ListPortfoliosOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPortfolios, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListPortfoliosAsync(ctx workflow.Context, input *servicecatalog.ListPortfoliosInput) *ServicecatalogListPortfoliosResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPortfolios, input)
    return &ServicecatalogListPortfoliosResult{Result: future}
}

func (a *ServiceCatalogStub) ListPortfoliosForProduct(ctx workflow.Context, input *servicecatalog.ListPortfoliosForProductInput) (*servicecatalog.ListPortfoliosForProductOutput, error) {
    var output servicecatalog.ListPortfoliosForProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPortfoliosForProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListPortfoliosForProductAsync(ctx workflow.Context, input *servicecatalog.ListPortfoliosForProductInput) *ServicecatalogListPortfoliosForProductResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPortfoliosForProduct, input)
    return &ServicecatalogListPortfoliosForProductResult{Result: future}
}

func (a *ServiceCatalogStub) ListPrincipalsForPortfolio(ctx workflow.Context, input *servicecatalog.ListPrincipalsForPortfolioInput) (*servicecatalog.ListPrincipalsForPortfolioOutput, error) {
    var output servicecatalog.ListPrincipalsForPortfolioOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPrincipalsForPortfolio, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListPrincipalsForPortfolioAsync(ctx workflow.Context, input *servicecatalog.ListPrincipalsForPortfolioInput) *ServicecatalogListPrincipalsForPortfolioResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPrincipalsForPortfolio, input)
    return &ServicecatalogListPrincipalsForPortfolioResult{Result: future}
}

func (a *ServiceCatalogStub) ListProvisionedProductPlans(ctx workflow.Context, input *servicecatalog.ListProvisionedProductPlansInput) (*servicecatalog.ListProvisionedProductPlansOutput, error) {
    var output servicecatalog.ListProvisionedProductPlansOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProvisionedProductPlans, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListProvisionedProductPlansAsync(ctx workflow.Context, input *servicecatalog.ListProvisionedProductPlansInput) *ServicecatalogListProvisionedProductPlansResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProvisionedProductPlans, input)
    return &ServicecatalogListProvisionedProductPlansResult{Result: future}
}

func (a *ServiceCatalogStub) ListProvisioningArtifacts(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsInput) (*servicecatalog.ListProvisioningArtifactsOutput, error) {
    var output servicecatalog.ListProvisioningArtifactsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProvisioningArtifacts, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListProvisioningArtifactsAsync(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsInput) *ServicecatalogListProvisioningArtifactsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProvisioningArtifacts, input)
    return &ServicecatalogListProvisioningArtifactsResult{Result: future}
}

func (a *ServiceCatalogStub) ListProvisioningArtifactsForServiceAction(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error) {
    var output servicecatalog.ListProvisioningArtifactsForServiceActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProvisioningArtifactsForServiceAction, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListProvisioningArtifactsForServiceActionAsync(ctx workflow.Context, input *servicecatalog.ListProvisioningArtifactsForServiceActionInput) *ServicecatalogListProvisioningArtifactsForServiceActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProvisioningArtifactsForServiceAction, input)
    return &ServicecatalogListProvisioningArtifactsForServiceActionResult{Result: future}
}

func (a *ServiceCatalogStub) ListRecordHistory(ctx workflow.Context, input *servicecatalog.ListRecordHistoryInput) (*servicecatalog.ListRecordHistoryOutput, error) {
    var output servicecatalog.ListRecordHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRecordHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListRecordHistoryAsync(ctx workflow.Context, input *servicecatalog.ListRecordHistoryInput) *ServicecatalogListRecordHistoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRecordHistory, input)
    return &ServicecatalogListRecordHistoryResult{Result: future}
}

func (a *ServiceCatalogStub) ListResourcesForTagOption(ctx workflow.Context, input *servicecatalog.ListResourcesForTagOptionInput) (*servicecatalog.ListResourcesForTagOptionOutput, error) {
    var output servicecatalog.ListResourcesForTagOptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourcesForTagOption, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListResourcesForTagOptionAsync(ctx workflow.Context, input *servicecatalog.ListResourcesForTagOptionInput) *ServicecatalogListResourcesForTagOptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResourcesForTagOption, input)
    return &ServicecatalogListResourcesForTagOptionResult{Result: future}
}

func (a *ServiceCatalogStub) ListServiceActions(ctx workflow.Context, input *servicecatalog.ListServiceActionsInput) (*servicecatalog.ListServiceActionsOutput, error) {
    var output servicecatalog.ListServiceActionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListServiceActions, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListServiceActionsAsync(ctx workflow.Context, input *servicecatalog.ListServiceActionsInput) *ServicecatalogListServiceActionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListServiceActions, input)
    return &ServicecatalogListServiceActionsResult{Result: future}
}

func (a *ServiceCatalogStub) ListServiceActionsForProvisioningArtifact(ctx workflow.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error) {
    var output servicecatalog.ListServiceActionsForProvisioningArtifactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListServiceActionsForProvisioningArtifact, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListServiceActionsForProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.ListServiceActionsForProvisioningArtifactInput) *ServicecatalogListServiceActionsForProvisioningArtifactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListServiceActionsForProvisioningArtifact, input)
    return &ServicecatalogListServiceActionsForProvisioningArtifactResult{Result: future}
}

func (a *ServiceCatalogStub) ListStackInstancesForProvisionedProduct(ctx workflow.Context, input *servicecatalog.ListStackInstancesForProvisionedProductInput) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error) {
    var output servicecatalog.ListStackInstancesForProvisionedProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListStackInstancesForProvisionedProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListStackInstancesForProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.ListStackInstancesForProvisionedProductInput) *ServicecatalogListStackInstancesForProvisionedProductResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListStackInstancesForProvisionedProduct, input)
    return &ServicecatalogListStackInstancesForProvisionedProductResult{Result: future}
}

func (a *ServiceCatalogStub) ListTagOptions(ctx workflow.Context, input *servicecatalog.ListTagOptionsInput) (*servicecatalog.ListTagOptionsOutput, error) {
    var output servicecatalog.ListTagOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ListTagOptionsAsync(ctx workflow.Context, input *servicecatalog.ListTagOptionsInput) *ServicecatalogListTagOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagOptions, input)
    return &ServicecatalogListTagOptionsResult{Result: future}
}

func (a *ServiceCatalogStub) ProvisionProduct(ctx workflow.Context, input *servicecatalog.ProvisionProductInput) (*servicecatalog.ProvisionProductOutput, error) {
    var output servicecatalog.ProvisionProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ProvisionProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ProvisionProductAsync(ctx workflow.Context, input *servicecatalog.ProvisionProductInput) *ServicecatalogProvisionProductResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ProvisionProduct, input)
    return &ServicecatalogProvisionProductResult{Result: future}
}

func (a *ServiceCatalogStub) RejectPortfolioShare(ctx workflow.Context, input *servicecatalog.RejectPortfolioShareInput) (*servicecatalog.RejectPortfolioShareOutput, error) {
    var output servicecatalog.RejectPortfolioShareOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RejectPortfolioShare, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) RejectPortfolioShareAsync(ctx workflow.Context, input *servicecatalog.RejectPortfolioShareInput) *ServicecatalogRejectPortfolioShareResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RejectPortfolioShare, input)
    return &ServicecatalogRejectPortfolioShareResult{Result: future}
}

func (a *ServiceCatalogStub) ScanProvisionedProducts(ctx workflow.Context, input *servicecatalog.ScanProvisionedProductsInput) (*servicecatalog.ScanProvisionedProductsOutput, error) {
    var output servicecatalog.ScanProvisionedProductsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ScanProvisionedProducts, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) ScanProvisionedProductsAsync(ctx workflow.Context, input *servicecatalog.ScanProvisionedProductsInput) *ServicecatalogScanProvisionedProductsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ScanProvisionedProducts, input)
    return &ServicecatalogScanProvisionedProductsResult{Result: future}
}

func (a *ServiceCatalogStub) SearchProducts(ctx workflow.Context, input *servicecatalog.SearchProductsInput) (*servicecatalog.SearchProductsOutput, error) {
    var output servicecatalog.SearchProductsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchProducts, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) SearchProductsAsync(ctx workflow.Context, input *servicecatalog.SearchProductsInput) *ServicecatalogSearchProductsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SearchProducts, input)
    return &ServicecatalogSearchProductsResult{Result: future}
}

func (a *ServiceCatalogStub) SearchProductsAsAdmin(ctx workflow.Context, input *servicecatalog.SearchProductsAsAdminInput) (*servicecatalog.SearchProductsAsAdminOutput, error) {
    var output servicecatalog.SearchProductsAsAdminOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchProductsAsAdmin, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) SearchProductsAsAdminAsync(ctx workflow.Context, input *servicecatalog.SearchProductsAsAdminInput) *ServicecatalogSearchProductsAsAdminResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SearchProductsAsAdmin, input)
    return &ServicecatalogSearchProductsAsAdminResult{Result: future}
}

func (a *ServiceCatalogStub) SearchProvisionedProducts(ctx workflow.Context, input *servicecatalog.SearchProvisionedProductsInput) (*servicecatalog.SearchProvisionedProductsOutput, error) {
    var output servicecatalog.SearchProvisionedProductsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SearchProvisionedProducts, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) SearchProvisionedProductsAsync(ctx workflow.Context, input *servicecatalog.SearchProvisionedProductsInput) *ServicecatalogSearchProvisionedProductsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SearchProvisionedProducts, input)
    return &ServicecatalogSearchProvisionedProductsResult{Result: future}
}

func (a *ServiceCatalogStub) TerminateProvisionedProduct(ctx workflow.Context, input *servicecatalog.TerminateProvisionedProductInput) (*servicecatalog.TerminateProvisionedProductOutput, error) {
    var output servicecatalog.TerminateProvisionedProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TerminateProvisionedProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) TerminateProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.TerminateProvisionedProductInput) *ServicecatalogTerminateProvisionedProductResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TerminateProvisionedProduct, input)
    return &ServicecatalogTerminateProvisionedProductResult{Result: future}
}

func (a *ServiceCatalogStub) UpdateConstraint(ctx workflow.Context, input *servicecatalog.UpdateConstraintInput) (*servicecatalog.UpdateConstraintOutput, error) {
    var output servicecatalog.UpdateConstraintOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConstraint, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) UpdateConstraintAsync(ctx workflow.Context, input *servicecatalog.UpdateConstraintInput) *ServicecatalogUpdateConstraintResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateConstraint, input)
    return &ServicecatalogUpdateConstraintResult{Result: future}
}

func (a *ServiceCatalogStub) UpdatePortfolio(ctx workflow.Context, input *servicecatalog.UpdatePortfolioInput) (*servicecatalog.UpdatePortfolioOutput, error) {
    var output servicecatalog.UpdatePortfolioOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePortfolio, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) UpdatePortfolioAsync(ctx workflow.Context, input *servicecatalog.UpdatePortfolioInput) *ServicecatalogUpdatePortfolioResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdatePortfolio, input)
    return &ServicecatalogUpdatePortfolioResult{Result: future}
}

func (a *ServiceCatalogStub) UpdateProduct(ctx workflow.Context, input *servicecatalog.UpdateProductInput) (*servicecatalog.UpdateProductOutput, error) {
    var output servicecatalog.UpdateProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) UpdateProductAsync(ctx workflow.Context, input *servicecatalog.UpdateProductInput) *ServicecatalogUpdateProductResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateProduct, input)
    return &ServicecatalogUpdateProductResult{Result: future}
}

func (a *ServiceCatalogStub) UpdateProvisionedProduct(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductInput) (*servicecatalog.UpdateProvisionedProductOutput, error) {
    var output servicecatalog.UpdateProvisionedProductOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateProvisionedProduct, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) UpdateProvisionedProductAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductInput) *ServicecatalogUpdateProvisionedProductResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateProvisionedProduct, input)
    return &ServicecatalogUpdateProvisionedProductResult{Result: future}
}

func (a *ServiceCatalogStub) UpdateProvisionedProductProperties(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductPropertiesInput) (*servicecatalog.UpdateProvisionedProductPropertiesOutput, error) {
    var output servicecatalog.UpdateProvisionedProductPropertiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateProvisionedProductProperties, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) UpdateProvisionedProductPropertiesAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisionedProductPropertiesInput) *ServicecatalogUpdateProvisionedProductPropertiesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateProvisionedProductProperties, input)
    return &ServicecatalogUpdateProvisionedProductPropertiesResult{Result: future}
}

func (a *ServiceCatalogStub) UpdateProvisioningArtifact(ctx workflow.Context, input *servicecatalog.UpdateProvisioningArtifactInput) (*servicecatalog.UpdateProvisioningArtifactOutput, error) {
    var output servicecatalog.UpdateProvisioningArtifactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateProvisioningArtifact, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) UpdateProvisioningArtifactAsync(ctx workflow.Context, input *servicecatalog.UpdateProvisioningArtifactInput) *ServicecatalogUpdateProvisioningArtifactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateProvisioningArtifact, input)
    return &ServicecatalogUpdateProvisioningArtifactResult{Result: future}
}

func (a *ServiceCatalogStub) UpdateServiceAction(ctx workflow.Context, input *servicecatalog.UpdateServiceActionInput) (*servicecatalog.UpdateServiceActionOutput, error) {
    var output servicecatalog.UpdateServiceActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateServiceAction, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) UpdateServiceActionAsync(ctx workflow.Context, input *servicecatalog.UpdateServiceActionInput) *ServicecatalogUpdateServiceActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateServiceAction, input)
    return &ServicecatalogUpdateServiceActionResult{Result: future}
}

func (a *ServiceCatalogStub) UpdateTagOption(ctx workflow.Context, input *servicecatalog.UpdateTagOptionInput) (*servicecatalog.UpdateTagOptionOutput, error) {
    var output servicecatalog.UpdateTagOptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTagOption, input).Get(ctx, &output)
    return &output, err
}

func (a *ServiceCatalogStub) UpdateTagOptionAsync(ctx workflow.Context, input *servicecatalog.UpdateTagOptionInput) *ServicecatalogUpdateTagOptionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTagOption, input)
    return &ServicecatalogUpdateTagOptionResult{Result: future}
}
