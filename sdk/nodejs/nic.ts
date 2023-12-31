// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Nic extends pulumi.CustomResource {
    /**
     * Get an existing Nic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NicState, opts?: pulumi.CustomResourceOptions): Nic {
        return new Nic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'outscale:index/nic:Nic';

    /**
     * Returns true if the given object is an instance of Nic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Nic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Nic.__pulumiType;
    }

    public /*out*/ readonly accountId!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string>;
    public /*out*/ readonly isSourceDestChecked!: pulumi.Output<boolean>;
    public /*out*/ readonly linkNic!: pulumi.Output<outputs.NicLinkNic>;
    public /*out*/ readonly linkPublicIp!: pulumi.Output<outputs.NicLinkPublicIp>;
    public /*out*/ readonly macAddress!: pulumi.Output<string>;
    public /*out*/ readonly netId!: pulumi.Output<string>;
    public /*out*/ readonly nicId!: pulumi.Output<string>;
    public /*out*/ readonly privateDnsName!: pulumi.Output<string>;
    public readonly privateIp!: pulumi.Output<string>;
    public readonly privateIps!: pulumi.Output<outputs.NicPrivateIp[]>;
    public /*out*/ readonly requestId!: pulumi.Output<string>;
    public /*out*/ readonly requesterManaged!: pulumi.Output<boolean>;
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly securityGroups!: pulumi.Output<outputs.NicSecurityGroup[]>;
    public /*out*/ readonly state!: pulumi.Output<string>;
    public readonly subnetId!: pulumi.Output<string>;
    public /*out*/ readonly subregionName!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.NicTag[] | undefined>;

    /**
     * Create a Nic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NicArgs | NicState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NicState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isSourceDestChecked"] = state ? state.isSourceDestChecked : undefined;
            resourceInputs["linkNic"] = state ? state.linkNic : undefined;
            resourceInputs["linkPublicIp"] = state ? state.linkPublicIp : undefined;
            resourceInputs["macAddress"] = state ? state.macAddress : undefined;
            resourceInputs["netId"] = state ? state.netId : undefined;
            resourceInputs["nicId"] = state ? state.nicId : undefined;
            resourceInputs["privateDnsName"] = state ? state.privateDnsName : undefined;
            resourceInputs["privateIp"] = state ? state.privateIp : undefined;
            resourceInputs["privateIps"] = state ? state.privateIps : undefined;
            resourceInputs["requestId"] = state ? state.requestId : undefined;
            resourceInputs["requesterManaged"] = state ? state.requesterManaged : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["subregionName"] = state ? state.subregionName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as NicArgs | undefined;
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["privateIp"] = args ? args.privateIp : undefined;
            resourceInputs["privateIps"] = args ? args.privateIps : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["isSourceDestChecked"] = undefined /*out*/;
            resourceInputs["linkNic"] = undefined /*out*/;
            resourceInputs["linkPublicIp"] = undefined /*out*/;
            resourceInputs["macAddress"] = undefined /*out*/;
            resourceInputs["netId"] = undefined /*out*/;
            resourceInputs["nicId"] = undefined /*out*/;
            resourceInputs["privateDnsName"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["requesterManaged"] = undefined /*out*/;
            resourceInputs["securityGroups"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["subregionName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Nic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Nic resources.
 */
export interface NicState {
    accountId?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    isSourceDestChecked?: pulumi.Input<boolean>;
    linkNic?: pulumi.Input<inputs.NicLinkNic>;
    linkPublicIp?: pulumi.Input<inputs.NicLinkPublicIp>;
    macAddress?: pulumi.Input<string>;
    netId?: pulumi.Input<string>;
    nicId?: pulumi.Input<string>;
    privateDnsName?: pulumi.Input<string>;
    privateIp?: pulumi.Input<string>;
    privateIps?: pulumi.Input<pulumi.Input<inputs.NicPrivateIp>[]>;
    requestId?: pulumi.Input<string>;
    requesterManaged?: pulumi.Input<boolean>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    securityGroups?: pulumi.Input<pulumi.Input<inputs.NicSecurityGroup>[]>;
    state?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    subregionName?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.NicTag>[]>;
}

/**
 * The set of arguments for constructing a Nic resource.
 */
export interface NicArgs {
    description?: pulumi.Input<string>;
    privateIp?: pulumi.Input<string>;
    privateIps?: pulumi.Input<pulumi.Input<inputs.NicPrivateIp>[]>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    subnetId: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.NicTag>[]>;
}
