// Copyright 2014 Google Inc. All Rights Reserved.
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

package validate

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/heapster/sinks"
	"github.com/GoogleCloudPlatform/heapster/sources/api"
	"github.com/GoogleCloudPlatform/heapster/version"
)

const (
	ValidatePage = "/validate/"
)

func HandleRequest(w http.ResponseWriter, sources []api.Source, sink sinks.ExternalSinkManager) error {
	out := fmt.Sprintf("Heapster Version: %v\n\n", version.HeapsterVersion)
	for _, source := range sources {
		out += source.DebugInfo()
	}
	out += sink.DebugInfo()
	_, err := w.Write([]byte(out))
	return err
}
