package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", eventsHandler)
	slog.Info("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i))
		time.Sleep(2 * time.Second)
		w.(http.Flusher).Flush()
	}

	closeNotify := w.(http.CloseNotifier).CloseNotify()
	<-closeNotify
	slog.Info("Connection closed")
}
