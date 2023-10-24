# go-skeleton.
Golang setup code

##### Add the below env variables to run the server side of application.
HOST<br />
PORT<br />
USER<br />
PASSWORD<br />
DBNAME<br />
MAX_OPEN_CONNECTION<br />
MAX_IDLE_CONNECTION<br />
REDIS_URL<br />

##### Docker command to buid & run the application.
```
$ docker build -t go-test
$ docker run -p 8000:8000 -it --rm go-test
```

