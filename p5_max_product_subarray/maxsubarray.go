package _maxsubarray

func maxProduct(nums []int) int {
	maxProduct := nums[0]
	minProduct := nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		n := nums[i]

		if n < 0 {
			maxProduct, minProduct = minProduct, maxProduct
		}
		maxProduct = max(maxProduct*n, n)
		minProduct = min(minProduct*n, n)

		res = max(res, maxProduct)
	}

	return res
}
