import sys

import click

from .auth_app.auth_app import AuthApp

CONTEXT_SETTINGS = dict(help_option_names=['-h', '--help'])

@click.group(context_settings=CONTEXT_SETTINGS)
def main():
    pass

@main.command()
@click.option("--config", required=False, type=str, help="Define config file full path (e.g: ./.configs.yml)")
@click.option("--log", required=False, type=str, help="Define log file full path (e.g: ./log/)")
def auth_app(config, log):
    cmd = AuthApp()
    cmd.run(config, log)