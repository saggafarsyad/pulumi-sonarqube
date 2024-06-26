# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AlmGitlabArgs', 'AlmGitlab']

@pulumi.input_type
class AlmGitlabArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 personal_access_token: pulumi.Input[str],
                 url: pulumi.Input[str]):
        """
        The set of arguments for constructing a AlmGitlab resource.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "personal_access_token", personal_access_token)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="personalAccessToken")
    def personal_access_token(self) -> pulumi.Input[str]:
        return pulumi.get(self, "personal_access_token")

    @personal_access_token.setter
    def personal_access_token(self, value: pulumi.Input[str]):
        pulumi.set(self, "personal_access_token", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class _AlmGitlabState:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 personal_access_token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AlmGitlab resources.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if personal_access_token is not None:
            pulumi.set(__self__, "personal_access_token", personal_access_token)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="personalAccessToken")
    def personal_access_token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "personal_access_token")

    @personal_access_token.setter
    def personal_access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "personal_access_token", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class AlmGitlab(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 personal_access_token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # AlmGitlab

        Provides a Sonarqube GitLab Alm/Devops Platform Integration resource. This can be used to create and manage a Alm/Devops
        Platform Integration for GitLab.

        ## Example: Create a GitHub Alm Integration

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        gitlab_alm = sonarqube.AlmGitlab("gitlab-alm",
            key="myalm",
            personal_access_token="my_personal_access_token",
            url="https://gitlab.com/api/v4")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AlmGitlabArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # AlmGitlab

        Provides a Sonarqube GitLab Alm/Devops Platform Integration resource. This can be used to create and manage a Alm/Devops
        Platform Integration for GitLab.

        ## Example: Create a GitHub Alm Integration

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        gitlab_alm = sonarqube.AlmGitlab("gitlab-alm",
            key="myalm",
            personal_access_token="my_personal_access_token",
            url="https://gitlab.com/api/v4")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param AlmGitlabArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlmGitlabArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 personal_access_token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AlmGitlabArgs.__new__(AlmGitlabArgs)

            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            if personal_access_token is None and not opts.urn:
                raise TypeError("Missing required property 'personal_access_token'")
            __props__.__dict__["personal_access_token"] = None if personal_access_token is None else pulumi.Output.secret(personal_access_token)
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["personalAccessToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AlmGitlab, __self__).__init__(
            'sonarqube:index/almGitlab:AlmGitlab',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key: Optional[pulumi.Input[str]] = None,
            personal_access_token: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'AlmGitlab':
        """
        Get an existing AlmGitlab resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AlmGitlabState.__new__(_AlmGitlabState)

        __props__.__dict__["key"] = key
        __props__.__dict__["personal_access_token"] = personal_access_token
        __props__.__dict__["url"] = url
        return AlmGitlab(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="personalAccessToken")
    def personal_access_token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "personal_access_token")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")

