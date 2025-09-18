package go_channels

import "fmt"

func ChannelRunner() {
	pages := []string{"A", "B", "C"}

	ch := make(chan string)

	go download(pages, ch)

	extract(ch)

}

func download(pages []string, ch chan string) {
	for _, page := range(pages) {
		// download the content
		content := page + " content"

		fmt.Println(content + " downloaded")

		ch <- content
	}
	close(ch)
}

func extract (ch chan string) {
	for content := range(ch) {
		fmt.Println(content + " extraced")
	}
} 