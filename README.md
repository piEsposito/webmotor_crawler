# WebMotors crawler in Go

A very fast, scalable and extendible crawler for webmotors.com based in concurrent goroutine requests.

## Useful info
It requests using ubuntu-firefox header from the API that the web-app uses to render the page. As Go handling of Json files is very painful, we save, for each consult, the .json file on the folder. You then can store it into a database or do whatever you want. 

## Extendibility
I don't know why you would do it, but if you want you cant extend it to use with other APIs 

## Installation
After git-cloning it, edit the goroutine number, page_start and page_end parameters on **main.go** (which correspond the start and end page for the API request), the run the following snippet on the project root.

```
mkdir json_files
cd json_files
go build ../main.go
./main
```

Made by **Pi Esposito**
