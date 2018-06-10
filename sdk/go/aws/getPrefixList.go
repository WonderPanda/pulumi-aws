// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `aws_prefix_list` provides details about a specific prefix list (PL)
// in the current region.
// 
// This can be used both to validate a prefix list given in a variable
// and to obtain the CIDR blocks (IP address ranges) for the associated
// AWS service. The latter may be useful e.g. for adding network ACL
// rules.
func LookuprefixList(ctx *pulumi.Context, args *GetPrefixListArgs) (*GetPrefixListResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["prefixListId"] = args.PrefixListId
	}
	outputs, err := ctx.Invoke("aws:index/getPrefixList:getPrefixList", inputs)
	if err != nil {
		return nil, err
	}
	ret := GetPrefixListResult{}
	if v, ok := outputs["cidrBlocks"]; ok {
		ret.CidrBlocks = v
	}
	if v, ok := outputs["name"]; ok {
		ret.Name = v
	}
	return &ret, nil
}

// A collection of arguments for invoking getPrefixList.
type GetPrefixListArgs struct {
	// The name of the prefix list to select.
	Name interface{}
	// The ID of the prefix list to select.
	PrefixListId interface{}
}

// A collection of values returned by getPrefixList.
type GetPrefixListResult struct {
	// The list of CIDR blocks for the AWS service associated
	// with the prefix list.
	CidrBlocks interface{}
	// The name of the selected prefix list.
	Name interface{}
}
