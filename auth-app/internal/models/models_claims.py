from pydantic import BaseModel, StrictStr

class RequestClaims(BaseModel):
    phone: StrictStr
    password: StrictStr