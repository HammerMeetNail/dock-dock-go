# dock-dock-go
A simple golang project for playing with the Docker API

# Table of Contents

- [Overview](#overview)
- [Build](#build)
- [Run](#run)
- [Options](#options)
- [Examples](#example)

# Overview

# Build
`docker build -t dock-dock-go:local .`

# Run
`docker run -v /var/run/docker.sock:/var/run/docker.sock dock-dock-go:local`

# Options

* **DOCKER_VERSION** 
    * Summary: Docker API and Client version
    * Valid Values: Version >= `1.38`
    * Default: `1.39`

* **LOG_LEVEL** 
    * Summary: Level at which to output logs
    * Valid Values: `[warn, debug]`
    * Default: `warn`

* **OUTPUT_TYPE** 
    * Summary: Where to send output
    * Valid Values: `[statsd, stdout]`
    * Default: `stdout`

* **OUTPUT_INTERVAL** 
    * Summary: How frequently, in milliseconds, to send output
    * Valid Values: Integer >= `0`
    * Default: `2000`

* **STATSD_SERVER_URI** 
    * Summary: URI of StatsD server
    * Valid Values: URL
    * Default: `http://localhost`

* **STATSD_SERVER_PORT** 
    * Summary: Port of StatsD server
    * Valid Values: Integer > `0`
    * Default: `8125`

# Examples
`docker-compose up -d`

`echo "deploys.test.myservice:1|c" | nc -w 1 -u 127.0.0.1 8125`
`echo "automagic.vol1.size:71|g" | nc -w 1 -u 127.0.0.1 8125`