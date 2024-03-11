// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage the [default AWS DHCP Options Set](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_DHCP_Options.html#AmazonDNS)
// in the current region.
//
// Each AWS region comes with a default set of DHCP options.
// **This is an advanced resource**, and has special caveats to be aware of when
// using it. Please read this document in its entirety before using this resource.
//
// The `ec2.DefaultVpcDhcpOptions` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead "adopts" it
// into management.
//
// ## Example Usage
//
// Basic usage with tags:
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
//			_, err := ec2.NewDefaultVpcDhcpOptions(ctx, "default", &ec2.DefaultVpcDhcpOptionsArgs{
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("Default DHCP Option Set"),
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
//
// ## Import
//
// Using `pulumi import`, import VPC DHCP Options using the DHCP Options `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2/defaultVpcDhcpOptions:DefaultVpcDhcpOptions default_options dopt-d9070ebb
//
// ```
type DefaultVpcDhcpOptions struct {
	pulumi.CustomResourceState

	// The ARN of the DHCP Options Set.
	Arn               pulumi.StringOutput `pulumi:"arn"`
	DomainName        pulumi.StringOutput `pulumi:"domainName"`
	DomainNameServers pulumi.StringOutput `pulumi:"domainNameServers"`
	// List of NETBIOS name servers.
	NetbiosNameServers pulumi.StringOutput `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType pulumi.StringOutput `pulumi:"netbiosNodeType"`
	NtpServers      pulumi.StringOutput `pulumi:"ntpServers"`
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewDefaultVpcDhcpOptions registers a new resource with the given unique name, arguments, and options.
func NewDefaultVpcDhcpOptions(ctx *pulumi.Context,
	name string, args *DefaultVpcDhcpOptionsArgs, opts ...pulumi.ResourceOption) (*DefaultVpcDhcpOptions, error) {
	if args == nil {
		args = &DefaultVpcDhcpOptionsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DefaultVpcDhcpOptions
	err := ctx.RegisterResource("aws:ec2/defaultVpcDhcpOptions:DefaultVpcDhcpOptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultVpcDhcpOptions gets an existing DefaultVpcDhcpOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultVpcDhcpOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultVpcDhcpOptionsState, opts ...pulumi.ResourceOption) (*DefaultVpcDhcpOptions, error) {
	var resource DefaultVpcDhcpOptions
	err := ctx.ReadResource("aws:ec2/defaultVpcDhcpOptions:DefaultVpcDhcpOptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultVpcDhcpOptions resources.
type defaultVpcDhcpOptionsState struct {
	// The ARN of the DHCP Options Set.
	Arn               *string `pulumi:"arn"`
	DomainName        *string `pulumi:"domainName"`
	DomainNameServers *string `pulumi:"domainNameServers"`
	// List of NETBIOS name servers.
	NetbiosNameServers *string `pulumi:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType *string `pulumi:"netbiosNodeType"`
	NtpServers      *string `pulumi:"ntpServers"`
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type DefaultVpcDhcpOptionsState struct {
	// The ARN of the DHCP Options Set.
	Arn               pulumi.StringPtrInput
	DomainName        pulumi.StringPtrInput
	DomainNameServers pulumi.StringPtrInput
	// List of NETBIOS name servers.
	NetbiosNameServers pulumi.StringPtrInput
	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType pulumi.StringPtrInput
	NtpServers      pulumi.StringPtrInput
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (DefaultVpcDhcpOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultVpcDhcpOptionsState)(nil)).Elem()
}

type defaultVpcDhcpOptionsArgs struct {
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DefaultVpcDhcpOptions resource.
type DefaultVpcDhcpOptionsArgs struct {
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DefaultVpcDhcpOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultVpcDhcpOptionsArgs)(nil)).Elem()
}

type DefaultVpcDhcpOptionsInput interface {
	pulumi.Input

	ToDefaultVpcDhcpOptionsOutput() DefaultVpcDhcpOptionsOutput
	ToDefaultVpcDhcpOptionsOutputWithContext(ctx context.Context) DefaultVpcDhcpOptionsOutput
}

func (*DefaultVpcDhcpOptions) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultVpcDhcpOptions)(nil)).Elem()
}

func (i *DefaultVpcDhcpOptions) ToDefaultVpcDhcpOptionsOutput() DefaultVpcDhcpOptionsOutput {
	return i.ToDefaultVpcDhcpOptionsOutputWithContext(context.Background())
}

func (i *DefaultVpcDhcpOptions) ToDefaultVpcDhcpOptionsOutputWithContext(ctx context.Context) DefaultVpcDhcpOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultVpcDhcpOptionsOutput)
}

// DefaultVpcDhcpOptionsArrayInput is an input type that accepts DefaultVpcDhcpOptionsArray and DefaultVpcDhcpOptionsArrayOutput values.
// You can construct a concrete instance of `DefaultVpcDhcpOptionsArrayInput` via:
//
//	DefaultVpcDhcpOptionsArray{ DefaultVpcDhcpOptionsArgs{...} }
type DefaultVpcDhcpOptionsArrayInput interface {
	pulumi.Input

	ToDefaultVpcDhcpOptionsArrayOutput() DefaultVpcDhcpOptionsArrayOutput
	ToDefaultVpcDhcpOptionsArrayOutputWithContext(context.Context) DefaultVpcDhcpOptionsArrayOutput
}

type DefaultVpcDhcpOptionsArray []DefaultVpcDhcpOptionsInput

func (DefaultVpcDhcpOptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultVpcDhcpOptions)(nil)).Elem()
}

func (i DefaultVpcDhcpOptionsArray) ToDefaultVpcDhcpOptionsArrayOutput() DefaultVpcDhcpOptionsArrayOutput {
	return i.ToDefaultVpcDhcpOptionsArrayOutputWithContext(context.Background())
}

func (i DefaultVpcDhcpOptionsArray) ToDefaultVpcDhcpOptionsArrayOutputWithContext(ctx context.Context) DefaultVpcDhcpOptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultVpcDhcpOptionsArrayOutput)
}

// DefaultVpcDhcpOptionsMapInput is an input type that accepts DefaultVpcDhcpOptionsMap and DefaultVpcDhcpOptionsMapOutput values.
// You can construct a concrete instance of `DefaultVpcDhcpOptionsMapInput` via:
//
//	DefaultVpcDhcpOptionsMap{ "key": DefaultVpcDhcpOptionsArgs{...} }
type DefaultVpcDhcpOptionsMapInput interface {
	pulumi.Input

	ToDefaultVpcDhcpOptionsMapOutput() DefaultVpcDhcpOptionsMapOutput
	ToDefaultVpcDhcpOptionsMapOutputWithContext(context.Context) DefaultVpcDhcpOptionsMapOutput
}

type DefaultVpcDhcpOptionsMap map[string]DefaultVpcDhcpOptionsInput

func (DefaultVpcDhcpOptionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultVpcDhcpOptions)(nil)).Elem()
}

func (i DefaultVpcDhcpOptionsMap) ToDefaultVpcDhcpOptionsMapOutput() DefaultVpcDhcpOptionsMapOutput {
	return i.ToDefaultVpcDhcpOptionsMapOutputWithContext(context.Background())
}

func (i DefaultVpcDhcpOptionsMap) ToDefaultVpcDhcpOptionsMapOutputWithContext(ctx context.Context) DefaultVpcDhcpOptionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultVpcDhcpOptionsMapOutput)
}

type DefaultVpcDhcpOptionsOutput struct{ *pulumi.OutputState }

func (DefaultVpcDhcpOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultVpcDhcpOptions)(nil)).Elem()
}

func (o DefaultVpcDhcpOptionsOutput) ToDefaultVpcDhcpOptionsOutput() DefaultVpcDhcpOptionsOutput {
	return o
}

func (o DefaultVpcDhcpOptionsOutput) ToDefaultVpcDhcpOptionsOutputWithContext(ctx context.Context) DefaultVpcDhcpOptionsOutput {
	return o
}

// The ARN of the DHCP Options Set.
func (o DefaultVpcDhcpOptionsOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultVpcDhcpOptions) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o DefaultVpcDhcpOptionsOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultVpcDhcpOptions) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o DefaultVpcDhcpOptionsOutput) DomainNameServers() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultVpcDhcpOptions) pulumi.StringOutput { return v.DomainNameServers }).(pulumi.StringOutput)
}

// List of NETBIOS name servers.
func (o DefaultVpcDhcpOptionsOutput) NetbiosNameServers() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultVpcDhcpOptions) pulumi.StringOutput { return v.NetbiosNameServers }).(pulumi.StringOutput)
}

// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
func (o DefaultVpcDhcpOptionsOutput) NetbiosNodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultVpcDhcpOptions) pulumi.StringOutput { return v.NetbiosNodeType }).(pulumi.StringOutput)
}

func (o DefaultVpcDhcpOptionsOutput) NtpServers() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultVpcDhcpOptions) pulumi.StringOutput { return v.NtpServers }).(pulumi.StringOutput)
}

// The ID of the AWS account that owns the DHCP options set.
func (o DefaultVpcDhcpOptionsOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultVpcDhcpOptions) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource.
func (o DefaultVpcDhcpOptionsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DefaultVpcDhcpOptions) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Deprecated: Please use `tags` instead.
func (o DefaultVpcDhcpOptionsOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DefaultVpcDhcpOptions) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type DefaultVpcDhcpOptionsArrayOutput struct{ *pulumi.OutputState }

func (DefaultVpcDhcpOptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultVpcDhcpOptions)(nil)).Elem()
}

func (o DefaultVpcDhcpOptionsArrayOutput) ToDefaultVpcDhcpOptionsArrayOutput() DefaultVpcDhcpOptionsArrayOutput {
	return o
}

func (o DefaultVpcDhcpOptionsArrayOutput) ToDefaultVpcDhcpOptionsArrayOutputWithContext(ctx context.Context) DefaultVpcDhcpOptionsArrayOutput {
	return o
}

func (o DefaultVpcDhcpOptionsArrayOutput) Index(i pulumi.IntInput) DefaultVpcDhcpOptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultVpcDhcpOptions {
		return vs[0].([]*DefaultVpcDhcpOptions)[vs[1].(int)]
	}).(DefaultVpcDhcpOptionsOutput)
}

type DefaultVpcDhcpOptionsMapOutput struct{ *pulumi.OutputState }

func (DefaultVpcDhcpOptionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultVpcDhcpOptions)(nil)).Elem()
}

func (o DefaultVpcDhcpOptionsMapOutput) ToDefaultVpcDhcpOptionsMapOutput() DefaultVpcDhcpOptionsMapOutput {
	return o
}

func (o DefaultVpcDhcpOptionsMapOutput) ToDefaultVpcDhcpOptionsMapOutputWithContext(ctx context.Context) DefaultVpcDhcpOptionsMapOutput {
	return o
}

func (o DefaultVpcDhcpOptionsMapOutput) MapIndex(k pulumi.StringInput) DefaultVpcDhcpOptionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultVpcDhcpOptions {
		return vs[0].(map[string]*DefaultVpcDhcpOptions)[vs[1].(string)]
	}).(DefaultVpcDhcpOptionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultVpcDhcpOptionsInput)(nil)).Elem(), &DefaultVpcDhcpOptions{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultVpcDhcpOptionsArrayInput)(nil)).Elem(), DefaultVpcDhcpOptionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultVpcDhcpOptionsMapInput)(nil)).Elem(), DefaultVpcDhcpOptionsMap{})
	pulumi.RegisterOutputType(DefaultVpcDhcpOptionsOutput{})
	pulumi.RegisterOutputType(DefaultVpcDhcpOptionsArrayOutput{})
	pulumi.RegisterOutputType(DefaultVpcDhcpOptionsMapOutput{})
}
