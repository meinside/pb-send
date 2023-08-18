package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/meinside/infisical-go"
	"github.com/meinside/infisical-go/helper"

	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
)

const (
	applicationName = "pb-send"
	configFilename  = "config.json" // $HOME/.config/pb-send/config.json
)

type config struct {
	// Pushbullet API Access Token,
	AccessToken string `json:"access_token,omitempty"`

	// or Infisical settings
	Infisical *struct {
		WorkspaceID        string               `json:"workspace_id"`
		Token              string               `json:"token"`
		Environment        string               `json:"environment"`
		SecretType         infisical.SecretType `json:"secret_type"`
		AccessTokenKeyPath string               `json:"key_path"`
	} `json:"infisical,omitempty"`
}

// loggers
var _stderr = log.New(os.Stderr, "", 0)

// load config file
func loadConf() (conf config, err error) {
	// https://xdgbasedirectoryspecification.com
	configDir := os.Getenv("XDG_CONFIG_HOME")

	// If the value of the environment variable is unset, empty, or not an absolute path, use the default
	if configDir == "" || configDir[0:1] != "/" {
		var homeDir string
		if homeDir, err = os.UserHomeDir(); err == nil {
			configDir = filepath.Join(homeDir, ".config", applicationName)
		}
	} else {
		configDir = filepath.Join(configDir, applicationName)
	}

	if err == nil {
		configFilepath := filepath.Join(configDir, configFilename)

		var bytes []byte
		if bytes, err = os.ReadFile(configFilepath); err == nil {
			if err = json.Unmarshal(bytes, &conf); err == nil {
				if conf.AccessToken == "" && conf.Infisical != nil {
					// read access token from infisical
					var accessToken string
					accessToken, err = helper.Value(
						conf.Infisical.WorkspaceID,
						conf.Infisical.Token,
						conf.Infisical.Environment,
						conf.Infisical.SecretType, conf.Infisical.AccessTokenKeyPath,
					)
					conf.AccessToken = accessToken
				}

				return conf, err
			}
		}
	}

	return config{}, err
}

func main() {
	var conf config
	var err error

	if conf, err = loadConf(); err == nil {
		switch len(os.Args) {
		case 1: // no params
			err = fmt.Errorf(`> usage:

$ %s [message to send]`, os.Args[0])
		default: // one or more params
			str := strings.Join(os.Args[1:], " ")

			client := pushbullet.New(conf.AccessToken)

			note := requests.NewNote()
			note.Title = str
			note.Body = str

			_, err = client.PostPushesNote(note)
		}
	}

	if err != nil {
		_stderr.Fatalf(err.Error())
	}
}
