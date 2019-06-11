package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
)

const (
	configFilepath = ".config/pb-send.json" // $HOME/.config/pb-send.json
)

type config struct {
	AccessToken string `json:"access_token"`
}

// loggers
var _stderr = log.New(os.Stderr, "", 0)

// load config file
func loadConf() (conf config, err error) {
	var usr *user.User
	if usr, err = user.Current(); err == nil {
		fpath := filepath.Join(usr.HomeDir, configFilepath)

		var bytes []byte
		if bytes, err = ioutil.ReadFile(fpath); err == nil {
			if err = json.Unmarshal(bytes, &conf); err == nil {
				return conf, nil
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
