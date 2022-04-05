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

func TestChatServer_CheckUserName(t *testing.T) {
	type fields struct {
		UserConnections []net.Conn
		Users           map[string]User
	}
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chat := &ChatServer{
				UserConnections: tt.fields.UserConnections,
				Users:           tt.fields.Users,
			}
			got, got1 := chat.CheckUserName(tt.args.conn)
			if got != tt.want {
				t.Errorf("CheckUserName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckUserName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
