package iterators

import (
	"fmt"
	"strings"
	"regexp"
)

func Strings() {
	fmt.Println("Strings as slices:")
	stringsAsSlices()

	fmt.Println("String transformation:")
	transformStrings()

	fmt.Println("String regexp:")
	useRegexp()
}

func stringsAsSlices() {
	var str = "Hello, World!"
	fmt.Println(str[0:5])

	for i, c := range str {
		// Range is iterating over the unicode code points of the string.
		// So str[i] is the i'th byte of the string
		fmt.Println("i: ", i, "| unicode char: ", c, "| char: ", string(c))
	}
}

func transformStrings() {
	str := "hello, world!"
	fmt.Println("strings.ToUpper: ", strings.ToUpper(str))
	fmt.Println("strings.ToLower: ", strings.ToLower(str))
	fmt.Println("string contains 'z': ", strings.ContainsRune(str, 'z'))
	fmt.Println(str[strings.Index(str, "world!"):])
	fmt.Println("last 3 characters: ", str[len(str) -3:])
}

func useRegexp() {
	phoneNumber := "+11234567890"
	re := regexp.MustCompile(`(\d{1})(\d{3})(\d{3})(\d{4})`)
	newStr := re.ReplaceAllString(phoneNumber, "$1 ($2) $3-$4")
	fmt.Println("newStr: ", string(newStr))
}