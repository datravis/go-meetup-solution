import grpc

import pipeline_pb2
import pipeline_pb2_grpc


def generate_messages():
    messages = [
        "First message",
        "Second message",
        "Third message",
        "Fourth message",
        "Fifth message",
    ]
    for msg in messages:
        print("Sending %s" % (msg))
        yield pipeline_pb2.IngestRequest(
            input=msg,
        )

channel = grpc.insecure_channel('localhost:9002')
stub = pipeline_pb2_grpc.PipelineServiceStub(channel)

resp = stub.Ingest(generate_messages())
print("Received message %s" % (resp))