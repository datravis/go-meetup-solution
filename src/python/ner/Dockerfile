FROM grpc-base

ADD ./src/python/ner/app.py /

ADD /src/python/protogen /protogen
RUN pip install --no-cache-dir /protogen

RUN python -m spacy download en_core_web_sm

EXPOSE 9000

CMD [ "python", "./app.py" ]