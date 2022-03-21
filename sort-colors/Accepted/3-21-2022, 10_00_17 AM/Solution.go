// https://leetcode.com/problems/sort-colors

func sortColors(a []int) {
	i, j, k := 0, 0, len(a)-1

	for j <= k {
		switch a[j] {
		case 0:
			a[i], a[j] = a[j], a[i]
			i++
			j++
		case 1:
			j++
		case 2:
			a[j], a[k] = a[k], a[j]
			k--
		}
	}

}