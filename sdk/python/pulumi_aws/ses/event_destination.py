# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class EventDestination(pulumi.CustomResource):
    """
    Provides an SES event destination
    """
    def __init__(__self__, __name__, __opts__=None, cloudwatch_destination=None, configuration_set_name=None, enabled=None, kinesis_destination=None, matching_types=None, name=None, sns_destination=None):
        """Create a EventDestination resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if cloudwatch_destination and not isinstance(cloudwatch_destination, dict):
            raise TypeError('Expected property cloudwatch_destination to be a dict')
        __self__.cloudwatch_destination = cloudwatch_destination
        """
        CloudWatch destination for the events
        """
        __props__['cloudwatchDestination'] = cloudwatch_destination

        if not configuration_set_name:
            raise TypeError('Missing required property configuration_set_name')
        elif not isinstance(configuration_set_name, basestring):
            raise TypeError('Expected property configuration_set_name to be a basestring')
        __self__.configuration_set_name = configuration_set_name
        """
        The name of the configuration set
        """
        __props__['configurationSetName'] = configuration_set_name

        if enabled and not isinstance(enabled, bool):
            raise TypeError('Expected property enabled to be a bool')
        __self__.enabled = enabled
        """
        If true, the event destination will be enabled
        """
        __props__['enabled'] = enabled

        if kinesis_destination and not isinstance(kinesis_destination, dict):
            raise TypeError('Expected property kinesis_destination to be a dict')
        __self__.kinesis_destination = kinesis_destination
        """
        Send the events to a kinesis firehose destination
        """
        __props__['kinesisDestination'] = kinesis_destination

        if not matching_types:
            raise TypeError('Missing required property matching_types')
        elif not isinstance(matching_types, list):
            raise TypeError('Expected property matching_types to be a list')
        __self__.matching_types = matching_types
        """
        A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
        """
        __props__['matchingTypes'] = matching_types

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the event destination
        """
        __props__['name'] = name

        if sns_destination and not isinstance(sns_destination, dict):
            raise TypeError('Expected property sns_destination to be a dict')
        __self__.sns_destination = sns_destination
        """
        Send the events to an SNS Topic destination
        """
        __props__['snsDestination'] = sns_destination

        super(EventDestination, __self__).__init__(
            'aws:ses/eventDestination:EventDestination',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'cloudwatchDestination' in outs:
            self.cloudwatch_destination = outs['cloudwatchDestination']
        if 'configurationSetName' in outs:
            self.configuration_set_name = outs['configurationSetName']
        if 'enabled' in outs:
            self.enabled = outs['enabled']
        if 'kinesisDestination' in outs:
            self.kinesis_destination = outs['kinesisDestination']
        if 'matchingTypes' in outs:
            self.matching_types = outs['matchingTypes']
        if 'name' in outs:
            self.name = outs['name']
        if 'snsDestination' in outs:
            self.sns_destination = outs['snsDestination']
