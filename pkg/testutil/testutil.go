package testutil

import (
	"math/rand"

	"github.com/boltdb/bolt"
)

func RandomString() string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 10)

	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}

	return string(b)
}

func NewBoltDB() (*bolt.DB, string) {
	dbName := RandomString() + ".db"
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		panic(err)
	}

	return db, dbName
}

func CreateBucket(db *bolt.DB, name string) {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}

		return nil
	})
}
