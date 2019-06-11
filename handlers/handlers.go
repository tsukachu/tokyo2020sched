package handlers

import "gopkg.in/gorp.v2"

type Handler struct {
	DbMap *gorp.DbMap
}
