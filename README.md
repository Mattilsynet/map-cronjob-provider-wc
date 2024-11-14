# map-cronjob-provider-wc

Responsible for triggering a component every X seconds, i.e., provider is the source and component is the target

> Disclaimer: This is WIP and an experiment, use in your environment at your own risk.

## Requires
tinygo >= 0.33.0 (homebrew)
wrpc (cargo)
wasm-tools (cargo)
wash-cli (homebrew) 

## Optional
wit-deps

## How to get started
1. `wash build` in root  
2. `cd component`
3. mkdir gen/
4. go generate ./...
3. `wash build`  
4. Terminal 1: `wash up` 
5. Terminal 2: In root of project: `wash app deploy wadm.yaml --replace`  
6. Terminal 2: `watch wash app status cronjob-logger` # to view progress of cronjob-logger  

### Expected output:
#### Every two seconds (without altering wadm.yaml)
2024-11-14T08:40:51.990755Z  INFO log: wasmcloud_host::wasmbus::handler: Cronjob handler called component_id="cronjob_logger-log_component" level=Level::Info context="cronjob-handler"
2024-11-14T08:40:53.989890Z  INFO log: wasmcloud_host::wasmbus::handler: Cronjob handler called component_id="cronjob_logger-log_component" level=Level::Info context="cronjob-handler"


