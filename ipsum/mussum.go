package ipsum

import (
	"math/rand"
	"strings"
	"time"
)

// GetMussum returns a Mussum lorem ipsum text
func GetMussum(size int) (string, error) {
	// if required size is bigger than maximum
	// acceptable capacity, an error is returned
	if size > maxCap {
		return "", ErrLenghtTooBig
	}

	// seed rand
	rand.Seed(time.Now().UTC().Unix())

	// Builder is used to emprove performance
	var b strings.Builder
	// If size is negative, it will panic
	b.Grow(size)

	content, err := getIpsums()
	if err != nil {
		return "", err
	}

	list := content.Mussum

	// use count to control string size
	count := 0

	for size > count {
		// get a rand phrase from list
		r := rand.Intn(len(list))

		// this check avoids unnecessary memory allocation.
		// If write on buffer will exceed allocated memory
		// size, it will adjust content to fit on it.
		if len(list[r])+count > size {
			list[r] = (list[r])[:(size - count)]
		}

		// write to builder. The write function will not
		// allocate more memory.
		n, err := b.WriteString(list[r])
		if err != nil {
			return "", err
		}

		// update count, n bytes from string
		count += n

		// it is needed adjust the phrases to return
		// a valid text; so a space is added
		// (if possible) after phrases addition
		if count < size {
			// add a dot or/and space.
			err = b.WriteByte(byte(32))
			if err != nil {
				return "", err
			}

			count++
		}
	}

	return b.String(), nil
}
