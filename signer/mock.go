package signer

// Mock is used to test signing and verifying content
type Mock struct {
	salt        byte
	SignError   error
	VerifyError error
}

// NewMock instantiates a new Mock signer
func NewMock(salt byte) *Mock {
	return &Mock{salt: salt}
}

// Sign signs content
func (m *Mock) Sign(content []byte) ([]byte, error) {
	if m.SignError != nil {
		return nil, m.SignError
	}
	checksum := byte(0)
	for i, v := range content {
		checksum += m.salt + byte(i) + v
	}
	return []byte{checksum}, nil
}

// Verify verifies if a signature is correct
func (m *Mock) Verify(content []byte, signature []byte) (bool, error) {
	if m.SignError != nil {
		return false, m.SignError
	}
	checksum := byte(0)
	for i, v := range content {
		checksum += m.salt + byte(i) + v
	}

	if len(signature) != 1 {
		return false, nil
	}

	result := signature[0] == checksum
	return result, nil
}
