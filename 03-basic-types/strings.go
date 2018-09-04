package main
import "fmt"

func main() {
	var website = "\thttps://www.callicoder.com\t\n"

	var siteDescription = `\t\tCalliCoder is a programming blog where you can find
                           practical guides and tutorials on programming languages, 
                           web development, and desktop app development.\t\n`

	fmt.Println(website, siteDescription)
}
