package main

import (
	"fmt"
	"math/rand"

	// "strconv"
	"time"
)

func main() {
	fmt.Println("12块金子，其中一块重量小于其他，限使用天平3次将其找出。")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	golds := [12]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}
	golds[r.Int31n(12)] = 99
	fmt.Println("上帝视角：")
	for i := 0; i < 12; i++ {
		fmt.Printf("第%d块金子重量为：%d\n", (i + 1), golds[i])
	}
	fmt.Println("首先将金子分为4等份，每份3块")
	group := [][]int{{golds[0], golds[1], golds[2]}, {golds[3], golds[4], golds[5]}, {golds[6], golds[7], golds[8]}, {golds[9], golds[10], golds[11]}}
	a := sum(group[0])
	fmt.Printf("第一份金子重量为：%d\n", a)
	b := sum(group[1])
	fmt.Printf("第二份金子重量为：%d\n", b)
	c := sum(group[2])
	fmt.Printf("第三份金子重量为：%d\n", c)
	d := sum(group[3])
	fmt.Printf("第四份金子重量为：%d\n", d)

	fmt.Println("第1次使用天平比较第一份金子和第二份金子")
	r1 := compare(a, b)
	switch r1 {
	case 0:
		fmt.Println("第一份金子和第二份金子相等")
		fmt.Println("第2次使用天平比较第三份金子和第四份金子")
		r2 := compare(c, d)
		if r2 < 0 {
			fmt.Println("第三份金子较轻")
			fmt.Println("第3次使用天平比较第7块金子和第8块金子")
			r3 := compare(golds[6], golds[7])
			if r3 < 0 {
				fmt.Println("第7块金子较轻")
			} else if r3 > 0 {
				fmt.Println("第8块金子较轻")
			} else {
				fmt.Println("第9块金子较轻")
			}
		} else {
			fmt.Println("第四份金子较轻")
			fmt.Println("第3次使用天平比较第10块金子和第11块金子")
			r3 := compare(golds[9], golds[10])
			if r3 < 0 {
				fmt.Println("第10块金子较轻")
			} else if r3 > 0 {
				fmt.Println("第11块金子较轻")
			} else {
				fmt.Println("第12块金子较轻")
			}
		}
		break
	case -1:
		fmt.Println("第一份金子较轻")
		fmt.Println("第2次使用天平比较第1块金子和第2块金子")
		r2 := compare(golds[0], golds[1])
		if r2 < 0 {
			fmt.Println("第1块金子较轻")
		} else if r2 > 0 {
			fmt.Println("第2块金子较轻")
		} else {
			fmt.Println("第3块金子较轻")
		}
		break
	case 1:
		fmt.Println("第二份金子较轻")
		fmt.Println("第2次使用天平比较第4块金子和第5块金子")
		r2 := compare(golds[3], golds[4])
		if r2 < 0 {
			fmt.Println("第4块金子较轻")
		} else if r2 > 0 {
			fmt.Println("第5块金子较轻")
		} else {
			fmt.Println("第6块金子较轻")
		}
		break
	default:
		fmt.Println("无效的取值")
	}
}

func sum(items []int) int {
	r := 0
	for _, value := range items {
		r += value
	}
	return r
}

func compare(a int, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
