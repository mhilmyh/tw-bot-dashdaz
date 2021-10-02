package env

import "os"

func Load() {
	f, err := os.Open(".env")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	env := ParseFile(f)

	for k, v := range env {
		os.Setenv(k, v)
	}
}