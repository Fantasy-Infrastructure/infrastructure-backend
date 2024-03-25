// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The VPC Endpoint Service data source details about a specific service that
// can be specified when creating a VPC endpoint within the region configured in the provider.
//
// ## Example Usage
//
// ### AWS Service
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
//			s3, err := ec2.LookupVpcEndpointService(ctx, &ec2.LookupVpcEndpointServiceArgs{
//				Service:     pulumi.StringRef("s3"),
//				ServiceType: pulumi.StringRef("Gateway"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Create a VPC
//			foo, err := ec2.NewVpc(ctx, "foo", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			// Create a VPC endpoint
//			_, err = ec2.NewVpcEndpoint(ctx, "ep", &ec2.VpcEndpointArgs{
//				VpcId:       foo.ID(),
//				ServiceName: *pulumi.String(s3.ServiceName),
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
//
// ### Non-AWS Service
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
//			_, err := ec2.LookupVpcEndpointService(ctx, &ec2.LookupVpcEndpointServiceArgs{
//				ServiceName: pulumi.StringRef("com.amazonaws.vpce.us-west-2.vpce-svc-0e87519c997c63cd8"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Filter
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
//			_, err := ec2.LookupVpcEndpointService(ctx, &ec2.LookupVpcEndpointServiceArgs{
//				Filters: []ec2.GetVpcEndpointServiceFilter{
//					{
//						Name: "service-name",
//						Values: []string{
//							"some-service",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupVpcEndpointService(ctx *pulumi.Context, args *LookupVpcEndpointServiceArgs, opts ...pulumi.InvokeOption) (*LookupVpcEndpointServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcEndpointServiceResult
	err := ctx.Invoke("aws:ec2/getVpcEndpointService:getVpcEndpointService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcEndpointService.
type LookupVpcEndpointServiceArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetVpcEndpointServiceFilter `pulumi:"filters"`
	// Common name of an AWS service (e.g., `s3`).
	Service *string `pulumi:"service"`
	// Service name that is specified when creating a VPC endpoint. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName *string `pulumi:"serviceName"`
	// Service type, `Gateway` or `Interface`.
	ServiceType *string `pulumi:"serviceType"`
	// Map of tags, each pair of which must exactly match a pair on the desired VPC Endpoint Service.
	//
	// > **NOTE:** Specifying `service` will not work for non-AWS services or AWS services that don't follow the standard `serviceName` pattern of `com.amazonaws.<region>.<service>`.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getVpcEndpointService.
type LookupVpcEndpointServiceResult struct {
	// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
	AcceptanceRequired bool `pulumi:"acceptanceRequired"`
	// ARN of the VPC endpoint service.
	Arn string `pulumi:"arn"`
	// Availability Zones in which the service is available.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The DNS names for the service.
	BaseEndpointDnsNames []string                      `pulumi:"baseEndpointDnsNames"`
	Filters              []GetVpcEndpointServiceFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Whether or not the service manages its VPC endpoints - `true` or `false`.
	ManagesVpcEndpoints bool `pulumi:"managesVpcEndpoints"`
	// AWS account ID of the service owner or `amazon`.
	Owner string `pulumi:"owner"`
	// Private DNS name for the service.
	PrivateDnsName string  `pulumi:"privateDnsName"`
	Service        *string `pulumi:"service"`
	// ID of the endpoint service.
	ServiceId   string `pulumi:"serviceId"`
	ServiceName string `pulumi:"serviceName"`
	ServiceType string `pulumi:"serviceType"`
	// The supported IP address types.
	SupportedIpAddressTypes []string `pulumi:"supportedIpAddressTypes"`
	// Map of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not the service supports endpoint policies - `true` or `false`.
	VpcEndpointPolicySupported bool `pulumi:"vpcEndpointPolicySupported"`
}

func LookupVpcEndpointServiceOutput(ctx *pulumi.Context, args LookupVpcEndpointServiceOutputArgs, opts ...pulumi.InvokeOption) LookupVpcEndpointServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcEndpointServiceResult, error) {
			args := v.(LookupVpcEndpointServiceArgs)
			r, err := LookupVpcEndpointService(ctx, &args, opts...)
			var s LookupVpcEndpointServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcEndpointServiceResultOutput)
}

// A collection of arguments for invoking getVpcEndpointService.
type LookupVpcEndpointServiceOutputArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters GetVpcEndpointServiceFilterArrayInput `pulumi:"filters"`
	// Common name of an AWS service (e.g., `s3`).
	Service pulumi.StringPtrInput `pulumi:"service"`
	// Service name that is specified when creating a VPC endpoint. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
	// Service type, `Gateway` or `Interface`.
	ServiceType pulumi.StringPtrInput `pulumi:"serviceType"`
	// Map of tags, each pair of which must exactly match a pair on the desired VPC Endpoint Service.
	//
	// > **NOTE:** Specifying `service` will not work for non-AWS services or AWS services that don't follow the standard `serviceName` pattern of `com.amazonaws.<region>.<service>`.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupVpcEndpointServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointServiceArgs)(nil)).Elem()
}

// A collection of values returned by getVpcEndpointService.
type LookupVpcEndpointServiceResultOutput struct{ *pulumi.OutputState }

func (LookupVpcEndpointServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcEndpointServiceResult)(nil)).Elem()
}

func (o LookupVpcEndpointServiceResultOutput) ToLookupVpcEndpointServiceResultOutput() LookupVpcEndpointServiceResultOutput {
	return o
}

func (o LookupVpcEndpointServiceResultOutput) ToLookupVpcEndpointServiceResultOutputWithContext(ctx context.Context) LookupVpcEndpointServiceResultOutput {
	return o
}

// Whether or not VPC endpoint connection requests to the service must be accepted by the service owner - `true` or `false`.
func (o LookupVpcEndpointServiceResultOutput) AcceptanceRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) bool { return v.AcceptanceRequired }).(pulumi.BoolOutput)
}

// ARN of the VPC endpoint service.
func (o LookupVpcEndpointServiceResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Availability Zones in which the service is available.
func (o LookupVpcEndpointServiceResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// The DNS names for the service.
func (o LookupVpcEndpointServiceResultOutput) BaseEndpointDnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) []string { return v.BaseEndpointDnsNames }).(pulumi.StringArrayOutput)
}

func (o LookupVpcEndpointServiceResultOutput) Filters() GetVpcEndpointServiceFilterArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) []GetVpcEndpointServiceFilter { return v.Filters }).(GetVpcEndpointServiceFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcEndpointServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether or not the service manages its VPC endpoints - `true` or `false`.
func (o LookupVpcEndpointServiceResultOutput) ManagesVpcEndpoints() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) bool { return v.ManagesVpcEndpoints }).(pulumi.BoolOutput)
}

// AWS account ID of the service owner or `amazon`.
func (o LookupVpcEndpointServiceResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.Owner }).(pulumi.StringOutput)
}

// Private DNS name for the service.
func (o LookupVpcEndpointServiceResultOutput) PrivateDnsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.PrivateDnsName }).(pulumi.StringOutput)
}

func (o LookupVpcEndpointServiceResultOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) *string { return v.Service }).(pulumi.StringPtrOutput)
}

// ID of the endpoint service.
func (o LookupVpcEndpointServiceResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func (o LookupVpcEndpointServiceResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o LookupVpcEndpointServiceResultOutput) ServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) string { return v.ServiceType }).(pulumi.StringOutput)
}

// The supported IP address types.
func (o LookupVpcEndpointServiceResultOutput) SupportedIpAddressTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) []string { return v.SupportedIpAddressTypes }).(pulumi.StringArrayOutput)
}

// Map of tags assigned to the resource.
func (o LookupVpcEndpointServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Whether or not the service supports endpoint policies - `true` or `false`.
func (o LookupVpcEndpointServiceResultOutput) VpcEndpointPolicySupported() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcEndpointServiceResult) bool { return v.VpcEndpointPolicySupported }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcEndpointServiceResultOutput{})
}
