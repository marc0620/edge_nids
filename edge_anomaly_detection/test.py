from pickle import dump, load
import numpy as np
import types
import sklearn
import time

income_data = [10025, 45, 0, 6, 0, 0, 133333.3333, 44444.44444, 45, 45, 0, 0, 0, 40, 22222.22222, 29200, 0]
model = load(open("cicids_model", "rb"))
a = input()
print(time.time()*1000)
a = a.replace("[", "")
a = a.replace("]", "")

a = a.split(",")
a=[float(i) for i in a]
print(time.time()*1000)
print(model.predict([a]))
