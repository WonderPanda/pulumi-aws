// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a VPC/Subnet/ENI Flow Log to capture IP traffic for a specific network
// interface, subnet, or VPC. Logs are sent to a CloudWatch Log Group.
type FlowLog struct {
	s *pulumi.ResourceState
}

// NewFlowLog registers a new resource with the given unique name, arguments, and options.
func NewFlowLog(ctx *pulumi.Context,
	name string, args *FlowLogArgs, opts ...pulumi.ResourceOpt) (*FlowLog, error) {
	if args == nil || args.IamRoleArn == nil {
		return nil, errors.New("missing required argument 'IamRoleArn'")
	}
	if args == nil || args.LogGroupName == nil {
		return nil, errors.New("missing required argument 'LogGroupName'")
	}
	if args == nil || args.TrafficType == nil {
		return nil, errors.New("missing required argument 'TrafficType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["eniId"] = nil
		inputs["iamRoleArn"] = nil
		inputs["logGroupName"] = nil
		inputs["subnetId"] = nil
		inputs["trafficType"] = nil
		inputs["vpcId"] = nil
	} else {
		inputs["eniId"] = args.EniId
		inputs["iamRoleArn"] = args.IamRoleArn
		inputs["logGroupName"] = args.LogGroupName
		inputs["subnetId"] = args.SubnetId
		inputs["trafficType"] = args.TrafficType
		inputs["vpcId"] = args.VpcId
	}
	s, err := ctx.RegisterResource("aws:ec2/flowLog:FlowLog", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FlowLog{s: s}, nil
}

// GetFlowLog gets an existing FlowLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowLog(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FlowLogState, opts ...pulumi.ResourceOpt) (*FlowLog, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["eniId"] = state.EniId
		inputs["iamRoleArn"] = state.IamRoleArn
		inputs["logGroupName"] = state.LogGroupName
		inputs["subnetId"] = state.SubnetId
		inputs["trafficType"] = state.TrafficType
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("aws:ec2/flowLog:FlowLog", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FlowLog{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FlowLog) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FlowLog) ID() *pulumi.IDOutput {
	return r.s.ID
}

// Elastic Network Interface ID to attach to
func (r *FlowLog) EniId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["eniId"])
}

// The ARN for the IAM role that's used to post flow
// logs to a CloudWatch Logs log group
func (r *FlowLog) IamRoleArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["iamRoleArn"])
}

// The name of the CloudWatch log group
func (r *FlowLog) LogGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["logGroupName"])
}

// Subnet ID to attach to
func (r *FlowLog) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// The type of traffic to capture. Valid values:
// `ACCEPT`,`REJECT`, `ALL`
func (r *FlowLog) TrafficType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["trafficType"])
}

// VPC ID to attach to
func (r *FlowLog) VpcId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering FlowLog resources.
type FlowLogState struct {
	// Elastic Network Interface ID to attach to
	EniId interface{}
	// The ARN for the IAM role that's used to post flow
	// logs to a CloudWatch Logs log group
	IamRoleArn interface{}
	// The name of the CloudWatch log group
	LogGroupName interface{}
	// Subnet ID to attach to
	SubnetId interface{}
	// The type of traffic to capture. Valid values:
	// `ACCEPT`,`REJECT`, `ALL`
	TrafficType interface{}
	// VPC ID to attach to
	VpcId interface{}
}

// The set of arguments for constructing a FlowLog resource.
type FlowLogArgs struct {
	// Elastic Network Interface ID to attach to
	EniId interface{}
	// The ARN for the IAM role that's used to post flow
	// logs to a CloudWatch Logs log group
	IamRoleArn interface{}
	// The name of the CloudWatch log group
	LogGroupName interface{}
	// Subnet ID to attach to
	SubnetId interface{}
	// The type of traffic to capture. Valid values:
	// `ACCEPT`,`REJECT`, `ALL`
	TrafficType interface{}
	// VPC ID to attach to
	VpcId interface{}
}
