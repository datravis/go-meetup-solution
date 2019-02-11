import concurrent.futures
import grpc
import os
import praw
import sys

import pipeline_pb2
import pipeline_pb2_grpc

client_id = os.environ['CLIENT_ID']
client_secret = os.environ['CLIENT_SECRET']
username = os.environ['USERNAME']
password = os.environ['PASSWORD']
user_agent = os.environ['USER_AGENT']

subreddits = os.environ['SUBREDDITS']
subreddits = subreddits.split(',')

reddit = praw.Reddit(client_id=client_id,
                     client_secret=client_secret, password=password,
                     user_agent=user_agent, username=username)


channel = grpc.insecure_channel('pipeline:9000')
stub = pipeline_pb2_grpc.PipelineServiceStub(channel)

def yield_comments(subreddit):
    for comment in reddit.subreddit(subreddit).stream.comments():
        yield pipeline_pb2.IngestRequest(
            source=subreddit,
            input=comment.body
        )

def write_comments(subreddit):
    # for comment in yield_comments(subreddit):
    #     print(comment)
    while True:
        try:
            resp = stub.Ingest(yield_comments(subreddit))
            print(resp)
        except Exception as e:
            print(e)

executor = concurrent.futures.ThreadPoolExecutor(max_workers=10)
for subreddit in subreddits:
    executor.submit(write_comments, subreddit)

executor.shutdown(wait=True)
    