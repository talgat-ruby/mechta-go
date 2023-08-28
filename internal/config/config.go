package config

import "flag"

type Config struct {
	Max  int
	File string
}

func NewConfig() Config {
	goroutines := flag.Int("goroutines", 0, "number of goroutines")
	file := flag.String("file", "", "json file to load")
	flag.Parse()

	return Config{
		Max:  *goroutines,
		File: *file,
	}
}
