#! /bin/bash

OutType=$1
ROOT=$PWD
BIN_FOLDER=$ROOT/bin
API_OUTPUT=$ROOT/bin/external
PROTO_OUTPUT=$ROOT/bin/internal
ETC_OUTPUT=$ROOT/bin/etc
etc_output_api=$ETC_OUTPUT/external
etc_output_proto=$ETC_OUTPUT/internal
LOG_FOLDER=$ROOT/bin/logs/


if [ ! -d $BIN_FOLDER ]; then
    mkdir -p $API_OUTPUT
    mkdir -p $PROTO_OUTPUT
    mkdir -p $etc_output_api
    mkdir -p $etc_output_proto
    mkdir -p $LOG_FOLDER
fi

create_api(){
    cd external
    goctl api go -api api/baseinterface.api -dir baseinterface/ -style gozero
    goctl api go -api api/commaction.api -dir commaction/ -style gozero
    goctl api go -api api/relationfollow.api -dir relationfollow/ -style gozero
    cd -
}
create_proto(){
    cd internal
    goctl_cmd1="goctl rpc protoc "
    for f in $(find . -name "*.proto" -exec basename {} \;)
    do
        ft=$(basename $f .proto)
        proto_file="proto/$f"
        goctl_cmd2=" --go_out=MicroServices/$ft/pkg --go-grpc_out=MicroServices/$ft/pkg --zrpc_out=MicroServices/$ft"
        cmd=$goctl_cmd1$proto_file$goctl_cmd2
        echo $cmd
        eval $cmd
        # echo $f
    done
    cd -
}
if [[ $OutType == "create" ]];then 
    case $2 in
    "api") 
        create_api
    ;;
    "proto") 
        create_proto
    ;;
    *) 
        create_api 
        create_proto
        echo "useage: ./GoZeroUse create api"
        echo "        ./GoZeroUse create proto"
    ;;
    esac
fi

# -v	编译时显示包名
# -p n	开启并发编译，默认情况下该值为 CPU 逻辑核数
# -a	强制重新构建
# -n	打印编译时会用到的所有命令，但不真正执行
# -x	打印编译时会用到的所有命令
build_api(){
    is_build=$1
    if [ ! -d $API_OUTPUT ]; then
        mkdir -p $API_OUTPUT
    fi
    if [ ! -d $etc_output_api ]; then
        mkdir -p $etc_output_api
    fi
    cd external
    go build -o $API_OUTPUT -v -p 2 $is_build baseinterface/baseinterface.go  
    go build -o $API_OUTPUT -v -p 2 $is_build commaction/commactioninterface.go  
    go build -o $API_OUTPUT -v -p 2 $is_build relationfollow/relationfollowinterface.go  
    cp baseinterface/etc/* $etc_output_api
    cp commaction/etc/* $etc_output_api
    cp relationfollow/etc/* $etc_output_api
    cd -
}
build_proto(){
    is_build=$1
    if [ ! -d $PROTO_OUTPUT ]; then
        mkdir -p $PROTO_OUTPUT
    fi
    if [ ! -d $etc_output_proto ]; then
        mkdir -p $etc_output_proto
    fi
    cd internal/MicroServices
    # goctl_cmd1="go build -o $PROTO_OUTPUT -v -p 2 $IS_REBUILD "
    # for f in $(find . -name "*.go" -maxdepth 1 -exec basename {} \;)
    # do
    #     cmd=$goctl_cmd1$f
    #     echo $cmd
    #     eval $cmd
    #     echo $f
    # done
    # cp -r etc $PROTO_OUTPUT
    goctl_cmd1="go build -o $PROTO_OUTPUT -v -p 2 $is_build "
    for f in $(find . -name "*.go" -maxdepth 2)
    do 
        cmd=$goctl_cmd1$f
        eval $cmd
        echo $cmd
    done
    # go build -o $PROTO_OUTPUT -v -p 2 ./utilserver/utilserver.go
    for f in $(find . -name "*.yaml")
    do 
        cp $f $etc_output_proto
        echo $f
    done
    # cp minio/etc/* $etc_output_proto
    cd -
}
if [[ $OutType == "build" ]];then
    case $2 in
    "api") 
        build_api
    ;;
    "proto") 
        build_proto
    ;;
    *) 
        build_api 
        build_proto
        echo "useage: ./GoZeroUse build api"
        echo "        ./GoZeroUse build api -a"
    ;;
    esac
fi



run_logs_create(){
    datatime="$(date +%Y%m%d)_$(date +%H%M%S)/"
    if [[ ! -d datatime ]];then
        mkdir -p $LOG_FOLDER$datatime
    fi
    LOG_FOLDER=$LOG_FOLDER$datatime
}
# 需要使用go build生成的exe文件来执行，这样os.Executable()获取到的才是正确的路径
# go run来运行，会将可执行文件默认放到/tmp/go-build...
# 需要配置GOTMPDIR=""来改变go run生成可执行文件的位置
# go run Baseinterface-Api.go -f etc/BaseInterface-Api.yaml
run_api(){
    build_api
    cd $API_OUTPUT
    if [[ $1 == "-c" ]];then
        case $2 in
        "base")
            filename="$(date +%Y%m%d)_$(date +%H%M%S)_baseinterface.log"
            ./baseinterface -f $etc_output_api/baseinterface.yaml > $LOG_FOLDER$filename 2>&1 
        ;;
        "com")
            ./commactioninterface -f $etc_output_api/commactioninterface.yaml
        ;;
        "rela")
            ./relationfollowinterface -f $etc_output_api/relationfollowinterface.yaml
        ;;
        *)
            echo "useage: ./GoZeroUse run api -c"
            echo "        ./GoZeroUse run api -c [base,com,rela]"
            echo "        ./GoZeroUse run api # 全部运行"
        esac
    else
        for f in $(find .  -type f -exec basename {} \;)
        do
            filename="$(date +%Y%m%d)_$(date +%H%M%S)_$f.log"
            cmd="./$f -f $etc_output_api/$f.yaml > $LOG_FOLDER$filename 2>&1 &"
            echo $cmd
            eval $cmd
        done
    fi
    # ./baseinterface -f $etc_output_api/baseinterface.yaml  &
    # ./commactioninterface -f $etc_output_api/commactioninterface.yaml  &
    # ./relationfollowinterface -f $etc_output_api/relationfollowinterface.yaml  &
    cd -
}
run_proto(){
    build_proto
    filename="$(date +%Y%m%d)_$(date +%H%M%S)_etcd.log"
    etcd > $LOG_FOLDER$filename 2>&1 &
    cd $PROTO_OUTPUT
    for f in $(find .  -type f -exec basename {} \;)
    do
        filename="$(date +%Y%m%d)_$(date +%H%M%S)_$f.log"
        cmd="./$f -f $etc_output_proto/$f.yaml > $LOG_FOLDER$filename 2>&1  &"
        echo $cmd
        eval $cmd
    done
    cd -
}
if [[ $OutType == "run" ]];then 
    cp -r config bin  # 复制配置文件
    case $2 in
    "api" )
        run_logs_create # 生成日志文件夹
        run_api $3 $4
    ;;
    "proto" )
        run_logs_create # 生成日志文件夹
        run_proto
    ;;
    "all")
        run_logs_create # 生成日志文件夹
        run_api
        run_proto
    ;;
    *)
        echo "useage: ./GoZeroUse run  "
        echo "        ./GoZeroUse run all"
        echo "        ./GoZeroUse run api "
        echo "        ./GoZeroUse run api -c base"
    ;;
    esac
fi



# pid=$(ps -ef | grep "./baseinterface" | grep -v grep | awk '{print $2}')
# kill -f $pid
kill_api(){
    cd $API_OUTPUT
    for f in $(find . ! -name "*.yaml" -type f -exec basename {} \;)
    do
        echo $f
        pkill -f $f 
    done
    cd -
}
kill_proto(){
    cd $PROTO_OUTPUT
    for f in $(find . ! -name "*.yaml" -type f -exec basename {} \;)
    do
        echo $f
        pkill -f $f 
    done
    cd -
}
if [[ $OutType == "kill" ]];then 
    # ps -ef | grep "baseinterface" | grep -v grep # 显示进程详细信息
    # pkill -f baseinterface # 通过进程名字杀死进程
    # pkill -f commactioninterface 
    # pkill -f relationfollowinterface 
    case $2 in
    "api")
        kill_api
        ps -f
    ;;
    "proto")
        kill_proto
        ps -f
    ;;
    "etcd")
        pkill -f etcd
        ps -f
    ;;
    "all")
        pkill -f etcd
        kill_api
        kill_proto
        ps -f
    ;;
    *)
        echo "useage: ./GoZeroUse kill "
        echo "        ./GoZeroUse kill api"
        echo "        ./GoZeroUse kill all"
    ;;
    esac
fi



if [[ $OutType == "clear" ]];then 
    rm -rf $LOG_FOLDER/*
fi


