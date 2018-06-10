// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a S3 bucket [metrics configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html) resource.
type BucketMetric struct {
	s *pulumi.ResourceState
}

// NewBucketMetric registers a new resource with the given unique name, arguments, and options.
func NewBucketMetric(ctx *pulumi.Context,
	name string, args *BucketMetricArgs, opts ...pulumi.ResourceOpt) (*BucketMetric, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["bucket"] = nil
		inputs["filter"] = nil
		inputs["name"] = nil
	} else {
		inputs["bucket"] = args.Bucket
		inputs["filter"] = args.Filter
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("aws:s3/bucketMetric:BucketMetric", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BucketMetric{s: s}, nil
}

// GetBucketMetric gets an existing BucketMetric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketMetric(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BucketMetricState, opts ...pulumi.ResourceOpt) (*BucketMetric, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["bucket"] = state.Bucket
		inputs["filter"] = state.Filter
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("aws:s3/bucketMetric:BucketMetric", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BucketMetric{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *BucketMetric) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *BucketMetric) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The name of the bucket to put metric configuration.
func (r *BucketMetric) Bucket() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["bucket"])
}

// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
func (r *BucketMetric) Filter() *pulumi.Output {
	return r.s.State["filter"]
}

// Unique identifier of the metrics configuration for the bucket.
func (r *BucketMetric) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering BucketMetric resources.
type BucketMetricState struct {
	// The name of the bucket to put metric configuration.
	Bucket interface{}
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter interface{}
	// Unique identifier of the metrics configuration for the bucket.
	Name interface{}
}

// The set of arguments for constructing a BucketMetric resource.
type BucketMetricArgs struct {
	// The name of the bucket to put metric configuration.
	Bucket interface{}
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter interface{}
	// Unique identifier of the metrics configuration for the bucket.
	Name interface{}
}
