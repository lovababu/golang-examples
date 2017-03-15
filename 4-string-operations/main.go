package main

import (
     "fmt"
     "strings"
     "unicode"
)

func main() {
	fmt.Println("============== Package strings usage =================")
	var str1 string = "Hello, World..!"
	//concatenation.
	fmt.Println("String : " + str1)

	str1 = str1 + "Welcome to Go Lang."

	fmt.Printf("After concatenation: %s \n", str1)

	//string functions.
	
	//calculate length of given string.
	fmt.Printf("String Length: %d \n", len(str1))
	
	var string1 string = "one"
	var string2 string = "two"
	
	//strings.Compare
	fmt.Println("is string1 and string2 equal: %d", strings.Compare(string1, string2))
	
	string1 = "go lang contains demo!"
	
	//strings.Contains
	fmt.Println("Is string1 contains lang : ", strings.Contains(string1, "lang"))
	fmt.Println("Is string1 contains Lang : ", strings.Contains(string1, "Lang"))
	
	//strings.ContainsAny 
	fmt.Println("Is string1 contains any (- or i) :", strings.ContainsAny(string1, "-i"))
	
	//strings.ContainsRune
	//var r rune = '0002'
	//fmt.Println("Is string1 contains rune 0002", strings.ContainsRune(string1, r))
	
	//strings.Count
	fmt.Println("Count of a char in string1: ", strings.Count(string1, "a")) 
	
	//strings.EqualFold
	fmt.Println("Is both string interpreted as UTF-8: ", strings.EqualFold("Go", "go"))
	
	//strings.Fields
	fmt.Printf("Splits the string1: %q \n", strings.Fields(string1))
	
	//strings.FieldsFunc
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	
	fmt.Printf("Fileds Are : %q \n", strings.FieldsFunc(string1, f))
	
	//strings.HasPrefix
	fmt.Printf("string1 has prefix go:", strings.HasPrefix(string1, "go"))
	fmt.Printf("string1 has prefix og:", strings.HasPrefix(string1, "og"))
	
	//strings.Index	
	fmt.Println("demo index in string1:", strings.Index(string1, "demo"))
	fmt.Println("hello index in string1:", strings.Index(string1, "hello"))
	
	//if any char from second string found in first, returns index value.
	fmt.Println("hello find index any in string1: ", strings.IndexAny(string1, "hello"))
	
	var b byte = 32 //space char
	fmt.Println("index byte: ", strings.IndexByte(string1, b))
	
	s := []string{"Go", "Lang", "by", "Google"}
	fmt.Println("After concatinating array: ", strings.Join(s, "-"))
	
	//strings.LastIndex
	fmt.Println("last index of o in string1: ", strings.LastIndex(string1, "o"))
	
	//strings.LastIndexAny
	fmt.Println("last index of any(oin) in string1: ", strings.LastIndexAny(string1, "xnz"))
	
	//strings.LastIndexFunc
	f1 := func(c rune) bool {
		return unicode.IsLetter(c) && c == 'm'
	}
	 
	fmt.Println("last index func:", strings.LastIndexFunc(string1, f1))
	
	//strings.Map
	//transform function, turns the letter to upper case.
	transformfun := func(c rune) rune {
		if unicode.IsLetter(c) {
			return unicode.ToUpper(c)
		} else {
			return c
		} 
	}
	
	fmt.Println("after transforming : ", strings.Map(transformfun, string1))
	
	//strings.Repeat
	fmt.Println("repeated \"go \"  3times:", strings.Repeat("go ",3))
	
	//strings.Replace
	fmt.Println("replace demo with live in string1: ", strings.Replace(string1, "demo", "live", 1))
	
	//strings.Split
	fmt.Printf("slice the string1 with space : %q \n", strings.Split(string1, " "))
	
	//strings.SplitAfter
	fmt.Printf("slices hello-world-welcome-go by - : %q \n", strings.SplitAfter("hello-world-welcome-go", "-"))
	
	//strings.SplitAfterN
	fmt.Printf("slices hello-world-welcome-go by - : %q \n", strings.SplitAfterN("hello-world-welcome-go", "-", 2))

	//strings.SplitN
	fmt.Printf("2 slices of hello-world-welcome-go : %q \n", strings.SplitN("hello-world-welcome-go","-", 2))
	
	//strings.Title
	fmt.Println("title of string1 : ", strings.Title(string1))
	
	//strings.ToLower
	fmt.Println("GO LANG in lower case: ", strings.ToLower("GO LaNg"))
	
	//strings.ToLowerSpecial
	fmt.Println("GO LANG to lower special: ", strings.ToLowerSpecial( unicode.TurkishCase, "GO LaNg"))


	var reader = strings.NewReader("String from Reader.");
	fmt.Println("Length is: " , reader.Len())
	var bArray []byte
	bArray = make([]byte, 6)
	var i, err = reader.ReadAt(bArray, 0);
	fmt.Println("No of bytes Read: ", i)
	fmt.Println("Byte value is: ",  string(bArray))
	fmt.Println("Error: ", err)
}