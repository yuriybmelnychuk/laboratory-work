package hash

import (
	"encoding/binary"
	"math/bits"
)

const (
	// in bytes (512 bits)
	chunkSize = 64

	// in bytes (32 bits)
	wordSize = 4

	_h0 uint32 = 0x67452301
	_h1 uint32 = 0xEFCDAB89
	_h2 uint32 = 0x98BADCFE
	_h3 uint32 = 0x10325476
	_h4 uint32 = 0xC3D2E1F0

	_k0 uint32 = 0x5A827999
	_k1 uint32 = 0x6ED9EBA1
	_k2 uint32 = 0x8F1BBCDC
	_k3 uint32 = 0xCA62C1D6
)

func Sha1Sum(data []byte) [20]byte {
	// Calculate original message length in bytes
	l := len(data)

	// Calculate 0 bits padding length
	diff := (l + 1 + 8) % chunkSize
	p := chunkSize - diff

	// Write message length in bits
	lBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(lBytes, uint64(l*8))

	// Prepare message so that bit length mod 512 equals 0
	message := make([]byte, l+p+1+8)
	copy(message[:], data)
	copy(message[l:], []byte{0b10000000})
	copy(message[l+1+p:], lBytes)

	// Initialize h0, h1, h2, h3, h4 with predefined constants
	h0 := _h0
	h1 := _h1
	h2 := _h2
	h3 := _h3
	h4 := _h4

	// Process the message in successive 512-bit (64-byte) chunks
	for i := 0; i < len(message); i += chunkSize {
		chunk := message[i:]

		w := make([]uint32, 80)

		// Break chunk into sixteen 32-bit big-endian words
		for i := 0; i < 16; i++ {
			iBytes := i * wordSize
			w[i] = uint32(chunk[iBytes+0])<<24 | uint32(chunk[iBytes+1])<<16 | uint32(chunk[iBytes+2])<<8 | uint32(chunk[iBytes+3])
		}

		// Extend the sixteen 32-bit words into eighty 32-bit words
		for i := 16; i < 80; i++ {
			w[i] = bits.RotateLeft32(w[i-3]^w[i-8]^w[i-14]^w[i-16], 1)
		}

		// Initialize hash value for this chunk
		a := h0
		b := h1
		c := h2
		d := h3
		e := h4

		// Main loop
		for i := 0; i < 80; i++ {
			var f uint32
			var k uint32

			if i < 20 {
				f = (b & c) | ((^b) & d)
				k = _k0
			} else if i < 40 {
				f = b ^ c ^ d
				k = _k1
			} else if i < 60 {
				f = (b & c) | (b & d) | (c & d)
				k = _k2
			} else if i < 80 {
				f = b ^ c ^ d
				k = _k3
			}

			temp := bits.RotateLeft32(a, 5) + f + e + k + w[i]
			e = d
			d = c
			c = bits.RotateLeft32(b, 30)
			b = a
			a = temp
		}

		// Add this chunk's hash to result so far
		h0 = h0 + a
		h1 = h1 + b
		h2 = h2 + c
		h3 = h3 + d
		h4 = h4 + e
	}

	// Produce the final hash value (big-endian) as a 160-bit number
	var hash [20]byte
	binary.BigEndian.PutUint32(hash[:], h0)
	binary.BigEndian.PutUint32(hash[4:], h1)
	binary.BigEndian.PutUint32(hash[8:], h2)
	binary.BigEndian.PutUint32(hash[12:], h3)
	binary.BigEndian.PutUint32(hash[16:], h4)

	return hash
}
