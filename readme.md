# Toast

A go package for Windows 10 toast notifications.

[![Go Reference](https://pkg.go.dev/badge/github.com/Shinyhero36/go-toast.svg)](https://pkg.go.dev/github.com/Shinyhero36/go-toast)

## CLI

As well as using go-toast within your Go projects, you can also utilise the CLI.


```cmd
pwd=$(pwd)

./toast.exe \
  --app-id "Example App" \
  --title "Hello World" \
  --message "Lorem ipsum dolor sit amet, consectetur adipiscing elit." \
  --icon "${pwd}\icon.png" \
  --audio "default" --loop \
  --duration "long" \
  --activation-arg "https://google.com" \
  --action "Open maps" --action-arg "bingmaps:?q=sushi" \
  --action "Open browser" --action-arg "http://github.com/Shinyhero36/go-toast/"
```

## Example

Examples are available in [examples](./examples).

## Screenshots

