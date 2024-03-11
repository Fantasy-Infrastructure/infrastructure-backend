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

// Attaches a policy to an S3 bucket resource.
//
// > Policies can be attached to both S3 general purpose buckets and S3 directory buckets.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := s3.NewBucketV2(ctx, "example", &s3.BucketV2Args{
//				Bucket: pulumi.String("my-tf-test-bucket"),
//			})
//			if err != nil {
//				return err
//			}
//			allowAccessFromAnotherAccount := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
//				Statements: iam.GetPolicyDocumentStatementArray{
//					&iam.GetPolicyDocumentStatementArgs{
//						Principals: iam.GetPolicyDocumentStatementPrincipalArray{
//							&iam.GetPolicyDocumentStatementPrincipalArgs{
//								Type: pulumi.String("AWS"),
//								Identifiers: pulumi.StringArray{
//									pulumi.String("123456789012"),
//								},
//							},
//						},
//						Actions: pulumi.StringArray{
//							pulumi.String("s3:GetObject"),
//							pulumi.String("s3:ListBucket"),
//						},
//						Resources: pulumi.StringArray{
//							example.Arn,
//							example.Arn.ApplyT(func(arn string) (string, error) {
//								return fmt.Sprintf("%v/*", arn), nil
//							}).(pulumi.StringOutput),
//						},
//					},
//				},
//			}, nil)
//			_, err = s3.NewBucketPolicy(ctx, "allow_access_from_another_account", &s3.BucketPolicyArgs{
//				Bucket: example.ID(),
//				Policy: allowAccessFromAnotherAccount.ApplyT(func(allowAccessFromAnotherAccount iam.GetPolicyDocumentResult) (*string, error) {
//					return &allowAccessFromAnotherAccount.Json, nil
//				}).(pulumi.StringPtrOutput),
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
// Using `pulumi import`, import S3 bucket policies using the bucket name. For example:
//
// ```sh
//
//	$ pulumi import aws:s3/bucketPolicy:BucketPolicy allow_access_from_another_account my-tf-test-bucket
//
// ```
type BucketPolicy struct {
	pulumi.CustomResourceState

	// Name of the bucket to which to apply the policy.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Text of the policy. Although this is a bucket policy rather than an IAM policy, the `iam.getPolicyDocument` data source may be used, so long as it specifies a principal. For more information about building AWS IAM policy documents, see the AWS IAM Policy Document Guide. Note: Bucket policies are limited to 20 KB in size.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewBucketPolicy registers a new resource with the given unique name, arguments, and options.
func NewBucketPolicy(ctx *pulumi.Context,
	name string, args *BucketPolicyArgs, opts ...pulumi.ResourceOption) (*BucketPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketPolicy
	err := ctx.RegisterResource("aws:s3/bucketPolicy:BucketPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketPolicy gets an existing BucketPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketPolicyState, opts ...pulumi.ResourceOption) (*BucketPolicy, error) {
	var resource BucketPolicy
	err := ctx.ReadResource("aws:s3/bucketPolicy:BucketPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketPolicy resources.
type bucketPolicyState struct {
	// Name of the bucket to which to apply the policy.
	Bucket *string `pulumi:"bucket"`
	// Text of the policy. Although this is a bucket policy rather than an IAM policy, the `iam.getPolicyDocument` data source may be used, so long as it specifies a principal. For more information about building AWS IAM policy documents, see the AWS IAM Policy Document Guide. Note: Bucket policies are limited to 20 KB in size.
	Policy interface{} `pulumi:"policy"`
}

type BucketPolicyState struct {
	// Name of the bucket to which to apply the policy.
	Bucket pulumi.StringPtrInput
	// Text of the policy. Although this is a bucket policy rather than an IAM policy, the `iam.getPolicyDocument` data source may be used, so long as it specifies a principal. For more information about building AWS IAM policy documents, see the AWS IAM Policy Document Guide. Note: Bucket policies are limited to 20 KB in size.
	Policy pulumi.Input
}

func (BucketPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPolicyState)(nil)).Elem()
}

type bucketPolicyArgs struct {
	// Name of the bucket to which to apply the policy.
	Bucket string `pulumi:"bucket"`
	// Text of the policy. Although this is a bucket policy rather than an IAM policy, the `iam.getPolicyDocument` data source may be used, so long as it specifies a principal. For more information about building AWS IAM policy documents, see the AWS IAM Policy Document Guide. Note: Bucket policies are limited to 20 KB in size.
	Policy interface{} `pulumi:"policy"`
}

// The set of arguments for constructing a BucketPolicy resource.
type BucketPolicyArgs struct {
	// Name of the bucket to which to apply the policy.
	Bucket pulumi.StringInput
	// Text of the policy. Although this is a bucket policy rather than an IAM policy, the `iam.getPolicyDocument` data source may be used, so long as it specifies a principal. For more information about building AWS IAM policy documents, see the AWS IAM Policy Document Guide. Note: Bucket policies are limited to 20 KB in size.
	Policy pulumi.Input
}

func (BucketPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPolicyArgs)(nil)).Elem()
}

type BucketPolicyInput interface {
	pulumi.Input

	ToBucketPolicyOutput() BucketPolicyOutput
	ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput
}

func (*BucketPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketPolicy)(nil)).Elem()
}

func (i *BucketPolicy) ToBucketPolicyOutput() BucketPolicyOutput {
	return i.ToBucketPolicyOutputWithContext(context.Background())
}

func (i *BucketPolicy) ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyOutput)
}

// BucketPolicyArrayInput is an input type that accepts BucketPolicyArray and BucketPolicyArrayOutput values.
// You can construct a concrete instance of `BucketPolicyArrayInput` via:
//
//	BucketPolicyArray{ BucketPolicyArgs{...} }
type BucketPolicyArrayInput interface {
	pulumi.Input

	ToBucketPolicyArrayOutput() BucketPolicyArrayOutput
	ToBucketPolicyArrayOutputWithContext(context.Context) BucketPolicyArrayOutput
}

type BucketPolicyArray []BucketPolicyInput

func (BucketPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketPolicy)(nil)).Elem()
}

func (i BucketPolicyArray) ToBucketPolicyArrayOutput() BucketPolicyArrayOutput {
	return i.ToBucketPolicyArrayOutputWithContext(context.Background())
}

func (i BucketPolicyArray) ToBucketPolicyArrayOutputWithContext(ctx context.Context) BucketPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyArrayOutput)
}

// BucketPolicyMapInput is an input type that accepts BucketPolicyMap and BucketPolicyMapOutput values.
// You can construct a concrete instance of `BucketPolicyMapInput` via:
//
//	BucketPolicyMap{ "key": BucketPolicyArgs{...} }
type BucketPolicyMapInput interface {
	pulumi.Input

	ToBucketPolicyMapOutput() BucketPolicyMapOutput
	ToBucketPolicyMapOutputWithContext(context.Context) BucketPolicyMapOutput
}

type BucketPolicyMap map[string]BucketPolicyInput

func (BucketPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketPolicy)(nil)).Elem()
}

func (i BucketPolicyMap) ToBucketPolicyMapOutput() BucketPolicyMapOutput {
	return i.ToBucketPolicyMapOutputWithContext(context.Background())
}

func (i BucketPolicyMap) ToBucketPolicyMapOutputWithContext(ctx context.Context) BucketPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyMapOutput)
}

type BucketPolicyOutput struct{ *pulumi.OutputState }

func (BucketPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketPolicy)(nil)).Elem()
}

func (o BucketPolicyOutput) ToBucketPolicyOutput() BucketPolicyOutput {
	return o
}

func (o BucketPolicyOutput) ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput {
	return o
}

// Name of the bucket to which to apply the policy.
func (o BucketPolicyOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketPolicy) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Text of the policy. Although this is a bucket policy rather than an IAM policy, the `iam.getPolicyDocument` data source may be used, so long as it specifies a principal. For more information about building AWS IAM policy documents, see the AWS IAM Policy Document Guide. Note: Bucket policies are limited to 20 KB in size.
func (o BucketPolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketPolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

type BucketPolicyArrayOutput struct{ *pulumi.OutputState }

func (BucketPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketPolicy)(nil)).Elem()
}

func (o BucketPolicyArrayOutput) ToBucketPolicyArrayOutput() BucketPolicyArrayOutput {
	return o
}

func (o BucketPolicyArrayOutput) ToBucketPolicyArrayOutputWithContext(ctx context.Context) BucketPolicyArrayOutput {
	return o
}

func (o BucketPolicyArrayOutput) Index(i pulumi.IntInput) BucketPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketPolicy {
		return vs[0].([]*BucketPolicy)[vs[1].(int)]
	}).(BucketPolicyOutput)
}

type BucketPolicyMapOutput struct{ *pulumi.OutputState }

func (BucketPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketPolicy)(nil)).Elem()
}

func (o BucketPolicyMapOutput) ToBucketPolicyMapOutput() BucketPolicyMapOutput {
	return o
}

func (o BucketPolicyMapOutput) ToBucketPolicyMapOutputWithContext(ctx context.Context) BucketPolicyMapOutput {
	return o
}

func (o BucketPolicyMapOutput) MapIndex(k pulumi.StringInput) BucketPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketPolicy {
		return vs[0].(map[string]*BucketPolicy)[vs[1].(string)]
	}).(BucketPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPolicyInput)(nil)).Elem(), &BucketPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPolicyArrayInput)(nil)).Elem(), BucketPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPolicyMapInput)(nil)).Elem(), BucketPolicyMap{})
	pulumi.RegisterOutputType(BucketPolicyOutput{})
	pulumi.RegisterOutputType(BucketPolicyArrayOutput{})
	pulumi.RegisterOutputType(BucketPolicyMapOutput{})
}
