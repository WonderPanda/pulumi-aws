// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Amazon CloudFront origin access identity.
// 
// For information about CloudFront distributions, see the
// [Amazon CloudFront Developer Guide][1]. For more information on generating
// origin access identities, see
// [Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content][2].
type OriginAccessIdentity struct {
	s *pulumi.ResourceState
}

// NewOriginAccessIdentity registers a new resource with the given unique name, arguments, and options.
func NewOriginAccessIdentity(ctx *pulumi.Context,
	name string, args *OriginAccessIdentityArgs, opts ...pulumi.ResourceOpt) (*OriginAccessIdentity, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["comment"] = nil
	} else {
		inputs["comment"] = args.Comment
	}
	inputs["callerReference"] = nil
	inputs["cloudfrontAccessIdentityPath"] = nil
	inputs["etag"] = nil
	inputs["iamArn"] = nil
	inputs["s3CanonicalUserId"] = nil
	s, err := ctx.RegisterResource("aws:cloudfront/originAccessIdentity:OriginAccessIdentity", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OriginAccessIdentity{s: s}, nil
}

// GetOriginAccessIdentity gets an existing OriginAccessIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOriginAccessIdentity(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OriginAccessIdentityState, opts ...pulumi.ResourceOpt) (*OriginAccessIdentity, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["callerReference"] = state.CallerReference
		inputs["cloudfrontAccessIdentityPath"] = state.CloudfrontAccessIdentityPath
		inputs["comment"] = state.Comment
		inputs["etag"] = state.Etag
		inputs["iamArn"] = state.IamArn
		inputs["s3CanonicalUserId"] = state.S3CanonicalUserId
	}
	s, err := ctx.ReadResource("aws:cloudfront/originAccessIdentity:OriginAccessIdentity", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OriginAccessIdentity{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OriginAccessIdentity) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OriginAccessIdentity) ID() *pulumi.IDOutput {
	return r.s.ID
}

// Internal value used by CloudFront to allow future
// updates to the origin access identity.
func (r *OriginAccessIdentity) CallerReference() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["callerReference"])
}

// A shortcut to the full path for the
// origin access identity to use in CloudFront, see below.
func (r *OriginAccessIdentity) CloudfrontAccessIdentityPath() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cloudfrontAccessIdentityPath"])
}

// An optional comment for the origin access identity.
func (r *OriginAccessIdentity) Comment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["comment"])
}

// The current version of the origin access identity's information.
// For example: `E2QWRUHAPOMQZL`.
func (r *OriginAccessIdentity) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// A pre-generated ARN for use in S3 bucket policies (see below).
// Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
// E2QWRUHAPOMQZL`.
func (r *OriginAccessIdentity) IamArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["iamArn"])
}

// The Amazon S3 canonical user ID for the origin
// access identity, which you use when giving the origin access identity read
// permission to an object in Amazon S3.
func (r *OriginAccessIdentity) S3CanonicalUserId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["s3CanonicalUserId"])
}

// Input properties used for looking up and filtering OriginAccessIdentity resources.
type OriginAccessIdentityState struct {
	// Internal value used by CloudFront to allow future
	// updates to the origin access identity.
	CallerReference interface{}
	// A shortcut to the full path for the
	// origin access identity to use in CloudFront, see below.
	CloudfrontAccessIdentityPath interface{}
	// An optional comment for the origin access identity.
	Comment interface{}
	// The current version of the origin access identity's information.
	// For example: `E2QWRUHAPOMQZL`.
	Etag interface{}
	// A pre-generated ARN for use in S3 bucket policies (see below).
	// Example: `arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity
	// E2QWRUHAPOMQZL`.
	IamArn interface{}
	// The Amazon S3 canonical user ID for the origin
	// access identity, which you use when giving the origin access identity read
	// permission to an object in Amazon S3.
	S3CanonicalUserId interface{}
}

// The set of arguments for constructing a OriginAccessIdentity resource.
type OriginAccessIdentityArgs struct {
	// An optional comment for the origin access identity.
	Comment interface{}
}
