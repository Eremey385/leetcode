func sumAndMultiply(n int) int64 {
    sl := []int{}
    x := 0
    sum := 0
  for n > 0 {
    ost := n % 10
    n /= 10 
    if ost != 0 {
        sl = append(sl,ost)
    }
  }  
  for i := len(sl) - 1; i >= 0; i--{
    x = x * 10 + sl[i]
    sum += sl[i]
  } 
  var final int64 = int64(sum) * int64(x)
  return final
}