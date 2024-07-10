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
package queue

import "context"

// BlockingQueue is a generic interface for a blocking queue.
// Blocking queues are not always FIFO, it depends on the implementation.
type BlockingQueue[T any] interface {
	// Push adds an element to the the queue.
	// when cancelled or timeout, return context.Canceled or context.DeadlineExceeded
	// Shall always use errors.Is(err, context.Canceled) or errors.Is(err, context.DeadlineExceeded) to check the error
	Push(ctx context.Context, t T) error

	// TryPop removes and returns a element in the queue.
	// when cancelled or timeout, return context.Canceled or context.DeadlineExceeded
	// Shall always use errors.Is(err, context.Canceled) or errors.Is(err, context.DeadlineExceeded) to check the error
	TryPop(ctx context.Context) (T, error)
}
