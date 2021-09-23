#! /bin/bash

path="/Users/lzz/Documents/094熊逸说苏轼"

function read_dir() {
    for file in $(#注意此处这是两个反引号，表示运行系统命令
        ls $path
    ); do
        if [ -d $path"/"$file ]; then #注意此处之间一定要加上空格，否则会报错
            read_dir $path"/"$file
        else
            #  echo $path"/"$file #在此处处理文件即可
            result=$(echo $file | grep ".m4a")
            if [[ "$result" != "" ]]; then
                #  echo ${file%%.*}
                ffmpeg -i $path"/"$file -f mp3 -acodec libmp3lame -y $path"/"${file%%.*}.mp3
            fi
        fi
    done
}
#读取第一个参数
read_dir
