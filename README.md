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

### Using Infisical

You can use [Infisical](https://infisical.com/) for retrieving your access token:

```json
{
  "infisical": {
    "client_id": "abcd-efgh-ijkl-mnop",
    "client_secret": "0123456789abcdefghijklmnop",

    "workspace_id": "012345abcdefg",
    "environment": "dev",
    "secret_type": "shared",

    "key_path": "/path/to/your/KEY_TO_ACCESS_TOKEN"
  }
}
```

## Run

```bash
$ pb-send [any message to send]
```

## License

MIT

