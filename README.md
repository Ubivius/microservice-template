# microservice-template
Template for microservices.

Beginning of code was based off of Nic Jackson's building microservices in Go : https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_

## TODO
- Refactor to have a game related item instead of random product example
- Tests for http handlers
- Make validate take an interface instead of a specific struct
- Improve setup documentation
- Automated tests for non handler functions
- Only current options to interact with API is with curl
- Move routing to another file? Or keep in main.go
- Check gorilla/mux docs for how to fix warning in main.go that Nic Jackson ignores
