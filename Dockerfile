FROM python:3.10
COPY . /home/
WORKDIR /home/
RUN apt-get update -y
# RUN python3 /home/app.py
