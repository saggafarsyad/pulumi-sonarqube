// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # sonarqube.AzureBinding
 *
 * Provides a Sonarqube Azure Devops binding resource. This can be used to create and manage the binding between an
 * Azure Devops repository and a SonarQube project
 *
 * ## Example: Create an Azure Devops binding
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as sonarqube from "@pulumi/sonarqube";
 *
 * const az1 = new sonarqube.AlmAzure("az1", {
 *     key: "az1",
 *     personalAccessToken: "my_pat",
 *     url: "https://dev.azure.com/my-org",
 * });
 * const mainProject = new sonarqube.Project("mainProject", {
 *     project: "main",
 *     visibility: "public",
 * });
 * const mainAzureBinding = new sonarqube.AzureBinding("mainAzureBinding", {
 *     almSetting: az1.key,
 *     project: mainProject.project,
 *     projectName: "my_azure_project",
 *     repositoryName: "my_repo",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Bindings can be imported using their ID
 *
 * terraform
 *
 * ```sh
 * $ pulumi import sonarqube:index/azureBinding:AzureBinding main project/project_name/repository
 * ```
 */
export class AzureBinding extends pulumi.CustomResource {
    /**
     * Get an existing AzureBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureBindingState, opts?: pulumi.CustomResourceOptions): AzureBinding {
        return new AzureBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sonarqube:index/azureBinding:AzureBinding';

    /**
     * Returns true if the given object is an instance of AzureBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureBinding.__pulumiType;
    }

    /**
     * Azure DevOps setting key
     */
    public readonly almSetting!: pulumi.Output<string>;
    /**
     * Is this project part of a monorepo
     */
    public readonly monorepo!: pulumi.Output<boolean | undefined>;
    /**
     * SonarQube project key
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Azure project name
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * Azure repository name
     */
    public readonly repositoryName!: pulumi.Output<string>;

    /**
     * Create a AzureBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureBindingArgs | AzureBindingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AzureBindingState | undefined;
            resourceInputs["almSetting"] = state ? state.almSetting : undefined;
            resourceInputs["monorepo"] = state ? state.monorepo : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["repositoryName"] = state ? state.repositoryName : undefined;
        } else {
            const args = argsOrState as AzureBindingArgs | undefined;
            if ((!args || args.almSetting === undefined) && !opts.urn) {
                throw new Error("Missing required property 'almSetting'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.projectName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectName'");
            }
            if ((!args || args.repositoryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryName'");
            }
            resourceInputs["almSetting"] = args ? args.almSetting : undefined;
            resourceInputs["monorepo"] = args ? args.monorepo : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["repositoryName"] = args ? args.repositoryName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AzureBinding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureBinding resources.
 */
export interface AzureBindingState {
    /**
     * Azure DevOps setting key
     */
    almSetting?: pulumi.Input<string>;
    /**
     * Is this project part of a monorepo
     */
    monorepo?: pulumi.Input<boolean>;
    /**
     * SonarQube project key
     */
    project?: pulumi.Input<string>;
    /**
     * Azure project name
     */
    projectName?: pulumi.Input<string>;
    /**
     * Azure repository name
     */
    repositoryName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AzureBinding resource.
 */
export interface AzureBindingArgs {
    /**
     * Azure DevOps setting key
     */
    almSetting: pulumi.Input<string>;
    /**
     * Is this project part of a monorepo
     */
    monorepo?: pulumi.Input<boolean>;
    /**
     * SonarQube project key
     */
    project: pulumi.Input<string>;
    /**
     * Azure project name
     */
    projectName: pulumi.Input<string>;
    /**
     * Azure repository name
     */
    repositoryName: pulumi.Input<string>;
}
