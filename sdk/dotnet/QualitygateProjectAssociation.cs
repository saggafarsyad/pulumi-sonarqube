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
    /// ## # sonarqube.QualitygateProjectAssociation
    /// 
    /// Provides a Sonarqube Quality Gate Project association resource. This can be used to associate a Quality Gate to a Project
    /// 
    /// ## Example: create a quality gate project association
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
    ///     var mainQualitygate = new Sonarqube.Qualitygate("mainQualitygate", new()
    ///     {
    ///         Conditions = new[]
    ///         {
    ///             new Sonarqube.Inputs.QualitygateConditionArgs
    ///             {
    ///                 Metric = "new_coverage",
    ///                 Op = "LT",
    ///                 Threshold = "30",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var mainProject = new Sonarqube.Project("mainProject", new()
    ///     {
    ///         ProjectName = "my_project",
    ///         Visibility = "public",
    ///     });
    /// 
    ///     var mainQualitygateProjectAssociation = new Sonarqube.QualitygateProjectAssociation("mainQualitygateProjectAssociation", new()
    ///     {
    ///         Gatename = mainQualitygate.Id,
    ///         Projectkey = mainProject.ProjectName,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [SonarqubeResourceType("sonarqube:index/qualitygateProjectAssociation:QualitygateProjectAssociation")]
    public partial class QualitygateProjectAssociation : global::Pulumi.CustomResource
    {
        [Output("gateid")]
        public Output<string?> Gateid { get; private set; } = null!;

        [Output("gatename")]
        public Output<string?> Gatename { get; private set; } = null!;

        [Output("projectkey")]
        public Output<string> Projectkey { get; private set; } = null!;


        /// <summary>
        /// Create a QualitygateProjectAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QualitygateProjectAssociation(string name, QualitygateProjectAssociationArgs args, CustomResourceOptions? options = null)
            : base("sonarqube:index/qualitygateProjectAssociation:QualitygateProjectAssociation", name, args ?? new QualitygateProjectAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QualitygateProjectAssociation(string name, Input<string> id, QualitygateProjectAssociationState? state = null, CustomResourceOptions? options = null)
            : base("sonarqube:index/qualitygateProjectAssociation:QualitygateProjectAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QualitygateProjectAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QualitygateProjectAssociation Get(string name, Input<string> id, QualitygateProjectAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new QualitygateProjectAssociation(name, id, state, options);
        }
    }

    public sealed class QualitygateProjectAssociationArgs : global::Pulumi.ResourceArgs
    {
        [Input("gateid")]
        public Input<string>? Gateid { get; set; }

        [Input("gatename")]
        public Input<string>? Gatename { get; set; }

        [Input("projectkey", required: true)]
        public Input<string> Projectkey { get; set; } = null!;

        public QualitygateProjectAssociationArgs()
        {
        }
        public static new QualitygateProjectAssociationArgs Empty => new QualitygateProjectAssociationArgs();
    }

    public sealed class QualitygateProjectAssociationState : global::Pulumi.ResourceArgs
    {
        [Input("gateid")]
        public Input<string>? Gateid { get; set; }

        [Input("gatename")]
        public Input<string>? Gatename { get; set; }

        [Input("projectkey")]
        public Input<string>? Projectkey { get; set; }

        public QualitygateProjectAssociationState()
        {
        }
        public static new QualitygateProjectAssociationState Empty => new QualitygateProjectAssociationState();
    }
}
