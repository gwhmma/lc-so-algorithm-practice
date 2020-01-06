package main

import "fmt"

/*
LL今天心情特别好,因为他去买了一副扑克牌,发现里面居然有2个大王,2个小王(一副牌原本是54张^_^)...
他随机从中抽出了5张牌,想测测自己的手气,看看能不能抽到顺子,如果抽到的话,他决定去买体育彩票,嘿嘿！！
“红心A,黑桃3,小王,大王,方片5”,“Oh My God!”不是顺子.....LL不高兴了,
他想了想,决定大\小 王可以看成任何数字,并且A看作1,J为11,Q为12,K为13。
上面的5张牌就可以变成“1,2,3,4,5”(大小王分别看作2和4),“So Lucky!”。LL决定去买体育彩票啦。
现在,要求你使用这幅牌模拟上面的过程,然后告诉我们LL的运气如何， 如果牌能组成顺子就输出true，
否则就输出false。为了方便起见,你可以认为大小王是0。
*/

// 数组中数字范围: 1--13 有4组   有4个0
// 给定一个长度为5的数组 判断这个数组里的数字是不是连续的  0可以代表任意数
// 数组中最大值 max 最小值 min 满足 max - min = 4
// 数组中除了0之外不能有重复的数字
func continuous(num []int) bool {
	var repeat = make(map[int]int)
	min := 14
	max := -1

	if len(num) < 5 {
		return false
	}

	for _, v := range num {
		if v < min && v != 0 {
			min = v
		}
		if v > max && v != 0 {
			max = v
		}
		repeat[v]++
	}

	for k, v := range repeat {
		if k != 0 && v > 1 {
			return false
		}
	}

	if max-min == 4 {
		return true
	}

	// 需要统计在min和max之间的数字出现了几个
	// 例如这种情况(8, 7, 0, 10, 0) ---> (7,8,9,10,11)  max-min = 3 但是 0 只出现了2次
	// 为了下面  max-min-mid == repeat[0] 的判断  所以需要减去min 和 max之间出现数字的个数
	mid := 0
	for k, _ := range repeat {
		if k != 0 && k > min && k < max {
			mid++
		}
	}

	if max-min-mid == repeat[0] {
		return true
	}

	return false
}

func main() {
	num := []int{8, 7, 0, 10, 0}
	fmt.Println(continuous(num))
}
