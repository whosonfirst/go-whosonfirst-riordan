// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the properties of available patches organized by product, product family,
// classification, severity, and other properties of available patches. You can use
// the reported properties in the filters you specify in requests for operations
// such as CreatePatchBaseline, UpdatePatchBaseline, DescribeAvailablePatches, and DescribePatchBaselines.
//
// The following section lists the properties that can be used in filters for each
// major operating system type:
//
// AMAZON_LINUX Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// AMAZON_LINUX_2 Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// AMAZON_LINUX_2023 Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// CENTOS Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// DEBIAN Valid properties: PRODUCT | PRIORITY
//
// MACOS Valid properties: PRODUCT | CLASSIFICATION
//
// ORACLE_LINUX Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// REDHAT_ENTERPRISE_LINUX Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// SUSE Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// UBUNTU Valid properties: PRODUCT | PRIORITY
//
// WINDOWS Valid properties: PRODUCT | PRODUCT_FAMILY | CLASSIFICATION |
// MSRC_SEVERITY
func (c *Client) DescribePatchProperties(ctx context.Context, params *DescribePatchPropertiesInput, optFns ...func(*Options)) (*DescribePatchPropertiesOutput, error) {
	if params == nil {
		params = &DescribePatchPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePatchProperties", params, optFns, c.addOperationDescribePatchPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePatchPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePatchPropertiesInput struct {

	// The operating system type for which to list patches.
	//
	// This member is required.
	OperatingSystem types.OperatingSystem

	// The patch property for which you want to view patch details.
	//
	// This member is required.
	Property types.PatchProperty

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// Indicates whether to list patches for the Windows operating system or for
	// applications released by Microsoft. Not applicable for the Linux or macOS
	// operating systems.
	PatchSet types.PatchSet

	noSmithyDocumentSerde
}

type DescribePatchPropertiesOutput struct {

	// The token for the next set of items to return. (You use this token in the next
	// call.)
	NextToken *string

	// A list of the properties for patches matching the filter request parameters.
	Properties []map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePatchPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePatchProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePatchProperties{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePatchProperties"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpDescribePatchPropertiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePatchProperties(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// DescribePatchPropertiesPaginatorOptions is the paginator options for
// DescribePatchProperties
type DescribePatchPropertiesPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePatchPropertiesPaginator is a paginator for DescribePatchProperties
type DescribePatchPropertiesPaginator struct {
	options   DescribePatchPropertiesPaginatorOptions
	client    DescribePatchPropertiesAPIClient
	params    *DescribePatchPropertiesInput
	nextToken *string
	firstPage bool
}

// NewDescribePatchPropertiesPaginator returns a new
// DescribePatchPropertiesPaginator
func NewDescribePatchPropertiesPaginator(client DescribePatchPropertiesAPIClient, params *DescribePatchPropertiesInput, optFns ...func(*DescribePatchPropertiesPaginatorOptions)) *DescribePatchPropertiesPaginator {
	if params == nil {
		params = &DescribePatchPropertiesInput{}
	}

	options := DescribePatchPropertiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePatchPropertiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePatchPropertiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribePatchProperties page.
func (p *DescribePatchPropertiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePatchPropertiesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribePatchProperties(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribePatchPropertiesAPIClient is a client that implements the
// DescribePatchProperties operation.
type DescribePatchPropertiesAPIClient interface {
	DescribePatchProperties(context.Context, *DescribePatchPropertiesInput, ...func(*Options)) (*DescribePatchPropertiesOutput, error)
}

var _ DescribePatchPropertiesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribePatchProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePatchProperties",
	}
}
