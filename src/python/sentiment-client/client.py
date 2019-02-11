import grpc

import sentiment_pb2
import sentiment_pb2_grpc

channel = grpc.insecure_channel('localhost:9001')
stub = sentiment_pb2_grpc.SentimentServiceStub(channel)

resp = stub.Analyze(
    sentiment_pb2.SentimentRequest(
        input=" Have looked at the DOW 30?"
    )
)
print(resp)