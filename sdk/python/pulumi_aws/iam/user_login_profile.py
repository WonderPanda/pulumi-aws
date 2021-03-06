# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class UserLoginProfile(pulumi.CustomResource):
    """
    Provides one-time creation of a IAM user login profile, and uses PGP to
    encrypt the password for safe transport to the user. PGP keys can be
    obtained from Keybase.
    """
    def __init__(__self__, __name__, __opts__=None, password_length=None, password_reset_required=None, pgp_key=None, user=None):
        """Create a UserLoginProfile resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if password_length and not isinstance(password_length, int):
            raise TypeError('Expected property password_length to be a int')
        __self__.password_length = password_length
        """
        The length of the generated
        password.
        """
        __props__['passwordLength'] = password_length

        if password_reset_required and not isinstance(password_reset_required, bool):
            raise TypeError('Expected property password_reset_required to be a bool')
        __self__.password_reset_required = password_reset_required
        """
        Whether the
        user should be forced to reset the generated password on first login.
        """
        __props__['passwordResetRequired'] = password_reset_required

        if not pgp_key:
            raise TypeError('Missing required property pgp_key')
        elif not isinstance(pgp_key, basestring):
            raise TypeError('Expected property pgp_key to be a basestring')
        __self__.pgp_key = pgp_key
        """
        Either a base-64 encoded PGP public key, or a
        keybase username in the form `keybase:username`.
        """
        __props__['pgpKey'] = pgp_key

        if not user:
            raise TypeError('Missing required property user')
        elif not isinstance(user, basestring):
            raise TypeError('Expected property user to be a basestring')
        __self__.user = user
        """
        The IAM user's name.
        """
        __props__['user'] = user

        __self__.encrypted_password = pulumi.runtime.UNKNOWN
        """
        The encrypted password, base64 encoded.
        """
        __self__.key_fingerprint = pulumi.runtime.UNKNOWN
        """
        The fingerprint of the PGP key used to encrypt
        the password
        """

        super(UserLoginProfile, __self__).__init__(
            'aws:iam/userLoginProfile:UserLoginProfile',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'encryptedPassword' in outs:
            self.encrypted_password = outs['encryptedPassword']
        if 'keyFingerprint' in outs:
            self.key_fingerprint = outs['keyFingerprint']
        if 'passwordLength' in outs:
            self.password_length = outs['passwordLength']
        if 'passwordResetRequired' in outs:
            self.password_reset_required = outs['passwordResetRequired']
        if 'pgpKey' in outs:
            self.pgp_key = outs['pgpKey']
        if 'user' in outs:
            self.user = outs['user']
