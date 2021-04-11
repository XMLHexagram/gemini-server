package debug

import (
	"fmt"
	"os"
	"strings"
)

func Print(format string, values ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(os.Stdout, "[Gemini-debug] "+format, values...)

}

func PrintRoute(path string) {
	Print("%-6s \n", path)
}

func PrintLoadFile(path, absolutePath string) {
	Print("%-6s %-6s", path, absolutePath)
}

func PrintLoadDir(path, absolutePath string) {
	Print("%-6s %-6s", path, absolutePath)
}

func PrintProxy(path, url string) {
	Print("%-6s %-6s", path, url)
}

func PrintError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stdout, "[Gemini-debug] [ERROR] %v\n", err)
	}
}
