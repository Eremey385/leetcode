func lengthOfLongestSubstring(s string) int {
l := 0
r := 0
mp := map[byte]int{}
curlen:= 0
best := 0
for r < len(s){
    value ,ok:= mp[s[r]]

    if ok {
        for l <= value{
        l ++
        curlen -= 1
        }
        mp[s[r]] = r
        r++
        curlen += 1
    }else{
        mp[s[r]] = r
        r++
        curlen += 1
    }
    if curlen > best {best = curlen}
    
}
	return best
}