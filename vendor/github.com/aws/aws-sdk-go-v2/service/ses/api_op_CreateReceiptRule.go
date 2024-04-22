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

// Creates a receipt rule. For information about setting up receipt rules, see the
// Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html)
// . You can execute this operation no more than once per second.
func (c *Client) CreateReceiptRule(ctx context.Context, params *CreateReceiptRuleInput, optFns ...func(*Options)) (*CreateReceiptRuleOutput, error) {
	if params == nil {
		params = &CreateReceiptRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReceiptRule", params, optFns, c.addOperationCreateReceiptRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReceiptRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to create a receipt rule. You use receipt rules to receive
// email with Amazon SES. For more information, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/receiving-email-concepts.html)
// .
type CreateReceiptRuleInput struct {

	// A data structure that contains the specified rule's name, actions, recipients,
	// domains, enabled status, scan status, and TLS policy.
	//
	// This member is required.
	Rule *types.ReceiptRule

	// The name of the rule set where the receipt rule is added.
	//
	// This member is required.
	RuleSetName *string

	// The name of an existing rule after which the new rule is placed. If this
	// parameter is null, the new rule is inserted at the beginning of the rule list.
	After *string

	noSmithyDocumentSerde
}

// An empty element returned on a successful request.
type CreateReceiptRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReceiptRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateReceiptRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateReceiptRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateReceiptRule"); err != nil {
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
	if err = addOpCreateReceiptRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReceiptRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReceiptRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateReceiptRule",
	}
}
