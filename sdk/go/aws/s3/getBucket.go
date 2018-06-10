// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides details about a specific S3 bucket.
// 
// This resource may prove useful when setting up a Route53 record, or an origin for a CloudFront
// Distribution.
func Lookupucket(ctx *pulumi.Context, args *GetBucketArgs) (*GetBucketResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["bucket"] = args.Bucket
	}
	outputs, err := ctx.Invoke("aws:s3/getBucket:getBucket", inputs)
	if err != nil {
		return nil, err
	}
	ret := GetBucketResult{}
	if v, ok := outputs["arn"]; ok {
		ret.Arn = v
	}
	if v, ok := outputs["bucketDomainName"]; ok {
		ret.BucketDomainName = v
	}
	if v, ok := outputs["hostedZoneId"]; ok {
		ret.HostedZoneId = v
	}
	if v, ok := outputs["region"]; ok {
		ret.Region = v
	}
	if v, ok := outputs["websiteDomain"]; ok {
		ret.WebsiteDomain = v
	}
	if v, ok := outputs["websiteEndpoint"]; ok {
		ret.WebsiteEndpoint = v
	}
	return &ret, nil
}

// A collection of arguments for invoking getBucket.
type GetBucketArgs struct {
	// The name of the bucket
	Bucket interface{}
}

// A collection of values returned by getBucket.
type GetBucketResult struct {
	// The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
	Arn interface{}
	// The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
	BucketDomainName interface{}
	// The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
	HostedZoneId interface{}
	// The AWS region this bucket resides in.
	Region interface{}
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
	WebsiteDomain interface{}
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint interface{}
}
