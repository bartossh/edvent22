package day20

import (
	"fmt"
	"strconv"
	"strings"
)

func CircularBufferDecryptWithEncryptionKey(d string, key, rounds int) int {
	coordsAfterZero := []int{1000, 2000, 3000}
	rows := strings.Split(d, "\n")
	original := make([]*int, 0)
	circular := make([]*int, 0)
	for _, r := range rows {
		if r == "" {
			continue
		}
		n, err := strconv.Atoi(r)
		if err != nil {
			panic(fmt.Errorf("unexpected, %w", err))
		}
		n = n * key
		original = append(original, &n)
		circular = append(circular, &n)
	}

	for rounds > 0 {
		for _, member := range original {
			circular = rearrangeCircBuf(member, circular)
		}
		rounds--
	}

	idxOfZero := findByValue(0, circular)
	if idxOfZero == -1 {
		panic("cannot find idx of zero")
	}

	sum := 0
	for _, v := range coordsAfterZero {
		offset := (idxOfZero + v) % len(circular)
		value := circular[offset]
		sum += *value
	}

	return sum
}

func CircularBufferDecrypt(d string) int {
	coordsAfterZero := []int{1000, 2000, 3000}
	rows := strings.Split(d, "\n")
	original := make([]*int, 0)
	circular := make([]*int, 0)
	for _, r := range rows {
		if r == "" {
			continue
		}
		n, err := strconv.Atoi(r)
		if err != nil {
			panic(fmt.Errorf("unexpected, %w", err))
		}
		original = append(original, &n)
		circular = append(circular, &n)
	}

	for _, member := range original {
		circular = rearrangeCircBuf(member, circular)
	}

	idxOfZero := findByValue(0, circular)
	if idxOfZero == -1 {
		panic("cannot find idx of zero")
	}

	sum := 0
	for _, v := range coordsAfterZero {
		offset := (idxOfZero + v) % len(circular)
		value := circular[offset]
		sum += *value
	}

	return sum
}

func rearrangeCircBuf(m *int, buf []*int) []*int {
	pos := findInCircBuf(m, buf)
	if pos == -1 {
		panic("member not found")
	}

	offset := (*m) % (len(buf) - 1)

	switch {
	case offset > 0:
		buf = moveInCircularBufferRight(offset, pos, buf)
	case offset < 0:
		buf = moveInCircularBufferLeft(offset, pos, buf)
	}

	return buf
}

func moveInCircularBufferLeft(offset, pos int, buf []*int) []*int {
	m := buf[pos]
	for offset < 0 {
		if pos == 0 {
			buf = append(buf[1:len(buf)-1], append([]*int{m}, buf[len(buf)-1])...)
			pos = len(buf) - 2
			offset++
			continue
		}

		buf = append(buf[:pos], buf[pos+1:]...)
		buf = append(buf[:pos-1], append([]*int{m}, buf[pos-1:]...)...)
		pos--
		offset++
	}
	return buf
}

func moveInCircularBufferRight(offset, pos int, buf []*int) []*int {
	m := buf[pos]
	for offset > 0 {
		if pos == len(buf)-1 {
			buf = append(buf[:1], append([]*int{m}, buf[1:len(buf)-1]...)...)

			pos = 1
			offset--
			continue
		}

		buf = append(buf[:pos], buf[pos+1:]...)
		buf = append(buf[:pos+1], append([]*int{m}, buf[pos+1:]...)...)
		pos++
		offset--
	}

	return buf
}

func findInCircBuf(m *int, buf []*int) int {
	for i, mm := range buf {
		if mm == m {
			return i
		}
	}
	return -1
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func findPosByValue(v int, buf []*int) int {
	for i, vv := range buf {
		if v == *vv {
			return i
		}
	}
	return -1
}

func findByValue(v int, buf []*int) int {
	for i, vv := range buf {
		if *vv == v {
			return i
		}
	}
	return -1
}
