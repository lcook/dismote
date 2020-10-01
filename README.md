- [Getting started](#getting-started)
  - [Configuration](#configuration)
  - [Building/running](#buildingrunning)
  - [Commands](#commands)

## dismote ("emote stealer") bot

Minimal Discord bot which listens on the provided channels as defined in the configuration file, parses a users message accordingly and adds any found emotes to the server. Both static and animated emotes are supported granted the server has enough emote slots left.

## Getting started

### Configuration

You'll find a [yaml configuration](config.yaml) in this repository that the bot will use. The variables are as follows:

- **bot (bool)**: **true** if running as a bot and **false** if running as a self-bot (not recommended since self-bots are in violation of the Discord ToS).
- **token (str)**: Bot/account token.
- **channels (array)**: Array of channels to listen on.
- **prefix (str, default: +)**: Command prefix.

**Note**: the configuration file ("config.yaml") must be in the same diectory as the binary or it won't be found, you can specify a custom path by passing the config flag ("-c") to dismote.

You can conveniently reload the current settings by sending a *SIGHUP* signal to the PID of dismote. Likewise for *SIGTERM* that terminates the appplication. 

```console
$ kill -HUP `pidof dismote` # Reload the settings
$ kill -TERM `pidof dismote` # Gracefully terminate the application
```

### Building/running

I've include a Makefile to save a few keystrokes. To build, it's as simple as:

```console
$ make build # Build the resulting ("dismote") binary
/usr/local/bin/go build -ldflags="-s -w" -o dismote
$ ./dismote
2020/09/17 15:34:18 Loaded configuration file ("config.yaml")
2020/09/17 15:34:19 Successfully started Discord session
2020/09/17 15:34:19 Now listening on 2 channel(s)
--- truncated ---
```

Logs are outputted to the terminal, so it's suggested running the application inside tmux or screen to background the process.

### Commands

Whilst dismote offers little functionality it does come with a few handy commands. Command functions can be found [here](internal/commands).

| Command  | Description |
| ------------- | ------------- |
| help | Prints available commands |
| info | Displays servers statistics |
| clear | Bulk delete previous messages |
| listen | Set listening status |

Implicity, the "stealer" command is set for the default command to be ran. To readily add emotes to a server you must be in a channel that's listed in the settings, from there type out the emotes you want added. **Note**: messages that don't start with an emote are ignored and will not be added. 
