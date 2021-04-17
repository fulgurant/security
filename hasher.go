package hasher

type Hasher interface {
	Hash(s string) string
}
