# coding: utf-8

"""
    Bacalhau API

    This page is the reference of the Bacalhau REST API. Project docs are available at https://docs.bacalhau.org/. Find more information about Bacalhau at https://github.com/bacalhau-project/bacalhau.  # noqa: E501

    OpenAPI spec version: 0.3.18.post4
    Contact: team@bacalhau.org
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from bacalhau_apiclient.configuration import Configuration


class NodeInfo(object):
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
        'compute_node_info': 'ComputeNodeInfo',
        'labels': 'dict(str, str)',
        'node_type': 'NodeType',
        'peer_info': 'PeerAddrInfo'
    }

    attribute_map = {
        'compute_node_info': 'ComputeNodeInfo',
        'labels': 'Labels',
        'node_type': 'NodeType',
        'peer_info': 'PeerInfo'
    }

    def __init__(self, compute_node_info=None, labels=None, node_type=None, peer_info=None, _configuration=None):  # noqa: E501
        """NodeInfo - a model defined in Swagger"""  # noqa: E501
        if _configuration is None:
            _configuration = Configuration()
        self._configuration = _configuration

        self._compute_node_info = None
        self._labels = None
        self._node_type = None
        self._peer_info = None
        self.discriminator = None

        if compute_node_info is not None:
            self.compute_node_info = compute_node_info
        if labels is not None:
            self.labels = labels
        if node_type is not None:
            self.node_type = node_type
        if peer_info is not None:
            self.peer_info = peer_info

    @property
    def compute_node_info(self):
        """Gets the compute_node_info of this NodeInfo.  # noqa: E501


        :return: The compute_node_info of this NodeInfo.  # noqa: E501
        :rtype: ComputeNodeInfo
        """
        return self._compute_node_info

    @compute_node_info.setter
    def compute_node_info(self, compute_node_info):
        """Sets the compute_node_info of this NodeInfo.


        :param compute_node_info: The compute_node_info of this NodeInfo.  # noqa: E501
        :type: ComputeNodeInfo
        """

        self._compute_node_info = compute_node_info

    @property
    def labels(self):
        """Gets the labels of this NodeInfo.  # noqa: E501


        :return: The labels of this NodeInfo.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._labels

    @labels.setter
    def labels(self, labels):
        """Sets the labels of this NodeInfo.


        :param labels: The labels of this NodeInfo.  # noqa: E501
        :type: dict(str, str)
        """

        self._labels = labels

    @property
    def node_type(self):
        """Gets the node_type of this NodeInfo.  # noqa: E501


        :return: The node_type of this NodeInfo.  # noqa: E501
        :rtype: NodeType
        """
        return self._node_type

    @node_type.setter
    def node_type(self, node_type):
        """Sets the node_type of this NodeInfo.


        :param node_type: The node_type of this NodeInfo.  # noqa: E501
        :type: NodeType
        """

        self._node_type = node_type

    @property
    def peer_info(self):
        """Gets the peer_info of this NodeInfo.  # noqa: E501


        :return: The peer_info of this NodeInfo.  # noqa: E501
        :rtype: PeerAddrInfo
        """
        return self._peer_info

    @peer_info.setter
    def peer_info(self, peer_info):
        """Sets the peer_info of this NodeInfo.


        :param peer_info: The peer_info of this NodeInfo.  # noqa: E501
        :type: PeerAddrInfo
        """

        self._peer_info = peer_info

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
        if issubclass(NodeInfo, dict):
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
        if not isinstance(other, NodeInfo):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, NodeInfo):
            return True

        return self.to_dict() != other.to_dict()
