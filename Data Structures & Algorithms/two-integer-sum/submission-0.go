func twoSum(nums []int, target int) []int {
    var result []int
    for i := 0; i < len(nums); i++{
        currV := nums[i]
        for j := i+1; j < len(nums); j++{
            if(nums[j] + currV == target){
                result = []int{i, j}
                return result
            }
        }
    }
    return result
}
