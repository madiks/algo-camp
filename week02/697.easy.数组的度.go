package week02

// https://leetcode-cn.com/problems/degree-of-an-array/
// https://leetcode-cn.com/submissions/detail/190787561/

type numstat struct {
	cnt   int
	start int
	end   int
}

func findShortestSubArray(nums []int) int {
	dict := map[int]*numstat{}
	for k, v := range nums {
		if _, ok := dict[v]; !ok {
			dict[v] = &numstat{
				start: k,
			}
		}
		dict[v].cnt++
		dict[v].end = k
	}
	ans, maxCnt := 0, 0
	for _, v := range dict {
		vl := v.end - v.start + 1
		if v.cnt > maxCnt {
			maxCnt = v.cnt
			ans = vl
		} else if v.cnt == maxCnt {
			if ans > vl {
				ans = vl
			}
		}
	}
	return ans
}
