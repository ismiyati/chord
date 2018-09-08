package chord

import (
	"bytes"
	"errors"
	"math/rand"
	"time"
)

var (
	ERR_NO_SUCCESSOR = errors.New("cannot find successor")
	ERR_NODE_EXISTS  = errors.New("node with id already exists")
)

func isEqual(a, b []byte) bool {
	return bytes.Compare(a, b) == 0
}

func isPowerOfTwo(num int) bool {
	return (num != 0) && ((num & (num - 1)) == 0)
}

func randStabilize(min, max time.Duration) time.Duration {
	r := rand.Float64()
	return time.Duration((r * float64(max-min)) + float64(min))
}

// check if key is between a and b, right inclusive
func betweenRightIncl(key, a, b []byte) bool {
	return between(key, a, b) || bytes.Equal(key, b)
}

// Checks if a key is STRICTLY between two ID's exclusively
func between(key, a, b []byte) bool {
	switch bytes.Compare(a, b) {
	case 1:
		return bytes.Compare(a, key) == -1 || bytes.Compare(b, key) >= 0
	case -1:
		return bytes.Compare(a, key) == -1 && bytes.Compare(b, key) >= 0
	case 0:
		return bytes.Compare(a, key) != 0
	}
	return false
}

func (n *Node) hashKey(key string) ([]byte, error) {
	h := n.cnf.Hash
	if _, err := h.Write([]byte(key)); err != nil {
		return nil, err
	}
	val := h.Sum(nil)
	h.Reset()
	return val, nil
}
