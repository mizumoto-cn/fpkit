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
package slice

import (
	"github.com/mizumoto-cn/fpkit/internal/err"
)

func Insert[T any](s []T, index int, value T) ([]T, error) {
	if index < 0 || index > len(s) {
		return nil, err.NewIndexOutOfRangeError(index, len(s))
	}
	var zeroValue T
	s = append(s, zeroValue)
	// `copy` operates the memory directly, so it is faster than a for loop
	copy(s[index+1:], s[index:])
	s[index] = value
	return s, nil
}
