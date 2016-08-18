#!/usr/bin/env bash

set -e

echo "Collect static files"
python3 manage.py collectstatic --noinput

echo "Apply database migrations"
python3 manage.py migrate

echo "Starting server"
python3 manage.py runserver 0.0.0.0:8000
