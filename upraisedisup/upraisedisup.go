package upraisedisup

import (
	"fmt"
	"net/http"
)

func CheckUpraised() {
	links := []string{
		"https://www.upraised.co/",
		"https://www.upraised.co/app/",
		"https://www.upraised.co/blog/",
		"https://hire.upraised.co/",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("Oops! %v is down, please check \n", link)
		return
	}
	fmt.Printf("Yahoo! %v is up and running, chai break ? \n", link)

}
