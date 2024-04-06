func mergeAlternately(word1 string, word2 string) string {
    length := 0
    var mergeString []byte 
    if len(word1) >= len(word2){
        length = len(word1)
        for i := 0; i < length; i++{
            mergeString = append(mergeString, word1[i])
            if len(word2) > i{
                mergeString = append(mergeString, word2[i])
            }
        }
    }else{
        length = len(word2)
        for i := 0; i < length; i++{
            if len(word1) > i{
                mergeString = append(mergeString, word1[i])
            }
            mergeString = append(mergeString, word2[i])    
        }
    }


    return string(mergeString)   
}