// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

func firstUniqChar(s string) int {
   letters :=  make([]int, 26)
    idx := -1 

   for _, c := range s {
       letters[int(c-'a')]++
   }

    for i, c := range s {
       if letters[int(c-'a')] == 1 {
           idx = i
           break
       }
   }

    return idx
}
