# pb-send

Sends a message through [Pushbullet](https://docs.pushbullet.com/).

## install

```bash
$ go get -u github.com/meinside/pb-send
```

## setup

Get your access token from [here](https://www.pushbullet.com/#settings/account),

then create a file named `pb-send.json` in your `$HOME/.config/` directory:

```json
{
	"access_token": "PUT_YOUR_ACCESS_TOKEN_HERE"
}
```

## run

```bash
$ $GOPATH/bin/pb-send [any message to send]
```

