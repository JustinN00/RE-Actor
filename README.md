# RE-Actor

Hi! Thanks for coming to look at RE-Actor! This application is intended to apply a Regular Expression based pattern to the logs of a Docker container. If the pattern matches the log, then the values are interpolated into the supplied template. Once that is complete, the resulting value will be sent to the selected action. Currently, the only supported action is posting a message to a Discord webhook.

Here is a list of environment variables that are used to configure RE-Actor:

| Variable Name | Description |
|---------------|-------------|
| `CONTAINER_NAME` | This is the name of the container you would like RE-Actor to monitor. It should only match one container. |
| `PATTERN` | This is the Regular Expression pattern to run against each log from the respective container. Please ensure that this pattern is valid and does what you expect. Here is a great resource for creating your pattern: [Regex101](https://regex101.com). |
| `TEMPLATE` | if the current log message matches your pattern, this defines how the output will look. Please refer to the previous linked resource for help on this. |
| `DISCORD_WEBHOOK_URL` | This is the webhook that RE-Actor will post the resulting message to. |
| `LOG_LEVEL` | This variable is optional, but determines how verobose the logging will be. Here are the valid options: `DEBUG`, `INFO`, `WARNING`, and `ERROR`. The default value is `INFO`, if not set. |