// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates and manages an AWS IoT certificate.
type Certificate struct {
	s *pulumi.ResourceState
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOpt) (*Certificate, error) {
	if args == nil || args.Active == nil {
		return nil, errors.New("missing required argument 'Active'")
	}
	if args == nil || args.Csr == nil {
		return nil, errors.New("missing required argument 'Csr'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["active"] = nil
		inputs["csr"] = nil
	} else {
		inputs["active"] = args.Active
		inputs["csr"] = args.Csr
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:iot/certificate:Certificate", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Certificate{s: s}, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CertificateState, opts ...pulumi.ResourceOpt) (*Certificate, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["active"] = state.Active
		inputs["arn"] = state.Arn
		inputs["csr"] = state.Csr
	}
	s, err := ctx.ReadResource("aws:iot/certificate:Certificate", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Certificate{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Certificate) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Certificate) ID() *pulumi.IDOutput {
	return r.s.ID
}

// Boolean flag to indicate if the certificate should be active
func (r *Certificate) Active() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["active"])
}

// The ARN of the created AWS IoT certificate
func (r *Certificate) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The certificate signing request. Review the
// [IoT API Reference Guide] (http://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
// for more information on creating a certificate from a certificate signing request (CSR).
func (r *Certificate) Csr() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["csr"])
}

// Input properties used for looking up and filtering Certificate resources.
type CertificateState struct {
	// Boolean flag to indicate if the certificate should be active
	Active interface{}
	// The ARN of the created AWS IoT certificate
	Arn interface{}
	// The certificate signing request. Review the
	// [IoT API Reference Guide] (http://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
	// for more information on creating a certificate from a certificate signing request (CSR).
	Csr interface{}
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Boolean flag to indicate if the certificate should be active
	Active interface{}
	// The certificate signing request. Review the
	// [IoT API Reference Guide] (http://docs.aws.amazon.com/iot/latest/apireference/API_CreateCertificateFromCsr.html)
	// for more information on creating a certificate from a certificate signing request (CSR).
	Csr interface{}
}
