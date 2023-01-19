package day20

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestCircularBufferEncryptionTestData(t *testing.T) {
	res := CircularBufferDecrypt(data.EncryptedCoordinatesTest)
	fmt.Printf("Result of encrypting test data: %v\n", res)
}

func BenchmarkCircularBufferEncryptionTestData(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CircularBufferDecrypt(data.EncryptedCoordinatesTest)
	}
}

func TestCircularBufferEncryptionKeyTestData(t *testing.T) {
	res := CircularBufferDecryptWithEncryptionKey(data.EncryptedCoordinatesTest, 811589153, 10)
	fmt.Printf("Result of encrypting test data: %v\n", res)
}

func TestCircularBufferEncryptionKeyPuzzleData(t *testing.T) {
	res := CircularBufferDecryptWithEncryptionKey(data.EncryptedCoordinatesPuzzle, 811589153, 10)
	fmt.Printf("Result of encrypting test data: %v\n", res)
}

func TestCircularBufferEncryptionPuzzleData(t *testing.T) {
	res := CircularBufferDecrypt(data.EncryptedCoordinatesPuzzle)
	fmt.Printf("Result of encrypting test data: %v\n", res)
}

func TestMoveCircularBufferLeft(*testing.T) {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	circular := make([]*int, 0, len(test))

	for i := range test {

		circular = append(circular, &test[i])
	}

	offset := -20

	moveInCircularBufferLeft(offset, 10, circular)

}

func TestMoveCircularBufferRight(*testing.T) {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	circular := make([]*int, 0, len(test))

	for i := range test {

		circular = append(circular, &test[i])
	}

	offset := 20

	moveInCircularBufferRight(offset, 0, circular)

}
