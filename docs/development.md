# Development environment

> Cool for looking by - hopfully you'll find this helpful. If not: _the source is everything you need_ :)

## Backend

> IT's go :D nothing really to explain, install dependencies and build LEL.

```
lel@srv:~/lel$ go get
lel@srv:~/lel$ goxc
lel@srv:~/lel/cmd$ go build . -t LEL
```

For running configuration you can use dev.yml

```bash
lel@srv:~/lel/cmd$ ./LEL -d -c ../dev-config.yml
```

## Frontend

Since I'm no frontend expert i'd love some tips :D

Start LEL and the npm serve

```bash

lel@srv:~/lel/cmd$ go run *.go -d
lel@srv:~/lel/static$ npm install <stuff>
lel@srv:~/lel/static$ npm run serve
```

Then go to [http://127.0.0.1:8080/dev.html](http://127.0.0.1:8080/dev.html) to see the react component.

## Contributing

Thanks for looking at this file. If you want to contribute feel free to do so and give something back to the community.

:)

## Errors & Pitfalls

if compiling isn't working increase the inodes: [Answer](https://webpack.js.org/configuration/watch/#not-enough-watchers)

```
echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p
```
