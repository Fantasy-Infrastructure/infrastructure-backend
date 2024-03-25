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

// Authorizes a VPC in a different account to be associated with a local Route53 Hosted Zone.
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
//				CidrBlock:          pulumi.String("10.6.0.0/16"),
//				EnableDnsHostnames: pulumi.Bool(true),
//				EnableDnsSupport:   pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			exampleZone, err := route53.NewZone(ctx, "example", &route53.ZoneArgs{
//				Name: pulumi.String("example.com"),
//				Vpcs: route53.ZoneVpcArray{
//					&route53.ZoneVpcArgs{
//						VpcId: example.ID(),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			alternate, err := ec2.NewVpc(ctx, "alternate", &ec2.VpcArgs{
//				CidrBlock:          pulumi.String("10.7.0.0/16"),
//				EnableDnsHostnames: pulumi.Bool(true),
//				EnableDnsSupport:   pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			exampleVpcAssociationAuthorization, err := route53.NewVpcAssociationAuthorization(ctx, "example", &route53.VpcAssociationAuthorizationArgs{
//				VpcId:  alternate.ID(),
//				ZoneId: exampleZone.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = route53.NewZoneAssociation(ctx, "example", &route53.ZoneAssociationArgs{
//				VpcId:  exampleVpcAssociationAuthorization.VpcId,
//				ZoneId: exampleVpcAssociationAuthorization.ZoneId,
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
// Using `pulumi import`, import Route 53 VPC Association Authorizations using the Hosted Zone ID and VPC ID, separated by a colon (`:`). For example:
//
// ```sh
// $ pulumi import aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization example Z123456ABCDEFG:vpc-12345678
// ```
type VpcAssociationAuthorization struct {
	pulumi.CustomResourceState

	// The VPC to authorize for association with the private hosted zone.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringOutput `pulumi:"vpcRegion"`
	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewVpcAssociationAuthorization registers a new resource with the given unique name, arguments, and options.
func NewVpcAssociationAuthorization(ctx *pulumi.Context,
	name string, args *VpcAssociationAuthorizationArgs, opts ...pulumi.ResourceOption) (*VpcAssociationAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcAssociationAuthorization
	err := ctx.RegisterResource("aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcAssociationAuthorization gets an existing VpcAssociationAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcAssociationAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcAssociationAuthorizationState, opts ...pulumi.ResourceOption) (*VpcAssociationAuthorization, error) {
	var resource VpcAssociationAuthorization
	err := ctx.ReadResource("aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcAssociationAuthorization resources.
type vpcAssociationAuthorizationState struct {
	// The VPC to authorize for association with the private hosted zone.
	VpcId *string `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion *string `pulumi:"vpcRegion"`
	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneId *string `pulumi:"zoneId"`
}

type VpcAssociationAuthorizationState struct {
	// The VPC to authorize for association with the private hosted zone.
	VpcId pulumi.StringPtrInput
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringPtrInput
	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneId pulumi.StringPtrInput
}

func (VpcAssociationAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcAssociationAuthorizationState)(nil)).Elem()
}

type vpcAssociationAuthorizationArgs struct {
	// The VPC to authorize for association with the private hosted zone.
	VpcId string `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion *string `pulumi:"vpcRegion"`
	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a VpcAssociationAuthorization resource.
type VpcAssociationAuthorizationArgs struct {
	// The VPC to authorize for association with the private hosted zone.
	VpcId pulumi.StringInput
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringPtrInput
	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneId pulumi.StringInput
}

func (VpcAssociationAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcAssociationAuthorizationArgs)(nil)).Elem()
}

type VpcAssociationAuthorizationInput interface {
	pulumi.Input

	ToVpcAssociationAuthorizationOutput() VpcAssociationAuthorizationOutput
	ToVpcAssociationAuthorizationOutputWithContext(ctx context.Context) VpcAssociationAuthorizationOutput
}

func (*VpcAssociationAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcAssociationAuthorization)(nil)).Elem()
}

func (i *VpcAssociationAuthorization) ToVpcAssociationAuthorizationOutput() VpcAssociationAuthorizationOutput {
	return i.ToVpcAssociationAuthorizationOutputWithContext(context.Background())
}

func (i *VpcAssociationAuthorization) ToVpcAssociationAuthorizationOutputWithContext(ctx context.Context) VpcAssociationAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcAssociationAuthorizationOutput)
}

// VpcAssociationAuthorizationArrayInput is an input type that accepts VpcAssociationAuthorizationArray and VpcAssociationAuthorizationArrayOutput values.
// You can construct a concrete instance of `VpcAssociationAuthorizationArrayInput` via:
//
//	VpcAssociationAuthorizationArray{ VpcAssociationAuthorizationArgs{...} }
type VpcAssociationAuthorizationArrayInput interface {
	pulumi.Input

	ToVpcAssociationAuthorizationArrayOutput() VpcAssociationAuthorizationArrayOutput
	ToVpcAssociationAuthorizationArrayOutputWithContext(context.Context) VpcAssociationAuthorizationArrayOutput
}

type VpcAssociationAuthorizationArray []VpcAssociationAuthorizationInput

func (VpcAssociationAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcAssociationAuthorization)(nil)).Elem()
}

func (i VpcAssociationAuthorizationArray) ToVpcAssociationAuthorizationArrayOutput() VpcAssociationAuthorizationArrayOutput {
	return i.ToVpcAssociationAuthorizationArrayOutputWithContext(context.Background())
}

func (i VpcAssociationAuthorizationArray) ToVpcAssociationAuthorizationArrayOutputWithContext(ctx context.Context) VpcAssociationAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcAssociationAuthorizationArrayOutput)
}

// VpcAssociationAuthorizationMapInput is an input type that accepts VpcAssociationAuthorizationMap and VpcAssociationAuthorizationMapOutput values.
// You can construct a concrete instance of `VpcAssociationAuthorizationMapInput` via:
//
//	VpcAssociationAuthorizationMap{ "key": VpcAssociationAuthorizationArgs{...} }
type VpcAssociationAuthorizationMapInput interface {
	pulumi.Input

	ToVpcAssociationAuthorizationMapOutput() VpcAssociationAuthorizationMapOutput
	ToVpcAssociationAuthorizationMapOutputWithContext(context.Context) VpcAssociationAuthorizationMapOutput
}

type VpcAssociationAuthorizationMap map[string]VpcAssociationAuthorizationInput

func (VpcAssociationAuthorizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcAssociationAuthorization)(nil)).Elem()
}

func (i VpcAssociationAuthorizationMap) ToVpcAssociationAuthorizationMapOutput() VpcAssociationAuthorizationMapOutput {
	return i.ToVpcAssociationAuthorizationMapOutputWithContext(context.Background())
}

func (i VpcAssociationAuthorizationMap) ToVpcAssociationAuthorizationMapOutputWithContext(ctx context.Context) VpcAssociationAuthorizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcAssociationAuthorizationMapOutput)
}

type VpcAssociationAuthorizationOutput struct{ *pulumi.OutputState }

func (VpcAssociationAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcAssociationAuthorization)(nil)).Elem()
}

func (o VpcAssociationAuthorizationOutput) ToVpcAssociationAuthorizationOutput() VpcAssociationAuthorizationOutput {
	return o
}

func (o VpcAssociationAuthorizationOutput) ToVpcAssociationAuthorizationOutputWithContext(ctx context.Context) VpcAssociationAuthorizationOutput {
	return o
}

// The VPC to authorize for association with the private hosted zone.
func (o VpcAssociationAuthorizationOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcAssociationAuthorization) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The VPC's region. Defaults to the region of the AWS provider.
func (o VpcAssociationAuthorizationOutput) VpcRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcAssociationAuthorization) pulumi.StringOutput { return v.VpcRegion }).(pulumi.StringOutput)
}

// The ID of the private hosted zone that you want to authorize associating a VPC with.
func (o VpcAssociationAuthorizationOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcAssociationAuthorization) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type VpcAssociationAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (VpcAssociationAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcAssociationAuthorization)(nil)).Elem()
}

func (o VpcAssociationAuthorizationArrayOutput) ToVpcAssociationAuthorizationArrayOutput() VpcAssociationAuthorizationArrayOutput {
	return o
}

func (o VpcAssociationAuthorizationArrayOutput) ToVpcAssociationAuthorizationArrayOutputWithContext(ctx context.Context) VpcAssociationAuthorizationArrayOutput {
	return o
}

func (o VpcAssociationAuthorizationArrayOutput) Index(i pulumi.IntInput) VpcAssociationAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcAssociationAuthorization {
		return vs[0].([]*VpcAssociationAuthorization)[vs[1].(int)]
	}).(VpcAssociationAuthorizationOutput)
}

type VpcAssociationAuthorizationMapOutput struct{ *pulumi.OutputState }

func (VpcAssociationAuthorizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcAssociationAuthorization)(nil)).Elem()
}

func (o VpcAssociationAuthorizationMapOutput) ToVpcAssociationAuthorizationMapOutput() VpcAssociationAuthorizationMapOutput {
	return o
}

func (o VpcAssociationAuthorizationMapOutput) ToVpcAssociationAuthorizationMapOutputWithContext(ctx context.Context) VpcAssociationAuthorizationMapOutput {
	return o
}

func (o VpcAssociationAuthorizationMapOutput) MapIndex(k pulumi.StringInput) VpcAssociationAuthorizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcAssociationAuthorization {
		return vs[0].(map[string]*VpcAssociationAuthorization)[vs[1].(string)]
	}).(VpcAssociationAuthorizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcAssociationAuthorizationInput)(nil)).Elem(), &VpcAssociationAuthorization{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcAssociationAuthorizationArrayInput)(nil)).Elem(), VpcAssociationAuthorizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcAssociationAuthorizationMapInput)(nil)).Elem(), VpcAssociationAuthorizationMap{})
	pulumi.RegisterOutputType(VpcAssociationAuthorizationOutput{})
	pulumi.RegisterOutputType(VpcAssociationAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(VpcAssociationAuthorizationMapOutput{})
}
