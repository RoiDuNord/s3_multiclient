package main

import "s3_multiclient/app"

func main() {
	if err := app.Run(); err != nil {
		return
	}
}
