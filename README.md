Go gin micro server for uploading files


# How to run

### Running with docker CLI
```bash
docker run \
	--env TOKEN=seCreTt0keN \
	--env MAX_MB=2 \
	--name test-uploader \
	--publish 8080:8080/tcp \
	--volume "$(pwd)/test:/data" \
	0009000/go-upload-server
```

### Running through docker-compose
```yml
services:
  upload-server:
    image: 0009000/go-upload-server
    container_name: upload-server
    environment:
      - TOKEN=seCreTt0keN 
      - MAX_MB=2
    ports:
      - 8080:8080
    volumes:
      - '/srv/data:/data'
```

# How to use

Once your docker is running you can upload files to the server making simple request
```bash
curl \
  -H "Token: seCreTt0keN" \
  -F "path=dump/folder" \
  -F "upload=@/home/root/file.txt" \
  -F "upload=@/home/root/another_file.png" \
  localhost:8080/upload
```


# Environment variables

There are 2 optional environment variables:

1. `TOKEN` can be used to make service a bit more safe. If not set, it will not be checked at all.
2. `MAX_MB` configuration property for gin's **MaxMultipartMemory**. Default is 32