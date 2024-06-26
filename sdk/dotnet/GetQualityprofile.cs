// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sonarqube
{
    public static class GetQualityprofile
    {
        /// <summary>
        /// ## # Data Source: sonarqube.Qualityprofile
        /// 
        /// Use this data source to get a Sonarqube qualityprofile resource
        /// 
        /// ## Example Usage
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
        ///     var main = Sonarqube.GetQualityprofile.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetQualityprofileResult> InvokeAsync(GetQualityprofileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQualityprofileResult>("sonarqube:index/getQualityprofile:getQualityprofile", args ?? new GetQualityprofileArgs(), options.WithDefaults());

        /// <summary>
        /// ## # Data Source: sonarqube.Qualityprofile
        /// 
        /// Use this data source to get a Sonarqube qualityprofile resource
        /// 
        /// ## Example Usage
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
        ///     var main = Sonarqube.GetQualityprofile.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetQualityprofileResult> Invoke(GetQualityprofileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQualityprofileResult>("sonarqube:index/getQualityprofile:getQualityprofile", args ?? new GetQualityprofileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQualityprofileArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetQualityprofileArgs()
        {
        }
        public static new GetQualityprofileArgs Empty => new GetQualityprofileArgs();
    }

    public sealed class GetQualityprofileInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetQualityprofileInvokeArgs()
        {
        }
        public static new GetQualityprofileInvokeArgs Empty => new GetQualityprofileInvokeArgs();
    }


    [OutputType]
    public sealed class GetQualityprofileResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsDefault;
        public readonly string Key;
        public readonly string Language;
        public readonly string Name;

        [OutputConstructor]
        private GetQualityprofileResult(
            string id,

            bool isDefault,

            string key,

            string language,

            string name)
        {
            Id = id;
            IsDefault = isDefault;
            Key = key;
            Language = language;
            Name = name;
        }
    }
}
