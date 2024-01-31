/* 
150. You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22

Stack the numbers, when a sign is found, unstack two numbers, evaluate, stack the result.
 */

 import "strconv"

 type Stack struct {
	 items []int
 }
 
 func (s *Stack) Push(data int) {
	 s.items = append(s.items, data)
 }
 
 func (s *Stack) Pop() {
	 if s.IsEmpty() {
		 return
	 }
	 s.items = s.items[:len(s.items)-1]
 }
 
 func (s *Stack) Top() (int, error) {
	 if s.IsEmpty() {
		 return 0, fmt.Errorf("stack is empty")
	 }
	 return s.items[len(s.items)-1], nil
 } 
 
 func (s *Stack) Len() int {
	 return len(s.items)
 }
 
 func (s *Stack) IsEmpty() bool {
	 if len(s.items) == 0 {
		 return true
	 }
	 return false
 }
 
 func evalRPN(tokens []string) int {
	 stack := Stack{}
 
	 for _, token := range tokens {
		 if isSignal(token) {
			 evaluate(string(token), &stack)
		 } else {
			 val, _ := strconv.Atoi(token)
			 stack.Push(val)
		 }
	 } 
 
	 val, _ := stack.Top()
	 return val;
 }
 
 func isSignal(token string) bool {
	 return (token == "*" || token == "/" || token == "+" || token == "-");
 }
 
 
 func evaluate(token string, s *Stack) {
	 result := 0
	 a,b,error := popOperantors(s)
 
	 if(error != nil) {
		 return
	 }
 
	 if token == "*" {
		 result = b*a
	 } else if token == "/" {
		 result = b/a
	 } else if token == "+" {
		 result = a+b
	 } else if token == "-" {
		 result = b-a
	 }
 
	 s.Push(result)
 }
 
 func popOperantors(s *Stack) (int, int, error) {
	 a, e1 := s.Top()
	 s.Pop()
	 b, e2 := s.Top()
	 s.Pop()
 
	 if e1 != nil || e2 != nil {
		 return 0, 0, fmt.Errorf("stack is empty")
	 }
 
	 return a,b,nil
 }
 