package inspec

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/chrisevett/stool/verifiers"
)

func LoadOutput(r io.Reader) (error, InspecOutputModel) {
	bytes, err := ioutil.ReadAll(r)
	var i InspecOutputModel
	err = json.Unmarshal(bytes, &i)
	return err, i
}

func InspecOutputToTestSuite(i *InspecOutputModel) (error, verifiers.TestSuite) {
	suite := verifiers.TestSuite{}
	// platform
	suite.Name = i.Profiles[0].Name
	suite.Platform = i.Platform.Name

	testcases := []verifiers.TestCase{}

	for _, control := range i.Controls {
		t := verifiers.TestCase{}
		t.Name = control.CodeDesc
		t.Message = control.Message
		if control.Status == "passed" {
			t.Result = true
		} else {
			t.Result = false
		}
		testcases = append(testcases, t)
	}

	suite.Tests = testcases

	// name from first profile because
	return nil, suite

}

func OutputToSuite(r io.Reader) (error, verifiers.TestSuite) {
	err, model := LoadOutput(r)
	if err != nil {
		fmt.Println("error loading json output from inspec test data")
		return err, verifiers.TestSuite{}
	}
	err, suite := InspecOutputToTestSuite(&model)
	if err != nil {
		fmt.Println("error converting inspec output")
		return err, verifiers.TestSuite{}
	}

	return err, suite

}
