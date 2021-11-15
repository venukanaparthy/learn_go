//wait-group-example.go
package concurrency

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func waitExample() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	filesInHomeDir, err := ioutil.ReadDir(homeDir)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(filesInHomeDir))

	fmt.Println("Files in home", homeDir)
	for _, file := range filesInHomeDir {
		go func(f os.FileInfo) {
			defer wg.Done()
			fmt.Println(f.Name())
		}(file)
	}

	wg.Wait()
	fmt.Println("Done !!")
}
