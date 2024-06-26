FROM golang:1.22.2-bookworm
# 已经测试过alpine
LABEL authors="zen"
RUN cp /etc/apt/sources.list.d/debian.sources /etc/apt/sources.list.d/debian.sources.bak
RUN sed -i 's/deb.debian.org/mirrors4.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list.d/debian.sources
RUN apt update
RUN apt install -y dos2unix
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/root/go/bin
RUN mkdir -p /root/app
WORKDIR /root/app
COPY . .
RUN go build -o /usr/local/bin/conv main.go
RUN chmod +x /usr/local/bin/conv
WORKDIR /usr/local/bin
RUN dos2unix /root/app/install-retry.sh
RUN chmod +x /root/app/install-retry.sh
RUN /root/app/install-retry.sh ffmpeg nano mediainfo build-essential

CMD ["conv"]
# docker build --no-cache -t audios:latest .
# docker run --name audio -dit -v /c/Users/zen/Github/FastBS4/rss/从乞丐到元首:/data audios:latest
# docker run -itd  --cpus=1 --memory=2048M --name test -v /f/large/bodysuit:/data test:5