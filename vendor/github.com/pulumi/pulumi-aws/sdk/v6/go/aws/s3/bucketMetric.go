// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a S3 bucket [metrics configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html) resource.
//
// > This resource cannot be used with S3 directory buckets.
//
// ## Example Usage
// ### Add metrics configuration for entire S3 bucket
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := s3.NewBucketV2(ctx, "example", nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketMetric(ctx, "example-entire-bucket", &s3.BucketMetricArgs{
//				Bucket: example.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Add metrics configuration with S3 object filter
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := s3.NewBucketV2(ctx, "example", nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketMetric(ctx, "example-filtered", &s3.BucketMetricArgs{
//				Bucket: example.ID(),
//				Filter: &s3.BucketMetricFilterArgs{
//					Prefix: pulumi.String("documents/"),
//					Tags: pulumi.StringMap{
//						"priority": pulumi.String("high"),
//						"class":    pulumi.String("blue"),
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
// ### Add metrics configuration with S3 object filter for S3 Access Point
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := s3.NewBucketV2(ctx, "example", nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewAccessPoint(ctx, "example-access-point", &s3.AccessPointArgs{
//				Bucket: example.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketMetric(ctx, "example-filtered", &s3.BucketMetricArgs{
//				Bucket: example.ID(),
//				Filter: &s3.BucketMetricFilterArgs{
//					AccessPoint: example_access_point.Arn,
//					Tags: pulumi.StringMap{
//						"priority": pulumi.String("high"),
//						"class":    pulumi.String("blue"),
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
//
// ## Import
//
// Using `pulumi import`, import S3 bucket metric configurations using `bucket:metric`. For example:
//
// ```sh
//
//	$ pulumi import aws:s3/bucketMetric:BucketMetric my-bucket-entire-bucket my-bucket:EntireBucket
//
// ```
type BucketMetric struct {
	pulumi.CustomResourceState

	// Name of the bucket to put metric configuration.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter BucketMetricFilterPtrOutput `pulumi:"filter"`
	// Unique identifier of the metrics configuration for the bucket. Must be less than or equal to 64 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewBucketMetric registers a new resource with the given unique name, arguments, and options.
func NewBucketMetric(ctx *pulumi.Context,
	name string, args *BucketMetricArgs, opts ...pulumi.ResourceOption) (*BucketMetric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketMetric
	err := ctx.RegisterResource("aws:s3/bucketMetric:BucketMetric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketMetric gets an existing BucketMetric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketMetric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketMetricState, opts ...pulumi.ResourceOption) (*BucketMetric, error) {
	var resource BucketMetric
	err := ctx.ReadResource("aws:s3/bucketMetric:BucketMetric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketMetric resources.
type bucketMetricState struct {
	// Name of the bucket to put metric configuration.
	Bucket *string `pulumi:"bucket"`
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter *BucketMetricFilter `pulumi:"filter"`
	// Unique identifier of the metrics configuration for the bucket. Must be less than or equal to 64 characters in length.
	Name *string `pulumi:"name"`
}

type BucketMetricState struct {
	// Name of the bucket to put metric configuration.
	Bucket pulumi.StringPtrInput
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter BucketMetricFilterPtrInput
	// Unique identifier of the metrics configuration for the bucket. Must be less than or equal to 64 characters in length.
	Name pulumi.StringPtrInput
}

func (BucketMetricState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketMetricState)(nil)).Elem()
}

type bucketMetricArgs struct {
	// Name of the bucket to put metric configuration.
	Bucket string `pulumi:"bucket"`
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter *BucketMetricFilter `pulumi:"filter"`
	// Unique identifier of the metrics configuration for the bucket. Must be less than or equal to 64 characters in length.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a BucketMetric resource.
type BucketMetricArgs struct {
	// Name of the bucket to put metric configuration.
	Bucket pulumi.StringInput
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter BucketMetricFilterPtrInput
	// Unique identifier of the metrics configuration for the bucket. Must be less than or equal to 64 characters in length.
	Name pulumi.StringPtrInput
}

func (BucketMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketMetricArgs)(nil)).Elem()
}

type BucketMetricInput interface {
	pulumi.Input

	ToBucketMetricOutput() BucketMetricOutput
	ToBucketMetricOutputWithContext(ctx context.Context) BucketMetricOutput
}

func (*BucketMetric) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketMetric)(nil)).Elem()
}

func (i *BucketMetric) ToBucketMetricOutput() BucketMetricOutput {
	return i.ToBucketMetricOutputWithContext(context.Background())
}

func (i *BucketMetric) ToBucketMetricOutputWithContext(ctx context.Context) BucketMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMetricOutput)
}

// BucketMetricArrayInput is an input type that accepts BucketMetricArray and BucketMetricArrayOutput values.
// You can construct a concrete instance of `BucketMetricArrayInput` via:
//
//	BucketMetricArray{ BucketMetricArgs{...} }
type BucketMetricArrayInput interface {
	pulumi.Input

	ToBucketMetricArrayOutput() BucketMetricArrayOutput
	ToBucketMetricArrayOutputWithContext(context.Context) BucketMetricArrayOutput
}

type BucketMetricArray []BucketMetricInput

func (BucketMetricArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketMetric)(nil)).Elem()
}

func (i BucketMetricArray) ToBucketMetricArrayOutput() BucketMetricArrayOutput {
	return i.ToBucketMetricArrayOutputWithContext(context.Background())
}

func (i BucketMetricArray) ToBucketMetricArrayOutputWithContext(ctx context.Context) BucketMetricArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMetricArrayOutput)
}

// BucketMetricMapInput is an input type that accepts BucketMetricMap and BucketMetricMapOutput values.
// You can construct a concrete instance of `BucketMetricMapInput` via:
//
//	BucketMetricMap{ "key": BucketMetricArgs{...} }
type BucketMetricMapInput interface {
	pulumi.Input

	ToBucketMetricMapOutput() BucketMetricMapOutput
	ToBucketMetricMapOutputWithContext(context.Context) BucketMetricMapOutput
}

type BucketMetricMap map[string]BucketMetricInput

func (BucketMetricMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketMetric)(nil)).Elem()
}

func (i BucketMetricMap) ToBucketMetricMapOutput() BucketMetricMapOutput {
	return i.ToBucketMetricMapOutputWithContext(context.Background())
}

func (i BucketMetricMap) ToBucketMetricMapOutputWithContext(ctx context.Context) BucketMetricMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMetricMapOutput)
}

type BucketMetricOutput struct{ *pulumi.OutputState }

func (BucketMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketMetric)(nil)).Elem()
}

func (o BucketMetricOutput) ToBucketMetricOutput() BucketMetricOutput {
	return o
}

func (o BucketMetricOutput) ToBucketMetricOutputWithContext(ctx context.Context) BucketMetricOutput {
	return o
}

// Name of the bucket to put metric configuration.
func (o BucketMetricOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketMetric) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
func (o BucketMetricOutput) Filter() BucketMetricFilterPtrOutput {
	return o.ApplyT(func(v *BucketMetric) BucketMetricFilterPtrOutput { return v.Filter }).(BucketMetricFilterPtrOutput)
}

// Unique identifier of the metrics configuration for the bucket. Must be less than or equal to 64 characters in length.
func (o BucketMetricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketMetric) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type BucketMetricArrayOutput struct{ *pulumi.OutputState }

func (BucketMetricArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketMetric)(nil)).Elem()
}

func (o BucketMetricArrayOutput) ToBucketMetricArrayOutput() BucketMetricArrayOutput {
	return o
}

func (o BucketMetricArrayOutput) ToBucketMetricArrayOutputWithContext(ctx context.Context) BucketMetricArrayOutput {
	return o
}

func (o BucketMetricArrayOutput) Index(i pulumi.IntInput) BucketMetricOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketMetric {
		return vs[0].([]*BucketMetric)[vs[1].(int)]
	}).(BucketMetricOutput)
}

type BucketMetricMapOutput struct{ *pulumi.OutputState }

func (BucketMetricMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketMetric)(nil)).Elem()
}

func (o BucketMetricMapOutput) ToBucketMetricMapOutput() BucketMetricMapOutput {
	return o
}

func (o BucketMetricMapOutput) ToBucketMetricMapOutputWithContext(ctx context.Context) BucketMetricMapOutput {
	return o
}

func (o BucketMetricMapOutput) MapIndex(k pulumi.StringInput) BucketMetricOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketMetric {
		return vs[0].(map[string]*BucketMetric)[vs[1].(string)]
	}).(BucketMetricOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketMetricInput)(nil)).Elem(), &BucketMetric{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketMetricArrayInput)(nil)).Elem(), BucketMetricArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketMetricMapInput)(nil)).Elem(), BucketMetricMap{})
	pulumi.RegisterOutputType(BucketMetricOutput{})
	pulumi.RegisterOutputType(BucketMetricArrayOutput{})
	pulumi.RegisterOutputType(BucketMetricMapOutput{})
}
