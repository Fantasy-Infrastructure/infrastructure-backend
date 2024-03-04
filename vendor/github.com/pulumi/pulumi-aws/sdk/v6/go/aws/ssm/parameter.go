// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SSM Parameter resource.
//
// > **Note:** `overwrite` also makes it possible to overwrite an existing SSM Parameter that's not created by the provider before. This argument has been deprecated and will be removed in v6.0.0 of the provider. For more information on how this affects the behavior of this resource, see this issue comment.
//
// ## Example Usage
// ### Basic example
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
//			_, err := ssm.NewParameter(ctx, "foo", &ssm.ParameterArgs{
//				Type:  pulumi.String("String"),
//				Value: pulumi.String("bar"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Encrypted string using default SSM KMS key
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds.NewInstance(ctx, "default", &rds.InstanceArgs{
//				AllocatedStorage:   pulumi.Int(10),
//				StorageType:        pulumi.String("gp2"),
//				Engine:             pulumi.String("mysql"),
//				EngineVersion:      pulumi.String("5.7.16"),
//				InstanceClass:      pulumi.String("db.t2.micro"),
//				DbName:             pulumi.String("mydb"),
//				Username:           pulumi.String("foo"),
//				Password:           pulumi.Any(_var.Database_master_password),
//				DbSubnetGroupName:  pulumi.String("my_database_subnet_group"),
//				ParameterGroupName: pulumi.String("default.mysql5.7"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ssm.NewParameter(ctx, "secret", &ssm.ParameterArgs{
//				Description: pulumi.String("The parameter description"),
//				Type:        pulumi.String("SecureString"),
//				Value:       pulumi.Any(_var.Database_master_password),
//				Tags: pulumi.StringMap{
//					"environment": pulumi.String("production"),
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
// Using `pulumi import`, import SSM Parameters using the parameter store `name`. For example:
//
// ```sh
//
//	$ pulumi import aws:ssm/parameter:Parameter my_param /my_path/my_paramname
//
// ```
type Parameter struct {
	pulumi.CustomResourceState

	// Regular expression used to validate the parameter value.
	AllowedPattern pulumi.StringPtrOutput `pulumi:"allowedPattern"`
	// ARN of the parameter.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Data type of the parameter. Valid values: `text`, `aws:ssm:integration` and `aws:ec2:image` for AMI format, see the [Native parameter support for Amazon Machine Image IDs](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html).
	DataType pulumi.StringOutput `pulumi:"dataType"`
	// Description of the parameter.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Value of the parameter. **Use caution:** This value is _never_ marked as sensitive in the pulumi preview output. This argument is not valid with a `type` of `SecureString`.
	InsecureValue pulumi.StringOutput `pulumi:"insecureValue"`
	// KMS key ID or ARN for encrypting a SecureString.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// Name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name pulumi.StringOutput `pulumi:"name"`
	// Overwrite an existing parameter. If not specified, defaults to `false` if the resource has not been created by Pulumi to avoid overwrite of existing resource, and will default to `true` otherwise (Pulumi lifecycle rules should then be used to manage the update behavior).
	//
	// Deprecated: this attribute has been deprecated
	Overwrite pulumi.BoolPtrOutput `pulumi:"overwrite"`
	// Map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. Downgrading an `Advanced` tier parameter to `Standard` will recreate the resource. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier pulumi.StringPtrOutput `pulumi:"tier"`
	// Type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	//
	// The following arguments are optional:
	Type pulumi.StringOutput `pulumi:"type"`
	// Value of the parameter. This value is always marked as sensitive in the pulumi preview output, regardless of `type`.
	//
	// > **NOTE:** `aws:ssm:integration` dataType parameters must be of the type `SecureString` and the name must start with the prefix `/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/`. See [here](https://docs.aws.amazon.com/systems-manager/latest/userguide/creating-integrations.html) for information on the usage of `aws:ssm:integration` parameters.
	Value pulumi.StringOutput `pulumi:"value"`
	// Version of the parameter.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewParameter registers a new resource with the given unique name, arguments, and options.
func NewParameter(ctx *pulumi.Context,
	name string, args *ParameterArgs, opts ...pulumi.ResourceOption) (*Parameter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value != nil {
		args.Value = pulumi.ToSecret(args.Value).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"value",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Parameter
	err := ctx.RegisterResource("aws:ssm/parameter:Parameter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameter gets an existing Parameter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterState, opts ...pulumi.ResourceOption) (*Parameter, error) {
	var resource Parameter
	err := ctx.ReadResource("aws:ssm/parameter:Parameter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Parameter resources.
type parameterState struct {
	// Regular expression used to validate the parameter value.
	AllowedPattern *string `pulumi:"allowedPattern"`
	// ARN of the parameter.
	Arn *string `pulumi:"arn"`
	// Data type of the parameter. Valid values: `text`, `aws:ssm:integration` and `aws:ec2:image` for AMI format, see the [Native parameter support for Amazon Machine Image IDs](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html).
	DataType *string `pulumi:"dataType"`
	// Description of the parameter.
	Description *string `pulumi:"description"`
	// Value of the parameter. **Use caution:** This value is _never_ marked as sensitive in the pulumi preview output. This argument is not valid with a `type` of `SecureString`.
	InsecureValue *string `pulumi:"insecureValue"`
	// KMS key ID or ARN for encrypting a SecureString.
	KeyId *string `pulumi:"keyId"`
	// Name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name *string `pulumi:"name"`
	// Overwrite an existing parameter. If not specified, defaults to `false` if the resource has not been created by Pulumi to avoid overwrite of existing resource, and will default to `true` otherwise (Pulumi lifecycle rules should then be used to manage the update behavior).
	//
	// Deprecated: this attribute has been deprecated
	Overwrite *bool `pulumi:"overwrite"`
	// Map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. Downgrading an `Advanced` tier parameter to `Standard` will recreate the resource. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier *string `pulumi:"tier"`
	// Type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	//
	// The following arguments are optional:
	Type *string `pulumi:"type"`
	// Value of the parameter. This value is always marked as sensitive in the pulumi preview output, regardless of `type`.
	//
	// > **NOTE:** `aws:ssm:integration` dataType parameters must be of the type `SecureString` and the name must start with the prefix `/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/`. See [here](https://docs.aws.amazon.com/systems-manager/latest/userguide/creating-integrations.html) for information on the usage of `aws:ssm:integration` parameters.
	Value *string `pulumi:"value"`
	// Version of the parameter.
	Version *int `pulumi:"version"`
}

type ParameterState struct {
	// Regular expression used to validate the parameter value.
	AllowedPattern pulumi.StringPtrInput
	// ARN of the parameter.
	Arn pulumi.StringPtrInput
	// Data type of the parameter. Valid values: `text`, `aws:ssm:integration` and `aws:ec2:image` for AMI format, see the [Native parameter support for Amazon Machine Image IDs](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html).
	DataType pulumi.StringPtrInput
	// Description of the parameter.
	Description pulumi.StringPtrInput
	// Value of the parameter. **Use caution:** This value is _never_ marked as sensitive in the pulumi preview output. This argument is not valid with a `type` of `SecureString`.
	InsecureValue pulumi.StringPtrInput
	// KMS key ID or ARN for encrypting a SecureString.
	KeyId pulumi.StringPtrInput
	// Name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name pulumi.StringPtrInput
	// Overwrite an existing parameter. If not specified, defaults to `false` if the resource has not been created by Pulumi to avoid overwrite of existing resource, and will default to `true` otherwise (Pulumi lifecycle rules should then be used to manage the update behavior).
	//
	// Deprecated: this attribute has been deprecated
	Overwrite pulumi.BoolPtrInput
	// Map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. Downgrading an `Advanced` tier parameter to `Standard` will recreate the resource. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier pulumi.StringPtrInput
	// Type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	//
	// The following arguments are optional:
	Type pulumi.StringPtrInput
	// Value of the parameter. This value is always marked as sensitive in the pulumi preview output, regardless of `type`.
	//
	// > **NOTE:** `aws:ssm:integration` dataType parameters must be of the type `SecureString` and the name must start with the prefix `/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/`. See [here](https://docs.aws.amazon.com/systems-manager/latest/userguide/creating-integrations.html) for information on the usage of `aws:ssm:integration` parameters.
	Value pulumi.StringPtrInput
	// Version of the parameter.
	Version pulumi.IntPtrInput
}

func (ParameterState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterState)(nil)).Elem()
}

type parameterArgs struct {
	// Regular expression used to validate the parameter value.
	AllowedPattern *string `pulumi:"allowedPattern"`
	// ARN of the parameter.
	Arn *string `pulumi:"arn"`
	// Data type of the parameter. Valid values: `text`, `aws:ssm:integration` and `aws:ec2:image` for AMI format, see the [Native parameter support for Amazon Machine Image IDs](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html).
	DataType *string `pulumi:"dataType"`
	// Description of the parameter.
	Description *string `pulumi:"description"`
	// Value of the parameter. **Use caution:** This value is _never_ marked as sensitive in the pulumi preview output. This argument is not valid with a `type` of `SecureString`.
	InsecureValue *string `pulumi:"insecureValue"`
	// KMS key ID or ARN for encrypting a SecureString.
	KeyId *string `pulumi:"keyId"`
	// Name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name *string `pulumi:"name"`
	// Overwrite an existing parameter. If not specified, defaults to `false` if the resource has not been created by Pulumi to avoid overwrite of existing resource, and will default to `true` otherwise (Pulumi lifecycle rules should then be used to manage the update behavior).
	//
	// Deprecated: this attribute has been deprecated
	Overwrite *bool `pulumi:"overwrite"`
	// Map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. Downgrading an `Advanced` tier parameter to `Standard` will recreate the resource. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier *string `pulumi:"tier"`
	// Type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	//
	// The following arguments are optional:
	Type string `pulumi:"type"`
	// Value of the parameter. This value is always marked as sensitive in the pulumi preview output, regardless of `type`.
	//
	// > **NOTE:** `aws:ssm:integration` dataType parameters must be of the type `SecureString` and the name must start with the prefix `/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/`. See [here](https://docs.aws.amazon.com/systems-manager/latest/userguide/creating-integrations.html) for information on the usage of `aws:ssm:integration` parameters.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a Parameter resource.
type ParameterArgs struct {
	// Regular expression used to validate the parameter value.
	AllowedPattern pulumi.StringPtrInput
	// ARN of the parameter.
	Arn pulumi.StringPtrInput
	// Data type of the parameter. Valid values: `text`, `aws:ssm:integration` and `aws:ec2:image` for AMI format, see the [Native parameter support for Amazon Machine Image IDs](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html).
	DataType pulumi.StringPtrInput
	// Description of the parameter.
	Description pulumi.StringPtrInput
	// Value of the parameter. **Use caution:** This value is _never_ marked as sensitive in the pulumi preview output. This argument is not valid with a `type` of `SecureString`.
	InsecureValue pulumi.StringPtrInput
	// KMS key ID or ARN for encrypting a SecureString.
	KeyId pulumi.StringPtrInput
	// Name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name pulumi.StringPtrInput
	// Overwrite an existing parameter. If not specified, defaults to `false` if the resource has not been created by Pulumi to avoid overwrite of existing resource, and will default to `true` otherwise (Pulumi lifecycle rules should then be used to manage the update behavior).
	//
	// Deprecated: this attribute has been deprecated
	Overwrite pulumi.BoolPtrInput
	// Map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. Downgrading an `Advanced` tier parameter to `Standard` will recreate the resource. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier pulumi.StringPtrInput
	// Type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	//
	// The following arguments are optional:
	Type pulumi.StringInput
	// Value of the parameter. This value is always marked as sensitive in the pulumi preview output, regardless of `type`.
	//
	// > **NOTE:** `aws:ssm:integration` dataType parameters must be of the type `SecureString` and the name must start with the prefix `/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/`. See [here](https://docs.aws.amazon.com/systems-manager/latest/userguide/creating-integrations.html) for information on the usage of `aws:ssm:integration` parameters.
	Value pulumi.StringPtrInput
}

func (ParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterArgs)(nil)).Elem()
}

type ParameterInput interface {
	pulumi.Input

	ToParameterOutput() ParameterOutput
	ToParameterOutputWithContext(ctx context.Context) ParameterOutput
}

func (*Parameter) ElementType() reflect.Type {
	return reflect.TypeOf((**Parameter)(nil)).Elem()
}

func (i *Parameter) ToParameterOutput() ParameterOutput {
	return i.ToParameterOutputWithContext(context.Background())
}

func (i *Parameter) ToParameterOutputWithContext(ctx context.Context) ParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterOutput)
}

// ParameterArrayInput is an input type that accepts ParameterArray and ParameterArrayOutput values.
// You can construct a concrete instance of `ParameterArrayInput` via:
//
//	ParameterArray{ ParameterArgs{...} }
type ParameterArrayInput interface {
	pulumi.Input

	ToParameterArrayOutput() ParameterArrayOutput
	ToParameterArrayOutputWithContext(context.Context) ParameterArrayOutput
}

type ParameterArray []ParameterInput

func (ParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parameter)(nil)).Elem()
}

func (i ParameterArray) ToParameterArrayOutput() ParameterArrayOutput {
	return i.ToParameterArrayOutputWithContext(context.Background())
}

func (i ParameterArray) ToParameterArrayOutputWithContext(ctx context.Context) ParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterArrayOutput)
}

// ParameterMapInput is an input type that accepts ParameterMap and ParameterMapOutput values.
// You can construct a concrete instance of `ParameterMapInput` via:
//
//	ParameterMap{ "key": ParameterArgs{...} }
type ParameterMapInput interface {
	pulumi.Input

	ToParameterMapOutput() ParameterMapOutput
	ToParameterMapOutputWithContext(context.Context) ParameterMapOutput
}

type ParameterMap map[string]ParameterInput

func (ParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parameter)(nil)).Elem()
}

func (i ParameterMap) ToParameterMapOutput() ParameterMapOutput {
	return i.ToParameterMapOutputWithContext(context.Background())
}

func (i ParameterMap) ToParameterMapOutputWithContext(ctx context.Context) ParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterMapOutput)
}

type ParameterOutput struct{ *pulumi.OutputState }

func (ParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Parameter)(nil)).Elem()
}

func (o ParameterOutput) ToParameterOutput() ParameterOutput {
	return o
}

func (o ParameterOutput) ToParameterOutputWithContext(ctx context.Context) ParameterOutput {
	return o
}

// Regular expression used to validate the parameter value.
func (o ParameterOutput) AllowedPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringPtrOutput { return v.AllowedPattern }).(pulumi.StringPtrOutput)
}

// ARN of the parameter.
func (o ParameterOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Data type of the parameter. Valid values: `text`, `aws:ssm:integration` and `aws:ec2:image` for AMI format, see the [Native parameter support for Amazon Machine Image IDs](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html).
func (o ParameterOutput) DataType() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.DataType }).(pulumi.StringOutput)
}

// Description of the parameter.
func (o ParameterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Value of the parameter. **Use caution:** This value is _never_ marked as sensitive in the pulumi preview output. This argument is not valid with a `type` of `SecureString`.
func (o ParameterOutput) InsecureValue() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.InsecureValue }).(pulumi.StringOutput)
}

// KMS key ID or ARN for encrypting a SecureString.
func (o ParameterOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.KeyId }).(pulumi.StringOutput)
}

// Name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
func (o ParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Overwrite an existing parameter. If not specified, defaults to `false` if the resource has not been created by Pulumi to avoid overwrite of existing resource, and will default to `true` otherwise (Pulumi lifecycle rules should then be used to manage the update behavior).
//
// Deprecated: this attribute has been deprecated
func (o ParameterOutput) Overwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Parameter) pulumi.BoolPtrOutput { return v.Overwrite }).(pulumi.BoolPtrOutput)
}

// Map of tags to assign to the object. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ParameterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o ParameterOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. Downgrading an `Advanced` tier parameter to `Standard` will recreate the resource. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
func (o ParameterOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringPtrOutput { return v.Tier }).(pulumi.StringPtrOutput)
}

// Type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
//
// The following arguments are optional:
func (o ParameterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Value of the parameter. This value is always marked as sensitive in the pulumi preview output, regardless of `type`.
//
// > **NOTE:** `aws:ssm:integration` dataType parameters must be of the type `SecureString` and the name must start with the prefix `/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/`. See [here](https://docs.aws.amazon.com/systems-manager/latest/userguide/creating-integrations.html) for information on the usage of `aws:ssm:integration` parameters.
func (o ParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameter) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

// Version of the parameter.
func (o ParameterOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *Parameter) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type ParameterArrayOutput struct{ *pulumi.OutputState }

func (ParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parameter)(nil)).Elem()
}

func (o ParameterArrayOutput) ToParameterArrayOutput() ParameterArrayOutput {
	return o
}

func (o ParameterArrayOutput) ToParameterArrayOutputWithContext(ctx context.Context) ParameterArrayOutput {
	return o
}

func (o ParameterArrayOutput) Index(i pulumi.IntInput) ParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Parameter {
		return vs[0].([]*Parameter)[vs[1].(int)]
	}).(ParameterOutput)
}

type ParameterMapOutput struct{ *pulumi.OutputState }

func (ParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parameter)(nil)).Elem()
}

func (o ParameterMapOutput) ToParameterMapOutput() ParameterMapOutput {
	return o
}

func (o ParameterMapOutput) ToParameterMapOutputWithContext(ctx context.Context) ParameterMapOutput {
	return o
}

func (o ParameterMapOutput) MapIndex(k pulumi.StringInput) ParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Parameter {
		return vs[0].(map[string]*Parameter)[vs[1].(string)]
	}).(ParameterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterInput)(nil)).Elem(), &Parameter{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterArrayInput)(nil)).Elem(), ParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterMapInput)(nil)).Elem(), ParameterMap{})
	pulumi.RegisterOutputType(ParameterOutput{})
	pulumi.RegisterOutputType(ParameterArrayOutput{})
	pulumi.RegisterOutputType(ParameterMapOutput{})
}
