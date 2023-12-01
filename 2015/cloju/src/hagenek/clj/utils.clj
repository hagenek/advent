(ns hagenek.clj.utils
  (:require [clojure.java.io :as io]))

(defn read-resource [resource]
  (slurp (io/resource resource)))