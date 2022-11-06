#!/bin/sh

conda activate auth_app_env

# Hand off to the CMD
exec "$@"
