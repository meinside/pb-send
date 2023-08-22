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
    "workspace_id": "012345abcdefg",
    "token": "st.xyzwabcd.0987654321.abcdefghijklmnop",
    "environment": "dev",
    "secret_type": "shared",
    "key_path": "/path/to/your/KEY_TO_ACCESS_TOKEN"
  }
}
```

If your Infisical workspace's E2EE setting is enabled, you also need to provide your API key:

```json
{
  "infisical": {
    "e2ee": true,
    "api_key": "ak.1234567890.abcdefghijk",

    "workspace_id": "012345abcdefg",
    "token": "st.xyzwabcd.0987654321.abcdefghijklmnop",
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

