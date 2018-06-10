// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides the ability to register instances and containers with a LB
// target group
// 
// ~> **Note:** `aws_alb_target_group_attachment` is known as `aws_lb_target_group_attachment`. The functionality is identical.
type TargetGroupAttachment struct {
	s *pulumi.ResourceState
}

// NewTargetGroupAttachment registers a new resource with the given unique name, arguments, and options.
func NewTargetGroupAttachment(ctx *pulumi.Context,
	name string, args *TargetGroupAttachmentArgs, opts ...pulumi.ResourceOpt) (*TargetGroupAttachment, error) {
	if args == nil || args.TargetGroupArn == nil {
		return nil, errors.New("missing required argument 'TargetGroupArn'")
	}
	if args == nil || args.TargetId == nil {
		return nil, errors.New("missing required argument 'TargetId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["availabilityZone"] = nil
		inputs["port"] = nil
		inputs["targetGroupArn"] = nil
		inputs["targetId"] = nil
	} else {
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["port"] = args.Port
		inputs["targetGroupArn"] = args.TargetGroupArn
		inputs["targetId"] = args.TargetId
	}
	s, err := ctx.RegisterResource("aws:elasticloadbalancingv2/targetGroupAttachment:TargetGroupAttachment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TargetGroupAttachment{s: s}, nil
}

// GetTargetGroupAttachment gets an existing TargetGroupAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetGroupAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TargetGroupAttachmentState, opts ...pulumi.ResourceOpt) (*TargetGroupAttachment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["port"] = state.Port
		inputs["targetGroupArn"] = state.TargetGroupArn
		inputs["targetId"] = state.TargetId
	}
	s, err := ctx.ReadResource("aws:elasticloadbalancingv2/targetGroupAttachment:TargetGroupAttachment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TargetGroupAttachment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TargetGroupAttachment) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TargetGroupAttachment) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The Availability Zone where the IP address of the target is to be registered.
func (r *TargetGroupAttachment) AvailabilityZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["availabilityZone"])
}

// The port on which targets receive traffic.
func (r *TargetGroupAttachment) Port() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["port"])
}

// The ARN of the target group with which to register targets
func (r *TargetGroupAttachment) TargetGroupArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["targetGroupArn"])
}

// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address.
func (r *TargetGroupAttachment) TargetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["targetId"])
}

// Input properties used for looking up and filtering TargetGroupAttachment resources.
type TargetGroupAttachmentState struct {
	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone interface{}
	// The port on which targets receive traffic.
	Port interface{}
	// The ARN of the target group with which to register targets
	TargetGroupArn interface{}
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address.
	TargetId interface{}
}

// The set of arguments for constructing a TargetGroupAttachment resource.
type TargetGroupAttachmentArgs struct {
	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone interface{}
	// The port on which targets receive traffic.
	Port interface{}
	// The ARN of the target group with which to register targets
	TargetGroupArn interface{}
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address.
	TargetId interface{}
}
