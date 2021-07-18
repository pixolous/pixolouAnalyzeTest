package main

import (
	"fmt"
	"os"

	"github.com/pixolous/pixolousAnalyze"

	"gocv.io/x/gocv"
)

func main() {
	//groups := make([][]string, 0)

	pathToHash := make(map[string]string)
	for i := 2014; i <= 2040; i++ {

		path := fmt.Sprintf("%s%d%s", "pics/IMG-", i, ".jpg")
		image := gocv.IMRead(path, gocv.IMReadColor)
		if image.Empty() {
			fmt.Printf("Failed to read image: %s\n", path)
			os.Exit(1)
		}
		defer image.Close()

		hash := pixolousAnalyze.AHash(image)
		pathToHash[path] = hash

		fmt.Println(hash)

	}

	hashes := make([]string, 0, len(pathToHash))
	for p := range pathToHash {
		hashes = append(hashes, pathToHash[p])
	}
	// for i := range pathToHash {
	// 	for j := range pathToHash {
	// 		if i != j {
	// 			fmt.Println("Similarity", i, j, pixolousAnalyze.hashSimilarity(pathToHash[i], pathToHash[j]))
	// 		}
	// 	}
	// }

	fmt.Println(pixolousAnalyze.GetSimilarGrouped(pathToHash))

}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
