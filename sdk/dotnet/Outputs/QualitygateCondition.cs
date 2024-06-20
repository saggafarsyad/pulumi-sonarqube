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
    public sealed class QualitygateCondition
    {
        public readonly string? Id;
        public readonly string Metric;
        public readonly string Op;
        public readonly string Threshold;

        [OutputConstructor]
        private QualitygateCondition(
            string? id,

            string metric,

            string op,

            string threshold)
        {
            Id = id;
            Metric = metric;
            Op = op;
            Threshold = threshold;
        }
    }
}
