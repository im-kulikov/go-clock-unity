package main

// #cgo pkg-config: gtk+-2.0
// #cgo CFLAGS: -I/usr/include/libappindicator-0.1
// #cgo LDFLAGS: -lappindicator
// #include <gtk/gtk.h>
// #include <libappindicator/app-indicator.h>
// #include <stdlib.h>
// #include "clock.go.h"
import "C"
import (
	"time"
)

func updateclock(indicator *C.AppIndicator, t time.Time) {
	msg := GString(t.Format("3:04:05 PM"))
	C.app_indicator_set_label(indicator, msg, msg)
}
func GString(data string) *C.gchar {
	ptr := C.CString(data)
	msg := make([]C.gchar, len(data))
	C.g_stpcpy(&msg[0], ptr)
	return &msg[0]
}
func GtkMainLoop() chan bool {
	q := make(chan bool)
	go func() {
		C.gtk_main()
		q <- true
	}()
	return q
}
func main() {
	C.gtk_init(nil, nil)
	indicator := C.app_indicator_new(
		GString("simple-clock-client"),
		GString("clock"),
		C.APP_INDICATOR_CATEGORY_APPLICATION_STATUS)
	C.app_indicator_set_status(indicator, C.APP_INDICATOR_STATUS_ACTIVE)
	updateclock(indicator, time.Now())
	menu := C.gtk_menu_new()
	C.app_indicator_set_menu(indicator, C.toGMenu(menu))
	ticker := time.Tick(time.Second)
	quit := GtkMainLoop()
	for {
		select {
		case t := <-ticker:
			updateclock(indicator, t)
		case <-quit:
			break
		}
	}
}
