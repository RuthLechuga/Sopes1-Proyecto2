FROM python

WORKDIR /usr/src/app
RUN pip install --upgrade pip
RUN pip install pika
RUN pip install pymongo
RUN pip install redis
RUN pip install grpcio
RUN pip install grpcio-tools
COPY . .
EXPOSE 8080
CMD ["python","main.py"]