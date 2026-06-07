func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    smap := map[rune]int{}

    for _, v := range s {
        smap[v]++
    }

    for _, v := range t {
        smap[v]--
    }

    for _, v := range smap {
        if v != 0 {
            return false
        }
    }

    return true
}
