package day6

func CalcMarkerPosition(stream string) int {
	return calcDistinctConsecutive(4, stream)
}

func CalcMessagePosition(stream string) int {
	return calcDistinctConsecutive(14, stream)
}

func calcDistinctConsecutive(distinct int, stream string) int {
	var buf = make([]rune, distinct)
Outer:
	for i, r := range stream {
		buf = append(buf[1:], r)
		if i < distinct {
			continue
		}
		for j := range buf[:len(buf)-1] {
			for k := j + 1; k < len(buf); k++ {
				if buf[j] == buf[k] {
					continue Outer
				}
			}
		}
		return i + 1
	}

	return -1
}
