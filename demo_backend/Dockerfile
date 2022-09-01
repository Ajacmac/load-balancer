FROM golang

WORKDIR /backend

COPY * /backend/

RUN go build .

CMD ["go", "run", "."]