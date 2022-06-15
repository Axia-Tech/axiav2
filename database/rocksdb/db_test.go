//go:build linux && amd64 && rocksdballowed
// +build linux,amd64,rocksdballowed

// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rocksdb

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/Axia-Tech/axiav2/database"
	"github.com/Axia-Tech/axiav2/utils/logging"
)

func TestInterface(t *testing.T) {
	for _, test := range database.Tests {
		folder := t.TempDir()
		db, err := New(folder, nil, logging.NoLog{}, "", prometheus.NewRegistry())
		if err != nil {
			t.Fatalf("rocksdb.New(%q, logging.NoLog{}) erred with %s", folder, err)
		}

		test(t, db)

		// The database may have been closed by the test, so we don't care if it
		// errors here.
		_ = db.Close()
	}
}

func BenchmarkInterface(b *testing.B) {
	for _, size := range database.BenchmarkSizes {
		keys, values := database.SetupBenchmark(b, size[0], size[1], size[2])
		for _, bench := range database.Benchmarks {
			folder := b.TempDir()

			db, err := New(folder, nil, logging.NoLog{}, "", prometheus.NewRegistry())
			if err != nil {
				b.Fatal(err)
			}

			bench(b, db, "rocksdb", keys, values)

			// The database may have been closed by the test, so we don't care if it
			// errors here.
			_ = db.Close()
		}
	}
}