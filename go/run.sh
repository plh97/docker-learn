docker run \
--name mygolang \
--rm -it \
-d \
-p 8080:80 \
-p 8081:81 \
-p 8082:82 \
-p 8083:83 \
-v $(pwd):/root/app \
-w /root/app \
pengliheng/golang:latest