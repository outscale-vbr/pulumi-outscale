// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class NetPeering extends pulumi.CustomResource {
    /**
     * Get an existing NetPeering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetPeeringState, opts?: pulumi.CustomResourceOptions): NetPeering {
        return new NetPeering(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'outscale:index/netPeering:NetPeering';

    /**
     * Returns true if the given object is an instance of NetPeering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetPeering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetPeering.__pulumiType;
    }

    public /*out*/ readonly accepterNet!: pulumi.Output<outputs.NetPeeringAccepterNet>;
    public readonly accepterNetId!: pulumi.Output<string>;
    public /*out*/ readonly netPeeringId!: pulumi.Output<string>;
    public /*out*/ readonly requestId!: pulumi.Output<string>;
    public /*out*/ readonly sourceNet!: pulumi.Output<outputs.NetPeeringSourceNet>;
    public readonly sourceNetAccountId!: pulumi.Output<string>;
    public readonly sourceNetId!: pulumi.Output<string>;
    public /*out*/ readonly state!: pulumi.Output<outputs.NetPeeringState>;
    public readonly tags!: pulumi.Output<outputs.NetPeeringTag[] | undefined>;

    /**
     * Create a NetPeering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetPeeringArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetPeeringArgs | NetPeeringState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetPeeringState | undefined;
            resourceInputs["accepterNet"] = state ? state.accepterNet : undefined;
            resourceInputs["accepterNetId"] = state ? state.accepterNetId : undefined;
            resourceInputs["netPeeringId"] = state ? state.netPeeringId : undefined;
            resourceInputs["requestId"] = state ? state.requestId : undefined;
            resourceInputs["sourceNet"] = state ? state.sourceNet : undefined;
            resourceInputs["sourceNetAccountId"] = state ? state.sourceNetAccountId : undefined;
            resourceInputs["sourceNetId"] = state ? state.sourceNetId : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as NetPeeringArgs | undefined;
            if ((!args || args.accepterNetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accepterNetId'");
            }
            if ((!args || args.sourceNetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceNetId'");
            }
            resourceInputs["accepterNetId"] = args ? args.accepterNetId : undefined;
            resourceInputs["sourceNetAccountId"] = args ? args.sourceNetAccountId : undefined;
            resourceInputs["sourceNetId"] = args ? args.sourceNetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["accepterNet"] = undefined /*out*/;
            resourceInputs["netPeeringId"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["sourceNet"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetPeering.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetPeering resources.
 */
export interface NetPeeringState {
    accepterNet?: pulumi.Input<inputs.NetPeeringAccepterNet>;
    accepterNetId?: pulumi.Input<string>;
    netPeeringId?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    sourceNet?: pulumi.Input<inputs.NetPeeringSourceNet>;
    sourceNetAccountId?: pulumi.Input<string>;
    sourceNetId?: pulumi.Input<string>;
    state?: pulumi.Input<inputs.NetPeeringState>;
    tags?: pulumi.Input<pulumi.Input<inputs.NetPeeringTag>[]>;
}

/**
 * The set of arguments for constructing a NetPeering resource.
 */
export interface NetPeeringArgs {
    accepterNetId: pulumi.Input<string>;
    sourceNetAccountId?: pulumi.Input<string>;
    sourceNetId: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.NetPeeringTag>[]>;
}
