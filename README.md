# POSTGO

**POSTGO** - is a fetherweight golang postgresql quiry builder that makes process of making SQL queries easy, fast and repeatable. 

This is a home project in a early development stage. If you are looking for finished libraries, please, take a look at such libraries like [goqu](https://github.com/doug-martin/goqu), [loukoum](https://github.com/ulule/loukoum) and [buildsqlx](https://github.com/arthurkushman/buildsqlx).

Main goal of **POSTGO** library is to make regular golang developer be able to utilize the most powerfull postgresql tools with the lowest possible syntax complexity.

## Ver 0.1. Goals
### JOIN

```
fmt.Println(query.ToSQL())

// Output
SELECT * FROM "users" INNER JOIN "orders" ON ("users"."id" = "orders"."user_id");
```

### JOIN USING

```
fmt.Println(query.ToSQL())

// Output
SELECT * FROM "users" INNER JOIN "orders" USING ("card_number");
```

### UNION

```
fmt.Println(query.ToSQL())

// Output
SELECT "distributors"."name"
    FROM "distributors"
    WHERE "distributors"."name" LIKE 'W%'
UNION
SELECT "actors"."name"
    FROM "actors"
    WHERE "actors"."name" LIKE 'W%';
```

### EXCEPT

```
fmt.Println(query.ToSQL())

// Output
SELECT "title", "release_year"
	FROM "top_films"
	WHERE "yearly_income" > '1000000' 
EXCEPT
SELECT "title", "release_year"
	FROM "films_by_genre"
	WHERE "genre" = 'action'
```

### INTERSECT

```
fmt.Println(query.ToSQL())

// Output
SELECT "title", "release_year"
	FROM "most_popular_films" 
INTERSECT
SELECT "title", "release_year"
	FROM "top_rated_films";
```

### WITH closures

```
fmt.Println(query.ToSQL())

// Output
WITH "regional_sales" AS (
        SELECT "region", SUM("amount") AS "total_sales"
        FROM "orders"
        GROUP BY "region"
     ), "top_regions" AS (
        SELECT "region"
        FROM "regional_sales"
        WHERE "total_sales" > (SELECT SUM("total_sales")/10 FROM "regional_sales")
     )
SELECT "region", 
		"product",
       	SUM("quantity") AS "product_units",
       	SUM("amount") AS "product_sales"
	FROM "orders"
	WHERE "region" IN (SELECT "region" FROM "top_regions")
	GROUP BY "region", "product";
```