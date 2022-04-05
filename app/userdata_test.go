package ChatServer

import (
	"bufio"
	"bytes"
	"net"
	"testing"
)

func TestPromptQuestion(t *testing.T) {
	writeBuff := bytes.NewBuffer([]byte{})
	readBuff := bytes.NewBuffer([]byte{})

	readWriter := bufio.NewReadWriter(
		bufio.NewReader(readBuff),
		bufio.NewWriter(writeBuff),
	)

	t.Run("PromptQuestion", func(t *testing.T) {
		// write what ?
		PromptQuestion("What is your username?", readWriter)
		// read what ?

		//if got := PromptQuestion("What is your username?", readWriter); got != "willem" {
		//	t.Errorf("PromptQuestion() = %v, want %v", got)
		//}
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
