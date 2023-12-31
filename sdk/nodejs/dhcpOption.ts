// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class DhcpOption extends pulumi.CustomResource {
    /**
     * Get an existing DhcpOption resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DhcpOptionState, opts?: pulumi.CustomResourceOptions): DhcpOption {
        return new DhcpOption(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'outscale:index/dhcpOption:DhcpOption';

    /**
     * Returns true if the given object is an instance of DhcpOption.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DhcpOption {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DhcpOption.__pulumiType;
    }

    public /*out*/ readonly default!: pulumi.Output<boolean>;
    public /*out*/ readonly dhcpOptionsSetId!: pulumi.Output<string>;
    public readonly domainName!: pulumi.Output<string>;
    public readonly domainNameServers!: pulumi.Output<string[]>;
    public readonly logServers!: pulumi.Output<string[]>;
    public readonly ntpServers!: pulumi.Output<string[]>;
    public /*out*/ readonly requestId!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<outputs.DhcpOptionTag[] | undefined>;

    /**
     * Create a DhcpOption resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DhcpOptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DhcpOptionArgs | DhcpOptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DhcpOptionState | undefined;
            resourceInputs["default"] = state ? state.default : undefined;
            resourceInputs["dhcpOptionsSetId"] = state ? state.dhcpOptionsSetId : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["domainNameServers"] = state ? state.domainNameServers : undefined;
            resourceInputs["logServers"] = state ? state.logServers : undefined;
            resourceInputs["ntpServers"] = state ? state.ntpServers : undefined;
            resourceInputs["requestId"] = state ? state.requestId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DhcpOptionArgs | undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["domainNameServers"] = args ? args.domainNameServers : undefined;
            resourceInputs["logServers"] = args ? args.logServers : undefined;
            resourceInputs["ntpServers"] = args ? args.ntpServers : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["default"] = undefined /*out*/;
            resourceInputs["dhcpOptionsSetId"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DhcpOption.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DhcpOption resources.
 */
export interface DhcpOptionState {
    default?: pulumi.Input<boolean>;
    dhcpOptionsSetId?: pulumi.Input<string>;
    domainName?: pulumi.Input<string>;
    domainNameServers?: pulumi.Input<pulumi.Input<string>[]>;
    logServers?: pulumi.Input<pulumi.Input<string>[]>;
    ntpServers?: pulumi.Input<pulumi.Input<string>[]>;
    requestId?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.DhcpOptionTag>[]>;
}

/**
 * The set of arguments for constructing a DhcpOption resource.
 */
export interface DhcpOptionArgs {
    domainName?: pulumi.Input<string>;
    domainNameServers?: pulumi.Input<pulumi.Input<string>[]>;
    logServers?: pulumi.Input<pulumi.Input<string>[]>;
    ntpServers?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.DhcpOptionTag>[]>;
}
