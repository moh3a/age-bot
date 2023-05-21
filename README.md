# Slack bots in golang

- botcmd: for greetings and to calculate age
- upload: for image uploading to slack
- ai: a smart bot with wit.ai and wolfram

## ai bot setup

### create wolfram app

- copy app ID and add it as an environment variable

### create wit.ai app

- intents: create `wolfram` intent
- add built-in entities `wit/wolfram_search_query`
- settings: copy server access token and add it as an environment variable
