// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Previews a CIDR from an IPAM address pool. Only works for private IPv4.
//
// > **NOTE:** This functionality is also encapsulated in a resource sharing the same name. The data source can be used when you need to use the cidr in a calculation of the same Root module, `count` for example. However, once a cidr range has been allocated that was previewed, the next refresh will find a **new** cidr and may force new resources downstream. Make sure to use `ignoreChanges` if this is undesirable.
//
// ## Example Usage
//
// Basic usage:
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
//			test, err := ec2.GetIpamPreviewNextCidr(ctx, &ec2.GetIpamPreviewNextCidrArgs{
//				IpamPoolId:    testAwsVpcIpamPool.Id,
//				NetmaskLength: pulumi.IntRef(28),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcIpamPoolCidrAllocation(ctx, "test", &ec2.VpcIpamPoolCidrAllocationArgs{
//				IpamPoolId: pulumi.Any(testAwsVpcIpamPool.Id),
//				Cidr:       *pulumi.String(test.Cidr),
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
func GetIpamPreviewNextCidr(ctx *pulumi.Context, args *GetIpamPreviewNextCidrArgs, opts ...pulumi.InvokeOption) (*GetIpamPreviewNextCidrResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIpamPreviewNextCidrResult
	err := ctx.Invoke("aws:ec2/getIpamPreviewNextCidr:getIpamPreviewNextCidr", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpamPreviewNextCidr.
type GetIpamPreviewNextCidrArgs struct {
	// Exclude a particular CIDR range from being returned by the pool.
	DisallowedCidrs []string `pulumi:"disallowedCidrs"`
	// ID of the pool to which you want to assign a CIDR.
	IpamPoolId string `pulumi:"ipamPoolId"`
	// Netmask length of the CIDR you would like to preview from the IPAM pool.
	NetmaskLength *int `pulumi:"netmaskLength"`
}

// A collection of values returned by getIpamPreviewNextCidr.
type GetIpamPreviewNextCidrResult struct {
	// Previewed CIDR from the pool.
	Cidr            string   `pulumi:"cidr"`
	DisallowedCidrs []string `pulumi:"disallowedCidrs"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	IpamPoolId    string `pulumi:"ipamPoolId"`
	NetmaskLength *int   `pulumi:"netmaskLength"`
}

func GetIpamPreviewNextCidrOutput(ctx *pulumi.Context, args GetIpamPreviewNextCidrOutputArgs, opts ...pulumi.InvokeOption) GetIpamPreviewNextCidrResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIpamPreviewNextCidrResult, error) {
			args := v.(GetIpamPreviewNextCidrArgs)
			r, err := GetIpamPreviewNextCidr(ctx, &args, opts...)
			var s GetIpamPreviewNextCidrResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIpamPreviewNextCidrResultOutput)
}

// A collection of arguments for invoking getIpamPreviewNextCidr.
type GetIpamPreviewNextCidrOutputArgs struct {
	// Exclude a particular CIDR range from being returned by the pool.
	DisallowedCidrs pulumi.StringArrayInput `pulumi:"disallowedCidrs"`
	// ID of the pool to which you want to assign a CIDR.
	IpamPoolId pulumi.StringInput `pulumi:"ipamPoolId"`
	// Netmask length of the CIDR you would like to preview from the IPAM pool.
	NetmaskLength pulumi.IntPtrInput `pulumi:"netmaskLength"`
}

func (GetIpamPreviewNextCidrOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpamPreviewNextCidrArgs)(nil)).Elem()
}

// A collection of values returned by getIpamPreviewNextCidr.
type GetIpamPreviewNextCidrResultOutput struct{ *pulumi.OutputState }

func (GetIpamPreviewNextCidrResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpamPreviewNextCidrResult)(nil)).Elem()
}

func (o GetIpamPreviewNextCidrResultOutput) ToGetIpamPreviewNextCidrResultOutput() GetIpamPreviewNextCidrResultOutput {
	return o
}

func (o GetIpamPreviewNextCidrResultOutput) ToGetIpamPreviewNextCidrResultOutputWithContext(ctx context.Context) GetIpamPreviewNextCidrResultOutput {
	return o
}

// Previewed CIDR from the pool.
func (o GetIpamPreviewNextCidrResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamPreviewNextCidrResult) string { return v.Cidr }).(pulumi.StringOutput)
}

func (o GetIpamPreviewNextCidrResultOutput) DisallowedCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpamPreviewNextCidrResult) []string { return v.DisallowedCidrs }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpamPreviewNextCidrResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamPreviewNextCidrResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIpamPreviewNextCidrResultOutput) IpamPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpamPreviewNextCidrResult) string { return v.IpamPoolId }).(pulumi.StringOutput)
}

func (o GetIpamPreviewNextCidrResultOutput) NetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetIpamPreviewNextCidrResult) *int { return v.NetmaskLength }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpamPreviewNextCidrResultOutput{})
}
