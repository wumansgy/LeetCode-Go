package main

func main(){

}

// 方法一 暴力法
func twoSum(nums []int, target int) []int {
	result:=make([]int,2)  //定义一个答案的切片
	n:=len(nums)  //切片的长度
	for i:=0;i<n;i++{   //外层循环
		for j:=i+1;j<n;j++{
			if nums[i]+nums[j]==target{
				result[0],result[1]=i,j
				return result
			}
		}
	}
	return result

}

