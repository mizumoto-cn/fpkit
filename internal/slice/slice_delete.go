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

import "github.com/mizumoto-cn/fpkit/internal/err"

func Delete[T any](src []T, index int) ([]T, T, error) {
	l := len(src)
	if index < 0 || index >= l {
		var zero T
		return nil, zero, err.NewIndexOutOfRangeError(index, l)
	}
	deleted := src[index]
	copy(src[index:], src[index+1:])
	return src[:l-1], deleted, nil
}
