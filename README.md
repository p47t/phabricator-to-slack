# Phabricator-to-Slack

[![Build Status](https://travis-ci.org/yinghau76/phabricator-to-slack.svg?branch=master)](https://travis-ci.org/yinghau76/phabricator-to-slack)

## What?

The code demostrates sending Phabricator notifications to Slack. It is implemented with Go majorly for advatange of single-binary deployment.

## Why?

Phabricator and Slack are among my favorite tools. It will be nice if they can work together.

## How?

Phabricator-to-Slack works by installing itself as a HTTP hook of Phabricator. After receiving notification from Phabricator, it digs more information from Phabricator via its Conduit API and send the message to specified Slack channel.

### Build the code

    go get github.com/yinghau76/phabricator-to-slack

### Install as system service

    sudo phabricator-to-slack install

#### Ubuntu 12.04

Ubuntu 12.04 is the only tested platform. However, it should work only any platform that Go supports.

If you are using upstart, you should edit `/etc/init/phabricator-to-slack.conf` and add the following environment variables:

    env PORT=...             # e.g. 8000
    env SLACK_CHANNEL=...    # e.g. #random
    env SLACK_TOKEN=...      # find it in https://api.slack.com/
    env PHABRICATOR_HOST=... # e.g. https://your.phabricator.host
    env PHABRICATOR_USER=... # Phabricator username
    env PHABRICATOR_CERT=... # find it in /settings/panel/conduit/ of your Phabricator site

### Phabricator Setup

Remember to configure Phabricator (`/config/edit/feed.http-hooks/`) so that notifications can be published to phabricator-to-slack.

## Thanks

Justin Tulloss - for him blog post ["Authenticating With Phabricator"](https://justin.harmonize.fm/development/2013/06/29/authenticating-with-phabricator.html)
