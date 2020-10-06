#!/bin/bash
export PORT=20208

if [ -f main ]; then
    make || exit 1
fi

./main &
curl localhost:${PORT} || wget -O - localhost:${PORT}

kill $!
