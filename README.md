# nof

**nof** is a tiny CLI tool that runs shell commands defined in simple YAML templates.

:exclamation: This is a work in progress, so the code might be messy, tests might be lacking 
and things might break frequently. If you want a better version that has many more features visit [this](https://github.com/go-task/task) repo.
I just wanted something much more minimalistic that has a more readable and simpler yaml syntax.

## Example

The yaml file for the example below is present inside [examples/find.yaml](examples/find.yaml)


To run the command:

```
find /var/log "*.log" -mtime -3
```

Use the following YAML file:

```yaml
# find.yaml
find:
  - "/var/log"
  - "*.log"
  - "-mtime"
  - "-3"
```

Then run:

```
nof ./examples/find.yaml
```

## Installation

1. Clone the repo:

   ```
   git clone https://github.com/hosmir/nof.git
   cd nof
   ```

2. Build the binary (requires Go):

   ```
   go build -o ./bin/ -v ./...
   ```

OR

Just download the binary from the [releases](/releases) page and add it to your PATH. (The binary is built on for linux amd64)

## Usage

```
nof /path/to/command.yaml
```

## License

GNU GPL-3.0
