package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	infisical "github.com/infisical/go-sdk"
	"github.com/infisical/go-sdk/packages/models"
	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
	"github.com/tailscale/hujson"
)

const (
	applicationName = "pb-send"
	configFilename  = "config.json" // $HOME/.config/pb-send/config.json
)

type config struct {
	// Pushbullet API Access Token,
	AccessToken *string `json:"access_token,omitempty"`

	// or Infisical settings
	Infisical *struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`

		ProjectID   string `json:"project_id"`
		Environment string `json:"environment"`
		SecretType  string `json:"secret_type"`

		AccessTokenKeyPath string `json:"key_path"`
	} `json:"infisical,omitempty"`
}

// GetAccessToken returns your access token of Pushbullet
//
// (retrieve it from infisical if needed)
func (c *config) GetAccessToken() (accessToken *string, err error) {
	if (c.AccessToken == nil || len(*c.AccessToken) == 0) &&
		c.Infisical != nil {
		// read access token from infisical
		client := infisical.NewInfisicalClient(infisical.Config{
			SiteUrl: "https://app.infisical.com",
		})

		_, err = client.Auth().UniversalAuthLogin(c.Infisical.ClientID, c.Infisical.ClientSecret)
		if err != nil {
			_stderr.Printf("* failed to authenticate with Infisical: %s", err)
			return nil, err
		}

		var secret models.Secret
		secret, err = client.Secrets().Retrieve(infisical.RetrieveSecretOptions{
			SecretKey:   path.Base(c.Infisical.AccessTokenKeyPath),
			SecretPath:  path.Dir(c.Infisical.AccessTokenKeyPath),
			ProjectID:   c.Infisical.ProjectID,
			Type:        c.Infisical.SecretType,
			Environment: c.Infisical.Environment,
		})
		if err != nil {
			_stderr.Printf("* failed to retrieve Pushbullet access token from infisical: %s\n", err)
			return nil, err
		}

		c.AccessToken = &secret.SecretValue
	}

	return c.AccessToken, nil
}

// loggers
var _stderr = log.New(os.Stderr, "", 0)

// standardize given JSON (JWCC) bytes
func standardizeJSON(b []byte) ([]byte, error) {
	ast, err := hujson.Parse(b)
	if err != nil {
		return b, err
	}
	ast.Standardize()

	return ast.Pack(), nil
}

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
			if bytes, err = standardizeJSON(bytes); err == nil {
				if err = json.Unmarshal(bytes, &conf); err == nil {
					return conf, err
				}
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

			var accessToken *string
			if accessToken, err = conf.GetAccessToken(); err == nil {
				client := pushbullet.New(*accessToken)

				note := requests.NewNote()
				note.Title = str
				note.Body = str

				_, err = client.PostPushesNote(note)
			}
		}
	}

	if err != nil {
		_stderr.Fatalf(err.Error())
	}
}
