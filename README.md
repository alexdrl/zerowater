# zerowater

I was having fun playing with Go, the [Zerolog](https://github.com/rs/zerolog) library and [Watermill](https://github.com/ThreeDotsLabs/watermill)... did not find a good adapter for the last one, so why not building a new one as a way of playing more with Go? 

This is a pet project as a .NET developer so feel free to criticize 😉

## Example Usage:
```go
logger := zerowater.NewZerologLoggerAdapter(log.Logger.With().Str("component", "windmill").Logger())

pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
```

## Ideas:
- Implement tests with mocks and not using the actual log.
- Check why the stringer test on the original adapter for the log package on watermill does not work on zerolog.