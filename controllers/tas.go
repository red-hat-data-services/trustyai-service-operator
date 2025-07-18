/*
Copyright 2024.

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

package controllers

import (
	"github.com/trustyai-explainability/trustyai-service-operator/controllers/tas"
)

func init() {
	// Register just the v1alpha1 controller
	// Kubernetes API server will automatically convert  v1 (storage version) to v1alpha1 for controller to watch
	registerService(tas.ServiceName, tas.ControllerSetUp)
}
