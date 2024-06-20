// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sonarqube
{
    /// <summary>
    /// ## # sonarqube.AlmAzure
    /// 
    /// Provides a Sonarqube Azure Devops Alm/Devops Platform Integration resource. This can be used to create and manage a Alm/Devops
    /// Platform Integration for Azure Devops.
    /// 
    /// ## Example: Create an Azure Devops Alm Integration
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Sonarqube = Pulumi.Sonarqube;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var az1 = new Sonarqube.AlmAzure("az1", new()
    ///     {
    ///         Key = "az1",
    ///         PersonalAccessToken = "my_pat",
    ///         Url = "https://dev.azure.com/my-org",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Resource can be imported providing their Alm Instance Key and Azure DevOps Personal Access Token
    /// 
    /// terraform
    /// 
    /// ```sh
    /// $ pulumi import sonarqube:index/almAzure:AlmAzure az1 key/personal_access_token
    /// ```
    /// </summary>
    [SonarqubeResourceType("sonarqube:index/almAzure:AlmAzure")]
    public partial class AlmAzure : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Unique key of the Azure Devops instance setting
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// Azure Devops personal access token
        /// </summary>
        [Output("personalAccessToken")]
        public Output<string> PersonalAccessToken { get; private set; } = null!;

        /// <summary>
        /// Azure API URL
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a AlmAzure resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlmAzure(string name, AlmAzureArgs args, CustomResourceOptions? options = null)
            : base("sonarqube:index/almAzure:AlmAzure", name, args ?? new AlmAzureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlmAzure(string name, Input<string> id, AlmAzureState? state = null, CustomResourceOptions? options = null)
            : base("sonarqube:index/almAzure:AlmAzure", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "personalAccessToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AlmAzure resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlmAzure Get(string name, Input<string> id, AlmAzureState? state = null, CustomResourceOptions? options = null)
        {
            return new AlmAzure(name, id, state, options);
        }
    }

    public sealed class AlmAzureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique key of the Azure Devops instance setting
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("personalAccessToken", required: true)]
        private Input<string>? _personalAccessToken;

        /// <summary>
        /// Azure Devops personal access token
        /// </summary>
        public Input<string>? PersonalAccessToken
        {
            get => _personalAccessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _personalAccessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Azure API URL
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public AlmAzureArgs()
        {
        }
        public static new AlmAzureArgs Empty => new AlmAzureArgs();
    }

    public sealed class AlmAzureState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique key of the Azure Devops instance setting
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("personalAccessToken")]
        private Input<string>? _personalAccessToken;

        /// <summary>
        /// Azure Devops personal access token
        /// </summary>
        public Input<string>? PersonalAccessToken
        {
            get => _personalAccessToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _personalAccessToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Azure API URL
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public AlmAzureState()
        {
        }
        public static new AlmAzureState Empty => new AlmAzureState();
    }
}
