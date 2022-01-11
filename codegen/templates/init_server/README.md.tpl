{{ $gofiber_port := .ProjectInfo.Port -}}
# gomakeme

```bash
# build the image
docker build -t gofiber .

# run and publish with the name of gofiber
docker run --publish {{$gofiber_port}}:{{$gofiber_port}} --name gofiber gofiber

# stop
docker stop gofiber

# remove
docker image rm gofiber
```