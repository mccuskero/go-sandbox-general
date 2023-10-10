package bloom

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"hash"
	"errors"
)


type SimpleBloom struct {
	Bitmap []bool
	Sha1Hasher hash.Hash
}

func NewSimpleBloom() *SimpleBloom {
	simpleBloom := SimpleBloom{}

	bitmap := make([]bool, 100)
  
	simpleBloom.Sha1Hasher = sha1.New()
	simpleBloom.Bitmap = bitmap

	return &simpleBloom
}

func (sb *SimpleBloom) Set(uuid string) error {
	hashPos, err := sb.getHashPosition(uuid)

	if err != nil {
		return errors.New("error setting uuid in bloom filter")
	}

	sb.Bitmap[hashPos] = true

	return nil
}

func (sb *SimpleBloom) Get(uuid string) (bool, error) {
	hashPos, err := sb.getHashPosition(uuid)

	if err != nil {
		return false, errors.New("error getting value from bitmap")
	}

	return sb.Bitmap[hashPos], nil
}

func (sb *SimpleBloom) getHashPosition(uuid string) (int, error) {
	hash, err := sb.createHash(uuid)

	if err != nil {
		return 0, errors.New("error getting hash position")
	}

	if hash < 0 {
		hash = -hash
	}

	hashPos := hash % len(sb.Bitmap)

	return hashPos, nil
}

func (sb *SimpleBloom) createHash(uuid string) (int, error) {
	bits := sb.Sha1Hasher.Sum([]byte(uuid))
	buffer := bytes.NewBuffer(bits)
	result, err := binary.ReadVarint(buffer)

	if err != nil {
		return 0, errors.New("error creating hash")
	}

	return int(result), nil
}