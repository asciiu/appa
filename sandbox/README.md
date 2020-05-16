# Youtube downloader. 

### Dependencies 
* go - v1.12
* brew install ffmpeg (OSX)
* apt install ffmpeg (Linux)

### Run
```
$ go run main.go https://www.youtube.com/watch?v=LWE79K2Ii-s 
```

### Build command
```
$ export SSH_PRIVATE_KEY="$(cat ../id_rsa)"
$ make build
$ make push
```