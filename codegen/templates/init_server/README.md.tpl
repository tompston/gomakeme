{{ $GOLANG_PORT := .ProjectInfo.Port -}}
# gomakeme

```bash
# build the image
docker build -t gofiber .

# run and publish with the name of gofiber
docker run --publish {{$GOLANG_PORT}}:{{$GOLANG_PORT}} --name gofiber gofiber

# stop
docker stop gofiber

# remove
docker image rm gofiber
```