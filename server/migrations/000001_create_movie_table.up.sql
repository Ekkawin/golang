
CREATE TABLE "directors" (
    "director_id" text NOT NULL,
    "first_name" text,
    "last_name" text,
    PRIMARY KEY ("director_id")
);

CREATE TABLE "movies" (
    "movie_id" text NOT NULL,
    "title" text,
    "description" text,
    "director_id" text,
    "rank" text,
    "year" text,
    "run_time" text,
    "vote" text,
    "rating" text,
    "revenue" text,
    "meta_score" text,
    CONSTRAINT "fk_directors_movies" FOREIGN KEY ("director_id") REFERENCES "directors"("director_id"),
    PRIMARY KEY ("movie_id")
);

CREATE TABLE "actors" (
    "actor_id" text NOT NULL,
    "first_name" text,
    "last_name" text,
    PRIMARY KEY ("actor_id")
);

CREATE TABLE "generes" (
    "genere_id" text NOT NULL,
    "name" text,
    PRIMARY KEY ("genere_id")
);

CREATE TABLE "movie_actors" (
    "actor_actor_id" text NOT NULL,
    "movie_movie_id" text NOT NULL,
    CONSTRAINT "fk_movie_actors_movie" FOREIGN KEY ("movie_movie_id") REFERENCES "movies"("movie_id"),
    CONSTRAINT "fk_movie_actors_actor" FOREIGN KEY ("actor_actor_id") REFERENCES "actors"("actor_id"),
    PRIMARY KEY ("actor_actor_id","movie_movie_id")
);

CREATE TABLE "movie_generes" (
    "genere_genere_id" text NOT NULL,
    "movie_movie_id" text NOT NULL,
    CONSTRAINT "fk_movie_generes_genere" FOREIGN KEY ("genere_genere_id") REFERENCES "generes"("genere_id"),
    CONSTRAINT "fk_movie_generes_movie" FOREIGN KEY ("movie_movie_id") REFERENCES "movies"("movie_id"),
    PRIMARY KEY ("genere_genere_id","movie_movie_id")
);