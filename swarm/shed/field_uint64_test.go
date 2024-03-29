// Copyright 2018 The go-solidum Authors
// This file is part of the go-solidum library.
//
// The go-solidum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-solidum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-solidum library. If not, see <http://www.gnu.org/licenses/>.

package shed

import (
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

// TestUint64Field validates put and get operations
// of the Uint64Field.
func TestUint64Field(t *testing.T) {
	db, cleanupFunc := newTestDB(t)
	defer cleanupFunc()

	counter, err := db.NewUint64Field("counter")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("get empty", func(t *testing.T) {
		got, err := counter.Get()
		if err != nil {
			t.Fatal(err)
		}
		var want uint64
		if got != want {
			t.Errorf("got uint64 %v, want %v", got, want)
		}
	})

	t.Run("put", func(t *testing.T) {
		var want uint64 = 42
		err = counter.Put(want)
		if err != nil {
			t.Fatal(err)
		}
		got, err := counter.Get()
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Errorf("got uint64 %v, want %v", got, want)
		}

		t.Run("overwrite", func(t *testing.T) {
			var want uint64 = 84
			err = counter.Put(want)
			if err != nil {
				t.Fatal(err)
			}
			got, err := counter.Get()
			if err != nil {
				t.Fatal(err)
			}
			if got != want {
				t.Errorf("got uint64 %v, want %v", got, want)
			}
		})
	})

	t.Run("put in batch", func(t *testing.T) {
		batch := new(leveldb.Batch)
		var want uint64 = 42
		counter.PutInBatch(batch, want)
		err = db.WriteBatch(batch)
		if err != nil {
			t.Fatal(err)
		}
		got, err := counter.Get()
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Errorf("got uint64 %v, want %v", got, want)
		}

		t.Run("overwrite", func(t *testing.T) {
			batch := new(leveldb.Batch)
			var want uint64 = 84
			counter.PutInBatch(batch, want)
			err = db.WriteBatch(batch)
			if err != nil {
				t.Fatal(err)
			}
			got, err := counter.Get()
			if err != nil {
				t.Fatal(err)
			}
			if got != want {
				t.Errorf("got uint64 %v, want %v", got, want)
			}
		})
	})
}

// TestUint64Field_Inc validates Inc operation
// of the Uint64Field.
func TestUint64Field_Inc(t *testing.T) {
	db, cleanupFunc := newTestDB(t)
	defer cleanupFunc()

	counter, err := db.NewUint64Field("counter")
	if err != nil {
		t.Fatal(err)
	}

	var want uint64 = 1
	got, err := counter.Inc()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got uint64 %v, want %v", got, want)
	}

	want = 2
	got, err = counter.Inc()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got uint64 %v, want %v", got, want)
	}
}

// TestUint64Field_IncInBatch validates IncInBatch operation
// of the Uint64Field.
func TestUint64Field_IncInBatch(t *testing.T) {
	db, cleanupFunc := newTestDB(t)
	defer cleanupFunc()

	counter, err := db.NewUint64Field("counter")
	if err != nil {
		t.Fatal(err)
	}

	batch := new(leveldb.Batch)
	var want uint64 = 1
	got, err := counter.IncInBatch(batch)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got uint64 %v, want %v", got, want)
	}
	err = db.WriteBatch(batch)
	if err != nil {
		t.Fatal(err)
	}
	got, err = counter.Get()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got uint64 %v, want %v", got, want)
	}

	batch2 := new(leveldb.Batch)
	want = 2
	got, err = counter.IncInBatch(batch2)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got uint64 %v, want %v", got, want)
	}
	err = db.WriteBatch(batch2)
	if err != nil {
		t.Fatal(err)
	}
	got, err = counter.Get()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got uint64 %v, want %v", got, want)
	}
}
