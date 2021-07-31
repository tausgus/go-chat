# go-chat

go-chat is a small server written in Go that behaves like a messaging service.
It can be used with any TCP client like `netcat`,

## Usage
- Simply run `/path/to/go-chat portnum`, where `portnum` is the port that should be used for communications.
- Connecting via netcat: `netcat ip_address port`
- The server will prompt you for a username which will be displayed to other clients.
Using `rlwrap` will prevent messages others send from overlapping what you are currently writing:
`rlwrap netcat ip_address port`

## Features
- Usernames visible to other clients
- Clients do not get their own messages sent back to them
- Very simple and minimalistic, the built binary is under 5 MB.
