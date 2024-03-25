// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `route53.ResolverEndpoint` provides details about a specific Route53 Resolver Endpoint.
//
// This data source allows to find a list of IPaddresses associated with a specific Route53 Resolver Endpoint.
//
// ## Example Usage
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
//			_, err := route53.LookupResolverEndpoint(ctx, &route53.LookupResolverEndpointArgs{
//				ResolverEndpointId: pulumi.StringRef("rslvr-in-1abc2345ef678g91h"),
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
//			_, err := route53.LookupResolverEndpoint(ctx, &route53.LookupResolverEndpointArgs{
//				Filters: []route53.GetResolverEndpointFilter{
//					{
//						Name: "NAME",
//						Values: []string{
//							"MyResolverExampleName",
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
func LookupResolverEndpoint(ctx *pulumi.Context, args *LookupResolverEndpointArgs, opts ...pulumi.InvokeOption) (*LookupResolverEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResolverEndpointResult
	err := ctx.Invoke("aws:route53/getResolverEndpoint:getResolverEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResolverEndpoint.
type LookupResolverEndpointArgs struct {
	// One or more name/value pairs to use as filters. There are
	// several valid keys, for a full reference, check out
	// [Route53resolver Filter value in the AWS API reference][1].
	//
	// In addition to all arguments above, the following attributes are exported:
	Filters []GetResolverEndpointFilter `pulumi:"filters"`
	// ID of the Route53 Resolver Endpoint.
	ResolverEndpointId *string `pulumi:"resolverEndpointId"`
}

// A collection of values returned by getResolverEndpoint.
type LookupResolverEndpointResult struct {
	Arn       string                      `pulumi:"arn"`
	Direction string                      `pulumi:"direction"`
	Filters   []GetResolverEndpointFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string   `pulumi:"id"`
	IpAddresses          []string `pulumi:"ipAddresses"`
	Name                 string   `pulumi:"name"`
	Protocols            []string `pulumi:"protocols"`
	ResolverEndpointId   *string  `pulumi:"resolverEndpointId"`
	ResolverEndpointType string   `pulumi:"resolverEndpointType"`
	Status               string   `pulumi:"status"`
	VpcId                string   `pulumi:"vpcId"`
}

func LookupResolverEndpointOutput(ctx *pulumi.Context, args LookupResolverEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupResolverEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResolverEndpointResult, error) {
			args := v.(LookupResolverEndpointArgs)
			r, err := LookupResolverEndpoint(ctx, &args, opts...)
			var s LookupResolverEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResolverEndpointResultOutput)
}

// A collection of arguments for invoking getResolverEndpoint.
type LookupResolverEndpointOutputArgs struct {
	// One or more name/value pairs to use as filters. There are
	// several valid keys, for a full reference, check out
	// [Route53resolver Filter value in the AWS API reference][1].
	//
	// In addition to all arguments above, the following attributes are exported:
	Filters GetResolverEndpointFilterArrayInput `pulumi:"filters"`
	// ID of the Route53 Resolver Endpoint.
	ResolverEndpointId pulumi.StringPtrInput `pulumi:"resolverEndpointId"`
}

func (LookupResolverEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverEndpointArgs)(nil)).Elem()
}

// A collection of values returned by getResolverEndpoint.
type LookupResolverEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupResolverEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverEndpointResult)(nil)).Elem()
}

func (o LookupResolverEndpointResultOutput) ToLookupResolverEndpointResultOutput() LookupResolverEndpointResultOutput {
	return o
}

func (o LookupResolverEndpointResultOutput) ToLookupResolverEndpointResultOutputWithContext(ctx context.Context) LookupResolverEndpointResultOutput {
	return o
}

func (o LookupResolverEndpointResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResolverEndpointResult) string { return v.Arn }).(pulumi.StringOutput)
}

func (o LookupResolverEndpointResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResolverEndpointResult) string { return v.Direction }).(pulumi.StringOutput)
}

func (o LookupResolverEndpointResultOutput) Filters() GetResolverEndpointFilterArrayOutput {
	return o.ApplyT(func(v LookupResolverEndpointResult) []GetResolverEndpointFilter { return v.Filters }).(GetResolverEndpointFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupResolverEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResolverEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResolverEndpointResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupResolverEndpointResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupResolverEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResolverEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupResolverEndpointResultOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupResolverEndpointResult) []string { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o LookupResolverEndpointResultOutput) ResolverEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverEndpointResult) *string { return v.ResolverEndpointId }).(pulumi.StringPtrOutput)
}

func (o LookupResolverEndpointResultOutput) ResolverEndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResolverEndpointResult) string { return v.ResolverEndpointType }).(pulumi.StringOutput)
}

func (o LookupResolverEndpointResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResolverEndpointResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupResolverEndpointResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResolverEndpointResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResolverEndpointResultOutput{})
}
