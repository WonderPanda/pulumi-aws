// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.
type WebAcl struct {
	s *pulumi.ResourceState
}

// NewWebAcl registers a new resource with the given unique name, arguments, and options.
func NewWebAcl(ctx *pulumi.Context,
	name string, args *WebAclArgs, opts ...pulumi.ResourceOpt) (*WebAcl, error) {
	if args == nil || args.DefaultAction == nil {
		return nil, errors.New("missing required argument 'DefaultAction'")
	}
	if args == nil || args.MetricName == nil {
		return nil, errors.New("missing required argument 'MetricName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["defaultAction"] = nil
		inputs["metricName"] = nil
		inputs["name"] = nil
		inputs["rules"] = nil
	} else {
		inputs["defaultAction"] = args.DefaultAction
		inputs["metricName"] = args.MetricName
		inputs["name"] = args.Name
		inputs["rules"] = args.Rules
	}
	s, err := ctx.RegisterResource("aws:wafregional/webAcl:WebAcl", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &WebAcl{s: s}, nil
}

// GetWebAcl gets an existing WebAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAcl(ctx *pulumi.Context,
	name string, id pulumi.ID, state *WebAclState, opts ...pulumi.ResourceOpt) (*WebAcl, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["defaultAction"] = state.DefaultAction
		inputs["metricName"] = state.MetricName
		inputs["name"] = state.Name
		inputs["rules"] = state.Rules
	}
	s, err := ctx.ReadResource("aws:wafregional/webAcl:WebAcl", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &WebAcl{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *WebAcl) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *WebAcl) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
func (r *WebAcl) DefaultAction() *pulumi.Output {
	return r.s.State["defaultAction"]
}

// The name or description for the Amazon CloudWatch metric of this web ACL.
func (r *WebAcl) MetricName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["metricName"])
}

// The name or description of the web ACL.
func (r *WebAcl) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The rules to associate with the web ACL and the settings for each rule.
func (r *WebAcl) Rules() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["rules"])
}

// Input properties used for looking up and filtering WebAcl resources.
type WebAclState struct {
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction interface{}
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName interface{}
	// The name or description of the web ACL.
	Name interface{}
	// The rules to associate with the web ACL and the settings for each rule.
	Rules interface{}
}

// The set of arguments for constructing a WebAcl resource.
type WebAclArgs struct {
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction interface{}
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName interface{}
	// The name or description of the web ACL.
	Name interface{}
	// The rules to associate with the web ACL and the settings for each rule.
	Rules interface{}
}
