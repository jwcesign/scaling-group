/*
Copyright 2023 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	// The set of controllers this controller process runs.
	"knative.dev/sample-controller/pkg/reconciler/addressableservice"
	"knative.dev/sample-controller/pkg/reconciler/simpledeployment"
	"knative.dev/sample-controller/vendor/knative.dev/pkg/injection"

	// This defines the shared main for injected controllers.
	"knative.dev/pkg/injection/sharedmain"
)

var ctors = []injection.ControllerConstructor{}

func main() {

	sharedmain.Main("scaling-group-controller",
		addressableservice.NewController,
		simpledeployment.NewController,
	)
}
