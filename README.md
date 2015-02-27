## What

Demostrates sending Phabricator notifications to Slack with Go.

## Why

Phabricator, Slack, and Go are my favorite tools.

## How

### Build the code

    go get github.com/yinghau76/phabricator-to-slack

### Run it directly

    PORT=... \
    PHABRICATOR_HOST=... PHABRICATOR_USER=... PHABRICATOR_CERT=... \
    SLACK_TOKEN=... SLACK_CHANNEL= ... \
    phabricator-slack

### Install as system service

    sudo phabricator-to-slack install
    
#### Ubuntu 12.04

Ubuntu 12.04 is the only tested platform.

You should edit `/etc/init/phabricator-to-slack.conf` and add the following environment variables:
        
    env PORT=...             # e.g. 3003 
    env SLACK_CHANNEL=...    # e.g. #random
    env SLACK_TOKEN=... 
    env PHABRICATOR_HOST=... # e.g. https://your.phabricator.host
    env PHABRICATOR_USER=...
    env PHABRICATOR_CERT=...

### Phabricator Setup

Remember to configure Phabricator (`/config/edit/feed.http-hooks/`) so that notifications can be published to phabricator-to-slack.

## Thanks

Justin Tulloss - for him blog post ["Authenticating With Phabricator"](https://justin.harmonize.fm/development/2013/06/29/authenticating-with-phabricator.html)
