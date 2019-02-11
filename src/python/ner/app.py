from concurrent import futures
import grpc
import spacy
import time

import ner_pb2
import ner_pb2_grpc

class NerServiceServicer(ner_pb2_grpc.NerServiceServicer):
    def __init__(self):
        self.nlp = spacy.load('en_core_web_sm')

    def ExtractSubjectOrgs(self, request, context):
        doc = self.nlp(request.input)

        orgs = []
        for ent in doc.ents:
            if ent.label_ == "ORG":
                orgs.append(ent.text)
        
        return ner_pb2.NerResponse(
            entities=orgs,
        )

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    ner_pb2_grpc.add_NerServiceServicer_to_server(
        NerServiceServicer(), server)
    server.add_insecure_port('[::]:9000')
    server.start()
    try:
        while True:
            time.sleep(5)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()