/*
An integer has sequential digits if and only if each digit in the number is one more than the previous digit.

Return a sorted list of all the integers in the range [low, high] inclusive that have sequential digits.

 

Example 1:

Input: low = 100, high = 300
Output: [123,234]

*/
// nextSequential => 123 return 234 given 234 return 345
//1 123456789
import "strconv"

func firstNum(num int) int {
    lenNum := len(strconv.Itoa(num))
    num = 0
    for i := 1; i <= lenNum; i++ {
        num *= 10
        num += i
    }

    return num
}

func findAllSequences(low int, high int, num int, ans *[]int)  {
    if num > high {
        return
    }

    tmp := num
    tmp_str := strconv.Itoa(num)
    lenNum := len(tmp_str)
    i := lenNum+1

    if tmp >= low {
        *ans = append(*ans, tmp)
    }
    
    // rotate number and add
    for tmp <= high && i <= 9 {
        tmp_str = tmp_str[1:lenNum]+strconv.Itoa(i)
        tmp, _ = strconv.Atoi(tmp_str)

        if tmp <= high && tmp >= low {
            *ans = append(*ans, tmp)
        }

        i++
    }

    findAllSequences(low, high, (num*10)+lenNum+1, ans);
}

func sequentialDigits(low int, high int) []int {
    startNum := firstNum(low)
    ans :=  make([]int, 0, 0)
    
    findAllSequences(low, high, startNum, &ans)

    return ans
}