docker run \
--name mygolang \
--rm -it -d \
-p 8080:80 \
-p 8081:81 \
-v $(pwd):/go/src \
-w /go/src \
pengliheng/golang:latest