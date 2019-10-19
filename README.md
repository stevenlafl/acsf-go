# About
ACSF Go is a command-line tool designed to interface with Acquia Site Factory

# Getting Started

Edit the `api/config.go` file to reflect your ACSF credentials and environment URLs:
    
	...
	    USE_PROXY  =  false
	    PROXY_ADDR  =  "127.0.0.1:10000"
	    API_USER  =  ""
	    API_PASS  =  ""
	...
	    APIUrls["dev"] =  "https://www.dev-example.acsitefactory.com"
	    APIUrls["test"] =  "https://www.test-example.acsitefactory.com"
	    APIUrls["live"] =  "https://www.example.acsitefactory.com"
	...

> **Note:** `PROXY_ADDR` is the address for a SOCKS5 proxy (e.g. `ssh -D` tunnel)

Build is simple as always:

	$ cd acsf-go
	$ go build

# Usage
	$ ./acsf-go
	Usage: ./acsf-go [OPTION]... COMMAND [ARGUMENT]...
	Options: 
	  --environment=[dev*,test,prod]
	    Select an environment to act on
	Available commands:
	  stacks
	    Show available stacks for the environment
	  vcs [STACK ID]
	    Show available vcs refs and currently deployed vcs ref for the environment
	  updates
	    Show a list of updates for the environment, previous and currently running
	  status [UPDATE ID]
	    Show update status of a specific update for the environment
