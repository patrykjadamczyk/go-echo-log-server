package logdb

import (
	"sync"
)

type LogRequest struct {
	// Start time
	Start string
	// Request Information
	RequestInfo string
	// Request Data Form as JSON String
	RequestDataForm string
	// Request Data
	RequestData string
}

type DB struct {
	mu   sync.RWMutex
	data map[string]LogRequest
}

func Create() *DB {
	db := &DB{
		data: make(map[string]LogRequest),
	}
	return db
}

type Tx struct {
	db       *DB
	writable bool
}

func (tx *Tx) lock() {
	if tx.writable {
		tx.db.mu.Lock()
	} else {
		tx.db.mu.RLock()
	}
}

func (tx *Tx) unlock() {
	if tx.writable {
		tx.db.mu.Unlock()
	} else {
		tx.db.mu.RUnlock()
	}
}

func (tx *Tx) Set(key string, value LogRequest) {
	tx.db.data[key] = value
}

func (tx *Tx) Get(key string) LogRequest {
	return tx.db.data[key]
}

func (tx *Tx) GetAll() map[string]LogRequest {
	return tx.db.data
}

func (db *DB) View(fn func(tx *Tx) error) error {
	return db.managed(false, fn)
}

func (db *DB) Update(fn func(tx *Tx) error) error {
	return db.managed(true, fn)
}

func (db *DB) Begin(writable bool) *Tx {
	tx := &Tx{
		db:       db,
		writable: writable,
	}
	tx.lock()

	return tx
}

func (db *DB) managed(writable bool, fn func(tx *Tx) error) (err error) {
	var tx *Tx
	tx = db.Begin(writable)
	defer func() {
		if writable {
			tx.unlock()
		} else {
			tx.unlock()
		}
	}()

	_ = fn(tx)
	return
}

// Example
/*
func main() {

	db := Create()

	go db.Update(func(tx *Tx) error {
		tx.Set("mykey", "go")
		tx.Set("mykey2", "is")
		tx.Set("mykey3", "awesome")
		return nil
	})

	go db.View(func(tx *Tx) error {
		fmt.Println("value is")
		fmt.Println(tx.Get("mykey3"))
		return nil
	})

	time.Sleep(20000000000)
}
*/
