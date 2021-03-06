# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Key(pulumi.CustomResource):
    """
    Provides a KMS customer master key.
    """
    def __init__(__self__, __name__, __opts__=None, deletion_window_in_days=None, description=None, enable_key_rotation=None, is_enabled=None, key_usage=None, policy=None, tags=None):
        """Create a Key resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if deletion_window_in_days and not isinstance(deletion_window_in_days, int):
            raise TypeError('Expected property deletion_window_in_days to be a int')
        __self__.deletion_window_in_days = deletion_window_in_days
        """
        Duration in days after which the key is deleted
        after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
        """
        __props__['deletionWindowInDays'] = deletion_window_in_days

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        The description of the key as viewed in AWS console.
        """
        __props__['description'] = description

        if enable_key_rotation and not isinstance(enable_key_rotation, bool):
            raise TypeError('Expected property enable_key_rotation to be a bool')
        __self__.enable_key_rotation = enable_key_rotation
        """
        Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
        is enabled. Defaults to false.
        """
        __props__['enableKeyRotation'] = enable_key_rotation

        if is_enabled and not isinstance(is_enabled, bool):
            raise TypeError('Expected property is_enabled to be a bool')
        __self__.is_enabled = is_enabled
        """
        Specifies whether the key is enabled. Defaults to true.
        """
        __props__['isEnabled'] = is_enabled

        if key_usage and not isinstance(key_usage, basestring):
            raise TypeError('Expected property key_usage to be a basestring')
        __self__.key_usage = key_usage
        """
        Specifies the intended use of the key.
        Defaults to ENCRYPT_DECRYPT, and only symmetric encryption and decryption are supported.
        """
        __props__['keyUsage'] = key_usage

        if policy and not isinstance(policy, basestring):
            raise TypeError('Expected property policy to be a basestring')
        __self__.policy = policy
        """
        A valid policy JSON document.
        """
        __props__['policy'] = policy

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the object.
        """
        __props__['tags'] = tags

        __self__.arn = pulumi.runtime.UNKNOWN
        """
        The Amazon Resource Name (ARN) of the key.
        """
        __self__.key_id = pulumi.runtime.UNKNOWN
        """
        The globally unique identifier for the key.
        """

        super(Key, __self__).__init__(
            'aws:kms/key:Key',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'arn' in outs:
            self.arn = outs['arn']
        if 'deletionWindowInDays' in outs:
            self.deletion_window_in_days = outs['deletionWindowInDays']
        if 'description' in outs:
            self.description = outs['description']
        if 'enableKeyRotation' in outs:
            self.enable_key_rotation = outs['enableKeyRotation']
        if 'isEnabled' in outs:
            self.is_enabled = outs['isEnabled']
        if 'keyId' in outs:
            self.key_id = outs['keyId']
        if 'keyUsage' in outs:
            self.key_usage = outs['keyUsage']
        if 'policy' in outs:
            self.policy = outs['policy']
        if 'tags' in outs:
            self.tags = outs['tags']
