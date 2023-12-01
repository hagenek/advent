(ns hagenek.clj.core
  (:require [clojure.java.io :as io]))


(def input (slurp (io/resource "day2.txt")))

(require '[clojure.string :as str])

(count (str/split input #"\n"))
