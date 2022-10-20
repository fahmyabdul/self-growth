import logging
import sys
import os

from logging.handlers import TimedRotatingFileHandler

log = None
console_handler = None
file_handler = None
log_services = {}

FORMATTER = logging.Formatter("%(asctime)s [%(name)s] [%(levelname)s] %(message)s")

def get_console_handler():
    global console_handler
    console_handler = logging.StreamHandler(sys.stdout)
    console_handler.setFormatter(FORMATTER)
    return console_handler

def get_file_handler(logger_path, logger_name):
    global file_handler
    if not os.path.exists(logger_path):
        os.mkdir(logger_path)

    file_handler = TimedRotatingFileHandler(logger_path+logger_name+'.log', when='midnight')
    file_handler.setFormatter(FORMATTER)
    return file_handler

def set_logger(logger_name, logger_path, logger_tofile):
    global log
    log = logging.getLogger(logger_name)
    log.setLevel(logging.DEBUG)
    log.addHandler(get_console_handler())
    if logger_tofile is True:
        log.addHandler(get_file_handler(logger_path+"/", logger_name))
    log.propagate = False
    return log

def set_logger_services(logger_name, logger_path, logger_tofile):
    global log_services
    log_services[logger_name] = logging.getLogger(logger_name)
    log_services[logger_name].setLevel(logging.DEBUG)
    log_services[logger_name].addHandler(get_console_handler())
    if logger_tofile is True:
        log_services[logger_name].addHandler(get_file_handler(logger_path, logger_name))
    log_services[logger_name].propagate = False
    return log_services[logger_name]
