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
package err

import (
	"fmt"
	"time"
)

func NewIndexOutOfRangeError(index, length int) error {
	return fmt.Errorf("fpkit: index out of range: [%d] with length: %d", index, length)
}

func NewQueueFullError(cap, l int) error {
	return fmt.Errorf("fpkit: queue is full, capacity is %d, current length is %d", cap, l)
}

func NewQueueResizeError(cap, l int) error {
	return fmt.Errorf("fpkit: cannot resize queue with new capacity %d, current limitation is %d", cap, l)
}

func NewQueueCapacityError(cap int) error {
	return fmt.Errorf("fpkit: invalid queue capacity: %d", cap)
}

func NewTypeCastError(from any, to string) error {
	return fmt.Errorf("fpkit: cannot cast type %#v to %s", from, to)
}

func NewInvalidTimeIntervalError(interval time.Duration) error {
	return fmt.Errorf("fpkit: invalid time interval: [%v]", interval)
}
