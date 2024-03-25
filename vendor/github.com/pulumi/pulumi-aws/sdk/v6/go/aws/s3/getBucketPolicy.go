// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The bucket policy data source returns IAM policy of an S3 bucket.
//
// ## Example Usage
//
// The following example retrieves IAM policy of a specified S3 bucket.
//
// <!--Start PulumiCodeChooser -->
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
//			example, err := s3.LookupBucketPolicy(ctx, &s3.LookupBucketPolicyArgs{
//				Bucket: "example-bucket-name",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("foo", example.Policy)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupBucketPolicy(ctx *pulumi.Context, args *LookupBucketPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBucketPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBucketPolicyResult
	err := ctx.Invoke("aws:s3/getBucketPolicy:getBucketPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBucketPolicy.
type LookupBucketPolicyArgs struct {
	// Bucket name.
	Bucket string `pulumi:"bucket"`
}

// A collection of values returned by getBucketPolicy.
type LookupBucketPolicyResult struct {
	Bucket string `pulumi:"bucket"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IAM bucket policy.
	Policy string `pulumi:"policy"`
}

func LookupBucketPolicyOutput(ctx *pulumi.Context, args LookupBucketPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBucketPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBucketPolicyResult, error) {
			args := v.(LookupBucketPolicyArgs)
			r, err := LookupBucketPolicy(ctx, &args, opts...)
			var s LookupBucketPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBucketPolicyResultOutput)
}

// A collection of arguments for invoking getBucketPolicy.
type LookupBucketPolicyOutputArgs struct {
	// Bucket name.
	Bucket pulumi.StringInput `pulumi:"bucket"`
}

func (LookupBucketPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getBucketPolicy.
type LookupBucketPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBucketPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketPolicyResult)(nil)).Elem()
}

func (o LookupBucketPolicyResultOutput) ToLookupBucketPolicyResultOutput() LookupBucketPolicyResultOutput {
	return o
}

func (o LookupBucketPolicyResultOutput) ToLookupBucketPolicyResultOutputWithContext(ctx context.Context) LookupBucketPolicyResultOutput {
	return o
}

func (o LookupBucketPolicyResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketPolicyResult) string { return v.Bucket }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBucketPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// IAM bucket policy.
func (o LookupBucketPolicyResultOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketPolicyResult) string { return v.Policy }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBucketPolicyResultOutput{})
}
