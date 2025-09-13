FROM golang:latest AS builder
WORKDIR /app
COPY . .