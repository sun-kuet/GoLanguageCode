package main

import "bufio"
import "os"
import "fmt"
import "strconv"

//Stack reference : http://play.golang.org/p/5LdPGqpdt0
type stack [] string

func (s stack) empty() bool { return len(s) == 0 }
func (s stack) top() string  { return s[len(s)-1] }
func (s *stack) push(i string)  { (*s) = append((*s), i) }
func (s *stack) pop() string {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}
//fuction for modifing a string
func newString(str string) string{
	ret := ""
	for i:=0 ; i< len(str)-2; i = i +1{
		ret = ret + string(str[i])
		//fmt.Println(ret)
	}
	ret = ret + string(')')
	return ret
}

// Evaluate the value of the modified result string
func evaluateValue(result [] string) string{
	var st stack
	result = append(result,")")

	for i:= 0; i< len(result); i = i+1{
		str := result[i]
		if str[0]>=48 && str[0]<=57{
			st.push(str)
		}else if(str=="+" || str=="-" || str=="*" || str=="/"){
			a := st.pop()
			b := st.pop()
			var res int
			var aint int
			var bint int
			if t, err := strconv.Atoi(a); err == nil {
				aint = t
			}
			if t, err := strconv.Atoi(b); err == nil {
				bint = t
			}

			if(str=="+"){
				res = aint + bint;
			}else if(str=="-"){
				res = bint - aint;
			}else if(str=="*"){
				res = aint * bint;
			}else if(str=="/"){
				res = bint / aint;
			}

			st.push(strconv.Itoa(res))
		}
	}
	return st.pop()
}

func main() {

	var st stack
	var result []string
	//next 2 lines for input 
    reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the Equation: ")
	input, _ := reader.ReadString('\n')

	input = newString(input)
	st.push("(")

	for i:=0; i< len(input); i = i +1 {

		if(input[i]>=48 && input[i]<=57){           // if numbers then
			temp:=""
			for input[i]>=48 && input[i]<=57{
				temp = temp + string(input[i])
				i = i+1
			}
			i = i-1 
			result = append(result,temp)
		}else if(string(input[i])=="("){  // if right parenthesis then
			st.push("(")
		}else if(string (input[i])=="+" || string (input[i])=="-" || string (input[i])=="*" || string (input[i])=="/"){   // if operator then
			if(string (input[i])=="+" || string (input[i])=="-"){
				for !st.empty() && st.top()!="("{
					result = append(result,st.pop())
				}
			}else{
				for !st.empty() && st.top()!="(" && st.top()!="+" && st.top()!="-"{
					result = append(result,st.pop())
				}
			}
			st.push(string(input[i]))
		}else if(string(input[i])==")"){                // if right parenthesis then
			for !st.empty() && st.top()!="(" {
				result = append(result,st.pop())
			}
			st.pop()
		}
	}

	finalValue := evaluateValue(result)                // evaluate the result

	fmt.Println("The Calculated Value :",finalValue)   // Print the result

}