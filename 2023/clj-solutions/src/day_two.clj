(ns day-two
  (:require [clojure.string :as s]))

(def demo-data (s/split-lines (slurp "resources/day2_demo.txt")))
(def real-data (s/split-lines (slurp "resources/day2.txt")))

(defn parse-game [game]
  (map (fn [draw]
         (reduce (fn [m [_ n color]]
                   (assoc m (keyword color) (Integer/parseInt n)))
                 {}
                 (re-seq #"(\d+) (\w+)" draw)))
       (s/split game #";")))


(defn possible-draw? [game-data]
  (false? (or
           (> (get game-data :red 0) 12)
           (> (get game-data :green 0) 13)
           (> (get game-data :blue 0) 14)
           false)))

(defn possible-game? [game]
  (every? identity (map possible-draw? (parse-game game))))

(defn attach-id [bool-list]
  (into {} (map-indexed (fn [id val]
                          [(inc id) val]) bool-list)))

(defn solv1 [game-data]
  (let [possible-games (filter
                        (fn [game] (second game)) (attach-id (map possible-game? game-data)))
        possible-games-ids (map first possible-games)
        id-sum (reduce + 0 possible-games-ids)]
    id-sum))

(solv1 real-data)

;; => 2156

;; PART 2

(defn max-draws [game-data]
  (reduce (fn [m draw]
            (merge-with max m draw))
          {}
          game-data))

(defn powerize [game]
  (* (:red game 0) (:green game 0) (:blue game 0)))

(defn part-2 [game-data]
  (->> game-data
       (map parse-game)
       (map max-draws)
       (map powerize)
       (apply +)))

(part-2 real-data)
;; => 66909
