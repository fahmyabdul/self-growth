#!/bin/sh

conda activate efishery_auth_app_env

# Hand off to the CMD
exec "$@"
