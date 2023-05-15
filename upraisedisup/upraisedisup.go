package upraisedisup

import (
	"fmt"
	"net/http"
	"time"
)

func CheckUpraised() {
	startTime := time.Now()

	links := []string{
		"https://www.upraised.co/",
		"https://www.upraised.co/app/",
		"https://www.upraised.co/blog/",
		"https://hire.upraised.co/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// receiving response msgs via channel
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	//checking infinitely
	for l := range c {
		go checkLink(l, c)
	}

	fmt.Println("Total time taken for execution: ", time.Since(startTime))
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	result := ""
	if err != nil {
		result = fmt.Sprintf("Oops! %v is down, please check \n", link)
		fmt.Printf(result)
		c <- link
		return
	}
	result = fmt.Sprintf("Yahoo! %v is up and running, chai break ? \n", link)
	fmt.Printf(result)
	c <- link
	return

}
