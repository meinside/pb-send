# pb-send

Send messages through [Pushbullet](https://docs.pushbullet.com/) API.

## Install

```bash
$ go install github.com/meinside/pb-send@latest
```

## Configuration

Get your access token from [here](https://www.pushbullet.com/#settings/account),

then create a file named `config.json` in your `$XDG_CONFIG_HOME/pb-send/` directory:

```json
{
	"access_token": "PUT_YOUR_ACCESS_TOKEN_HERE"
}
```

You can also use [Infisical](https://infisical.com/) for retrieving your access token:

```json
{
  "infisical": {
    "workspace_id": "012345abcdefg",
    "token": "st.xyzwabcd.0987654321.abcdefghijklmnop",
    "environment": "dev",
    "secret_type": "shared",
    "key_path": "/path/to/your/KEY_TO_ACCESS_TOKEN"
  }
}
```

NOTE: It only supports E2EE-disabled Infisical workspaces for now.

## Run

```bash
$ pb-send [any message to send]
```

## License

MIT

