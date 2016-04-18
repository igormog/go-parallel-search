func searchAndProcessResults(timeout int64, done <-chan struct{},
	results <-chan Result) {
	finish := time.After(time.Duration(timeout))
	for working := workers; working > 0; {
		select { // blocked
		case result := <-results:
			fmt.Printf("%s:%d:%s\n", result.filename, result.lino,
			result.line)
		case <-finish:
			fmt.Println("timed out")
			return // Time's up, so to finish with the existing results
		case <-done:
			working--
		}
	}
	for {
		select { // Not blocked
		case result := <-results:
		fmt.Printf("%s:%d:%s\n", result.filename, result.lino,
			result.line)
		case <-finish:
			fmt.Println("timed out")
			return // Time's up, so to finish with the existing results
		default:
			return
		}
	}
}