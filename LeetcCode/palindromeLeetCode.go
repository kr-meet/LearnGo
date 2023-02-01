//leetcode - 9

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	rev := 0
	copy := x
	for x > 0 {
		rem := x % 10
		rev = (rev * 10) + rem
		x /= 10
	}

	if rev == copy {
		return true
	}
	return false
}
