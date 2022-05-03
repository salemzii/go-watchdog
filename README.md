# go-watchdog
Go-watchdog is a web application observability tool built for Go, that exposes a status endpoint for application services like databases, caches, message-brokers, mails and even storages.


Go-watchdog us meant to be a replica of django-watchman, but for golang; Go-watchdog is pretty much loosely-coupled thereby allowing developers the ease to customize it to their use cases.

The base schema Go-watchdog operates on is configure in the watchdogConfig struct, this struct contains details on each service to be monitored and is  to be initiated in your project's main.go/application entry file.
