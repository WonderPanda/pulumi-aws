// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage AWS EMR Security Configurations
type SecurityConfiguration struct {
	s *pulumi.ResourceState
}

// NewSecurityConfiguration registers a new resource with the given unique name, arguments, and options.
func NewSecurityConfiguration(ctx *pulumi.Context,
	name string, args *SecurityConfigurationArgs, opts ...pulumi.ResourceOpt) (*SecurityConfiguration, error) {
	if args == nil || args.Configuration == nil {
		return nil, errors.New("missing required argument 'Configuration'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["configuration"] = nil
		inputs["name"] = nil
		inputs["namePrefix"] = nil
	} else {
		inputs["configuration"] = args.Configuration
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
	}
	inputs["creationDate"] = nil
	s, err := ctx.RegisterResource("aws:emr/securityConfiguration:SecurityConfiguration", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurityConfiguration{s: s}, nil
}

// GetSecurityConfiguration gets an existing SecurityConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityConfiguration(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecurityConfigurationState, opts ...pulumi.ResourceOpt) (*SecurityConfiguration, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["configuration"] = state.Configuration
		inputs["creationDate"] = state.CreationDate
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
	}
	s, err := ctx.ReadResource("aws:emr/securityConfiguration:SecurityConfiguration", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurityConfiguration{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecurityConfiguration) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecurityConfiguration) ID() *pulumi.IDOutput {
	return r.s.ID
}

// A JSON formatted Security Configuration
func (r *SecurityConfiguration) Configuration() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["configuration"])
}

// Date the Security Configuration was created
func (r *SecurityConfiguration) CreationDate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationDate"])
}

// The name of the EMR Security Configuration. By default generated by Terraform.
func (r *SecurityConfiguration) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Creates a unique name beginning with the specified
// prefix. Conflicts with `name`.
func (r *SecurityConfiguration) NamePrefix() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["namePrefix"])
}

// Input properties used for looking up and filtering SecurityConfiguration resources.
type SecurityConfigurationState struct {
	// A JSON formatted Security Configuration
	Configuration interface{}
	// Date the Security Configuration was created
	CreationDate interface{}
	// The name of the EMR Security Configuration. By default generated by Terraform.
	Name interface{}
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix interface{}
}

// The set of arguments for constructing a SecurityConfiguration resource.
type SecurityConfigurationArgs struct {
	// A JSON formatted Security Configuration
	Configuration interface{}
	// The name of the EMR Security Configuration. By default generated by Terraform.
	Name interface{}
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix interface{}
}
