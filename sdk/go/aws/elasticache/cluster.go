// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an ElastiCache Cluster resource, which manages a Memcached cluster or Redis instance.
// For working with Redis (Cluster Mode Enabled) replication groups, see the
// [`aws_elasticache_replication_group` resource](/docs/providers/aws/r/elasticache_replication_group.html).
// 
// ~> **Note:** When you change an attribute, such as `node_type`, by default
// it is applied in the next maintenance window. Because of this, Terraform may report
// a difference in its planning phase because the actual modification has not yet taken
// place. You can use the `apply_immediately` flag to instruct the service to apply the
// change immediately. Using `apply_immediately` can result in a brief downtime as the server reboots.
// See the AWS Docs on [Modifying an ElastiCache Cache Cluster][2] for more information.
type Cluster struct {
	s *pulumi.ResourceState
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["applyImmediately"] = nil
		inputs["availabilityZone"] = nil
		inputs["availabilityZones"] = nil
		inputs["azMode"] = nil
		inputs["clusterId"] = nil
		inputs["engine"] = nil
		inputs["engineVersion"] = nil
		inputs["maintenanceWindow"] = nil
		inputs["nodeType"] = nil
		inputs["notificationTopicArn"] = nil
		inputs["numCacheNodes"] = nil
		inputs["parameterGroupName"] = nil
		inputs["port"] = nil
		inputs["replicationGroupId"] = nil
		inputs["securityGroupIds"] = nil
		inputs["securityGroupNames"] = nil
		inputs["snapshotArns"] = nil
		inputs["snapshotName"] = nil
		inputs["snapshotRetentionLimit"] = nil
		inputs["snapshotWindow"] = nil
		inputs["subnetGroupName"] = nil
		inputs["tags"] = nil
	} else {
		inputs["applyImmediately"] = args.ApplyImmediately
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["availabilityZones"] = args.AvailabilityZones
		inputs["azMode"] = args.AzMode
		inputs["clusterId"] = args.ClusterId
		inputs["engine"] = args.Engine
		inputs["engineVersion"] = args.EngineVersion
		inputs["maintenanceWindow"] = args.MaintenanceWindow
		inputs["nodeType"] = args.NodeType
		inputs["notificationTopicArn"] = args.NotificationTopicArn
		inputs["numCacheNodes"] = args.NumCacheNodes
		inputs["parameterGroupName"] = args.ParameterGroupName
		inputs["port"] = args.Port
		inputs["replicationGroupId"] = args.ReplicationGroupId
		inputs["securityGroupIds"] = args.SecurityGroupIds
		inputs["securityGroupNames"] = args.SecurityGroupNames
		inputs["snapshotArns"] = args.SnapshotArns
		inputs["snapshotName"] = args.SnapshotName
		inputs["snapshotRetentionLimit"] = args.SnapshotRetentionLimit
		inputs["snapshotWindow"] = args.SnapshotWindow
		inputs["subnetGroupName"] = args.SubnetGroupName
		inputs["tags"] = args.Tags
	}
	inputs["cacheNodes"] = nil
	inputs["clusterAddress"] = nil
	inputs["configurationEndpoint"] = nil
	s, err := ctx.RegisterResource("aws:elasticache/cluster:Cluster", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterState, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["applyImmediately"] = state.ApplyImmediately
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["availabilityZones"] = state.AvailabilityZones
		inputs["azMode"] = state.AzMode
		inputs["cacheNodes"] = state.CacheNodes
		inputs["clusterAddress"] = state.ClusterAddress
		inputs["clusterId"] = state.ClusterId
		inputs["configurationEndpoint"] = state.ConfigurationEndpoint
		inputs["engine"] = state.Engine
		inputs["engineVersion"] = state.EngineVersion
		inputs["maintenanceWindow"] = state.MaintenanceWindow
		inputs["nodeType"] = state.NodeType
		inputs["notificationTopicArn"] = state.NotificationTopicArn
		inputs["numCacheNodes"] = state.NumCacheNodes
		inputs["parameterGroupName"] = state.ParameterGroupName
		inputs["port"] = state.Port
		inputs["replicationGroupId"] = state.ReplicationGroupId
		inputs["securityGroupIds"] = state.SecurityGroupIds
		inputs["securityGroupNames"] = state.SecurityGroupNames
		inputs["snapshotArns"] = state.SnapshotArns
		inputs["snapshotName"] = state.SnapshotName
		inputs["snapshotRetentionLimit"] = state.SnapshotRetentionLimit
		inputs["snapshotWindow"] = state.SnapshotWindow
		inputs["subnetGroupName"] = state.SubnetGroupName
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:elasticache/cluster:Cluster", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Cluster) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Cluster) ID() *pulumi.IDOutput {
	return r.s.ID
}

// Specifies whether any database modifications
// are applied immediately, or during the next maintenance window. Default is
// `false`. See [Amazon ElastiCache Documentation for more information.][1]
// (Available since v0.6.0)
func (r *Cluster) ApplyImmediately() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["applyImmediately"])
}

// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `availability_zones`
func (r *Cluster) AvailabilityZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["availabilityZone"])
}

// List of Availability Zones in which the cache nodes will be created. If you want to create cache nodes in single-az, use `availability_zone`
func (r *Cluster) AvailabilityZones() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["availabilityZones"])
}

// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `num_cache_nodes` must be greater than `1`
func (r *Cluster) AzMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["azMode"])
}

// List of node objects including `id`, `address`, `port` and `availability_zone`.
// Referenceable e.g. as `${aws_elasticache_cluster.bar.cache_nodes.0.address}`
func (r *Cluster) CacheNodes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["cacheNodes"])
}

// (Memcached only) The DNS name of the cache cluster without the port appended.
func (r *Cluster) ClusterAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterAddress"])
}

// Group identifier. ElastiCache converts
// this name to lowercase
func (r *Cluster) ClusterId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterId"])
}

// (Memcached only) The configuration endpoint to allow host discovery.
func (r *Cluster) ConfigurationEndpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["configurationEndpoint"])
}

// Name of the cache engine to be used for this cache cluster.
// Valid values for this parameter are `memcached` or `redis`
func (r *Cluster) Engine() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["engine"])
}

// Version number of the cache engine to be used.
// See [Selecting a Cache Engine and Version](https://docs.aws.amazon.com/AmazonElastiCache/latest/UserGuide/SelectEngine.html)
// in the AWS Documentation center for supported versions
func (r *Cluster) EngineVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["engineVersion"])
}

// Specifies the weekly time range for when maintenance
// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
func (r *Cluster) MaintenanceWindow() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["maintenanceWindow"])
}

// The compute and memory capacity of the nodes. See
// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
// supported node types
func (r *Cluster) NodeType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["nodeType"])
}

// An Amazon Resource Name (ARN) of an
// SNS topic to send ElastiCache notifications to. Example:
// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
func (r *Cluster) NotificationTopicArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["notificationTopicArn"])
}

// The initial number of cache nodes that the
// cache cluster will have. For Redis, this value must be 1. For Memcache, this
// value must be between 1 and 20. If this number is reduced on subsequent runs,
// the highest numbered nodes will be removed.
func (r *Cluster) NumCacheNodes() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["numCacheNodes"])
}

// Name of the parameter group to associate
// with this cache cluster
func (r *Cluster) ParameterGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["parameterGroupName"])
}

// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replication_group_id`.
func (r *Cluster) Port() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["port"])
}

// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
func (r *Cluster) ReplicationGroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["replicationGroupId"])
}

// One or more VPC security groups associated
// with the cache cluster
func (r *Cluster) SecurityGroupIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["securityGroupIds"])
}

// List of security group
// names to associate with this cache cluster
func (r *Cluster) SecurityGroupNames() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["securityGroupNames"])
}

// A single-element string list containing an
// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
func (r *Cluster) SnapshotArns() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["snapshotArns"])
}

// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshot_name` forces a new resource.
func (r *Cluster) SnapshotName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["snapshotName"])
}

// The number of days for which ElastiCache will
// retain automatic cache cluster snapshots before deleting them. For example, if you set
// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
// Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro or cache.t2.* cache nodes
func (r *Cluster) SnapshotRetentionLimit() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["snapshotRetentionLimit"])
}

// The daily time range (in UTC) during which ElastiCache will
// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
func (r *Cluster) SnapshotWindow() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["snapshotWindow"])
}

// Name of the subnet group to be used
// for the cache cluster.
func (r *Cluster) SubnetGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetGroupName"])
}

// A mapping of tags to assign to the resource
func (r *Cluster) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Cluster resources.
type ClusterState struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon ElastiCache Documentation for more information.][1]
	// (Available since v0.6.0)
	ApplyImmediately interface{}
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `availability_zones`
	AvailabilityZone interface{}
	// List of Availability Zones in which the cache nodes will be created. If you want to create cache nodes in single-az, use `availability_zone`
	AvailabilityZones interface{}
	// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `num_cache_nodes` must be greater than `1`
	AzMode interface{}
	// List of node objects including `id`, `address`, `port` and `availability_zone`.
	// Referenceable e.g. as `${aws_elasticache_cluster.bar.cache_nodes.0.address}`
	CacheNodes interface{}
	// (Memcached only) The DNS name of the cache cluster without the port appended.
	ClusterAddress interface{}
	// Group identifier. ElastiCache converts
	// this name to lowercase
	ClusterId interface{}
	// (Memcached only) The configuration endpoint to allow host discovery.
	ConfigurationEndpoint interface{}
	// Name of the cache engine to be used for this cache cluster.
	// Valid values for this parameter are `memcached` or `redis`
	Engine interface{}
	// Version number of the cache engine to be used.
	// See [Selecting a Cache Engine and Version](https://docs.aws.amazon.com/AmazonElastiCache/latest/UserGuide/SelectEngine.html)
	// in the AWS Documentation center for supported versions
	EngineVersion interface{}
	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow interface{}
	// The compute and memory capacity of the nodes. See
	// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
	// supported node types
	NodeType interface{}
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send ElastiCache notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn interface{}
	// The initial number of cache nodes that the
	// cache cluster will have. For Redis, this value must be 1. For Memcache, this
	// value must be between 1 and 20. If this number is reduced on subsequent runs,
	// the highest numbered nodes will be removed.
	NumCacheNodes interface{}
	// Name of the parameter group to associate
	// with this cache cluster
	ParameterGroupName interface{}
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replication_group_id`.
	Port interface{}
	// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupId interface{}
	// One or more VPC security groups associated
	// with the cache cluster
	SecurityGroupIds interface{}
	// List of security group
	// names to associate with this cache cluster
	SecurityGroupNames interface{}
	// A single-element string list containing an
	// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
	// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns interface{}
	// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshot_name` forces a new resource.
	SnapshotName interface{}
	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them. For example, if you set
	// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
	// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
	// Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro or cache.t2.* cache nodes
	SnapshotRetentionLimit interface{}
	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow interface{}
	// Name of the subnet group to be used
	// for the cache cluster.
	SubnetGroupName interface{}
	// A mapping of tags to assign to the resource
	Tags interface{}
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is
	// `false`. See [Amazon ElastiCache Documentation for more information.][1]
	// (Available since v0.6.0)
	ApplyImmediately interface{}
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `availability_zones`
	AvailabilityZone interface{}
	// List of Availability Zones in which the cache nodes will be created. If you want to create cache nodes in single-az, use `availability_zone`
	AvailabilityZones interface{}
	// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `num_cache_nodes` must be greater than `1`
	AzMode interface{}
	// Group identifier. ElastiCache converts
	// this name to lowercase
	ClusterId interface{}
	// Name of the cache engine to be used for this cache cluster.
	// Valid values for this parameter are `memcached` or `redis`
	Engine interface{}
	// Version number of the cache engine to be used.
	// See [Selecting a Cache Engine and Version](https://docs.aws.amazon.com/AmazonElastiCache/latest/UserGuide/SelectEngine.html)
	// in the AWS Documentation center for supported versions
	EngineVersion interface{}
	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
	// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
	MaintenanceWindow interface{}
	// The compute and memory capacity of the nodes. See
	// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
	// supported node types
	NodeType interface{}
	// An Amazon Resource Name (ARN) of an
	// SNS topic to send ElastiCache notifications to. Example:
	// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
	NotificationTopicArn interface{}
	// The initial number of cache nodes that the
	// cache cluster will have. For Redis, this value must be 1. For Memcache, this
	// value must be between 1 and 20. If this number is reduced on subsequent runs,
	// the highest numbered nodes will be removed.
	NumCacheNodes interface{}
	// Name of the parameter group to associate
	// with this cache cluster
	ParameterGroupName interface{}
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replication_group_id`.
	Port interface{}
	// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
	ReplicationGroupId interface{}
	// One or more VPC security groups associated
	// with the cache cluster
	SecurityGroupIds interface{}
	// List of security group
	// names to associate with this cache cluster
	SecurityGroupNames interface{}
	// A single-element string list containing an
	// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
	// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns interface{}
	// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshot_name` forces a new resource.
	SnapshotName interface{}
	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them. For example, if you set
	// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
	// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
	// Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro or cache.t2.* cache nodes
	SnapshotRetentionLimit interface{}
	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
	SnapshotWindow interface{}
	// Name of the subnet group to be used
	// for the cache cluster.
	SubnetGroupName interface{}
	// A mapping of tags to assign to the resource
	Tags interface{}
}
