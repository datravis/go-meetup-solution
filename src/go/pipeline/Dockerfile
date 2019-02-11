FROM ubuntu

RUN mkdir /app
WORKDIR /app 
ADD pipeline/service /app/.
EXPOSE 9000
CMD ["/app/service"]