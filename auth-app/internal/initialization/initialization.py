from app import logger

from internal.models.models_users import RequestUser, Users

def initialize():
    check_table_users()

def check_table_users():
    models_user = Users(phone="test", name="test", role="test", password="test", created_at="0000-00-00 00:00:00")
    err = models_user.create_table()
    if err != None:
        logger.log.warn("Check Table Users | Error: {}".format(err))
        return err
    
    logger.log.warn("Check Table Users | Done")

    return None