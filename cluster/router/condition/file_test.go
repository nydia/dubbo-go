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

package condition

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestLoadYmlConfig(t *testing.T) {
	router, e := NewFileConditionRouter([]byte(`priority: 1
force: true
conditions :
  - "a => b"
  - "c => d"`))
	assert.Nil(t, e)
	assert.NotNil(t, router)
	assert.Equal(t, router.routerRule.Priority, 1)
	assert.Equal(t, router.routerRule.Force, true)
	assert.Equal(t, len(router.routerRule.Conditions), 2)
}

func TestParseCondition(t *testing.T) {
	s := make([]string, 2)
	s = append(s, "a => b")
	s = append(s, "c => d")
	condition := parseCondition(s)
	assert.Equal(t, "a & c => b & d", condition)
}
