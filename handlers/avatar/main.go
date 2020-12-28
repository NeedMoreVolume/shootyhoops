package avatar

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func SetAvatar(s *discordgo.Session, message string) error {
	urlCheck := strings.Split(message, "set-avatar")

	if len(urlCheck) < 2 {
		return errors.New("You need to provide a URL to set a new avatar.")
	}

	imageUrl := strings.TrimSpace(urlCheck[1])
	if !isValidUrl(imageUrl) {
		return errors.New("Is shootyhoops a joke to you? Do you think this is a game? Give me a valid URL to get the new avatar image from, or you will get nothing... and like it.")
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", imageUrl, nil)
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("I can't seem to get the image from the URL provided... Are you sure you gave me a URL I can actually get?")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("I can't read the body of the response from the URL you asked me to get the avatar image from.")
	}

	contentType := http.DetectContentType(body)

	fmtedAvatar := fmt.Sprintf("data:%s;base64,%s", contentType, base64.StdEncoding.EncodeToString(body))
	_, err = s.UserUpdate("", "", "", fmtedAvatar, "")
	if err != nil {
		return err
	}
	return nil
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
