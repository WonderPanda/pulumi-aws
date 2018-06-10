// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage a VPC peering connection.
// 
// ~> **NOTE on VPC Peering Connections and VPC Peering Connection Options:** Terraform provides
// both a standalone [VPC Peering Connection Options](vpc_peering_options.html) and a VPC Peering Connection
// resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
// connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
// Doing so will cause a conflict of options and will overwrite the options.
// Using a VPC Peering Connection Options resource decouples management of the connection options from
// management of the VPC Peering Connection and allows options to be set correctly in cross-account scenarios.
// 
// -> **Note:** For cross-account (requester's AWS account differs from the accepter's AWS account) or inter-region
// VPC Peering Connections use the `aws_vpc_peering_connection` resource to manage the requester's side of the
// connection and use the `aws_vpc_peering_connection_accepter` resource to manage the accepter's side of the connection.
type VpcPeeringConnection struct {
	s *pulumi.ResourceState
}

// NewVpcPeeringConnection registers a new resource with the given unique name, arguments, and options.
func NewVpcPeeringConnection(ctx *pulumi.Context,
	name string, args *VpcPeeringConnectionArgs, opts ...pulumi.ResourceOpt) (*VpcPeeringConnection, error) {
	if args == nil || args.PeerVpcId == nil {
		return nil, errors.New("missing required argument 'PeerVpcId'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accepter"] = nil
		inputs["autoAccept"] = nil
		inputs["peerOwnerId"] = nil
		inputs["peerRegion"] = nil
		inputs["peerVpcId"] = nil
		inputs["requester"] = nil
		inputs["tags"] = nil
		inputs["vpcId"] = nil
	} else {
		inputs["accepter"] = args.Accepter
		inputs["autoAccept"] = args.AutoAccept
		inputs["peerOwnerId"] = args.PeerOwnerId
		inputs["peerRegion"] = args.PeerRegion
		inputs["peerVpcId"] = args.PeerVpcId
		inputs["requester"] = args.Requester
		inputs["tags"] = args.Tags
		inputs["vpcId"] = args.VpcId
	}
	inputs["acceptStatus"] = nil
	s, err := ctx.RegisterResource("aws:ec2/vpcPeeringConnection:VpcPeeringConnection", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcPeeringConnection{s: s}, nil
}

// GetVpcPeeringConnection gets an existing VpcPeeringConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPeeringConnection(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VpcPeeringConnectionState, opts ...pulumi.ResourceOpt) (*VpcPeeringConnection, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["acceptStatus"] = state.AcceptStatus
		inputs["accepter"] = state.Accepter
		inputs["autoAccept"] = state.AutoAccept
		inputs["peerOwnerId"] = state.PeerOwnerId
		inputs["peerRegion"] = state.PeerRegion
		inputs["peerVpcId"] = state.PeerVpcId
		inputs["requester"] = state.Requester
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("aws:ec2/vpcPeeringConnection:VpcPeeringConnection", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcPeeringConnection{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VpcPeeringConnection) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VpcPeeringConnection) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The status of the VPC Peering Connection request.
func (r *VpcPeeringConnection) AcceptStatus() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["acceptStatus"])
}

// An optional configuration block that allows for [VPC Peering Connection]
// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
// the peering connection (a maximum of one).
func (r *VpcPeeringConnection) Accepter() *pulumi.Output {
	return r.s.State["accepter"]
}

// Accept the peering (both VPCs need to be in the same AWS account).
func (r *VpcPeeringConnection) AutoAccept() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoAccept"])
}

// The AWS account ID of the owner of the peer VPC.
// Defaults to the account ID the [AWS provider][1] is currently connected to.
func (r *VpcPeeringConnection) PeerOwnerId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["peerOwnerId"])
}

// The region of the accepter VPC of the [VPC Peering Connection]. `auto_accept` must be `false`,
// and use the `aws_vpc_peering_connection_accepter` to manage the accepter side.
func (r *VpcPeeringConnection) PeerRegion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["peerRegion"])
}

// The ID of the VPC with which you are creating the VPC Peering Connection.
func (r *VpcPeeringConnection) PeerVpcId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["peerVpcId"])
}

// A optional configuration block that allows for [VPC Peering Connection]
// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
// the peering connection (a maximum of one).
func (r *VpcPeeringConnection) Requester() *pulumi.Output {
	return r.s.State["requester"]
}

// A mapping of tags to assign to the resource.
func (r *VpcPeeringConnection) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The ID of the requester VPC.
func (r *VpcPeeringConnection) VpcId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering VpcPeeringConnection resources.
type VpcPeeringConnectionState struct {
	// The status of the VPC Peering Connection request.
	AcceptStatus interface{}
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter interface{}
	// Accept the peering (both VPCs need to be in the same AWS account).
	AutoAccept interface{}
	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the [AWS provider][1] is currently connected to.
	PeerOwnerId interface{}
	// The region of the accepter VPC of the [VPC Peering Connection]. `auto_accept` must be `false`,
	// and use the `aws_vpc_peering_connection_accepter` to manage the accepter side.
	PeerRegion interface{}
	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVpcId interface{}
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The ID of the requester VPC.
	VpcId interface{}
}

// The set of arguments for constructing a VpcPeeringConnection resource.
type VpcPeeringConnectionArgs struct {
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter interface{}
	// Accept the peering (both VPCs need to be in the same AWS account).
	AutoAccept interface{}
	// The AWS account ID of the owner of the peer VPC.
	// Defaults to the account ID the [AWS provider][1] is currently connected to.
	PeerOwnerId interface{}
	// The region of the accepter VPC of the [VPC Peering Connection]. `auto_accept` must be `false`,
	// and use the `aws_vpc_peering_connection_accepter` to manage the accepter side.
	PeerRegion interface{}
	// The ID of the VPC with which you are creating the VPC Peering Connection.
	PeerVpcId interface{}
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The ID of the requester VPC.
	VpcId interface{}
}
