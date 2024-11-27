# Use an ARG to specify the Python version (build-time variable)
ARG PYTHON_VERSION=3.9-slim

# Use the Python version as the base image
FROM python:${PYTHON_VERSION}

# Set an environment variable for the greeting message (runtime variable)


# Copy the Python script into the container
WORKDIR /app
COPY app6.py .

# Default command to run the Python app
CMD ["python", "app6.py"]
