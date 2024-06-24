package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const baseURL = "http://localhost:10769"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cidercli <command> [arguments]")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "active":
		getRequest("/active")
	case "currentPlayingSong":
		getRequest("/currentPlayingSong")
	case "addToLibrary":
		getRequest("/addToLibrary")
	case "isPlaying":
		getRequest("/isPlaying")
	case "toggleAutoplay":
		getRequest("/toggleAutoplay")
	case "playPause":
		getRequest("/playPause")
	case "play":
		getRequest("/play")
	case "pause":
		getRequest("/pause")
	case "stop":
		getRequest("/stop")
	case "next":
		getRequest("/next")
	case "previous":
		getRequest("/previous")
	case "seekto":
		if len(args) < 1 {
			fmt.Println("Usage: cidercli seekto <time>")
			return
		}
		getRequest("/seekto/" + args[0])
	case "show":
		getRequest("/show")
	case "hide":
		getRequest("/hide")
	case "album":
		if len(args) < 1 {
			fmt.Println("Usage: cidercli album <id>")
			return
		}
		getRequest("/album/" + args[0])
	case "rating":
		if len(args) < 3 {
			fmt.Println("Usage: cidercli rating <type> <id> <rating>")
			return
		}
		putRequest("/rating/" + args[0] + "/" + args[1] + "/" + args[2])
	case "audio":
		if len(args) < 1 {
			getRequest("/audio")
		} else {
			getRequest("/audio/" + args[0])
		}
	default:
		fmt.Println("Unknown command:", command)
	}
}

func getRequest(endpoint string) {
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Response (%d): %s\n", resp.StatusCode, string(body))
}

func putRequest(endpoint string) {
	req, err := http.NewRequest(http.MethodPut, baseURL+endpoint, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Response (%d): %s\n", resp.StatusCode, string(body))
}
