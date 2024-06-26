# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GithubBindingArgs', 'GithubBinding']

@pulumi.input_type
class GithubBindingArgs:
    def __init__(__self__, *,
                 alm_setting: pulumi.Input[str],
                 project: pulumi.Input[str],
                 repository: pulumi.Input[str],
                 monorepo: Optional[pulumi.Input[str]] = None,
                 summary_comment_enabled: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GithubBinding resource.
        """
        pulumi.set(__self__, "alm_setting", alm_setting)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "repository", repository)
        if monorepo is not None:
            pulumi.set(__self__, "monorepo", monorepo)
        if summary_comment_enabled is not None:
            pulumi.set(__self__, "summary_comment_enabled", summary_comment_enabled)

    @property
    @pulumi.getter(name="almSetting")
    def alm_setting(self) -> pulumi.Input[str]:
        return pulumi.get(self, "alm_setting")

    @alm_setting.setter
    def alm_setting(self, value: pulumi.Input[str]):
        pulumi.set(self, "alm_setting", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def monorepo(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "monorepo")

    @monorepo.setter
    def monorepo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "monorepo", value)

    @property
    @pulumi.getter(name="summaryCommentEnabled")
    def summary_comment_enabled(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "summary_comment_enabled")

    @summary_comment_enabled.setter
    def summary_comment_enabled(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "summary_comment_enabled", value)


@pulumi.input_type
class _GithubBindingState:
    def __init__(__self__, *,
                 alm_setting: Optional[pulumi.Input[str]] = None,
                 monorepo: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 summary_comment_enabled: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GithubBinding resources.
        """
        if alm_setting is not None:
            pulumi.set(__self__, "alm_setting", alm_setting)
        if monorepo is not None:
            pulumi.set(__self__, "monorepo", monorepo)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if summary_comment_enabled is not None:
            pulumi.set(__self__, "summary_comment_enabled", summary_comment_enabled)

    @property
    @pulumi.getter(name="almSetting")
    def alm_setting(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "alm_setting")

    @alm_setting.setter
    def alm_setting(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alm_setting", value)

    @property
    @pulumi.getter
    def monorepo(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "monorepo")

    @monorepo.setter
    def monorepo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "monorepo", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="summaryCommentEnabled")
    def summary_comment_enabled(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "summary_comment_enabled")

    @summary_comment_enabled.setter
    def summary_comment_enabled(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "summary_comment_enabled", value)


class GithubBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alm_setting: Optional[pulumi.Input[str]] = None,
                 monorepo: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 summary_comment_enabled: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # GithubBinding

        Provides a Sonarqube GitHub binding resource. This can be used to create and manage the binding between a
        GitHub repository and a SonarQube project

        ## Example: Create a GitHub binding

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        github_alm = sonarqube.AlmGithub("github-alm",
            app_id="12345",
            client_id="56789",
            client_secret="secret",
            key="myalm",
            private_key="myprivate_key",
            url="https://api.github.com",
            webhook_secret="mysecret")
        main = sonarqube.Project("main",
            project="my_project",
            visibility="public")
        github_binding = sonarqube.GithubBinding("github-binding",
            alm_setting=github_alm.key,
            project="my_project",
            repository="myorg/myrepo")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Bindings can be imported using their ID

        terraform

        ```sh
        $ pulumi import sonarqube:index/githubBinding:GithubBinding github-binding project/repository
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GithubBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # GithubBinding

        Provides a Sonarqube GitHub binding resource. This can be used to create and manage the binding between a
        GitHub repository and a SonarQube project

        ## Example: Create a GitHub binding

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        github_alm = sonarqube.AlmGithub("github-alm",
            app_id="12345",
            client_id="56789",
            client_secret="secret",
            key="myalm",
            private_key="myprivate_key",
            url="https://api.github.com",
            webhook_secret="mysecret")
        main = sonarqube.Project("main",
            project="my_project",
            visibility="public")
        github_binding = sonarqube.GithubBinding("github-binding",
            alm_setting=github_alm.key,
            project="my_project",
            repository="myorg/myrepo")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Bindings can be imported using their ID

        terraform

        ```sh
        $ pulumi import sonarqube:index/githubBinding:GithubBinding github-binding project/repository
        ```

        :param str resource_name: The name of the resource.
        :param GithubBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GithubBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alm_setting: Optional[pulumi.Input[str]] = None,
                 monorepo: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 summary_comment_enabled: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GithubBindingArgs.__new__(GithubBindingArgs)

            if alm_setting is None and not opts.urn:
                raise TypeError("Missing required property 'alm_setting'")
            __props__.__dict__["alm_setting"] = alm_setting
            __props__.__dict__["monorepo"] = monorepo
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["summary_comment_enabled"] = summary_comment_enabled
        super(GithubBinding, __self__).__init__(
            'sonarqube:index/githubBinding:GithubBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alm_setting: Optional[pulumi.Input[str]] = None,
            monorepo: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            summary_comment_enabled: Optional[pulumi.Input[str]] = None) -> 'GithubBinding':
        """
        Get an existing GithubBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GithubBindingState.__new__(_GithubBindingState)

        __props__.__dict__["alm_setting"] = alm_setting
        __props__.__dict__["monorepo"] = monorepo
        __props__.__dict__["project"] = project
        __props__.__dict__["repository"] = repository
        __props__.__dict__["summary_comment_enabled"] = summary_comment_enabled
        return GithubBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="almSetting")
    def alm_setting(self) -> pulumi.Output[str]:
        return pulumi.get(self, "alm_setting")

    @property
    @pulumi.getter
    def monorepo(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "monorepo")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter(name="summaryCommentEnabled")
    def summary_comment_enabled(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "summary_comment_enabled")

