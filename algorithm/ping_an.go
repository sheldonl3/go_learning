package main

func is_huiwen(str string) bool {
	l := len(str)
	if l == 0 {
		return false
	}
	tmp := make([]rune, 0)
	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch >= 'A' && ch <= 'Z' {
			tmp = append(tmp, ch)
		}
	}
	l2 := len(tmp)

	for i := 0; i < l2/2; i++ {
		if tmp[i] != tmp[l2-1-i] {
			return false
		}
	}
	return true
}

func find2(str string) int {
	l := len(str)
	if l <= 1 {
		return -1
	}
	map_string := make(map[rune]int)
	for _, v := range str {
		_, got := map_string[v]
		if !got {
			map_string[v] = 1
		} else {
			map_string[v] += 1
		}
	}

	time := 0
	for k, v := range str {
		if map_string[v] == 1 {
			time++
		}
		if time == 2 {
			return k
		}
	}
	return -1
}


//func main() {
//	//reader := bufio.NewReader(os.Stdin)
//	//str, err := reader.ReadString('\n')
//	//if err != nil {
//	//	return
//	//}
//	input := ScanLine()
//	fmt.Println(is_huiwen(input))
//	fmt.Printf("%d", find2(input))
//}
