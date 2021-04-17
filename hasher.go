package simplehash

type Hasher interface {
	Hash(s string) string
}
