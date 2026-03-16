package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func repoGet(id int) error {
	// TODO: 返回 ErrNotFound
	return nil
}

func serviceGet(id int) error {
	// TODO: 包装 repoGet
	return nil
}

func handler(id int) error {
	// TODO: 包装 serviceGet
	return nil
}

func main() {
	err := handler(1)
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrNotFound))
}
