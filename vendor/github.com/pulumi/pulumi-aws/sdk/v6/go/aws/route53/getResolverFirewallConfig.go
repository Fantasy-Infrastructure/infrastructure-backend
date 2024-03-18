// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `route53.ResolverFirewallConfig` provides details about a specific a Route 53 Resolver DNS Firewall config.
//
// This data source allows to find a details about a specific a Route 53 Resolver DNS Firewall config.
//
// ## Example Usage
//
// The following example shows how to get a firewall config using the VPC ID.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := route53.LookupResolverFirewallConfig(ctx, &route53.LookupResolverFirewallConfigArgs{
//				ResourceId: "vpc-exampleid",
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
func LookupResolverFirewallConfig(ctx *pulumi.Context, args *LookupResolverFirewallConfigArgs, opts ...pulumi.InvokeOption) (*LookupResolverFirewallConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResolverFirewallConfigResult
	err := ctx.Invoke("aws:route53/getResolverFirewallConfig:getResolverFirewallConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResolverFirewallConfig.
type LookupResolverFirewallConfigArgs struct {
	// The ID of the VPC from Amazon VPC that the configuration is for.
	//
	// The following attribute is additionally exported:
	ResourceId string `pulumi:"resourceId"`
}

// A collection of values returned by getResolverFirewallConfig.
type LookupResolverFirewallConfigResult struct {
	FirewallFailOpen string `pulumi:"firewallFailOpen"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	OwnerId    string `pulumi:"ownerId"`
	ResourceId string `pulumi:"resourceId"`
}

func LookupResolverFirewallConfigOutput(ctx *pulumi.Context, args LookupResolverFirewallConfigOutputArgs, opts ...pulumi.InvokeOption) LookupResolverFirewallConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResolverFirewallConfigResult, error) {
			args := v.(LookupResolverFirewallConfigArgs)
			r, err := LookupResolverFirewallConfig(ctx, &args, opts...)
			var s LookupResolverFirewallConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResolverFirewallConfigResultOutput)
}

// A collection of arguments for invoking getResolverFirewallConfig.
type LookupResolverFirewallConfigOutputArgs struct {
	// The ID of the VPC from Amazon VPC that the configuration is for.
	//
	// The following attribute is additionally exported:
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
}

func (LookupResolverFirewallConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverFirewallConfigArgs)(nil)).Elem()
}

// A collection of values returned by getResolverFirewallConfig.
type LookupResolverFirewallConfigResultOutput struct{ *pulumi.OutputState }

func (LookupResolverFirewallConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverFirewallConfigResult)(nil)).Elem()
}

func (o LookupResolverFirewallConfigResultOutput) ToLookupResolverFirewallConfigResultOutput() LookupResolverFirewallConfigResultOutput {
	return o
}

func (o LookupResolverFirewallConfigResultOutput) ToLookupResolverFirewallConfigResultOutputWithContext(ctx context.Context) LookupResolverFirewallConfigResultOutput {
	return o
}

func (o LookupResolverFirewallConfigResultOutput) FirewallFailOpen() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResolverFirewallConfigResult) string { return v.FirewallFailOpen }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupResolverFirewallConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResolverFirewallConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResolverFirewallConfigResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResolverFirewallConfigResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o LookupResolverFirewallConfigResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResolverFirewallConfigResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResolverFirewallConfigResultOutput{})
}
