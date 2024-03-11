// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can be useful for getting back a set of subnet IDs.
//
// ## Example Usage
//
// The following shows outputting all CIDR blocks for every subnet ID in a VPC.
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
//	func notImplemented(message string) pulumi.AnyOutput {
//	  panic(message)
//	}
//
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _, err := ec2.GetSubnets(ctx, &ec2.GetSubnetsArgs{
// Filters: []ec2.GetSubnetsFilter{
// {
// Name: "vpc-id",
// Values: interface{}{
// vpcId,
// },
// },
// },
// }, nil);
// if err != nil {
// return err
// }
// exampleGetSubnet := "TODO: For expression";
// ctx.Export("subnetCidrBlocks", "TODO: For expression")
// return nil
// })
// }
// ```
//
// The following example retrieves a set of all subnets in a VPC with a custom
// tag of `Tier` set to a value of "Private" so that the `ec2.Instance` resource
// can loop through the subnets, putting instances across availability zones.
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
//	func notImplemented(message string) pulumi.AnyOutput {
//	  panic(message)
//	}
//
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _, err := ec2.GetSubnets(ctx, &ec2.GetSubnetsArgs{
// Filters: []ec2.GetSubnetsFilter{
// {
// Name: "vpc-id",
// Values: interface{}{
// vpcId,
// },
// },
// },
// Tags: map[string]interface{}{
// "Tier": "Private",
// },
// }, nil);
// if err != nil {
// return err
// }
// var app []*ec2.Instance
//
//	for index := 0; index < notImplemented("toset(data.aws_subnets.private.ids)"); index++ {
//	    key0 := index
//	    val0 := index
//
// __res, err := ec2.NewInstance(ctx, fmt.Sprintf("app-%v", key0), &ec2.InstanceArgs{
// Ami: pulumi.Any(ami),
// InstanceType: pulumi.String("t2.micro"),
// SubnetId: pulumi.Any(val0),
// })
// if err != nil {
// return err
// }
// app = append(app, __res)
// }
// return nil
// })
// }
// ```
func GetSubnets(ctx *pulumi.Context, args *GetSubnetsArgs, opts ...pulumi.InvokeOption) (*GetSubnetsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSubnetsResult
	err := ctx.Invoke("aws:ec2/getSubnets:getSubnets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnets.
type GetSubnetsArgs struct {
	// Custom filter block as described below.
	Filters []GetSubnetsFilter `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired subnets.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getSubnets.
type GetSubnetsResult struct {
	Filters []GetSubnetsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of all the subnet ids found.
	Ids  []string          `pulumi:"ids"`
	Tags map[string]string `pulumi:"tags"`
}

func GetSubnetsOutput(ctx *pulumi.Context, args GetSubnetsOutputArgs, opts ...pulumi.InvokeOption) GetSubnetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSubnetsResult, error) {
			args := v.(GetSubnetsArgs)
			r, err := GetSubnets(ctx, &args, opts...)
			var s GetSubnetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSubnetsResultOutput)
}

// A collection of arguments for invoking getSubnets.
type GetSubnetsOutputArgs struct {
	// Custom filter block as described below.
	Filters GetSubnetsFilterArrayInput `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired subnets.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetSubnetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetsArgs)(nil)).Elem()
}

// A collection of values returned by getSubnets.
type GetSubnetsResultOutput struct{ *pulumi.OutputState }

func (GetSubnetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetsResult)(nil)).Elem()
}

func (o GetSubnetsResultOutput) ToGetSubnetsResultOutput() GetSubnetsResultOutput {
	return o
}

func (o GetSubnetsResultOutput) ToGetSubnetsResultOutputWithContext(ctx context.Context) GetSubnetsResultOutput {
	return o
}

func (o GetSubnetsResultOutput) Filters() GetSubnetsFilterArrayOutput {
	return o.ApplyT(func(v GetSubnetsResult) []GetSubnetsFilter { return v.Filters }).(GetSubnetsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSubnetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of all the subnet ids found.
func (o GetSubnetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSubnetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetSubnetsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSubnetsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSubnetsResultOutput{})
}
