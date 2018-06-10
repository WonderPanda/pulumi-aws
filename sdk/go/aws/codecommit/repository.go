// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codecommit

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CodeCommit Repository Resource.
// 
// ~> **NOTE on CodeCommit Availability**: The CodeCommit is not yet rolled out
// in all regions - available regions are listed
// [the AWS Docs](https://docs.aws.amazon.com/general/latest/gr/rande.html#codecommit_region).
type Repository struct {
	s *pulumi.ResourceState
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOpt) (*Repository, error) {
	if args == nil || args.RepositoryName == nil {
		return nil, errors.New("missing required argument 'RepositoryName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["defaultBranch"] = nil
		inputs["description"] = nil
		inputs["repositoryName"] = nil
	} else {
		inputs["defaultBranch"] = args.DefaultBranch
		inputs["description"] = args.Description
		inputs["repositoryName"] = args.RepositoryName
	}
	inputs["arn"] = nil
	inputs["cloneUrlHttp"] = nil
	inputs["cloneUrlSsh"] = nil
	inputs["repositoryId"] = nil
	s, err := ctx.RegisterResource("aws:codecommit/repository:Repository", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Repository{s: s}, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RepositoryState, opts ...pulumi.ResourceOpt) (*Repository, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["cloneUrlHttp"] = state.CloneUrlHttp
		inputs["cloneUrlSsh"] = state.CloneUrlSsh
		inputs["defaultBranch"] = state.DefaultBranch
		inputs["description"] = state.Description
		inputs["repositoryId"] = state.RepositoryId
		inputs["repositoryName"] = state.RepositoryName
	}
	s, err := ctx.ReadResource("aws:codecommit/repository:Repository", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Repository{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Repository) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Repository) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The ARN of the repository
func (r *Repository) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The URL to use for cloning the repository over HTTPS.
func (r *Repository) CloneUrlHttp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cloneUrlHttp"])
}

// The URL to use for cloning the repository over SSH.
func (r *Repository) CloneUrlSsh() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cloneUrlSsh"])
}

// The default branch of the repository. The branch specified here needs to exist.
func (r *Repository) DefaultBranch() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultBranch"])
}

// The description of the repository. This needs to be less than 1000 characters
func (r *Repository) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The ID of the repository
func (r *Repository) RepositoryId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["repositoryId"])
}

// The name for the repository. This needs to be less than 100 characters.
func (r *Repository) RepositoryName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["repositoryName"])
}

// Input properties used for looking up and filtering Repository resources.
type RepositoryState struct {
	// The ARN of the repository
	Arn interface{}
	// The URL to use for cloning the repository over HTTPS.
	CloneUrlHttp interface{}
	// The URL to use for cloning the repository over SSH.
	CloneUrlSsh interface{}
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch interface{}
	// The description of the repository. This needs to be less than 1000 characters
	Description interface{}
	// The ID of the repository
	RepositoryId interface{}
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName interface{}
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// The default branch of the repository. The branch specified here needs to exist.
	DefaultBranch interface{}
	// The description of the repository. This needs to be less than 1000 characters
	Description interface{}
	// The name for the repository. This needs to be less than 100 characters.
	RepositoryName interface{}
}
