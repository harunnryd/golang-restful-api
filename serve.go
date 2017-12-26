package main

import (
  "golang-restful-api/routing"
)

func main() {
	serve := routing.WebService {}
	serve.Run()
}
