# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NewCodePeriodsArgs', 'NewCodePeriods']

@pulumi.input_type
class NewCodePeriodsArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 branch: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NewCodePeriods resource.
        """
        pulumi.set(__self__, "type", type)
        if branch is not None:
            pulumi.set(__self__, "branch", branch)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class _NewCodePeriodsState:
    def __init__(__self__, *,
                 branch: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NewCodePeriods resources.
        """
        if branch is not None:
            pulumi.set(__self__, "branch", branch)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class NewCodePeriods(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # NewCodePeriods

        Provides a Sonarqube New Code Periods resource. This can be used to manage Sonarqube New Code Periods.

        ## Example: Set the global new code period to a number of days

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        code_period = sonarqube.NewCodePeriods("codePeriod",
            type="NUMBER_OF_DAYS",
            value="7")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NewCodePeriodsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # NewCodePeriods

        Provides a Sonarqube New Code Periods resource. This can be used to manage Sonarqube New Code Periods.

        ## Example: Set the global new code period to a number of days

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        code_period = sonarqube.NewCodePeriods("codePeriod",
            type="NUMBER_OF_DAYS",
            value="7")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param NewCodePeriodsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NewCodePeriodsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NewCodePeriodsArgs.__new__(NewCodePeriodsArgs)

            __props__.__dict__["branch"] = branch
            __props__.__dict__["project"] = project
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["value"] = value
        super(NewCodePeriods, __self__).__init__(
            'sonarqube:index/newCodePeriods:NewCodePeriods',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'NewCodePeriods':
        """
        Get an existing NewCodePeriods resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NewCodePeriodsState.__new__(_NewCodePeriodsState)

        __props__.__dict__["branch"] = branch
        __props__.__dict__["project"] = project
        __props__.__dict__["type"] = type
        __props__.__dict__["value"] = value
        return NewCodePeriods(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "value")

