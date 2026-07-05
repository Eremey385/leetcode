func isPalindrome(x int) bool {
 if x < 0 {return false} 

    result := 0
    y:=x
 for y > 0 {
    num := y % 10
    y = y / 10
    result = (result * 10) + num
 }  
 if result == x {
    return true
 }else {
    return false
    }
}