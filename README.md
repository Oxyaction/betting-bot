# Betting telegram bot

## Local development
### Setting webhook
1. Register on [ngrok](ngrok.com)  
2. Run ngrok locally and grab `https` url
3. Register webhook running
```
curl --location --request POST 'https://api.telegram.org/bot<token>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "<ngrok_url>"
}'
```
[More info](https://core.telegram.org/bots/api#setwebhook)  
### Running
```
API_TOKEN=<token> make start
```
