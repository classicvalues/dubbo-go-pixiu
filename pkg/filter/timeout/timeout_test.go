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

package timeout

import (
	"testing"
	"time"
)

import (
	"github.com/apache/dubbo-go-pixiu/pkg/common/extension/filter"
	"github.com/apache/dubbo-go-pixiu/pkg/context/mock"
	"github.com/apache/dubbo-go-pixiu/pkg/filter/recovery"
)

func timeoutFilterFunc(wait time.Duration) filter.HttpFilter {
	config := &Config{Timeout: wait}
	t := &Filter{cfg: config}
	_ = t.Apply()
	return t
}

func TestPanic(t *testing.T) {
	c := mock.GetMockHTTPContext(nil, timeoutFilterFunc(0), recovery.GetMock(), timeoutFilterFunc(time.Millisecond*100))
	c.Next()
	// print
	// 500
	// "timeout filter test panic"
}

func TestTimeout(t *testing.T) {
	c := mock.GetMockHTTPContext(nil, timeoutFilterFunc(0), timeoutFilterFunc(time.Second*3))
	c.Next()
	// print
	// 503
	// {"code":"S005","message":"http: Handler timeout"}
}

func TestNormal(t *testing.T) {
	c := mock.GetMockHTTPContext(nil, timeoutFilterFunc(time.Millisecond*200))
	c.Next()
}
