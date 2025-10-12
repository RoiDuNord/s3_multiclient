package main

import "s3_multiclient/app"

func main() {
	if err := app.MustRun(); err != nil {
		return
	}
}
