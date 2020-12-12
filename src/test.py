import json
import math
import socket

import numpy as np
from scipy.optimize import curve_fit
from matplotlib import pyplot as plt

data = [6847, 6878, 6894, 6918, 7039, 7057, 7097, 7125, 7144, 7158, 7201, 7268, 7284, 7430, 7534, 7544, 7569, 7645,
        7645, 7561, 7651, 7657, 7678, 7682, 7685, 7688, 7716, 7912, 7919, 7933, 7943, 7956, 7982, 8071, 8078, 8094,
        8106, 8111, 8118, 8134, 8146, 8151, 8159, 8173, 8178, 8186, 8201, 8218, 8228, 8246, 8370, 8411, 8486, 8502,
        8564, 8612, 8641, 8673, 8783, 8817, 8846, 8889, 9002, 9070, 9141, 9236, 9277, 9372, 9434, 9515, 9601, 9664,
        9708, 9775, 9922, 10038, 10115, 10259, 10359, 10440, 10496, 10687, 11093, 11270, 11366, 11381, 11555, 11682,
        11751, 11911, 11957, 12201, 12624, 12716, 12913]

# x = list(range(len(data)))
# plt.plot(x, data)
# plt.savefig("test.png", dpi=600)
# plt.show()
# exit(0)


def f(x):
    return math.e ** x + 10


def exponential(x, a, b, c):
    return a * (b ** x) + c


def run(future, values):
    x = list(range(len(values)))
    # try:
    # print(values)
    pars, cov = curve_fit(f=exponential, xdata=x, ydata=values, maxfev=5000, p0=[1, 1, 0], bounds=(-np.inf, np.inf))
    # except:
    #     exit(0)
    day_to_predict = x[-1] + future
    prediction = exponential(day_to_predict, pars[0], pars[1], pars[2])
    # print(round(prediction))
    return round(prediction)
    # return lambda n: exponential(n, pars[0], pars[1], pars[2])


def main():
    host = '127.0.0.1'  # Standard loopback interface address (localhost)
    port = 5000  # Port to listen on (non-privileged ports are > 1023)

    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.setblocking(True)
        s.bind((host, port))
        s.listen()
        while True:
            conn, addr = s.accept()
            with conn:
                print('Connected by', addr)
                while True:
                    data = conn.recv(2048)
                    if not data:
                        break
                    r = json.loads(data.decode('utf-8'))
                    future = r[0]
                    values = r[1]
                    result = run(future, values)
                    conn.send(str(result).encode('utf-8'))
                print("Client disconnected", addr)

    # while True:
    #     r = json.loads(input())
    #     future = r[0]
    #     values = r[1]
    #     run(future, values)
    # g = run(1, data)
    # x = list(range(len(data)))
    # plt.plot(x, data)
    # plt.plot(x, [g(x) for x in x])
    # plt.savefig("test.png", dpi=600)
    # plt.show()


if __name__ == '__main__':
    main()
