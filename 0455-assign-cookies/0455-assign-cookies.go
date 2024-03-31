func findContentChildren(g []int, s []int) int {
    counter := 0
    sort.Ints(g)
	sort.Ints(s)

	for _, cookie := range s {
        if counter >= len(g){
            break
        }else if cookie >= g[counter]{
            counter++
        }
	}

    return counter
}