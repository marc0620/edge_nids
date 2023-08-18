import pandas as pd
import seaborn as sns
import numpy as np
import re
import sklearn

import warnings

warnings.filterwarnings("ignore")

import matplotlib.pyplot as plt
import matplotlib as matplot
from IPython.core.interactiveshell import InteractiveShell

InteractiveShell.ast_node_interactivity = "all"

from sklearn.model_selection import train_test_split

# from sklearn import linear_model
# from sklearn.ensemble import VotingClassifier
# from sklearn.ensemble import AdaBoostClassifier
# from sklearn.ensemble import RandomForestClassifier
# from sklearn.ensemble import BaggingClassifier
# from sklearn.ensemble import ExtraTreesClassifier
# from sklearn.ensemble import GradientBoostingClassifier
# from sklearn.ensemble import IsolationForest
# from sklearn.svm import SVC
# from sklearn.tree import DecisionTreeClassifier
from pickle import dump, load

train = pd.read_csv("./KDDTrain+.csv")
test = pd.read_csv("./KDDTest+.csv")
print(train.shape)
test.shape
train.columns = range(train.shape[1])
test.columns = range(test.shape[1])
labels = [
    "duration",
    "protocol_type",
    "service",
    "flag",
    "src_bytes",
    "dst_bytes",
    "land",
    "wrong_fragment",
    "urgent",
    "hot",
    "num_failed_logins",
    "logged_in",
    "num_compromised",
    "root_shell",
    "su_attempted",
    "num_root",
    "num_file_creations",
    "num_shells",
    "num_access_files",
    "num_outbound_cmds",
    "is_host_login",
    "is_guest_login",
    "count",
    "srv_count",
    "serror_rate",
    "srv_serror_rate",
    "rerror_rate",
    "srv_rerror_rate",
    "same_srv_rate",
    "diff_srv_rate",
    "srv_diff_host_rate",
    "dst_host_count",
    "dst_host_srv_count",
    "dst_host_same_srv_rate",
    "dst_host_diff_srv_rate",
    "dst_host_same_src_port_rate",
    "dst_host_srv_diff_host_rate",
    "dst_host_serror_rate",
    "dst_host_srv_serror_rate",
    "dst_host_rerror_rate",
    "dst_host_srv_rerror_rate",
    "attack_type",
    "difficulty_level",
]  # subclass - > attack_type
combined_data = pd.concat([train, test])
combined_data.head(5)
combined_data.columns = labels
combined_data = combined_data.drop("difficulty_level", 1)
combined_data.head(3)
from sklearn import preprocessing

le = preprocessing.LabelEncoder()

print(
    set(list(combined_data["attack_type"]))
)  # use print to make it print on single line
combined_data["attack_type"] = le.fit_transform(combined_data["attack_type"])
combined_data["protocol_type"] = le.fit_transform(combined_data["protocol_type"])
combined_data["service"] = le.fit_transform(combined_data["service"])
combined_data["flag"] = le.fit_transform(combined_data["flag"])

print("\nDescribing attack_type: ")
print("min", combined_data["attack_type"].min())
print("max", combined_data["attack_type"].max())
print("mean", combined_data["attack_type"].mean())
print("mode", combined_data["attack_type"].mode())
print("looks like 16 is 'normal' ")
# select least correlated
corr_matrix = combined_data.corr().abs().sort_values("attack_type")
# tmp.head(10) # to view CORR matrix
leastCorrelated = corr_matrix["attack_type"].nsmallest(10)
leastCorrelated = list(leastCorrelated.index)

# select least correlated
leastSTD = combined_data.std().to_frame().nsmallest(5, columns=0)
leastSTD = list(
    leastSTD.transpose().columns
)  # fuckin pandas.core.indexes.base.Index   -_-
# tmp = tmp.append('num_outbound_cmds')  # might not work...
featureElimination = set(leastCorrelated + leastSTD)
len(featureElimination)
featureElimination
# dont change combined_data, we will neeed it latter
combined_data_reduced = combined_data.drop(featureElimination, axis=1)
data_x = combined_data_reduced.drop("attack_type", axis=1)
data_y = combined_data_reduced.loc[:, ["attack_type"]]
# del combined_data # free mem

import gc

gc.collect()


X_train, X_test, y_train, y_test = train_test_split(
    data_x, data_y, test_size=0.5, random_state=42
)  # TODO
print(
    "Thats how to rid rid of {0} dimentions of data, from the 10 lowest STD and 5 lowest correlation".format(
        len(featureElimination)
    )
)

print(X_test.shape)
RF = load(open("ff", "rb"))
print(type(X_train))
test_list = X_train.values.tolist()
income_data = test_list[0]
print((income_data))
print(RF.predict([income_data]))
