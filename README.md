# webhookserver

execute any (predefined) command through a webhook

this is an API to execute commands on your computer/server through HTTP requests. it's made in Go and uses JWTs to authenticate.

## how to set up?

- download a binary from the Releases page (available for Linux and macOS, both for amd64 and arm64) or build from source (`go build`)
- set the following environment variables (or save them to a .env file):
  - `WS_JWT_SECRET`: a random secure string to be used for jwt creation/validation
  - `WS_LISTENING_ADDRESS`: to set a custom listening address. by default it's `127.0.0.1:3000`
  - `WS_CONFIG_PATH`: custom path to the configuration json file. by default it's `{cwd}/config.json`
- create a `config.json` file: see [config.example.json](./config.example.json)
  - `name`: a random unique name. make sure it's safe to use it in an http url
  - `command`, `args`: the command/args to be used
  - `environment`: a list of environment variables to be set during execution of the command. it's an array of strings in the following format: `key=value`
  - `workingDirectory`: the directory where the command is going to be run in.
  - `timeout`: seconds. zero means no timeout.
- create a jwt by running the program with the "jwt" option
  - the `--time` parameter is optional and sets the longevity of the jwt. by default, the jwt won't have an expiration date

```sh
go run main.go jwt \
    --name something-that-identifies-this [--time nSeconds]
```

```sh
./webhookserver jwt --name something-that-identifies-this [--time nSeconds]
```

- do HTTP POST requests to `{url}/execute/{id of the job you want to execute}`
  - set the `Authorization` header to `Bearer {your jwt token}`

# license

MIT