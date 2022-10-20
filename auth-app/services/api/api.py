import logging

from flask import Flask, jsonify, abort, request
from flasgger import Swagger
from flask_restful import Api, Resource

from configs import configs
from app import logger
from app.app import current_version

from .controllers.ctrl_heartbeat import CtrlHeartbeat
from .controllers.ctrl_register import CtrlRegister
from .controllers.ctrl_claims import CtrlClaims
from .controllers.ctrl_validate import CtrlValidate

class ServicesApi:
    def __init__(self):
        # Loading the configuration file
        self.configs = configs.properties['services']['api']

        # Initialize Flask
        self.app = Flask(__name__)
        self.app.log = logging.getLogger('werkzeug')
        self.app.log.addHandler(logger.console_handler)
        if configs.properties['logger']['to_file'] is True:
            self.app.log.addHandler(logger.file_handler)
        self.app.log.setLevel(logging.INFO)
        self.app.config['SWAGGER'] = self.configs['swagger']
        self.app.config['SWAGGER']['version'] = current_version
        self.app.config['SWAGGER']['specs_route'] = "{}/swagger/index.html".format(self.configs['swagger']['basePath'])
        self.app.config['SWAGGER']['specs'] = [
            {
                "endpoint": 'apispec_1',
                'route': "{}/swagger/docs.json".format(self.configs['swagger']['basePath'])
            }
        ]

        self.api = Api(self.app)
        self.swagger = Swagger(self.app, template=self.configs['swagger'])

        self.flask_routes()

    def start(self):
        logger.log.info("API | Starting | Host: {} | Port: {}".format(self.configs['host'], self.configs['port']))
        self.app.run(host=self.configs['host'], port=self.configs['port'], debug=False, threaded=True)
    
    def flask_routes(self):
        self.main_routes="/api/v1/auth-app"
        
        self.api.add_resource(CtrlHeartbeat, "{}/{}".format(self.main_routes, 'heartbeat'))

        self.api.add_resource(CtrlRegister, "{}/{}".format(self.main_routes, 'register'))

        self.api.add_resource(CtrlClaims, "{}/{}".format(self.main_routes, 'claims'))
        
        self.api.add_resource(CtrlValidate, "{}/{}".format(self.main_routes, 'validate'))
