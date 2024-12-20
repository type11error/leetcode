import (
    "strings"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


func mergeAlternately(word1 string, word2 string) string {
    len1 := len(word1)
    len2 := len(word2)

    l := min(len1, len2)

    var sb strings.Builder
    for i := range l {
        sb.WriteString(string(word1[i]))
        sb.WriteString(string(word2[i]))
    }

    if len1 > len2 {
        sb.WriteString(string(word1[l:]))
    }

    if len2 > len1 {
        sb.WriteString(string(word2[l:]))
    }
    
    return sb.String()
    
}
