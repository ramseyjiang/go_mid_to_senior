package strtransform

const MOD = 1000000007

func numberOfWays(s, t string, k int64) int {
	n := int64(len(s))
	f := func(k int64) int64 {
		return (pow(n-1, k, MOD) - pow(-1, k, MOD)) * pow(n, MOD-2, MOD) % MOD
	}

	ans := int64(0)
	if s == t {
		ans += (n - 1) * f(k-1)
	}
	cnt := int64(kmp(t[1:]+t[:len(t)-1], s))
	ans += cnt * f(k)
	return int(ans % MOD)
}

func pow(x, y, m int64) int64 {
	result := int64(1)
	for y > 0 {
		if y%2 == 1 {
			result = result * x % m
		}
		x = x * x % m
		y /= 2
	}
	return result
}

func kmp(text, word string) int {
	cnt := 0
	z := calcZ(word)
	i, j := 0, 0
	for i < len(text) {
		if word[j] != text[i] {
			j = z[j]
			if j < 0 {
				i++
				j++
			}
			continue
		}
		i++
		j++
		if j == len(word) {
			cnt++
			j = z[j]
		}
	}
	return cnt
}

func calcZ(s string) []int {
	z := make([]int, len(s)+1)
	z[0] = -1
	i, j := 1, 0
	for i < len(s) {
		if s[i] == s[j] {
			z[i] = z[j]
		} else {
			z[i] = j
			for j >= 0 && s[i] != s[j] {
				j = z[j]
			}
		}
		i++
		j++
	}
	z[i] = j
	return z
}
