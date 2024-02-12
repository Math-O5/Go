/*
    The problem is to identify elements in an integer array, nums, of size n, that appear more than ⌊n/3⌋ times and return them as an output.

*/

// Applying Boyer-Moore Majority Voting Algorithm.
func majorityElement(nums []int) []int {
    voteNumbers := []int{0,0} 
    candidateNumbers := []int{0,0}
    var majorNumbers []int

    for _, num := range nums {
        if voteNumbers[0] == 0 && num != candidateNumbers[1] {
           voteNumbers[0] = 1
           candidateNumbers[0] = num
        } else if voteNumbers[1] == 0 && num != candidateNumbers[0] {
           voteNumbers[1] = 1
           candidateNumbers[1] = num
        } else if candidateNumbers[0] == num {
            voteNumbers[0]++
        } else if candidateNumbers[1] == num {
            voteNumbers[1]++
        } else {
            voteNumbers[0]--
            voteNumbers[1]--
        }    
    } 

    c1, c2 := 0,0
    for _, value := range nums {
        if value == candidateNumbers[0] {
            c1 += 1
        } else if value == candidateNumbers[1] {
            c2 += 1
        }
    }

    if c1 > len(nums)/3 {
        majorNumbers = append(majorNumbers, candidateNumbers[0])
    }

    if c2 > len(nums)/3 {
        majorNumbers = append(majorNumbers, candidateNumbers[1])
    }

    return majorNumbers
}

// Simple intuitive Mapping
func majorityElement2(nums []int) []int {
    freqMap := make(map[int]int)
    var majorNumbers []int

    for _, num := range nums {
        freqMap[num] += 1
    } 

    for key, value := range freqMap {
        if value > len(nums)/3 {
            majorNumbers = append(majorNumbers, key)
        }
    }

    return majorNumbers
}