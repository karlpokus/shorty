package dbmock

type StoreMock map[string]string

func (sm StoreMock) List() map[string]string {
  return sm
}

func (sm StoreMock) Add(url string) (string, error) {
  return "abc123", nil
}

func (sm StoreMock) Find(key string) (string, error) {
  return "www.google.com", nil
}
