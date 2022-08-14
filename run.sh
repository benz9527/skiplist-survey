./skiplist-survey --filter=^[a-z0-9]+Inserts$ | tee inserts.csv
./skiplist-survey --filter=^[a-z0-9]+WorstInserts$ | tee worst-inserts.csv
./skiplist-survey --filter=^[a-z0-9]+AvgSearch$ | tee avg-search.csv
./skiplist-survey --filter=^[a-z0-9]+SearchEnd$ | tee search-end.csv
./skiplist-survey --filter=^[a-z0-9]+Delete$ | tee delete.csv
./skiplist-survey --filter=^[a-z0-9]+WorstDelete$ | tee worst-delete.csv
