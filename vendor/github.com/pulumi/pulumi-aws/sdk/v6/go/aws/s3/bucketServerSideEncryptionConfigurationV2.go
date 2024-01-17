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

// Provides a S3 bucket server-side encryption configuration resource.
//
// > **NOTE:** Destroying an `s3.BucketServerSideEncryptionConfigurationV2` resource resets the bucket to [Amazon S3 bucket default encryption](https://docs.aws.amazon.com/AmazonS3/latest/userguide/default-encryption-faq.html).
//
// > This resource cannot be used with S3 directory buckets.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mykey, err := kms.NewKey(ctx, "mykey", &kms.KeyArgs{
//				Description:          pulumi.String("This key is used to encrypt bucket objects"),
//				DeletionWindowInDays: pulumi.Int(10),
//			})
//			if err != nil {
//				return err
//			}
//			mybucket, err := s3.NewBucketV2(ctx, "mybucket", nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketServerSideEncryptionConfigurationV2(ctx, "example", &s3.BucketServerSideEncryptionConfigurationV2Args{
//				Bucket: mybucket.ID(),
//				Rules: s3.BucketServerSideEncryptionConfigurationV2RuleArray{
//					&s3.BucketServerSideEncryptionConfigurationV2RuleArgs{
//						ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs{
//							KmsMasterKeyId: mykey.Arn,
//							SseAlgorithm:   pulumi.String("aws:kms"),
//						},
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
// If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):
//
// __Using `pulumi import` to import__ S3 bucket server-side encryption configuration using the `bucket` or using the `bucket` and `expected_bucket_owner` separated by a comma (`,`). For example:
//
// If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, import using the `bucket`:
//
// ```sh
//
//	$ pulumi import aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2 example bucket-name
//
// ```
//
//	If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):
//
// ```sh
//
//	$ pulumi import aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2 example bucket-name,123456789012
//
// ```
type BucketServerSideEncryptionConfigurationV2 struct {
	pulumi.CustomResourceState

	// ID (name) of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrOutput `pulumi:"expectedBucketOwner"`
	// Set of server-side encryption configuration rules. See below. Currently, only a single rule is supported.
	Rules BucketServerSideEncryptionConfigurationV2RuleArrayOutput `pulumi:"rules"`
}

// NewBucketServerSideEncryptionConfigurationV2 registers a new resource with the given unique name, arguments, and options.
func NewBucketServerSideEncryptionConfigurationV2(ctx *pulumi.Context,
	name string, args *BucketServerSideEncryptionConfigurationV2Args, opts ...pulumi.ResourceOption) (*BucketServerSideEncryptionConfigurationV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketServerSideEncryptionConfigurationV2
	err := ctx.RegisterResource("aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketServerSideEncryptionConfigurationV2 gets an existing BucketServerSideEncryptionConfigurationV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketServerSideEncryptionConfigurationV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketServerSideEncryptionConfigurationV2State, opts ...pulumi.ResourceOption) (*BucketServerSideEncryptionConfigurationV2, error) {
	var resource BucketServerSideEncryptionConfigurationV2
	err := ctx.ReadResource("aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketServerSideEncryptionConfigurationV2 resources.
type bucketServerSideEncryptionConfigurationV2State struct {
	// ID (name) of the bucket.
	Bucket *string `pulumi:"bucket"`
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// Set of server-side encryption configuration rules. See below. Currently, only a single rule is supported.
	Rules []BucketServerSideEncryptionConfigurationV2Rule `pulumi:"rules"`
}

type BucketServerSideEncryptionConfigurationV2State struct {
	// ID (name) of the bucket.
	Bucket pulumi.StringPtrInput
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrInput
	// Set of server-side encryption configuration rules. See below. Currently, only a single rule is supported.
	Rules BucketServerSideEncryptionConfigurationV2RuleArrayInput
}

func (BucketServerSideEncryptionConfigurationV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketServerSideEncryptionConfigurationV2State)(nil)).Elem()
}

type bucketServerSideEncryptionConfigurationV2Args struct {
	// ID (name) of the bucket.
	Bucket string `pulumi:"bucket"`
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// Set of server-side encryption configuration rules. See below. Currently, only a single rule is supported.
	Rules []BucketServerSideEncryptionConfigurationV2Rule `pulumi:"rules"`
}

// The set of arguments for constructing a BucketServerSideEncryptionConfigurationV2 resource.
type BucketServerSideEncryptionConfigurationV2Args struct {
	// ID (name) of the bucket.
	Bucket pulumi.StringInput
	// Account ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrInput
	// Set of server-side encryption configuration rules. See below. Currently, only a single rule is supported.
	Rules BucketServerSideEncryptionConfigurationV2RuleArrayInput
}

func (BucketServerSideEncryptionConfigurationV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketServerSideEncryptionConfigurationV2Args)(nil)).Elem()
}

type BucketServerSideEncryptionConfigurationV2Input interface {
	pulumi.Input

	ToBucketServerSideEncryptionConfigurationV2Output() BucketServerSideEncryptionConfigurationV2Output
	ToBucketServerSideEncryptionConfigurationV2OutputWithContext(ctx context.Context) BucketServerSideEncryptionConfigurationV2Output
}

func (*BucketServerSideEncryptionConfigurationV2) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketServerSideEncryptionConfigurationV2)(nil)).Elem()
}

func (i *BucketServerSideEncryptionConfigurationV2) ToBucketServerSideEncryptionConfigurationV2Output() BucketServerSideEncryptionConfigurationV2Output {
	return i.ToBucketServerSideEncryptionConfigurationV2OutputWithContext(context.Background())
}

func (i *BucketServerSideEncryptionConfigurationV2) ToBucketServerSideEncryptionConfigurationV2OutputWithContext(ctx context.Context) BucketServerSideEncryptionConfigurationV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(BucketServerSideEncryptionConfigurationV2Output)
}

// BucketServerSideEncryptionConfigurationV2ArrayInput is an input type that accepts BucketServerSideEncryptionConfigurationV2Array and BucketServerSideEncryptionConfigurationV2ArrayOutput values.
// You can construct a concrete instance of `BucketServerSideEncryptionConfigurationV2ArrayInput` via:
//
//	BucketServerSideEncryptionConfigurationV2Array{ BucketServerSideEncryptionConfigurationV2Args{...} }
type BucketServerSideEncryptionConfigurationV2ArrayInput interface {
	pulumi.Input

	ToBucketServerSideEncryptionConfigurationV2ArrayOutput() BucketServerSideEncryptionConfigurationV2ArrayOutput
	ToBucketServerSideEncryptionConfigurationV2ArrayOutputWithContext(context.Context) BucketServerSideEncryptionConfigurationV2ArrayOutput
}

type BucketServerSideEncryptionConfigurationV2Array []BucketServerSideEncryptionConfigurationV2Input

func (BucketServerSideEncryptionConfigurationV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketServerSideEncryptionConfigurationV2)(nil)).Elem()
}

func (i BucketServerSideEncryptionConfigurationV2Array) ToBucketServerSideEncryptionConfigurationV2ArrayOutput() BucketServerSideEncryptionConfigurationV2ArrayOutput {
	return i.ToBucketServerSideEncryptionConfigurationV2ArrayOutputWithContext(context.Background())
}

func (i BucketServerSideEncryptionConfigurationV2Array) ToBucketServerSideEncryptionConfigurationV2ArrayOutputWithContext(ctx context.Context) BucketServerSideEncryptionConfigurationV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketServerSideEncryptionConfigurationV2ArrayOutput)
}

// BucketServerSideEncryptionConfigurationV2MapInput is an input type that accepts BucketServerSideEncryptionConfigurationV2Map and BucketServerSideEncryptionConfigurationV2MapOutput values.
// You can construct a concrete instance of `BucketServerSideEncryptionConfigurationV2MapInput` via:
//
//	BucketServerSideEncryptionConfigurationV2Map{ "key": BucketServerSideEncryptionConfigurationV2Args{...} }
type BucketServerSideEncryptionConfigurationV2MapInput interface {
	pulumi.Input

	ToBucketServerSideEncryptionConfigurationV2MapOutput() BucketServerSideEncryptionConfigurationV2MapOutput
	ToBucketServerSideEncryptionConfigurationV2MapOutputWithContext(context.Context) BucketServerSideEncryptionConfigurationV2MapOutput
}

type BucketServerSideEncryptionConfigurationV2Map map[string]BucketServerSideEncryptionConfigurationV2Input

func (BucketServerSideEncryptionConfigurationV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketServerSideEncryptionConfigurationV2)(nil)).Elem()
}

func (i BucketServerSideEncryptionConfigurationV2Map) ToBucketServerSideEncryptionConfigurationV2MapOutput() BucketServerSideEncryptionConfigurationV2MapOutput {
	return i.ToBucketServerSideEncryptionConfigurationV2MapOutputWithContext(context.Background())
}

func (i BucketServerSideEncryptionConfigurationV2Map) ToBucketServerSideEncryptionConfigurationV2MapOutputWithContext(ctx context.Context) BucketServerSideEncryptionConfigurationV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketServerSideEncryptionConfigurationV2MapOutput)
}

type BucketServerSideEncryptionConfigurationV2Output struct{ *pulumi.OutputState }

func (BucketServerSideEncryptionConfigurationV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketServerSideEncryptionConfigurationV2)(nil)).Elem()
}

func (o BucketServerSideEncryptionConfigurationV2Output) ToBucketServerSideEncryptionConfigurationV2Output() BucketServerSideEncryptionConfigurationV2Output {
	return o
}

func (o BucketServerSideEncryptionConfigurationV2Output) ToBucketServerSideEncryptionConfigurationV2OutputWithContext(ctx context.Context) BucketServerSideEncryptionConfigurationV2Output {
	return o
}

// ID (name) of the bucket.
func (o BucketServerSideEncryptionConfigurationV2Output) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketServerSideEncryptionConfigurationV2) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Account ID of the expected bucket owner.
func (o BucketServerSideEncryptionConfigurationV2Output) ExpectedBucketOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketServerSideEncryptionConfigurationV2) pulumi.StringPtrOutput {
		return v.ExpectedBucketOwner
	}).(pulumi.StringPtrOutput)
}

// Set of server-side encryption configuration rules. See below. Currently, only a single rule is supported.
func (o BucketServerSideEncryptionConfigurationV2Output) Rules() BucketServerSideEncryptionConfigurationV2RuleArrayOutput {
	return o.ApplyT(func(v *BucketServerSideEncryptionConfigurationV2) BucketServerSideEncryptionConfigurationV2RuleArrayOutput {
		return v.Rules
	}).(BucketServerSideEncryptionConfigurationV2RuleArrayOutput)
}

type BucketServerSideEncryptionConfigurationV2ArrayOutput struct{ *pulumi.OutputState }

func (BucketServerSideEncryptionConfigurationV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketServerSideEncryptionConfigurationV2)(nil)).Elem()
}

func (o BucketServerSideEncryptionConfigurationV2ArrayOutput) ToBucketServerSideEncryptionConfigurationV2ArrayOutput() BucketServerSideEncryptionConfigurationV2ArrayOutput {
	return o
}

func (o BucketServerSideEncryptionConfigurationV2ArrayOutput) ToBucketServerSideEncryptionConfigurationV2ArrayOutputWithContext(ctx context.Context) BucketServerSideEncryptionConfigurationV2ArrayOutput {
	return o
}

func (o BucketServerSideEncryptionConfigurationV2ArrayOutput) Index(i pulumi.IntInput) BucketServerSideEncryptionConfigurationV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketServerSideEncryptionConfigurationV2 {
		return vs[0].([]*BucketServerSideEncryptionConfigurationV2)[vs[1].(int)]
	}).(BucketServerSideEncryptionConfigurationV2Output)
}

type BucketServerSideEncryptionConfigurationV2MapOutput struct{ *pulumi.OutputState }

func (BucketServerSideEncryptionConfigurationV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketServerSideEncryptionConfigurationV2)(nil)).Elem()
}

func (o BucketServerSideEncryptionConfigurationV2MapOutput) ToBucketServerSideEncryptionConfigurationV2MapOutput() BucketServerSideEncryptionConfigurationV2MapOutput {
	return o
}

func (o BucketServerSideEncryptionConfigurationV2MapOutput) ToBucketServerSideEncryptionConfigurationV2MapOutputWithContext(ctx context.Context) BucketServerSideEncryptionConfigurationV2MapOutput {
	return o
}

func (o BucketServerSideEncryptionConfigurationV2MapOutput) MapIndex(k pulumi.StringInput) BucketServerSideEncryptionConfigurationV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketServerSideEncryptionConfigurationV2 {
		return vs[0].(map[string]*BucketServerSideEncryptionConfigurationV2)[vs[1].(string)]
	}).(BucketServerSideEncryptionConfigurationV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketServerSideEncryptionConfigurationV2Input)(nil)).Elem(), &BucketServerSideEncryptionConfigurationV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketServerSideEncryptionConfigurationV2ArrayInput)(nil)).Elem(), BucketServerSideEncryptionConfigurationV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketServerSideEncryptionConfigurationV2MapInput)(nil)).Elem(), BucketServerSideEncryptionConfigurationV2Map{})
	pulumi.RegisterOutputType(BucketServerSideEncryptionConfigurationV2Output{})
	pulumi.RegisterOutputType(BucketServerSideEncryptionConfigurationV2ArrayOutput{})
	pulumi.RegisterOutputType(BucketServerSideEncryptionConfigurationV2MapOutput{})
}
