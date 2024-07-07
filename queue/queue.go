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

// Queue is a generic interface for a queue.
// Queues are not always FIFO, it depends on the implementation.
type Queue[T any] interface {
	// Push adds an element to the the queue.
	Push(T) error

	// Pop removes and returns a element in the queue.
	Pop() (T, error)
}

// TODO: reference of C++ 11 std::queue
// member type		definition									notes
// value_type		The first template parameter (T)			Type of the elements
// container_type	The second template parameter (Container)	Type of the underlying container
// reference		container_type::reference					usually, value_type&
// const_reference	container_type::const_reference				usually, const value_type&
// size_type		an unsigned integral type					usually, the same as size_t
// (constructor)	Construct queue (public member function)
// empty			Test whether container is empty (public member function)
// size				Return size (public member function)
// front			Access next element (public member function)
// back				Access last element (public member function)
// push				Insert element (public member function)
// emplace			Construct and insert element (public member function)
// pop				Remove next element (public member function)
// swap				Swap contents (public member function)
