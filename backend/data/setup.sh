#!/bin/bash
psql -f ./setup.sql -U user -d chitchat
echo "Database setup complete."