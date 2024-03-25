// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Information about most recent Spot Price for a given EC2 instance.
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
//			_, err := ec2.GetSpotPrice(ctx, &ec2.GetSpotPriceArgs{
//				InstanceType:     pulumi.StringRef("t3.medium"),
//				AvailabilityZone: pulumi.StringRef("us-west-2a"),
//				Filters: []ec2.GetSpotPriceFilter{
//					{
//						Name: "product-description",
//						Values: []string{
//							"Linux/UNIX",
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
func GetSpotPrice(ctx *pulumi.Context, args *GetSpotPriceArgs, opts ...pulumi.InvokeOption) (*GetSpotPriceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSpotPriceResult
	err := ctx.Invoke("aws:ec2/getSpotPrice:getSpotPrice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSpotPrice.
type GetSpotPriceArgs struct {
	// Availability zone in which to query Spot price information.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSpotPriceHistory.html) for supported filters. Detailed below.
	Filters []GetSpotPriceFilter `pulumi:"filters"`
	// Type of instance for which to query Spot Price information.
	InstanceType *string `pulumi:"instanceType"`
}

// A collection of values returned by getSpotPrice.
type GetSpotPriceResult struct {
	AvailabilityZone *string              `pulumi:"availabilityZone"`
	Filters          []GetSpotPriceFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	InstanceType *string `pulumi:"instanceType"`
	// Most recent Spot Price value for the given instance type and AZ.
	SpotPrice string `pulumi:"spotPrice"`
	// The timestamp at which the Spot Price value was published.
	SpotPriceTimestamp string `pulumi:"spotPriceTimestamp"`
}

func GetSpotPriceOutput(ctx *pulumi.Context, args GetSpotPriceOutputArgs, opts ...pulumi.InvokeOption) GetSpotPriceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSpotPriceResult, error) {
			args := v.(GetSpotPriceArgs)
			r, err := GetSpotPrice(ctx, &args, opts...)
			var s GetSpotPriceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSpotPriceResultOutput)
}

// A collection of arguments for invoking getSpotPrice.
type GetSpotPriceOutputArgs struct {
	// Availability zone in which to query Spot price information.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSpotPriceHistory.html) for supported filters. Detailed below.
	Filters GetSpotPriceFilterArrayInput `pulumi:"filters"`
	// Type of instance for which to query Spot Price information.
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
}

func (GetSpotPriceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSpotPriceArgs)(nil)).Elem()
}

// A collection of values returned by getSpotPrice.
type GetSpotPriceResultOutput struct{ *pulumi.OutputState }

func (GetSpotPriceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSpotPriceResult)(nil)).Elem()
}

func (o GetSpotPriceResultOutput) ToGetSpotPriceResultOutput() GetSpotPriceResultOutput {
	return o
}

func (o GetSpotPriceResultOutput) ToGetSpotPriceResultOutputWithContext(ctx context.Context) GetSpotPriceResultOutput {
	return o
}

func (o GetSpotPriceResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSpotPriceResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o GetSpotPriceResultOutput) Filters() GetSpotPriceFilterArrayOutput {
	return o.ApplyT(func(v GetSpotPriceResult) []GetSpotPriceFilter { return v.Filters }).(GetSpotPriceFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSpotPriceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSpotPriceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSpotPriceResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSpotPriceResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// Most recent Spot Price value for the given instance type and AZ.
func (o GetSpotPriceResultOutput) SpotPrice() pulumi.StringOutput {
	return o.ApplyT(func(v GetSpotPriceResult) string { return v.SpotPrice }).(pulumi.StringOutput)
}

// The timestamp at which the Spot Price value was published.
func (o GetSpotPriceResultOutput) SpotPriceTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v GetSpotPriceResult) string { return v.SpotPriceTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSpotPriceResultOutput{})
}
