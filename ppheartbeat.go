// Copyright 2019 PiaoYun/P.Y.G. All rights reserved.
// Use of this source code is governed by a BSD-style
// https://www.chinapyg.com

package client

import (
	"log"
	"time"
)

type PPHeartbeat struct {
	duration time.Duration
	callback func()
	done     chan struct{}
}

func NewHeartbeat(d time.Duration, cb func()) *PPHeartbeat {
	return &PPHeartbeat{d, cb, make(chan struct{}, 1)}
}

func (h *PPHeartbeat) Start() {
	log.Println("=====call Start()=====")
	go func() {
		ticker := time.NewTicker(h.duration)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				h.callback()
			case <-h.done:
				return
			}
		}
	}()
}

func (h *PPHeartbeat) Stop() {
	log.Println("=====call pingStop()=====")
	if h.done != nil {
		close(h.done)
	}
}

func init() {
	log.Println("=====go PPHeartbeat by PiaoYun/P.Y.G=====")
}