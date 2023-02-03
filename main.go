package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func checkUsers(users []string) {
	var wg sync.WaitGroup
	for u := range users {
		user := users[u]
		url := fmt.Sprintf("https://instagram.com/%s", user)
		wg.Add(1)
		go func() {
			defer wg.Done()
			output, err := exec.Command("curl", "-L", url).CombinedOutput()
			if err != nil {
				print("err", err)
			}
			if strings.Contains(string(output), "og:title") {
				fmt.Printf("[-] Username taken. | @%s\n", user)
			} else {
				fmt.Printf("[+] Username available. | @%s\n", user)
			}
		}()
	}
	wg.Wait()
}

func main() {
	if len(os.Args) > 1 {
		filePath := os.Args[1]
		readFile, err := os.Open(filePath)

		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var fileLines []string

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}
		readFile.Close()
		checkUsers(fileLines)
	} else {
		users := []string{"test", "hiuhquihfquehiuashdfu", "AvrilLavigne"}
		checkUsers(users)
	}
}
