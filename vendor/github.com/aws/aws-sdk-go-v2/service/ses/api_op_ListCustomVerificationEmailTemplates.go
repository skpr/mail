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

// Lists the existing custom verification email templates for your account in the
// current Amazon Web Services Region. For more information about custom
// verification email templates, see Using Custom Verification Email Templates (https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second.
func (c *Client) ListCustomVerificationEmailTemplates(ctx context.Context, params *ListCustomVerificationEmailTemplatesInput, optFns ...func(*Options)) (*ListCustomVerificationEmailTemplatesOutput, error) {
	if params == nil {
		params = &ListCustomVerificationEmailTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomVerificationEmailTemplates", params, optFns, c.addOperationListCustomVerificationEmailTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomVerificationEmailTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list the existing custom verification email templates
// for your account. For more information about custom verification email
// templates, see Using Custom Verification Email Templates (https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom)
// in the Amazon SES Developer Guide.
type ListCustomVerificationEmailTemplatesInput struct {

	// The maximum number of custom verification email templates to return. This value
	// must be at least 1 and less than or equal to 50. If you do not specify a value,
	// or if you specify a value less than 1 or greater than 50, the operation returns
	// up to 50 results.
	MaxResults *int32

	// An array the contains the name and creation time stamp for each template in
	// your Amazon SES account.
	NextToken *string

	noSmithyDocumentSerde
}

// A paginated list of custom verification email templates.
type ListCustomVerificationEmailTemplatesOutput struct {

	// A list of the custom verification email templates that exist in your account.
	CustomVerificationEmailTemplates []types.CustomVerificationEmailTemplate

	// A token indicating that there are additional custom verification email
	// templates available to be listed. Pass this token to a subsequent call to
	// ListTemplates to retrieve the next 50 custom verification email templates.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomVerificationEmailTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListCustomVerificationEmailTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListCustomVerificationEmailTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCustomVerificationEmailTemplates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomVerificationEmailTemplates(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListCustomVerificationEmailTemplatesAPIClient is a client that implements the
// ListCustomVerificationEmailTemplates operation.
type ListCustomVerificationEmailTemplatesAPIClient interface {
	ListCustomVerificationEmailTemplates(context.Context, *ListCustomVerificationEmailTemplatesInput, ...func(*Options)) (*ListCustomVerificationEmailTemplatesOutput, error)
}

var _ ListCustomVerificationEmailTemplatesAPIClient = (*Client)(nil)

// ListCustomVerificationEmailTemplatesPaginatorOptions is the paginator options
// for ListCustomVerificationEmailTemplates
type ListCustomVerificationEmailTemplatesPaginatorOptions struct {
	// The maximum number of custom verification email templates to return. This value
	// must be at least 1 and less than or equal to 50. If you do not specify a value,
	// or if you specify a value less than 1 or greater than 50, the operation returns
	// up to 50 results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomVerificationEmailTemplatesPaginator is a paginator for
// ListCustomVerificationEmailTemplates
type ListCustomVerificationEmailTemplatesPaginator struct {
	options   ListCustomVerificationEmailTemplatesPaginatorOptions
	client    ListCustomVerificationEmailTemplatesAPIClient
	params    *ListCustomVerificationEmailTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListCustomVerificationEmailTemplatesPaginator returns a new
// ListCustomVerificationEmailTemplatesPaginator
func NewListCustomVerificationEmailTemplatesPaginator(client ListCustomVerificationEmailTemplatesAPIClient, params *ListCustomVerificationEmailTemplatesInput, optFns ...func(*ListCustomVerificationEmailTemplatesPaginatorOptions)) *ListCustomVerificationEmailTemplatesPaginator {
	if params == nil {
		params = &ListCustomVerificationEmailTemplatesInput{}
	}

	options := ListCustomVerificationEmailTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomVerificationEmailTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomVerificationEmailTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomVerificationEmailTemplates page.
func (p *ListCustomVerificationEmailTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomVerificationEmailTemplatesOutput, error) {
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

	result, err := p.client.ListCustomVerificationEmailTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomVerificationEmailTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCustomVerificationEmailTemplates",
	}
}
