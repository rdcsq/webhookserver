package utils

import (
	"encoding/json"
	"net/http"
)

func JsonError(w http.ResponseWriter, message string) {
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(map[string]string{"code": "other", "error": message})
}

func JsonExecInitError(w http.ResponseWriter, message string) {
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(map[string]string{"code": "exec_init", "error": message})
}

func JsonExecError(w http.ResponseWriter, message, out string) {
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(map[string]string{
		"code":  "exec",
		"error": message,
		"out":   out,
	})
}

func JsonExecSuccess(w http.ResponseWriter, stdout string) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"stdout": stdout})
}
