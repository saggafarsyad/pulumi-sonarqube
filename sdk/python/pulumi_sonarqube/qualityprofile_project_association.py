# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['QualityprofileProjectAssociationArgs', 'QualityprofileProjectAssociation']

@pulumi.input_type
class QualityprofileProjectAssociationArgs:
    def __init__(__self__, *,
                 language: pulumi.Input[str],
                 project: pulumi.Input[str],
                 quality_profile: pulumi.Input[str]):
        """
        The set of arguments for constructing a QualityprofileProjectAssociation resource.
        :param pulumi.Input[str] language: Quality profile language
        :param pulumi.Input[str] project: Project key
        :param pulumi.Input[str] quality_profile: Quality profile name
        """
        pulumi.set(__self__, "language", language)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "quality_profile", quality_profile)

    @property
    @pulumi.getter
    def language(self) -> pulumi.Input[str]:
        """
        Quality profile language
        """
        return pulumi.get(self, "language")

    @language.setter
    def language(self, value: pulumi.Input[str]):
        pulumi.set(self, "language", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        Project key
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="qualityProfile")
    def quality_profile(self) -> pulumi.Input[str]:
        """
        Quality profile name
        """
        return pulumi.get(self, "quality_profile")

    @quality_profile.setter
    def quality_profile(self, value: pulumi.Input[str]):
        pulumi.set(self, "quality_profile", value)


@pulumi.input_type
class _QualityprofileProjectAssociationState:
    def __init__(__self__, *,
                 language: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 quality_profile: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering QualityprofileProjectAssociation resources.
        :param pulumi.Input[str] language: Quality profile language
        :param pulumi.Input[str] project: Project key
        :param pulumi.Input[str] quality_profile: Quality profile name
        """
        if language is not None:
            pulumi.set(__self__, "language", language)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if quality_profile is not None:
            pulumi.set(__self__, "quality_profile", quality_profile)

    @property
    @pulumi.getter
    def language(self) -> Optional[pulumi.Input[str]]:
        """
        Quality profile language
        """
        return pulumi.get(self, "language")

    @language.setter
    def language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        Project key
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="qualityProfile")
    def quality_profile(self) -> Optional[pulumi.Input[str]]:
        """
        Quality profile name
        """
        return pulumi.get(self, "quality_profile")

    @quality_profile.setter
    def quality_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quality_profile", value)


class QualityprofileProjectAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 language: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 quality_profile: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # QualityprofileProjectAssociation

        Provides a Sonarqube Quality Profile Project association resource. This can be used to associate a Quality Profile to a Project

        ## Example: create a quality profile project association

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        main_qualityprofile = sonarqube.Qualityprofile("mainQualityprofile", language="js")
        main_project = sonarqube.Project("mainProject",
            project="my_project",
            visibility="public")
        main_qualityprofile_project_association = sonarqube.QualityprofileProjectAssociation("mainQualityprofileProjectAssociation",
            quality_profile=main_qualityprofile.name,
            project=main_project.project,
            language="js")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] language: Quality profile language
        :param pulumi.Input[str] project: Project key
        :param pulumi.Input[str] quality_profile: Quality profile name
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QualityprofileProjectAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # QualityprofileProjectAssociation

        Provides a Sonarqube Quality Profile Project association resource. This can be used to associate a Quality Profile to a Project

        ## Example: create a quality profile project association

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        main_qualityprofile = sonarqube.Qualityprofile("mainQualityprofile", language="js")
        main_project = sonarqube.Project("mainProject",
            project="my_project",
            visibility="public")
        main_qualityprofile_project_association = sonarqube.QualityprofileProjectAssociation("mainQualityprofileProjectAssociation",
            quality_profile=main_qualityprofile.name,
            project=main_project.project,
            language="js")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param QualityprofileProjectAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QualityprofileProjectAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 language: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 quality_profile: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QualityprofileProjectAssociationArgs.__new__(QualityprofileProjectAssociationArgs)

            if language is None and not opts.urn:
                raise TypeError("Missing required property 'language'")
            __props__.__dict__["language"] = language
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if quality_profile is None and not opts.urn:
                raise TypeError("Missing required property 'quality_profile'")
            __props__.__dict__["quality_profile"] = quality_profile
        super(QualityprofileProjectAssociation, __self__).__init__(
            'sonarqube:index/qualityprofileProjectAssociation:QualityprofileProjectAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            language: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            quality_profile: Optional[pulumi.Input[str]] = None) -> 'QualityprofileProjectAssociation':
        """
        Get an existing QualityprofileProjectAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] language: Quality profile language
        :param pulumi.Input[str] project: Project key
        :param pulumi.Input[str] quality_profile: Quality profile name
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QualityprofileProjectAssociationState.__new__(_QualityprofileProjectAssociationState)

        __props__.__dict__["language"] = language
        __props__.__dict__["project"] = project
        __props__.__dict__["quality_profile"] = quality_profile
        return QualityprofileProjectAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def language(self) -> pulumi.Output[str]:
        """
        Quality profile language
        """
        return pulumi.get(self, "language")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        Project key
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="qualityProfile")
    def quality_profile(self) -> pulumi.Output[str]:
        """
        Quality profile name
        """
        return pulumi.get(self, "quality_profile")

