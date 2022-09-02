## URL-Shortner

> Tech Stack Used - <br/>
>
> - Golang
> - MongoDB

### How to send data:-
Example :-
```json
{
    "url": "https://www.youtube.com/watch?v=i_23KUAEtUM"
}
```

Response we get :-
```json
{
    "url": "https://www.youtube.com/watch?v=i_23KUAEtUM",
    "shortUrl": "localhost:3200/6799fa",
    "urlId": "6799fa",
    "createdTime": "2022-09-02T15:09:54.2226656+05:30"
}
```

Here ***shortUrl*** is the shorted url of the original URL. 

