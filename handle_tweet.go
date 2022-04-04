package main

func handleTweet(bytes []byte, client Client) bool {
	var done bool = false

	// This check handles sporadic empty messages
	if len(bytes) >= 0 {
		tweet := Tweet{
			Error: false,
		}
		tweet = extractValues(bytes, tweet)

		// Check for empty tweet.MediaUrl to
		// prevent crash from panic in processing
		// images
		if tweet.MediaUrl != "" {

			// tmpFile, _ := ioutil.TempFile("", "*")
			// defer tmpFile.Close()

			// _, err := tmpFile.WriteString("hello")
			// if err != nil {
			// 	fmt.Printf("%v\n", err)
			// }

			urlAndSize := tweet.MediaUrl + "@-@" + tweet.Text

			// tmpFile.Seek(0, 0)
			// s := bufio.NewScanner(tmpFile)
			// for s.Scan() {
			// 	println(s.Text())
			// }

			client.imageUrlAndSizes <- urlAndSize
			client.imageUrlAndSizes <- ""

			// storeImageInServer(tweet, client)

			done = true
		}
	}

	return done
}

// func storeImageInServer(tweet Tweet, client Client) {
// 	// Download Image
// 	fileName, b, err := downloadImage(tweet.MediaUrl)
// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 	}

// 	// I should be able to just create the temp file
// 	// and temp file here instead of an actual file?

// 	// Resize the image
// 	resizeImage(fileName, b, tweet.Text)

// 	sendProcessedImageToServer(fileName, client)

// 	// and then hopefully delete the temp file?

// 	// Reply to tweet with URL to download
// 	// Excel pattern

// 	// fileLock := flock.New(fileName)

// 	// println(strconv.FormatBool(fileLock.Locked()))

// }
