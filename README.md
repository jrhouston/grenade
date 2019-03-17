# Grenade

Grenade a tool for creating deployment failures. I created this tool at [Wattpad](http://wattpad.com) to answer questions like _What happens when a developer deploys a service that chews up a ton of memory?_ and _How quickly do I get an alert when log volume in production spikes dramatically?_.

This is a useful activity to do particularly when operating at a scale where the blast radius of a failure in production can be quite large. For example, if you run a service in production with 300 replicas, it's useful to see how your system behaves when all 300 start misbehaving. This can help you figure out what checks to put in place before a service is allowed to completely roll out in production.

## Get grenade

### Download the binary

### Docker image

### Using `go install`

## Usage

`grenade` is a simple command line utility, see `grenade --help`

## Supported failure modes

TODO

## Deploying to Kubernetes

### Using plain kubernetes YAML manifest

TODO

### Using helm

TODO