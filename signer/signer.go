package signer

type Signer interface {
	Sign(content []byte) ([]byte, error)
	Verify(content []byte, signature []byte) (bool, error)
}
