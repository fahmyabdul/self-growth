FROM continuumio/miniconda3

WORKDIR /auth-app

COPY . .

RUN apt-get update && apt-get install -y build-essential libpq-dev

RUN conda env create -f conda.yml

SHELL ["conda", "run", "-n", "auth_app_env", "/bin/bash", "-c"]

ENTRYPOINT ["conda", "run", "--no-capture-output", "-n", "auth_app_env", "/bin/bash", "scripts/run.sh"]
