FROM python:3.10
WORKDIR /app
COPY . .
RUN pip install -r requirements.txt
RUN chmod +x /app/skripta.sh
ENTRYPOINT ["/app/skripta.sh"]
