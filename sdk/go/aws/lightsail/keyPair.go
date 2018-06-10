// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Lightsail Key Pair, for use with Lightsail Instances. These key pairs
// are seperate from EC2 Key Pairs, and must be created or imported for use with
// Lightsail.
// 
// ~> **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
type KeyPair struct {
	s *pulumi.ResourceState
}

// NewKeyPair registers a new resource with the given unique name, arguments, and options.
func NewKeyPair(ctx *pulumi.Context,
	name string, args *KeyPairArgs, opts ...pulumi.ResourceOpt) (*KeyPair, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["namePrefix"] = nil
		inputs["pgpKey"] = nil
		inputs["publicKey"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["pgpKey"] = args.PgpKey
		inputs["publicKey"] = args.PublicKey
	}
	inputs["arn"] = nil
	inputs["encryptedFingerprint"] = nil
	inputs["encryptedPrivateKey"] = nil
	inputs["fingerprint"] = nil
	inputs["privateKey"] = nil
	s, err := ctx.RegisterResource("aws:lightsail/keyPair:KeyPair", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KeyPair{s: s}, nil
}

// GetKeyPair gets an existing KeyPair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyPair(ctx *pulumi.Context,
	name string, id pulumi.ID, state *KeyPairState, opts ...pulumi.ResourceOpt) (*KeyPair, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["encryptedFingerprint"] = state.EncryptedFingerprint
		inputs["encryptedPrivateKey"] = state.EncryptedPrivateKey
		inputs["fingerprint"] = state.Fingerprint
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["pgpKey"] = state.PgpKey
		inputs["privateKey"] = state.PrivateKey
		inputs["publicKey"] = state.PublicKey
	}
	s, err := ctx.ReadResource("aws:lightsail/keyPair:KeyPair", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KeyPair{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *KeyPair) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *KeyPair) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The ARN of the Lightsail key pair
func (r *KeyPair) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The MD5 public key fingerprint for the encrypted
// private key
func (r *KeyPair) EncryptedFingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["encryptedFingerprint"])
}

// the private key material, base 64 encoded and
// encrypted with the given `pgp_key`. This is only populated when creating a new
// key and `pgp_key` is supplied
func (r *KeyPair) EncryptedPrivateKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["encryptedPrivateKey"])
}

// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
func (r *KeyPair) Fingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fingerprint"])
}

// The name of the Lightsail Key Pair. If omitted, a unique
// name will be generated by Terraform
func (r *KeyPair) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *KeyPair) NamePrefix() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["namePrefix"])
}

// An optional PGP key to encrypt the resulting private
// key material. Only used when creating a new key pair
func (r *KeyPair) PgpKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["pgpKey"])
}

// the private key, base64 encoded. This is only populated
// when creating a new key, and when no `pgp_key` is provided
func (r *KeyPair) PrivateKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateKey"])
}

// The public key material. This public key will be
// imported into Lightsail
func (r *KeyPair) PublicKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publicKey"])
}

// Input properties used for looking up and filtering KeyPair resources.
type KeyPairState struct {
	// The ARN of the Lightsail key pair
	Arn interface{}
	// The MD5 public key fingerprint for the encrypted
	// private key
	EncryptedFingerprint interface{}
	// the private key material, base 64 encoded and
	// encrypted with the given `pgp_key`. This is only populated when creating a new
	// key and `pgp_key` is supplied
	EncryptedPrivateKey interface{}
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	Fingerprint interface{}
	// The name of the Lightsail Key Pair. If omitted, a unique
	// name will be generated by Terraform
	Name interface{}
	NamePrefix interface{}
	// An optional PGP key to encrypt the resulting private
	// key material. Only used when creating a new key pair
	PgpKey interface{}
	// the private key, base64 encoded. This is only populated
	// when creating a new key, and when no `pgp_key` is provided
	PrivateKey interface{}
	// The public key material. This public key will be
	// imported into Lightsail
	PublicKey interface{}
}

// The set of arguments for constructing a KeyPair resource.
type KeyPairArgs struct {
	// The name of the Lightsail Key Pair. If omitted, a unique
	// name will be generated by Terraform
	Name interface{}
	NamePrefix interface{}
	// An optional PGP key to encrypt the resulting private
	// key material. Only used when creating a new key pair
	PgpKey interface{}
	// The public key material. This public key will be
	// imported into Lightsail
	PublicKey interface{}
}
