func canMakeArithmeticProgression(arr []int) bool {
    diff := 0
for j:= 0; j < len(arr); j++{
    for i := 0; i < len(arr)-1; i++{
        if arr[i] > arr[i+1]{
            arr[i],arr[i+1] = arr[i+1],arr[i]
        }
    }
}
diff = arr[1] - arr[0]
for i := 0;i < len(arr)-1;i++{
    if arr[i+1]-arr[i] == diff{
    }else {
       return false
    }
}
return true
}