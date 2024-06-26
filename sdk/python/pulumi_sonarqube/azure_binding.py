# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AzureBindingArgs', 'AzureBinding']

@pulumi.input_type
class AzureBindingArgs:
    def __init__(__self__, *,
                 alm_setting: pulumi.Input[str],
                 project: pulumi.Input[str],
                 project_name: pulumi.Input[str],
                 repository_name: pulumi.Input[str],
                 monorepo: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AzureBinding resource.
        :param pulumi.Input[str] alm_setting: Azure DevOps setting key
        :param pulumi.Input[str] project: SonarQube project key
        :param pulumi.Input[str] project_name: Azure project name
        :param pulumi.Input[str] repository_name: Azure repository name
        :param pulumi.Input[bool] monorepo: Is this project part of a monorepo
        """
        pulumi.set(__self__, "alm_setting", alm_setting)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "project_name", project_name)
        pulumi.set(__self__, "repository_name", repository_name)
        if monorepo is not None:
            pulumi.set(__self__, "monorepo", monorepo)

    @property
    @pulumi.getter(name="almSetting")
    def alm_setting(self) -> pulumi.Input[str]:
        """
        Azure DevOps setting key
        """
        return pulumi.get(self, "alm_setting")

    @alm_setting.setter
    def alm_setting(self, value: pulumi.Input[str]):
        pulumi.set(self, "alm_setting", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        SonarQube project key
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Input[str]:
        """
        Azure project name
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> pulumi.Input[str]:
        """
        Azure repository name
        """
        return pulumi.get(self, "repository_name")

    @repository_name.setter
    def repository_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_name", value)

    @property
    @pulumi.getter
    def monorepo(self) -> Optional[pulumi.Input[bool]]:
        """
        Is this project part of a monorepo
        """
        return pulumi.get(self, "monorepo")

    @monorepo.setter
    def monorepo(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "monorepo", value)


@pulumi.input_type
class _AzureBindingState:
    def __init__(__self__, *,
                 alm_setting: Optional[pulumi.Input[str]] = None,
                 monorepo: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 repository_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AzureBinding resources.
        :param pulumi.Input[str] alm_setting: Azure DevOps setting key
        :param pulumi.Input[bool] monorepo: Is this project part of a monorepo
        :param pulumi.Input[str] project: SonarQube project key
        :param pulumi.Input[str] project_name: Azure project name
        :param pulumi.Input[str] repository_name: Azure repository name
        """
        if alm_setting is not None:
            pulumi.set(__self__, "alm_setting", alm_setting)
        if monorepo is not None:
            pulumi.set(__self__, "monorepo", monorepo)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if repository_name is not None:
            pulumi.set(__self__, "repository_name", repository_name)

    @property
    @pulumi.getter(name="almSetting")
    def alm_setting(self) -> Optional[pulumi.Input[str]]:
        """
        Azure DevOps setting key
        """
        return pulumi.get(self, "alm_setting")

    @alm_setting.setter
    def alm_setting(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alm_setting", value)

    @property
    @pulumi.getter
    def monorepo(self) -> Optional[pulumi.Input[bool]]:
        """
        Is this project part of a monorepo
        """
        return pulumi.get(self, "monorepo")

    @monorepo.setter
    def monorepo(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "monorepo", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        SonarQube project key
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        Azure project name
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> Optional[pulumi.Input[str]]:
        """
        Azure repository name
        """
        return pulumi.get(self, "repository_name")

    @repository_name.setter
    def repository_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_name", value)


class AzureBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alm_setting: Optional[pulumi.Input[str]] = None,
                 monorepo: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 repository_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # AzureBinding

        Provides a Sonarqube Azure Devops binding resource. This can be used to create and manage the binding between an
        Azure Devops repository and a SonarQube project

        ## Example: Create an Azure Devops binding

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        az1 = sonarqube.AlmAzure("az1",
            key="az1",
            personal_access_token="my_pat",
            url="https://dev.azure.com/my-org")
        main_project = sonarqube.Project("mainProject",
            project="main",
            visibility="public")
        main_azure_binding = sonarqube.AzureBinding("mainAzureBinding",
            alm_setting=az1.key,
            project=main_project.project,
            project_name="my_azure_project",
            repository_name="my_repo")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Bindings can be imported using their ID

        terraform

        ```sh
        $ pulumi import sonarqube:index/azureBinding:AzureBinding main project/project_name/repository
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alm_setting: Azure DevOps setting key
        :param pulumi.Input[bool] monorepo: Is this project part of a monorepo
        :param pulumi.Input[str] project: SonarQube project key
        :param pulumi.Input[str] project_name: Azure project name
        :param pulumi.Input[str] repository_name: Azure repository name
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AzureBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # AzureBinding

        Provides a Sonarqube Azure Devops binding resource. This can be used to create and manage the binding between an
        Azure Devops repository and a SonarQube project

        ## Example: Create an Azure Devops binding

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        az1 = sonarqube.AlmAzure("az1",
            key="az1",
            personal_access_token="my_pat",
            url="https://dev.azure.com/my-org")
        main_project = sonarqube.Project("mainProject",
            project="main",
            visibility="public")
        main_azure_binding = sonarqube.AzureBinding("mainAzureBinding",
            alm_setting=az1.key,
            project=main_project.project,
            project_name="my_azure_project",
            repository_name="my_repo")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Bindings can be imported using their ID

        terraform

        ```sh
        $ pulumi import sonarqube:index/azureBinding:AzureBinding main project/project_name/repository
        ```

        :param str resource_name: The name of the resource.
        :param AzureBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AzureBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alm_setting: Optional[pulumi.Input[str]] = None,
                 monorepo: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 repository_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AzureBindingArgs.__new__(AzureBindingArgs)

            if alm_setting is None and not opts.urn:
                raise TypeError("Missing required property 'alm_setting'")
            __props__.__dict__["alm_setting"] = alm_setting
            __props__.__dict__["monorepo"] = monorepo
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if project_name is None and not opts.urn:
                raise TypeError("Missing required property 'project_name'")
            __props__.__dict__["project_name"] = project_name
            if repository_name is None and not opts.urn:
                raise TypeError("Missing required property 'repository_name'")
            __props__.__dict__["repository_name"] = repository_name
        super(AzureBinding, __self__).__init__(
            'sonarqube:index/azureBinding:AzureBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alm_setting: Optional[pulumi.Input[str]] = None,
            monorepo: Optional[pulumi.Input[bool]] = None,
            project: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            repository_name: Optional[pulumi.Input[str]] = None) -> 'AzureBinding':
        """
        Get an existing AzureBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alm_setting: Azure DevOps setting key
        :param pulumi.Input[bool] monorepo: Is this project part of a monorepo
        :param pulumi.Input[str] project: SonarQube project key
        :param pulumi.Input[str] project_name: Azure project name
        :param pulumi.Input[str] repository_name: Azure repository name
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AzureBindingState.__new__(_AzureBindingState)

        __props__.__dict__["alm_setting"] = alm_setting
        __props__.__dict__["monorepo"] = monorepo
        __props__.__dict__["project"] = project
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["repository_name"] = repository_name
        return AzureBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="almSetting")
    def alm_setting(self) -> pulumi.Output[str]:
        """
        Azure DevOps setting key
        """
        return pulumi.get(self, "alm_setting")

    @property
    @pulumi.getter
    def monorepo(self) -> pulumi.Output[Optional[bool]]:
        """
        Is this project part of a monorepo
        """
        return pulumi.get(self, "monorepo")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        SonarQube project key
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        Azure project name
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> pulumi.Output[str]:
        """
        Azure repository name
        """
        return pulumi.get(self, "repository_name")

