#!/usr/bin/env python

from distutils.core import setup

setup(name='protogen',
      version='1.0',
      description='',
      author='Dan Travis',
      author_email='dan.travis@gmail.com',
      url='',
      py_modules=['ner_pb2', 'ner_pb2_grpc', 'sentiment_pb2', 'sentiment_pb2_grpc', 'pipeline_pb2', 'pipeline_pb2_grpc'],
     )