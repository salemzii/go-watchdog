# go-watchdog
Go-watchdog is a simple web application monitoring tool for Go, that exposes an endpoint for application services like databases, caches, message-brokers and even storages.


Go-watchdog us meant to be a replica of django-watchman, but for golang; Go-watchdog is pretty much loosely-coupled thereby allowing developers the ease to customize it to their use cases.

The base schema Go-watchdog operates on is configure in the watchdogConfig.json file, this file contains details on each service to be monitored and is expected to be in the same directory as your project's main.go/application entry file.
