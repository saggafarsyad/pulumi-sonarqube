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
    'PortfolioSelectedProject',
    'ProjectSetting',
    'QualitygateCondition',
    'GetQualitygateConditionResult',
]

@pulumi.output_type
class PortfolioSelectedProject(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "projectKey":
            suggest = "project_key"
        elif key == "selectedBranches":
            suggest = "selected_branches"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PortfolioSelectedProject. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PortfolioSelectedProject.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PortfolioSelectedProject.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 project_key: str,
                 selected_branches: Optional[Sequence[str]] = None):
        """
        :param str project_key: The key of a project to add to the portfolio
        :param Sequence[str] selected_branches: A list of branches of the project to add to the portfolio. Defaults to the `MAIN BRANCH` of the repo if omitted
               
               Here is an example of how this option can be leveraged:
        """
        pulumi.set(__self__, "project_key", project_key)
        if selected_branches is not None:
            pulumi.set(__self__, "selected_branches", selected_branches)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> str:
        """
        The key of a project to add to the portfolio
        """
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter(name="selectedBranches")
    def selected_branches(self) -> Optional[Sequence[str]]:
        """
        A list of branches of the project to add to the portfolio. Defaults to the `MAIN BRANCH` of the repo if omitted

        Here is an example of how this option can be leveraged:
        """
        return pulumi.get(self, "selected_branches")


@pulumi.output_type
class ProjectSetting(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fieldValues":
            suggest = "field_values"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ProjectSetting. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ProjectSetting.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ProjectSetting.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 key: str,
                 field_values: Optional[Sequence[Mapping[str, Any]]] = None,
                 value: Optional[str] = None,
                 values: Optional[Sequence[str]] = None):
        """
        :param str key: Setting key
        :param Sequence[Mapping[str, Any]] field_values: Setting field values for the supplied key
        :param str value: Setting a value for the supplied key
        :param Sequence[str] values: Setting multi values for the supplied key
        """
        pulumi.set(__self__, "key", key)
        if field_values is not None:
            pulumi.set(__self__, "field_values", field_values)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Setting key
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="fieldValues")
    def field_values(self) -> Optional[Sequence[Mapping[str, Any]]]:
        """
        Setting field values for the supplied key
        """
        return pulumi.get(self, "field_values")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        Setting a value for the supplied key
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter
    def values(self) -> Optional[Sequence[str]]:
        """
        Setting multi values for the supplied key
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class QualitygateCondition(dict):
    def __init__(__self__, *,
                 metric: str,
                 op: str,
                 threshold: str,
                 id: Optional[str] = None):
        pulumi.set(__self__, "metric", metric)
        pulumi.set(__self__, "op", op)
        pulumi.set(__self__, "threshold", threshold)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def metric(self) -> str:
        return pulumi.get(self, "metric")

    @property
    @pulumi.getter
    def op(self) -> str:
        return pulumi.get(self, "op")

    @property
    @pulumi.getter
    def threshold(self) -> str:
        return pulumi.get(self, "threshold")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


@pulumi.output_type
class GetQualitygateConditionResult(dict):
    def __init__(__self__, *,
                 id: str,
                 metric: str,
                 op: str,
                 threshold: str):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "metric", metric)
        pulumi.set(__self__, "op", op)
        pulumi.set(__self__, "threshold", threshold)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def metric(self) -> str:
        return pulumi.get(self, "metric")

    @property
    @pulumi.getter
    def op(self) -> str:
        return pulumi.get(self, "op")

    @property
    @pulumi.getter
    def threshold(self) -> str:
        return pulumi.get(self, "threshold")


