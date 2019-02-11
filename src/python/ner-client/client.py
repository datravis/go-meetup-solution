import grpc

import ner_pb2
import ner_pb2_grpc

channel = grpc.insecure_channel('localhost:9000')
stub = ner_pb2_grpc.NerServiceStub(channel)

resp = stub.ExtractSubjectOrgs(
    ner_pb2.NerRequest(
        input="Idk what to expect from ACB on Monday"
    )
)
print(resp)