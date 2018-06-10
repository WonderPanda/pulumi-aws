// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `aws_subnet` provides details about a specific VPC subnet.
// 
// This resource can prove useful when a module accepts a subnet id as
// an input variable and needs to, for example, determine the id of the
// VPC that the subnet belongs to.
func Lookupubnet(ctx *pulumi.Context, args *GetSubnetArgs) (*GetSubnetResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["cidrBlock"] = args.CidrBlock
		inputs["defaultForAz"] = args.DefaultForAz
		inputs["filters"] = args.Filters
		inputs["id"] = args.Id
		inputs["ipv6CidrBlock"] = args.Ipv6CidrBlock
		inputs["state"] = args.State
		inputs["tags"] = args.Tags
		inputs["vpcId"] = args.VpcId
	}
	outputs, err := ctx.Invoke("aws:ec2/getSubnet:getSubnet", inputs)
	if err != nil {
		return nil, err
	}
	ret := GetSubnetResult{}
	if v, ok := outputs["assignIpv6AddressOnCreation"]; ok {
		ret.AssignIpv6AddressOnCreation = v
	}
	if v, ok := outputs["availabilityZone"]; ok {
		ret.AvailabilityZone = v
	}
	if v, ok := outputs["cidrBlock"]; ok {
		ret.CidrBlock = v
	}
	if v, ok := outputs["defaultForAz"]; ok {
		ret.DefaultForAz = v
	}
	if v, ok := outputs["id"]; ok {
		ret.Id = v
	}
	if v, ok := outputs["ipv6CidrBlock"]; ok {
		ret.Ipv6CidrBlock = v
	}
	if v, ok := outputs["ipv6CidrBlockAssociationId"]; ok {
		ret.Ipv6CidrBlockAssociationId = v
	}
	if v, ok := outputs["mapPublicIpOnLaunch"]; ok {
		ret.MapPublicIpOnLaunch = v
	}
	if v, ok := outputs["state"]; ok {
		ret.State = v
	}
	if v, ok := outputs["tags"]; ok {
		ret.Tags = v
	}
	if v, ok := outputs["vpcId"]; ok {
		ret.VpcId = v
	}
	return &ret, nil
}

// A collection of arguments for invoking getSubnet.
type GetSubnetArgs struct {
	// The availability zone where the
	// subnet must reside.
	AvailabilityZone interface{}
	// The cidr block of the desired subnet.
	CidrBlock interface{}
	// Boolean constraint for whether the desired
	// subnet must be the default subnet for its associated availability zone.
	DefaultForAz interface{}
	// Custom filter block as described below.
	Filters interface{}
	// The id of the specific subnet to retrieve.
	Id interface{}
	// The Ipv6 cidr block of the desired subnet
	Ipv6CidrBlock interface{}
	// The state that the desired subnet must have.
	State interface{}
	// A mapping of tags, each pair of which must exactly match
	// a pair on the desired subnet.
	Tags interface{}
	// The id of the VPC that the desired subnet belongs to.
	VpcId interface{}
}

// A collection of values returned by getSubnet.
type GetSubnetResult struct {
	AssignIpv6AddressOnCreation interface{}
	AvailabilityZone interface{}
	CidrBlock interface{}
	DefaultForAz interface{}
	Id interface{}
	Ipv6CidrBlock interface{}
	Ipv6CidrBlockAssociationId interface{}
	MapPublicIpOnLaunch interface{}
	State interface{}
	Tags interface{}
	VpcId interface{}
}
