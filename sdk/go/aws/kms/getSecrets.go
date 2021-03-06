// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Decrypt multiple secrets from data encrypted with the AWS KMS service.
// 
// ~> **NOTE**: Using this data provider will allow you to conceal secret data within your resource definitions but does not take care of protecting that data in all Terraform logging and state output. Please take care to secure your secret data beyond just the Terraform configuration.
func LookupSecrets(ctx *pulumi.Context, args *GetSecretsArgs) (*GetSecretsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["secrets"] = args.Secrets
	}
	outputs, err := ctx.Invoke("aws:kms/getSecrets:getSecrets", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSecretsResult{
		Plaintext: outputs["plaintext"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSecrets.
type GetSecretsArgs struct {
	// One or more encrypted payload definitions from the KMS service. See the Secret Definitions below.
	Secrets interface{}
}

// A collection of values returned by getSecrets.
type GetSecretsResult struct {
	// Map containing each `secret` `name` as the key with its decrypted plaintext value
	Plaintext interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
