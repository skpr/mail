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

// Enables or disables the custom MAIL FROM domain setup for a verified identity
// (an email address or a domain). To send emails using the specified MAIL FROM
// domain, you must add an MX record to your MAIL FROM domain's DNS settings. To
// ensure that your emails pass Sender Policy Framework (SPF) checks, you must also
// add or update an SPF record. For more information, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/dg/mail-from.html) . You can
// execute this operation no more than once per second.
func (c *Client) SetIdentityMailFromDomain(ctx context.Context, params *SetIdentityMailFromDomainInput, optFns ...func(*Options)) (*SetIdentityMailFromDomainOutput, error) {
	if params == nil {
		params = &SetIdentityMailFromDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetIdentityMailFromDomain", params, optFns, c.addOperationSetIdentityMailFromDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetIdentityMailFromDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to enable or disable the Amazon SES custom MAIL FROM
// domain setup for a verified identity. For information about using a custom MAIL
// FROM domain, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/mail-from.html)
// .
type SetIdentityMailFromDomainInput struct {

	// The verified identity.
	//
	// This member is required.
	Identity *string

	// The action for Amazon SES to take if it cannot successfully read the required
	// MX record when you send an email. If you choose UseDefaultValue , Amazon SES
	// uses amazonses.com (or a subdomain of that) as the MAIL FROM domain. If you
	// choose RejectMessage , Amazon SES returns a MailFromDomainNotVerified error and
	// not send the email. The action specified in BehaviorOnMXFailure is taken when
	// the custom MAIL FROM domain setup is in the Pending , Failed , and
	// TemporaryFailure states.
	BehaviorOnMXFailure types.BehaviorOnMXFailure

	// The custom MAIL FROM domain for the verified identity to use. The MAIL FROM
	// domain must 1) be a subdomain of the verified identity, 2) not be used in a
	// "From" address if the MAIL FROM domain is the destination of email feedback
	// forwarding (for more information, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/mail-from.html)
	// ), and 3) not be used to receive emails. A value of null disables the custom
	// MAIL FROM setting for the identity.
	MailFromDomain *string

	noSmithyDocumentSerde
}

// An empty element returned on a successful request.
type SetIdentityMailFromDomainOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetIdentityMailFromDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetIdentityMailFromDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetIdentityMailFromDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetIdentityMailFromDomain"); err != nil {
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
	if err = addOpSetIdentityMailFromDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetIdentityMailFromDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetIdentityMailFromDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetIdentityMailFromDomain",
	}
}
