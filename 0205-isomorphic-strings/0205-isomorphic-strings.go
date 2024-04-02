func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sourceMap := make(map[byte]int)
    targetMap := make(map[byte]int) 

    for i := 0; i < len(s); i++ {
        if sourceMap[s[i]] != targetMap[t[i]] {
            return false 
        }
        sourceMap[s[i]] = i + 1
        targetMap[t[i]] = i + 1
    }

    return true
}