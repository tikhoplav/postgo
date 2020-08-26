# Contributing to Postgo

Thanks for your interest to Postgo. 

If you are interested in contributing to Postgo project you are totaly welcome to do that. Please send a [pull request](https://github.com/tikhoplav/postgo/pulls) with clear description of code you've created, it's goals, problems it solves, decisions you have made and any other detail which help to understand it's impact. Please use as much links to other resources such as official golang documentation, stack overflow or others, as you find comfortable. Each and every pull request will be processed with full passion, so more info you provide more probability of your work to be integrated forever into this project.

Please provide a message for each of your commits, if it adds some new functionality to library, removes or fixes, for example:
```
$ git commit -m 'Added ability to add columns to select query'
```

The bigger commit, the bigger description it requires:
```
$ git commit -m 'Added SelectQuery class which acts like a SQL string build for select queries;
Added Select function for column specifications;
Added SelectAs function for column with alias specifications;
Added related test cases;
Added function From to package with which new SelectQueries should be instantiated;
'
```