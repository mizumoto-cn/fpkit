/*
 * Copyright (c) 2024 Ruiyuan "mizumoto-cn" Xu
 *
 * This file is part of "github.com/mizumoto-cn/fpkit".
 *
 * Licensed under the Mizumoto General Public License v1.5 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://github.com/mizumoto-cn/fpkit/blob/main/LICENSE
 *     https://github.com/mizumoto-cn/fpkit/blob/main/licensing
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package functional

import "sync/atomic"

// AtomicBool is a boolean value that can be atomically set and read.
type AtomicBool struct {
	value int32
}

// Set sets the value of the atomic boolean.
func (ab *AtomicBool) Set(value bool) {
	var intVal int32
	if value {
		intVal = 1
	} else {
		intVal = 0
	}
	atomic.StoreInt32(&ab.value, intVal)
}

// Get gets the value of the atomic boolean.
func (ab *AtomicBool) Get() bool {
	return atomic.LoadInt32(&ab.value) == 1
}
