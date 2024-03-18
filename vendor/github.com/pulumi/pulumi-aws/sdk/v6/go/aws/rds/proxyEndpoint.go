// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an RDS DB proxy endpoint resource. For additional information, see the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy-endpoints.html).
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// var splat0 []interface{}
// for _, val0 := range testAwsSubnet {
// splat0 = append(splat0, val0.Id)
// }
// _, err := rds.NewProxyEndpoint(ctx, "example", &rds.ProxyEndpointArgs{
// DbProxyName: pulumi.Any(test.Name),
// DbProxyEndpointName: pulumi.String("example"),
// VpcSubnetIds: toPulumiArray(splat0),
// TargetRole: pulumi.String("READ_ONLY"),
// })
// if err != nil {
// return err
// }
// return nil
// })
// }
// func toPulumiArray(arr []) pulumi.Array {
// var pulumiArr pulumi.Array
// for _, v := range arr {
// pulumiArr = append(pulumiArr, pulumi.(v))
// }
// return pulumiArr
// }
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Using `pulumi import`, import DB proxy endpoints using the `DB-PROXY-NAME/DB-PROXY-ENDPOINT-NAME`. For example:
//
// ```sh
// $ pulumi import aws:rds/proxyEndpoint:ProxyEndpoint example example/example
// ```
type ProxyEndpoint struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) for the proxy endpoint.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	DbProxyEndpointName pulumi.StringOutput `pulumi:"dbProxyEndpointName"`
	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	DbProxyName pulumi.StringOutput `pulumi:"dbProxyName"`
	// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Indicates whether this endpoint is the default endpoint for the associated DB proxy.
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
	TargetRole pulumi.StringPtrOutput `pulumi:"targetRole"`
	// The VPC ID of the DB proxy endpoint.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds pulumi.StringArrayOutput `pulumi:"vpcSecurityGroupIds"`
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds pulumi.StringArrayOutput `pulumi:"vpcSubnetIds"`
}

// NewProxyEndpoint registers a new resource with the given unique name, arguments, and options.
func NewProxyEndpoint(ctx *pulumi.Context,
	name string, args *ProxyEndpointArgs, opts ...pulumi.ResourceOption) (*ProxyEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbProxyEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'DbProxyEndpointName'")
	}
	if args.DbProxyName == nil {
		return nil, errors.New("invalid value for required argument 'DbProxyName'")
	}
	if args.VpcSubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'VpcSubnetIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProxyEndpoint
	err := ctx.RegisterResource("aws:rds/proxyEndpoint:ProxyEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProxyEndpoint gets an existing ProxyEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProxyEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProxyEndpointState, opts ...pulumi.ResourceOption) (*ProxyEndpoint, error) {
	var resource ProxyEndpoint
	err := ctx.ReadResource("aws:rds/proxyEndpoint:ProxyEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProxyEndpoint resources.
type proxyEndpointState struct {
	// The Amazon Resource Name (ARN) for the proxy endpoint.
	Arn *string `pulumi:"arn"`
	// The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	DbProxyEndpointName *string `pulumi:"dbProxyEndpointName"`
	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	DbProxyName *string `pulumi:"dbProxyName"`
	// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint *string `pulumi:"endpoint"`
	// Indicates whether this endpoint is the default endpoint for the associated DB proxy.
	IsDefault *bool `pulumi:"isDefault"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
	TargetRole *string `pulumi:"targetRole"`
	// The VPC ID of the DB proxy endpoint.
	VpcId *string `pulumi:"vpcId"`
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds []string `pulumi:"vpcSubnetIds"`
}

type ProxyEndpointState struct {
	// The Amazon Resource Name (ARN) for the proxy endpoint.
	Arn pulumi.StringPtrInput
	// The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	DbProxyEndpointName pulumi.StringPtrInput
	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	DbProxyName pulumi.StringPtrInput
	// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
	Endpoint pulumi.StringPtrInput
	// Indicates whether this endpoint is the default endpoint for the associated DB proxy.
	IsDefault pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
	TargetRole pulumi.StringPtrInput
	// The VPC ID of the DB proxy endpoint.
	VpcId pulumi.StringPtrInput
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds pulumi.StringArrayInput
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds pulumi.StringArrayInput
}

func (ProxyEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyEndpointState)(nil)).Elem()
}

type proxyEndpointArgs struct {
	// The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	DbProxyEndpointName string `pulumi:"dbProxyEndpointName"`
	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	DbProxyName string `pulumi:"dbProxyName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
	TargetRole *string `pulumi:"targetRole"`
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds []string `pulumi:"vpcSubnetIds"`
}

// The set of arguments for constructing a ProxyEndpoint resource.
type ProxyEndpointArgs struct {
	// The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	DbProxyEndpointName pulumi.StringInput
	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	DbProxyName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
	TargetRole pulumi.StringPtrInput
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds pulumi.StringArrayInput
	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds pulumi.StringArrayInput
}

func (ProxyEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyEndpointArgs)(nil)).Elem()
}

type ProxyEndpointInput interface {
	pulumi.Input

	ToProxyEndpointOutput() ProxyEndpointOutput
	ToProxyEndpointOutputWithContext(ctx context.Context) ProxyEndpointOutput
}

func (*ProxyEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyEndpoint)(nil)).Elem()
}

func (i *ProxyEndpoint) ToProxyEndpointOutput() ProxyEndpointOutput {
	return i.ToProxyEndpointOutputWithContext(context.Background())
}

func (i *ProxyEndpoint) ToProxyEndpointOutputWithContext(ctx context.Context) ProxyEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyEndpointOutput)
}

// ProxyEndpointArrayInput is an input type that accepts ProxyEndpointArray and ProxyEndpointArrayOutput values.
// You can construct a concrete instance of `ProxyEndpointArrayInput` via:
//
//	ProxyEndpointArray{ ProxyEndpointArgs{...} }
type ProxyEndpointArrayInput interface {
	pulumi.Input

	ToProxyEndpointArrayOutput() ProxyEndpointArrayOutput
	ToProxyEndpointArrayOutputWithContext(context.Context) ProxyEndpointArrayOutput
}

type ProxyEndpointArray []ProxyEndpointInput

func (ProxyEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProxyEndpoint)(nil)).Elem()
}

func (i ProxyEndpointArray) ToProxyEndpointArrayOutput() ProxyEndpointArrayOutput {
	return i.ToProxyEndpointArrayOutputWithContext(context.Background())
}

func (i ProxyEndpointArray) ToProxyEndpointArrayOutputWithContext(ctx context.Context) ProxyEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyEndpointArrayOutput)
}

// ProxyEndpointMapInput is an input type that accepts ProxyEndpointMap and ProxyEndpointMapOutput values.
// You can construct a concrete instance of `ProxyEndpointMapInput` via:
//
//	ProxyEndpointMap{ "key": ProxyEndpointArgs{...} }
type ProxyEndpointMapInput interface {
	pulumi.Input

	ToProxyEndpointMapOutput() ProxyEndpointMapOutput
	ToProxyEndpointMapOutputWithContext(context.Context) ProxyEndpointMapOutput
}

type ProxyEndpointMap map[string]ProxyEndpointInput

func (ProxyEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProxyEndpoint)(nil)).Elem()
}

func (i ProxyEndpointMap) ToProxyEndpointMapOutput() ProxyEndpointMapOutput {
	return i.ToProxyEndpointMapOutputWithContext(context.Background())
}

func (i ProxyEndpointMap) ToProxyEndpointMapOutputWithContext(ctx context.Context) ProxyEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyEndpointMapOutput)
}

type ProxyEndpointOutput struct{ *pulumi.OutputState }

func (ProxyEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyEndpoint)(nil)).Elem()
}

func (o ProxyEndpointOutput) ToProxyEndpointOutput() ProxyEndpointOutput {
	return o
}

func (o ProxyEndpointOutput) ToProxyEndpointOutputWithContext(ctx context.Context) ProxyEndpointOutput {
	return o
}

// The Amazon Resource Name (ARN) for the proxy endpoint.
func (o ProxyEndpointOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ProxyEndpoint) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
func (o ProxyEndpointOutput) DbProxyEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *ProxyEndpoint) pulumi.StringOutput { return v.DbProxyEndpointName }).(pulumi.StringOutput)
}

// The name of the DB proxy associated with the DB proxy endpoint that you create.
func (o ProxyEndpointOutput) DbProxyName() pulumi.StringOutput {
	return o.ApplyT(func(v *ProxyEndpoint) pulumi.StringOutput { return v.DbProxyName }).(pulumi.StringOutput)
}

// The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application.
func (o ProxyEndpointOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ProxyEndpoint) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Indicates whether this endpoint is the default endpoint for the associated DB proxy.
func (o ProxyEndpointOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProxyEndpoint) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// A mapping of tags to assign to the resource.
func (o ProxyEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProxyEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Deprecated: Please use `tags` instead.
func (o ProxyEndpointOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProxyEndpoint) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
func (o ProxyEndpointOutput) TargetRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyEndpoint) pulumi.StringPtrOutput { return v.TargetRole }).(pulumi.StringPtrOutput)
}

// The VPC ID of the DB proxy endpoint.
func (o ProxyEndpointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProxyEndpoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// One or more VPC security group IDs to associate with the new proxy.
func (o ProxyEndpointOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProxyEndpoint) pulumi.StringArrayOutput { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

// One or more VPC subnet IDs to associate with the new proxy.
func (o ProxyEndpointOutput) VpcSubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProxyEndpoint) pulumi.StringArrayOutput { return v.VpcSubnetIds }).(pulumi.StringArrayOutput)
}

type ProxyEndpointArrayOutput struct{ *pulumi.OutputState }

func (ProxyEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProxyEndpoint)(nil)).Elem()
}

func (o ProxyEndpointArrayOutput) ToProxyEndpointArrayOutput() ProxyEndpointArrayOutput {
	return o
}

func (o ProxyEndpointArrayOutput) ToProxyEndpointArrayOutputWithContext(ctx context.Context) ProxyEndpointArrayOutput {
	return o
}

func (o ProxyEndpointArrayOutput) Index(i pulumi.IntInput) ProxyEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProxyEndpoint {
		return vs[0].([]*ProxyEndpoint)[vs[1].(int)]
	}).(ProxyEndpointOutput)
}

type ProxyEndpointMapOutput struct{ *pulumi.OutputState }

func (ProxyEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProxyEndpoint)(nil)).Elem()
}

func (o ProxyEndpointMapOutput) ToProxyEndpointMapOutput() ProxyEndpointMapOutput {
	return o
}

func (o ProxyEndpointMapOutput) ToProxyEndpointMapOutputWithContext(ctx context.Context) ProxyEndpointMapOutput {
	return o
}

func (o ProxyEndpointMapOutput) MapIndex(k pulumi.StringInput) ProxyEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProxyEndpoint {
		return vs[0].(map[string]*ProxyEndpoint)[vs[1].(string)]
	}).(ProxyEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyEndpointInput)(nil)).Elem(), &ProxyEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyEndpointArrayInput)(nil)).Elem(), ProxyEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyEndpointMapInput)(nil)).Elem(), ProxyEndpointMap{})
	pulumi.RegisterOutputType(ProxyEndpointOutput{})
	pulumi.RegisterOutputType(ProxyEndpointArrayOutput{})
	pulumi.RegisterOutputType(ProxyEndpointMapOutput{})
}
