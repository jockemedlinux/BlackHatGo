package main
import (
"fmt"
"net/http"
"github.com/stacktitan/ldapauth"
"golang.org/x/lint/golint"
)

// go get from book is deprecated. I'll use the go install instead wich places for instance github.com/stacktitan/ldapauth in my $GOROOT-path/github.com/stacktitan.
// also the go get -u golang.org/x/lint/golint won't work. work with go install instead. (get install golang.org/x/lint/golint@latest)
