// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Route 53 Resolver config resource.
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
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := ec2.NewVpc(ctx, "example", &ec2.VpcArgs{
//				CidrBlock:          pulumi.String("10.0.0.0/16"),
//				EnableDnsSupport:   pulumi.Bool(true),
//				EnableDnsHostnames: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = route53.NewResolverConfig(ctx, "example", &route53.ResolverConfigArgs{
//				ResourceId:             example.ID(),
//				AutodefinedReverseFlag: pulumi.String("DISABLE"),
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
// ## Import
//
// Using `pulumi import`, import Route 53 Resolver configs using the Route 53 Resolver config ID. For example:
//
// ```sh
// $ pulumi import aws:route53/resolverConfig:ResolverConfig example rslvr-rc-715aa20c73a23da7
// ```
type ResolverConfig struct {
	pulumi.CustomResourceState

	// Indicates whether or not the Resolver will create autodefined rules for reverse DNS lookups. Valid values: `ENABLE`, `DISABLE`.
	AutodefinedReverseFlag pulumi.StringOutput `pulumi:"autodefinedReverseFlag"`
	// The AWS account ID of the owner of the VPC that this resolver configuration applies to.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The ID of the VPC that the configuration is for.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
}

// NewResolverConfig registers a new resource with the given unique name, arguments, and options.
func NewResolverConfig(ctx *pulumi.Context,
	name string, args *ResolverConfigArgs, opts ...pulumi.ResourceOption) (*ResolverConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutodefinedReverseFlag == nil {
		return nil, errors.New("invalid value for required argument 'AutodefinedReverseFlag'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResolverConfig
	err := ctx.RegisterResource("aws:route53/resolverConfig:ResolverConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverConfig gets an existing ResolverConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverConfigState, opts ...pulumi.ResourceOption) (*ResolverConfig, error) {
	var resource ResolverConfig
	err := ctx.ReadResource("aws:route53/resolverConfig:ResolverConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverConfig resources.
type resolverConfigState struct {
	// Indicates whether or not the Resolver will create autodefined rules for reverse DNS lookups. Valid values: `ENABLE`, `DISABLE`.
	AutodefinedReverseFlag *string `pulumi:"autodefinedReverseFlag"`
	// The AWS account ID of the owner of the VPC that this resolver configuration applies to.
	OwnerId *string `pulumi:"ownerId"`
	// The ID of the VPC that the configuration is for.
	ResourceId *string `pulumi:"resourceId"`
}

type ResolverConfigState struct {
	// Indicates whether or not the Resolver will create autodefined rules for reverse DNS lookups. Valid values: `ENABLE`, `DISABLE`.
	AutodefinedReverseFlag pulumi.StringPtrInput
	// The AWS account ID of the owner of the VPC that this resolver configuration applies to.
	OwnerId pulumi.StringPtrInput
	// The ID of the VPC that the configuration is for.
	ResourceId pulumi.StringPtrInput
}

func (ResolverConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverConfigState)(nil)).Elem()
}

type resolverConfigArgs struct {
	// Indicates whether or not the Resolver will create autodefined rules for reverse DNS lookups. Valid values: `ENABLE`, `DISABLE`.
	AutodefinedReverseFlag string `pulumi:"autodefinedReverseFlag"`
	// The ID of the VPC that the configuration is for.
	ResourceId string `pulumi:"resourceId"`
}

// The set of arguments for constructing a ResolverConfig resource.
type ResolverConfigArgs struct {
	// Indicates whether or not the Resolver will create autodefined rules for reverse DNS lookups. Valid values: `ENABLE`, `DISABLE`.
	AutodefinedReverseFlag pulumi.StringInput
	// The ID of the VPC that the configuration is for.
	ResourceId pulumi.StringInput
}

func (ResolverConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverConfigArgs)(nil)).Elem()
}

type ResolverConfigInput interface {
	pulumi.Input

	ToResolverConfigOutput() ResolverConfigOutput
	ToResolverConfigOutputWithContext(ctx context.Context) ResolverConfigOutput
}

func (*ResolverConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverConfig)(nil)).Elem()
}

func (i *ResolverConfig) ToResolverConfigOutput() ResolverConfigOutput {
	return i.ToResolverConfigOutputWithContext(context.Background())
}

func (i *ResolverConfig) ToResolverConfigOutputWithContext(ctx context.Context) ResolverConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverConfigOutput)
}

// ResolverConfigArrayInput is an input type that accepts ResolverConfigArray and ResolverConfigArrayOutput values.
// You can construct a concrete instance of `ResolverConfigArrayInput` via:
//
//	ResolverConfigArray{ ResolverConfigArgs{...} }
type ResolverConfigArrayInput interface {
	pulumi.Input

	ToResolverConfigArrayOutput() ResolverConfigArrayOutput
	ToResolverConfigArrayOutputWithContext(context.Context) ResolverConfigArrayOutput
}

type ResolverConfigArray []ResolverConfigInput

func (ResolverConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverConfig)(nil)).Elem()
}

func (i ResolverConfigArray) ToResolverConfigArrayOutput() ResolverConfigArrayOutput {
	return i.ToResolverConfigArrayOutputWithContext(context.Background())
}

func (i ResolverConfigArray) ToResolverConfigArrayOutputWithContext(ctx context.Context) ResolverConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverConfigArrayOutput)
}

// ResolverConfigMapInput is an input type that accepts ResolverConfigMap and ResolverConfigMapOutput values.
// You can construct a concrete instance of `ResolverConfigMapInput` via:
//
//	ResolverConfigMap{ "key": ResolverConfigArgs{...} }
type ResolverConfigMapInput interface {
	pulumi.Input

	ToResolverConfigMapOutput() ResolverConfigMapOutput
	ToResolverConfigMapOutputWithContext(context.Context) ResolverConfigMapOutput
}

type ResolverConfigMap map[string]ResolverConfigInput

func (ResolverConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverConfig)(nil)).Elem()
}

func (i ResolverConfigMap) ToResolverConfigMapOutput() ResolverConfigMapOutput {
	return i.ToResolverConfigMapOutputWithContext(context.Background())
}

func (i ResolverConfigMap) ToResolverConfigMapOutputWithContext(ctx context.Context) ResolverConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverConfigMapOutput)
}

type ResolverConfigOutput struct{ *pulumi.OutputState }

func (ResolverConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverConfig)(nil)).Elem()
}

func (o ResolverConfigOutput) ToResolverConfigOutput() ResolverConfigOutput {
	return o
}

func (o ResolverConfigOutput) ToResolverConfigOutputWithContext(ctx context.Context) ResolverConfigOutput {
	return o
}

// Indicates whether or not the Resolver will create autodefined rules for reverse DNS lookups. Valid values: `ENABLE`, `DISABLE`.
func (o ResolverConfigOutput) AutodefinedReverseFlag() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverConfig) pulumi.StringOutput { return v.AutodefinedReverseFlag }).(pulumi.StringOutput)
}

// The AWS account ID of the owner of the VPC that this resolver configuration applies to.
func (o ResolverConfigOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverConfig) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The ID of the VPC that the configuration is for.
func (o ResolverConfigOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverConfig) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

type ResolverConfigArrayOutput struct{ *pulumi.OutputState }

func (ResolverConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverConfig)(nil)).Elem()
}

func (o ResolverConfigArrayOutput) ToResolverConfigArrayOutput() ResolverConfigArrayOutput {
	return o
}

func (o ResolverConfigArrayOutput) ToResolverConfigArrayOutputWithContext(ctx context.Context) ResolverConfigArrayOutput {
	return o
}

func (o ResolverConfigArrayOutput) Index(i pulumi.IntInput) ResolverConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResolverConfig {
		return vs[0].([]*ResolverConfig)[vs[1].(int)]
	}).(ResolverConfigOutput)
}

type ResolverConfigMapOutput struct{ *pulumi.OutputState }

func (ResolverConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverConfig)(nil)).Elem()
}

func (o ResolverConfigMapOutput) ToResolverConfigMapOutput() ResolverConfigMapOutput {
	return o
}

func (o ResolverConfigMapOutput) ToResolverConfigMapOutputWithContext(ctx context.Context) ResolverConfigMapOutput {
	return o
}

func (o ResolverConfigMapOutput) MapIndex(k pulumi.StringInput) ResolverConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResolverConfig {
		return vs[0].(map[string]*ResolverConfig)[vs[1].(string)]
	}).(ResolverConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverConfigInput)(nil)).Elem(), &ResolverConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverConfigArrayInput)(nil)).Elem(), ResolverConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverConfigMapInput)(nil)).Elem(), ResolverConfigMap{})
	pulumi.RegisterOutputType(ResolverConfigOutput{})
	pulumi.RegisterOutputType(ResolverConfigArrayOutput{})
	pulumi.RegisterOutputType(ResolverConfigMapOutput{})
}
