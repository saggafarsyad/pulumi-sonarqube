// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sonarqube.Outputs
{

    [OutputType]
    public sealed class PortfolioSelectedProject
    {
        /// <summary>
        /// The key of a project to add to the portfolio
        /// </summary>
        public readonly string ProjectKey;
        /// <summary>
        /// A list of branches of the project to add to the portfolio. Defaults to the `MAIN BRANCH` of the repo if omitted
        /// 
        /// Here is an example of how this option can be leveraged:
        /// </summary>
        public readonly ImmutableArray<string> SelectedBranches;

        [OutputConstructor]
        private PortfolioSelectedProject(
            string projectKey,

            ImmutableArray<string> selectedBranches)
        {
            ProjectKey = projectKey;
            SelectedBranches = selectedBranches;
        }
    }
}
