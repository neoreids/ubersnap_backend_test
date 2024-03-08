## UBERSNAP Middle Backend Test

Build image docker with command 

```docker-compose build```

Run http service by docker container with command

```docker-compose up -d```

If you want to run test file just run command 

```docker exec ubersnap go test ./...```

For the link endpoints each service

[compress image](#compress-image)
```
http://localhost:8000/image/compress [POST]

multipart/form-data
image=<binary image>
quality=<int>
```

[convert PNG to JPG](#convert-image)
```
http://localhost:8000/image/convert [POST]

multipart/form-data
image=<binary image>
```

[resize dimension image](#resize-image)
```
http://localhost:8000/image/resize [POST]

multipart/form-data
image=<binary image>
width=<int>
height=<int>
```

