## What

Demostrates sending Phabricator notifications to Slack with Go.

## Why

Phabricator, Slack, and Go are my favorite tools.

## How

Build the code:

    go get github.com/yinghau76/phabricator-to-slack

Run it directly:

    PORT=... \
    PHABRICATOR_HOST=... PHABRICATOR_USER=... PHABRICATOR_CERT=... \
    SLACK_TOKEN=... SLACK_CHANNEL= ... \
    phabricator-slack

Or install it:

    sudo phabricator-to-slack install
    
For Ubuntu 12.04, edit `/etc/init/phabricator-to-slack.conf` to configure settings
        
    env PORT=...
    env SLACK_CHANNEL="#team"
    env SLACK_TOKEN=...
    env PHABRICATOR_HOST=...
    env PHABRICATOR_USER=...
    env PHABRICATOR_CERT=...

Finally configure Phabricator to publish notifications in `/config/edit/feed.http-hooks/`.

## Thanks

Justin Tulloss - for him blog post ["Authenticating With Phabricator"](https://justin.harmonize.fm/development/2013/06/29/authenticating-with-phabricator.html)
