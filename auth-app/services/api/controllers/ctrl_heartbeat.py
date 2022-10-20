from datetime import datetime
from app import logger
from app.app import current_version
from flask import request
from flask_restful import Resource

class CtrlHeartbeat(Resource):
    def get(self):
        """
        Get the heartbeat of the service
        ---
        tags: ['heartbeat']
        responses:
          200:
            description: Endpoint that return the heartbeat of the service
            schema:
              id: heartbeat
              properties:
                version:
                  type: string
                  example: v2022.10-17-1
                last_check:
                  type: string
                  example: "2022-10-19 06:50:29"
        """
        last_check = datetime.now()
        last_check = datetime.strftime(last_check, '%Y-%m-%d %H:%M:%S')
        
        return {'version': current_version, 'last_check': last_check }, 200