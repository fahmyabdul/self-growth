from configs import configs
from app.logger import set_logger
from app import databases
from internal.initialization import initialization

current_version = "v2022.10-17-1"
log_path = None

def initialize(config_file, log_path_input, service_name) -> Exception:
    global log_path
    # Initializing Application
    try:
        if config_file == None:
            config_file = ".configs.yml"

        # Loading Config File
        config = configs.init_config(config_file)
        log_file = service_name.lower().replace(" ", "_")

        log_path = log_path_input
        if log_path == None:
            log_path = config["logger"]["path"]

        # Setting application logger
        set_logger(log_file, log_path, config["logger"]["to_file"])

        # Connection databases
        connect_databases()

        # Do custom initialization
        initialization.initialize()
    except Exception as e:
        return e

def connect_databases():
    # Connecting to PostgreSQL Database
    psql_conf = configs.properties["databases"]["postgre"]
    if psql_conf["status"] == True:
        pool_postgre = databases.postgresql_conn(
                psql_conf["host"], 
                psql_conf["port"], 
                psql_conf["user"], 
                psql_conf["pass"], 
                psql_conf["db"], 
                psql_conf["schema"], 
                psql_conf["min_conn"], 
                psql_conf["max_conn"]
            )
            
        if pool_postgre == None:
            exit()
    
    # Connecting to Redis Database
    redis_conf = configs.properties["databases"]["redis"]
    if redis_conf["status"] == True:        
        redisx_main = databases.redis_connx("main",redis_conf['host'], redis_conf['port'], redis_conf['auth'], redis_conf['db'], 0, configs.properties["debugging"])
        if redisx_main == None:
            exit()
