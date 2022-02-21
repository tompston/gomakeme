# gomakeme

```bash
# build the image
docker build -t gofiber .

# run and publish with the name of gofiber
docker run --publish 4444:4444 --name gofiber gofiber

# stop
docker stop gofiber

# remove
docker image rm gofiber
```