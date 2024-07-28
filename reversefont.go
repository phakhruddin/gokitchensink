package main

import "fmt"

func reverse(s string) string {
	c := []byte(s)
	fmt.Println("s ", s, "c ", c)
	for i := 0; i < len(c)/2; i++ {
		//fmt.Println("c[i]:", c[i], ",s[i]:", s[i], "string(c[i]), ", string(c[i]), ",string s[i]:", string(s[i]))
		//fmt.Println("before swap at i,", i, ": c[i]:", c[i], ",c[len(c)-1-i],:", c[len(c)-1-i], "string(c[i]), ", string(c[i]), ",string c[len(c)-1-i]:", string(c[len(c)-1-i]), "string(c):", string(c))
		c[len(c)-1-i], c[i] = c[i], c[len(c)-1-i]
		//fmt.Println("AFTER SWAP at i,", i, ": c[i]:", c[i], ",c[len(c)-1-i],:", c[len(c)-1-i], "string(c[i]), ", string(c[i]), ",string c[len(c)-1-i]:", string(c[len(c)-1-i]), "string(c):", string(c))
	}
	return string(c)
}

func main() {
	fmt.Println(reverse("Google"))
	fmt.Println(reverse("ABCD"))
	fmt.Println(reverse("Plan"))
	fmt.Println(reverse("Planning"))
	fmt.Println(reverse("Interaction"))
	fmt.Println(reverse("AmirHaziq"))
}

/* other solutions
//variant: using slice of bytes and conversions
func reverse(s string) string {
	sl := []byte(s)
	var rev [100]byte
	j := 0
	for i:=len(sl)-1; i >= 0; i-- {
		rev[j] = sl[i]
		j++
	}
	strRev := string(rev[:len(sl)])
	return strRev
}

// variant: "in place" using swapping _one slice
func reverse(s string) string {
	sl2 := []byte(s)
	for i, j := 0, len(sl2) - 1; i < j; i, j = i+1, j-1 {
		sl2[i], sl2[j] = sl2[j], sl2[i]
	}
	return string(sl2)
}

//variant: using [] int for runes (necessary for Unicode-strings!):
func reverse(s string) string {
	runes := []rune(s)
	n, h := len(runes), len(runes)/2
	for i:= 0; i < h; i ++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}*/
