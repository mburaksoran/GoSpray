package main

import (
	"fmt"
	"github.com/mburaksoran/GoSpray/goftp"
	"sync"
)

func ftpSpray(wg *sync.WaitGroup, channelToCommunicate chan string, taskToRun task, storeResult *int) {
	defer wg.Done()
	internalCounter := 0
	if taskToRun.target.port == 0 {
		taskToRun.target.port = 21
	}
	for _, password := range taskToRun.passwords {
		for _, username := range taskToRun.usernames {
			if internalCounter >= *storeResult {
				ftpClient, err := goftp.NewFtp(stringifyTarget(taskToRun.target))
				if err != nil {
					panic(err)
				}

				if err = ftpClient.Login(username, password); err != nil {
					fmt.Print("-")
				} else {
					fmt.Print("+")
					channelToCommunicate <- username + ":" + password
				}
				ftpClient.Close()
				*storeResult++
			} else {

			}
			internalCounter++
		}
	}
}
