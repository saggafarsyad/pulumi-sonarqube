# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserTokenArgs', 'UserToken']

@pulumi.input_type
class UserTokenArgs:
    def __init__(__self__, *,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UserToken resource.
        """
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if login_name is not None:
            pulumi.set(__self__, "login_name", login_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="loginName")
    def login_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "login_name")

    @login_name.setter
    def login_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _UserTokenState:
    def __init__(__self__, *,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserToken resources.
        """
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if login_name is not None:
            pulumi.set(__self__, "login_name", login_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="loginName")
    def login_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "login_name")

    @login_name.setter
    def login_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class UserToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # UserToken

        Provides a Sonarqube User token resource. This can be used to manage Sonarqube User tokens.

        ## Example: create a user, user token and output the token value

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        user = sonarqube.User("user",
            login_name="terraform-test",
            password="secret-sauce37!")
        token = sonarqube.UserToken("token", login_name=user.login_name)
        pulumi.export("userToken", token.token)
        ```
        <!--End PulumiCodeChooser -->

        ## Example: create an expiring global analysis token and output the token value

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        token = sonarqube.UserToken("token",
            type="GLOBAL_ANALYSIS_TOKEN",
            expiration_date="2099-01-01")
        pulumi.export("globalAnalysisToken", token.token)
        ```
        <!--End PulumiCodeChooser -->

        ## Example: create a project, project analysis token, and output the token value

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        token = sonarqube.UserToken("token",
            type="PROJECT_ANALYSIS_TOKEN",
            project_key="my-project")
        pulumi.export("projectAnalysisToken", token.token)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Import is not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UserTokenArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # UserToken

        Provides a Sonarqube User token resource. This can be used to manage Sonarqube User tokens.

        ## Example: create a user, user token and output the token value

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        user = sonarqube.User("user",
            login_name="terraform-test",
            password="secret-sauce37!")
        token = sonarqube.UserToken("token", login_name=user.login_name)
        pulumi.export("userToken", token.token)
        ```
        <!--End PulumiCodeChooser -->

        ## Example: create an expiring global analysis token and output the token value

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        token = sonarqube.UserToken("token",
            type="GLOBAL_ANALYSIS_TOKEN",
            expiration_date="2099-01-01")
        pulumi.export("globalAnalysisToken", token.token)
        ```
        <!--End PulumiCodeChooser -->

        ## Example: create a project, project analysis token, and output the token value

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_sonarqube as sonarqube

        token = sonarqube.UserToken("token",
            type="PROJECT_ANALYSIS_TOKEN",
            project_key="my-project")
        pulumi.export("projectAnalysisToken", token.token)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Import is not supported for this resource.

        :param str resource_name: The name of the resource.
        :param UserTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 login_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserTokenArgs.__new__(UserTokenArgs)

            __props__.__dict__["expiration_date"] = expiration_date
            __props__.__dict__["login_name"] = login_name
            __props__.__dict__["name"] = name
            __props__.__dict__["project_key"] = project_key
            __props__.__dict__["type"] = type
            __props__.__dict__["token"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(UserToken, __self__).__init__(
            'sonarqube:index/userToken:UserToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expiration_date: Optional[pulumi.Input[str]] = None,
            login_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_key: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'UserToken':
        """
        Get an existing UserToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserTokenState.__new__(_UserTokenState)

        __props__.__dict__["expiration_date"] = expiration_date
        __props__.__dict__["login_name"] = login_name
        __props__.__dict__["name"] = name
        __props__.__dict__["project_key"] = project_key
        __props__.__dict__["token"] = token
        __props__.__dict__["type"] = type
        return UserToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> pulumi.Output[str]:
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter(name="loginName")
    def login_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "login_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "type")

