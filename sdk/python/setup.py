# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call

class InstallPluginCommand(install):
    def run(self):
        install.run(self)
        check_call(['pulumi', 'plugin', 'install', 'resource', 'aws', 'v0.14.0-dev-1528512373-gdbeff2d-dirty'])

setup(name='pulumi_aws',
      version='0.14.0.dev1528512373+gdbeff2d.dirty',
      description='A Pulumi package for creating and managing Amazon Web Services (AWS) cloud resources.',
      cmdclass={
          'install': InstallPluginCommand,
      },
      keywords='pulumi aws',
      url='https://pulumi.io/aws',
      project_urls={
          'Repository': 'https://github.com/pulumi/pulumi-aws'
      },
      packages=find_packages(),
      install_requires=[
          'pulumi>=0.12.2,<0.13.0'
      ],
      zip_safe=False)
