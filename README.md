WASMCLOUD_OCI_REGISTRY=europe-north1-docker.pkg.dev WASMCLOUD_OCI_REGISTRY_USER=oauth2accesstoken WASMCLOUD_OCI_REGISTRY_PASSWORD=$(gcloud auth print-access-token) wash up

gcloud auth print-access-token | oras push -u oauth2accesstoken --password-stdin --artifact-type application/vnd.wasmcloud.provider.archive.layer.v1+par europe-north1-docker.pkg.dev/artifacts-352708/map/cronjob-provider:0.0.1 cronjob.par.gz
