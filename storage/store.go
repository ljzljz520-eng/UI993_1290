package storage

import (
	"encoding/json"
	bb "go.etcd.io/bbolt"
	"os"
	"sync"
)

var buckets = []byte("referees")
var qb = []byte("qualifications")
var tb = []byte("trainings")
var ab = []byte("audits")

type Store struct {
	db *bb.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, e := bb.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	e = db.Update(func(tx *bb.Tx) error {
		for _, b := range [][]byte{buckets, qb, tb, ab} {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return &Store{db: db}, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func put[T any](s *Store, b []byte, key string, v T) error {
	raw, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bb.Tx) error { return tx.Bucket(b).Put([]byte(key), raw) })
}
func get[T any](s *Store, b []byte, key string) (T, error) {
	var out T
	e := s.db.View(func(tx *bb.Tx) error {
		v := tx.Bucket(b).Get([]byte(key))
		if v == nil {
			return os.ErrNotExist
		}
		return json.Unmarshal(v, &out)
	})
	return out, e
}
func del(s *Store, b []byte, key string) error {
	return s.db.Update(func(tx *bb.Tx) error { return tx.Bucket(b).Delete([]byte(key)) })
}
func scan[T any](s *Store, b []byte) ([]T, error) {
	out := []T{}
	e := s.db.View(func(tx *bb.Tx) error {
		return tx.Bucket(b).ForEach(func(_, v []byte) error {
			var x T
			if e := json.Unmarshal(v, &x); e != nil {
				return e
			}
			out = append(out, x)
			return nil
		})
	})
	return out, e
}
