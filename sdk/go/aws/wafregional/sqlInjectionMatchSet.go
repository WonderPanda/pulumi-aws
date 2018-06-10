// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Regional SQL Injection Match Set Resource for use with Application Load Balancer.
type SqlInjectionMatchSet struct {
	s *pulumi.ResourceState
}

// NewSqlInjectionMatchSet registers a new resource with the given unique name, arguments, and options.
func NewSqlInjectionMatchSet(ctx *pulumi.Context,
	name string, args *SqlInjectionMatchSetArgs, opts ...pulumi.ResourceOpt) (*SqlInjectionMatchSet, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["sqlInjectionMatchTuples"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["sqlInjectionMatchTuples"] = args.SqlInjectionMatchTuples
	}
	s, err := ctx.RegisterResource("aws:wafregional/sqlInjectionMatchSet:SqlInjectionMatchSet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SqlInjectionMatchSet{s: s}, nil
}

// GetSqlInjectionMatchSet gets an existing SqlInjectionMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlInjectionMatchSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SqlInjectionMatchSetState, opts ...pulumi.ResourceOpt) (*SqlInjectionMatchSet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["sqlInjectionMatchTuples"] = state.SqlInjectionMatchTuples
	}
	s, err := ctx.ReadResource("aws:wafregional/sqlInjectionMatchSet:SqlInjectionMatchSet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SqlInjectionMatchSet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SqlInjectionMatchSet) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SqlInjectionMatchSet) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The name or description of the SizeConstraintSet.
func (r *SqlInjectionMatchSet) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
func (r *SqlInjectionMatchSet) SqlInjectionMatchTuples() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["sqlInjectionMatchTuples"])
}

// Input properties used for looking up and filtering SqlInjectionMatchSet resources.
type SqlInjectionMatchSetState struct {
	// The name or description of the SizeConstraintSet.
	Name interface{}
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples interface{}
}

// The set of arguments for constructing a SqlInjectionMatchSet resource.
type SqlInjectionMatchSetArgs struct {
	// The name or description of the SizeConstraintSet.
	Name interface{}
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples interface{}
}
