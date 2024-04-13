/*
Copyright 2023 The Kubernetes Authors.

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

package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {

	type testCase[T interface {
		session.Session
	}] struct {
		name    string
		config  AWSSessionConfig
		want    T
		wantErr assert.ErrorAssertionFunc
	}

	tests := []testCase[session.Session]{
		{
			name:    "Sample test case",
			config:  AWSSessionConfig{},
			want:    session.Session{},
			wantErr: assert.NoError,
		},
		// {
		// 	name:    "Sample error test case",
		// 	config:    AWSSessionConfig{},
		// 	want:    session.Session{},
		// 	wantErr: assert.Error,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSession(tt.config)
			if !tt.wantErr(t, err, fmt.Sprintf("NewSession(%v)", tt.config)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewSession(%v)", tt.config)
		})
	}
}
