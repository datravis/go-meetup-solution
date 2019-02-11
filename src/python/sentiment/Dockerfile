FROM grpc-base

ADD ./src/python/sentiment/app.py /

ADD /src/python/protogen /protogen
RUN pip install --no-cache-dir /protogen
RUN pip install textblob

EXPOSE 9000

CMD [ "python", "./app.py" ]