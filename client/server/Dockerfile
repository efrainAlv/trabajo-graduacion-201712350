FROM python:latest

WORKDIR /user/src/app

COPY ./src .

RUN pip install flask
RUN pip install python-dotenv
RUN pip install flask_cors
RUN pip install pandas
RUN pip install statsmodels

ENV FLASK_APP=app.py
ENV FLASK_DEBUG=true

EXPOSE 8083

CMD ["flask", "run", "--host", "0.0.0.0", "--port", "8083"]
