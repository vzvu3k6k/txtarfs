package main

import (
	"net/http"

	"github.com/vzvu3k6k/txtarfs"
)

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.FS(tfs)))
}

var tfs = txtarfs.Parse([]byte(
	`-- page.html --
<script src="/main.js"></script>
-- main.js --
setInterval(() => { document.body.textContent = new Date() }, 1000)`))
