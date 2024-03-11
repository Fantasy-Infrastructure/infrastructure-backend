// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an RDS database instance snapshot. For managing RDS database cluster snapshots, see the `rds.ClusterSnapshot` resource.
//
// ## Example Usage
//
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
//			bar, err := rds.NewInstance(ctx, "bar", &rds.InstanceArgs{
//				AllocatedStorage:      pulumi.Int(10),
//				Engine:                pulumi.String("mysql"),
//				EngineVersion:         pulumi.String("5.6.21"),
//				InstanceClass:         pulumi.String("db.t2.micro"),
//				DbName:                pulumi.String("baz"),
//				Password:              pulumi.String("barbarbarbar"),
//				Username:              pulumi.String("foo"),
//				MaintenanceWindow:     pulumi.String("Fri:09:00-Fri:09:30"),
//				BackupRetentionPeriod: pulumi.Int(0),
//				ParameterGroupName:    pulumi.String("default.mysql5.6"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewSnapshot(ctx, "test", &rds.SnapshotArgs{
//				DbInstanceIdentifier: bar.Identifier,
//				DbSnapshotIdentifier: pulumi.String("testsnapshot1234"),
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
// Using `pulumi import`, import `aws_db_snapshot` using the snapshot identifier. For example:
//
// ```sh
//
//	$ pulumi import aws:rds/snapshot:Snapshot example my-snapshot
//
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage pulumi.IntOutput `pulumi:"allocatedStorage"`
	// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The DB Instance Identifier from which to take the snapshot.
	DbInstanceIdentifier pulumi.StringOutput `pulumi:"dbInstanceIdentifier"`
	// The Amazon Resource Name (ARN) for the DB snapshot.
	DbSnapshotArn pulumi.StringOutput `pulumi:"dbSnapshotArn"`
	// The Identifier for the snapshot.
	DbSnapshotIdentifier pulumi.StringOutput `pulumi:"dbSnapshotIdentifier"`
	// Specifies whether the DB snapshot is encrypted.
	Encrypted pulumi.BoolOutput `pulumi:"encrypted"`
	// Specifies the name of the database engine.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Specifies the version of the database engine.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops pulumi.IntOutput `pulumi:"iops"`
	// The ARN for the KMS encryption key.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// License model information for the restored DB instance.
	LicenseModel pulumi.StringOutput `pulumi:"licenseModel"`
	// Provides the option group name for the DB snapshot.
	OptionGroupName pulumi.StringOutput `pulumi:"optionGroupName"`
	Port            pulumi.IntOutput    `pulumi:"port"`
	// List of AWS Account ids to share snapshot with, use `all` to make snaphot public.
	SharedAccounts pulumi.StringArrayOutput `pulumi:"sharedAccounts"`
	SnapshotType   pulumi.StringOutput      `pulumi:"snapshotType"`
	// The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
	SourceDbSnapshotIdentifier pulumi.StringOutput `pulumi:"sourceDbSnapshotIdentifier"`
	// The region that the DB snapshot was created in or copied from.
	SourceRegion pulumi.StringOutput `pulumi:"sourceRegion"`
	// Specifies the status of this DB snapshot.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the storage type associated with DB snapshot.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Provides the VPC ID associated with the DB snapshot.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbInstanceIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceIdentifier'")
	}
	if args.DbSnapshotIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'DbSnapshotIdentifier'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("aws:rds/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("aws:rds/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage *int `pulumi:"allocatedStorage"`
	// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The DB Instance Identifier from which to take the snapshot.
	DbInstanceIdentifier *string `pulumi:"dbInstanceIdentifier"`
	// The Amazon Resource Name (ARN) for the DB snapshot.
	DbSnapshotArn *string `pulumi:"dbSnapshotArn"`
	// The Identifier for the snapshot.
	DbSnapshotIdentifier *string `pulumi:"dbSnapshotIdentifier"`
	// Specifies whether the DB snapshot is encrypted.
	Encrypted *bool `pulumi:"encrypted"`
	// Specifies the name of the database engine.
	Engine *string `pulumi:"engine"`
	// Specifies the version of the database engine.
	EngineVersion *string `pulumi:"engineVersion"`
	// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops *int `pulumi:"iops"`
	// The ARN for the KMS encryption key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// License model information for the restored DB instance.
	LicenseModel *string `pulumi:"licenseModel"`
	// Provides the option group name for the DB snapshot.
	OptionGroupName *string `pulumi:"optionGroupName"`
	Port            *int    `pulumi:"port"`
	// List of AWS Account ids to share snapshot with, use `all` to make snaphot public.
	SharedAccounts []string `pulumi:"sharedAccounts"`
	SnapshotType   *string  `pulumi:"snapshotType"`
	// The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
	SourceDbSnapshotIdentifier *string `pulumi:"sourceDbSnapshotIdentifier"`
	// The region that the DB snapshot was created in or copied from.
	SourceRegion *string `pulumi:"sourceRegion"`
	// Specifies the status of this DB snapshot.
	Status *string `pulumi:"status"`
	// Specifies the storage type associated with DB snapshot.
	StorageType *string `pulumi:"storageType"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Provides the VPC ID associated with the DB snapshot.
	VpcId *string `pulumi:"vpcId"`
}

type SnapshotState struct {
	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage pulumi.IntPtrInput
	// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone pulumi.StringPtrInput
	// The DB Instance Identifier from which to take the snapshot.
	DbInstanceIdentifier pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the DB snapshot.
	DbSnapshotArn pulumi.StringPtrInput
	// The Identifier for the snapshot.
	DbSnapshotIdentifier pulumi.StringPtrInput
	// Specifies whether the DB snapshot is encrypted.
	Encrypted pulumi.BoolPtrInput
	// Specifies the name of the database engine.
	Engine pulumi.StringPtrInput
	// Specifies the version of the database engine.
	EngineVersion pulumi.StringPtrInput
	// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops pulumi.IntPtrInput
	// The ARN for the KMS encryption key.
	KmsKeyId pulumi.StringPtrInput
	// License model information for the restored DB instance.
	LicenseModel pulumi.StringPtrInput
	// Provides the option group name for the DB snapshot.
	OptionGroupName pulumi.StringPtrInput
	Port            pulumi.IntPtrInput
	// List of AWS Account ids to share snapshot with, use `all` to make snaphot public.
	SharedAccounts pulumi.StringArrayInput
	SnapshotType   pulumi.StringPtrInput
	// The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
	SourceDbSnapshotIdentifier pulumi.StringPtrInput
	// The region that the DB snapshot was created in or copied from.
	SourceRegion pulumi.StringPtrInput
	// Specifies the status of this DB snapshot.
	Status pulumi.StringPtrInput
	// Specifies the storage type associated with DB snapshot.
	StorageType pulumi.StringPtrInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// Provides the VPC ID associated with the DB snapshot.
	VpcId pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// The DB Instance Identifier from which to take the snapshot.
	DbInstanceIdentifier string `pulumi:"dbInstanceIdentifier"`
	// The Identifier for the snapshot.
	DbSnapshotIdentifier string `pulumi:"dbSnapshotIdentifier"`
	// List of AWS Account ids to share snapshot with, use `all` to make snaphot public.
	SharedAccounts []string `pulumi:"sharedAccounts"`
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// The DB Instance Identifier from which to take the snapshot.
	DbInstanceIdentifier pulumi.StringInput
	// The Identifier for the snapshot.
	DbSnapshotIdentifier pulumi.StringInput
	// List of AWS Account ids to share snapshot with, use `all` to make snaphot public.
	SharedAccounts pulumi.StringArrayInput
	// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//	SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//	SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// Specifies the allocated storage size in gigabytes (GB).
func (o SnapshotOutput) AllocatedStorage() pulumi.IntOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntOutput { return v.AllocatedStorage }).(pulumi.IntOutput)
}

// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
func (o SnapshotOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The DB Instance Identifier from which to take the snapshot.
func (o SnapshotOutput) DbInstanceIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.DbInstanceIdentifier }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) for the DB snapshot.
func (o SnapshotOutput) DbSnapshotArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.DbSnapshotArn }).(pulumi.StringOutput)
}

// The Identifier for the snapshot.
func (o SnapshotOutput) DbSnapshotIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.DbSnapshotIdentifier }).(pulumi.StringOutput)
}

// Specifies whether the DB snapshot is encrypted.
func (o SnapshotOutput) Encrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolOutput { return v.Encrypted }).(pulumi.BoolOutput)
}

// Specifies the name of the database engine.
func (o SnapshotOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Specifies the version of the database engine.
func (o SnapshotOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
func (o SnapshotOutput) Iops() pulumi.IntOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntOutput { return v.Iops }).(pulumi.IntOutput)
}

// The ARN for the KMS encryption key.
func (o SnapshotOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.KmsKeyId }).(pulumi.StringOutput)
}

// License model information for the restored DB instance.
func (o SnapshotOutput) LicenseModel() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.LicenseModel }).(pulumi.StringOutput)
}

// Provides the option group name for the DB snapshot.
func (o SnapshotOutput) OptionGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.OptionGroupName }).(pulumi.StringOutput)
}

func (o SnapshotOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// List of AWS Account ids to share snapshot with, use `all` to make snaphot public.
func (o SnapshotOutput) SharedAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringArrayOutput { return v.SharedAccounts }).(pulumi.StringArrayOutput)
}

func (o SnapshotOutput) SnapshotType() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SnapshotType }).(pulumi.StringOutput)
}

// The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
func (o SnapshotOutput) SourceDbSnapshotIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SourceDbSnapshotIdentifier }).(pulumi.StringOutput)
}

// The region that the DB snapshot was created in or copied from.
func (o SnapshotOutput) SourceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SourceRegion }).(pulumi.StringOutput)
}

// Specifies the status of this DB snapshot.
func (o SnapshotOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the storage type associated with DB snapshot.
func (o SnapshotOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.StorageType }).(pulumi.StringOutput)
}

// Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o SnapshotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o SnapshotOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Provides the VPC ID associated with the DB snapshot.
func (o SnapshotOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].([]*Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].(map[string]*Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotArrayInput)(nil)).Elem(), SnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotMapInput)(nil)).Elem(), SnapshotMap{})
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}
