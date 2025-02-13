/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mock

import (
	"context"
	"fmt"
	"net/http"
)

import (
	"github.com/apache/dubbo-go-pixiu/pkg/common/extension/filter"
	contexthttp "github.com/apache/dubbo-go-pixiu/pkg/context/http"
)

// GetMockHTTPContext mock context for test.
func GetMockHTTPContext(r *http.Request, fc ...filter.HttpFilter) *contexthttp.HttpContext {
	result := &contexthttp.HttpContext{
		Index:   -1,
		Request: r,
	}

	w := mockWriter{header: map[string][]string{}}
	result.Writer = &w
	result.Reset()
	result.Ctx = context.Background()
	for i := range fc {
		result.Filters = append(result.Filters, fc[i].Handle)
	}

	return result
}

type mockWriter struct {
	header http.Header
}

func (w *mockWriter) Header() http.Header {
	return w.header
}

func (w *mockWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	return -1, nil
}

func (w *mockWriter) WriteHeader(statusCode int) {
	fmt.Println(statusCode)
}
