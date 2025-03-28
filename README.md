# map-cronjob-provider-wc

Uses library: `https://github.com/robfig/cron` to handle cron expressions.   

> Disclaimer: This is WIP, use in your environment at your own risk.

# Note
There are [two packages](https://github.com/orgs/Mattilsynet/packages?repo_name=map-cronjob-provider-wc) in this github repo:  
1. `map-cronjob-provider-wc:v{m:m:p}`
2. `cronjob:{m:m:p}`  

First is the provider you'll link to from your Wadm manifest    
Second is the Wit package `wash wit fetch` will use to fetch dependencies in your `/wit/deps` folder as part of your `wash build` command    

# How to use the cronjob provider in your application

## WADM
Add the yaml-snippet underneath to your `wadm.yaml` or check out full example in [local.wadm.yaml](./local.wadm.yaml)

```yaml
- name: cronjob
      type: capability
      properties:
        image: ghcr.io/mattilsynet/map-cronjob-provider-wc:v0.0.17
      traits:
        - type: link
          properties:
            target: log-component
            namespace: mattilsynet
            package: test
            interfaces: [cron-handler]
            source_config:
              - name: log-component-cron-expression
                properties:
                  expression: "@every 2s"  # <--- This is what you're looking for!  
```

Replace "expression" with examples from: [pkg.go.dev/github.com/robfig/cron](https://pkg.go.dev/github.com/robfig/cron) to suit your own cron needs,  
| **description**                          | **expression**     |
|------------------------------------------|-------------------------|
| **Run Every Minute**                     | `@every 1m`             |
| **Run Every Hour**                       | `@every 1h`             |
| **Run Every Day at Midnight**            | `0 0 * * *`             |
| **Run Every Sunday at 6:00 PM**          | `0 18 * * SUN`          |
| **Run Every Monday at 9:00 AM**          | `0 9 * * MON`           |
| **Run Every 5 Minutes**                  | `*/5 * * * *`           |
| **Run at 12 PM and 6 PM Every Day**      | `0 12,18 * * *`         |
| **Run at 3:15 AM on the 1st Day of Every Month** | `15 3 1 * *`     |
| **Run Every Hour on the 15th Minute**   | `15 * * * *`            |
| **Run Every Week on Saturday at 9 AM**   | `0 9 * * SAT`           |
| **Run Every Day at Noon (12:00 PM)**     | `@daily`                |
| **Run Every Week at Midnight on Sunday** | `@weekly`               |
| **Run Every 15 Minutes**                 | `@every 15m`            |  

## WIT
Add wkg config, if you haven't installed the tool `wkg` then look under "Requires" under "Development" 
1. `wkg config --edit`
2. Add:
```toml
[namespace_registries]
mattilsynet = "ghcr.io"
```
3. Add `export mattilsynet:cronjob/cronjob@0.0.1;` to your `wit/world.wit`, e.g.,  
```WIT
package mattilsynet:logger;

world component {
  include wasmcloud:component-go/imports@0.1.0;
  export mattilsynet:cronjob/cronjob@0.0.1; # <-- look at me, I'm what you want!
}
```
4. `wash build` your component (you'll get some error saying your compoonent hasn't implemented the import requirement)
5. Add required code accordingly, look at [this code](./component/logger.go) for an example in Go.
# Development
If you'd like to test this provider together with a working component, then `git clone git@github.com:Mattilsynet/map-cronjob-provider-wc.git`, install requirements underneath and follow steps in "Quick start"

## Requires
tinygo >= 0.33.0 (`go install github.com/tinygo-org/tinygo@latest`)  
wrpc (`cargo install wrpc`)  
wasm-tools (`cargo install wasm-tools`)  
wash-cli (`cargo install wash`)   
wkg (`cargo install wkg`)

## Quick start
1. Add wkg config as described under "How to use the cronjob provider in your application"
2. `wash build` in root  
3. `cd component`
4. `wash build`  
5. Terminal 1: `wash up` 
6. Terminal 2: In root of project: `wash app deploy local.wadm.yaml --replace`  
7. Terminal 2: `watch wash app status cronjob-logger` # to view progress of cronjob-logger   
### Expected output:
#### Every two seconds (without altering local.wadm.yaml)
2024-11-14T08:40:51.990755Z  INFO log: wasmcloud_host::wasmbus::handler: Cronjob handler called component_id="cronjob_logger-log_component" level=Level::Info context="cronjob-handler"
2024-11-14T08:40:53.989890Z  INFO log: wasmcloud_host::wasmbus::handler: Cronjob handler called component_id="cronjob_logger-log_component" level=Level::Info context="cronjob-handler"

## Problems/Questions?
If you have any issues with the setup or questions, please don't hesitate to [create an issue](https://github.com/Mattilsynet/map-cronjob-provider-wc/issues)


