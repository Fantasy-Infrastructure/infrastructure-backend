// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ec2.Route` provides details about a specific Route.
//
// This resource can prove useful when finding the resource associated with a CIDR. For example, finding the peering connection associated with a CIDR value.
//
// ## Example Usage
//
// The following example shows how one might use a CIDR value to find a network interface id and use this to create a data source of that network interface.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			subnetId := cfg.RequireObject("subnetId")
//			_, err := ec2.LookupRouteTable(ctx, &ec2.LookupRouteTableArgs{
//				SubnetId: pulumi.StringRef(subnetId),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			route, err := ec2.LookupRoute(ctx, &ec2.LookupRouteArgs{
//				RouteTableId:         selectedAwsRouteTable.Id,
//				DestinationCidrBlock: pulumi.StringRef("10.0.1.0/24"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ec2.LookupNetworkInterface(ctx, &ec2.LookupNetworkInterfaceArgs{
//				Id: pulumi.StringRef(route.NetworkInterfaceId),
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
func LookupRoute(ctx *pulumi.Context, args *LookupRouteArgs, opts ...pulumi.InvokeOption) (*LookupRouteResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteResult
	err := ctx.Invoke("aws:ec2/getRoute:getRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoute.
type LookupRouteArgs struct {
	// EC2 Carrier Gateway ID of the Route belonging to the Route Table.
	CarrierGatewayId *string `pulumi:"carrierGatewayId"`
	// Core network ARN of the Route belonging to the Route Table.
	CoreNetworkArn *string `pulumi:"coreNetworkArn"`
	// CIDR block of the Route belonging to the Route Table.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// IPv6 CIDR block of the Route belonging to the Route Table.
	DestinationIpv6CidrBlock *string `pulumi:"destinationIpv6CidrBlock"`
	// ID of a managed prefix list destination of the Route belonging to the Route Table.
	DestinationPrefixListId *string `pulumi:"destinationPrefixListId"`
	// Egress Only Gateway ID of the Route belonging to the Route Table.
	EgressOnlyGatewayId *string `pulumi:"egressOnlyGatewayId"`
	// Gateway ID of the Route belonging to the Route Table.
	GatewayId *string `pulumi:"gatewayId"`
	// Instance ID of the Route belonging to the Route Table.
	InstanceId *string `pulumi:"instanceId"`
	// Local Gateway ID of the Route belonging to the Route Table.
	LocalGatewayId *string `pulumi:"localGatewayId"`
	// NAT Gateway ID of the Route belonging to the Route Table.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// Network Interface ID of the Route belonging to the Route Table.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// ID of the specific Route Table containing the Route entry.
	//
	// The following arguments are optional:
	RouteTableId string `pulumi:"routeTableId"`
	// EC2 Transit Gateway ID of the Route belonging to the Route Table.
	TransitGatewayId *string `pulumi:"transitGatewayId"`
	// VPC Peering Connection ID of the Route belonging to the Route Table.
	VpcPeeringConnectionId *string `pulumi:"vpcPeeringConnectionId"`
}

// A collection of values returned by getRoute.
type LookupRouteResult struct {
	CarrierGatewayId         string `pulumi:"carrierGatewayId"`
	CoreNetworkArn           string `pulumi:"coreNetworkArn"`
	DestinationCidrBlock     string `pulumi:"destinationCidrBlock"`
	DestinationIpv6CidrBlock string `pulumi:"destinationIpv6CidrBlock"`
	DestinationPrefixListId  string `pulumi:"destinationPrefixListId"`
	EgressOnlyGatewayId      string `pulumi:"egressOnlyGatewayId"`
	GatewayId                string `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string `pulumi:"id"`
	InstanceId             string `pulumi:"instanceId"`
	LocalGatewayId         string `pulumi:"localGatewayId"`
	NatGatewayId           string `pulumi:"natGatewayId"`
	NetworkInterfaceId     string `pulumi:"networkInterfaceId"`
	RouteTableId           string `pulumi:"routeTableId"`
	TransitGatewayId       string `pulumi:"transitGatewayId"`
	VpcPeeringConnectionId string `pulumi:"vpcPeeringConnectionId"`
}

func LookupRouteOutput(ctx *pulumi.Context, args LookupRouteOutputArgs, opts ...pulumi.InvokeOption) LookupRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteResult, error) {
			args := v.(LookupRouteArgs)
			r, err := LookupRoute(ctx, &args, opts...)
			var s LookupRouteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteResultOutput)
}

// A collection of arguments for invoking getRoute.
type LookupRouteOutputArgs struct {
	// EC2 Carrier Gateway ID of the Route belonging to the Route Table.
	CarrierGatewayId pulumi.StringPtrInput `pulumi:"carrierGatewayId"`
	// Core network ARN of the Route belonging to the Route Table.
	CoreNetworkArn pulumi.StringPtrInput `pulumi:"coreNetworkArn"`
	// CIDR block of the Route belonging to the Route Table.
	DestinationCidrBlock pulumi.StringPtrInput `pulumi:"destinationCidrBlock"`
	// IPv6 CIDR block of the Route belonging to the Route Table.
	DestinationIpv6CidrBlock pulumi.StringPtrInput `pulumi:"destinationIpv6CidrBlock"`
	// ID of a managed prefix list destination of the Route belonging to the Route Table.
	DestinationPrefixListId pulumi.StringPtrInput `pulumi:"destinationPrefixListId"`
	// Egress Only Gateway ID of the Route belonging to the Route Table.
	EgressOnlyGatewayId pulumi.StringPtrInput `pulumi:"egressOnlyGatewayId"`
	// Gateway ID of the Route belonging to the Route Table.
	GatewayId pulumi.StringPtrInput `pulumi:"gatewayId"`
	// Instance ID of the Route belonging to the Route Table.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// Local Gateway ID of the Route belonging to the Route Table.
	LocalGatewayId pulumi.StringPtrInput `pulumi:"localGatewayId"`
	// NAT Gateway ID of the Route belonging to the Route Table.
	NatGatewayId pulumi.StringPtrInput `pulumi:"natGatewayId"`
	// Network Interface ID of the Route belonging to the Route Table.
	NetworkInterfaceId pulumi.StringPtrInput `pulumi:"networkInterfaceId"`
	// ID of the specific Route Table containing the Route entry.
	//
	// The following arguments are optional:
	RouteTableId pulumi.StringInput `pulumi:"routeTableId"`
	// EC2 Transit Gateway ID of the Route belonging to the Route Table.
	TransitGatewayId pulumi.StringPtrInput `pulumi:"transitGatewayId"`
	// VPC Peering Connection ID of the Route belonging to the Route Table.
	VpcPeeringConnectionId pulumi.StringPtrInput `pulumi:"vpcPeeringConnectionId"`
}

func (LookupRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteArgs)(nil)).Elem()
}

// A collection of values returned by getRoute.
type LookupRouteResultOutput struct{ *pulumi.OutputState }

func (LookupRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteResult)(nil)).Elem()
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutput() LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutputWithContext(ctx context.Context) LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) CarrierGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.CarrierGatewayId }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) CoreNetworkArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.CoreNetworkArn }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) DestinationCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.DestinationCidrBlock }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) DestinationIpv6CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.DestinationIpv6CidrBlock }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) DestinationPrefixListId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.DestinationPrefixListId }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) EgressOnlyGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.EgressOnlyGatewayId }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) LocalGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.LocalGatewayId }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) NatGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NatGatewayId }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.RouteTableId }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) TransitGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.TransitGatewayId }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) VpcPeeringConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.VpcPeeringConnectionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteResultOutput{})
}
