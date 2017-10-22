# slack-wifi-status

## Install

```
go get github.com/polidog/slack-wifi-status
```

## Setting

```config.toml
// ~/.slack-wifi-status.toml

[slack]
token = "your-slack-token"

[wifi]
    [[wifi.list]]
    name = "your-target-wifi-name"
    message = "status message"
    emoji = ":office:" // your emoji 

    [[wifi.list]]
    name = "home-wifi"
    message = "status message"
    emoji = ":house:" // your emoji 
```

## Use

```
$ go run $GOPATH/srcslack-wifi-status
```


