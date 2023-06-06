package utils

import "log"

func Handle(e error) {
	if e != nil {
		log.Panic(e)
	}
}
