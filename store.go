package shorty

import "fmt"

const keyLength = 5

type Store interface {
	// List returns a map of key:long url
	List() map[string]string
	// Add accepts a long a url, saves it and returns its key
	Add(string) (string, error)
	// Find accepts a key and returns a long url
	Find(string) (string, error)
}

type database map[string]string

func (db database) List() map[string]string {
	return db
}

func (db database) Add(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("Empty url string")
	}
	for k, v := range db {
		if v == url {
			return k, nil
		}
	}
	bytes := randSeq(keyLength)
	key := string(bytes)
	db[key] = url
	return key, nil
}

func (db database) Find(key string) (string, error) {
	url, ok := db[key]
	if ok {
		return url, nil
	}
	return "", fmt.Errorf("%s does not exist", key)
}
