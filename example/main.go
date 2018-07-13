package main

import "github.com/makki0205/ses"

func main() {
	ses := ses.NewSES("AWS_ACCESS_KEY_ID", "AWS_SECRET_KEY", "us-west-2")
	err := ses.Send("hoge@hoge.com", "hoge@domain.me", "sub test", "body test")
	if err != nil {
		panic(err)
	}
}
