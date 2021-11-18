package cli

import (
	"bufio"
	"os"
	"strings"
)

func ConsoleInput(out *chan Message) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		// convert CRLF to LF (WINDOWS)
		input = strings.Replace(input, "\r\n", "", -1)
		*out <- Message{Sender: "cli", Arguments: strings.Fields(input)}
	}
}
