# Grenade 💣 🔥

![](https://media.giphy.com/media/9PkfGzhKwBDHPTnDSj/giphy.gif)

Grenade a tool for creating deployment failures. I created this tool at [Wattpad](http://wattpad.com) to answer questions like _What happens when a developer deploys a service that chews up a ton of memory?_ and _How quickly do I get an alert when log volume in production spikes dramatically?_

This is a useful activity to do particularly when operating at a scale where the blast radius of a failure in production can be quite large. For example, if you run a service in production with 300 replicas, it's useful to see how your system behaves when all 300 start misbehaving. This can help you figure out what checks to put in place before a service is allowed to completely roll out in production, and help you tune your alert monitors.

**DISCLAIMER:** grenade does not "simulate" problems through monitoring trickery, it actually causes destructive problems. Don't deploy this to production unless you really know what you're doing. 

## Get grenade

### Download the binary

### Docker image

### Using `go install`

## Usage

`grenade` is a simple command line utility, see `grenade --help`

```
usage: grenade [<flags>]

Flags:
  --help                   Show context-sensitive help (also try --help-long and --help-man).
  --delay=DELAY            trigger the grenade after specified delay e.g 5m30s
  --crash                  Process crash
  --exit                   Exit process with non-zero exit code
  --mem-spike              Cause a memory spike
  --mem-leak               Cause a memory leak
  --cpu-spike              Cause a CPU spike
  --goroutine-leak         Cause goroutines to be leaked
  --conn-spike=CONN-SPIKE  Cause a connect spike to an address
  --conn-leak=CONN-LEAK    Cause a connect leak to an address
  --dns-spike              Cause a spike in DNS lookups
  --log-spike              Cause a spike in log volume
  --fd-spike               Cause a spike in open file descriptors
  --disk-spike             Cause a spike in disk space usage
  --disk-leak              Cause a leak in disk space usage
```

## Supported failure modes

TODO

## Deploying to Kubernetes

### Using plain kubernetes YAML manifest

TODO

### Using helm

TODO