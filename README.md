# slack-wifi-status

## Install

download file is [here](https://github.com/polidog/slack-wifi-status/releases).  
Unsupported Windows and Linux.

## Setting

```
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

or 

```
$ slack-wifi-status --config=~/foo.toml
```


### help

```
$ slack-wifi-status -h
```