FROM golang:onbuild
COPY . /code
LABEL pengliheng "pengliheng111@gmail.com"
WORKDIR /code

RUN go install



CMD ["gin", "run" , "main.go"]