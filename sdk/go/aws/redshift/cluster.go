// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Redshift Cluster Resource.
// 
// ~> **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](/docs/state/sensitive-data.html).
type Cluster struct {
	s *pulumi.ResourceState
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	if args == nil || args.ClusterIdentifier == nil {
		return nil, errors.New("missing required argument 'ClusterIdentifier'")
	}
	if args == nil || args.NodeType == nil {
		return nil, errors.New("missing required argument 'NodeType'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allowVersionUpgrade"] = nil
		inputs["automatedSnapshotRetentionPeriod"] = nil
		inputs["availabilityZone"] = nil
		inputs["bucketName"] = nil
		inputs["clusterIdentifier"] = nil
		inputs["clusterParameterGroupName"] = nil
		inputs["clusterPublicKey"] = nil
		inputs["clusterRevisionNumber"] = nil
		inputs["clusterSecurityGroups"] = nil
		inputs["clusterSubnetGroupName"] = nil
		inputs["clusterType"] = nil
		inputs["clusterVersion"] = nil
		inputs["databaseName"] = nil
		inputs["elasticIp"] = nil
		inputs["enableLogging"] = nil
		inputs["encrypted"] = nil
		inputs["endpoint"] = nil
		inputs["enhancedVpcRouting"] = nil
		inputs["finalSnapshotIdentifier"] = nil
		inputs["iamRoles"] = nil
		inputs["kmsKeyId"] = nil
		inputs["logging"] = nil
		inputs["masterPassword"] = nil
		inputs["masterUsername"] = nil
		inputs["nodeType"] = nil
		inputs["numberOfNodes"] = nil
		inputs["ownerAccount"] = nil
		inputs["port"] = nil
		inputs["preferredMaintenanceWindow"] = nil
		inputs["publiclyAccessible"] = nil
		inputs["s3KeyPrefix"] = nil
		inputs["skipFinalSnapshot"] = nil
		inputs["snapshotClusterIdentifier"] = nil
		inputs["snapshotCopy"] = nil
		inputs["snapshotIdentifier"] = nil
		inputs["tags"] = nil
		inputs["vpcSecurityGroupIds"] = nil
	} else {
		inputs["allowVersionUpgrade"] = args.AllowVersionUpgrade
		inputs["automatedSnapshotRetentionPeriod"] = args.AutomatedSnapshotRetentionPeriod
		inputs["availabilityZone"] = args.AvailabilityZone
		inputs["bucketName"] = args.BucketName
		inputs["clusterIdentifier"] = args.ClusterIdentifier
		inputs["clusterParameterGroupName"] = args.ClusterParameterGroupName
		inputs["clusterPublicKey"] = args.ClusterPublicKey
		inputs["clusterRevisionNumber"] = args.ClusterRevisionNumber
		inputs["clusterSecurityGroups"] = args.ClusterSecurityGroups
		inputs["clusterSubnetGroupName"] = args.ClusterSubnetGroupName
		inputs["clusterType"] = args.ClusterType
		inputs["clusterVersion"] = args.ClusterVersion
		inputs["databaseName"] = args.DatabaseName
		inputs["elasticIp"] = args.ElasticIp
		inputs["enableLogging"] = args.EnableLogging
		inputs["encrypted"] = args.Encrypted
		inputs["endpoint"] = args.Endpoint
		inputs["enhancedVpcRouting"] = args.EnhancedVpcRouting
		inputs["finalSnapshotIdentifier"] = args.FinalSnapshotIdentifier
		inputs["iamRoles"] = args.IamRoles
		inputs["kmsKeyId"] = args.KmsKeyId
		inputs["logging"] = args.Logging
		inputs["masterPassword"] = args.MasterPassword
		inputs["masterUsername"] = args.MasterUsername
		inputs["nodeType"] = args.NodeType
		inputs["numberOfNodes"] = args.NumberOfNodes
		inputs["ownerAccount"] = args.OwnerAccount
		inputs["port"] = args.Port
		inputs["preferredMaintenanceWindow"] = args.PreferredMaintenanceWindow
		inputs["publiclyAccessible"] = args.PubliclyAccessible
		inputs["s3KeyPrefix"] = args.S3KeyPrefix
		inputs["skipFinalSnapshot"] = args.SkipFinalSnapshot
		inputs["snapshotClusterIdentifier"] = args.SnapshotClusterIdentifier
		inputs["snapshotCopy"] = args.SnapshotCopy
		inputs["snapshotIdentifier"] = args.SnapshotIdentifier
		inputs["tags"] = args.Tags
		inputs["vpcSecurityGroupIds"] = args.VpcSecurityGroupIds
	}
	inputs["dnsName"] = nil
	s, err := ctx.RegisterResource("aws:redshift/cluster:Cluster", name, true, inputs, opts...)
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
		inputs["allowVersionUpgrade"] = state.AllowVersionUpgrade
		inputs["automatedSnapshotRetentionPeriod"] = state.AutomatedSnapshotRetentionPeriod
		inputs["availabilityZone"] = state.AvailabilityZone
		inputs["bucketName"] = state.BucketName
		inputs["clusterIdentifier"] = state.ClusterIdentifier
		inputs["clusterParameterGroupName"] = state.ClusterParameterGroupName
		inputs["clusterPublicKey"] = state.ClusterPublicKey
		inputs["clusterRevisionNumber"] = state.ClusterRevisionNumber
		inputs["clusterSecurityGroups"] = state.ClusterSecurityGroups
		inputs["clusterSubnetGroupName"] = state.ClusterSubnetGroupName
		inputs["clusterType"] = state.ClusterType
		inputs["clusterVersion"] = state.ClusterVersion
		inputs["databaseName"] = state.DatabaseName
		inputs["dnsName"] = state.DnsName
		inputs["elasticIp"] = state.ElasticIp
		inputs["enableLogging"] = state.EnableLogging
		inputs["encrypted"] = state.Encrypted
		inputs["endpoint"] = state.Endpoint
		inputs["enhancedVpcRouting"] = state.EnhancedVpcRouting
		inputs["finalSnapshotIdentifier"] = state.FinalSnapshotIdentifier
		inputs["iamRoles"] = state.IamRoles
		inputs["kmsKeyId"] = state.KmsKeyId
		inputs["logging"] = state.Logging
		inputs["masterPassword"] = state.MasterPassword
		inputs["masterUsername"] = state.MasterUsername
		inputs["nodeType"] = state.NodeType
		inputs["numberOfNodes"] = state.NumberOfNodes
		inputs["ownerAccount"] = state.OwnerAccount
		inputs["port"] = state.Port
		inputs["preferredMaintenanceWindow"] = state.PreferredMaintenanceWindow
		inputs["publiclyAccessible"] = state.PubliclyAccessible
		inputs["s3KeyPrefix"] = state.S3KeyPrefix
		inputs["skipFinalSnapshot"] = state.SkipFinalSnapshot
		inputs["snapshotClusterIdentifier"] = state.SnapshotClusterIdentifier
		inputs["snapshotCopy"] = state.SnapshotCopy
		inputs["snapshotIdentifier"] = state.SnapshotIdentifier
		inputs["tags"] = state.Tags
		inputs["vpcSecurityGroupIds"] = state.VpcSecurityGroupIds
	}
	s, err := ctx.ReadResource("aws:redshift/cluster:Cluster", name, id, inputs, opts...)
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

// If true , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default is true
func (r *Cluster) AllowVersionUpgrade() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["allowVersionUpgrade"])
}

// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with create-cluster-snapshot. Default is 1.
func (r *Cluster) AutomatedSnapshotRetentionPeriod() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["automatedSnapshotRetentionPeriod"])
}

// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency.
func (r *Cluster) AvailabilityZone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["availabilityZone"])
}

// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
// For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
func (r *Cluster) BucketName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["bucketName"])
}

// The Cluster Identifier. Must be a lower case
// string.
func (r *Cluster) ClusterIdentifier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterIdentifier"])
}

// The name of the parameter group to be associated with this cluster.
func (r *Cluster) ClusterParameterGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterParameterGroupName"])
}

// The public key for the cluster
func (r *Cluster) ClusterPublicKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterPublicKey"])
}

// The specific revision number of the database in the cluster
func (r *Cluster) ClusterRevisionNumber() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterRevisionNumber"])
}

// A list of security groups to be associated with this cluster.
func (r *Cluster) ClusterSecurityGroups() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["clusterSecurityGroups"])
}

// The name of a cluster subnet group to be associated with this cluster. If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
func (r *Cluster) ClusterSubnetGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterSubnetGroupName"])
}

// The cluster type to use. Either `single-node` or `multi-node`.
func (r *Cluster) ClusterType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterType"])
}

// The version of the Amazon Redshift engine software that you want to deploy on the cluster.
// The version selected runs on all the nodes in the cluster.
func (r *Cluster) ClusterVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterVersion"])
}

// The name of the first database to be created when the cluster is created.
// If you do not provide a name, Amazon Redshift will create a default database called `dev`.
func (r *Cluster) DatabaseName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["databaseName"])
}

// The DNS name of the cluster
func (r *Cluster) DnsName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dnsName"])
}

// The Elastic IP (EIP) address for the cluster.
func (r *Cluster) ElasticIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["elasticIp"])
}

func (r *Cluster) EnableLogging() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableLogging"])
}

// If true , the data in the cluster is encrypted at rest.
func (r *Cluster) Encrypted() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["encrypted"])
}

// The connection endpoint
func (r *Cluster) Endpoint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endpoint"])
}

// If true , enhanced VPC routing is enabled.
func (r *Cluster) EnhancedVpcRouting() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enhancedVpcRouting"])
}

// The identifier of the final snapshot that is to be created immediately before deleting the cluster. If this parameter is provided, `skip_final_snapshot` must be false.
func (r *Cluster) FinalSnapshotIdentifier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["finalSnapshotIdentifier"])
}

// A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
func (r *Cluster) IamRoles() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["iamRoles"])
}

// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true.
func (r *Cluster) KmsKeyId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["kmsKeyId"])
}

// Logging, documented below.
func (r *Cluster) Logging() *pulumi.Output {
	return r.s.State["logging"]
}

// Password for the master DB user.
// Note that this may show up in logs, and it will be stored in the state file. Password must contain at least 8 chars and
// contain at least one uppercase letter, one lowercase letter, and one number.
func (r *Cluster) MasterPassword() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["masterPassword"])
}

// Username for the master DB user.
func (r *Cluster) MasterUsername() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["masterUsername"])
}

// The node type to be provisioned for the cluster.
func (r *Cluster) NodeType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["nodeType"])
}

// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node. Default is 1.
func (r *Cluster) NumberOfNodes() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["numberOfNodes"])
}

// The AWS customer account used to create or copy the snapshot. Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
func (r *Cluster) OwnerAccount() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ownerAccount"])
}

// The port number on which the cluster accepts incoming connections.
// The cluster is accessible only via the JDBC and ODBC connection strings. Part of the connection string requires the port on which the cluster will listen for incoming connections. Default port is 5439.
func (r *Cluster) Port() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["port"])
}

// The weekly time range (in UTC) during which automated cluster maintenance can occur.
// Format: ddd:hh24:mi-ddd:hh24:mi
func (r *Cluster) PreferredMaintenanceWindow() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["preferredMaintenanceWindow"])
}

// If true, the cluster can be accessed from a public network. Default is `true`.
func (r *Cluster) PubliclyAccessible() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["publiclyAccessible"])
}

// The prefix applied to the log file names.
func (r *Cluster) S3KeyPrefix() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["s3KeyPrefix"])
}

// Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster. If true , a final cluster snapshot is not created. If false , a final cluster snapshot is created before the cluster is deleted. Default is false.
func (r *Cluster) SkipFinalSnapshot() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["skipFinalSnapshot"])
}

// The name of the cluster the source snapshot was created from.
func (r *Cluster) SnapshotClusterIdentifier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["snapshotClusterIdentifier"])
}

// Configuration of automatic copy of snapshots from one region to another. Documented below.
func (r *Cluster) SnapshotCopy() *pulumi.Output {
	return r.s.State["snapshotCopy"]
}

// The name of the snapshot from which to create the new cluster.
func (r *Cluster) SnapshotIdentifier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["snapshotIdentifier"])
}

// A mapping of tags to assign to the resource.
func (r *Cluster) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
func (r *Cluster) VpcSecurityGroupIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["vpcSecurityGroupIds"])
}

// Input properties used for looking up and filtering Cluster resources.
type ClusterState struct {
	// If true , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default is true
	AllowVersionUpgrade interface{}
	// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with create-cluster-snapshot. Default is 1.
	AutomatedSnapshotRetentionPeriod interface{}
	// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency.
	AvailabilityZone interface{}
	// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
	// For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
	BucketName interface{}
	// The Cluster Identifier. Must be a lower case
	// string.
	ClusterIdentifier interface{}
	// The name of the parameter group to be associated with this cluster.
	ClusterParameterGroupName interface{}
	// The public key for the cluster
	ClusterPublicKey interface{}
	// The specific revision number of the database in the cluster
	ClusterRevisionNumber interface{}
	// A list of security groups to be associated with this cluster.
	ClusterSecurityGroups interface{}
	// The name of a cluster subnet group to be associated with this cluster. If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
	ClusterSubnetGroupName interface{}
	// The cluster type to use. Either `single-node` or `multi-node`.
	ClusterType interface{}
	// The version of the Amazon Redshift engine software that you want to deploy on the cluster.
	// The version selected runs on all the nodes in the cluster.
	ClusterVersion interface{}
	// The name of the first database to be created when the cluster is created.
	// If you do not provide a name, Amazon Redshift will create a default database called `dev`.
	DatabaseName interface{}
	// The DNS name of the cluster
	DnsName interface{}
	// The Elastic IP (EIP) address for the cluster.
	ElasticIp interface{}
	EnableLogging interface{}
	// If true , the data in the cluster is encrypted at rest.
	Encrypted interface{}
	// The connection endpoint
	Endpoint interface{}
	// If true , enhanced VPC routing is enabled.
	EnhancedVpcRouting interface{}
	// The identifier of the final snapshot that is to be created immediately before deleting the cluster. If this parameter is provided, `skip_final_snapshot` must be false.
	FinalSnapshotIdentifier interface{}
	// A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
	IamRoles interface{}
	// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true.
	KmsKeyId interface{}
	// Logging, documented below.
	Logging interface{}
	// Password for the master DB user.
	// Note that this may show up in logs, and it will be stored in the state file. Password must contain at least 8 chars and
	// contain at least one uppercase letter, one lowercase letter, and one number.
	MasterPassword interface{}
	// Username for the master DB user.
	MasterUsername interface{}
	// The node type to be provisioned for the cluster.
	NodeType interface{}
	// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node. Default is 1.
	NumberOfNodes interface{}
	// The AWS customer account used to create or copy the snapshot. Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
	OwnerAccount interface{}
	// The port number on which the cluster accepts incoming connections.
	// The cluster is accessible only via the JDBC and ODBC connection strings. Part of the connection string requires the port on which the cluster will listen for incoming connections. Default port is 5439.
	Port interface{}
	// The weekly time range (in UTC) during which automated cluster maintenance can occur.
	// Format: ddd:hh24:mi-ddd:hh24:mi
	PreferredMaintenanceWindow interface{}
	// If true, the cluster can be accessed from a public network. Default is `true`.
	PubliclyAccessible interface{}
	// The prefix applied to the log file names.
	S3KeyPrefix interface{}
	// Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster. If true , a final cluster snapshot is not created. If false , a final cluster snapshot is created before the cluster is deleted. Default is false.
	SkipFinalSnapshot interface{}
	// The name of the cluster the source snapshot was created from.
	SnapshotClusterIdentifier interface{}
	// Configuration of automatic copy of snapshots from one region to another. Documented below.
	SnapshotCopy interface{}
	// The name of the snapshot from which to create the new cluster.
	SnapshotIdentifier interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
	VpcSecurityGroupIds interface{}
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// If true , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default is true
	AllowVersionUpgrade interface{}
	// The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with create-cluster-snapshot. Default is 1.
	AutomatedSnapshotRetentionPeriod interface{}
	// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency.
	AvailabilityZone interface{}
	// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
	// For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
	BucketName interface{}
	// The Cluster Identifier. Must be a lower case
	// string.
	ClusterIdentifier interface{}
	// The name of the parameter group to be associated with this cluster.
	ClusterParameterGroupName interface{}
	// The public key for the cluster
	ClusterPublicKey interface{}
	// The specific revision number of the database in the cluster
	ClusterRevisionNumber interface{}
	// A list of security groups to be associated with this cluster.
	ClusterSecurityGroups interface{}
	// The name of a cluster subnet group to be associated with this cluster. If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
	ClusterSubnetGroupName interface{}
	// The cluster type to use. Either `single-node` or `multi-node`.
	ClusterType interface{}
	// The version of the Amazon Redshift engine software that you want to deploy on the cluster.
	// The version selected runs on all the nodes in the cluster.
	ClusterVersion interface{}
	// The name of the first database to be created when the cluster is created.
	// If you do not provide a name, Amazon Redshift will create a default database called `dev`.
	DatabaseName interface{}
	// The Elastic IP (EIP) address for the cluster.
	ElasticIp interface{}
	EnableLogging interface{}
	// If true , the data in the cluster is encrypted at rest.
	Encrypted interface{}
	// The connection endpoint
	Endpoint interface{}
	// If true , enhanced VPC routing is enabled.
	EnhancedVpcRouting interface{}
	// The identifier of the final snapshot that is to be created immediately before deleting the cluster. If this parameter is provided, `skip_final_snapshot` must be false.
	FinalSnapshotIdentifier interface{}
	// A list of IAM Role ARNs to associate with the cluster. A Maximum of 10 can be associated to the cluster at any time.
	IamRoles interface{}
	// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true.
	KmsKeyId interface{}
	// Logging, documented below.
	Logging interface{}
	// Password for the master DB user.
	// Note that this may show up in logs, and it will be stored in the state file. Password must contain at least 8 chars and
	// contain at least one uppercase letter, one lowercase letter, and one number.
	MasterPassword interface{}
	// Username for the master DB user.
	MasterUsername interface{}
	// The node type to be provisioned for the cluster.
	NodeType interface{}
	// The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node. Default is 1.
	NumberOfNodes interface{}
	// The AWS customer account used to create or copy the snapshot. Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
	OwnerAccount interface{}
	// The port number on which the cluster accepts incoming connections.
	// The cluster is accessible only via the JDBC and ODBC connection strings. Part of the connection string requires the port on which the cluster will listen for incoming connections. Default port is 5439.
	Port interface{}
	// The weekly time range (in UTC) during which automated cluster maintenance can occur.
	// Format: ddd:hh24:mi-ddd:hh24:mi
	PreferredMaintenanceWindow interface{}
	// If true, the cluster can be accessed from a public network. Default is `true`.
	PubliclyAccessible interface{}
	// The prefix applied to the log file names.
	S3KeyPrefix interface{}
	// Determines whether a final snapshot of the cluster is created before Amazon Redshift deletes the cluster. If true , a final cluster snapshot is not created. If false , a final cluster snapshot is created before the cluster is deleted. Default is false.
	SkipFinalSnapshot interface{}
	// The name of the cluster the source snapshot was created from.
	SnapshotClusterIdentifier interface{}
	// Configuration of automatic copy of snapshots from one region to another. Documented below.
	SnapshotCopy interface{}
	// The name of the snapshot from which to create the new cluster.
	SnapshotIdentifier interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
	VpcSecurityGroupIds interface{}
}
