package avatar

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetAvatar(message string) string {
	urlCheck := strings.Split(message, "set-avatar")

	if len(urlCheck) < 2 {
		return "not enough arguments to get a new avatar"
	}

	url := urlCheck[1]
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return "I can't seem to get the image from the url provided... Are you sure you gave me a URL I can actually get?"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Hey man, let's not be telling me to go get random shit on the internet, this is serious fucking shootyhoops business. I can't read the body of the response from the thing you asked me to get."
	}


}
