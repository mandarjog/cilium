// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteTransitGatewayRouteTableInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the transit gateway route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTransitGatewayRouteTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTransitGatewayRouteTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTransitGatewayRouteTableInput"}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTransitGatewayRouteTableOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deleted transit gateway route table.
	TransitGatewayRouteTable *TransitGatewayRouteTable `locationName:"transitGatewayRouteTable" type:"structure"`
}

// String returns the string representation
func (s DeleteTransitGatewayRouteTableOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTransitGatewayRouteTable = "DeleteTransitGatewayRouteTable"

// DeleteTransitGatewayRouteTableRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified transit gateway route table. You must disassociate
// the route table from any transit gateway route tables before you can delete
// it.
//
//    // Example sending a request using DeleteTransitGatewayRouteTableRequest.
//    req := client.DeleteTransitGatewayRouteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGatewayRouteTable
func (c *Client) DeleteTransitGatewayRouteTableRequest(input *DeleteTransitGatewayRouteTableInput) DeleteTransitGatewayRouteTableRequest {
	op := &aws.Operation{
		Name:       opDeleteTransitGatewayRouteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTransitGatewayRouteTableInput{}
	}

	req := c.newRequest(op, input, &DeleteTransitGatewayRouteTableOutput{})

	return DeleteTransitGatewayRouteTableRequest{Request: req, Input: input, Copy: c.DeleteTransitGatewayRouteTableRequest}
}

// DeleteTransitGatewayRouteTableRequest is the request type for the
// DeleteTransitGatewayRouteTable API operation.
type DeleteTransitGatewayRouteTableRequest struct {
	*aws.Request
	Input *DeleteTransitGatewayRouteTableInput
	Copy  func(*DeleteTransitGatewayRouteTableInput) DeleteTransitGatewayRouteTableRequest
}

// Send marshals and sends the DeleteTransitGatewayRouteTable API request.
func (r DeleteTransitGatewayRouteTableRequest) Send(ctx context.Context) (*DeleteTransitGatewayRouteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTransitGatewayRouteTableResponse{
		DeleteTransitGatewayRouteTableOutput: r.Request.Data.(*DeleteTransitGatewayRouteTableOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTransitGatewayRouteTableResponse is the response type for the
// DeleteTransitGatewayRouteTable API operation.
type DeleteTransitGatewayRouteTableResponse struct {
	*DeleteTransitGatewayRouteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTransitGatewayRouteTable request.
func (r *DeleteTransitGatewayRouteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
