#! /bin/bash

OutType=$1
ROOT=$PWD
API_OUTPUT=$ROOT/bin/external
PROTO_OUTPUT=$ROOT/bin/internal
ETC_OUTPUT=$ROOT/bin/etc
etc_output_api=$ETC_OUTPUT/external
etc_output_proto=$ETC_OUTPUT/internal

if [ ! -d $ETC_OUTPUT ]; then
    mkdir -p $ETC_OUTPUT
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
        goctl_cmd2=" --go_out=MicroServices/pkg --go-grpc_out=MicroServices/pkg --zrpc_out=MicroServices/$ft"
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
    go build -o $PROTO_OUTPUT -v -p 2 $is_build minio/minio.go  
    go build -o $PROTO_OUTPUT -v -p 2 $is_build mysql/mysql.go  
    cp minio/etc/* $etc_output_proto
    cp mysql/etc/* $etc_output_proto
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

# 需要使用go build生成的exe文件来执行，这样os.Executable()获取到的才是正确的路径
# go run来运行，会将可执行文件默认放到/tmp/go-build...
# 需要配置GOTMPDIR=""来改变go run生成可执行文件的位置
# go run Baseinterface-Api.go -f etc/BaseInterface-Api.yaml

run_api(){
    build_api
    cd $API_OUTPUT
    for f in $(find .  -type f -exec basename {} \;)
    do
        cmd="./$f -f $etc_output_api/$f.yaml  &"
        echo $cmd
        eval $cmd
    done
    # ./baseinterface -f $etc_output_api/baseinterface.yaml  &
    # ./commactioninterface -f $etc_output_api/commactioninterface.yaml  &
    # ./relationfollowinterface -f $etc_output_api/relationfollowinterface.yaml  &
    cd -
}
run_proto(){
    build_proto
    cd $PROTO_OUTPUT
    for f in $(find .  -type f -exec basename {} \;)
    do
        cmd="./$f -f $etc_output_proto/$f.yaml  &"
        echo $cmd
        eval $cmd
    done
    cd -
}
if [[ $OutType == "run" ]];then 
    case $2 in
    "api" )
        run_api
    ;;
    "proto" )
        etcd &
        run_proto
    ;;
    *)
        run_api
        run_proto
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


