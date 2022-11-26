#!/usr/bin/env bash

## Complete the following steps to get Docker running locally

# Step 1:
# Build image and add a descriptive tag
docker build -t dunglh9/my-app .

# Step 2: 
# List docker images
docker images ls

# Step 3: 
# Run Go app
docker run -p 8080:8080 dunglh9/my-app