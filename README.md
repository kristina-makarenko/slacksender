# slacksender

The program reads the json-file and sends messages to Slack channels by interacting with the bot. The json-file contains
information about the token and messages to different Slack channels.

## Installation

1. You need to first install Go and set up a Go workspace.
2. Clone the repo:

 ```shell
 git clone https://github.com/kristina-makarenko/slacksender.git
 ```

3. Sign up for Slack and create a workspace
4. Create channels `test1`, `test2`, `test3` and create a bot in Slack
5. Copy `message-example.json` file to `message.json` with updating bot token:

 ```json
{
  "bot_token": "your_bot_token"
}
 ```

## Usage

1. Download dependencies

```shell
make dep
```

2. Build

 ```shell
make build
 ```

3. Run

```shell
make run
```

4. Check for messages in channels in Slack
