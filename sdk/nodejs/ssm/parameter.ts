// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {Tags} from "../index";

/**
 * Provides an SSM Parameter resource.
 */
export class Parameter extends pulumi.CustomResource {
    /**
     * Get an existing Parameter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ParameterState): Parameter {
        return new Parameter(name, <any>state, { id });
    }

    /**
     * A regular expression used to validate the parameter value.
     */
    public readonly allowedPattern: pulumi.Output<string | undefined>;
    /**
     * The ARN of the parameter.
     */
    public readonly arn: pulumi.Output<string>;
    /**
     * The description of the parameter.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * The KMS key id or arn for encrypting a SecureString.
     */
    public readonly keyId: pulumi.Output<string>;
    /**
     * The name of the parameter.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by terraform to avoid overwrite of existing resource and will default to `true` otherwise (terraform lifecycle rules should then be used to manage the update behavior).
     */
    public readonly overwrite: pulumi.Output<boolean | undefined>;
    /**
     * A mapping of tags to assign to the object.
     */
    public readonly tags: pulumi.Output<Tags | undefined>;
    /**
     * The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
     */
    public readonly type: pulumi.Output<string>;
    /**
     * The value of the parameter.
     */
    public readonly value: pulumi.Output<string>;

    /**
     * Create a Parameter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ParameterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ParameterArgs | ParameterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ParameterState = argsOrState as ParameterState | undefined;
            inputs["allowedPattern"] = state ? state.allowedPattern : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["keyId"] = state ? state.keyId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["overwrite"] = state ? state.overwrite : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ParameterArgs | undefined;
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            if (!args || args.value === undefined) {
                throw new Error("Missing required property 'value'");
            }
            inputs["allowedPattern"] = args ? args.allowedPattern : undefined;
            inputs["arn"] = args ? args.arn : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["keyId"] = args ? args.keyId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["overwrite"] = args ? args.overwrite : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        super("aws:ssm/parameter:Parameter", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Parameter resources.
 */
export interface ParameterState {
    /**
     * A regular expression used to validate the parameter value.
     */
    readonly allowedPattern?: pulumi.Input<string>;
    /**
     * The ARN of the parameter.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The description of the parameter.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The KMS key id or arn for encrypting a SecureString.
     */
    readonly keyId?: pulumi.Input<string>;
    /**
     * The name of the parameter.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by terraform to avoid overwrite of existing resource and will default to `true` otherwise (terraform lifecycle rules should then be used to manage the update behavior).
     */
    readonly overwrite?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<Tags>;
    /**
     * The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The value of the parameter.
     */
    readonly value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Parameter resource.
 */
export interface ParameterArgs {
    /**
     * A regular expression used to validate the parameter value.
     */
    readonly allowedPattern?: pulumi.Input<string>;
    /**
     * The ARN of the parameter.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The description of the parameter.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The KMS key id or arn for encrypting a SecureString.
     */
    readonly keyId?: pulumi.Input<string>;
    /**
     * The name of the parameter.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by terraform to avoid overwrite of existing resource and will default to `true` otherwise (terraform lifecycle rules should then be used to manage the update behavior).
     */
    readonly overwrite?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<Tags>;
    /**
     * The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
     */
    readonly type: pulumi.Input<string>;
    /**
     * The value of the parameter.
     */
    readonly value: pulumi.Input<string>;
}
