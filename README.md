# map-cronjob-provider-wc

Uses `https://github.com/robfig/cron` to handle cron expressions.   


> Disclaimer: This is WIP, use in your environment at your own risk.


# How to use the cronjob provider in your application

Look at `local.wadm.yaml` on how the link configuration is expected to look, use the commented image block in your wadm    
Replace "expression" with examples from: `https://pkg.go.dev/github.com/robfig/cron` to suit your own cron needs   
Add wkg config, if you haven't installed the tool `wkg` then look under "Requires" under "Development" 
1. `wkg config --edit`
2. Add:
```
[namespace_registries]
mattilsynet = "ghcr.io"
```

# Development
If you'd like to test this provider together with a working component, then clone the project, install requirements underneath and follow steps in "Quick start"

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

