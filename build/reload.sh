#!/usr/bin/bash

test -x $0 || chmod +x $0


program_path=`pwd`
app_name=${program_path##*/}

#bucket_name=`echo ${app_name} | sed 's/_//g'`
bucket_name="tflow"

cd ${program_path}/tools/op/ && ./stop.sh
rclone --config /etc/rclone.conf copy minio:joybuild/${bucket_name}/build/${app_name} ${program_path}/bin
rclone --config /etc/rclone.conf sync minio:joybuild/${bucket_name}/tools ${program_path}/tools


cd ${program_path}
find ./tools  -type f -name "*.sh" | xargs dos2unix > /dev/null 2>&1
find ./tools  -type f -name "*.sh" | xargs -i chmod +x {} > /dev/null 2>&1
test -d log || mkdir log

chmod +x ${program_path}/bin/${app_name}
cd ${program_path}/tools/op/ && ./start.sh
md5sum ${program_path}/bin/${app_name}
sleep 3
ps -ef | grep ${app_name} | grep conf | grep -v grep
