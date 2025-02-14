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


class TimeKeeperState(object):
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
        'frequency': 'int',
        'frequency_error': 'int',
        'local_frequency': 'int',
        'local_frequency_error': 'int',
        'offset': 'float',
        'synced': 'bool',
        'theta': 'float'
    }

    attribute_map = {
        'frequency': 'frequency',
        'frequency_error': 'frequency_error',
        'local_frequency': 'local_frequency',
        'local_frequency_error': 'local_frequency_error',
        'offset': 'offset',
        'synced': 'synced',
        'theta': 'theta'
    }

    def __init__(self, frequency=None, frequency_error=None, local_frequency=None, local_frequency_error=None, offset=None, synced=None, theta=None):  # noqa: E501
        """TimeKeeperState - a model defined in Swagger"""  # noqa: E501

        self._frequency = None
        self._frequency_error = None
        self._local_frequency = None
        self._local_frequency_error = None
        self._offset = None
        self._synced = None
        self._theta = None
        self.discriminator = None

        if frequency is not None:
            self.frequency = frequency
        if frequency_error is not None:
            self.frequency_error = frequency_error
        if local_frequency is not None:
            self.local_frequency = local_frequency
        if local_frequency_error is not None:
            self.local_frequency_error = local_frequency_error
        self.offset = offset
        self.synced = synced
        if theta is not None:
            self.theta = theta

    @property
    def frequency(self):
        """Gets the frequency of this TimeKeeperState.  # noqa: E501

        The time counter frequency as measured by the interval between the two best timestamp exchanges with the time source over the past two hours, in Hz.   # noqa: E501

        :return: The frequency of this TimeKeeperState.  # noqa: E501
        :rtype: int
        """
        return self._frequency

    @frequency.setter
    def frequency(self, frequency):
        """Sets the frequency of this TimeKeeperState.

        The time counter frequency as measured by the interval between the two best timestamp exchanges with the time source over the past two hours, in Hz.   # noqa: E501

        :param frequency: The frequency of this TimeKeeperState.  # noqa: E501
        :type: int
        """
        self._frequency = frequency

    @property
    def frequency_error(self):
        """Gets the frequency_error of this TimeKeeperState.  # noqa: E501

        The estimated error of the time counter frequency measurement, in Parts Per Billion (PPB).   # noqa: E501

        :return: The frequency_error of this TimeKeeperState.  # noqa: E501
        :rtype: int
        """
        return self._frequency_error

    @frequency_error.setter
    def frequency_error(self, frequency_error):
        """Sets the frequency_error of this TimeKeeperState.

        The estimated error of the time counter frequency measurement, in Parts Per Billion (PPB).   # noqa: E501

        :param frequency_error: The frequency_error of this TimeKeeperState.  # noqa: E501
        :type: int
        """
        self._frequency_error = frequency_error

    @property
    def local_frequency(self):
        """Gets the local_frequency of this TimeKeeperState.  # noqa: E501

        The time counter frequency as measured by the interval between the two best timestamp exchanges with the time source over the past hour, in Hz. This value is used to help determine time stamp error due to time counter frequency drift.   # noqa: E501

        :return: The local_frequency of this TimeKeeperState.  # noqa: E501
        :rtype: int
        """
        return self._local_frequency

    @local_frequency.setter
    def local_frequency(self, local_frequency):
        """Sets the local_frequency of this TimeKeeperState.

        The time counter frequency as measured by the interval between the two best timestamp exchanges with the time source over the past hour, in Hz. This value is used to help determine time stamp error due to time counter frequency drift.   # noqa: E501

        :param local_frequency: The local_frequency of this TimeKeeperState.  # noqa: E501
        :type: int
        """
        self._local_frequency = local_frequency

    @property
    def local_frequency_error(self):
        """Gets the local_frequency_error of this TimeKeeperState.  # noqa: E501

        The estimated error of the local time counter frequency measurement, int Parts Per Billion (PPB).   # noqa: E501

        :return: The local_frequency_error of this TimeKeeperState.  # noqa: E501
        :rtype: int
        """
        return self._local_frequency_error

    @local_frequency_error.setter
    def local_frequency_error(self, local_frequency_error):
        """Sets the local_frequency_error of this TimeKeeperState.

        The estimated error of the local time counter frequency measurement, int Parts Per Billion (PPB).   # noqa: E501

        :param local_frequency_error: The local_frequency_error of this TimeKeeperState.  # noqa: E501
        :type: int
        """
        self._local_frequency_error = local_frequency_error

    @property
    def offset(self):
        """Gets the offset of this TimeKeeperState.  # noqa: E501

        The offset applied to time counter derived timestamp values, in seconds.  This value comes from the system host clock.   # noqa: E501

        :return: The offset of this TimeKeeperState.  # noqa: E501
        :rtype: float
        """
        return self._offset

    @offset.setter
    def offset(self, offset):
        """Sets the offset of this TimeKeeperState.

        The offset applied to time counter derived timestamp values, in seconds.  This value comes from the system host clock.   # noqa: E501

        :param offset: The offset of this TimeKeeperState.  # noqa: E501
        :type: float
        """
        self._offset = offset

    @property
    def synced(self):
        """Gets the synced of this TimeKeeperState.  # noqa: E501

        The time keeper is considered to be synced to the time source if a clock offset, theta, has been calculated and applied within the past 20 minutes.   # noqa: E501

        :return: The synced of this TimeKeeperState.  # noqa: E501
        :rtype: bool
        """
        return self._synced

    @synced.setter
    def synced(self, synced):
        """Sets the synced of this TimeKeeperState.

        The time keeper is considered to be synced to the time source if a clock offset, theta, has been calculated and applied within the past 20 minutes.   # noqa: E501

        :param synced: The synced of this TimeKeeperState.  # noqa: E501
        :type: bool
        """
        self._synced = synced

    @property
    def theta(self):
        """Gets the theta of this TimeKeeperState.  # noqa: E501

        The calculated correction to apply to the offset, based on the measured time counter frequency and time source timestamps.   # noqa: E501

        :return: The theta of this TimeKeeperState.  # noqa: E501
        :rtype: float
        """
        return self._theta

    @theta.setter
    def theta(self, theta):
        """Sets the theta of this TimeKeeperState.

        The calculated correction to apply to the offset, based on the measured time counter frequency and time source timestamps.   # noqa: E501

        :param theta: The theta of this TimeKeeperState.  # noqa: E501
        :type: float
        """
        self._theta = theta

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
        if issubclass(TimeKeeperState, dict):
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
        if not isinstance(other, TimeKeeperState):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
