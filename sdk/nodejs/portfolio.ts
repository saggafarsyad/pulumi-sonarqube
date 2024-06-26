// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # sonarqube.Portfolio
 *
 * Provides a Sonarqube Portfolio resource. This can be used to create and manage Sonarqube Portfolio. Note that the SonarQube API for Portfolios is called ``views``
 *
 * ## Example: create a portfolio
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sonarqube from "@pulumi/sonarqube";
 *
 * const main = new sonarqube.Portfolio("main", {
 *     description: "portfolio-description",
 *     key: "portfolio-key",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Portfolios can be imported using their portfolio key
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class Portfolio extends pulumi.CustomResource {
    /**
     * Get an existing Portfolio resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PortfolioState, opts?: pulumi.CustomResourceOptions): Portfolio {
        return new Portfolio(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sonarqube:index/portfolio:Portfolio';

    /**
     * Returns true if the given object is an instance of Portfolio.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Portfolio {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Portfolio.__pulumiType;
    }

    /**
     * Which branch to analyze. If nothing, or '' is specified, the main branch is used.
     */
    public readonly branch!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string>;
    public readonly key!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly qualifier!: pulumi.Output<string>;
    public readonly regexp!: pulumi.Output<string | undefined>;
    /**
     * A set of projects to add to the portfolio.
     */
    public readonly selectedProjects!: pulumi.Output<outputs.PortfolioSelectedProject[] | undefined>;
    public readonly selectionMode!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly visibility!: pulumi.Output<string | undefined>;

    /**
     * Create a Portfolio resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PortfolioArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PortfolioArgs | PortfolioState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PortfolioState | undefined;
            resourceInputs["branch"] = state ? state.branch : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["qualifier"] = state ? state.qualifier : undefined;
            resourceInputs["regexp"] = state ? state.regexp : undefined;
            resourceInputs["selectedProjects"] = state ? state.selectedProjects : undefined;
            resourceInputs["selectionMode"] = state ? state.selectionMode : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as PortfolioArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["branch"] = args ? args.branch : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regexp"] = args ? args.regexp : undefined;
            resourceInputs["selectedProjects"] = args ? args.selectedProjects : undefined;
            resourceInputs["selectionMode"] = args ? args.selectionMode : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["qualifier"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Portfolio.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Portfolio resources.
 */
export interface PortfolioState {
    /**
     * Which branch to analyze. If nothing, or '' is specified, the main branch is used.
     */
    branch?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    key?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    qualifier?: pulumi.Input<string>;
    regexp?: pulumi.Input<string>;
    /**
     * A set of projects to add to the portfolio.
     */
    selectedProjects?: pulumi.Input<pulumi.Input<inputs.PortfolioSelectedProject>[]>;
    selectionMode?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Portfolio resource.
 */
export interface PortfolioArgs {
    /**
     * Which branch to analyze. If nothing, or '' is specified, the main branch is used.
     */
    branch?: pulumi.Input<string>;
    description: pulumi.Input<string>;
    key: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    regexp?: pulumi.Input<string>;
    /**
     * A set of projects to add to the portfolio.
     */
    selectedProjects?: pulumi.Input<pulumi.Input<inputs.PortfolioSelectedProject>[]>;
    selectionMode?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    visibility?: pulumi.Input<string>;
}
