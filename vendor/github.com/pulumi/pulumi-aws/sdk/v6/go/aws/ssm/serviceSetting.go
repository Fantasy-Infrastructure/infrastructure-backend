// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This setting defines how a user interacts with or uses a service or a feature of a service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssm.NewServiceSetting(ctx, "testSetting", &ssm.ServiceSettingArgs{
//				SettingId:    pulumi.String("arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled"),
//				SettingValue: pulumi.String("true"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import AWS SSM Service Setting using the `setting_id`. For example:
//
// ```sh
//
//	$ pulumi import aws:ssm/serviceSetting:ServiceSetting example arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled
//
// ```
type ServiceSetting struct {
	pulumi.CustomResourceState

	// ARN of the service setting.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ID of the service setting.
	SettingId pulumi.StringOutput `pulumi:"settingId"`
	// Value of the service setting.
	SettingValue pulumi.StringOutput `pulumi:"settingValue"`
	// Status of the service setting. Value can be `Default`, `Customized` or `PendingUpdate`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewServiceSetting registers a new resource with the given unique name, arguments, and options.
func NewServiceSetting(ctx *pulumi.Context,
	name string, args *ServiceSettingArgs, opts ...pulumi.ResourceOption) (*ServiceSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SettingId == nil {
		return nil, errors.New("invalid value for required argument 'SettingId'")
	}
	if args.SettingValue == nil {
		return nil, errors.New("invalid value for required argument 'SettingValue'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceSetting
	err := ctx.RegisterResource("aws:ssm/serviceSetting:ServiceSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceSetting gets an existing ServiceSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceSettingState, opts ...pulumi.ResourceOption) (*ServiceSetting, error) {
	var resource ServiceSetting
	err := ctx.ReadResource("aws:ssm/serviceSetting:ServiceSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceSetting resources.
type serviceSettingState struct {
	// ARN of the service setting.
	Arn *string `pulumi:"arn"`
	// ID of the service setting.
	SettingId *string `pulumi:"settingId"`
	// Value of the service setting.
	SettingValue *string `pulumi:"settingValue"`
	// Status of the service setting. Value can be `Default`, `Customized` or `PendingUpdate`.
	Status *string `pulumi:"status"`
}

type ServiceSettingState struct {
	// ARN of the service setting.
	Arn pulumi.StringPtrInput
	// ID of the service setting.
	SettingId pulumi.StringPtrInput
	// Value of the service setting.
	SettingValue pulumi.StringPtrInput
	// Status of the service setting. Value can be `Default`, `Customized` or `PendingUpdate`.
	Status pulumi.StringPtrInput
}

func (ServiceSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceSettingState)(nil)).Elem()
}

type serviceSettingArgs struct {
	// ID of the service setting.
	SettingId string `pulumi:"settingId"`
	// Value of the service setting.
	SettingValue string `pulumi:"settingValue"`
}

// The set of arguments for constructing a ServiceSetting resource.
type ServiceSettingArgs struct {
	// ID of the service setting.
	SettingId pulumi.StringInput
	// Value of the service setting.
	SettingValue pulumi.StringInput
}

func (ServiceSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceSettingArgs)(nil)).Elem()
}

type ServiceSettingInput interface {
	pulumi.Input

	ToServiceSettingOutput() ServiceSettingOutput
	ToServiceSettingOutputWithContext(ctx context.Context) ServiceSettingOutput
}

func (*ServiceSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceSetting)(nil)).Elem()
}

func (i *ServiceSetting) ToServiceSettingOutput() ServiceSettingOutput {
	return i.ToServiceSettingOutputWithContext(context.Background())
}

func (i *ServiceSetting) ToServiceSettingOutputWithContext(ctx context.Context) ServiceSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSettingOutput)
}

// ServiceSettingArrayInput is an input type that accepts ServiceSettingArray and ServiceSettingArrayOutput values.
// You can construct a concrete instance of `ServiceSettingArrayInput` via:
//
//	ServiceSettingArray{ ServiceSettingArgs{...} }
type ServiceSettingArrayInput interface {
	pulumi.Input

	ToServiceSettingArrayOutput() ServiceSettingArrayOutput
	ToServiceSettingArrayOutputWithContext(context.Context) ServiceSettingArrayOutput
}

type ServiceSettingArray []ServiceSettingInput

func (ServiceSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceSetting)(nil)).Elem()
}

func (i ServiceSettingArray) ToServiceSettingArrayOutput() ServiceSettingArrayOutput {
	return i.ToServiceSettingArrayOutputWithContext(context.Background())
}

func (i ServiceSettingArray) ToServiceSettingArrayOutputWithContext(ctx context.Context) ServiceSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSettingArrayOutput)
}

// ServiceSettingMapInput is an input type that accepts ServiceSettingMap and ServiceSettingMapOutput values.
// You can construct a concrete instance of `ServiceSettingMapInput` via:
//
//	ServiceSettingMap{ "key": ServiceSettingArgs{...} }
type ServiceSettingMapInput interface {
	pulumi.Input

	ToServiceSettingMapOutput() ServiceSettingMapOutput
	ToServiceSettingMapOutputWithContext(context.Context) ServiceSettingMapOutput
}

type ServiceSettingMap map[string]ServiceSettingInput

func (ServiceSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceSetting)(nil)).Elem()
}

func (i ServiceSettingMap) ToServiceSettingMapOutput() ServiceSettingMapOutput {
	return i.ToServiceSettingMapOutputWithContext(context.Background())
}

func (i ServiceSettingMap) ToServiceSettingMapOutputWithContext(ctx context.Context) ServiceSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceSettingMapOutput)
}

type ServiceSettingOutput struct{ *pulumi.OutputState }

func (ServiceSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceSetting)(nil)).Elem()
}

func (o ServiceSettingOutput) ToServiceSettingOutput() ServiceSettingOutput {
	return o
}

func (o ServiceSettingOutput) ToServiceSettingOutputWithContext(ctx context.Context) ServiceSettingOutput {
	return o
}

// ARN of the service setting.
func (o ServiceSettingOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceSetting) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// ID of the service setting.
func (o ServiceSettingOutput) SettingId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceSetting) pulumi.StringOutput { return v.SettingId }).(pulumi.StringOutput)
}

// Value of the service setting.
func (o ServiceSettingOutput) SettingValue() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceSetting) pulumi.StringOutput { return v.SettingValue }).(pulumi.StringOutput)
}

// Status of the service setting. Value can be `Default`, `Customized` or `PendingUpdate`.
func (o ServiceSettingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceSetting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type ServiceSettingArrayOutput struct{ *pulumi.OutputState }

func (ServiceSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceSetting)(nil)).Elem()
}

func (o ServiceSettingArrayOutput) ToServiceSettingArrayOutput() ServiceSettingArrayOutput {
	return o
}

func (o ServiceSettingArrayOutput) ToServiceSettingArrayOutputWithContext(ctx context.Context) ServiceSettingArrayOutput {
	return o
}

func (o ServiceSettingArrayOutput) Index(i pulumi.IntInput) ServiceSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceSetting {
		return vs[0].([]*ServiceSetting)[vs[1].(int)]
	}).(ServiceSettingOutput)
}

type ServiceSettingMapOutput struct{ *pulumi.OutputState }

func (ServiceSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceSetting)(nil)).Elem()
}

func (o ServiceSettingMapOutput) ToServiceSettingMapOutput() ServiceSettingMapOutput {
	return o
}

func (o ServiceSettingMapOutput) ToServiceSettingMapOutputWithContext(ctx context.Context) ServiceSettingMapOutput {
	return o
}

func (o ServiceSettingMapOutput) MapIndex(k pulumi.StringInput) ServiceSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceSetting {
		return vs[0].(map[string]*ServiceSetting)[vs[1].(string)]
	}).(ServiceSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceSettingInput)(nil)).Elem(), &ServiceSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceSettingArrayInput)(nil)).Elem(), ServiceSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceSettingMapInput)(nil)).Elem(), ServiceSettingMap{})
	pulumi.RegisterOutputType(ServiceSettingOutput{})
	pulumi.RegisterOutputType(ServiceSettingArrayOutput{})
	pulumi.RegisterOutputType(ServiceSettingMapOutput{})
}