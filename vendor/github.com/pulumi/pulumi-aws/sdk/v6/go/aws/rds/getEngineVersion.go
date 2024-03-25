// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Information about an RDS engine version.
//
// ## Example Usage
//
// ### Basic Usage
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
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds.GetEngineVersion(ctx, &rds.GetEngineVersionArgs{
//				Engine: "mysql",
//				PreferredVersions: []string{
//					"8.0.27",
//					"8.0.26",
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
// <!--End PulumiCodeChooser -->
//
// ### With `filter`
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
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds.GetEngineVersion(ctx, &rds.GetEngineVersionArgs{
//				Engine:     "aurora-postgresql",
//				Version:    pulumi.StringRef("10.14"),
//				IncludeAll: pulumi.BoolRef(true),
//				Filters: []rds.GetEngineVersionFilter{
//					{
//						Name: "engine-mode",
//						Values: []string{
//							"serverless",
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
// <!--End PulumiCodeChooser -->
func GetEngineVersion(ctx *pulumi.Context, args *GetEngineVersionArgs, opts ...pulumi.InvokeOption) (*GetEngineVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEngineVersionResult
	err := ctx.Invoke("aws:rds/getEngineVersion:getEngineVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEngineVersion.
type GetEngineVersionArgs struct {
	// Whether the engine version must be an AWS-defined default version. Some engines have multiple default versions, such as for each major version. Using `defaultOnly` may help avoid `multiple RDS engine versions` errors. See also `latest`.
	DefaultOnly *bool `pulumi:"defaultOnly"`
	// Database engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
	//
	// The following arguments are optional:
	Engine string `pulumi:"engine"`
	// One or more name/value pairs to use in filtering versions. There are several valid keys; for a full reference, check out [describe-db-engine-versions in the AWS CLI reference](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-engine-versions.html).
	Filters []GetEngineVersionFilter `pulumi:"filters"`
	// Whether the engine version must have one or more major upgrade targets. Not including `hasMajorTarget` or setting it to `false` doesn't imply that there's no corresponding major upgrade target for the engine version.
	HasMajorTarget *bool `pulumi:"hasMajorTarget"`
	// Whether the engine version must have one or more minor upgrade targets. Not including `hasMinorTarget` or setting it to `false` doesn't imply that there's no corresponding minor upgrade target for the engine version.
	HasMinorTarget *bool `pulumi:"hasMinorTarget"`
	// Whether the engine version `status` can either be `deprecated` or `available`. When not set or set to `false`, the engine version `status` will always be `available`.
	IncludeAll *bool `pulumi:"includeAll"`
	// Whether the engine version is the most recent version matching the other criteria. This is different from `defaultOnly` in important ways: "default" relies on AWS-defined defaults, the latest version isn't always the default, and AWS might have multiple default versions for an engine. As a result, `defaultOnly` might not prevent errors from `multiple RDS engine versions`, while `latest` will. (`latest` can be used with `defaultOnly`.) **Note:** The data source uses a best-effort approach at selecting the latest version. Due to the complexity of version identifiers across engines and incomplete version date information provided by AWS, using `latest` may not always result in the engine version being the actual latest version.
	Latest *bool `pulumi:"latest"`
	// Name of a specific database parameter group family. Examples of parameter group families are `mysql8.0`, `mariadb10.4`, and `postgres12`.
	ParameterGroupFamily *string `pulumi:"parameterGroupFamily"`
	// Ordered list of preferred major version upgrade targets. The engine version will be the first match in the list unless the `latest` parameter is set to `true`. The engine version will be the default version if you don't include any criteria, such as `preferredMajorTargets`.
	PreferredMajorTargets []string `pulumi:"preferredMajorTargets"`
	// Ordered list of preferred version upgrade targets. The engine version will be the first match in this list unless the `latest` parameter is set to `true`. The engine version will be the default version if you don't include any criteria, such as `preferredUpgradeTargets`.
	PreferredUpgradeTargets []string `pulumi:"preferredUpgradeTargets"`
	// Ordered list of preferred versions. The engine version will be the first match in this list unless the `latest` parameter is set to `true`. The engine version will be the default version if you don't include any criteria, such as `preferredVersions`.
	PreferredVersions []string `pulumi:"preferredVersions"`
	Version           *string  `pulumi:"version"`
}

// A collection of values returned by getEngineVersion.
type GetEngineVersionResult struct {
	// Default character set for new instances of the engine version.
	DefaultCharacterSet string `pulumi:"defaultCharacterSet"`
	DefaultOnly         *bool  `pulumi:"defaultOnly"`
	Engine              string `pulumi:"engine"`
	// Description of the engine.
	EngineDescription string `pulumi:"engineDescription"`
	// Set of log types that the engine version has available for export to CloudWatch Logs.
	ExportableLogTypes []string                 `pulumi:"exportableLogTypes"`
	Filters            []GetEngineVersionFilter `pulumi:"filters"`
	HasMajorTarget     *bool                    `pulumi:"hasMajorTarget"`
	HasMinorTarget     *bool                    `pulumi:"hasMinorTarget"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string   `pulumi:"id"`
	IncludeAll              *bool    `pulumi:"includeAll"`
	Latest                  *bool    `pulumi:"latest"`
	ParameterGroupFamily    string   `pulumi:"parameterGroupFamily"`
	PreferredMajorTargets   []string `pulumi:"preferredMajorTargets"`
	PreferredUpgradeTargets []string `pulumi:"preferredUpgradeTargets"`
	PreferredVersions       []string `pulumi:"preferredVersions"`
	// Status of the engine version, either `available` or `deprecated`.
	Status string `pulumi:"status"`
	// Set of character sets supported by th engine version.
	SupportedCharacterSets []string `pulumi:"supportedCharacterSets"`
	// Set of features supported by the engine version.
	SupportedFeatureNames []string `pulumi:"supportedFeatureNames"`
	// Set of supported engine version modes.
	SupportedModes []string `pulumi:"supportedModes"`
	// Set of the time zones supported by the engine version.
	SupportedTimezones []string `pulumi:"supportedTimezones"`
	// Whether you can use Aurora global databases with the engine version.
	SupportsGlobalDatabases bool `pulumi:"supportsGlobalDatabases"`
	// Whether the engine version supports exporting the log types specified by `exportableLogTypes` to CloudWatch Logs.
	SupportsLogExportsToCloudwatch bool `pulumi:"supportsLogExportsToCloudwatch"`
	// Whether you can use Aurora parallel query with the engine version.
	SupportsParallelQuery bool `pulumi:"supportsParallelQuery"`
	// Whether the engine version supports read replicas.
	SupportsReadReplica bool `pulumi:"supportsReadReplica"`
	// Set of versions that are valid major version upgrades for the engine version.
	ValidMajorTargets []string `pulumi:"validMajorTargets"`
	// Set of versions that are valid minor version upgrades for the engine version.
	ValidMinorTargets []string `pulumi:"validMinorTargets"`
	// Set of versions that are valid major or minor upgrades for the engine version.
	ValidUpgradeTargets []string `pulumi:"validUpgradeTargets"`
	Version             string   `pulumi:"version"`
	// Complete engine version.
	VersionActual string `pulumi:"versionActual"`
	// Description of the engine version.
	VersionDescription string `pulumi:"versionDescription"`
}

func GetEngineVersionOutput(ctx *pulumi.Context, args GetEngineVersionOutputArgs, opts ...pulumi.InvokeOption) GetEngineVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEngineVersionResult, error) {
			args := v.(GetEngineVersionArgs)
			r, err := GetEngineVersion(ctx, &args, opts...)
			var s GetEngineVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEngineVersionResultOutput)
}

// A collection of arguments for invoking getEngineVersion.
type GetEngineVersionOutputArgs struct {
	// Whether the engine version must be an AWS-defined default version. Some engines have multiple default versions, such as for each major version. Using `defaultOnly` may help avoid `multiple RDS engine versions` errors. See also `latest`.
	DefaultOnly pulumi.BoolPtrInput `pulumi:"defaultOnly"`
	// Database engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
	//
	// The following arguments are optional:
	Engine pulumi.StringInput `pulumi:"engine"`
	// One or more name/value pairs to use in filtering versions. There are several valid keys; for a full reference, check out [describe-db-engine-versions in the AWS CLI reference](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-engine-versions.html).
	Filters GetEngineVersionFilterArrayInput `pulumi:"filters"`
	// Whether the engine version must have one or more major upgrade targets. Not including `hasMajorTarget` or setting it to `false` doesn't imply that there's no corresponding major upgrade target for the engine version.
	HasMajorTarget pulumi.BoolPtrInput `pulumi:"hasMajorTarget"`
	// Whether the engine version must have one or more minor upgrade targets. Not including `hasMinorTarget` or setting it to `false` doesn't imply that there's no corresponding minor upgrade target for the engine version.
	HasMinorTarget pulumi.BoolPtrInput `pulumi:"hasMinorTarget"`
	// Whether the engine version `status` can either be `deprecated` or `available`. When not set or set to `false`, the engine version `status` will always be `available`.
	IncludeAll pulumi.BoolPtrInput `pulumi:"includeAll"`
	// Whether the engine version is the most recent version matching the other criteria. This is different from `defaultOnly` in important ways: "default" relies on AWS-defined defaults, the latest version isn't always the default, and AWS might have multiple default versions for an engine. As a result, `defaultOnly` might not prevent errors from `multiple RDS engine versions`, while `latest` will. (`latest` can be used with `defaultOnly`.) **Note:** The data source uses a best-effort approach at selecting the latest version. Due to the complexity of version identifiers across engines and incomplete version date information provided by AWS, using `latest` may not always result in the engine version being the actual latest version.
	Latest pulumi.BoolPtrInput `pulumi:"latest"`
	// Name of a specific database parameter group family. Examples of parameter group families are `mysql8.0`, `mariadb10.4`, and `postgres12`.
	ParameterGroupFamily pulumi.StringPtrInput `pulumi:"parameterGroupFamily"`
	// Ordered list of preferred major version upgrade targets. The engine version will be the first match in the list unless the `latest` parameter is set to `true`. The engine version will be the default version if you don't include any criteria, such as `preferredMajorTargets`.
	PreferredMajorTargets pulumi.StringArrayInput `pulumi:"preferredMajorTargets"`
	// Ordered list of preferred version upgrade targets. The engine version will be the first match in this list unless the `latest` parameter is set to `true`. The engine version will be the default version if you don't include any criteria, such as `preferredUpgradeTargets`.
	PreferredUpgradeTargets pulumi.StringArrayInput `pulumi:"preferredUpgradeTargets"`
	// Ordered list of preferred versions. The engine version will be the first match in this list unless the `latest` parameter is set to `true`. The engine version will be the default version if you don't include any criteria, such as `preferredVersions`.
	PreferredVersions pulumi.StringArrayInput `pulumi:"preferredVersions"`
	Version           pulumi.StringPtrInput   `pulumi:"version"`
}

func (GetEngineVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEngineVersionArgs)(nil)).Elem()
}

// A collection of values returned by getEngineVersion.
type GetEngineVersionResultOutput struct{ *pulumi.OutputState }

func (GetEngineVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEngineVersionResult)(nil)).Elem()
}

func (o GetEngineVersionResultOutput) ToGetEngineVersionResultOutput() GetEngineVersionResultOutput {
	return o
}

func (o GetEngineVersionResultOutput) ToGetEngineVersionResultOutputWithContext(ctx context.Context) GetEngineVersionResultOutput {
	return o
}

// Default character set for new instances of the engine version.
func (o GetEngineVersionResultOutput) DefaultCharacterSet() pulumi.StringOutput {
	return o.ApplyT(func(v GetEngineVersionResult) string { return v.DefaultCharacterSet }).(pulumi.StringOutput)
}

func (o GetEngineVersionResultOutput) DefaultOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEngineVersionResult) *bool { return v.DefaultOnly }).(pulumi.BoolPtrOutput)
}

func (o GetEngineVersionResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetEngineVersionResult) string { return v.Engine }).(pulumi.StringOutput)
}

// Description of the engine.
func (o GetEngineVersionResultOutput) EngineDescription() pulumi.StringOutput {
	return o.ApplyT(func(v GetEngineVersionResult) string { return v.EngineDescription }).(pulumi.StringOutput)
}

// Set of log types that the engine version has available for export to CloudWatch Logs.
func (o GetEngineVersionResultOutput) ExportableLogTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []string { return v.ExportableLogTypes }).(pulumi.StringArrayOutput)
}

func (o GetEngineVersionResultOutput) Filters() GetEngineVersionFilterArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []GetEngineVersionFilter { return v.Filters }).(GetEngineVersionFilterArrayOutput)
}

func (o GetEngineVersionResultOutput) HasMajorTarget() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEngineVersionResult) *bool { return v.HasMajorTarget }).(pulumi.BoolPtrOutput)
}

func (o GetEngineVersionResultOutput) HasMinorTarget() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEngineVersionResult) *bool { return v.HasMinorTarget }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEngineVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEngineVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEngineVersionResultOutput) IncludeAll() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEngineVersionResult) *bool { return v.IncludeAll }).(pulumi.BoolPtrOutput)
}

func (o GetEngineVersionResultOutput) Latest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEngineVersionResult) *bool { return v.Latest }).(pulumi.BoolPtrOutput)
}

func (o GetEngineVersionResultOutput) ParameterGroupFamily() pulumi.StringOutput {
	return o.ApplyT(func(v GetEngineVersionResult) string { return v.ParameterGroupFamily }).(pulumi.StringOutput)
}

func (o GetEngineVersionResultOutput) PreferredMajorTargets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []string { return v.PreferredMajorTargets }).(pulumi.StringArrayOutput)
}

func (o GetEngineVersionResultOutput) PreferredUpgradeTargets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []string { return v.PreferredUpgradeTargets }).(pulumi.StringArrayOutput)
}

func (o GetEngineVersionResultOutput) PreferredVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []string { return v.PreferredVersions }).(pulumi.StringArrayOutput)
}

// Status of the engine version, either `available` or `deprecated`.
func (o GetEngineVersionResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetEngineVersionResult) string { return v.Status }).(pulumi.StringOutput)
}

// Set of character sets supported by th engine version.
func (o GetEngineVersionResultOutput) SupportedCharacterSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []string { return v.SupportedCharacterSets }).(pulumi.StringArrayOutput)
}

// Set of features supported by the engine version.
func (o GetEngineVersionResultOutput) SupportedFeatureNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []string { return v.SupportedFeatureNames }).(pulumi.StringArrayOutput)
}

// Set of supported engine version modes.
func (o GetEngineVersionResultOutput) SupportedModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []string { return v.SupportedModes }).(pulumi.StringArrayOutput)
}

// Set of the time zones supported by the engine version.
func (o GetEngineVersionResultOutput) SupportedTimezones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []string { return v.SupportedTimezones }).(pulumi.StringArrayOutput)
}

// Whether you can use Aurora global databases with the engine version.
func (o GetEngineVersionResultOutput) SupportsGlobalDatabases() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEngineVersionResult) bool { return v.SupportsGlobalDatabases }).(pulumi.BoolOutput)
}

// Whether the engine version supports exporting the log types specified by `exportableLogTypes` to CloudWatch Logs.
func (o GetEngineVersionResultOutput) SupportsLogExportsToCloudwatch() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEngineVersionResult) bool { return v.SupportsLogExportsToCloudwatch }).(pulumi.BoolOutput)
}

// Whether you can use Aurora parallel query with the engine version.
func (o GetEngineVersionResultOutput) SupportsParallelQuery() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEngineVersionResult) bool { return v.SupportsParallelQuery }).(pulumi.BoolOutput)
}

// Whether the engine version supports read replicas.
func (o GetEngineVersionResultOutput) SupportsReadReplica() pulumi.BoolOutput {
	return o.ApplyT(func(v GetEngineVersionResult) bool { return v.SupportsReadReplica }).(pulumi.BoolOutput)
}

// Set of versions that are valid major version upgrades for the engine version.
func (o GetEngineVersionResultOutput) ValidMajorTargets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []string { return v.ValidMajorTargets }).(pulumi.StringArrayOutput)
}

// Set of versions that are valid minor version upgrades for the engine version.
func (o GetEngineVersionResultOutput) ValidMinorTargets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []string { return v.ValidMinorTargets }).(pulumi.StringArrayOutput)
}

// Set of versions that are valid major or minor upgrades for the engine version.
func (o GetEngineVersionResultOutput) ValidUpgradeTargets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEngineVersionResult) []string { return v.ValidUpgradeTargets }).(pulumi.StringArrayOutput)
}

func (o GetEngineVersionResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetEngineVersionResult) string { return v.Version }).(pulumi.StringOutput)
}

// Complete engine version.
func (o GetEngineVersionResultOutput) VersionActual() pulumi.StringOutput {
	return o.ApplyT(func(v GetEngineVersionResult) string { return v.VersionActual }).(pulumi.StringOutput)
}

// Description of the engine version.
func (o GetEngineVersionResultOutput) VersionDescription() pulumi.StringOutput {
	return o.ApplyT(func(v GetEngineVersionResult) string { return v.VersionDescription }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEngineVersionResultOutput{})
}
