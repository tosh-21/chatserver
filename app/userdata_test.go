package ChatServer

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net"
	"testing"
	"time"
)

func TestPromptQuestion(t *testing.T) {
	client, server := net.Pipe()
	t.Run("PromptQuestion", func(t *testing.T) {
		question := "What is your username!?"
		go func() {
			buf := make([]byte, len(question))
			client.Read(buf)

			assert.Equal(t, string(buf), question)
			written, err := client.Write([]byte("willy \n"))
			log.Println(written, err)
		}()
		actual := PromptQuestion(question, server)
		assert.Equal(t, "willy", actual)
		//if got := PromptQuestion("What is your username?", readWriter); got != "willem" {
		//	t.Errorf("PromptQuestion() = %v, want %v", got)
		//}
	})

	t.Run("PromptQuestion Test Error", func(t *testing.T) {
		// Implement error test
		time.Sleep(500 * time.Millisecond)
		client.Close()
		question := "What is your username!?"
		go func() {
			buf := make([]byte, len(question))
			client.Read(buf)

			assert.Equal(t, string(buf), question)
			written, err := client.Write([]byte("willy \n"))
			log.Println(written, err)
		}()
		actual := PromptQuestion(question, server)
		assert.Equal(t, "Error while reading input: EOF", actual)

	})
}
