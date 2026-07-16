func isSubsequence(s string, t string) bool {
    k := 0
    m := 0
    for k < len(s) && m < len(t){

    if s[k] == t[m] {
        k++
        m++
    }else {
        m++
    }
}
if k == len(s){
    return true
}
return false
}