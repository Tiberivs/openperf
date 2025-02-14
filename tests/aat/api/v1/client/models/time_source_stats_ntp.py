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


class TimeSourceStatsNtp(object):
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
        'last_rx_accepted': 'datetime',
        'last_rx_ignored': 'datetime',
        'poll_period': 'int',
        'rx_ignored': 'int',
        'rx_packets': 'int',
        'tx_packets': 'int',
        'stratum': 'int'
    }

    attribute_map = {
        'last_rx_accepted': 'last_rx_accepted',
        'last_rx_ignored': 'last_rx_ignored',
        'poll_period': 'poll_period',
        'rx_ignored': 'rx_ignored',
        'rx_packets': 'rx_packets',
        'tx_packets': 'tx_packets',
        'stratum': 'stratum'
    }

    def __init__(self, last_rx_accepted=None, last_rx_ignored=None, poll_period=None, rx_ignored=None, rx_packets=None, tx_packets=None, stratum=None):  # noqa: E501
        """TimeSourceStatsNtp - a model defined in Swagger"""  # noqa: E501

        self._last_rx_accepted = None
        self._last_rx_ignored = None
        self._poll_period = None
        self._rx_ignored = None
        self._rx_packets = None
        self._tx_packets = None
        self._stratum = None
        self.discriminator = None

        if last_rx_accepted is not None:
            self.last_rx_accepted = last_rx_accepted
        if last_rx_ignored is not None:
            self.last_rx_ignored = last_rx_ignored
        self.poll_period = poll_period
        self.rx_ignored = rx_ignored
        self.rx_packets = rx_packets
        self.tx_packets = tx_packets
        if stratum is not None:
            self.stratum = stratum

    @property
    def last_rx_accepted(self):
        """Gets the last_rx_accepted of this TimeSourceStatsNtp.  # noqa: E501

        the time and date of the last accepted NTP reply, in ISO8601 format  # noqa: E501

        :return: The last_rx_accepted of this TimeSourceStatsNtp.  # noqa: E501
        :rtype: datetime
        """
        return self._last_rx_accepted

    @last_rx_accepted.setter
    def last_rx_accepted(self, last_rx_accepted):
        """Sets the last_rx_accepted of this TimeSourceStatsNtp.

        the time and date of the last accepted NTP reply, in ISO8601 format  # noqa: E501

        :param last_rx_accepted: The last_rx_accepted of this TimeSourceStatsNtp.  # noqa: E501
        :type: datetime
        """
        self._last_rx_accepted = last_rx_accepted

    @property
    def last_rx_ignored(self):
        """Gets the last_rx_ignored of this TimeSourceStatsNtp.  # noqa: E501

        The time and date of the last ignored NTP reply, in ISO8601 format  # noqa: E501

        :return: The last_rx_ignored of this TimeSourceStatsNtp.  # noqa: E501
        :rtype: datetime
        """
        return self._last_rx_ignored

    @last_rx_ignored.setter
    def last_rx_ignored(self, last_rx_ignored):
        """Sets the last_rx_ignored of this TimeSourceStatsNtp.

        The time and date of the last ignored NTP reply, in ISO8601 format  # noqa: E501

        :param last_rx_ignored: The last_rx_ignored of this TimeSourceStatsNtp.  # noqa: E501
        :type: datetime
        """
        self._last_rx_ignored = last_rx_ignored

    @property
    def poll_period(self):
        """Gets the poll_period of this TimeSourceStatsNtp.  # noqa: E501

        Current NTP server poll period, in seconds  # noqa: E501

        :return: The poll_period of this TimeSourceStatsNtp.  # noqa: E501
        :rtype: int
        """
        return self._poll_period

    @poll_period.setter
    def poll_period(self, poll_period):
        """Sets the poll_period of this TimeSourceStatsNtp.

        Current NTP server poll period, in seconds  # noqa: E501

        :param poll_period: The poll_period of this TimeSourceStatsNtp.  # noqa: E501
        :type: int
        """
        self._poll_period = poll_period

    @property
    def rx_ignored(self):
        """Gets the rx_ignored of this TimeSourceStatsNtp.  # noqa: E501

        Received packets that were ignored due to an invalid origin timestamp or stratum, e.g. a Kiss-o'-Death packet   # noqa: E501

        :return: The rx_ignored of this TimeSourceStatsNtp.  # noqa: E501
        :rtype: int
        """
        return self._rx_ignored

    @rx_ignored.setter
    def rx_ignored(self, rx_ignored):
        """Sets the rx_ignored of this TimeSourceStatsNtp.

        Received packets that were ignored due to an invalid origin timestamp or stratum, e.g. a Kiss-o'-Death packet   # noqa: E501

        :param rx_ignored: The rx_ignored of this TimeSourceStatsNtp.  # noqa: E501
        :type: int
        """
        self._rx_ignored = rx_ignored

    @property
    def rx_packets(self):
        """Gets the rx_packets of this TimeSourceStatsNtp.  # noqa: E501

        Received packets  # noqa: E501

        :return: The rx_packets of this TimeSourceStatsNtp.  # noqa: E501
        :rtype: int
        """
        return self._rx_packets

    @rx_packets.setter
    def rx_packets(self, rx_packets):
        """Sets the rx_packets of this TimeSourceStatsNtp.

        Received packets  # noqa: E501

        :param rx_packets: The rx_packets of this TimeSourceStatsNtp.  # noqa: E501
        :type: int
        """
        self._rx_packets = rx_packets

    @property
    def tx_packets(self):
        """Gets the tx_packets of this TimeSourceStatsNtp.  # noqa: E501

        Transmitted packets  # noqa: E501

        :return: The tx_packets of this TimeSourceStatsNtp.  # noqa: E501
        :rtype: int
        """
        return self._tx_packets

    @tx_packets.setter
    def tx_packets(self, tx_packets):
        """Sets the tx_packets of this TimeSourceStatsNtp.

        Transmitted packets  # noqa: E501

        :param tx_packets: The tx_packets of this TimeSourceStatsNtp.  # noqa: E501
        :type: int
        """
        self._tx_packets = tx_packets

    @property
    def stratum(self):
        """Gets the stratum of this TimeSourceStatsNtp.  # noqa: E501

        Time source distance from a NTP reference clock, in network hops.   # noqa: E501

        :return: The stratum of this TimeSourceStatsNtp.  # noqa: E501
        :rtype: int
        """
        return self._stratum

    @stratum.setter
    def stratum(self, stratum):
        """Sets the stratum of this TimeSourceStatsNtp.

        Time source distance from a NTP reference clock, in network hops.   # noqa: E501

        :param stratum: The stratum of this TimeSourceStatsNtp.  # noqa: E501
        :type: int
        """
        self._stratum = stratum

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
        if issubclass(TimeSourceStatsNtp, dict):
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
        if not isinstance(other, TimeSourceStatsNtp):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
