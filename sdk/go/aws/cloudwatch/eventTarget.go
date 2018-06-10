// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CloudWatch Event Target resource.
type EventTarget struct {
	s *pulumi.ResourceState
}

// NewEventTarget registers a new resource with the given unique name, arguments, and options.
func NewEventTarget(ctx *pulumi.Context,
	name string, args *EventTargetArgs, opts ...pulumi.ResourceOpt) (*EventTarget, error) {
	if args == nil || args.Arn == nil {
		return nil, errors.New("missing required argument 'Arn'")
	}
	if args == nil || args.Rule == nil {
		return nil, errors.New("missing required argument 'Rule'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["arn"] = nil
		inputs["batchTarget"] = nil
		inputs["ecsTarget"] = nil
		inputs["input"] = nil
		inputs["inputPath"] = nil
		inputs["inputTransformer"] = nil
		inputs["kinesisTarget"] = nil
		inputs["roleArn"] = nil
		inputs["rule"] = nil
		inputs["runCommandTargets"] = nil
		inputs["sqsTarget"] = nil
		inputs["targetId"] = nil
	} else {
		inputs["arn"] = args.Arn
		inputs["batchTarget"] = args.BatchTarget
		inputs["ecsTarget"] = args.EcsTarget
		inputs["input"] = args.Input
		inputs["inputPath"] = args.InputPath
		inputs["inputTransformer"] = args.InputTransformer
		inputs["kinesisTarget"] = args.KinesisTarget
		inputs["roleArn"] = args.RoleArn
		inputs["rule"] = args.Rule
		inputs["runCommandTargets"] = args.RunCommandTargets
		inputs["sqsTarget"] = args.SqsTarget
		inputs["targetId"] = args.TargetId
	}
	s, err := ctx.RegisterResource("aws:cloudwatch/eventTarget:EventTarget", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EventTarget{s: s}, nil
}

// GetEventTarget gets an existing EventTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventTarget(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EventTargetState, opts ...pulumi.ResourceOpt) (*EventTarget, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["batchTarget"] = state.BatchTarget
		inputs["ecsTarget"] = state.EcsTarget
		inputs["input"] = state.Input
		inputs["inputPath"] = state.InputPath
		inputs["inputTransformer"] = state.InputTransformer
		inputs["kinesisTarget"] = state.KinesisTarget
		inputs["roleArn"] = state.RoleArn
		inputs["rule"] = state.Rule
		inputs["runCommandTargets"] = state.RunCommandTargets
		inputs["sqsTarget"] = state.SqsTarget
		inputs["targetId"] = state.TargetId
	}
	s, err := ctx.ReadResource("aws:cloudwatch/eventTarget:EventTarget", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EventTarget{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *EventTarget) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *EventTarget) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The Amazon Resource Name (ARN) associated of the target.
func (r *EventTarget) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
func (r *EventTarget) BatchTarget() *pulumi.Output {
	return r.s.State["batchTarget"]
}

// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
func (r *EventTarget) EcsTarget() *pulumi.Output {
	return r.s.State["ecsTarget"]
}

// Valid JSON text passed to the target.
func (r *EventTarget) Input() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["input"])
}

// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
// that is used for extracting part of the matched event when passing it to the target.
func (r *EventTarget) InputPath() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["inputPath"])
}

// Parameters used when you are providing a custom input to a target based on certain event data.
func (r *EventTarget) InputTransformer() *pulumi.Output {
	return r.s.State["inputTransformer"]
}

// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
func (r *EventTarget) KinesisTarget() *pulumi.Output {
	return r.s.State["kinesisTarget"]
}

// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used.
func (r *EventTarget) RoleArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["roleArn"])
}

// The name of the rule you want to add targets to.
func (r *EventTarget) Rule() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["rule"])
}

// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
func (r *EventTarget) RunCommandTargets() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["runCommandTargets"])
}

// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
func (r *EventTarget) SqsTarget() *pulumi.Output {
	return r.s.State["sqsTarget"]
}

// The unique target assignment ID.  If missing, will generate a random, unique id.
func (r *EventTarget) TargetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["targetId"])
}

// Input properties used for looking up and filtering EventTarget resources.
type EventTargetState struct {
	// The Amazon Resource Name (ARN) associated of the target.
	Arn interface{}
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget interface{}
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget interface{}
	// Valid JSON text passed to the target.
	Input interface{}
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath interface{}
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer interface{}
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget interface{}
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used.
	RoleArn interface{}
	// The name of the rule you want to add targets to.
	Rule interface{}
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets interface{}
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget interface{}
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId interface{}
}

// The set of arguments for constructing a EventTarget resource.
type EventTargetArgs struct {
	// The Amazon Resource Name (ARN) associated of the target.
	Arn interface{}
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget interface{}
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget interface{}
	// Valid JSON text passed to the target.
	Input interface{}
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath interface{}
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer interface{}
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget interface{}
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used.
	RoleArn interface{}
	// The name of the rule you want to add targets to.
	Rule interface{}
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets interface{}
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget interface{}
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId interface{}
}
