// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Size Constraint Set Resource
type SizeConstraintSet struct {
	s *pulumi.ResourceState
}

// NewSizeConstraintSet registers a new resource with the given unique name, arguments, and options.
func NewSizeConstraintSet(ctx *pulumi.Context,
	name string, args *SizeConstraintSetArgs, opts ...pulumi.ResourceOpt) (*SizeConstraintSet, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["sizeConstraints"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["sizeConstraints"] = args.SizeConstraints
	}
	s, err := ctx.RegisterResource("aws:waf/sizeConstraintSet:SizeConstraintSet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SizeConstraintSet{s: s}, nil
}

// GetSizeConstraintSet gets an existing SizeConstraintSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSizeConstraintSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SizeConstraintSetState, opts ...pulumi.ResourceOpt) (*SizeConstraintSet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["sizeConstraints"] = state.SizeConstraints
	}
	s, err := ctx.ReadResource("aws:waf/sizeConstraintSet:SizeConstraintSet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SizeConstraintSet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SizeConstraintSet) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SizeConstraintSet) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The name or description of the Size Constraint Set.
func (r *SizeConstraintSet) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Specifies the parts of web requests that you want to inspect the size of.
func (r *SizeConstraintSet) SizeConstraints() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["sizeConstraints"])
}

// Input properties used for looking up and filtering SizeConstraintSet resources.
type SizeConstraintSetState struct {
	// The name or description of the Size Constraint Set.
	Name interface{}
	// Specifies the parts of web requests that you want to inspect the size of.
	SizeConstraints interface{}
}

// The set of arguments for constructing a SizeConstraintSet resource.
type SizeConstraintSetArgs struct {
	// The name or description of the Size Constraint Set.
	Name interface{}
	// Specifies the parts of web requests that you want to inspect the size of.
	SizeConstraints interface{}
}
