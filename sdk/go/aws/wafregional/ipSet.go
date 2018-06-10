// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Regional IPSet Resource for use with Application Load Balancer.
type IpSet struct {
	s *pulumi.ResourceState
}

// NewIpSet registers a new resource with the given unique name, arguments, and options.
func NewIpSet(ctx *pulumi.Context,
	name string, args *IpSetArgs, opts ...pulumi.ResourceOpt) (*IpSet, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["ipSetDescriptors"] = nil
		inputs["name"] = nil
	} else {
		inputs["ipSetDescriptors"] = args.IpSetDescriptors
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("aws:wafregional/ipSet:IpSet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IpSet{s: s}, nil
}

// GetIpSet gets an existing IpSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IpSetState, opts ...pulumi.ResourceOpt) (*IpSet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["ipSetDescriptors"] = state.IpSetDescriptors
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("aws:wafregional/ipSet:IpSet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IpSet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IpSet) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IpSet) ID() *pulumi.IDOutput {
	return r.s.ID
}

// One or more pairs specifying the IP address type (IPV4 or IPV5) and the IP address range (in CIDR notation) from which web requests originate.
func (r *IpSet) IpSetDescriptors() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ipSetDescriptors"])
}

// The name or description of the IPSet.
func (r *IpSet) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering IpSet resources.
type IpSetState struct {
	// One or more pairs specifying the IP address type (IPV4 or IPV5) and the IP address range (in CIDR notation) from which web requests originate.
	IpSetDescriptors interface{}
	// The name or description of the IPSet.
	Name interface{}
}

// The set of arguments for constructing a IpSet resource.
type IpSetArgs struct {
	// One or more pairs specifying the IP address type (IPV4 or IPV5) and the IP address range (in CIDR notation) from which web requests originate.
	IpSetDescriptors interface{}
	// The name or description of the IPSet.
	Name interface{}
}
