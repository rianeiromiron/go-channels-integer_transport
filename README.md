# Go Channels Integer Transport

This project is a simple Go example that demonstrates inter-process communication using channels.

## What it does

The program creates a channel and starts a producer goroutine that generates random integer values. Those values are sent through the channel and received by the main goroutine, which prints each one to the console. Once the producer finishes sending data, it closes the channel and the consumer stops reading.

## Requirements

- Go 1.24.5 or later

## Run the project

From the project root, execute:

```bash
go run main.go
```

## Project files

- [main.go](main.go): contains the producer/consumer channel example
- [go.mod](go.mod): Go module definition

## Example output

The program prints messages showing the producer sending values and the main function receiving them until the channel is closed.
