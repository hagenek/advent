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

(defn max-draws [game-data]
  (reduce (fn [m draw]
            (merge-with max m draw))
          {}
          game-data))



(map max-draws (map parse-game demo-data))
;; => ({:blue 6, :red 4, :green 2} {:blue 4, :green 3, :red 1} {:green 13, :blue 6, :red 20} {:green 3, :red 14, :blue 15} {:red 6, :blue 2, :green 3})

(first (map max-draws (map parse-game demo-data)))
;; => {:blue 6, :red 4, :green 2}

(defn powerize [draw]
  (* (:blue draw 0) (:red draw 0) (:green draw 0)))


(defn possible-draw? [game-data]
  (false? (or
           (> (get game-data :red 0) 12)
           (> (get game-data :green 0) 13)
           (> (get game-data :blue 0) 14)
           false)))

(map possible-draw? (parse-game (nth real-data 1)))

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



