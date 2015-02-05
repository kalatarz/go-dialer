// Copyright 2013 Alexandre Fiori
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Event Socket client that connects to FreeSWITCH to originate a new call.
package main

import (
	"fmt"
	"log"
//	"time"

	"github.com/fiorix/go-eventsocket/eventsocket"
)

const dest = "sofia/external/*010600@ekiga.net "
//const dialplan = "&socket(192.168.1.5:9090 async)"

func main() {
	var call = ""
	c, err := eventsocket.Dial("localhost:8021", "ClueCon")
	if err != nil {
		log.Fatal(err)
	}
	c.Send("events json ALL")
	c.Send(fmt.Sprintf("bgapi originate %s 9196 xml default", dest))
	for {
		ev, err := c.ReadEvent()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nNew event")
		ev.PrettyPrint()
		if ev.Get("Answer-State") == "hangup" {
		log.Fatal("Not working")
			break
		}
		if ev.Get("Answer-State") == "answered" {
		fmt.Println("We got the call: ", ev.Get("Unique-Id"))
		call = ev.Get("Unique-Id")
		}
		if ev.Get("Session-Count") >= "1" {
		fmt.Println("Call in progress: ", ev.Get("Session-Count"))
		//time.Sleep(3000 * time.Millisecond)
		fmt.Println("Bye bye: ", call)
		c.SendMsg(eventsocket.MSG{"call-command": "hangup", "hangup-cause": "we're done!", }, call, "")
			break
		}
	}
	fmt.Println("Closing socker")
	c.Close()
}
