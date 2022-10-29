from http.client import OK
from flask import Flask, request, make_response, jsonify
from flask_cors import CORS
from services import arima_forecast
import sys

app = Flask(__name__)
CORS(app)


@app.route("/")
def helloWorld():
    return "<h1>Hola mundo</h1>"


'''
# ROUTE FOR REGISTER SERVICES ================================================
@app.route("/register-service", methods=['POST'])
def registerService():
    
    body = request.get_json()

    try:

        result = flights.registerService(body)

    except:
        result = {
            'msg': "Error while registering service",
            'code': 500
        }

    return make_response(jsonify(result))

### END DEF ==================================================================
'''



# ROUTE FOR HEALTH CHECK =====================================================
@app.route("/healthcheck", methods=['GET'])
def healthcheck():
    
    result = {
        'msg': "Everything ok.",
        'code': 200
    }

    return make_response(jsonify(result["msg"]), result["code"])

### END DEF ==================================================================


# # ROUTE FOR CREATE FLIGHTS ===================================================
@app.route("/forecast/arima", methods=['POST'])
def forecastArima():
    
    body = request.get_json()

    try:

        result = arima_forecast.get_prediction(body)

    except Exception as e:
        print(e)
        
        result = {
            'msg': "Error while making arima forecast",
            'code': 500
        }

    return make_response(jsonify(result["msg"]), result["code"])

# ### END DEF ==================================================================



# # ROUTE FOR BOOK FLIGHTS =====================================================
# @app.route("/reservar-vuelo/<usuario>", methods=['GET'])
# def getReservationsByUser(usuario):
    
#     try:

#         result = flights.getReservationsByUser({'id_usuario':usuario})

#     except:
#         result = {
#             'msg': "Error while fetching reservations",
#             'code': 500
#         }

#     return make_response(jsonify(result["msg"]), result["code"])

# ### END DEF ==================================================================





# TESTING DATABASE CONNECTION ================================================

# if (db.connect()):
#     print("Database connected", file=sys.stderr)
# else:
#     print("Error connecting Database", file=sys.stderr)

if __name__ == '__main__':
    app.run(host='0.0.0.0',port=8083, debug=False)

# ============================================================================
