package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"sync"
	"time"
)

func omikuji() string {
	castNum := rand.Intn(6) + 1
	var result string
	switch castNum {
	case 1:
		result = "凶"
	case 2:
		result = "吉"
	case 3:
		result = "吉"
	case 4:
		result = "中吉"
	case 5:
		result = "中吉"
	case 6:
		result = "大吉"
	}
	return fmt.Sprintf("おみくじ結果：%v,さいのめ：%v\n", result, castNum)
}

func main() {
	http.Handle("/", &templateHandler{filename: "omikuji.html"})
	// webサーバー開始
	rand.Seed(time.Now().UnixNano())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates",
			t.filename)))
	})

	data := map[string]interface{}{
		"mon": omikuji(),
		"tue": omikuji(),
		"wed": omikuji(),
		"thu": omikuji(),
		"fri": omikuji(),
		"sat": omikuji(),
		"sun": omikuji(),
	}

	//log.Println(r.Host, r.Header)
	if err := t.templ.Execute(w, data); err != nil {
		log.Fatal("ServeHTTP:", err)
	}
}
