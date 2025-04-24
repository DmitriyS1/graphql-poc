FROM ubuntu:latest
LABEL authors="dmitrii_semenov"

ENTRYPOINT ["top", "-b"]