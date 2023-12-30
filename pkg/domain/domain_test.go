package domain_test

import (
	"nosu/pkg/domain"
	"nosu/pkg/testutil"
	"os"
	"testing"

	"github.com/boltdb/bolt"
)

func containsDomain(domains []domain.Domain, domain domain.Domain) bool {
	for _, d := range domains {
		if d.Name == domain.Name && d.Owner == domain.Owner {
			return true
		}
	}

	return false
}
func TestGetAll(t *testing.T) {
	t.Run("returns all domains", func(t *testing.T) {
		db, dbName := testutil.NewBoltDB()
		defer os.Remove(dbName)

		testutil.CreateBucket(db, "domains")

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("domains"))

			b.Put([]byte("test"), []byte("test"))

			b.Put([]byte("test2"), []byte("test2"))

			return nil
		})

		domains, err := domain.GetAll(db)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(domains) != 2 {
			t.Fatalf("expected 2 domain, got %v", len(domains))
		}

		if !containsDomain(domains, domain.Domain{Name: "test", Owner: "test"}) {
			t.Fatalf("expected domain with name test and owner test")
		}

		if !containsDomain(domains, domain.Domain{Name: "test2", Owner: "test2"}) {
			t.Fatalf("expected domain with name test2 and owner test2")
		}
	})
}

func TestRegister(t *testing.T) {
	t.Run("registers a domain", func(t *testing.T) {
		db, dbName := testutil.NewBoltDB()
		defer os.Remove(dbName)

		testutil.CreateBucket(db, "domains")

		dom, err := domain.Register("test", "test", db)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if dom.Name != "test" {
			t.Fatalf("expected domain name to be test, got %v", dom.Name)
		}

		if dom.Owner != "test" {
			t.Fatalf("expected domain owner to be test, got %v", dom.Owner)
		}
	})
}
