// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/kyokomi/emoji"
)

var (
	ecm             = emoji.CodeMap()
	currentEmoji    = []byte(ecm[":smile:"])
	validEmojiNames []string
	hub             = newHub()
)

func init() {
	for name := range ecm {
		validEmojiNames = append(validEmojiNames, name)
	}
}

func shuffleEmoji(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	newEmojiIndex := rand.Int63n(int64(len(validEmojiNames)))
	currentEmoji = []byte(ecm[validEmojiNames[newEmojiIndex]])
	hub.broadcast <- currentEmoji

	w.WriteHeader(http.StatusOK)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	go hub.run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/shuffle", shuffleEmoji)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
