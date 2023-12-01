(ns day-one
  (:require [clojure.string :as s]
            [clojure.edn :as edn]))


;; Part One

(slurp "resources/day1_part1_demo.txt")
;; => "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\n"

(def data (slurp "resources/day1_part1_demo.txt"))
;; => #'day-one/data

(re-seq #"\d" data)
;; => ("1" "2" "3" "8" "1" "2" "3" "4" "5" "7")

(s/split-lines data)
;; => ["1abc2" "pqr3stu8vwx" "a1b2c3d4e5f" "treb7uchet"]

(first (s/split-lines data))
;; => "1abc2"

(re-seq #"\d" (first (s/split-lines data)))
;; => ("1" "2")

(first (re-seq #"\d" (first (s/split-lines data))))
;; => "1"

(last (re-seq #"\d" (first (s/split-lines data))))
;; => "2"

(str (first (re-seq #"\d" (first (s/split-lines data))))
     (last (re-seq #"\d" (first (s/split-lines data)))))

;; => "12"

(Integer/parseInt (str (first (re-seq #"\d" (first (s/split-lines data))))
                       (last (re-seq #"\d" (first (s/split-lines data))))))
;; => 12

(map (fn [line] (Integer/parseInt (str (first (re-seq #"\d" line))
                                       (last (re-seq #"\d" line))))) (s/split-lines data))
;; => (12 38 15 77)

(apply + (map (fn [line]
                (Integer/parseInt (str (first (re-seq #"\d" line))
                                       (last (re-seq #"\d" line))))) (s/split-lines data)))
;; => 142

;; Defining functions

(defn parse-line [line]
  (Integer/parseInt (str (first (re-seq #"\d" line))
                         (last (re-seq #"\d" line)))))

(->> (slurp "resources/day1_part1_demo.txt")
     s/split-lines
     (map parse-line)
     (apply +))

;; => 142

;; Solution 
(defn parse-line [line]
  (Integer/parseInt (str (first (re-seq #"\d" line))
                         (last (re-seq #"\d" line)))))
(defn solve-part1
  [data]
  (->> data
       s/split-lines
       (map parse-line)
       (apply +)))

;; PART TWO

(def digit-dict {"one" "on1e" "two" "tw2o",
                 "three" "th3ree", "four" "4r",
                 "five" "5e", "six" "6e",
                 "seven" "7n", "eight" "eig8th",
                 "nine" "nin9e"})

(defn word->digit [str [word digit]]
  (s/replace str word digit))

(defn words->digit-string [string]
  (reduce word->digit
          string
          digit-dict))


(defn solve-part2 [data]
  (solve-part1 (s/join "\n" (map words->digit-string (s/split-lines data)))))

(solve-part2 (slurp "resources/day1.txt"))
;; => 54885

