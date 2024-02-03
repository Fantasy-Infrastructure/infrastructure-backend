// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SSM Patch Baseline data source. Useful if you wish to reuse the default baselines provided.
//
// ## Example Usage
//
// To retrieve a baseline provided by AWS:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssm.LookupPatchBaseline(ctx, &ssm.LookupPatchBaselineArgs{
//				NamePrefix:      pulumi.StringRef("AWS-"),
//				OperatingSystem: pulumi.StringRef("CENTOS"),
//				Owner:           "AWS",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// To retrieve a baseline on your account:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssm.LookupPatchBaseline(ctx, &ssm.LookupPatchBaselineArgs{
//				DefaultBaseline: pulumi.BoolRef(true),
//				NamePrefix:      pulumi.StringRef("MyCustomBaseline"),
//				OperatingSystem: pulumi.StringRef("WINDOWS"),
//				Owner:           "Self",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPatchBaseline(ctx *pulumi.Context, args *LookupPatchBaselineArgs, opts ...pulumi.InvokeOption) (*LookupPatchBaselineResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPatchBaselineResult
	err := ctx.Invoke("aws:ssm/getPatchBaseline:getPatchBaseline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPatchBaseline.
type LookupPatchBaselineArgs struct {
	// Filters the results against the baselines defaultBaseline field.
	DefaultBaseline *bool `pulumi:"defaultBaseline"`
	// Filter results by the baseline name prefix.
	NamePrefix *string `pulumi:"namePrefix"`
	// Specified OS for the baseline. Valid values: `AMAZON_LINUX`, `AMAZON_LINUX_2`, `UBUNTU`, `REDHAT_ENTERPRISE_LINUX`, `SUSE`, `CENTOS`, `ORACLE_LINUX`, `DEBIAN`, `MACOS`, `RASPBIAN` and `ROCKY_LINUX`.
	OperatingSystem *string `pulumi:"operatingSystem"`
	// Owner of the baseline. Valid values: `All`, `AWS`, `Self` (the current account).
	Owner string `pulumi:"owner"`
}

// A collection of values returned by getPatchBaseline.
type LookupPatchBaselineResult struct {
	// List of rules used to include patches in the baseline.
	ApprovalRules []GetPatchBaselineApprovalRule `pulumi:"approvalRules"`
	// List of explicitly approved patches for the baseline.
	ApprovedPatches []string `pulumi:"approvedPatches"`
	// The compliance level for approved patches.
	ApprovedPatchesComplianceLevel string `pulumi:"approvedPatchesComplianceLevel"`
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
	ApprovedPatchesEnableNonSecurity bool  `pulumi:"approvedPatchesEnableNonSecurity"`
	DefaultBaseline                  *bool `pulumi:"defaultBaseline"`
	// Description of the baseline.
	Description string `pulumi:"description"`
	// Set of global filters used to exclude patches from the baseline.
	GlobalFilters []GetPatchBaselineGlobalFilter `pulumi:"globalFilters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name specified to identify the patch source.
	Name            string  `pulumi:"name"`
	NamePrefix      *string `pulumi:"namePrefix"`
	OperatingSystem *string `pulumi:"operatingSystem"`
	Owner           string  `pulumi:"owner"`
	// List of rejected patches.
	RejectedPatches []string `pulumi:"rejectedPatches"`
	// The action specified to take on patches included in the `rejectedPatches` list.
	RejectedPatchesAction string `pulumi:"rejectedPatchesAction"`
	// Information about the patches to use to update the managed nodes, including target operating systems and source repositories.
	Sources []GetPatchBaselineSource `pulumi:"sources"`
}

func LookupPatchBaselineOutput(ctx *pulumi.Context, args LookupPatchBaselineOutputArgs, opts ...pulumi.InvokeOption) LookupPatchBaselineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPatchBaselineResult, error) {
			args := v.(LookupPatchBaselineArgs)
			r, err := LookupPatchBaseline(ctx, &args, opts...)
			var s LookupPatchBaselineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPatchBaselineResultOutput)
}

// A collection of arguments for invoking getPatchBaseline.
type LookupPatchBaselineOutputArgs struct {
	// Filters the results against the baselines defaultBaseline field.
	DefaultBaseline pulumi.BoolPtrInput `pulumi:"defaultBaseline"`
	// Filter results by the baseline name prefix.
	NamePrefix pulumi.StringPtrInput `pulumi:"namePrefix"`
	// Specified OS for the baseline. Valid values: `AMAZON_LINUX`, `AMAZON_LINUX_2`, `UBUNTU`, `REDHAT_ENTERPRISE_LINUX`, `SUSE`, `CENTOS`, `ORACLE_LINUX`, `DEBIAN`, `MACOS`, `RASPBIAN` and `ROCKY_LINUX`.
	OperatingSystem pulumi.StringPtrInput `pulumi:"operatingSystem"`
	// Owner of the baseline. Valid values: `All`, `AWS`, `Self` (the current account).
	Owner pulumi.StringInput `pulumi:"owner"`
}

func (LookupPatchBaselineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPatchBaselineArgs)(nil)).Elem()
}

// A collection of values returned by getPatchBaseline.
type LookupPatchBaselineResultOutput struct{ *pulumi.OutputState }

func (LookupPatchBaselineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPatchBaselineResult)(nil)).Elem()
}

func (o LookupPatchBaselineResultOutput) ToLookupPatchBaselineResultOutput() LookupPatchBaselineResultOutput {
	return o
}

func (o LookupPatchBaselineResultOutput) ToLookupPatchBaselineResultOutputWithContext(ctx context.Context) LookupPatchBaselineResultOutput {
	return o
}

// List of rules used to include patches in the baseline.
func (o LookupPatchBaselineResultOutput) ApprovalRules() GetPatchBaselineApprovalRuleArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []GetPatchBaselineApprovalRule { return v.ApprovalRules }).(GetPatchBaselineApprovalRuleArrayOutput)
}

// List of explicitly approved patches for the baseline.
func (o LookupPatchBaselineResultOutput) ApprovedPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []string { return v.ApprovedPatches }).(pulumi.StringArrayOutput)
}

// The compliance level for approved patches.
func (o LookupPatchBaselineResultOutput) ApprovedPatchesComplianceLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) string { return v.ApprovedPatchesComplianceLevel }).(pulumi.StringOutput)
}

// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
func (o LookupPatchBaselineResultOutput) ApprovedPatchesEnableNonSecurity() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) bool { return v.ApprovedPatchesEnableNonSecurity }).(pulumi.BoolOutput)
}

func (o LookupPatchBaselineResultOutput) DefaultBaseline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *bool { return v.DefaultBaseline }).(pulumi.BoolPtrOutput)
}

// Description of the baseline.
func (o LookupPatchBaselineResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) string { return v.Description }).(pulumi.StringOutput)
}

// Set of global filters used to exclude patches from the baseline.
func (o LookupPatchBaselineResultOutput) GlobalFilters() GetPatchBaselineGlobalFilterArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []GetPatchBaselineGlobalFilter { return v.GlobalFilters }).(GetPatchBaselineGlobalFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPatchBaselineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name specified to identify the patch source.
func (o LookupPatchBaselineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPatchBaselineResultOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *string { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

func (o LookupPatchBaselineResultOutput) OperatingSystem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) *string { return v.OperatingSystem }).(pulumi.StringPtrOutput)
}

func (o LookupPatchBaselineResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) string { return v.Owner }).(pulumi.StringOutput)
}

// List of rejected patches.
func (o LookupPatchBaselineResultOutput) RejectedPatches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []string { return v.RejectedPatches }).(pulumi.StringArrayOutput)
}

// The action specified to take on patches included in the `rejectedPatches` list.
func (o LookupPatchBaselineResultOutput) RejectedPatchesAction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) string { return v.RejectedPatchesAction }).(pulumi.StringOutput)
}

// Information about the patches to use to update the managed nodes, including target operating systems and source repositories.
func (o LookupPatchBaselineResultOutput) Sources() GetPatchBaselineSourceArrayOutput {
	return o.ApplyT(func(v LookupPatchBaselineResult) []GetPatchBaselineSource { return v.Sources }).(GetPatchBaselineSourceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPatchBaselineResultOutput{})
}