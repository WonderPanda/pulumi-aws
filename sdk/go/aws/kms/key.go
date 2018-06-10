// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a KMS customer master key.
type Key struct {
	s *pulumi.ResourceState
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOpt) (*Key, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["deletionWindowInDays"] = nil
		inputs["description"] = nil
		inputs["enableKeyRotation"] = nil
		inputs["isEnabled"] = nil
		inputs["keyUsage"] = nil
		inputs["policy"] = nil
		inputs["tags"] = nil
	} else {
		inputs["deletionWindowInDays"] = args.DeletionWindowInDays
		inputs["description"] = args.Description
		inputs["enableKeyRotation"] = args.EnableKeyRotation
		inputs["isEnabled"] = args.IsEnabled
		inputs["keyUsage"] = args.KeyUsage
		inputs["policy"] = args.Policy
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	inputs["keyId"] = nil
	s, err := ctx.RegisterResource("aws:kms/key:Key", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Key{s: s}, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *KeyState, opts ...pulumi.ResourceOpt) (*Key, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["deletionWindowInDays"] = state.DeletionWindowInDays
		inputs["description"] = state.Description
		inputs["enableKeyRotation"] = state.EnableKeyRotation
		inputs["isEnabled"] = state.IsEnabled
		inputs["keyId"] = state.KeyId
		inputs["keyUsage"] = state.KeyUsage
		inputs["policy"] = state.Policy
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:kms/key:Key", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Key{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Key) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Key) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The Amazon Resource Name (ARN) of the key.
func (r *Key) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Duration in days after which the key is deleted
// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
func (r *Key) DeletionWindowInDays() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["deletionWindowInDays"])
}

// The description of the key as viewed in AWS console.
func (r *Key) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
// is enabled. Defaults to false.
func (r *Key) EnableKeyRotation() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableKeyRotation"])
}

// Specifies whether the key is enabled. Defaults to true.
func (r *Key) IsEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["isEnabled"])
}

// The globally unique identifier for the key.
func (r *Key) KeyId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyId"])
}

// Specifies the intended use of the key.
// Defaults to ENCRYPT_DECRYPT, and only symmetric encryption and decryption are supported.
func (r *Key) KeyUsage() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyUsage"])
}

// A valid policy JSON document.
func (r *Key) Policy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policy"])
}

// A mapping of tags to assign to the object.
func (r *Key) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Key resources.
type KeyState struct {
	// The Amazon Resource Name (ARN) of the key.
	Arn interface{}
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays interface{}
	// The description of the key as viewed in AWS console.
	Description interface{}
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
	// is enabled. Defaults to false.
	EnableKeyRotation interface{}
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled interface{}
	// The globally unique identifier for the key.
	KeyId interface{}
	// Specifies the intended use of the key.
	// Defaults to ENCRYPT_DECRYPT, and only symmetric encryption and decryption are supported.
	KeyUsage interface{}
	// A valid policy JSON document.
	Policy interface{}
	// A mapping of tags to assign to the object.
	Tags interface{}
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays interface{}
	// The description of the key as viewed in AWS console.
	Description interface{}
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
	// is enabled. Defaults to false.
	EnableKeyRotation interface{}
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled interface{}
	// Specifies the intended use of the key.
	// Defaults to ENCRYPT_DECRYPT, and only symmetric encryption and decryption are supported.
	KeyUsage interface{}
	// A valid policy JSON document.
	Policy interface{}
	// A mapping of tags to assign to the object.
	Tags interface{}
}
