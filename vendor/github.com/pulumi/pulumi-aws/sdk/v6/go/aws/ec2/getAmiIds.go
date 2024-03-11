// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of AMI IDs matching the specified criteria.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.GetAmiIds(ctx, &ec2.GetAmiIdsArgs{
//				Owners: []string{
//					"099720109477",
//				},
//				Filters: []ec2.GetAmiIdsFilter{
//					{
//						Name: "name",
//						Values: []string{
//							"ubuntu/images/ubuntu-*-*-amd64-server-*",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAmiIds(ctx *pulumi.Context, args *GetAmiIdsArgs, opts ...pulumi.InvokeOption) (*GetAmiIdsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAmiIdsResult
	err := ctx.Invoke("aws:ec2/getAmiIds:getAmiIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAmiIds.
type GetAmiIdsArgs struct {
	// Limit search to users with *explicit* launch
	// permission on  the image. Valid items are the numeric account ID or `self`.
	ExecutableUsers []string `pulumi:"executableUsers"`
	// One or more name/value pairs to filter off of. There
	// are several valid keys, for a full reference, check out
	// [describe-images in the AWS CLI reference][1].
	Filters []GetAmiIdsFilter `pulumi:"filters"`
	// If true, all deprecated AMIs are included in the response.
	// If false, no deprecated AMIs are included in the response. If no value is specified, the default value is `false`.
	IncludeDeprecated *bool `pulumi:"includeDeprecated"`
	// Regex string to apply to the AMI list returned
	// by AWS. This allows more advanced filtering not supported from the AWS API.
	// This filtering is done locally on what AWS returns, and could have a performance
	// impact if the result is large. Combine this with other
	// options to narrow down the list AWS returns.
	NameRegex *string `pulumi:"nameRegex"`
	// List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g., `amazon`, `aws-marketplace`, `microsoft`).
	Owners []string `pulumi:"owners"`
	// Used to sort AMIs by creation time.
	// If no value is specified, the default value is `false`.
	SortAscending *bool `pulumi:"sortAscending"`
}

// A collection of values returned by getAmiIds.
type GetAmiIdsResult struct {
	ExecutableUsers []string          `pulumi:"executableUsers"`
	Filters         []GetAmiIdsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                string   `pulumi:"id"`
	Ids               []string `pulumi:"ids"`
	IncludeDeprecated *bool    `pulumi:"includeDeprecated"`
	NameRegex         *string  `pulumi:"nameRegex"`
	Owners            []string `pulumi:"owners"`
	SortAscending     *bool    `pulumi:"sortAscending"`
}

func GetAmiIdsOutput(ctx *pulumi.Context, args GetAmiIdsOutputArgs, opts ...pulumi.InvokeOption) GetAmiIdsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAmiIdsResult, error) {
			args := v.(GetAmiIdsArgs)
			r, err := GetAmiIds(ctx, &args, opts...)
			var s GetAmiIdsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAmiIdsResultOutput)
}

// A collection of arguments for invoking getAmiIds.
type GetAmiIdsOutputArgs struct {
	// Limit search to users with *explicit* launch
	// permission on  the image. Valid items are the numeric account ID or `self`.
	ExecutableUsers pulumi.StringArrayInput `pulumi:"executableUsers"`
	// One or more name/value pairs to filter off of. There
	// are several valid keys, for a full reference, check out
	// [describe-images in the AWS CLI reference][1].
	Filters GetAmiIdsFilterArrayInput `pulumi:"filters"`
	// If true, all deprecated AMIs are included in the response.
	// If false, no deprecated AMIs are included in the response. If no value is specified, the default value is `false`.
	IncludeDeprecated pulumi.BoolPtrInput `pulumi:"includeDeprecated"`
	// Regex string to apply to the AMI list returned
	// by AWS. This allows more advanced filtering not supported from the AWS API.
	// This filtering is done locally on what AWS returns, and could have a performance
	// impact if the result is large. Combine this with other
	// options to narrow down the list AWS returns.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g., `amazon`, `aws-marketplace`, `microsoft`).
	Owners pulumi.StringArrayInput `pulumi:"owners"`
	// Used to sort AMIs by creation time.
	// If no value is specified, the default value is `false`.
	SortAscending pulumi.BoolPtrInput `pulumi:"sortAscending"`
}

func (GetAmiIdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiIdsArgs)(nil)).Elem()
}

// A collection of values returned by getAmiIds.
type GetAmiIdsResultOutput struct{ *pulumi.OutputState }

func (GetAmiIdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiIdsResult)(nil)).Elem()
}

func (o GetAmiIdsResultOutput) ToGetAmiIdsResultOutput() GetAmiIdsResultOutput {
	return o
}

func (o GetAmiIdsResultOutput) ToGetAmiIdsResultOutputWithContext(ctx context.Context) GetAmiIdsResultOutput {
	return o
}

func (o GetAmiIdsResultOutput) ExecutableUsers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAmiIdsResult) []string { return v.ExecutableUsers }).(pulumi.StringArrayOutput)
}

func (o GetAmiIdsResultOutput) Filters() GetAmiIdsFilterArrayOutput {
	return o.ApplyT(func(v GetAmiIdsResult) []GetAmiIdsFilter { return v.Filters }).(GetAmiIdsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAmiIdsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiIdsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAmiIdsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAmiIdsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetAmiIdsResultOutput) IncludeDeprecated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAmiIdsResult) *bool { return v.IncludeDeprecated }).(pulumi.BoolPtrOutput)
}

func (o GetAmiIdsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAmiIdsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAmiIdsResultOutput) Owners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAmiIdsResult) []string { return v.Owners }).(pulumi.StringArrayOutput)
}

func (o GetAmiIdsResultOutput) SortAscending() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAmiIdsResult) *bool { return v.SortAscending }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAmiIdsResultOutput{})
}
