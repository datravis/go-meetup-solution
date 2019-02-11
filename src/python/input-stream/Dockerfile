FROM grpc-base

ADD ./src/python/input-stream/app.py /

ADD /src/python/protogen /protogen
RUN pip install --no-cache-dir /protogen

RUN pip install praw

CMD [ "python", "./app.py" ]