# slack-wifi-status

## Install

```
go install github.com/polidog/slack-wifi-status/cmd
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
$ slack-wifi-status
```
