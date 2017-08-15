package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApplicationAutoScaling::ScalingPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html
type AWSApplicationAutoScalingScalingPolicy struct {

	// PolicyName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-policyname
	PolicyName string `json:"PolicyName"`

	// PolicyType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-policytype
	PolicyType string `json:"PolicyType"`

	// ResourceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-resourceid
	ResourceId string `json:"ResourceId"`

	// ScalableDimension AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-scalabledimension
	ScalableDimension string `json:"ScalableDimension"`

	// ScalingTargetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-scalingtargetid
	ScalingTargetId string `json:"ScalingTargetId"`

	// ServiceNamespace AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-servicenamespace
	ServiceNamespace string `json:"ServiceNamespace"`

	// StepScalingPolicyConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration
	StepScalingPolicyConfiguration AWSApplicationAutoScalingScalingPolicy_StepScalingPolicyConfiguration `json:"StepScalingPolicyConfiguration"`

	// TargetTrackingScalingPolicyConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#cfn-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration
	TargetTrackingScalingPolicyConfiguration AWSApplicationAutoScalingScalingPolicy_TargetTrackingScalingPolicyConfiguration `json:"TargetTrackingScalingPolicyConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicy) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalingPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApplicationAutoScalingScalingPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApplicationAutoScalingScalingPolicyResources retrieves all AWSApplicationAutoScalingScalingPolicy items from a CloudFormation template
func GetAllAWSApplicationAutoScalingScalingPolicyResources(template *Template) map[string]*AWSApplicationAutoScalingScalingPolicy {

	results := map[string]*AWSApplicationAutoScalingScalingPolicy{}
	for name, resource := range template.Resources {
		result := &AWSApplicationAutoScalingScalingPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApplicationAutoScalingScalingPolicyWithName retrieves all AWSApplicationAutoScalingScalingPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSApplicationAutoScalingScalingPolicyWithName(name string, template *Template) (*AWSApplicationAutoScalingScalingPolicy, error) {

	result := &AWSApplicationAutoScalingScalingPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApplicationAutoScalingScalingPolicy{}, errors.New("resource not found")

}
