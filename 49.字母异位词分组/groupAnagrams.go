package groupAnagrams

//算术基本定理，又称为正整数的唯一分解定理，即：
//每个大于1的自然数，
//要么本身就是质数，
//要么可以写为2个以上的质数的积，
//而且这些质因子按大小排列之后，写法仅有一种方式。
func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	if len(strs) == 0 {
		return res
	}
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
	m := make(map[int][]string)

	for _, v := range strs {
		product := 1
		for _, c := range v {
			product *= primes[c - 'a']
		}
		m[product] = append(m[product], v)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
