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

// Provides an independent configuration resource for S3 bucket [lifecycle configuration](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html).
//
// An S3 Lifecycle configuration consists of one or more Lifecycle rules. Each rule consists of the following:
//
// * Rule metadata (`id` and `status`)
// * Filter identifying objects to which the rule applies
// * One or more transition or expiration actions
//
// For more information see the Amazon S3 User Guide on [`Lifecycle Configuration Elements`](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html).
//
// > **NOTE:** S3 Buckets only support a single lifecycle configuration. Declaring multiple `s3.BucketLifecycleConfigurationV2` resources to the same S3 Bucket will cause a perpetual difference in configuration.
//
// > **NOTE:** Lifecycle configurations may take some time to fully propagate to all AWS S3 systems.
// Running Pulumi operations shortly after creating a lifecycle configuration may result in changes that affect configuration idempotence.
// See the Amazon S3 User Guide on [setting lifecycle configuration on a bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/how-to-set-lifecycle-configuration-intro.html).
//
// > This resource cannot be used with S3 directory buckets.
//
// ## Example Usage
// ### With neither a filter nor prefix specified
//
// The Lifecycle rule applies to a subset of objects based on the key name prefix (`""`).
//
// This configuration is intended to replicate the default behavior of the `lifecycleRule`
// parameter in the AWS Provider `s3.BucketV2` resource prior to `v4.0`.
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
//			_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
//				Bucket: pulumi.Any(bucket.Id),
//				Rules: s3.BucketLifecycleConfigurationV2RuleArray{
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id:     pulumi.String("rule-1"),
//						Status: pulumi.String("Enabled"),
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
// ### Specifying an empty filter
//
// The Lifecycle rule applies to all objects in the bucket.
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
//			_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
//				Bucket: pulumi.Any(bucket.Id),
//				Rules: s3.BucketLifecycleConfigurationV2RuleArray{
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id:     pulumi.String("rule-1"),
//						Filter: nil,
//						Status: pulumi.String("Enabled"),
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
// ### Specifying a filter using key prefixes
//
// The Lifecycle rule applies to a subset of objects based on the key name prefix (`logs/`).
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
//			_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
//				Bucket: pulumi.Any(bucket.Id),
//				Rules: s3.BucketLifecycleConfigurationV2RuleArray{
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id: pulumi.String("rule-1"),
//						Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
//							Prefix: pulumi.String("logs/"),
//						},
//						Status: pulumi.String("Enabled"),
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
// If you want to apply a Lifecycle action to a subset of objects based on different key name prefixes, specify separate rules.
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
//			_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
//				Bucket: pulumi.Any(bucket.Id),
//				Rules: s3.BucketLifecycleConfigurationV2RuleArray{
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id: pulumi.String("rule-1"),
//						Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
//							Prefix: pulumi.String("logs/"),
//						},
//						Status: pulumi.String("Enabled"),
//					},
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id: pulumi.String("rule-2"),
//						Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
//							Prefix: pulumi.String("tmp/"),
//						},
//						Status: pulumi.String("Enabled"),
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
// ### Specifying a filter based on an object tag
//
// The Lifecycle rule specifies a filter based on a tag key and value. The rule then applies only to a subset of objects with the specific tag.
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
//			_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
//				Bucket: pulumi.Any(bucket.Id),
//				Rules: s3.BucketLifecycleConfigurationV2RuleArray{
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id: pulumi.String("rule-1"),
//						Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
//							Tag: &s3.BucketLifecycleConfigurationV2RuleFilterTagArgs{
//								Key:   pulumi.String("Name"),
//								Value: pulumi.String("Staging"),
//							},
//						},
//						Status: pulumi.String("Enabled"),
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
// ### Specifying a filter based on multiple tags
//
// The Lifecycle rule directs Amazon S3 to perform lifecycle actions on objects with two tags (with the specific tag keys and values). Notice `tags` is wrapped in the `and` configuration block.
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
//			_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
//				Bucket: pulumi.Any(bucket.Id),
//				Rules: s3.BucketLifecycleConfigurationV2RuleArray{
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id: pulumi.String("rule-1"),
//						Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
//							And: &s3.BucketLifecycleConfigurationV2RuleFilterAndArgs{
//								Tags: pulumi.StringMap{
//									"Key1": pulumi.String("Value1"),
//									"Key2": pulumi.String("Value2"),
//								},
//							},
//						},
//						Status: pulumi.String("Enabled"),
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
// ### Specifying a filter based on both prefix and one or more tags
//
// The Lifecycle rule directs Amazon S3 to perform lifecycle actions on objects with the specified prefix and two tags (with the specific tag keys and values). Notice both `prefix` and `tags` are wrapped in the `and` configuration block.
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
//			_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
//				Bucket: pulumi.Any(bucket.Id),
//				Rules: s3.BucketLifecycleConfigurationV2RuleArray{
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id: pulumi.String("rule-1"),
//						Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
//							And: &s3.BucketLifecycleConfigurationV2RuleFilterAndArgs{
//								Prefix: pulumi.String("logs/"),
//								Tags: pulumi.StringMap{
//									"Key1": pulumi.String("Value1"),
//									"Key2": pulumi.String("Value2"),
//								},
//							},
//						},
//						Status: pulumi.String("Enabled"),
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
// ### Specifying a filter based on object size
//
// Object size values are in bytes. Maximum filter size is 5TB. Some storage classes have minimum object size limitations, for more information, see [Comparing the Amazon S3 storage classes](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-class-intro.html#sc-compare).
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
//			_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
//				Bucket: pulumi.Any(bucket.Id),
//				Rules: s3.BucketLifecycleConfigurationV2RuleArray{
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id: pulumi.String("rule-1"),
//						Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
//							ObjectSizeGreaterThan: pulumi.String("500"),
//						},
//						Status: pulumi.String("Enabled"),
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
// ### Specifying a filter based on object size range and prefix
//
// The `objectSizeGreaterThan` must be less than the `objectSizeLessThan`. Notice both the object size range and prefix are wrapped in the `and` configuration block.
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
//			_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
//				Bucket: pulumi.Any(bucket.Id),
//				Rules: s3.BucketLifecycleConfigurationV2RuleArray{
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id: pulumi.String("rule-1"),
//						Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
//							And: &s3.BucketLifecycleConfigurationV2RuleFilterAndArgs{
//								Prefix:                pulumi.String("logs/"),
//								ObjectSizeGreaterThan: pulumi.Int(500),
//								ObjectSizeLessThan:    pulumi.Int(64000),
//							},
//						},
//						Status: pulumi.String("Enabled"),
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
// ### Creating a Lifecycle Configuration for a bucket with versioning
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
//			bucket, err := s3.NewBucketV2(ctx, "bucket", &s3.BucketV2Args{
//				Bucket: pulumi.String("my-bucket"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketAclV2(ctx, "bucket_acl", &s3.BucketAclV2Args{
//				Bucket: bucket.ID(),
//				Acl:    pulumi.String("private"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketLifecycleConfigurationV2(ctx, "bucket-config", &s3.BucketLifecycleConfigurationV2Args{
//				Bucket: bucket.ID(),
//				Rules: s3.BucketLifecycleConfigurationV2RuleArray{
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id: pulumi.String("log"),
//						Expiration: &s3.BucketLifecycleConfigurationV2RuleExpirationArgs{
//							Days: pulumi.Int(90),
//						},
//						Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
//							And: &s3.BucketLifecycleConfigurationV2RuleFilterAndArgs{
//								Prefix: pulumi.String("log/"),
//								Tags: pulumi.StringMap{
//									"rule":      pulumi.String("log"),
//									"autoclean": pulumi.String("true"),
//								},
//							},
//						},
//						Status: pulumi.String("Enabled"),
//						Transitions: s3.BucketLifecycleConfigurationV2RuleTransitionArray{
//							&s3.BucketLifecycleConfigurationV2RuleTransitionArgs{
//								Days:         pulumi.Int(30),
//								StorageClass: pulumi.String("STANDARD_IA"),
//							},
//							&s3.BucketLifecycleConfigurationV2RuleTransitionArgs{
//								Days:         pulumi.Int(60),
//								StorageClass: pulumi.String("GLACIER"),
//							},
//						},
//					},
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id: pulumi.String("tmp"),
//						Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
//							Prefix: pulumi.String("tmp/"),
//						},
//						Expiration: &s3.BucketLifecycleConfigurationV2RuleExpirationArgs{
//							Date: pulumi.String("2023-01-13T00:00:00Z"),
//						},
//						Status: pulumi.String("Enabled"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			versioningBucket, err := s3.NewBucketV2(ctx, "versioning_bucket", &s3.BucketV2Args{
//				Bucket: pulumi.String("my-versioning-bucket"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketAclV2(ctx, "versioning_bucket_acl", &s3.BucketAclV2Args{
//				Bucket: versioningBucket.ID(),
//				Acl:    pulumi.String("private"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketVersioningV2(ctx, "versioning", &s3.BucketVersioningV2Args{
//				Bucket: versioningBucket.ID(),
//				VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
//					Status: pulumi.String("Enabled"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketLifecycleConfigurationV2(ctx, "versioning-bucket-config", &s3.BucketLifecycleConfigurationV2Args{
//				Bucket: versioningBucket.ID(),
//				Rules: s3.BucketLifecycleConfigurationV2RuleArray{
//					&s3.BucketLifecycleConfigurationV2RuleArgs{
//						Id: pulumi.String("config"),
//						Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
//							Prefix: pulumi.String("config/"),
//						},
//						NoncurrentVersionExpiration: &s3.BucketLifecycleConfigurationV2RuleNoncurrentVersionExpirationArgs{
//							NoncurrentDays: pulumi.Int(90),
//						},
//						NoncurrentVersionTransitions: s3.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArray{
//							&s3.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs{
//								NoncurrentDays: pulumi.Int(30),
//								StorageClass:   pulumi.String("STANDARD_IA"),
//							},
//							&s3.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs{
//								NoncurrentDays: pulumi.Int(60),
//								StorageClass:   pulumi.String("GLACIER"),
//							},
//						},
//						Status: pulumi.String("Enabled"),
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
// __Using `pulumi import` to import__ S3 bucket lifecycle configuration using the `bucket` or using the `bucket` and `expected_bucket_owner` separated by a comma (`,`). For example:
//
// If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, import using the `bucket`:
//
// ```sh
//
//	$ pulumi import aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2 example bucket-name
//
// ```
//
//	If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):
//
// ```sh
//
//	$ pulumi import aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2 example bucket-name,123456789012
//
// ```
type BucketLifecycleConfigurationV2 struct {
	pulumi.CustomResourceState

	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner pulumi.StringPtrOutput `pulumi:"expectedBucketOwner"`
	// List of configuration blocks describing the rules managing the replication. See below.
	Rules BucketLifecycleConfigurationV2RuleArrayOutput `pulumi:"rules"`
}

// NewBucketLifecycleConfigurationV2 registers a new resource with the given unique name, arguments, and options.
func NewBucketLifecycleConfigurationV2(ctx *pulumi.Context,
	name string, args *BucketLifecycleConfigurationV2Args, opts ...pulumi.ResourceOption) (*BucketLifecycleConfigurationV2, error) {
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
	var resource BucketLifecycleConfigurationV2
	err := ctx.RegisterResource("aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketLifecycleConfigurationV2 gets an existing BucketLifecycleConfigurationV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketLifecycleConfigurationV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketLifecycleConfigurationV2State, opts ...pulumi.ResourceOption) (*BucketLifecycleConfigurationV2, error) {
	var resource BucketLifecycleConfigurationV2
	err := ctx.ReadResource("aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketLifecycleConfigurationV2 resources.
type bucketLifecycleConfigurationV2State struct {
	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket *string `pulumi:"bucket"`
	// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// List of configuration blocks describing the rules managing the replication. See below.
	Rules []BucketLifecycleConfigurationV2Rule `pulumi:"rules"`
}

type BucketLifecycleConfigurationV2State struct {
	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket pulumi.StringPtrInput
	// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner pulumi.StringPtrInput
	// List of configuration blocks describing the rules managing the replication. See below.
	Rules BucketLifecycleConfigurationV2RuleArrayInput
}

func (BucketLifecycleConfigurationV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketLifecycleConfigurationV2State)(nil)).Elem()
}

type bucketLifecycleConfigurationV2Args struct {
	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket string `pulumi:"bucket"`
	// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// List of configuration blocks describing the rules managing the replication. See below.
	Rules []BucketLifecycleConfigurationV2Rule `pulumi:"rules"`
}

// The set of arguments for constructing a BucketLifecycleConfigurationV2 resource.
type BucketLifecycleConfigurationV2Args struct {
	// Name of the source S3 bucket you want Amazon S3 to monitor.
	Bucket pulumi.StringInput
	// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner pulumi.StringPtrInput
	// List of configuration blocks describing the rules managing the replication. See below.
	Rules BucketLifecycleConfigurationV2RuleArrayInput
}

func (BucketLifecycleConfigurationV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketLifecycleConfigurationV2Args)(nil)).Elem()
}

type BucketLifecycleConfigurationV2Input interface {
	pulumi.Input

	ToBucketLifecycleConfigurationV2Output() BucketLifecycleConfigurationV2Output
	ToBucketLifecycleConfigurationV2OutputWithContext(ctx context.Context) BucketLifecycleConfigurationV2Output
}

func (*BucketLifecycleConfigurationV2) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketLifecycleConfigurationV2)(nil)).Elem()
}

func (i *BucketLifecycleConfigurationV2) ToBucketLifecycleConfigurationV2Output() BucketLifecycleConfigurationV2Output {
	return i.ToBucketLifecycleConfigurationV2OutputWithContext(context.Background())
}

func (i *BucketLifecycleConfigurationV2) ToBucketLifecycleConfigurationV2OutputWithContext(ctx context.Context) BucketLifecycleConfigurationV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationV2Output)
}

// BucketLifecycleConfigurationV2ArrayInput is an input type that accepts BucketLifecycleConfigurationV2Array and BucketLifecycleConfigurationV2ArrayOutput values.
// You can construct a concrete instance of `BucketLifecycleConfigurationV2ArrayInput` via:
//
//	BucketLifecycleConfigurationV2Array{ BucketLifecycleConfigurationV2Args{...} }
type BucketLifecycleConfigurationV2ArrayInput interface {
	pulumi.Input

	ToBucketLifecycleConfigurationV2ArrayOutput() BucketLifecycleConfigurationV2ArrayOutput
	ToBucketLifecycleConfigurationV2ArrayOutputWithContext(context.Context) BucketLifecycleConfigurationV2ArrayOutput
}

type BucketLifecycleConfigurationV2Array []BucketLifecycleConfigurationV2Input

func (BucketLifecycleConfigurationV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketLifecycleConfigurationV2)(nil)).Elem()
}

func (i BucketLifecycleConfigurationV2Array) ToBucketLifecycleConfigurationV2ArrayOutput() BucketLifecycleConfigurationV2ArrayOutput {
	return i.ToBucketLifecycleConfigurationV2ArrayOutputWithContext(context.Background())
}

func (i BucketLifecycleConfigurationV2Array) ToBucketLifecycleConfigurationV2ArrayOutputWithContext(ctx context.Context) BucketLifecycleConfigurationV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationV2ArrayOutput)
}

// BucketLifecycleConfigurationV2MapInput is an input type that accepts BucketLifecycleConfigurationV2Map and BucketLifecycleConfigurationV2MapOutput values.
// You can construct a concrete instance of `BucketLifecycleConfigurationV2MapInput` via:
//
//	BucketLifecycleConfigurationV2Map{ "key": BucketLifecycleConfigurationV2Args{...} }
type BucketLifecycleConfigurationV2MapInput interface {
	pulumi.Input

	ToBucketLifecycleConfigurationV2MapOutput() BucketLifecycleConfigurationV2MapOutput
	ToBucketLifecycleConfigurationV2MapOutputWithContext(context.Context) BucketLifecycleConfigurationV2MapOutput
}

type BucketLifecycleConfigurationV2Map map[string]BucketLifecycleConfigurationV2Input

func (BucketLifecycleConfigurationV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketLifecycleConfigurationV2)(nil)).Elem()
}

func (i BucketLifecycleConfigurationV2Map) ToBucketLifecycleConfigurationV2MapOutput() BucketLifecycleConfigurationV2MapOutput {
	return i.ToBucketLifecycleConfigurationV2MapOutputWithContext(context.Background())
}

func (i BucketLifecycleConfigurationV2Map) ToBucketLifecycleConfigurationV2MapOutputWithContext(ctx context.Context) BucketLifecycleConfigurationV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationV2MapOutput)
}

type BucketLifecycleConfigurationV2Output struct{ *pulumi.OutputState }

func (BucketLifecycleConfigurationV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketLifecycleConfigurationV2)(nil)).Elem()
}

func (o BucketLifecycleConfigurationV2Output) ToBucketLifecycleConfigurationV2Output() BucketLifecycleConfigurationV2Output {
	return o
}

func (o BucketLifecycleConfigurationV2Output) ToBucketLifecycleConfigurationV2OutputWithContext(ctx context.Context) BucketLifecycleConfigurationV2Output {
	return o
}

// Name of the source S3 bucket you want Amazon S3 to monitor.
func (o BucketLifecycleConfigurationV2Output) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketLifecycleConfigurationV2) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP 403 (Access Denied) error.
func (o BucketLifecycleConfigurationV2Output) ExpectedBucketOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketLifecycleConfigurationV2) pulumi.StringPtrOutput { return v.ExpectedBucketOwner }).(pulumi.StringPtrOutput)
}

// List of configuration blocks describing the rules managing the replication. See below.
func (o BucketLifecycleConfigurationV2Output) Rules() BucketLifecycleConfigurationV2RuleArrayOutput {
	return o.ApplyT(func(v *BucketLifecycleConfigurationV2) BucketLifecycleConfigurationV2RuleArrayOutput { return v.Rules }).(BucketLifecycleConfigurationV2RuleArrayOutput)
}

type BucketLifecycleConfigurationV2ArrayOutput struct{ *pulumi.OutputState }

func (BucketLifecycleConfigurationV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketLifecycleConfigurationV2)(nil)).Elem()
}

func (o BucketLifecycleConfigurationV2ArrayOutput) ToBucketLifecycleConfigurationV2ArrayOutput() BucketLifecycleConfigurationV2ArrayOutput {
	return o
}

func (o BucketLifecycleConfigurationV2ArrayOutput) ToBucketLifecycleConfigurationV2ArrayOutputWithContext(ctx context.Context) BucketLifecycleConfigurationV2ArrayOutput {
	return o
}

func (o BucketLifecycleConfigurationV2ArrayOutput) Index(i pulumi.IntInput) BucketLifecycleConfigurationV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketLifecycleConfigurationV2 {
		return vs[0].([]*BucketLifecycleConfigurationV2)[vs[1].(int)]
	}).(BucketLifecycleConfigurationV2Output)
}

type BucketLifecycleConfigurationV2MapOutput struct{ *pulumi.OutputState }

func (BucketLifecycleConfigurationV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketLifecycleConfigurationV2)(nil)).Elem()
}

func (o BucketLifecycleConfigurationV2MapOutput) ToBucketLifecycleConfigurationV2MapOutput() BucketLifecycleConfigurationV2MapOutput {
	return o
}

func (o BucketLifecycleConfigurationV2MapOutput) ToBucketLifecycleConfigurationV2MapOutputWithContext(ctx context.Context) BucketLifecycleConfigurationV2MapOutput {
	return o
}

func (o BucketLifecycleConfigurationV2MapOutput) MapIndex(k pulumi.StringInput) BucketLifecycleConfigurationV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketLifecycleConfigurationV2 {
		return vs[0].(map[string]*BucketLifecycleConfigurationV2)[vs[1].(string)]
	}).(BucketLifecycleConfigurationV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketLifecycleConfigurationV2Input)(nil)).Elem(), &BucketLifecycleConfigurationV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketLifecycleConfigurationV2ArrayInput)(nil)).Elem(), BucketLifecycleConfigurationV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketLifecycleConfigurationV2MapInput)(nil)).Elem(), BucketLifecycleConfigurationV2Map{})
	pulumi.RegisterOutputType(BucketLifecycleConfigurationV2Output{})
	pulumi.RegisterOutputType(BucketLifecycleConfigurationV2ArrayOutput{})
	pulumi.RegisterOutputType(BucketLifecycleConfigurationV2MapOutput{})
}
