from concurrent import futures
import grpc
import time

from textblob import TextBlob

import sentiment_pb2
import sentiment_pb2_grpc

class SentimentServiceServicer(sentiment_pb2_grpc.SentimentServiceServicer):
    def Analyze(self, request, context):
        blob = TextBlob(request.input)
        
        return sentiment_pb2.SentimentResponse(
            sentiment=blob.sentiment.polarity,
        )

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    sentiment_pb2_grpc.add_SentimentServiceServicer_to_server(
        SentimentServiceServicer(), server)
    server.add_insecure_port('[::]:9000')
    server.start()
    try:
        while True:
            time.sleep(5)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()