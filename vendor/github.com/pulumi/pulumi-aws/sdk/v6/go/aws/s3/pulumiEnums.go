// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// See https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
type CannedAcl string

const (
	CannedAclPrivate                = CannedAcl("private")
	CannedAclPublicRead             = CannedAcl("public-read")
	CannedAclPublicReadWrite        = CannedAcl("public-read-write")
	CannedAclAwsExecRead            = CannedAcl("aws-exec-read")
	CannedAclAuthenticatedRead      = CannedAcl("authenticated-read")
	CannedAclBucketOwnerRead        = CannedAcl("bucket-owner-read")
	CannedAclBucketOwnerFullControl = CannedAcl("bucket-owner-full-control")
	CannedAclLogDeliveryWrite       = CannedAcl("log-delivery-write")
)

func (CannedAcl) ElementType() reflect.Type {
	return reflect.TypeOf((*CannedAcl)(nil)).Elem()
}

func (e CannedAcl) ToCannedAclOutput() CannedAclOutput {
	return pulumi.ToOutput(e).(CannedAclOutput)
}

func (e CannedAcl) ToCannedAclOutputWithContext(ctx context.Context) CannedAclOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CannedAclOutput)
}

func (e CannedAcl) ToCannedAclPtrOutput() CannedAclPtrOutput {
	return e.ToCannedAclPtrOutputWithContext(context.Background())
}

func (e CannedAcl) ToCannedAclPtrOutputWithContext(ctx context.Context) CannedAclPtrOutput {
	return CannedAcl(e).ToCannedAclOutputWithContext(ctx).ToCannedAclPtrOutputWithContext(ctx)
}

func (e CannedAcl) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CannedAcl) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CannedAcl) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CannedAcl) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CannedAclOutput struct{ *pulumi.OutputState }

func (CannedAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CannedAcl)(nil)).Elem()
}

func (o CannedAclOutput) ToCannedAclOutput() CannedAclOutput {
	return o
}

func (o CannedAclOutput) ToCannedAclOutputWithContext(ctx context.Context) CannedAclOutput {
	return o
}

func (o CannedAclOutput) ToCannedAclPtrOutput() CannedAclPtrOutput {
	return o.ToCannedAclPtrOutputWithContext(context.Background())
}

func (o CannedAclOutput) ToCannedAclPtrOutputWithContext(ctx context.Context) CannedAclPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CannedAcl) *CannedAcl {
		return &v
	}).(CannedAclPtrOutput)
}

func (o CannedAclOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CannedAclOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CannedAcl) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CannedAclOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CannedAclOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CannedAcl) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CannedAclPtrOutput struct{ *pulumi.OutputState }

func (CannedAclPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CannedAcl)(nil)).Elem()
}

func (o CannedAclPtrOutput) ToCannedAclPtrOutput() CannedAclPtrOutput {
	return o
}

func (o CannedAclPtrOutput) ToCannedAclPtrOutputWithContext(ctx context.Context) CannedAclPtrOutput {
	return o
}

func (o CannedAclPtrOutput) Elem() CannedAclOutput {
	return o.ApplyT(func(v *CannedAcl) CannedAcl {
		if v != nil {
			return *v
		}
		var ret CannedAcl
		return ret
	}).(CannedAclOutput)
}

func (o CannedAclPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CannedAclPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CannedAcl) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CannedAclInput is an input type that accepts values of the CannedAcl enum
// A concrete instance of `CannedAclInput` can be one of the following:
//
//	CannedAclPrivate
//	CannedAclPublicRead
//	CannedAclPublicReadWrite
//	CannedAclAwsExecRead
//	CannedAclAuthenticatedRead
//	CannedAclBucketOwnerRead
//	CannedAclBucketOwnerFullControl
//	CannedAclLogDeliveryWrite
type CannedAclInput interface {
	pulumi.Input

	ToCannedAclOutput() CannedAclOutput
	ToCannedAclOutputWithContext(context.Context) CannedAclOutput
}

var cannedAclPtrType = reflect.TypeOf((**CannedAcl)(nil)).Elem()

type CannedAclPtrInput interface {
	pulumi.Input

	ToCannedAclPtrOutput() CannedAclPtrOutput
	ToCannedAclPtrOutputWithContext(context.Context) CannedAclPtrOutput
}

type cannedAclPtr string

func CannedAclPtr(v string) CannedAclPtrInput {
	return (*cannedAclPtr)(&v)
}

func (*cannedAclPtr) ElementType() reflect.Type {
	return cannedAclPtrType
}

func (in *cannedAclPtr) ToCannedAclPtrOutput() CannedAclPtrOutput {
	return pulumi.ToOutput(in).(CannedAclPtrOutput)
}

func (in *cannedAclPtr) ToCannedAclPtrOutputWithContext(ctx context.Context) CannedAclPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CannedAclPtrOutput)
}

func (in *cannedAclPtr) ToOutput(ctx context.Context) pulumix.Output[*CannedAcl] {
	return pulumix.Output[*CannedAcl]{
		OutputState: in.ToCannedAclPtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CannedAclInput)(nil)).Elem(), CannedAcl("private"))
	pulumi.RegisterInputType(reflect.TypeOf((*CannedAclPtrInput)(nil)).Elem(), CannedAcl("private"))
	pulumi.RegisterOutputType(CannedAclOutput{})
	pulumi.RegisterOutputType(CannedAclPtrOutput{})
}
