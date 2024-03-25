// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attaches a traffic source to an Auto Scaling group.
//
// > **NOTE on Auto Scaling Groups, Attachments and Traffic Source Attachments:** Pulumi provides standalone Attachment (for attaching Classic Load Balancers and Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target groups) and Traffic Source Attachment (for attaching Load Balancers and VPC Lattice target groups) resources and an Auto Scaling Group resource with `loadBalancers`, `targetGroupArns` and `trafficSource` attributes. Do not use the same traffic source in more than one of these resources. Doing so will cause a conflict of attachments. A `lifecycle` configuration block can be used to suppress differences if necessary.
//
// ## Example Usage
//
// ### Basic Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := autoscaling.NewTrafficSourceAttachment(ctx, "example", &autoscaling.TrafficSourceAttachmentArgs{
//				AutoscalingGroupName: pulumi.Any(exampleAwsAutoscalingGroup.Id),
//				TrafficSource: &autoscaling.TrafficSourceAttachmentTrafficSourceArgs{
//					Identifier: pulumi.Any(exampleAwsLbTargetGroup.Arn),
//					Type:       pulumi.String("elbv2"),
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
type TrafficSourceAttachment struct {
	pulumi.CustomResourceState

	// The name of the Auto Scaling group.
	AutoscalingGroupName pulumi.StringOutput `pulumi:"autoscalingGroupName"`
	// The unique identifiers of a traffic sources.
	TrafficSource TrafficSourceAttachmentTrafficSourcePtrOutput `pulumi:"trafficSource"`
}

// NewTrafficSourceAttachment registers a new resource with the given unique name, arguments, and options.
func NewTrafficSourceAttachment(ctx *pulumi.Context,
	name string, args *TrafficSourceAttachmentArgs, opts ...pulumi.ResourceOption) (*TrafficSourceAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoscalingGroupName == nil {
		return nil, errors.New("invalid value for required argument 'AutoscalingGroupName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TrafficSourceAttachment
	err := ctx.RegisterResource("aws:autoscaling/trafficSourceAttachment:TrafficSourceAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficSourceAttachment gets an existing TrafficSourceAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficSourceAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficSourceAttachmentState, opts ...pulumi.ResourceOption) (*TrafficSourceAttachment, error) {
	var resource TrafficSourceAttachment
	err := ctx.ReadResource("aws:autoscaling/trafficSourceAttachment:TrafficSourceAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficSourceAttachment resources.
type trafficSourceAttachmentState struct {
	// The name of the Auto Scaling group.
	AutoscalingGroupName *string `pulumi:"autoscalingGroupName"`
	// The unique identifiers of a traffic sources.
	TrafficSource *TrafficSourceAttachmentTrafficSource `pulumi:"trafficSource"`
}

type TrafficSourceAttachmentState struct {
	// The name of the Auto Scaling group.
	AutoscalingGroupName pulumi.StringPtrInput
	// The unique identifiers of a traffic sources.
	TrafficSource TrafficSourceAttachmentTrafficSourcePtrInput
}

func (TrafficSourceAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficSourceAttachmentState)(nil)).Elem()
}

type trafficSourceAttachmentArgs struct {
	// The name of the Auto Scaling group.
	AutoscalingGroupName string `pulumi:"autoscalingGroupName"`
	// The unique identifiers of a traffic sources.
	TrafficSource *TrafficSourceAttachmentTrafficSource `pulumi:"trafficSource"`
}

// The set of arguments for constructing a TrafficSourceAttachment resource.
type TrafficSourceAttachmentArgs struct {
	// The name of the Auto Scaling group.
	AutoscalingGroupName pulumi.StringInput
	// The unique identifiers of a traffic sources.
	TrafficSource TrafficSourceAttachmentTrafficSourcePtrInput
}

func (TrafficSourceAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficSourceAttachmentArgs)(nil)).Elem()
}

type TrafficSourceAttachmentInput interface {
	pulumi.Input

	ToTrafficSourceAttachmentOutput() TrafficSourceAttachmentOutput
	ToTrafficSourceAttachmentOutputWithContext(ctx context.Context) TrafficSourceAttachmentOutput
}

func (*TrafficSourceAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficSourceAttachment)(nil)).Elem()
}

func (i *TrafficSourceAttachment) ToTrafficSourceAttachmentOutput() TrafficSourceAttachmentOutput {
	return i.ToTrafficSourceAttachmentOutputWithContext(context.Background())
}

func (i *TrafficSourceAttachment) ToTrafficSourceAttachmentOutputWithContext(ctx context.Context) TrafficSourceAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficSourceAttachmentOutput)
}

// TrafficSourceAttachmentArrayInput is an input type that accepts TrafficSourceAttachmentArray and TrafficSourceAttachmentArrayOutput values.
// You can construct a concrete instance of `TrafficSourceAttachmentArrayInput` via:
//
//	TrafficSourceAttachmentArray{ TrafficSourceAttachmentArgs{...} }
type TrafficSourceAttachmentArrayInput interface {
	pulumi.Input

	ToTrafficSourceAttachmentArrayOutput() TrafficSourceAttachmentArrayOutput
	ToTrafficSourceAttachmentArrayOutputWithContext(context.Context) TrafficSourceAttachmentArrayOutput
}

type TrafficSourceAttachmentArray []TrafficSourceAttachmentInput

func (TrafficSourceAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficSourceAttachment)(nil)).Elem()
}

func (i TrafficSourceAttachmentArray) ToTrafficSourceAttachmentArrayOutput() TrafficSourceAttachmentArrayOutput {
	return i.ToTrafficSourceAttachmentArrayOutputWithContext(context.Background())
}

func (i TrafficSourceAttachmentArray) ToTrafficSourceAttachmentArrayOutputWithContext(ctx context.Context) TrafficSourceAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficSourceAttachmentArrayOutput)
}

// TrafficSourceAttachmentMapInput is an input type that accepts TrafficSourceAttachmentMap and TrafficSourceAttachmentMapOutput values.
// You can construct a concrete instance of `TrafficSourceAttachmentMapInput` via:
//
//	TrafficSourceAttachmentMap{ "key": TrafficSourceAttachmentArgs{...} }
type TrafficSourceAttachmentMapInput interface {
	pulumi.Input

	ToTrafficSourceAttachmentMapOutput() TrafficSourceAttachmentMapOutput
	ToTrafficSourceAttachmentMapOutputWithContext(context.Context) TrafficSourceAttachmentMapOutput
}

type TrafficSourceAttachmentMap map[string]TrafficSourceAttachmentInput

func (TrafficSourceAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficSourceAttachment)(nil)).Elem()
}

func (i TrafficSourceAttachmentMap) ToTrafficSourceAttachmentMapOutput() TrafficSourceAttachmentMapOutput {
	return i.ToTrafficSourceAttachmentMapOutputWithContext(context.Background())
}

func (i TrafficSourceAttachmentMap) ToTrafficSourceAttachmentMapOutputWithContext(ctx context.Context) TrafficSourceAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficSourceAttachmentMapOutput)
}

type TrafficSourceAttachmentOutput struct{ *pulumi.OutputState }

func (TrafficSourceAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficSourceAttachment)(nil)).Elem()
}

func (o TrafficSourceAttachmentOutput) ToTrafficSourceAttachmentOutput() TrafficSourceAttachmentOutput {
	return o
}

func (o TrafficSourceAttachmentOutput) ToTrafficSourceAttachmentOutputWithContext(ctx context.Context) TrafficSourceAttachmentOutput {
	return o
}

// The name of the Auto Scaling group.
func (o TrafficSourceAttachmentOutput) AutoscalingGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficSourceAttachment) pulumi.StringOutput { return v.AutoscalingGroupName }).(pulumi.StringOutput)
}

// The unique identifiers of a traffic sources.
func (o TrafficSourceAttachmentOutput) TrafficSource() TrafficSourceAttachmentTrafficSourcePtrOutput {
	return o.ApplyT(func(v *TrafficSourceAttachment) TrafficSourceAttachmentTrafficSourcePtrOutput { return v.TrafficSource }).(TrafficSourceAttachmentTrafficSourcePtrOutput)
}

type TrafficSourceAttachmentArrayOutput struct{ *pulumi.OutputState }

func (TrafficSourceAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficSourceAttachment)(nil)).Elem()
}

func (o TrafficSourceAttachmentArrayOutput) ToTrafficSourceAttachmentArrayOutput() TrafficSourceAttachmentArrayOutput {
	return o
}

func (o TrafficSourceAttachmentArrayOutput) ToTrafficSourceAttachmentArrayOutputWithContext(ctx context.Context) TrafficSourceAttachmentArrayOutput {
	return o
}

func (o TrafficSourceAttachmentArrayOutput) Index(i pulumi.IntInput) TrafficSourceAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficSourceAttachment {
		return vs[0].([]*TrafficSourceAttachment)[vs[1].(int)]
	}).(TrafficSourceAttachmentOutput)
}

type TrafficSourceAttachmentMapOutput struct{ *pulumi.OutputState }

func (TrafficSourceAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficSourceAttachment)(nil)).Elem()
}

func (o TrafficSourceAttachmentMapOutput) ToTrafficSourceAttachmentMapOutput() TrafficSourceAttachmentMapOutput {
	return o
}

func (o TrafficSourceAttachmentMapOutput) ToTrafficSourceAttachmentMapOutputWithContext(ctx context.Context) TrafficSourceAttachmentMapOutput {
	return o
}

func (o TrafficSourceAttachmentMapOutput) MapIndex(k pulumi.StringInput) TrafficSourceAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficSourceAttachment {
		return vs[0].(map[string]*TrafficSourceAttachment)[vs[1].(string)]
	}).(TrafficSourceAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficSourceAttachmentInput)(nil)).Elem(), &TrafficSourceAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficSourceAttachmentArrayInput)(nil)).Elem(), TrafficSourceAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficSourceAttachmentMapInput)(nil)).Elem(), TrafficSourceAttachmentMap{})
	pulumi.RegisterOutputType(TrafficSourceAttachmentOutput{})
	pulumi.RegisterOutputType(TrafficSourceAttachmentArrayOutput{})
	pulumi.RegisterOutputType(TrafficSourceAttachmentMapOutput{})
}
