# Translations

Internationalization for Go projects.

## Generating new labels

The gotext (`go install golang.org/x/text/cmd/gotext@latest`) utility must be installed first.
* Running `go generate ./internal/translations/translations.go` will generate the files. It scans for calls to `Printf()`.
* The files `out.gotext.json` are meant to be sent to translators. These will be received as `messages.gotext.json`.
* Once the translations have been received, the generation must be executed again `go generate ./internal/translations/translations.go`.
