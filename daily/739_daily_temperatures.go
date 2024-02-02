/*

739. Daily Temperatures

Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Solution: O(n)

In reverse order, stack the indexes. If the current value is greater than the top, unstack until there is some element greater or it is empty.
The ideia that once you found a number greater, the others will not be used anymore.

*/

type Stack struct {
    items []int
}

func (stack *Stack) Push(data int) {
    stack.items = append(stack.items, data)
}

func (stack *Stack) IsEmpty() bool {
    if len(stack.items) == 0 {
        return true
    }
    return false
}

func (stack *Stack) Pop() {
    if stack.IsEmpty() {
        return
    }
    stack.items = stack.items[:len(stack.items)-1]
}

func (stack *Stack) Top() int {
    return stack.items[len(stack.items)-1]
}

func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    nextHeat := make([]int, n)
    stack := Stack{}

    for i, _ := range temperatures {
        idx := n-i-1

        if(stack.IsEmpty()) {
            stack.Push(idx)
        } else {
            for !stack.IsEmpty() && temperatures[stack.Top()] <= temperatures[idx] {
                stack.Pop()
            }  

            if !stack.IsEmpty() {
                nextHeat[idx] = stack.Top()-idx
            }

            if stack.IsEmpty() || temperatures[stack.Top()] >= temperatures[idx] {
                stack.Push(idx)
            }
        }

    } 

    return nextHeat

}