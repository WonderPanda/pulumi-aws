// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage an API Gateway Documentation Version.
type DocumentationVersion struct {
	s *pulumi.ResourceState
}

// NewDocumentationVersion registers a new resource with the given unique name, arguments, and options.
func NewDocumentationVersion(ctx *pulumi.Context,
	name string, args *DocumentationVersionArgs, opts ...pulumi.ResourceOpt) (*DocumentationVersion, error) {
	if args == nil || args.RestApiId == nil {
		return nil, errors.New("missing required argument 'RestApiId'")
	}
	if args == nil || args.Version == nil {
		return nil, errors.New("missing required argument 'Version'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["restApiId"] = nil
		inputs["version"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["restApiId"] = args.RestApiId
		inputs["version"] = args.Version
	}
	s, err := ctx.RegisterResource("aws:apigateway/documentationVersion:DocumentationVersion", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DocumentationVersion{s: s}, nil
}

// GetDocumentationVersion gets an existing DocumentationVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentationVersion(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DocumentationVersionState, opts ...pulumi.ResourceOpt) (*DocumentationVersion, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["restApiId"] = state.RestApiId
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("aws:apigateway/documentationVersion:DocumentationVersion", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DocumentationVersion{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DocumentationVersion) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DocumentationVersion) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The description of the API documentation version.
func (r *DocumentationVersion) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The ID of the associated Rest API
func (r *DocumentationVersion) RestApiId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["restApiId"])
}

// The version identifier of the API documentation snapshot.
func (r *DocumentationVersion) Version() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering DocumentationVersion resources.
type DocumentationVersionState struct {
	// The description of the API documentation version.
	Description interface{}
	// The ID of the associated Rest API
	RestApiId interface{}
	// The version identifier of the API documentation snapshot.
	Version interface{}
}

// The set of arguments for constructing a DocumentationVersion resource.
type DocumentationVersionArgs struct {
	// The description of the API documentation version.
	Description interface{}
	// The ID of the associated Rest API
	RestApiId interface{}
	// The version identifier of the API documentation snapshot.
	Version interface{}
}
