// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupState, opts?: pulumi.CustomResourceOptions): SecurityGroup {
        return new SecurityGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'outscale:index/securityGroup:SecurityGroup';

    /**
     * Returns true if the given object is an instance of SecurityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroup.__pulumiType;
    }

    public /*out*/ readonly accountId!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly inboundRules!: pulumi.Output<outputs.SecurityGroupInboundRule[]>;
    public readonly netId!: pulumi.Output<string>;
    public /*out*/ readonly outboundRules!: pulumi.Output<outputs.SecurityGroupOutboundRule[]>;
    public readonly removeDefaultOutboundRule!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly requestId!: pulumi.Output<string>;
    public /*out*/ readonly securityGroupId!: pulumi.Output<string>;
    public readonly securityGroupName!: pulumi.Output<string>;
    public readonly tag!: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly tags!: pulumi.Output<outputs.SecurityGroupTag[] | undefined>;

    /**
     * Create a SecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecurityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupArgs | SecurityGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["inboundRules"] = state ? state.inboundRules : undefined;
            resourceInputs["netId"] = state ? state.netId : undefined;
            resourceInputs["outboundRules"] = state ? state.outboundRules : undefined;
            resourceInputs["removeDefaultOutboundRule"] = state ? state.removeDefaultOutboundRule : undefined;
            resourceInputs["requestId"] = state ? state.requestId : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["securityGroupName"] = state ? state.securityGroupName : undefined;
            resourceInputs["tag"] = state ? state.tag : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SecurityGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["netId"] = args ? args.netId : undefined;
            resourceInputs["removeDefaultOutboundRule"] = args ? args.removeDefaultOutboundRule : undefined;
            resourceInputs["securityGroupName"] = args ? args.securityGroupName : undefined;
            resourceInputs["tag"] = args ? args.tag : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["inboundRules"] = undefined /*out*/;
            resourceInputs["outboundRules"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["securityGroupId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroup resources.
 */
export interface SecurityGroupState {
    accountId?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    inboundRules?: pulumi.Input<pulumi.Input<inputs.SecurityGroupInboundRule>[]>;
    netId?: pulumi.Input<string>;
    outboundRules?: pulumi.Input<pulumi.Input<inputs.SecurityGroupOutboundRule>[]>;
    removeDefaultOutboundRule?: pulumi.Input<boolean>;
    requestId?: pulumi.Input<string>;
    securityGroupId?: pulumi.Input<string>;
    securityGroupName?: pulumi.Input<string>;
    tag?: pulumi.Input<{[key: string]: any}>;
    tags?: pulumi.Input<pulumi.Input<inputs.SecurityGroupTag>[]>;
}

/**
 * The set of arguments for constructing a SecurityGroup resource.
 */
export interface SecurityGroupArgs {
    description?: pulumi.Input<string>;
    netId?: pulumi.Input<string>;
    removeDefaultOutboundRule?: pulumi.Input<boolean>;
    securityGroupName?: pulumi.Input<string>;
    tag?: pulumi.Input<{[key: string]: any}>;
    tags?: pulumi.Input<pulumi.Input<inputs.SecurityGroupTag>[]>;
}
