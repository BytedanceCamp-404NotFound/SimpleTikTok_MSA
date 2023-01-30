#! /bin/bash


# goctl_cmd1="go build -o $PROTO_OUTPUT -v -p 2 $IS_REBUILD "
# for f in $(find . -name "*.go" -maxdepth 1 -exec basename {} \;)
# do
#     cmd=$goctl_cmd1$f
#     echo $cmd
#     eval $cmd
#     echo $f
# done
# cp -r etc $PROTO_OUTPUT

filename="$(date +%Y%m%d)_$(date +%H%M%S)_baseinterface"
echo $filename

