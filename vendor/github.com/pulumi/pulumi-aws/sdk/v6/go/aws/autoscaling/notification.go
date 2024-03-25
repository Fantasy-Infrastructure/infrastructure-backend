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

// Provides an AutoScaling Group with Notification support, via SNS Topics. Each of
// the `notifications` map to a [Notification Configuration](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeNotificationConfigurations.html) inside Amazon Web
// Services, and are applied to each AutoScaling Group you supply.
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
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := sns.NewTopic(ctx, "example", &sns.TopicArgs{
//				Name: pulumi.String("example-topic"),
//			})
//			if err != nil {
//				return err
//			}
//			bar, err := autoscaling.NewGroup(ctx, "bar", &autoscaling.GroupArgs{
//				Name: pulumi.String("foobar1-test"),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := autoscaling.NewGroup(ctx, "foo", &autoscaling.GroupArgs{
//				Name: pulumi.String("barfoo-test"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = autoscaling.NewNotification(ctx, "example_notifications", &autoscaling.NotificationArgs{
//				GroupNames: pulumi.StringArray{
//					bar.Name,
//					foo.Name,
//				},
//				Notifications: pulumi.StringArray{
//					pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH"),
//					pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE"),
//					pulumi.String("autoscaling:EC2_INSTANCE_LAUNCH_ERROR"),
//					pulumi.String("autoscaling:EC2_INSTANCE_TERMINATE_ERROR"),
//				},
//				TopicArn: example.Arn,
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
type Notification struct {
	pulumi.CustomResourceState

	// List of AutoScaling Group Names
	GroupNames pulumi.StringArrayOutput `pulumi:"groupNames"`
	// List of Notification Types that trigger
	// notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
	Notifications pulumi.StringArrayOutput `pulumi:"notifications"`
	// Topic ARN for notifications to be sent through
	TopicArn pulumi.StringOutput `pulumi:"topicArn"`
}

// NewNotification registers a new resource with the given unique name, arguments, and options.
func NewNotification(ctx *pulumi.Context,
	name string, args *NotificationArgs, opts ...pulumi.ResourceOption) (*Notification, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupNames == nil {
		return nil, errors.New("invalid value for required argument 'GroupNames'")
	}
	if args.Notifications == nil {
		return nil, errors.New("invalid value for required argument 'Notifications'")
	}
	if args.TopicArn == nil {
		return nil, errors.New("invalid value for required argument 'TopicArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Notification
	err := ctx.RegisterResource("aws:autoscaling/notification:Notification", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotification gets an existing Notification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotification(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationState, opts ...pulumi.ResourceOption) (*Notification, error) {
	var resource Notification
	err := ctx.ReadResource("aws:autoscaling/notification:Notification", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Notification resources.
type notificationState struct {
	// List of AutoScaling Group Names
	GroupNames []string `pulumi:"groupNames"`
	// List of Notification Types that trigger
	// notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
	Notifications []string `pulumi:"notifications"`
	// Topic ARN for notifications to be sent through
	TopicArn *string `pulumi:"topicArn"`
}

type NotificationState struct {
	// List of AutoScaling Group Names
	GroupNames pulumi.StringArrayInput
	// List of Notification Types that trigger
	// notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
	Notifications pulumi.StringArrayInput
	// Topic ARN for notifications to be sent through
	TopicArn pulumi.StringPtrInput
}

func (NotificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationState)(nil)).Elem()
}

type notificationArgs struct {
	// List of AutoScaling Group Names
	GroupNames []string `pulumi:"groupNames"`
	// List of Notification Types that trigger
	// notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
	Notifications []string `pulumi:"notifications"`
	// Topic ARN for notifications to be sent through
	TopicArn string `pulumi:"topicArn"`
}

// The set of arguments for constructing a Notification resource.
type NotificationArgs struct {
	// List of AutoScaling Group Names
	GroupNames pulumi.StringArrayInput
	// List of Notification Types that trigger
	// notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
	Notifications pulumi.StringArrayInput
	// Topic ARN for notifications to be sent through
	TopicArn pulumi.StringInput
}

func (NotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationArgs)(nil)).Elem()
}

type NotificationInput interface {
	pulumi.Input

	ToNotificationOutput() NotificationOutput
	ToNotificationOutputWithContext(ctx context.Context) NotificationOutput
}

func (*Notification) ElementType() reflect.Type {
	return reflect.TypeOf((**Notification)(nil)).Elem()
}

func (i *Notification) ToNotificationOutput() NotificationOutput {
	return i.ToNotificationOutputWithContext(context.Background())
}

func (i *Notification) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationOutput)
}

// NotificationArrayInput is an input type that accepts NotificationArray and NotificationArrayOutput values.
// You can construct a concrete instance of `NotificationArrayInput` via:
//
//	NotificationArray{ NotificationArgs{...} }
type NotificationArrayInput interface {
	pulumi.Input

	ToNotificationArrayOutput() NotificationArrayOutput
	ToNotificationArrayOutputWithContext(context.Context) NotificationArrayOutput
}

type NotificationArray []NotificationInput

func (NotificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Notification)(nil)).Elem()
}

func (i NotificationArray) ToNotificationArrayOutput() NotificationArrayOutput {
	return i.ToNotificationArrayOutputWithContext(context.Background())
}

func (i NotificationArray) ToNotificationArrayOutputWithContext(ctx context.Context) NotificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationArrayOutput)
}

// NotificationMapInput is an input type that accepts NotificationMap and NotificationMapOutput values.
// You can construct a concrete instance of `NotificationMapInput` via:
//
//	NotificationMap{ "key": NotificationArgs{...} }
type NotificationMapInput interface {
	pulumi.Input

	ToNotificationMapOutput() NotificationMapOutput
	ToNotificationMapOutputWithContext(context.Context) NotificationMapOutput
}

type NotificationMap map[string]NotificationInput

func (NotificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Notification)(nil)).Elem()
}

func (i NotificationMap) ToNotificationMapOutput() NotificationMapOutput {
	return i.ToNotificationMapOutputWithContext(context.Background())
}

func (i NotificationMap) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationMapOutput)
}

type NotificationOutput struct{ *pulumi.OutputState }

func (NotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Notification)(nil)).Elem()
}

func (o NotificationOutput) ToNotificationOutput() NotificationOutput {
	return o
}

func (o NotificationOutput) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return o
}

// List of AutoScaling Group Names
func (o NotificationOutput) GroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringArrayOutput { return v.GroupNames }).(pulumi.StringArrayOutput)
}

// List of Notification Types that trigger
// notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
func (o NotificationOutput) Notifications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringArrayOutput { return v.Notifications }).(pulumi.StringArrayOutput)
}

// Topic ARN for notifications to be sent through
func (o NotificationOutput) TopicArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.TopicArn }).(pulumi.StringOutput)
}

type NotificationArrayOutput struct{ *pulumi.OutputState }

func (NotificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Notification)(nil)).Elem()
}

func (o NotificationArrayOutput) ToNotificationArrayOutput() NotificationArrayOutput {
	return o
}

func (o NotificationArrayOutput) ToNotificationArrayOutputWithContext(ctx context.Context) NotificationArrayOutput {
	return o
}

func (o NotificationArrayOutput) Index(i pulumi.IntInput) NotificationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Notification {
		return vs[0].([]*Notification)[vs[1].(int)]
	}).(NotificationOutput)
}

type NotificationMapOutput struct{ *pulumi.OutputState }

func (NotificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Notification)(nil)).Elem()
}

func (o NotificationMapOutput) ToNotificationMapOutput() NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) MapIndex(k pulumi.StringInput) NotificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Notification {
		return vs[0].(map[string]*Notification)[vs[1].(string)]
	}).(NotificationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationInput)(nil)).Elem(), &Notification{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationArrayInput)(nil)).Elem(), NotificationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationMapInput)(nil)).Elem(), NotificationMap{})
	pulumi.RegisterOutputType(NotificationOutput{})
	pulumi.RegisterOutputType(NotificationArrayOutput{})
	pulumi.RegisterOutputType(NotificationMapOutput{})
}
