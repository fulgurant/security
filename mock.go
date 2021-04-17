package hasher

import "fmt"

type Mock struct {
	salt string
}

func NewMock(salt string) *Mock {
	return &Mock{salt: salt}
}

func (m *Mock) Hash(s string) string {
	return fmt.Sprintf("%s%s", m.salt, s)
}
