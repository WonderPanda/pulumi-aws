// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a DMS (Data Migration Service) replication task resource. DMS replication tasks can be created, updated, deleted, and imported.
type ReplicationTask struct {
	s *pulumi.ResourceState
}

// NewReplicationTask registers a new resource with the given unique name, arguments, and options.
func NewReplicationTask(ctx *pulumi.Context,
	name string, args *ReplicationTaskArgs, opts ...pulumi.ResourceOpt) (*ReplicationTask, error) {
	if args == nil || args.MigrationType == nil {
		return nil, errors.New("missing required argument 'MigrationType'")
	}
	if args == nil || args.ReplicationInstanceArn == nil {
		return nil, errors.New("missing required argument 'ReplicationInstanceArn'")
	}
	if args == nil || args.ReplicationTaskId == nil {
		return nil, errors.New("missing required argument 'ReplicationTaskId'")
	}
	if args == nil || args.SourceEndpointArn == nil {
		return nil, errors.New("missing required argument 'SourceEndpointArn'")
	}
	if args == nil || args.TableMappings == nil {
		return nil, errors.New("missing required argument 'TableMappings'")
	}
	if args == nil || args.TargetEndpointArn == nil {
		return nil, errors.New("missing required argument 'TargetEndpointArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cdcStartTime"] = nil
		inputs["migrationType"] = nil
		inputs["replicationInstanceArn"] = nil
		inputs["replicationTaskId"] = nil
		inputs["replicationTaskSettings"] = nil
		inputs["sourceEndpointArn"] = nil
		inputs["tableMappings"] = nil
		inputs["tags"] = nil
		inputs["targetEndpointArn"] = nil
	} else {
		inputs["cdcStartTime"] = args.CdcStartTime
		inputs["migrationType"] = args.MigrationType
		inputs["replicationInstanceArn"] = args.ReplicationInstanceArn
		inputs["replicationTaskId"] = args.ReplicationTaskId
		inputs["replicationTaskSettings"] = args.ReplicationTaskSettings
		inputs["sourceEndpointArn"] = args.SourceEndpointArn
		inputs["tableMappings"] = args.TableMappings
		inputs["tags"] = args.Tags
		inputs["targetEndpointArn"] = args.TargetEndpointArn
	}
	inputs["replicationTaskArn"] = nil
	s, err := ctx.RegisterResource("aws:dms/replicationTask:ReplicationTask", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ReplicationTask{s: s}, nil
}

// GetReplicationTask gets an existing ReplicationTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationTask(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ReplicationTaskState, opts ...pulumi.ResourceOpt) (*ReplicationTask, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cdcStartTime"] = state.CdcStartTime
		inputs["migrationType"] = state.MigrationType
		inputs["replicationInstanceArn"] = state.ReplicationInstanceArn
		inputs["replicationTaskArn"] = state.ReplicationTaskArn
		inputs["replicationTaskId"] = state.ReplicationTaskId
		inputs["replicationTaskSettings"] = state.ReplicationTaskSettings
		inputs["sourceEndpointArn"] = state.SourceEndpointArn
		inputs["tableMappings"] = state.TableMappings
		inputs["tags"] = state.Tags
		inputs["targetEndpointArn"] = state.TargetEndpointArn
	}
	s, err := ctx.ReadResource("aws:dms/replicationTask:ReplicationTask", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ReplicationTask{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ReplicationTask) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ReplicationTask) ID() *pulumi.IDOutput {
	return r.s.ID
}

// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
func (r *ReplicationTask) CdcStartTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cdcStartTime"])
}

// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
func (r *ReplicationTask) MigrationType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["migrationType"])
}

// The Amazon Resource Name (ARN) of the replication instance.
func (r *ReplicationTask) ReplicationInstanceArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["replicationInstanceArn"])
}

// The Amazon Resource Name (ARN) for the replication task.
func (r *ReplicationTask) ReplicationTaskArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["replicationTaskArn"])
}

// The replication task identifier.
func (r *ReplicationTask) ReplicationTaskId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["replicationTaskId"])
}

// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
func (r *ReplicationTask) ReplicationTaskSettings() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["replicationTaskSettings"])
}

// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
func (r *ReplicationTask) SourceEndpointArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceEndpointArn"])
}

// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
func (r *ReplicationTask) TableMappings() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tableMappings"])
}

// A mapping of tags to assign to the resource.
func (r *ReplicationTask) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
func (r *ReplicationTask) TargetEndpointArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["targetEndpointArn"])
}

// Input properties used for looking up and filtering ReplicationTask resources.
type ReplicationTaskState struct {
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime interface{}
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType interface{}
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn interface{}
	// The Amazon Resource Name (ARN) for the replication task.
	ReplicationTaskArn interface{}
	// The replication task identifier.
	ReplicationTaskId interface{}
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings interface{}
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn interface{}
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn interface{}
}

// The set of arguments for constructing a ReplicationTask resource.
type ReplicationTaskArgs struct {
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime interface{}
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType interface{}
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn interface{}
	// The replication task identifier.
	ReplicationTaskId interface{}
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings interface{}
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn interface{}
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn interface{}
}
