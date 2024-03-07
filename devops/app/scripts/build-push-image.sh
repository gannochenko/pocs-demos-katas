#!/usr/bin/env bash

REGION=asia-east1-docker.pkg.dev
PROJECT=go-app-390716
APP=devops-app

while getopts a:t:e: flag
do
    case "${flag}" in
        a) ACTION=${OPTARG};;
        t) TAG=${OPTARG};;
        e) ENV=${OPTARG};;
        *) exit 1
    esac
done

IMAGE="${REGION}"/"${PROJECT}"/devops-"${ENV}"/"${APP}":"${TAG}"

if [ "${ACTION}" = "build" ]
then
   docker build -t "${IMAGE}" .
fi

if [ "${ACTION}" = "push" ]
then
   gcloud config set project "${PROJECT}"
   docker push "${IMAGE}"
fi
