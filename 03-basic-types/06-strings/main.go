package main

import "fmt"

func main() {
	// Normal String (Can not contain newlines, and can have escape characters like `\n`, `\t` etc)
	var website = "\thttps://www.callicoder.com\t\n"

	// Raw String (Can span multiple lines. Escape characters are not interpreted)
	var siteDescription = `\t\tCalliCoder is a programming blog where you can find
                           practical guides and tutorials on programming languages, 
                           web development, and desktop app development.\t\n`

	fmt.Println(website, siteDescription)
}
