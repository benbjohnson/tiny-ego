package inmem

import (
	"sync"

	"github.com/benbjohnson/tiny-ego"
)

// DB represents a container for all data in the database.
type DB struct {
	sync.RWMutex

	seqs    map[string]int
	widgets map[int]*tinyego.Widget
}

// NewDB returns a new instance of DB.
func NewDB() *DB {
	return &DB{
		seqs:    make(map[string]int),
		widgets: make(map[int]*tinyego.Widget),
	}
}

func (db *DB) nextSeq(name string) int {
	db.seqs[name]++
	return db.seqs[name]
}
