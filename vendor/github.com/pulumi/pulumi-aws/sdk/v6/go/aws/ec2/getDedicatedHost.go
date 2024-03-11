// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about an EC2 Dedicated Host.
//
// ## Example Usage
//
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
//			testDedicatedHost, err := ec2.NewDedicatedHost(ctx, "test", &ec2.DedicatedHostArgs{
//				InstanceType:     pulumi.String("c5.18xlarge"),
//				AvailabilityZone: pulumi.String("us-west-2a"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = ec2.LookupDedicatedHostOutput(ctx, ec2.GetDedicatedHostOutputArgs{
//				HostId: testDedicatedHost.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// ### Filter Example
//
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
//			_, err := ec2.LookupDedicatedHost(ctx, &ec2.LookupDedicatedHostArgs{
//				Filters: []ec2.GetDedicatedHostFilter{
//					{
//						Name: "instance-type",
//						Values: []string{
//							"c5.18xlarge",
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
func LookupDedicatedHost(ctx *pulumi.Context, args *LookupDedicatedHostArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHostResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDedicatedHostResult
	err := ctx.Invoke("aws:ec2/getDedicatedHost:getDedicatedHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDedicatedHost.
type LookupDedicatedHostArgs struct {
	// Configuration block. Detailed below.
	Filters []GetDedicatedHostFilter `pulumi:"filters"`
	// ID of the Dedicated Host.
	HostId *string           `pulumi:"hostId"`
	Tags   map[string]string `pulumi:"tags"`
}

// A collection of values returned by getDedicatedHost.
type LookupDedicatedHostResult struct {
	// ARN of the Dedicated Host.
	Arn string `pulumi:"arn"`
	// The ID of the Outpost hardware asset on which the Dedicated Host is allocated.
	AssetId string `pulumi:"assetId"`
	// Whether auto-placement is on or off.
	AutoPlacement string `pulumi:"autoPlacement"`
	// Availability Zone of the Dedicated Host.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Number of cores on the Dedicated Host.
	Cores   int                      `pulumi:"cores"`
	Filters []GetDedicatedHostFilter `pulumi:"filters"`
	HostId  string                   `pulumi:"hostId"`
	// Whether host recovery is enabled or disabled for the Dedicated Host.
	HostRecovery string `pulumi:"hostRecovery"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Instance family supported by the Dedicated Host. For example, "m5".
	InstanceFamily string `pulumi:"instanceFamily"`
	// Instance type supported by the Dedicated Host. For example, "m5.large". If the host supports multiple instance types, no instanceType is returned.
	InstanceType string `pulumi:"instanceType"`
	// ARN of the AWS Outpost on which the Dedicated Host is allocated.
	OutpostArn string `pulumi:"outpostArn"`
	// ID of the AWS account that owns the Dedicated Host.
	OwnerId string `pulumi:"ownerId"`
	// Number of sockets on the Dedicated Host.
	Sockets int               `pulumi:"sockets"`
	Tags    map[string]string `pulumi:"tags"`
	// Total number of vCPUs on the Dedicated Host.
	TotalVcpus int `pulumi:"totalVcpus"`
}

func LookupDedicatedHostOutput(ctx *pulumi.Context, args LookupDedicatedHostOutputArgs, opts ...pulumi.InvokeOption) LookupDedicatedHostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDedicatedHostResult, error) {
			args := v.(LookupDedicatedHostArgs)
			r, err := LookupDedicatedHost(ctx, &args, opts...)
			var s LookupDedicatedHostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDedicatedHostResultOutput)
}

// A collection of arguments for invoking getDedicatedHost.
type LookupDedicatedHostOutputArgs struct {
	// Configuration block. Detailed below.
	Filters GetDedicatedHostFilterArrayInput `pulumi:"filters"`
	// ID of the Dedicated Host.
	HostId pulumi.StringPtrInput `pulumi:"hostId"`
	Tags   pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupDedicatedHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHostArgs)(nil)).Elem()
}

// A collection of values returned by getDedicatedHost.
type LookupDedicatedHostResultOutput struct{ *pulumi.OutputState }

func (LookupDedicatedHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDedicatedHostResult)(nil)).Elem()
}

func (o LookupDedicatedHostResultOutput) ToLookupDedicatedHostResultOutput() LookupDedicatedHostResultOutput {
	return o
}

func (o LookupDedicatedHostResultOutput) ToLookupDedicatedHostResultOutputWithContext(ctx context.Context) LookupDedicatedHostResultOutput {
	return o
}

// ARN of the Dedicated Host.
func (o LookupDedicatedHostResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The ID of the Outpost hardware asset on which the Dedicated Host is allocated.
func (o LookupDedicatedHostResultOutput) AssetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.AssetId }).(pulumi.StringOutput)
}

// Whether auto-placement is on or off.
func (o LookupDedicatedHostResultOutput) AutoPlacement() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.AutoPlacement }).(pulumi.StringOutput)
}

// Availability Zone of the Dedicated Host.
func (o LookupDedicatedHostResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Number of cores on the Dedicated Host.
func (o LookupDedicatedHostResultOutput) Cores() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) int { return v.Cores }).(pulumi.IntOutput)
}

func (o LookupDedicatedHostResultOutput) Filters() GetDedicatedHostFilterArrayOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) []GetDedicatedHostFilter { return v.Filters }).(GetDedicatedHostFilterArrayOutput)
}

func (o LookupDedicatedHostResultOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.HostId }).(pulumi.StringOutput)
}

// Whether host recovery is enabled or disabled for the Dedicated Host.
func (o LookupDedicatedHostResultOutput) HostRecovery() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.HostRecovery }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDedicatedHostResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.Id }).(pulumi.StringOutput)
}

// Instance family supported by the Dedicated Host. For example, "m5".
func (o LookupDedicatedHostResultOutput) InstanceFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.InstanceFamily }).(pulumi.StringOutput)
}

// Instance type supported by the Dedicated Host. For example, "m5.large". If the host supports multiple instance types, no instanceType is returned.
func (o LookupDedicatedHostResultOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.InstanceType }).(pulumi.StringOutput)
}

// ARN of the AWS Outpost on which the Dedicated Host is allocated.
func (o LookupDedicatedHostResultOutput) OutpostArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.OutpostArn }).(pulumi.StringOutput)
}

// ID of the AWS account that owns the Dedicated Host.
func (o LookupDedicatedHostResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

// Number of sockets on the Dedicated Host.
func (o LookupDedicatedHostResultOutput) Sockets() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) int { return v.Sockets }).(pulumi.IntOutput)
}

func (o LookupDedicatedHostResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Total number of vCPUs on the Dedicated Host.
func (o LookupDedicatedHostResultOutput) TotalVcpus() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDedicatedHostResult) int { return v.TotalVcpus }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDedicatedHostResultOutput{})
}
