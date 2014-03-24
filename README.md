go-clock-unity
==============

Warning, this is not a great example of interaction with C or GTK code, but it works for me (tm).

In particular, this code calls gtk functions outside the main event loop in away that might not be
threadsafe. Regular gtk timeouts would be better. It might leak memory or crash.

A low memory replacement for the ubuntu evolution based clock which took up way too much memory on
my system to show the time.
![screenshot](https://raw.githubusercontent.com/shanemhansen/go-clock-unity/screenshots/goclock.png)

