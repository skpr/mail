// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list containing all of the identities (email addresses and domains)
// for your Amazon Web Services account in the current Amazon Web Services Region,
// regardless of verification status.
//
// You can execute this operation no more than once per second.
//
// It's recommended that for successive pagination calls of this API, you continue
// to the use the same parameter/value pairs as used in the original call, e.g., if
// you used IdentityType=Domain in the the original call and received a NextToken
// in the response, you should continue providing the IdentityType=Domain
// parameter for further NextToken calls; however, if you didn't provide the
// IdentityType parameter in the original call, then continue to not provide it for
// successive pagination calls. Using this protocol will ensure consistent results.
func (c *Client) ListIdentities(ctx context.Context, params *ListIdentitiesInput, optFns ...func(*Options)) (*ListIdentitiesOutput, error) {
	if params == nil {
		params = &ListIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIdentities", params, optFns, c.addOperationListIdentitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return a list of all identities (email addresses and
// domains) that you have attempted to verify under your Amazon Web Services
// account, regardless of verification status.
type ListIdentitiesInput struct {

	// The type of the identities to list. Possible values are "EmailAddress" and
	// "Domain". If this parameter is omitted, then all identities are listed.
	IdentityType types.IdentityType

	// The maximum number of identities per page. Possible values are 1-1000 inclusive.
	MaxItems *int32

	// The token to use for pagination.
	NextToken *string

	noSmithyDocumentSerde
}

// A list of all identities that you have attempted to verify under your Amazon
// Web Services account, regardless of verification status.
type ListIdentitiesOutput struct {

	// A list of identities.
	//
	// This member is required.
	Identities []string

	// The token used for pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIdentitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListIdentities"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIdentities(options.Region), middleware.Before); err != nil {
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

// ListIdentitiesPaginatorOptions is the paginator options for ListIdentities
type ListIdentitiesPaginatorOptions struct {
	// The maximum number of identities per page. Possible values are 1-1000 inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIdentitiesPaginator is a paginator for ListIdentities
type ListIdentitiesPaginator struct {
	options   ListIdentitiesPaginatorOptions
	client    ListIdentitiesAPIClient
	params    *ListIdentitiesInput
	nextToken *string
	firstPage bool
}

// NewListIdentitiesPaginator returns a new ListIdentitiesPaginator
func NewListIdentitiesPaginator(client ListIdentitiesAPIClient, params *ListIdentitiesInput, optFns ...func(*ListIdentitiesPaginatorOptions)) *ListIdentitiesPaginator {
	if params == nil {
		params = &ListIdentitiesInput{}
	}

	options := ListIdentitiesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIdentitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIdentitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListIdentities page.
func (p *ListIdentitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIdentitiesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListIdentities(ctx, &params, optFns...)
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

// ListIdentitiesAPIClient is a client that implements the ListIdentities
// operation.
type ListIdentitiesAPIClient interface {
	ListIdentities(context.Context, *ListIdentitiesInput, ...func(*Options)) (*ListIdentitiesOutput, error)
}

var _ ListIdentitiesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListIdentities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListIdentities",
	}
}
