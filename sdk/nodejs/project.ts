// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # sonarqube.Project
 *
 * Provides a Sonarqube Project resource. This can be used to create and manage Sonarqube Project.
 *
 * ## Example: create a project
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sonarqube from "@pulumi/sonarqube";
 *
 * const main = new sonarqube.Project("main", {
 *     project: "my_project",
 *     visibility: "public",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Example: a project with associated settings
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sonarqube from "@pulumi/sonarqube";
 *
 * const main = new sonarqube.Project("main", {
 *     project: "my_project",
 *     settings: [{
 *         key: "sonar.demo",
 *         value: "sonarqube@example.org",
 *     }],
 *     visibility: "public",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Projects can be imported using their project key
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sonarqube:index/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * A list of settings associated to the project
     */
    public readonly settings!: pulumi.Output<outputs.ProjectSetting[] | undefined>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly visibility!: pulumi.Output<string | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * A list of settings associated to the project
     */
    settings?: pulumi.Input<pulumi.Input<inputs.ProjectSetting>[]>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    /**
     * A list of settings associated to the project
     */
    settings?: pulumi.Input<pulumi.Input<inputs.ProjectSetting>[]>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    visibility?: pulumi.Input<string>;
}