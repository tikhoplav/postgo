# POSTGO

**POSTGO** - is a fetherweight golang postgresql quiry builder that makes process of making SQL queries easy, fast and repeatable. 

This is a home project in a early development stage. If you are looking for finished libraries, please, take a look at such libraries like [goqu](https://github.com/doug-martin/goqu), [loukoum](https://github.com/ulule/loukoum) and [buildsqlx](https://github.com/arthurkushman/buildsqlx).

Main goal of **POSTGO** library is to make regular golang developer be able to utilize the most powerfull postgresql tools with the lowest possible syntax complexity.

## Ver 0.1. Goals

### SELECT queries

```
sql, _ := postgo.Select().From("users").ToSQL()
fmt.Println(sql)

// SELECT * FROM "users";
```

#### Columns with aliases

```
sql, _ := postgo.Select().
	Col("name").
	Col("email").
	From("users").
	ToSQL()

fmt.Println(sql)

// SELECT "users"."name", "users"."email" FROM "users";
```

Also columns can be retrieved with alias.

```
sql, _ := postgo.Select().
	Col("name", "user_name").
	Col("email", "login").
	From("users").
	ToSQL()

fmt.Println(sql)

// SELECT "users"."name" AS "user_name", "users"."email" AS "login" FROM "users";
``` 


### JOIN

```
sql, _ := postgo.Select().
	From("users").
	Join("orders").
	On("users.id", "orders.user_id").
	ToSQL()

fmt.Println(sql)

// SELECT * FROM "users" INNER JOIN "orders" ON ("users"."id" = "orders"."user_id");
```

### JOIN USING

```
sql, _ := postgo.Select()
	From("users").
	Join("orders").
	Using("card_number").
	ToSQL()

fmt.Println(sql)

// SELECT * FROM "users" INNER JOIN "orders" USING ("card_number");
```

### UNION

```
distributors := postgo.Select("name").
	From("distributors").
	WhereLike("name", "W%")

actors := postgo.Select("name").
	From("actors").
	WhereLike("name", "W%")

query, err := postgo.Union(distributors, actors).
	ToSQL()

if err != nil {
	panic(err)
}

fmt.Println(query)

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
topFilms := postgo.From("top_films").
	Select("title").
	Select("release_year").	
	WhereMore("yearly_income", 1000000)

filmsByGenre := postgo.From("films_by_genre").
	Select("title").
	Select("release_year").
	Where("genre", "action")

sql, _ := postgo.Except(topFilms, filmsByGenre).
	ToSQL()

fmt.Println(sql)

// SELECT "title", "release_year" 
// FROM "top_films" 
// WHERE "yearly_income" > '1000000' 
// EXCEPT SELECT "title", "release_year" 
// FROM "films_by_genre"
// WHERE "genre" = 'action';
```

### INTERSECT

```
sql, _ := postgo.Intersect(postgo.From("most_popular_films"), postgo.From("top_rated_films")).
	ToSQL()

fmt.Println(sql)

// SELECT * FROM "most_popular_films" INTERSECT SELECT * FROM "top_rated_films";
```

### WITH closures

```
regionalSales := postgo.
	Select("region").
	Select("SUM(amount)", "total_sales").
	From("orders").
	GroupBy("region")

topTen := postgo.
	From("regional_sales").
	Select("SUM(total_sales)/10")

topRegions := postgo.
	From("regional_sales").
	Select("region").
	WhereMore("total_sales", topTen)

query, _ := postgo.
	With(regionalSales, "regional_sales").
	With(topRegions, "top_regions").
	Select("region").
	Select("product").
	Select("SUM(quantity)", "product_units").
	Select("SUM(amount)", "product_sales").
	From("orders").
	WhereIn("region", postgo.Select("region").From("top_regions")).
	GroupBy("region").
	GroupBy("product").
	ToSQL()

fmt.Println(query)

// WITH "regional_sales" AS (
//         SELECT "region", SUM("amount") AS "total_sales"
//         FROM "orders"
//         GROUP BY "region"
//      ), "top_regions" AS (
//         SELECT "region"
//         FROM "regional_sales"
//         WHERE "total_sales" > (SELECT SUM("total_sales")/'10' FROM "regional_sales")
//      )
// SELECT "region", "product",
//    SUM("quantity") AS "product_units",
//    SUM("amount") AS "product_sales"
// FROM "orders"
// WHERE "region" IN (SELECT "region" FROM "top_regions")
// GROUP BY "region", "product";
```

### PREPARE

```
query := postgo.Select().

preapre, _ := query.Prepare()

fmt.Println(prepare)

// PREPARE usrrptplan (int) AS
//    SELECT * FROM users u, logs l WHERE u.usrid=$1 AND u.usrid=l.usrid
//    AND l.date = $2;

execute, _ := query.Execute()
fmt.Println(execute)

// EXECUTE usrrptplan(1, current_date);
```