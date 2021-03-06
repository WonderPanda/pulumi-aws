# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Preset(pulumi.CustomResource):
    """
    Provides an Elastic Transcoder preset resource.
    """
    def __init__(__self__, __name__, __opts__=None, audio=None, audio_codec_options=None, container=None, description=None, name=None, thumbnails=None, type=None, video=None, video_codec_options=None, video_watermarks=None):
        """Create a Preset resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if audio and not isinstance(audio, dict):
            raise TypeError('Expected property audio to be a dict')
        __self__.audio = audio
        """
        Audio parameters object (documented below).
        """
        __props__['audio'] = audio

        if audio_codec_options and not isinstance(audio_codec_options, dict):
            raise TypeError('Expected property audio_codec_options to be a dict')
        __self__.audio_codec_options = audio_codec_options
        """
        Codec options for the audio parameters (documented below)
        """
        __props__['audioCodecOptions'] = audio_codec_options

        if not container:
            raise TypeError('Missing required property container')
        elif not isinstance(container, basestring):
            raise TypeError('Expected property container to be a basestring')
        __self__.container = container
        """
        The container type for the output file. Valid values are `flac`, `flv`, `fmp4`, `gif`, `mp3`, `mp4`, `mpg`, `mxf`, `oga`, `ogg`, `ts`, and `webm`.
        """
        __props__['container'] = container

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        A description of the preset (maximum 255 characters)
        """
        __props__['description'] = description

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the preset. (maximum 40 characters)
        """
        __props__['name'] = name

        if thumbnails and not isinstance(thumbnails, dict):
            raise TypeError('Expected property thumbnails to be a dict')
        __self__.thumbnails = thumbnails
        """
        Thumbnail parameters object (documented below)
        """
        __props__['thumbnails'] = thumbnails

        if type and not isinstance(type, basestring):
            raise TypeError('Expected property type to be a basestring')
        __self__.type = type
        __props__['type'] = type

        if video and not isinstance(video, dict):
            raise TypeError('Expected property video to be a dict')
        __self__.video = video
        """
        Video parameters object (documented below)
        """
        __props__['video'] = video

        if video_codec_options and not isinstance(video_codec_options, dict):
            raise TypeError('Expected property video_codec_options to be a dict')
        __self__.video_codec_options = video_codec_options
        __props__['videoCodecOptions'] = video_codec_options

        if video_watermarks and not isinstance(video_watermarks, list):
            raise TypeError('Expected property video_watermarks to be a list')
        __self__.video_watermarks = video_watermarks
        """
        Watermark parameters for the video parameters (documented below)
        * `video_codec_options` (Optional, Forces new resource) Codec options for the video parameters
        """
        __props__['videoWatermarks'] = video_watermarks

        __self__.arn = pulumi.runtime.UNKNOWN

        super(Preset, __self__).__init__(
            'aws:elastictranscoder/preset:Preset',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'arn' in outs:
            self.arn = outs['arn']
        if 'audio' in outs:
            self.audio = outs['audio']
        if 'audioCodecOptions' in outs:
            self.audio_codec_options = outs['audioCodecOptions']
        if 'container' in outs:
            self.container = outs['container']
        if 'description' in outs:
            self.description = outs['description']
        if 'name' in outs:
            self.name = outs['name']
        if 'thumbnails' in outs:
            self.thumbnails = outs['thumbnails']
        if 'type' in outs:
            self.type = outs['type']
        if 'video' in outs:
            self.video = outs['video']
        if 'videoCodecOptions' in outs:
            self.video_codec_options = outs['videoCodecOptions']
        if 'videoWatermarks' in outs:
            self.video_watermarks = outs['videoWatermarks']
