// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The VPC Endpoint data source provides details about
// a specific VPC endpoint.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Declare the data source
//			s3, err := ec2.LookupVpcEndpoint(ctx, &ec2.LookupVpcEndpointArgs{
//				VpcId:       pulumi.StringRef(foo.Id),
//				ServiceName: pulumi.StringRef("com.amazonaws.us-west-2.s3"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcEndpointRouteTableAssociation(ctx, "private_s3", &ec2.VpcEndpointRouteTableAssociationArgs{
//				VpcEndpointId: *pulumi.String(s3.Id),
//				RouteTableId:  pulumi.Any(private.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupVpcEndpoint(ctx *pulumi.Context, args *LookupVpcEndpointArgs, opts ...pulumi.InvokeOption) (*LookupVpcEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcEndpointResult
	err := ctx.Invoke("aws:ec2/getVpcEndpoint:getVpcEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcEndpoint.
type LookupVpcEndpointArgs struct {
	// Custom filter block as described below.
	Filters []GetVpcEndpointFilter `pulumi:"filters"`
	// ID of the specific VPC Endpoint to retrieve.
	Id *string `pulumi:"id"`
	// Service name of the specific VPC Endpoint to retrieve. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName *string `pulumi:"serviceName"`
	// State of the specific VPC Endpoint to retrieve.
	State *string `pulumi:"state"`
	// Map of tags, each pair of which must exactly match
	// a pair on the specific VPC Endpoint to retrieve.
	Tags map[string]string `pulumi:"tags"`
	// ID of the VPC in which the specific VPC Endpoint is used.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getVpcEndpoint.
type LookupVpcEndpointResult struct {
	// ARN of the VPC endpoint.
	Arn string `pulumi:"arn"`
	// List of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
	CidrBlocks []string `pulumi:"cidrBlocks"`
	// DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS entry blocks are documented below.
	DnsEntries []GetVpcEndpointDnsEntry `pulumi:"dnsEntries"`
	// DNS options for the VPC Endpoint. DNS options blocks are documented below.
	DnsOptions    []GetVpcEndpointDnsOption `pulumi:"dnsOptions"`
	Filters       []GetVpcEndpointFilter    `pulumi:"filters"`
	Id            string                    `pulumi:"id"`
	IpAddressType string                    `pulumi:"ipAddressType"`
	// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// ID of the AWS account that owns the VPC endpoint.
	OwnerId string `pulumi:"ownerId"`
	// Policy document associated with the VPC Endpoint. Applicable for endpoints of type `Gateway`.
	Policy string `pulumi:"policy"`
	// Prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
	PrefixListId string `pulumi:"prefixListId"`
	// Whether or not the VPC is associated with a private hosted zone - `true` or `false`. Applicable for endpoints of type `Interface`.
	PrivateDnsEnabled bool `pulumi:"privateDnsEnabled"`
	// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
	RequesterManaged bool `pulumi:"requesterManaged"`
	// One or more route tables associated with the VPC Endpoint. Applicable for endpoints of type `Gateway`.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// One or more security groups associated with the network interfaces. Applicable for endpoints of type `Interface`.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	ServiceName      string   `pulumi:"serviceName"`
	State            string   `pulumi:"state"`
	// One or more subnets in which the VPC Endpoint is located. Applicable for endpoints of type `Interface`.
	SubnetIds []string          `pulumi:"subnetIds"`
	Tags      map[string]string `pulumi:"tags"`
	// VPC Endpoint type, `Gateway` or `Interface`.
	VpcEndpointType string `pulumi:"vpcEndpointType"`
	VpcId           string `pulumi:"vpcId"`
}

func LookupVpcEndpointOutput(ctx *pulumi.Context, args LookupVpcEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupVpcEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcEndpointResult, error) {
			args := v.(LookupVpcEndpointArgs)
			r, err := LookupVpcEndpoint(ctx, &args, opts...)
			var s LookupVpcEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcEndpointResultOutput)
}

// A collection of arguments for invoking getVpcEndpoint.
type LookupVpcEndpointOutputArgs struct {
	// Custom filter block as described below.
	Filters GetVpcEndpointFilterArrayInput `pulumi:"filters"`
	// ID of the specific VPC Endpoint to retrieve.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Service name of the specific VPC Endpoint to retrieve. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
	// State of the specific VPC Endpoint to retrieve.
	State pulumi.StringPtrInput `pulumi:"state"`
	// Map of tags, each pair of which must exactly match
	// a pair on the specific VPC Endpoint to retrieve.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// ID of the VPC in which the specific VPC Endpoint is used.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupVpcEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointArgs)(nil)).Elem()
}

// A collection of values returned by getVpcEndpoint.
type LookupVpcEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupVpcEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointResult)(nil)).Elem()
}

func (o LookupVpcEndpointResultOutput) ToLookupVpcEndpointResultOutput() LookupVpcEndpointResultOutput {
	return o
}

func (o LookupVpcEndpointResultOutput) ToLookupVpcEndpointResultOutputWithContext(ctx context.Context) LookupVpcEndpointResultOutput {
	return o
}

// ARN of the VPC endpoint.
func (o LookupVpcEndpointResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) string { return v.Arn }).(pulumi.StringOutput)
}

// List of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
func (o LookupVpcEndpointResultOutput) CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.CidrBlocks }).(pulumi.StringArrayOutput)
}

// DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS entry blocks are documented below.
func (o LookupVpcEndpointResultOutput) DnsEntries() GetVpcEndpointDnsEntryArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []GetVpcEndpointDnsEntry { return v.DnsEntries }).(GetVpcEndpointDnsEntryArrayOutput)
}

// DNS options for the VPC Endpoint. DNS options blocks are documented below.
func (o LookupVpcEndpointResultOutput) DnsOptions() GetVpcEndpointDnsOptionArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []GetVpcEndpointDnsOption { return v.DnsOptions }).(GetVpcEndpointDnsOptionArrayOutput)
}

func (o LookupVpcEndpointResultOutput) Filters() GetVpcEndpointFilterArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []GetVpcEndpointFilter { return v.Filters }).(GetVpcEndpointFilterArrayOutput)
}

func (o LookupVpcEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcEndpointResultOutput) IpAddressType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) string { return v.IpAddressType }).(pulumi.StringOutput)
}

// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
func (o LookupVpcEndpointResultOutput) NetworkInterfaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.NetworkInterfaceIds }).(pulumi.StringArrayOutput)
}

// ID of the AWS account that owns the VPC endpoint.
func (o LookupVpcEndpointResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

// Policy document associated with the VPC Endpoint. Applicable for endpoints of type `Gateway`.
func (o LookupVpcEndpointResultOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) string { return v.Policy }).(pulumi.StringOutput)
}

// Prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
func (o LookupVpcEndpointResultOutput) PrefixListId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) string { return v.PrefixListId }).(pulumi.StringOutput)
}

// Whether or not the VPC is associated with a private hosted zone - `true` or `false`. Applicable for endpoints of type `Interface`.
func (o LookupVpcEndpointResultOutput) PrivateDnsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) bool { return v.PrivateDnsEnabled }).(pulumi.BoolOutput)
}

// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
func (o LookupVpcEndpointResultOutput) RequesterManaged() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) bool { return v.RequesterManaged }).(pulumi.BoolOutput)
}

// One or more route tables associated with the VPC Endpoint. Applicable for endpoints of type `Gateway`.
func (o LookupVpcEndpointResultOutput) RouteTableIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.RouteTableIds }).(pulumi.StringArrayOutput)
}

// One or more security groups associated with the network interfaces. Applicable for endpoints of type `Interface`.
func (o LookupVpcEndpointResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupVpcEndpointResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o LookupVpcEndpointResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) string { return v.State }).(pulumi.StringOutput)
}

// One or more subnets in which the VPC Endpoint is located. Applicable for endpoints of type `Interface`.
func (o LookupVpcEndpointResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o LookupVpcEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// VPC Endpoint type, `Gateway` or `Interface`.
func (o LookupVpcEndpointResultOutput) VpcEndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) string { return v.VpcEndpointType }).(pulumi.StringOutput)
}

func (o LookupVpcEndpointResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcEndpointResultOutput{})
}
