FROM python:3.11.3

WORKDIR /app
ADD . /app
RUN pip install -r requirements.txt

ENV PORT=3000
EXPOSE 3000
CMD ["python", "app.py"]