package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

const index = `
<!doctype html>
<html>
  <head>
    <title>SSE Example</title>
  </head>
  <body>
    <div id="sse-data"></div>

    <script>
      const eventSource = new EventSource("{{ .ServerHOST }}/events");
      eventSource.onmessage = function (event) {
        const dataElement = document.getElementById("sse-data");
        dataElement.innerHTML += event.data + "<br>";
      };
    </script>
  </body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htmlTemp, err := template.New("inddx.html").Parse(index)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		serverHost := os.Getenv("SERVER_HOST")
		if serverHost == "" {
			serverHost = "http://localhost:8080"
		}
		if err := htmlTemp.Execute(w, struct {
			ServerHOST string
		}{
			ServerHOST: serverHost,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	slog.Info("Start server on 8000")
	_ = http.ListenAndServe(":8000", nil)
}
