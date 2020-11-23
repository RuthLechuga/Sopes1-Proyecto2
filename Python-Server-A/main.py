from concurrent import futures
import logging

import grpc

import helloworld_pb2
import helloworld_pb2_grpc

import redis
import pymongo
import json

import threading

myclient = pymongo.MongoClient("mongodb://54.87.43.11:27017")
mydb = myclient["sopes1"]
mycol = mydb["casos"]

r = redis.Redis(
    host='3.84.197.252',
    port=6379,
)

def insertData(data):
    r.sadd('casos',data.replace("\"","'"))
    x = json.loads(data)
    y= mycol.insert_one(x)

class Greeter(helloworld_pb2_grpc.GreeterServicer):
    def SayHello(self, request, context):
        print(">>>>>Si entre>>>>")
        insertData(request.name)
        return helloworld_pb2.HelloReply(message='Recibido!')

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    helloworld_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    print("Escuchando...")
    logging.basicConfig()
    serve()
