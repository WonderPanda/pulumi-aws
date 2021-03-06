# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class MetricAlarm(pulumi.CustomResource):
    """
    Provides a CloudWatch Metric Alarm resource.
    """
    def __init__(__self__, __name__, __opts__=None, actions_enabled=None, alarm_actions=None, alarm_description=None, name=None, comparison_operator=None, datapoints_to_alarm=None, dimensions=None, evaluate_low_sample_count_percentiles=None, evaluation_periods=None, extended_statistic=None, insufficient_data_actions=None, metric_name=None, namespace=None, ok_actions=None, period=None, statistic=None, threshold=None, treat_missing_data=None, unit=None):
        """Create a MetricAlarm resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if actions_enabled and not isinstance(actions_enabled, bool):
            raise TypeError('Expected property actions_enabled to be a bool')
        __self__.actions_enabled = actions_enabled
        """
        Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
        """
        __props__['actionsEnabled'] = actions_enabled

        if alarm_actions and not isinstance(alarm_actions, list):
            raise TypeError('Expected property alarm_actions to be a list')
        __self__.alarm_actions = alarm_actions
        """
        The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Number (ARN).
        """
        __props__['alarmActions'] = alarm_actions

        if alarm_description and not isinstance(alarm_description, basestring):
            raise TypeError('Expected property alarm_description to be a basestring')
        __self__.alarm_description = alarm_description
        """
        The description for the alarm.
        """
        __props__['alarmDescription'] = alarm_description

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The descriptive name for the alarm. This name must be unique within the user's AWS account
        """
        __props__['name'] = name

        if not comparison_operator:
            raise TypeError('Missing required property comparison_operator')
        elif not isinstance(comparison_operator, basestring):
            raise TypeError('Expected property comparison_operator to be a basestring')
        __self__.comparison_operator = comparison_operator
        """
        The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`.
        """
        __props__['comparisonOperator'] = comparison_operator

        if datapoints_to_alarm and not isinstance(datapoints_to_alarm, int):
            raise TypeError('Expected property datapoints_to_alarm to be a int')
        __self__.datapoints_to_alarm = datapoints_to_alarm
        """
        The number of datapoints that must be breaching to trigger the alarm.
        """
        __props__['datapointsToAlarm'] = datapoints_to_alarm

        if dimensions and not isinstance(dimensions, dict):
            raise TypeError('Expected property dimensions to be a dict')
        __self__.dimensions = dimensions
        """
        The dimensions for the alarm's associated metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        """
        __props__['dimensions'] = dimensions

        if evaluate_low_sample_count_percentiles and not isinstance(evaluate_low_sample_count_percentiles, basestring):
            raise TypeError('Expected property evaluate_low_sample_count_percentiles to be a basestring')
        __self__.evaluate_low_sample_count_percentiles = evaluate_low_sample_count_percentiles
        """
        Used only for alarms
        based on percentiles. If you specify `ignore`, the alarm state will not
        change during periods with too few data points to be statistically significant.
        If you specify `evaluate` or omit this parameter, the alarm will always be
        evaluated and possibly change state no matter how many data points are available.
        The following values are supported: `ignore`, and `evaluate`.
        """
        __props__['evaluateLowSampleCountPercentiles'] = evaluate_low_sample_count_percentiles

        if not evaluation_periods:
            raise TypeError('Missing required property evaluation_periods')
        elif not isinstance(evaluation_periods, int):
            raise TypeError('Expected property evaluation_periods to be a int')
        __self__.evaluation_periods = evaluation_periods
        """
        The number of periods over which data is compared to the specified threshold.
        """
        __props__['evaluationPeriods'] = evaluation_periods

        if extended_statistic and not isinstance(extended_statistic, basestring):
            raise TypeError('Expected property extended_statistic to be a basestring')
        __self__.extended_statistic = extended_statistic
        """
        The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
        """
        __props__['extendedStatistic'] = extended_statistic

        if insufficient_data_actions and not isinstance(insufficient_data_actions, list):
            raise TypeError('Expected property insufficient_data_actions to be a list')
        __self__.insufficient_data_actions = insufficient_data_actions
        """
        The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Number (ARN).
        """
        __props__['insufficientDataActions'] = insufficient_data_actions

        if not metric_name:
            raise TypeError('Missing required property metric_name')
        elif not isinstance(metric_name, basestring):
            raise TypeError('Expected property metric_name to be a basestring')
        __self__.metric_name = metric_name
        """
        The name for the alarm's associated metric.
        See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        """
        __props__['metricName'] = metric_name

        if not namespace:
            raise TypeError('Missing required property namespace')
        elif not isinstance(namespace, basestring):
            raise TypeError('Expected property namespace to be a basestring')
        __self__.namespace = namespace
        """
        The namespace for the alarm's associated metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
        See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        """
        __props__['namespace'] = namespace

        if ok_actions and not isinstance(ok_actions, list):
            raise TypeError('Expected property ok_actions to be a list')
        __self__.ok_actions = ok_actions
        """
        The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Number (ARN).
        """
        __props__['okActions'] = ok_actions

        if not period:
            raise TypeError('Missing required property period')
        elif not isinstance(period, int):
            raise TypeError('Expected property period to be a int')
        __self__.period = period
        """
        The period in seconds over which the specified `statistic` is applied.
        """
        __props__['period'] = period

        if statistic and not isinstance(statistic, basestring):
            raise TypeError('Expected property statistic to be a basestring')
        __self__.statistic = statistic
        """
        The statistic to apply to the alarm's associated metric.
        Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
        """
        __props__['statistic'] = statistic

        if not threshold:
            raise TypeError('Missing required property threshold')
        elif not isinstance(threshold, float):
            raise TypeError('Expected property threshold to be a float')
        __self__.threshold = threshold
        """
        The value against which the specified statistic is compared.
        """
        __props__['threshold'] = threshold

        if treat_missing_data and not isinstance(treat_missing_data, basestring):
            raise TypeError('Expected property treat_missing_data to be a basestring')
        __self__.treat_missing_data = treat_missing_data
        """
        Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
        """
        __props__['treatMissingData'] = treat_missing_data

        if unit and not isinstance(unit, basestring):
            raise TypeError('Expected property unit to be a basestring')
        __self__.unit = unit
        """
        The unit for the alarm's associated metric.
        """
        __props__['unit'] = unit

        __self__.arn = pulumi.runtime.UNKNOWN
        """
        The ARN of the cloudwatch metric alarm.
        """

        super(MetricAlarm, __self__).__init__(
            'aws:cloudwatch/metricAlarm:MetricAlarm',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'actionsEnabled' in outs:
            self.actions_enabled = outs['actionsEnabled']
        if 'alarmActions' in outs:
            self.alarm_actions = outs['alarmActions']
        if 'alarmDescription' in outs:
            self.alarm_description = outs['alarmDescription']
        if 'name' in outs:
            self.name = outs['name']
        if 'arn' in outs:
            self.arn = outs['arn']
        if 'comparisonOperator' in outs:
            self.comparison_operator = outs['comparisonOperator']
        if 'datapointsToAlarm' in outs:
            self.datapoints_to_alarm = outs['datapointsToAlarm']
        if 'dimensions' in outs:
            self.dimensions = outs['dimensions']
        if 'evaluateLowSampleCountPercentiles' in outs:
            self.evaluate_low_sample_count_percentiles = outs['evaluateLowSampleCountPercentiles']
        if 'evaluationPeriods' in outs:
            self.evaluation_periods = outs['evaluationPeriods']
        if 'extendedStatistic' in outs:
            self.extended_statistic = outs['extendedStatistic']
        if 'insufficientDataActions' in outs:
            self.insufficient_data_actions = outs['insufficientDataActions']
        if 'metricName' in outs:
            self.metric_name = outs['metricName']
        if 'namespace' in outs:
            self.namespace = outs['namespace']
        if 'okActions' in outs:
            self.ok_actions = outs['okActions']
        if 'period' in outs:
            self.period = outs['period']
        if 'statistic' in outs:
            self.statistic = outs['statistic']
        if 'threshold' in outs:
            self.threshold = outs['threshold']
        if 'treatMissingData' in outs:
            self.treat_missing_data = outs['treatMissingData']
        if 'unit' in outs:
            self.unit = outs['unit']
