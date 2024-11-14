# map-cronjob-provider-wc

## Responsible for triggering a component every X seconds, i.e., provider is the source and component is the target

## Requires
tinygo >= 0.33.0   
wit-bindgen-go  
wash-cli  
## Optional
wit-deps

## How to get started
1. `wash build` in root  
2. `cd component`  
3. `wash build`  
4. Terminal 1: `wash up`  
5. Terminal 2: `wash app deploy wadm.yaml --replace`  
6. Terminal 2: `watch wash app status cronjob-logger` # to view progress of cronjob-logger  

### Expected output:
#### Every two seconds (without altering wadm.yaml)
2024-11-14T08:40:51.990755Z  INFO log: wasmcloud_host::wasmbus::handler: Cronjob handler called component_id="cronjob_logger-log_component" level=Level::Info context="cronjob-handler"
2024-11-14T08:40:53.989890Z  INFO log: wasmcloud_host::wasmbus::handler: Cronjob handler called component_id="cronjob_logger-log_component" level=Level::Info context="cronjob-handler"


