import json
import math

import numpy as np
from scipy.optimize import curve_fit


def exponential(x, b, c):
    return (b ** x) + c


def run(future, values):
    x = list(range(len(values)))
    pars, cov = curve_fit(f=exponential, xdata=x, ydata=values, p0=[1, 0], bounds=(-np.inf, np.inf))
    day_to_predict = x[-1] + future
    prediction = exponential(day_to_predict, pars[0], pars[1])
    print(prediction)


def main():
    r = json.loads(input())
    future = r[0]
    values = r[1]
    run(future, values)


if __name__ == '__main__':
    main()
