#!/bin/bash

export PATH=$PATH:/sbin:/usr/sbin:/usr/local/bin:/usr/bin:/bin:/usr/local/sbin

CURFILE=$(readlink -f "$0")

CURDIR=$(dirname $CURFILE)

MODULE_DIR=${CURDIR%"/tools/op"}

MODULE=$(basename $MODULE_DIR)


cd $CURDIR

# stop may take minutes, so touch flag first to avoid auto restart during stop process
touch ${CURDIR}/.disable-autorestart

./stop.sh

if [[ "$?" -eq 0 ]];then
	exit 0
else
        # remove flag if stop failed
        rm ${CURDIR}/.disable-autorestart

	echo "stop proc failed!"
	exit 1
fi
