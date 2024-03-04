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

// Provides a Route 53 Resolver query logging configuration resource.
//
// ## Example Usage
//
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
//			_, err := route53.NewResolverQueryLogConfig(ctx, "example", &route53.ResolverQueryLogConfigArgs{
//				DestinationArn: pulumi.Any(aws_s3_bucket.Example.Arn),
//				Tags: pulumi.StringMap{
//					"Environment": pulumi.String("Prod"),
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
// # Using `pulumi import`, import
//
// Route 53 Resolver query logging configurations using the Route 53 Resolver query logging configuration ID. For example:
//
// ```sh
//
//	$ pulumi import aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig example rqlc-92edc3b1838248bf
//
// ```
type ResolverQueryLogConfig struct {
	pulumi.CustomResourceState

	// The ARN (Amazon Resource Name) of the Route 53 Resolver query logging configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ARN of the resource that you want Route 53 Resolver to send query logs.
	// You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
	DestinationArn pulumi.StringOutput `pulumi:"destinationArn"`
	// The name of the Route 53 Resolver query logging configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// The AWS account ID of the account that created the query logging configuration.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// An indication of whether the query logging configuration is shared with other AWS accounts, or was shared with the current account by another AWS account.
	// Sharing is configured through AWS Resource Access Manager (AWS RAM).
	// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
	ShareStatus pulumi.StringOutput `pulumi:"shareStatus"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewResolverQueryLogConfig registers a new resource with the given unique name, arguments, and options.
func NewResolverQueryLogConfig(ctx *pulumi.Context,
	name string, args *ResolverQueryLogConfigArgs, opts ...pulumi.ResourceOption) (*ResolverQueryLogConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationArn == nil {
		return nil, errors.New("invalid value for required argument 'DestinationArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResolverQueryLogConfig
	err := ctx.RegisterResource("aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverQueryLogConfig gets an existing ResolverQueryLogConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverQueryLogConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverQueryLogConfigState, opts ...pulumi.ResourceOption) (*ResolverQueryLogConfig, error) {
	var resource ResolverQueryLogConfig
	err := ctx.ReadResource("aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverQueryLogConfig resources.
type resolverQueryLogConfigState struct {
	// The ARN (Amazon Resource Name) of the Route 53 Resolver query logging configuration.
	Arn *string `pulumi:"arn"`
	// The ARN of the resource that you want Route 53 Resolver to send query logs.
	// You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
	DestinationArn *string `pulumi:"destinationArn"`
	// The name of the Route 53 Resolver query logging configuration.
	Name *string `pulumi:"name"`
	// The AWS account ID of the account that created the query logging configuration.
	OwnerId *string `pulumi:"ownerId"`
	// An indication of whether the query logging configuration is shared with other AWS accounts, or was shared with the current account by another AWS account.
	// Sharing is configured through AWS Resource Access Manager (AWS RAM).
	// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
	ShareStatus *string `pulumi:"shareStatus"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ResolverQueryLogConfigState struct {
	// The ARN (Amazon Resource Name) of the Route 53 Resolver query logging configuration.
	Arn pulumi.StringPtrInput
	// The ARN of the resource that you want Route 53 Resolver to send query logs.
	// You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
	DestinationArn pulumi.StringPtrInput
	// The name of the Route 53 Resolver query logging configuration.
	Name pulumi.StringPtrInput
	// The AWS account ID of the account that created the query logging configuration.
	OwnerId pulumi.StringPtrInput
	// An indication of whether the query logging configuration is shared with other AWS accounts, or was shared with the current account by another AWS account.
	// Sharing is configured through AWS Resource Access Manager (AWS RAM).
	// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
	ShareStatus pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (ResolverQueryLogConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverQueryLogConfigState)(nil)).Elem()
}

type resolverQueryLogConfigArgs struct {
	// The ARN of the resource that you want Route 53 Resolver to send query logs.
	// You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
	DestinationArn string `pulumi:"destinationArn"`
	// The name of the Route 53 Resolver query logging configuration.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ResolverQueryLogConfig resource.
type ResolverQueryLogConfigArgs struct {
	// The ARN of the resource that you want Route 53 Resolver to send query logs.
	// You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
	DestinationArn pulumi.StringInput
	// The name of the Route 53 Resolver query logging configuration.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ResolverQueryLogConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverQueryLogConfigArgs)(nil)).Elem()
}

type ResolverQueryLogConfigInput interface {
	pulumi.Input

	ToResolverQueryLogConfigOutput() ResolverQueryLogConfigOutput
	ToResolverQueryLogConfigOutputWithContext(ctx context.Context) ResolverQueryLogConfigOutput
}

func (*ResolverQueryLogConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverQueryLogConfig)(nil)).Elem()
}

func (i *ResolverQueryLogConfig) ToResolverQueryLogConfigOutput() ResolverQueryLogConfigOutput {
	return i.ToResolverQueryLogConfigOutputWithContext(context.Background())
}

func (i *ResolverQueryLogConfig) ToResolverQueryLogConfigOutputWithContext(ctx context.Context) ResolverQueryLogConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverQueryLogConfigOutput)
}

// ResolverQueryLogConfigArrayInput is an input type that accepts ResolverQueryLogConfigArray and ResolverQueryLogConfigArrayOutput values.
// You can construct a concrete instance of `ResolverQueryLogConfigArrayInput` via:
//
//	ResolverQueryLogConfigArray{ ResolverQueryLogConfigArgs{...} }
type ResolverQueryLogConfigArrayInput interface {
	pulumi.Input

	ToResolverQueryLogConfigArrayOutput() ResolverQueryLogConfigArrayOutput
	ToResolverQueryLogConfigArrayOutputWithContext(context.Context) ResolverQueryLogConfigArrayOutput
}

type ResolverQueryLogConfigArray []ResolverQueryLogConfigInput

func (ResolverQueryLogConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverQueryLogConfig)(nil)).Elem()
}

func (i ResolverQueryLogConfigArray) ToResolverQueryLogConfigArrayOutput() ResolverQueryLogConfigArrayOutput {
	return i.ToResolverQueryLogConfigArrayOutputWithContext(context.Background())
}

func (i ResolverQueryLogConfigArray) ToResolverQueryLogConfigArrayOutputWithContext(ctx context.Context) ResolverQueryLogConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverQueryLogConfigArrayOutput)
}

// ResolverQueryLogConfigMapInput is an input type that accepts ResolverQueryLogConfigMap and ResolverQueryLogConfigMapOutput values.
// You can construct a concrete instance of `ResolverQueryLogConfigMapInput` via:
//
//	ResolverQueryLogConfigMap{ "key": ResolverQueryLogConfigArgs{...} }
type ResolverQueryLogConfigMapInput interface {
	pulumi.Input

	ToResolverQueryLogConfigMapOutput() ResolverQueryLogConfigMapOutput
	ToResolverQueryLogConfigMapOutputWithContext(context.Context) ResolverQueryLogConfigMapOutput
}

type ResolverQueryLogConfigMap map[string]ResolverQueryLogConfigInput

func (ResolverQueryLogConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverQueryLogConfig)(nil)).Elem()
}

func (i ResolverQueryLogConfigMap) ToResolverQueryLogConfigMapOutput() ResolverQueryLogConfigMapOutput {
	return i.ToResolverQueryLogConfigMapOutputWithContext(context.Background())
}

func (i ResolverQueryLogConfigMap) ToResolverQueryLogConfigMapOutputWithContext(ctx context.Context) ResolverQueryLogConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverQueryLogConfigMapOutput)
}

type ResolverQueryLogConfigOutput struct{ *pulumi.OutputState }

func (ResolverQueryLogConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverQueryLogConfig)(nil)).Elem()
}

func (o ResolverQueryLogConfigOutput) ToResolverQueryLogConfigOutput() ResolverQueryLogConfigOutput {
	return o
}

func (o ResolverQueryLogConfigOutput) ToResolverQueryLogConfigOutputWithContext(ctx context.Context) ResolverQueryLogConfigOutput {
	return o
}

// The ARN (Amazon Resource Name) of the Route 53 Resolver query logging configuration.
func (o ResolverQueryLogConfigOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverQueryLogConfig) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ARN of the resource that you want Route 53 Resolver to send query logs.
// You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
func (o ResolverQueryLogConfigOutput) DestinationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverQueryLogConfig) pulumi.StringOutput { return v.DestinationArn }).(pulumi.StringOutput)
}

// The name of the Route 53 Resolver query logging configuration.
func (o ResolverQueryLogConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverQueryLogConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The AWS account ID of the account that created the query logging configuration.
func (o ResolverQueryLogConfigOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverQueryLogConfig) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// An indication of whether the query logging configuration is shared with other AWS accounts, or was shared with the current account by another AWS account.
// Sharing is configured through AWS Resource Access Manager (AWS RAM).
// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
func (o ResolverQueryLogConfigOutput) ShareStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverQueryLogConfig) pulumi.StringOutput { return v.ShareStatus }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ResolverQueryLogConfigOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResolverQueryLogConfig) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o ResolverQueryLogConfigOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResolverQueryLogConfig) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type ResolverQueryLogConfigArrayOutput struct{ *pulumi.OutputState }

func (ResolverQueryLogConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverQueryLogConfig)(nil)).Elem()
}

func (o ResolverQueryLogConfigArrayOutput) ToResolverQueryLogConfigArrayOutput() ResolverQueryLogConfigArrayOutput {
	return o
}

func (o ResolverQueryLogConfigArrayOutput) ToResolverQueryLogConfigArrayOutputWithContext(ctx context.Context) ResolverQueryLogConfigArrayOutput {
	return o
}

func (o ResolverQueryLogConfigArrayOutput) Index(i pulumi.IntInput) ResolverQueryLogConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResolverQueryLogConfig {
		return vs[0].([]*ResolverQueryLogConfig)[vs[1].(int)]
	}).(ResolverQueryLogConfigOutput)
}

type ResolverQueryLogConfigMapOutput struct{ *pulumi.OutputState }

func (ResolverQueryLogConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverQueryLogConfig)(nil)).Elem()
}

func (o ResolverQueryLogConfigMapOutput) ToResolverQueryLogConfigMapOutput() ResolverQueryLogConfigMapOutput {
	return o
}

func (o ResolverQueryLogConfigMapOutput) ToResolverQueryLogConfigMapOutputWithContext(ctx context.Context) ResolverQueryLogConfigMapOutput {
	return o
}

func (o ResolverQueryLogConfigMapOutput) MapIndex(k pulumi.StringInput) ResolverQueryLogConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResolverQueryLogConfig {
		return vs[0].(map[string]*ResolverQueryLogConfig)[vs[1].(string)]
	}).(ResolverQueryLogConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverQueryLogConfigInput)(nil)).Elem(), &ResolverQueryLogConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverQueryLogConfigArrayInput)(nil)).Elem(), ResolverQueryLogConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverQueryLogConfigMapInput)(nil)).Elem(), ResolverQueryLogConfigMap{})
	pulumi.RegisterOutputType(ResolverQueryLogConfigOutput{})
	pulumi.RegisterOutputType(ResolverQueryLogConfigArrayOutput{})
	pulumi.RegisterOutputType(ResolverQueryLogConfigMapOutput{})
}
