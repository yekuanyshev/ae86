package banner

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func Default(template string, data map[string]interface{}) {
	err := Show(os.Stdout, strings.NewReader(template), data)
	if err != nil {
		log.Println(err)
	}
}

// Show - load the banner and prints it to output
func Show(out io.Writer, in io.Reader, data map[string]interface{}) error {
	if in == nil {
		return fmt.Errorf("the input is nil")
	}

	banner, err := io.ReadAll(in)
	if err != nil {
		return fmt.Errorf("error trying to read the banner, err: %v", err)
	}

	t, err := template.New("banner").Parse(string(banner))
	if err != nil {
		return fmt.Errorf("error trying to parse the banner file, err: %v", err)
	}

	data["ansiColor"] = ansiColor{true}
	data["ansiBackground"] = ansiBackground{true}

	err = t.Execute(out, data)
	if err != nil {
		return fmt.Errorf("error trying to execute template: %v", err)
	}

	return nil
}
