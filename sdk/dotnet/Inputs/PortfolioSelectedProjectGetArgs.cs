// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sonarqube.Inputs
{

    public sealed class PortfolioSelectedProjectGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key of a project to add to the portfolio
        /// </summary>
        [Input("projectKey", required: true)]
        public Input<string> ProjectKey { get; set; } = null!;

        [Input("selectedBranches")]
        private InputList<string>? _selectedBranches;

        /// <summary>
        /// A list of branches of the project to add to the portfolio. Defaults to the `MAIN BRANCH` of the repo if omitted
        /// 
        /// Here is an example of how this option can be leveraged:
        /// </summary>
        public InputList<string> SelectedBranches
        {
            get => _selectedBranches ?? (_selectedBranches = new InputList<string>());
            set => _selectedBranches = value;
        }

        public PortfolioSelectedProjectGetArgs()
        {
        }
        public static new PortfolioSelectedProjectGetArgs Empty => new PortfolioSelectedProjectGetArgs();
    }
}