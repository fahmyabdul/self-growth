import time
import redis
import psycopg2
import sqlite3
import sqlalchemy.pool as pool
from contextlib import contextmanager
from pathlib import Path

from psycopg2.pool import ThreadedConnectionPool
from psycopg2.extras import DictCursor

from configs import configs
from app import logger

# sqlite_connection = None
postgre_pool = None
redis_connection = None
redisx_pools = {}

def sqlite_conn():
	sqlite_config = configs.properties["databases"]["sqlite"]

	Path(sqlite_config["path"]).mkdir(parents=True, exist_ok=True)

	sqlite_file = "{}/{}".format(sqlite_config["path"], sqlite_config["file"])
	try:
		sqlite_connection = sqlite3.connect(sqlite_file)
	except Exception as error:
		logger.log.error("SQLite | Unable to connect to SQLite, error: %s" % error)
		
		sqlite_connection = None
	else:
		logger.log.info("SQLite | Connected! | File: {}".format(sqlite_file))

	return sqlite_connection

def postgresql_conn(host, port, user, password, db, schema, min_conn, max_conn):
    global postgre_pool

    current_pool = None

    logger.log.info("PostgreSQL | Connecting to PostgreSQL Database connection pool...")

    try :
        current_pool = ThreadedConnectionPool(
			minconn=min_conn,
			maxconn=max_conn,
			database=db,
			user=user,
			password=password,
			host=host,
			port=port,
			cursor_factory=DictCursor,
			options="-c search_path={0}".format(schema),
		)
        
        postgre_pool = current_pool
    except Exception as error :
        logger.log.error("PostgreSQL | Unable to connect to PostgreSQL, error: %s" % error)
		
        postgre_pool = None
    else:
        logger.log.info("PostgreSQL | Connected!")

    return current_pool

@contextmanager
def postgre_pool_getconn():
	conn = postgre_pool.getconn()
	try:
		yield conn, conn.cursor()
	finally:
		postgre_pool.putconn(conn)

def redis_conn(host,port,auth,db):
    global redis_connection

    logger.log.info("Redis | Connecting to Redis Database...")

    
    redis_connection = redis.Redis(host,port,db,auth)

    try:
        redis_connection.ping()
    except Exception as e:
        logger.log.error("Redis | Unable to connect to Redis, error: {}".format(e))
        redis_connection = None
        return None
    else: 
        logger.log.info("Redis | Connected!")

    return redis_connection

def redis_connx(conn_name,host,port,auth,db,added_port,debugging):
	global redisx_pools

	if debugging is True:
		db = int(db)+added_port
	else:
		port = int(port)+added_port
	
	logger.log.info("Redis | Trying to connect to Redis [ %s ] : {host=%s,port=%s,db=%s}" % (conn_name, host, port, db))
	
	redisx_pools[conn_name] = redis.ConnectionPool(host=host,port=port,db=db,password=auth)

	while True:
		try:
			test_conn = redis.Redis(connection_pool=redisx_pools[conn_name])
			test_conn.ping()
		except Exception as ex:
			logger.log.error("Redis | Unable to connect to Redis, error: {}".format(ex))
			exit()
		else: 
			logger.log.info("Redis | Connected!")
			break

	time.sleep (0.1)

	return redisx_pools[conn_name]