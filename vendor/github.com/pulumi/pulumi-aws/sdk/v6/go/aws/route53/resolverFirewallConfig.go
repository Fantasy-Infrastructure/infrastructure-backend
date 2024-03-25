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

// Provides a Route 53 Resolver DNS Firewall config resource.
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
//			_, err = route53.NewResolverFirewallConfig(ctx, "example", &route53.ResolverFirewallConfigArgs{
//				ResourceId:       example.ID(),
//				FirewallFailOpen: pulumi.String("ENABLED"),
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
// Using `pulumi import`, import Route 53 Resolver DNS Firewall configs using the Route 53 Resolver DNS Firewall config ID. For example:
//
// ```sh
// $ pulumi import aws:route53/resolverFirewallConfig:ResolverFirewallConfig example rdsc-be1866ecc1683e95
// ```
type ResolverFirewallConfig struct {
	pulumi.CustomResourceState

	// Determines how Route 53 Resolver handles queries during failures, for example when all traffic that is sent to DNS Firewall fails to receive a reply. By default, fail open is disabled, which means the failure mode is closed. This approach favors security over availability. DNS Firewall blocks queries that it is unable to evaluate properly. If you enable this option, the failure mode is open. This approach favors availability over security. DNS Firewall allows queries to proceed if it is unable to properly evaluate them. Valid values: `ENABLED`, `DISABLED`.
	FirewallFailOpen pulumi.StringOutput `pulumi:"firewallFailOpen"`
	// The AWS account ID of the owner of the VPC that this firewall configuration applies to.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The ID of the VPC that the configuration is for.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
}

// NewResolverFirewallConfig registers a new resource with the given unique name, arguments, and options.
func NewResolverFirewallConfig(ctx *pulumi.Context,
	name string, args *ResolverFirewallConfigArgs, opts ...pulumi.ResourceOption) (*ResolverFirewallConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResolverFirewallConfig
	err := ctx.RegisterResource("aws:route53/resolverFirewallConfig:ResolverFirewallConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverFirewallConfig gets an existing ResolverFirewallConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverFirewallConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverFirewallConfigState, opts ...pulumi.ResourceOption) (*ResolverFirewallConfig, error) {
	var resource ResolverFirewallConfig
	err := ctx.ReadResource("aws:route53/resolverFirewallConfig:ResolverFirewallConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverFirewallConfig resources.
type resolverFirewallConfigState struct {
	// Determines how Route 53 Resolver handles queries during failures, for example when all traffic that is sent to DNS Firewall fails to receive a reply. By default, fail open is disabled, which means the failure mode is closed. This approach favors security over availability. DNS Firewall blocks queries that it is unable to evaluate properly. If you enable this option, the failure mode is open. This approach favors availability over security. DNS Firewall allows queries to proceed if it is unable to properly evaluate them. Valid values: `ENABLED`, `DISABLED`.
	FirewallFailOpen *string `pulumi:"firewallFailOpen"`
	// The AWS account ID of the owner of the VPC that this firewall configuration applies to.
	OwnerId *string `pulumi:"ownerId"`
	// The ID of the VPC that the configuration is for.
	ResourceId *string `pulumi:"resourceId"`
}

type ResolverFirewallConfigState struct {
	// Determines how Route 53 Resolver handles queries during failures, for example when all traffic that is sent to DNS Firewall fails to receive a reply. By default, fail open is disabled, which means the failure mode is closed. This approach favors security over availability. DNS Firewall blocks queries that it is unable to evaluate properly. If you enable this option, the failure mode is open. This approach favors availability over security. DNS Firewall allows queries to proceed if it is unable to properly evaluate them. Valid values: `ENABLED`, `DISABLED`.
	FirewallFailOpen pulumi.StringPtrInput
	// The AWS account ID of the owner of the VPC that this firewall configuration applies to.
	OwnerId pulumi.StringPtrInput
	// The ID of the VPC that the configuration is for.
	ResourceId pulumi.StringPtrInput
}

func (ResolverFirewallConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverFirewallConfigState)(nil)).Elem()
}

type resolverFirewallConfigArgs struct {
	// Determines how Route 53 Resolver handles queries during failures, for example when all traffic that is sent to DNS Firewall fails to receive a reply. By default, fail open is disabled, which means the failure mode is closed. This approach favors security over availability. DNS Firewall blocks queries that it is unable to evaluate properly. If you enable this option, the failure mode is open. This approach favors availability over security. DNS Firewall allows queries to proceed if it is unable to properly evaluate them. Valid values: `ENABLED`, `DISABLED`.
	FirewallFailOpen *string `pulumi:"firewallFailOpen"`
	// The ID of the VPC that the configuration is for.
	ResourceId string `pulumi:"resourceId"`
}

// The set of arguments for constructing a ResolverFirewallConfig resource.
type ResolverFirewallConfigArgs struct {
	// Determines how Route 53 Resolver handles queries during failures, for example when all traffic that is sent to DNS Firewall fails to receive a reply. By default, fail open is disabled, which means the failure mode is closed. This approach favors security over availability. DNS Firewall blocks queries that it is unable to evaluate properly. If you enable this option, the failure mode is open. This approach favors availability over security. DNS Firewall allows queries to proceed if it is unable to properly evaluate them. Valid values: `ENABLED`, `DISABLED`.
	FirewallFailOpen pulumi.StringPtrInput
	// The ID of the VPC that the configuration is for.
	ResourceId pulumi.StringInput
}

func (ResolverFirewallConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverFirewallConfigArgs)(nil)).Elem()
}

type ResolverFirewallConfigInput interface {
	pulumi.Input

	ToResolverFirewallConfigOutput() ResolverFirewallConfigOutput
	ToResolverFirewallConfigOutputWithContext(ctx context.Context) ResolverFirewallConfigOutput
}

func (*ResolverFirewallConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverFirewallConfig)(nil)).Elem()
}

func (i *ResolverFirewallConfig) ToResolverFirewallConfigOutput() ResolverFirewallConfigOutput {
	return i.ToResolverFirewallConfigOutputWithContext(context.Background())
}

func (i *ResolverFirewallConfig) ToResolverFirewallConfigOutputWithContext(ctx context.Context) ResolverFirewallConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallConfigOutput)
}

// ResolverFirewallConfigArrayInput is an input type that accepts ResolverFirewallConfigArray and ResolverFirewallConfigArrayOutput values.
// You can construct a concrete instance of `ResolverFirewallConfigArrayInput` via:
//
//	ResolverFirewallConfigArray{ ResolverFirewallConfigArgs{...} }
type ResolverFirewallConfigArrayInput interface {
	pulumi.Input

	ToResolverFirewallConfigArrayOutput() ResolverFirewallConfigArrayOutput
	ToResolverFirewallConfigArrayOutputWithContext(context.Context) ResolverFirewallConfigArrayOutput
}

type ResolverFirewallConfigArray []ResolverFirewallConfigInput

func (ResolverFirewallConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverFirewallConfig)(nil)).Elem()
}

func (i ResolverFirewallConfigArray) ToResolverFirewallConfigArrayOutput() ResolverFirewallConfigArrayOutput {
	return i.ToResolverFirewallConfigArrayOutputWithContext(context.Background())
}

func (i ResolverFirewallConfigArray) ToResolverFirewallConfigArrayOutputWithContext(ctx context.Context) ResolverFirewallConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallConfigArrayOutput)
}

// ResolverFirewallConfigMapInput is an input type that accepts ResolverFirewallConfigMap and ResolverFirewallConfigMapOutput values.
// You can construct a concrete instance of `ResolverFirewallConfigMapInput` via:
//
//	ResolverFirewallConfigMap{ "key": ResolverFirewallConfigArgs{...} }
type ResolverFirewallConfigMapInput interface {
	pulumi.Input

	ToResolverFirewallConfigMapOutput() ResolverFirewallConfigMapOutput
	ToResolverFirewallConfigMapOutputWithContext(context.Context) ResolverFirewallConfigMapOutput
}

type ResolverFirewallConfigMap map[string]ResolverFirewallConfigInput

func (ResolverFirewallConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverFirewallConfig)(nil)).Elem()
}

func (i ResolverFirewallConfigMap) ToResolverFirewallConfigMapOutput() ResolverFirewallConfigMapOutput {
	return i.ToResolverFirewallConfigMapOutputWithContext(context.Background())
}

func (i ResolverFirewallConfigMap) ToResolverFirewallConfigMapOutputWithContext(ctx context.Context) ResolverFirewallConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverFirewallConfigMapOutput)
}

type ResolverFirewallConfigOutput struct{ *pulumi.OutputState }

func (ResolverFirewallConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverFirewallConfig)(nil)).Elem()
}

func (o ResolverFirewallConfigOutput) ToResolverFirewallConfigOutput() ResolverFirewallConfigOutput {
	return o
}

func (o ResolverFirewallConfigOutput) ToResolverFirewallConfigOutputWithContext(ctx context.Context) ResolverFirewallConfigOutput {
	return o
}

// Determines how Route 53 Resolver handles queries during failures, for example when all traffic that is sent to DNS Firewall fails to receive a reply. By default, fail open is disabled, which means the failure mode is closed. This approach favors security over availability. DNS Firewall blocks queries that it is unable to evaluate properly. If you enable this option, the failure mode is open. This approach favors availability over security. DNS Firewall allows queries to proceed if it is unable to properly evaluate them. Valid values: `ENABLED`, `DISABLED`.
func (o ResolverFirewallConfigOutput) FirewallFailOpen() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverFirewallConfig) pulumi.StringOutput { return v.FirewallFailOpen }).(pulumi.StringOutput)
}

// The AWS account ID of the owner of the VPC that this firewall configuration applies to.
func (o ResolverFirewallConfigOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverFirewallConfig) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The ID of the VPC that the configuration is for.
func (o ResolverFirewallConfigOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverFirewallConfig) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

type ResolverFirewallConfigArrayOutput struct{ *pulumi.OutputState }

func (ResolverFirewallConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverFirewallConfig)(nil)).Elem()
}

func (o ResolverFirewallConfigArrayOutput) ToResolverFirewallConfigArrayOutput() ResolverFirewallConfigArrayOutput {
	return o
}

func (o ResolverFirewallConfigArrayOutput) ToResolverFirewallConfigArrayOutputWithContext(ctx context.Context) ResolverFirewallConfigArrayOutput {
	return o
}

func (o ResolverFirewallConfigArrayOutput) Index(i pulumi.IntInput) ResolverFirewallConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResolverFirewallConfig {
		return vs[0].([]*ResolverFirewallConfig)[vs[1].(int)]
	}).(ResolverFirewallConfigOutput)
}

type ResolverFirewallConfigMapOutput struct{ *pulumi.OutputState }

func (ResolverFirewallConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverFirewallConfig)(nil)).Elem()
}

func (o ResolverFirewallConfigMapOutput) ToResolverFirewallConfigMapOutput() ResolverFirewallConfigMapOutput {
	return o
}

func (o ResolverFirewallConfigMapOutput) ToResolverFirewallConfigMapOutputWithContext(ctx context.Context) ResolverFirewallConfigMapOutput {
	return o
}

func (o ResolverFirewallConfigMapOutput) MapIndex(k pulumi.StringInput) ResolverFirewallConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResolverFirewallConfig {
		return vs[0].(map[string]*ResolverFirewallConfig)[vs[1].(string)]
	}).(ResolverFirewallConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverFirewallConfigInput)(nil)).Elem(), &ResolverFirewallConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverFirewallConfigArrayInput)(nil)).Elem(), ResolverFirewallConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverFirewallConfigMapInput)(nil)).Elem(), ResolverFirewallConfigMap{})
	pulumi.RegisterOutputType(ResolverFirewallConfigOutput{})
	pulumi.RegisterOutputType(ResolverFirewallConfigArrayOutput{})
	pulumi.RegisterOutputType(ResolverFirewallConfigMapOutput{})
}
