# Homelab application configuration

Service configuration is grouped by application. HA, Node-RED and MQTT belong in `homelab-automation`; shared execution belongs in `homelab`; actual placement, exact pins and private-input references belong in private `homelab-site`.

The `compose/` groups retain the reviewed service definitions for integration work. This draft is not deployed and does not relocate any live bind mount or data. See `SPLIT-DRAFT.md` for remaining checks.

See [split input paths](docs/split-inputs.md) for required repository roots and private assembled runtime inputs.
