// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Elastic Beanstalk Configuration Template, which are associated with
// a specific application and are used to deploy different versions of the
// application with the same configuration settings.
type ConfigurationTemplate struct {
	s *pulumi.ResourceState
}

// NewConfigurationTemplate registers a new resource with the given unique name, arguments, and options.
func NewConfigurationTemplate(ctx *pulumi.Context,
	name string, args *ConfigurationTemplateArgs, opts ...pulumi.ResourceOpt) (*ConfigurationTemplate, error) {
	if args == nil || args.Application == nil {
		return nil, errors.New("missing required argument 'Application'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["application"] = nil
		inputs["description"] = nil
		inputs["environmentId"] = nil
		inputs["name"] = nil
		inputs["settings"] = nil
		inputs["solutionStackName"] = nil
	} else {
		inputs["application"] = args.Application
		inputs["description"] = args.Description
		inputs["environmentId"] = args.EnvironmentId
		inputs["name"] = args.Name
		inputs["settings"] = args.Settings
		inputs["solutionStackName"] = args.SolutionStackName
	}
	s, err := ctx.RegisterResource("aws:elasticbeanstalk/configurationTemplate:ConfigurationTemplate", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ConfigurationTemplate{s: s}, nil
}

// GetConfigurationTemplate gets an existing ConfigurationTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationTemplate(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ConfigurationTemplateState, opts ...pulumi.ResourceOpt) (*ConfigurationTemplate, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["application"] = state.Application
		inputs["description"] = state.Description
		inputs["environmentId"] = state.EnvironmentId
		inputs["name"] = state.Name
		inputs["settings"] = state.Settings
		inputs["solutionStackName"] = state.SolutionStackName
	}
	s, err := ctx.ReadResource("aws:elasticbeanstalk/configurationTemplate:ConfigurationTemplate", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ConfigurationTemplate{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ConfigurationTemplate) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ConfigurationTemplate) ID() *pulumi.IDOutput {
	return r.s.ID
}

// name of the application to associate with this configuration template
func (r *ConfigurationTemplate) Application() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["application"])
}

// Short description of the Template
func (r *ConfigurationTemplate) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The ID of the environment used with this configuration template
func (r *ConfigurationTemplate) EnvironmentId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["environmentId"])
}

// A unique name for this Template.
func (r *ConfigurationTemplate) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Option settings to configure the new Environment. These
// override specific values that are set as defaults. The format is detailed
// below in [Option Settings](#option-settings)
func (r *ConfigurationTemplate) Settings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["settings"])
}

// A solution stack to base your Template
// off of. Example stacks can be found in the [Amazon API documentation][1]
func (r *ConfigurationTemplate) SolutionStackName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["solutionStackName"])
}

// Input properties used for looking up and filtering ConfigurationTemplate resources.
type ConfigurationTemplateState struct {
	// name of the application to associate with this configuration template
	Application interface{}
	// Short description of the Template
	Description interface{}
	// The ID of the environment used with this configuration template
	EnvironmentId interface{}
	// A unique name for this Template.
	Name interface{}
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in [Option Settings](#option-settings)
	Settings interface{}
	// A solution stack to base your Template
	// off of. Example stacks can be found in the [Amazon API documentation][1]
	SolutionStackName interface{}
}

// The set of arguments for constructing a ConfigurationTemplate resource.
type ConfigurationTemplateArgs struct {
	// name of the application to associate with this configuration template
	Application interface{}
	// Short description of the Template
	Description interface{}
	// The ID of the environment used with this configuration template
	EnvironmentId interface{}
	// A unique name for this Template.
	Name interface{}
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in [Option Settings](#option-settings)
	Settings interface{}
	// A solution stack to base your Template
	// off of. Example stacks can be found in the [Amazon API documentation][1]
	SolutionStackName interface{}
}
