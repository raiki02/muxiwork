package savePaper

import (
	"fmt"

	"github.com/spf13/afero"
)

func SavePaper(path string, text string) {
	fs := afero.NewOsFs()
	file, err := fs.Create(path)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}
	fmt.Println("File saved successfully at", path)
}
