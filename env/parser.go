package env

import (
	"bufio"
	"io"
	"strings"
)

func ParseFile(r io.Reader) map[string]string {
	env := make(map[string]string)

	s := bufio.NewScanner(r)
	
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	err := s.Err()
	if err != nil {
		panic(err)
	}

	for _, text := range lines {
		text = strings.TrimSpace(text)

		if (len(text) < 3) {
			continue
		}

		i := strings.Index(text, "=")

		if (i == -1) {
			continue
		}

		k := text[:i]
		v := text[i+1:]

		env[k] = v
	}

	return env
}