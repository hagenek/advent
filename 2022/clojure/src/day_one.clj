(ns day-one
  (:require [clojure.string :as s]
            [clojure.edn :as edn]))
;; Part One

(def data  (slurp "resources/day-one-data.txt"))

(defn extract-elfs [data]
  (let [split-data (s/split data #"\n\n")]
    (->> split-data
         (map #(s/split % #"\n"))
         (map #(map edn/read-string %))
         (map (partial reduce +)))))

(reduce max (extract-elfs data))
;; => 68923

;; Part Two
(reduce + (take 3 (reverse (sort (extract-elfs data)))))
;; => 200044
