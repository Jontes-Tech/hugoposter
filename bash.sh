#!/bin/bash
cd /home/jonte/website/blog/
git fetch
cd /home/jonte/hugotweeter/
./main
cd /home/jonte/website/blog/
git pull
hugo
docker restart blog