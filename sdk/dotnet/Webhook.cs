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
    /// ## # sonarqube.Webhook
    /// 
    /// Provides a Sonarqube Webhook resource. This can be used to manage Sonarqube webhooks.
    /// 
    /// ## Example: create a webhook
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
    ///     var webhook = new Sonarqube.Webhook("webhook", new()
    ///     {
    ///         Url = "https://my-webhook-destination.example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Example: create a webhook owned by a project
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
    ///     var project = new Sonarqube.Project("project", new()
    ///     {
    ///         ProjectName = "project",
    ///         Visibility = "public",
    ///     });
    /// 
    ///     var webhook = new Sonarqube.Webhook("webhook", new()
    ///     {
    ///         Url = "https://my-webhook-destination.example.com",
    ///         Project = project.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Webhooks can be imported using their ID (key):
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import sonarqube:index/webhook:Webhook webhook AXnN9NuxdWLvsEEPOr2g
    /// ```
    /// </summary>
    [SonarqubeResourceType("sonarqube:index/webhook:Webhook")]
    public partial class Webhook : global::Pulumi.CustomResource
    {
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The key of the project that will own the webhook.
        /// </summary>
        [Output("project")]
        public Output<string?> Project { get; private set; } = null!;

        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Webhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Webhook(string name, WebhookArgs args, CustomResourceOptions? options = null)
            : base("sonarqube:index/webhook:Webhook", name, args ?? new WebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Webhook(string name, Input<string> id, WebhookState? state = null, CustomResourceOptions? options = null)
            : base("sonarqube:index/webhook:Webhook", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "secret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Webhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Webhook Get(string name, Input<string> id, WebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new Webhook(name, id, state, options);
        }
    }

    public sealed class WebhookArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The key of the project that will own the webhook.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("secret")]
        private Input<string>? _secret;
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public WebhookArgs()
        {
        }
        public static new WebhookArgs Empty => new WebhookArgs();
    }

    public sealed class WebhookState : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The key of the project that will own the webhook.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("secret")]
        private Input<string>? _secret;
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("url")]
        public Input<string>? Url { get; set; }

        public WebhookState()
        {
        }
        public static new WebhookState Empty => new WebhookState();
    }
}