// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file implements a simple system for testing the HCL converter.
package tfgen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HclConversion(t *testing.T) {
	g, err := NewGenerator(GeneratorOptions{
		Version:         "version",
		Language:        Language("go"),
		Debug:           false,
		SkipDocs:        false,
		SkipExamples:    false,
		CoverageTracker: newCoverageTracker("Provider", "Version"),
	})

	assert.NoError(t, err, "Failed to create generator")

	// Incorrect HCL
	hcl := "data \"aws_ec2_local_gateway_route_table\" \"foo\" {}\n\noutput \"foo\" {\nvalue = data.aws_ec2_local_gateway_route_table.foo.ids\n}"

	// Working HCL
	// hcl := "data \"aws_directory_service_directory\" \"example\" {\n  directory_id = aws_directory_service_directory.main.id\n}"

	name := "name"

	g.coverageTracker.foundExample(name, hcl)
	codeBlock, stderr, err := g.convertHCL(hcl, name)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(codeBlock)
	fmt.Println(stderr)
	assert.NoError(t, err, "Failed to convert")
}
