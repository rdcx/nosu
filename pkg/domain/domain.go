package domain

import "github.com/boltdb/bolt"

type Domain struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func GetAll(db *bolt.DB) ([]Domain, error) {
	var domains []Domain = []Domain{}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("domains"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			domains = append(domains, Domain{
				Name:  string(k),
				Owner: string(v),
			})
		}

		return nil
	})

	return domains, nil
}

func Register(name, owner string, db *bolt.DB) (*Domain, error) {
	dom := &Domain{
		Name:  name,
		Owner: owner,
	}

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("domains"))

		err := b.Put([]byte(name), []byte(owner))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return dom, nil
}
