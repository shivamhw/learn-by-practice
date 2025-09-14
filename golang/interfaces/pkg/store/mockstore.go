package store

import "fmt"

type MockStore struct {
	id string
}

func NewMockStore(id string) *MockStore {
	return &MockStore{
		id: id,
	}
}


func (m *MockStore) Read(f string) (data []byte, err error) {
	fmt.Printf("%s: reading %s", m.id, f)
	return
}

func (m *MockStore) Write(data []byte, f string) (err error) {
	fmt.Printf("%s: writing %d bytes to %s", m.id, len(data), f)
	return
}

func (m *MockStore) Exists(f string) (ok bool, err error) {
	fmt.Printf("%s: checking  %s file exists or not ", m.id, f)
	return
}

func (m *MockStore) ChangeId(id string) {
	m.id = id
}

func (m *MockStore) GetId(){
	fmt.Println(m.id)
}

func (m *MockStore) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"id":"%s","type":"mock"}`, m.id)), nil
}