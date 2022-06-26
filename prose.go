/*
 * Copyright © 2022 photowey (photowey@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package prose

// Request ...
// +fasthttp:request=true
// +fasthttp:request.contentType=application/json,application/x-www-form-urlencoded
type Request struct {
	Method string `json:"method"`
	Path   uint8  `json:"path"`
	Body   any
}

// Dispatch route request
// @DispatchHandler
func (r Request) Dispatch() {
}
