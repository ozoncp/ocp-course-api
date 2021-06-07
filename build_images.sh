#!/bin/bash

docker image rm ocp-course-api:dev
docker image rm ocp-lesson-api:dev

docker build -t ocp-course-api:dev --force-rm --build-arg service=ocp-course-api .
docker build -t ocp-lesson-api:dev --force-rm --build-arg service=ocp-lesson-api .
