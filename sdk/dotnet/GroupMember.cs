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
    /// ## # sonarqube.GroupMember
    /// 
    /// Provides a Sonarqube Group Member resource. This can be used to add or remove user to or from Sonarqube Groups.
    /// 
    /// ## Example: add a user to a group
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
    ///     var user = new Sonarqube.User("user", new()
    ///     {
    ///         LoginName = "terraform-test",
    ///         Password = "secret-sauce37!",
    ///     });
    /// 
    ///     var projectUsers = new Sonarqube.Group("projectUsers", new()
    ///     {
    ///         Description = "This is a group",
    ///     });
    /// 
    ///     var projectUsersMember = new Sonarqube.GroupMember("projectUsersMember", new()
    ///     {
    ///         LoginName = user.LoginName,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Group Members can be imported using their ID (`&lt;name&gt;[&lt;login_name&gt;]`):
    /// 
    /// terraform
    /// 
    /// ```sh
    /// $ pulumi import sonarqube:index/groupMember:GroupMember member group[user]
    /// ```
    /// </summary>
    [SonarqubeResourceType("sonarqube:index/groupMember:GroupMember")]
    public partial class GroupMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The `login_name` of the User to add as a member. Changing this forces a new resource to be created.
        /// </summary>
        [Output("loginName")]
        public Output<string> LoginName { get; private set; } = null!;

        /// <summary>
        /// The name of the Group to add a member to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMember(string name, GroupMemberArgs args, CustomResourceOptions? options = null)
            : base("sonarqube:index/groupMember:GroupMember", name, args ?? new GroupMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupMember(string name, Input<string> id, GroupMemberState? state = null, CustomResourceOptions? options = null)
            : base("sonarqube:index/groupMember:GroupMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMember Get(string name, Input<string> id, GroupMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMember(name, id, state, options);
        }
    }

    public sealed class GroupMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The `login_name` of the User to add as a member. Changing this forces a new resource to be created.
        /// </summary>
        [Input("loginName", required: true)]
        public Input<string> LoginName { get; set; } = null!;

        /// <summary>
        /// The name of the Group to add a member to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupMemberArgs()
        {
        }
        public static new GroupMemberArgs Empty => new GroupMemberArgs();
    }

    public sealed class GroupMemberState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The `login_name` of the User to add as a member. Changing this forces a new resource to be created.
        /// </summary>
        [Input("loginName")]
        public Input<string>? LoginName { get; set; }

        /// <summary>
        /// The name of the Group to add a member to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupMemberState()
        {
        }
        public static new GroupMemberState Empty => new GroupMemberState();
    }
}