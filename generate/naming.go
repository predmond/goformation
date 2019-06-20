package main

import (
	"path"
	"strings"
)

// filename takes a resource or property name (e.g. AWS::CloudFront::Distribution.Restrictions)
// and returns an appropriate filename for the generated struct (e.g. aws-cloudfront-distribution_restrictions.go)
func filename(input string) string {

	// Convert to lowercase
	output := strings.ToLower(input)

	// Replace :: with -
	output = strings.Replace(output, "::", "-", -1)

	// Replace . with _
	output = strings.Replace(output, ".", "_", -1)

	// Suffix with .go
	output += ".go"

	return output

}

// structName takes a resource or property name (e.g. AWS::CloudFront::Distribution.Restrictions)
// and returns an appropriate struct name for the generated struct (e.g. AWSCloudfrontDistributionRestrictions)
func structName(input string) string {

	// Remove ::
	output := strings.Replace(input, "::", "", -1)

	// Remove .
	output = strings.Replace(output, ".", "_", -1)

	return output

}

type structInfo struct {
	PackagePath string
	PackageName string
	FileName    string
	BaseName    string
	Name        string
}

func parseStructInfo(input string) structInfo {
	packageParts := strings.Split(input, "::")
	name := packageParts[len(packageParts)-1]
	nameParts := strings.Split(name, ".")
	packageParts = packageParts[:len(packageParts)-1]

	var result structInfo
	result.PackagePath = strings.ToLower(path.Join(packageParts...))
	if len(packageParts) == 0 {
		result.PackageName = "cloudformation"
	} else {
		result.PackageName = strings.ToLower(packageParts[len(packageParts)-1])
	}
	result.BaseName = nameParts[0]
	result.Name = strings.Join(nameParts, "_")
	result.FileName = strings.ToLower(result.Name) + ".go"
	return result
}
