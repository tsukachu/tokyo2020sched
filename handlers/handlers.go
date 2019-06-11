package handlers

import "gopkg.in/gorp.v1"

type Handler struct {
	DbMap *gorp.DbMap
}
