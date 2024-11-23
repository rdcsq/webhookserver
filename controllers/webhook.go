package controllers

import (
	"net/http"
	"os/exec"
	"syscall"
	"time"
	"webhookserver/structs"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(400)
		w.Write([]byte("id is required"))
		return
	}

	var config *structs.WSConfig
	for _, cfg := range structs.Config {
		if cfg.Name == id {
			config = &cfg
		}
	}
	if config == nil {
		w.WriteHeader(400)
		w.Write([]byte("id not found"))
		return
	}

	cmd := exec.Command(config.Command, config.Args...)
	cmd.Env = append(cmd.Env, config.Environment...)
	cmd.Dir = config.WorkingDirectory
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	var (
		waitTimer  *time.Timer
		waitTimerC <-chan time.Time
	)

	if config.Timeout > 0 {
		waitTimer = time.NewTimer(time.Duration(config.Timeout) * time.Second)
		waitTimerC = waitTimer.C
	}

	cmdDone := make(chan cmdResult, 1)

	go func() {
		outb, err := cmd.CombinedOutput()
		cmdDone <- cmdResult{string(outb), err}
	}()

	select {
	case <-waitTimerC:
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		w.WriteHeader(408)
		w.Write([]byte("command timed out"))
	case res := <-cmdDone:
		if res.err != nil {
			w.WriteHeader(500)
			w.Write([]byte(res.err.Error()))
			return
		}
		w.Write([]byte(res.out))
	}
}

type cmdResult struct {
	out string
	err error
}
