// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Lambda event source mapping. This allows Lambda functions to get events from Kinesis and DynamoDB.
// 
// For information about Lambda and how to use it, see [What is AWS Lambda?][1]
// For information about event source mappings, see [CreateEventSourceMapping][2] in the API docs.
type EventSourceMapping struct {
	s *pulumi.ResourceState
}

// NewEventSourceMapping registers a new resource with the given unique name, arguments, and options.
func NewEventSourceMapping(ctx *pulumi.Context,
	name string, args *EventSourceMappingArgs, opts ...pulumi.ResourceOpt) (*EventSourceMapping, error) {
	if args == nil || args.EventSourceArn == nil {
		return nil, errors.New("missing required argument 'EventSourceArn'")
	}
	if args == nil || args.FunctionName == nil {
		return nil, errors.New("missing required argument 'FunctionName'")
	}
	if args == nil || args.StartingPosition == nil {
		return nil, errors.New("missing required argument 'StartingPosition'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["batchSize"] = nil
		inputs["enabled"] = nil
		inputs["eventSourceArn"] = nil
		inputs["functionName"] = nil
		inputs["startingPosition"] = nil
	} else {
		inputs["batchSize"] = args.BatchSize
		inputs["enabled"] = args.Enabled
		inputs["eventSourceArn"] = args.EventSourceArn
		inputs["functionName"] = args.FunctionName
		inputs["startingPosition"] = args.StartingPosition
	}
	inputs["functionArn"] = nil
	inputs["lastModified"] = nil
	inputs["lastProcessingResult"] = nil
	inputs["state"] = nil
	inputs["stateTransitionReason"] = nil
	inputs["uuid"] = nil
	s, err := ctx.RegisterResource("aws:lambda/eventSourceMapping:EventSourceMapping", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EventSourceMapping{s: s}, nil
}

// GetEventSourceMapping gets an existing EventSourceMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventSourceMapping(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EventSourceMappingState, opts ...pulumi.ResourceOpt) (*EventSourceMapping, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["batchSize"] = state.BatchSize
		inputs["enabled"] = state.Enabled
		inputs["eventSourceArn"] = state.EventSourceArn
		inputs["functionArn"] = state.FunctionArn
		inputs["functionName"] = state.FunctionName
		inputs["lastModified"] = state.LastModified
		inputs["lastProcessingResult"] = state.LastProcessingResult
		inputs["startingPosition"] = state.StartingPosition
		inputs["state"] = state.State
		inputs["stateTransitionReason"] = state.StateTransitionReason
		inputs["uuid"] = state.Uuid
	}
	s, err := ctx.ReadResource("aws:lambda/eventSourceMapping:EventSourceMapping", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EventSourceMapping{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *EventSourceMapping) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *EventSourceMapping) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100`.
func (r *EventSourceMapping) BatchSize() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["batchSize"])
}

// Determines if the mapping will be enabled on creation. Defaults to `true`.
func (r *EventSourceMapping) Enabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enabled"])
}

// The event source ARN - can either be a Kinesis or DynamoDB stream.
func (r *EventSourceMapping) EventSourceArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["eventSourceArn"])
}

// The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from `function_name` above.)
func (r *EventSourceMapping) FunctionArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["functionArn"])
}

// The name or the ARN of the Lambda function that will be subscribing to events.
func (r *EventSourceMapping) FunctionName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["functionName"])
}

// The date this resource was last modified.
func (r *EventSourceMapping) LastModified() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lastModified"])
}

// The result of the last AWS Lambda invocation of your Lambda function.
func (r *EventSourceMapping) LastProcessingResult() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lastProcessingResult"])
}

// The position in the stream where AWS Lambda should start reading. Can be one of either `TRIM_HORIZON` or `LATEST`.
func (r *EventSourceMapping) StartingPosition() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["startingPosition"])
}

// The state of the event source mapping.
func (r *EventSourceMapping) State() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["state"])
}

// The reason the event source mapping is in its current state.
func (r *EventSourceMapping) StateTransitionReason() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["stateTransitionReason"])
}

// The UUID of the created event source mapping.
func (r *EventSourceMapping) Uuid() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["uuid"])
}

// Input properties used for looking up and filtering EventSourceMapping resources.
type EventSourceMappingState struct {
	// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100`.
	BatchSize interface{}
	// Determines if the mapping will be enabled on creation. Defaults to `true`.
	Enabled interface{}
	// The event source ARN - can either be a Kinesis or DynamoDB stream.
	EventSourceArn interface{}
	// The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from `function_name` above.)
	FunctionArn interface{}
	// The name or the ARN of the Lambda function that will be subscribing to events.
	FunctionName interface{}
	// The date this resource was last modified.
	LastModified interface{}
	// The result of the last AWS Lambda invocation of your Lambda function.
	LastProcessingResult interface{}
	// The position in the stream where AWS Lambda should start reading. Can be one of either `TRIM_HORIZON` or `LATEST`.
	StartingPosition interface{}
	// The state of the event source mapping.
	State interface{}
	// The reason the event source mapping is in its current state.
	StateTransitionReason interface{}
	// The UUID of the created event source mapping.
	Uuid interface{}
}

// The set of arguments for constructing a EventSourceMapping resource.
type EventSourceMappingArgs struct {
	// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100`.
	BatchSize interface{}
	// Determines if the mapping will be enabled on creation. Defaults to `true`.
	Enabled interface{}
	// The event source ARN - can either be a Kinesis or DynamoDB stream.
	EventSourceArn interface{}
	// The name or the ARN of the Lambda function that will be subscribing to events.
	FunctionName interface{}
	// The position in the stream where AWS Lambda should start reading. Can be one of either `TRIM_HORIZON` or `LATEST`.
	StartingPosition interface{}
}
