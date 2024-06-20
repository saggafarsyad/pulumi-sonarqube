# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetQualityprofileResult',
    'AwaitableGetQualityprofileResult',
    'get_qualityprofile',
    'get_qualityprofile_output',
]

@pulumi.output_type
class GetQualityprofileResult:
    """
    A collection of values returned by getQualityprofile.
    """
    def __init__(__self__, id=None, is_default=None, key=None, language=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        pulumi.set(__self__, "is_default", is_default)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if language and not isinstance(language, str):
            raise TypeError("Expected argument 'language' to be a str")
        pulumi.set(__self__, "language", language)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> bool:
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def language(self) -> str:
        return pulumi.get(self, "language")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetQualityprofileResult(GetQualityprofileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQualityprofileResult(
            id=self.id,
            is_default=self.is_default,
            key=self.key,
            language=self.language,
            name=self.name)


def get_qualityprofile(name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetQualityprofileResult:
    """
    ## # Data Source: Qualityprofile

    Use this data source to get a Sonarqube qualityprofile resource

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_sonarqube as sonarqube

    main = sonarqube.get_qualityprofile(name="example")
    ```
    <!--End PulumiCodeChooser -->
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('sonarqube:index/getQualityprofile:getQualityprofile', __args__, opts=opts, typ=GetQualityprofileResult).value

    return AwaitableGetQualityprofileResult(
        id=pulumi.get(__ret__, 'id'),
        is_default=pulumi.get(__ret__, 'is_default'),
        key=pulumi.get(__ret__, 'key'),
        language=pulumi.get(__ret__, 'language'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_qualityprofile)
def get_qualityprofile_output(name: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetQualityprofileResult]:
    """
    ## # Data Source: Qualityprofile

    Use this data source to get a Sonarqube qualityprofile resource

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_sonarqube as sonarqube

    main = sonarqube.get_qualityprofile(name="example")
    ```
    <!--End PulumiCodeChooser -->
    """
    ...