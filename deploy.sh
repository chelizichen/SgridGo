#!/usr/bin/env bash
# Written in [Amber](https://amber-lang.com/)
# version: 0.3.5-alpha
# date: 2025-04-02 16:18:48


exit__80_v0() {
    local code=$1
    exit "${code}";
    __AS=$?
}
__0_server_name="SgridGoServer"
build_to_server__94_v0() {
    echo "Building to server"
     rm -r ./app ;
    __AS=$?;
if [ $__AS != 0 ]; then
        echo "app is not exist, ignore"
fi
     GOOS=linux GOARCH=amd64  go build -o ${__0_server_name} ;
    __AS=$?;
if [ $__AS != 0 ]; then
        echo "Failed to build server"
        exit__80_v0 1;
        __AF_exit80_v0__12_9="$__AF_exit80_v0";
        echo "$__AF_exit80_v0__12_9" > /dev/null 2>&1
fi
}
tar__95_v0() {
    echo "Tar to dist"
     rm -r ./${__0_server_name}.tar.gz  ;
    __AS=$?;
if [ $__AS != 0 ]; then
        echo "${__0_server_name}.tar.gz is not exist, ignore"
fi
     tar -cvf ${__0_server_name}.tar.gz ./${__0_server_name} ;
    __AS=$?;
if [ $__AS != 0 ]; then
        echo "Failed to tar"
        exit__80_v0 1;
        __AF_exit80_v0__24_9="$__AF_exit80_v0";
        echo "$__AF_exit80_v0__24_9" > /dev/null 2>&1
fi
     rm -r ./${__0_server_name}  ;
    __AS=$?;
if [ $__AS != 0 ]; then
        echo "${__0_server_name} is not exist, ignore"
fi
    echo "Tar Success"
}
clean__96_v0() {
    echo "Clean"
     rm -r ./app ;
    __AS=$?;
if [ $__AS != 0 ]; then
        echo "app is not exist, ignore"
fi
}
build_to_server__94_v0 ;
__AF_build_to_server94_v0__40_1="$__AF_build_to_server94_v0";
echo "$__AF_build_to_server94_v0__40_1" > /dev/null 2>&1
tar__95_v0 ;
__AF_tar95_v0__42_1="$__AF_tar95_v0";
echo "$__AF_tar95_v0__42_1" > /dev/null 2>&1
clean__96_v0 ;
__AF_clean96_v0__44_1="$__AF_clean96_v0";
echo "$__AF_clean96_v0__44_1" > /dev/null 2>&1
echo "Done"
