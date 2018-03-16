package concurrency

type URLchecker func(string) bool

func WebsiteChecker(isOK URLchecker, urls []string) (results []bool) {
	for _, url := range urls {
		results = append(results, isOK(url))
	}

	return
}
