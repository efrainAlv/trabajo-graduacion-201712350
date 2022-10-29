FROM python:latest

WORKDIR /user/src/app

COPY ./src .

RUN pip install statsmodels
RUN pip install numpy
RUN pip install pandas



CMD ["python", "app.py"]
