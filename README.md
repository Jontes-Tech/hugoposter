# HugoPoster
## Let's break the files down!
* bash.sh - Bash script that runs at commit
* go.mod - Golang module stuff
* go.sum - More Golang module stuff
* main.go - Get Git commit url, then web scrape and run Python script
* README.md - This file
* twitter.py - Tweets with the data gathered in main.go
## Running it yourself
How do you use it? This isn't optimized for other's use, sorry, but I provide the steps below. This assumes the "Terminal" theme in Hugo.
* Change paths in bash.sh, and change the docker command to your web server restart command.
* Change paths and urls in main.go
* Rename twitter-sample.py to twitter.py
* Generate a Twitter API key and set it in twitter.py
If there is demand, create a Github issue about it and I will make it easy to selfhost.