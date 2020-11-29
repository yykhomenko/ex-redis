package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/mediocregopher/radix"
)

func main() {
	c, err := radix.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	if err := test(c); err != nil {
		log.Fatal(err)
	}
}

func test(client radix.Conn) error {

	key := "k"
	val := "v"

	if err := client.Do(radix.Cmd(nil, "SET", key, val)); err != nil {
		return err
	}

	var out string
	if err := client.Do(radix.Cmd(&out, "GET", key)); err != nil {
		return err
	} else if out != val {
		return errors.New("got wrong value")
	}

	fmt.Println(out)

	return nil
}
