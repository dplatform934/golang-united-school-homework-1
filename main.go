package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	emoji := emoji.Sprint("Hello:world_map:!")
	fmt.Println(emoji)
}
