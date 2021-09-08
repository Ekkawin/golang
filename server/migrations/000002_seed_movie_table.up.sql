COPY directors(director_id, first_name, last_name)
FROM '/Users/aek/Desktop/MyProj/golang/server/files/directors.csv'
DELIMITER ','
CSV HEADER;