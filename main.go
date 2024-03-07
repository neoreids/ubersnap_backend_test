package main

import "ubersnap/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}