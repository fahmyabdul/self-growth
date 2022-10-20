from typing import Optional
from pydantic import BaseModel, StrictStr, StrictInt, validator

from app import logger, databases

class RequestUser(BaseModel):
    phone: StrictStr
    name: StrictStr
    role: StrictStr
    password: Optional[StrictStr]
    created_at: Optional[StrictStr]
    logged_in: Optional[StrictStr]
    
    @validator("*", pre=False, always=True)
    def check_field_length(cls, value, field, values):
        max_length = {
            "phone": {
                "check_length": True,
                "length": [16],
            },
            "name": {
                "check_length": True,
                "length": [60],
            },
            "role": {
                "check_length": True,
                "length": [10],
            },
            "password": {
                "check_length": True,
                "length": [4],
            },
            "created_at": {
                "check_length": False,
                "length": [19],
            },
            "logged_in": {
                "check_length": False,
                "length": [19],
            },
        }

        field_config = max_length[field.name]
        if field_config["check_length"] == False:
            return value
        
        if len(str(value)) > field_config["length"][0]:
            raise ValueError('field length {} is invalid, max length: {}'.format(len(str(value)), field_config["length"][0]))
        
        return value

    class Config:
        validate_assignment = True

class Users(RequestUser):
    def create_table(self) -> Exception:
        with databases.sqlite_conn() as conn:
            try:
                cur = conn.cursor()
                cur.execute("CREATE TABLE IF NOT EXISTS t_users(id INTEGER PRIMARY KEY AUTOINCREMENT, phone TEXT, name TEXT, role TEXT, password TEXT, created_at TEXT)")
                conn.commit()
            except Exception as e:
                logger.log.info("Models Users | Create Table t_users | Failed, error: {}".format(e))
                return e
            
        return None

    def get_by_phone(self) -> Exception:
        with databases.sqlite_conn() as conn:
            try:
                cur = conn.cursor()
                cur.execute("Select phone from t_users where phone = ?", (self.phone,))
                rows = cur.fetchall()
            except Exception as e:
                logger.log.info("Models Users | Get By Phone | Failed, error: {}".format(e))
                return None, e
            
        return rows, None

    def check_password(self, password) -> Exception:
        with databases.sqlite_conn() as conn:
            try:
                cur = conn.cursor()
                cur.execute("Select password from t_users where password = ?", (password,))
                rows = cur.fetchall()
            except Exception as e:
                logger.log.info("Models Users | Check Password | Failed, error: {}".format(e))
                return None, e
            
        return rows, None

    def get_by_filter(self, filter) -> Exception:
        with databases.sqlite_conn() as conn:
            try:
                cur = conn.cursor()
                cur.execute("Select * from t_users where {}".format(" AND ".join(filter)))
                rows = cur.fetchall()
            except Exception as e:
                logger.log.info("Models Users | Get By Phone | Failed, error: {}".format(e))
                return None, e
            
        return rows, None

    def register(self):
        params = (self.phone, self.name, self.role, self.password, self.created_at)

        with databases.sqlite_conn() as conn:
            try:
                cur = conn.cursor()
                cur.execute("INSERT INTO t_users(phone, name, role, password, created_at) VALUES(?, ?, ?, ?, ?)", params)
                conn.commit()
            except Exception as e:
                logger.log.info("Models Users | Create Table t_users | Failed, error: {}".format(e))
                return e
            else:
                logger.log.info("Models Users | Create Table t_users | Success")
        
        return None