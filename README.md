# nof

**nof** is a tiny CLI tool that runs shell commands defined in simple YAML templates.

:exclamation: This is a working in progress, so the code might be messy, tests might be lacking 
and things might break frequently.

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
  - "\"*.log\""
  - "-mtime"
  - "-3"
```

Then run:

```
nof /path/to/find.yaml
```

## Installation

1. Clone the repo:

   ```
   git clone https://github.com/yourusername/nof.git
   cd nof
   ```

2. Build the binary (requires Go):

   ```
   go build -o nof
   ```

## Usage

```
nof /path/to/command.yaml
```

## License

GNU GPL-3.0
