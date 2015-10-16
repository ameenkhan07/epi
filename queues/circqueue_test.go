// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package queues

import "testing"

func TestIntArrayQueue(t *testing.T) {
	ifaceTests := []queueTest{
		{0, nil},
		{1, nil},
		{2, nil},
		{3, nil},
		{4, nil},
		{5, nil},
		{6, nil},
		{7, nil},
		{8, nil},
		{9, nil},
		{minInt, nil},
		{maxInt, nil},
		{"x", ErrType},
	}
	if err := testQueueInterface(t, NewIntArrayQueue(0), ifaceTests); err != nil {
		t.Error(err)
	}
}

func BenchmarkIntArrayQueue(b *testing.B) {
	b.StopTimer()
	q := NewIntArrayQueue(0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
		if i%64 == 0 {
			for i := 0; i < 16; i++ {
				q.Dequeue()
			}
		}
	}
}