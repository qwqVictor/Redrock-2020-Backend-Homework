#!/bin/bash

USERNAME=$(cat /dev/urandom | head -c 8 | base64 -w 0)
PASSWORD=$(cat /dev/urandom | head -c 11 | base64 -w 0)
PASSHASH=$(printf ${PASSWORD} | ( (md5sum | cut -d ' ' -f 1) || md5))
DB_FILENAME=$(mktemp)

if [ ! -f "main" ]; then
    make || exit 1
fi

printf "%s\n%s\n" ${USERNAME} ${PASSHASH} > ${DB_FILENAME}

printf "%s\n%s\n" ${USERNAME} ${PASSWORD} | ./main ${DB_FILENAME} 2> /dev/null

rm ${DB_FILENAME}