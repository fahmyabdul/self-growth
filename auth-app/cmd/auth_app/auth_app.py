from app import app
from services.api import api
from app import logger

class AuthApp:        
    def run(self, config_file: str, log_file: str):
        service_name = "Auth App"
                    
        init_error = app.initialize(config_file, log_file, service_name)
        
        logger.log.info("Starting {}".format(service_name))
        
        if init_error != None:
            return init_error

        obj_service = api.ServicesApi()
        obj_service.start()
        
        return None
