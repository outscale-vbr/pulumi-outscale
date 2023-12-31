// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class NetAttributes extends pulumi.CustomResource {
    /**
     * Get an existing NetAttributes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetAttributesState, opts?: pulumi.CustomResourceOptions): NetAttributes {
        return new NetAttributes(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'outscale:index/netAttributes:NetAttributes';

    /**
     * Returns true if the given object is an instance of NetAttributes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetAttributes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetAttributes.__pulumiType;
    }

    public readonly dhcpOptionsSetId!: pulumi.Output<string>;
    public /*out*/ readonly ipRange!: pulumi.Output<string>;
    public readonly netId!: pulumi.Output<string>;
    public /*out*/ readonly requestId!: pulumi.Output<string>;
    public /*out*/ readonly state!: pulumi.Output<string>;
    public /*out*/ readonly tags!: pulumi.Output<outputs.NetAttributesTag[]>;
    public /*out*/ readonly tenancy!: pulumi.Output<string>;

    /**
     * Create a NetAttributes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetAttributesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetAttributesArgs | NetAttributesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetAttributesState | undefined;
            resourceInputs["dhcpOptionsSetId"] = state ? state.dhcpOptionsSetId : undefined;
            resourceInputs["ipRange"] = state ? state.ipRange : undefined;
            resourceInputs["netId"] = state ? state.netId : undefined;
            resourceInputs["requestId"] = state ? state.requestId : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenancy"] = state ? state.tenancy : undefined;
        } else {
            const args = argsOrState as NetAttributesArgs | undefined;
            if ((!args || args.netId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'netId'");
            }
            resourceInputs["dhcpOptionsSetId"] = args ? args.dhcpOptionsSetId : undefined;
            resourceInputs["netId"] = args ? args.netId : undefined;
            resourceInputs["ipRange"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["tenancy"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetAttributes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetAttributes resources.
 */
export interface NetAttributesState {
    dhcpOptionsSetId?: pulumi.Input<string>;
    ipRange?: pulumi.Input<string>;
    netId?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.NetAttributesTag>[]>;
    tenancy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetAttributes resource.
 */
export interface NetAttributesArgs {
    dhcpOptionsSetId?: pulumi.Input<string>;
    netId: pulumi.Input<string>;
}
