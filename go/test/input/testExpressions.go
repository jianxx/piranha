/*
Copyright (c) 2021 Uber Technologies, Inc.
Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file

except in compliance with the License. You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the

License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/
package testfiles

import "fmt"

func testExpressions(ge GoExamples) {
	if ge.flagMthds.treatedBehaviour(staleFlag) {
		fmt.Println("then-branch of `ge.flagMthds.treatedBehaviour(staleFlag)`")
	} else {
		fmt.Println("else-branch of `ge.flagMthds.treatedBehaviour(staleFlag)`")
	}

	//global feature is not in properties right now. So this should not get treated
	if globalFeature(staleFlag) {
		fmt.Println("global treated behaviour")
	} else {
		fmt.Println("global control behaviour")
	}

	if ge.flagMthds.controlBehaviour(staleFlag) {
		fmt.Println("then-branch of `ge.flagMthds.controlBehaviour(staleFlag)`")
	} else {
		fmt.Println("else-branch of `ge.flagMthds.controlBehaviour(staleFlag)`")
	}
	var x, y bool = false, false

	if ge.flagMthds.treatedBehaviour(staleFlag) || x {
		fmt.Println("then-branch of `ge.flagMthds.treatedBehaviour(staleFlag) || x`")
	} else {
		fmt.Println("else-branch of `ge.flagMthds.treatedBehaviour(staleFlag) || x`")
	}

	if ge.flagMthds.treatedBehaviour(staleFlag) && x {
		fmt.Println("then-branch of `ge.flagMthds.treatedBehaviour(staleFlag) && x`")
	} else {
		fmt.Println("else-branch of `ge.flagMthds.treatedBehaviour(staleFlag) && x`")
	}

	if ge.flagMthds.treatedBehaviour(staleFlag) && (x || y) {
		fmt.Println("then-branch of `ge.flagMthds.treatedBehaviour(staleFlag) && (x || y)`")
	} else {
		fmt.Println("else-branch of `ge.flagMthds.treatedBehaviour(staleFlag) && (x || y)`")
	}

	if ge.flagMthds.treatedBehaviour(staleFlag) && (x && y) {
		fmt.Println("then-branch of `ge.flagMthds.treatedBehaviour(staleFlag) && (x && y)`")
	} else {
		fmt.Println("else-branch of `ge.flagMthds.treatedBehaviour(staleFlag) && (x && y)`")
	}

	if ge.flagMthds.treatedBehaviour(staleFlag) && y == x {
		fmt.Println("then-branch of `ge.flagMthds.treatedBehaviour(staleFlag) && y == x`")
	} else {
		fmt.Println("else-branch of `ge.flagMthds.treatedBehaviour(staleFlag) && y == x`")
	}

	if ge.flagMthds.controlBehaviour(staleFlag) || y == x {
		fmt.Println("then-branch of `ge.flagMthds.controlBehaviour(staleFlag) || y == x`")
	} else {
		fmt.Println("else-branch of `ge.flagMthds.controlBehaviour(staleFlag) || y == x`")
	}

	if ge.flagMthds.controlBehaviour(staleFlag) && y && x {
		fmt.Println("then-branch of `ge.flagMthds.controlBehaviour(staleFlag) && y && x`")
	} else {
		fmt.Println("else-branch of `ge.flagMthds.controlBehaviour(staleFlag) && y && x`")
	}

	if ge.flagMthds.controlBehaviour(staleFlag) || y || x {
		fmt.Println("then-branch of `ge.flagMthds.controlBehaviour(staleFlag) || y || x`")
	} else {
		fmt.Println("else-branch of `ge.flagMthds.controlBehaviour(staleFlag) || y || x`")
	}

	y = ge.flagMthds.treatedBehaviour(staleFlag)
	y = ge.flagMthds.controlBehaviour(staleFlag) && x
	y = ge.flagMthds.treatedBehaviour(staleFlag) || x

	if y {
		fmt.Println("y cleaned, so then-branch of y") 
	} else {
		fmt.Println("y cleaned, so else-branch of y")
	}

	y = ge.flagMthds.treatedBehaviour(staleFlag) && y == x
	// This is done on purpose to check deep clean work
	y = ge.flagMthds.treatedBehaviour(staleFlag)

	if y {
		fmt.Println("y not cleaned, so then-branch of y")
	} else {
		fmt.Println("y not cleaned, so else-branch of y")
	}

}
