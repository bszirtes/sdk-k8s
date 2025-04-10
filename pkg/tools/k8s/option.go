// Copyright (c) 2024-2025 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8s

type options struct {
	qps   float32
	burst int
}

// Option is an option pattern for NSM kubeutils
type Option func(o *options)

// WithQPS - QPS indicates the maximum QPS to the master from this client.
// See: https://github.com/kubernetes/client-go/blob/v0.28.3/rest/config.go#L116
func WithQPS(q float32) Option {
	return func(o *options) {
		o.qps = q
	}
}

// WithBurst - sets maximum burst for throttle
// See: https://github.com/kubernetes/client-go/blob/v0.28.3/rest/config.go#L120
func WithBurst(b int) Option {
	return func(o *options) {
		o.burst = b
	}
}
