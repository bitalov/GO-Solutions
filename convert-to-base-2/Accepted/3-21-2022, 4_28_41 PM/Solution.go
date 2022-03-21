// https://leetcode.com/problems/convert-to-base-2

func swap(B []byte) {
	i, j := 0, len(B)-1
	for i < j {
		B[i], B[j] = B[j], B[i]
		i++
		j--
	}
}

func baseNeg2(N int) string {
	if N == 0 {
		return "0"
	}

	B := make([]byte, 0, 30)
	for N > 0 {
		switch N & 3 {
		case 0, 1:
			B = append(B, byte(N&1)+'0', '0')
		default: // 2,3
			B = append(B, byte(N&1)+'0', '1')
			N += 4
		}
		N >>= 2
	}

	swap(B)

	if B[0] == '0' { 
		B = B[1:]
	}

	return string(B)
}

