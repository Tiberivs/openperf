# coding: utf-8

"""
    OpenPerf API

    REST API interface for OpenPerf  # noqa: E501

    OpenAPI spec version: 1
    Contact: support@spirent.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class CpuGeneratorCoreLoad(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'instruction_set': 'str',
        'data_type': 'str'
    }

    attribute_map = {
        'instruction_set': 'instruction_set',
        'data_type': 'data_type'
    }

    def __init__(self, instruction_set=None, data_type=None):  # noqa: E501
        """CpuGeneratorCoreLoad - a model defined in Swagger"""  # noqa: E501

        self._instruction_set = None
        self._data_type = None
        self.discriminator = None

        self.instruction_set = instruction_set
        self.data_type = data_type

    @property
    def instruction_set(self):
        """Gets the instruction_set of this CpuGeneratorCoreLoad.  # noqa: E501

        CPU load instruction set  # noqa: E501

        :return: The instruction_set of this CpuGeneratorCoreLoad.  # noqa: E501
        :rtype: str
        """
        return self._instruction_set

    @instruction_set.setter
    def instruction_set(self, instruction_set):
        """Sets the instruction_set of this CpuGeneratorCoreLoad.

        CPU load instruction set  # noqa: E501

        :param instruction_set: The instruction_set of this CpuGeneratorCoreLoad.  # noqa: E501
        :type: str
        """
        self._instruction_set = instruction_set

    @property
    def data_type(self):
        """Gets the data_type of this CpuGeneratorCoreLoad.  # noqa: E501

        CPU load target operation data type, actual for chosen instruction set  # noqa: E501

        :return: The data_type of this CpuGeneratorCoreLoad.  # noqa: E501
        :rtype: str
        """
        return self._data_type

    @data_type.setter
    def data_type(self, data_type):
        """Sets the data_type of this CpuGeneratorCoreLoad.

        CPU load target operation data type, actual for chosen instruction set  # noqa: E501

        :param data_type: The data_type of this CpuGeneratorCoreLoad.  # noqa: E501
        :type: str
        """
        self._data_type = data_type

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(CpuGeneratorCoreLoad, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CpuGeneratorCoreLoad):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other