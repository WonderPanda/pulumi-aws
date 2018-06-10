// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Service Discovery Private DNS Namespace resource.
type PrivateDnsNamespace struct {
	s *pulumi.ResourceState
}

// NewPrivateDnsNamespace registers a new resource with the given unique name, arguments, and options.
func NewPrivateDnsNamespace(ctx *pulumi.Context,
	name string, args *PrivateDnsNamespaceArgs, opts ...pulumi.ResourceOpt) (*PrivateDnsNamespace, error) {
	if args == nil || args.Vpc == nil {
		return nil, errors.New("missing required argument 'Vpc'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["vpc"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["vpc"] = args.Vpc
	}
	inputs["arn"] = nil
	inputs["hostedZone"] = nil
	s, err := ctx.RegisterResource("aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PrivateDnsNamespace{s: s}, nil
}

// GetPrivateDnsNamespace gets an existing PrivateDnsNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateDnsNamespace(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PrivateDnsNamespaceState, opts ...pulumi.ResourceOpt) (*PrivateDnsNamespace, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["description"] = state.Description
		inputs["hostedZone"] = state.HostedZone
		inputs["name"] = state.Name
		inputs["vpc"] = state.Vpc
	}
	s, err := ctx.ReadResource("aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PrivateDnsNamespace{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PrivateDnsNamespace) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PrivateDnsNamespace) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The ARN that Amazon Route 53 assigns to the namespace when you create it.
func (r *PrivateDnsNamespace) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The description that you specify for the namespace when you create it.
func (r *PrivateDnsNamespace) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
func (r *PrivateDnsNamespace) HostedZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["hostedZone"])
}

// The name of the namespace.
func (r *PrivateDnsNamespace) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of VPC that you want to associate the namespace with.
func (r *PrivateDnsNamespace) Vpc() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpc"])
}

// Input properties used for looking up and filtering PrivateDnsNamespace resources.
type PrivateDnsNamespaceState struct {
	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn interface{}
	// The description that you specify for the namespace when you create it.
	Description interface{}
	// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
	HostedZone interface{}
	// The name of the namespace.
	Name interface{}
	// The ID of VPC that you want to associate the namespace with.
	Vpc interface{}
}

// The set of arguments for constructing a PrivateDnsNamespace resource.
type PrivateDnsNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description interface{}
	// The name of the namespace.
	Name interface{}
	// The ID of VPC that you want to associate the namespace with.
	Vpc interface{}
}
