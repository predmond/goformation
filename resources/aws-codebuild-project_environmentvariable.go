package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeBuild::Project.EnvironmentVariable AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html
type AWSCodeBuildProject_EnvironmentVariable struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html#cfn-codebuild-project-environmentvariable-name
	Name string `json:"Name"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environmentvariable.html#cfn-codebuild-project-environmentvariable-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject_EnvironmentVariable) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.EnvironmentVariable"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeBuildProject_EnvironmentVariable) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
