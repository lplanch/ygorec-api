#!/bin/bash

# turn on bash's job control
set -m

# Start the server and put it in the background
./main &

# Start the helper process
make generate-views

# the make generate-views might need to know how to wait on the
# primary process to start before it does its work and returns

# now we bring the primary process back into the foreground
# and leave it there
fg %1
