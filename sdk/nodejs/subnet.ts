// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Subnet extends pulumi.CustomResource {
    /**
     * Get an existing Subnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetState, opts?: pulumi.CustomResourceOptions): Subnet {
        return new Subnet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'outscale:index/subnet:Subnet';

    /**
     * Returns true if the given object is an instance of Subnet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subnet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subnet.__pulumiType;
    }

    public /*out*/ readonly availableIpsCount!: pulumi.Output<number>;
    public readonly ipRange!: pulumi.Output<string>;
    public readonly mapPublicIpOnLaunch!: pulumi.Output<boolean | undefined>;
    public readonly netId!: pulumi.Output<string>;
    public /*out*/ readonly requestId!: pulumi.Output<string>;
    public /*out*/ readonly state!: pulumi.Output<string>;
    public /*out*/ readonly subnetId!: pulumi.Output<string>;
    public readonly subregionName!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.SubnetTag[] | undefined>;

    /**
     * Create a Subnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetArgs | SubnetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubnetState | undefined;
            resourceInputs["availableIpsCount"] = state ? state.availableIpsCount : undefined;
            resourceInputs["ipRange"] = state ? state.ipRange : undefined;
            resourceInputs["mapPublicIpOnLaunch"] = state ? state.mapPublicIpOnLaunch : undefined;
            resourceInputs["netId"] = state ? state.netId : undefined;
            resourceInputs["requestId"] = state ? state.requestId : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["subregionName"] = state ? state.subregionName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SubnetArgs | undefined;
            if ((!args || args.ipRange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipRange'");
            }
            if ((!args || args.netId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'netId'");
            }
            resourceInputs["ipRange"] = args ? args.ipRange : undefined;
            resourceInputs["mapPublicIpOnLaunch"] = args ? args.mapPublicIpOnLaunch : undefined;
            resourceInputs["netId"] = args ? args.netId : undefined;
            resourceInputs["subregionName"] = args ? args.subregionName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["availableIpsCount"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["subnetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Subnet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subnet resources.
 */
export interface SubnetState {
    availableIpsCount?: pulumi.Input<number>;
    ipRange?: pulumi.Input<string>;
    mapPublicIpOnLaunch?: pulumi.Input<boolean>;
    netId?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    state?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    subregionName?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.SubnetTag>[]>;
}

/**
 * The set of arguments for constructing a Subnet resource.
 */
export interface SubnetArgs {
    ipRange: pulumi.Input<string>;
    mapPublicIpOnLaunch?: pulumi.Input<boolean>;
    netId: pulumi.Input<string>;
    subregionName?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.SubnetTag>[]>;
}
