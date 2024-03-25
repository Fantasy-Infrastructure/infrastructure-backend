// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an ECS cluster capacity provider. More information can be found on the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html).
//
// > **NOTE:** Associating an ECS Capacity Provider to an Auto Scaling Group will automatically add the `AmazonECSManaged` tag to the Auto Scaling Group. This tag should be included in the `autoscaling.Group` resource configuration to prevent the provider from removing it in subsequent executions as well as ensuring the `AmazonECSManaged` tag is propagated to all EC2 Instances in the Auto Scaling Group if `minSize` is above 0 on creation. Any EC2 Instances in the Auto Scaling Group without this tag must be manually be updated, otherwise they may cause unexpected scaling behavior and metrics.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := autoscaling.NewGroup(ctx, "test", &autoscaling.GroupArgs{
//				Tags: autoscaling.GroupTagArray{
//					&autoscaling.GroupTagArgs{
//						Key:               pulumi.String("AmazonECSManaged"),
//						Value:             pulumi.String("true"),
//						PropagateAtLaunch: pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewCapacityProvider(ctx, "test", &ecs.CapacityProviderArgs{
//				Name: pulumi.String("test"),
//				AutoScalingGroupProvider: &ecs.CapacityProviderAutoScalingGroupProviderArgs{
//					AutoScalingGroupArn:          test.Arn,
//					ManagedTerminationProtection: pulumi.String("ENABLED"),
//					ManagedScaling: &ecs.CapacityProviderAutoScalingGroupProviderManagedScalingArgs{
//						MaximumScalingStepSize: pulumi.Int(1000),
//						MinimumScalingStepSize: pulumi.Int(1),
//						Status:                 pulumi.String("ENABLED"),
//						TargetCapacity:         pulumi.Int(10),
//					},
//				},
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
// Using `pulumi import`, import ECS Capacity Providers using the `name`. For example:
//
// ```sh
// $ pulumi import aws:ecs/capacityProvider:CapacityProvider example example
// ```
type CapacityProvider struct {
	pulumi.CustomResourceState

	// ARN that identifies the capacity provider.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block for the provider for the ECS auto scaling group. Detailed below.
	AutoScalingGroupProvider CapacityProviderAutoScalingGroupProviderOutput `pulumi:"autoScalingGroupProvider"`
	// Name of the capacity provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewCapacityProvider registers a new resource with the given unique name, arguments, and options.
func NewCapacityProvider(ctx *pulumi.Context,
	name string, args *CapacityProviderArgs, opts ...pulumi.ResourceOption) (*CapacityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoScalingGroupProvider == nil {
		return nil, errors.New("invalid value for required argument 'AutoScalingGroupProvider'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CapacityProvider
	err := ctx.RegisterResource("aws:ecs/capacityProvider:CapacityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacityProvider gets an existing CapacityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityProviderState, opts ...pulumi.ResourceOption) (*CapacityProvider, error) {
	var resource CapacityProvider
	err := ctx.ReadResource("aws:ecs/capacityProvider:CapacityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CapacityProvider resources.
type capacityProviderState struct {
	// ARN that identifies the capacity provider.
	Arn *string `pulumi:"arn"`
	// Configuration block for the provider for the ECS auto scaling group. Detailed below.
	AutoScalingGroupProvider *CapacityProviderAutoScalingGroupProvider `pulumi:"autoScalingGroupProvider"`
	// Name of the capacity provider.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type CapacityProviderState struct {
	// ARN that identifies the capacity provider.
	Arn pulumi.StringPtrInput
	// Configuration block for the provider for the ECS auto scaling group. Detailed below.
	AutoScalingGroupProvider CapacityProviderAutoScalingGroupProviderPtrInput
	// Name of the capacity provider.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (CapacityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityProviderState)(nil)).Elem()
}

type capacityProviderArgs struct {
	// Configuration block for the provider for the ECS auto scaling group. Detailed below.
	AutoScalingGroupProvider CapacityProviderAutoScalingGroupProvider `pulumi:"autoScalingGroupProvider"`
	// Name of the capacity provider.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CapacityProvider resource.
type CapacityProviderArgs struct {
	// Configuration block for the provider for the ECS auto scaling group. Detailed below.
	AutoScalingGroupProvider CapacityProviderAutoScalingGroupProviderInput
	// Name of the capacity provider.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (CapacityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityProviderArgs)(nil)).Elem()
}

type CapacityProviderInput interface {
	pulumi.Input

	ToCapacityProviderOutput() CapacityProviderOutput
	ToCapacityProviderOutputWithContext(ctx context.Context) CapacityProviderOutput
}

func (*CapacityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityProvider)(nil)).Elem()
}

func (i *CapacityProvider) ToCapacityProviderOutput() CapacityProviderOutput {
	return i.ToCapacityProviderOutputWithContext(context.Background())
}

func (i *CapacityProvider) ToCapacityProviderOutputWithContext(ctx context.Context) CapacityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityProviderOutput)
}

// CapacityProviderArrayInput is an input type that accepts CapacityProviderArray and CapacityProviderArrayOutput values.
// You can construct a concrete instance of `CapacityProviderArrayInput` via:
//
//	CapacityProviderArray{ CapacityProviderArgs{...} }
type CapacityProviderArrayInput interface {
	pulumi.Input

	ToCapacityProviderArrayOutput() CapacityProviderArrayOutput
	ToCapacityProviderArrayOutputWithContext(context.Context) CapacityProviderArrayOutput
}

type CapacityProviderArray []CapacityProviderInput

func (CapacityProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CapacityProvider)(nil)).Elem()
}

func (i CapacityProviderArray) ToCapacityProviderArrayOutput() CapacityProviderArrayOutput {
	return i.ToCapacityProviderArrayOutputWithContext(context.Background())
}

func (i CapacityProviderArray) ToCapacityProviderArrayOutputWithContext(ctx context.Context) CapacityProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityProviderArrayOutput)
}

// CapacityProviderMapInput is an input type that accepts CapacityProviderMap and CapacityProviderMapOutput values.
// You can construct a concrete instance of `CapacityProviderMapInput` via:
//
//	CapacityProviderMap{ "key": CapacityProviderArgs{...} }
type CapacityProviderMapInput interface {
	pulumi.Input

	ToCapacityProviderMapOutput() CapacityProviderMapOutput
	ToCapacityProviderMapOutputWithContext(context.Context) CapacityProviderMapOutput
}

type CapacityProviderMap map[string]CapacityProviderInput

func (CapacityProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CapacityProvider)(nil)).Elem()
}

func (i CapacityProviderMap) ToCapacityProviderMapOutput() CapacityProviderMapOutput {
	return i.ToCapacityProviderMapOutputWithContext(context.Background())
}

func (i CapacityProviderMap) ToCapacityProviderMapOutputWithContext(ctx context.Context) CapacityProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityProviderMapOutput)
}

type CapacityProviderOutput struct{ *pulumi.OutputState }

func (CapacityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityProvider)(nil)).Elem()
}

func (o CapacityProviderOutput) ToCapacityProviderOutput() CapacityProviderOutput {
	return o
}

func (o CapacityProviderOutput) ToCapacityProviderOutputWithContext(ctx context.Context) CapacityProviderOutput {
	return o
}

// ARN that identifies the capacity provider.
func (o CapacityProviderOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityProvider) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Configuration block for the provider for the ECS auto scaling group. Detailed below.
func (o CapacityProviderOutput) AutoScalingGroupProvider() CapacityProviderAutoScalingGroupProviderOutput {
	return o.ApplyT(func(v *CapacityProvider) CapacityProviderAutoScalingGroupProviderOutput {
		return v.AutoScalingGroupProvider
	}).(CapacityProviderAutoScalingGroupProviderOutput)
}

// Name of the capacity provider.
func (o CapacityProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o CapacityProviderOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CapacityProvider) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o CapacityProviderOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CapacityProvider) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type CapacityProviderArrayOutput struct{ *pulumi.OutputState }

func (CapacityProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CapacityProvider)(nil)).Elem()
}

func (o CapacityProviderArrayOutput) ToCapacityProviderArrayOutput() CapacityProviderArrayOutput {
	return o
}

func (o CapacityProviderArrayOutput) ToCapacityProviderArrayOutputWithContext(ctx context.Context) CapacityProviderArrayOutput {
	return o
}

func (o CapacityProviderArrayOutput) Index(i pulumi.IntInput) CapacityProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CapacityProvider {
		return vs[0].([]*CapacityProvider)[vs[1].(int)]
	}).(CapacityProviderOutput)
}

type CapacityProviderMapOutput struct{ *pulumi.OutputState }

func (CapacityProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CapacityProvider)(nil)).Elem()
}

func (o CapacityProviderMapOutput) ToCapacityProviderMapOutput() CapacityProviderMapOutput {
	return o
}

func (o CapacityProviderMapOutput) ToCapacityProviderMapOutputWithContext(ctx context.Context) CapacityProviderMapOutput {
	return o
}

func (o CapacityProviderMapOutput) MapIndex(k pulumi.StringInput) CapacityProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CapacityProvider {
		return vs[0].(map[string]*CapacityProvider)[vs[1].(string)]
	}).(CapacityProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CapacityProviderInput)(nil)).Elem(), &CapacityProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*CapacityProviderArrayInput)(nil)).Elem(), CapacityProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CapacityProviderMapInput)(nil)).Elem(), CapacityProviderMap{})
	pulumi.RegisterOutputType(CapacityProviderOutput{})
	pulumi.RegisterOutputType(CapacityProviderArrayOutput{})
	pulumi.RegisterOutputType(CapacityProviderMapOutput{})
}
