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
    /// ## # sonarqube.Project
    /// 
    /// Provides a Sonarqube Project resource. This can be used to create and manage Sonarqube Project.
    /// 
    /// ## Example: create a project
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
    ///     var main = new Sonarqube.Project("main", new()
    ///     {
    ///         ProjectName = "my_project",
    ///         Visibility = "public",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Example: a project with associated settings
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
    ///     var main = new Sonarqube.Project("main", new()
    ///     {
    ///         ProjectName = "my_project",
    ///         Settings = new[]
    ///         {
    ///             new Sonarqube.Inputs.ProjectSettingArgs
    ///             {
    ///                 Key = "sonar.demo",
    ///                 Value = "sonarqube@example.org",
    ///             },
    ///         },
    ///         Visibility = "public",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Projects can be imported using their project key
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [SonarqubeResourceType("sonarqube:index/project:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// A list of settings associated to the project
        /// </summary>
        [Output("settings")]
        public Output<ImmutableArray<Outputs.ProjectSetting>> Settings { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("visibility")]
        public Output<string?> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs args, CustomResourceOptions? options = null)
            : base("sonarqube:index/project:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
            : base("sonarqube:index/project:Project", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new Project(name, id, state, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        [Input("settings")]
        private InputList<Inputs.ProjectSettingArgs>? _settings;

        /// <summary>
        /// A list of settings associated to the project
        /// </summary>
        public InputList<Inputs.ProjectSettingArgs> Settings
        {
            get => _settings ?? (_settings = new InputList<Inputs.ProjectSettingArgs>());
            set => _settings = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }

    public sealed class ProjectState : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? ProjectName { get; set; }

        [Input("settings")]
        private InputList<Inputs.ProjectSettingGetArgs>? _settings;

        /// <summary>
        /// A list of settings associated to the project
        /// </summary>
        public InputList<Inputs.ProjectSettingGetArgs> Settings
        {
            get => _settings ?? (_settings = new InputList<Inputs.ProjectSettingGetArgs>());
            set => _settings = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ProjectState()
        {
        }
        public static new ProjectState Empty => new ProjectState();
    }
}
