# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Connection(pulumi.CustomResource):
    """
    Provides a Connection of Direct Connect.
    """
    def __init__(__self__, __name__, __opts__=None, bandwidth=None, location=None, name=None, tags=None):
        """Create a Connection resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not bandwidth:
            raise TypeError('Missing required property bandwidth')
        elif not isinstance(bandwidth, basestring):
            raise TypeError('Expected property bandwidth to be a basestring')
        __self__.bandwidth = bandwidth
        """
        The bandwidth of the connection. Available values: 1Gbps, 10Gbps. Case sensitive.
        """
        __props__['bandwidth'] = bandwidth

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        The AWS Direct Connect location where the connection is located. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
        """
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the connection.
        """
        __props__['name'] = name

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        __props__['tags'] = tags

        __self__.arn = pulumi.runtime.UNKNOWN
        """
        The ARN of the connection.
        """

        super(Connection, __self__).__init__(
            'aws:directconnect/connection:Connection',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'arn' in outs:
            self.arn = outs['arn']
        if 'bandwidth' in outs:
            self.bandwidth = outs['bandwidth']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'tags' in outs:
            self.tags = outs['tags']
