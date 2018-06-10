// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an SSM Patch Baseline resource
// 
// ~> **NOTE on Patch Baselines:** The `approved_patches` and `approval_rule` are 
// both marked as optional fields, but the Patch Baseline requires that at least one
// of them is specified.
type PatchBaseline struct {
	s *pulumi.ResourceState
}

// NewPatchBaseline registers a new resource with the given unique name, arguments, and options.
func NewPatchBaseline(ctx *pulumi.Context,
	name string, args *PatchBaselineArgs, opts ...pulumi.ResourceOpt) (*PatchBaseline, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["approvalRules"] = nil
		inputs["approvedPatches"] = nil
		inputs["approvedPatchesComplianceLevel"] = nil
		inputs["description"] = nil
		inputs["globalFilters"] = nil
		inputs["name"] = nil
		inputs["operatingSystem"] = nil
		inputs["rejectedPatches"] = nil
	} else {
		inputs["approvalRules"] = args.ApprovalRules
		inputs["approvedPatches"] = args.ApprovedPatches
		inputs["approvedPatchesComplianceLevel"] = args.ApprovedPatchesComplianceLevel
		inputs["description"] = args.Description
		inputs["globalFilters"] = args.GlobalFilters
		inputs["name"] = args.Name
		inputs["operatingSystem"] = args.OperatingSystem
		inputs["rejectedPatches"] = args.RejectedPatches
	}
	s, err := ctx.RegisterResource("aws:ssm/patchBaseline:PatchBaseline", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PatchBaseline{s: s}, nil
}

// GetPatchBaseline gets an existing PatchBaseline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPatchBaseline(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PatchBaselineState, opts ...pulumi.ResourceOpt) (*PatchBaseline, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["approvalRules"] = state.ApprovalRules
		inputs["approvedPatches"] = state.ApprovedPatches
		inputs["approvedPatchesComplianceLevel"] = state.ApprovedPatchesComplianceLevel
		inputs["description"] = state.Description
		inputs["globalFilters"] = state.GlobalFilters
		inputs["name"] = state.Name
		inputs["operatingSystem"] = state.OperatingSystem
		inputs["rejectedPatches"] = state.RejectedPatches
	}
	s, err := ctx.ReadResource("aws:ssm/patchBaseline:PatchBaseline", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PatchBaseline{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PatchBaseline) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PatchBaseline) ID() *pulumi.IDOutput {
	return r.s.ID
}

// A set of rules used to include patches in the baseline. up to 10 approval rules can be specified. Each approval_rule block requires the fields documented below.
func (r *PatchBaseline) ApprovalRules() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["approvalRules"])
}

// A list of explicitly approved patches for the baseline.
func (r *PatchBaseline) ApprovedPatches() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["approvedPatches"])
}

// Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid compliance levels include the following: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`. The default value is `UNSPECIFIED`.
func (r *PatchBaseline) ApprovedPatchesComplianceLevel() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["approvedPatchesComplianceLevel"])
}

// The description of the patch baseline.
func (r *PatchBaseline) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// A set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are `PRODUCT | CLASSIFICATION | MSRC_SEVERITY | PATCH_ID`.
func (r *PatchBaseline) GlobalFilters() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["globalFilters"])
}

// The name of the patch baseline.
func (r *PatchBaseline) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Defines the operating system the patch baseline applies to. Supported operating systems include `WINDOWS`, `AMAZON_LINUX`, `UBUNTU`, `CENTOS`, and `REDHAT_ENTERPRISE_LINUX`. The Default value is `WINDOWS`.
func (r *PatchBaseline) OperatingSystem() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["operatingSystem"])
}

// A list of rejected patches.
func (r *PatchBaseline) RejectedPatches() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["rejectedPatches"])
}

// Input properties used for looking up and filtering PatchBaseline resources.
type PatchBaselineState struct {
	// A set of rules used to include patches in the baseline. up to 10 approval rules can be specified. Each approval_rule block requires the fields documented below.
	ApprovalRules interface{}
	// A list of explicitly approved patches for the baseline.
	ApprovedPatches interface{}
	// Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid compliance levels include the following: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`. The default value is `UNSPECIFIED`.
	ApprovedPatchesComplianceLevel interface{}
	// The description of the patch baseline.
	Description interface{}
	// A set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are `PRODUCT | CLASSIFICATION | MSRC_SEVERITY | PATCH_ID`.
	GlobalFilters interface{}
	// The name of the patch baseline.
	Name interface{}
	// Defines the operating system the patch baseline applies to. Supported operating systems include `WINDOWS`, `AMAZON_LINUX`, `UBUNTU`, `CENTOS`, and `REDHAT_ENTERPRISE_LINUX`. The Default value is `WINDOWS`.
	OperatingSystem interface{}
	// A list of rejected patches.
	RejectedPatches interface{}
}

// The set of arguments for constructing a PatchBaseline resource.
type PatchBaselineArgs struct {
	// A set of rules used to include patches in the baseline. up to 10 approval rules can be specified. Each approval_rule block requires the fields documented below.
	ApprovalRules interface{}
	// A list of explicitly approved patches for the baseline.
	ApprovedPatches interface{}
	// Defines the compliance level for approved patches. This means that if an approved patch is reported as missing, this is the severity of the compliance violation. Valid compliance levels include the following: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, `INFORMATIONAL`, `UNSPECIFIED`. The default value is `UNSPECIFIED`.
	ApprovedPatchesComplianceLevel interface{}
	// The description of the patch baseline.
	Description interface{}
	// A set of global filters used to exclude patches from the baseline. Up to 4 global filters can be specified using Key/Value pairs. Valid Keys are `PRODUCT | CLASSIFICATION | MSRC_SEVERITY | PATCH_ID`.
	GlobalFilters interface{}
	// The name of the patch baseline.
	Name interface{}
	// Defines the operating system the patch baseline applies to. Supported operating systems include `WINDOWS`, `AMAZON_LINUX`, `UBUNTU`, `CENTOS`, and `REDHAT_ENTERPRISE_LINUX`. The Default value is `WINDOWS`.
	OperatingSystem interface{}
	// A list of rejected patches.
	RejectedPatches interface{}
}
