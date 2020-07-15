package leetcode

// 		字符          数值
// 		I             1
// 		V             5
// 		X             10
// 		L             50
// 		C             100
// 		D             500
// 		M             1000

// I 	II	 III	IV	 V	 VI	  VII	 VIII 	 IX	  X
//XI	XII	 XIII   XIV  XV  XVI  XVII   XVIII   XIX  XX

// 罗马数字转整数
func romanToInt(s string) int {

	m := make(map[rune]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	sum := 0
	lastV := 0
	for k, v := range s {
		if k > 0 && m[v] > lastV {
			sum -= 2 * lastV
		}
		lastV = m[v]
		sum += lastV
	}

	return sum
}
