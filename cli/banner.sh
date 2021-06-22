pwd=$(pwd)

bin/toast.exe \
  --app-id "Example App" \
  --title "Hello World" \
  --message "Lorem ipsum dolor sit amet, consectetur adipiscing elit." \
  --icon "${pwd}\icon.png" \
  --audio "default" --loop \
  --duration "long" \
  --activation-arg "https://google.com" \
  --banner "${pwd}/banner.jpg"
