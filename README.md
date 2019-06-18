# golang-article-app

## Setup
 - install govendor package
 - install air or gin for live reload
 - clone project to github subfolder under src
 - run `govendor fetch`
 
## Development
To start the application, use any of the options below:
 - run `go run *.go` 
 - if air is installed, run `~/.air` in project root / run `air` if alias is set in .bashrc file
 - for gin, run `gin`

## Testing
run `go test -v ./tests`