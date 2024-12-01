## Testing

### Unit Tests:

    1. Using testify package for assertions.
    2. Use stretchr package for mocking the DB

### e2e Testing:

    1. Dockerize the project with postgres, to enable local db connection for tests.
    2. Use httptest package for testing the routes
    3. Make combinations of all the possible wrong values to test if validations are working.
    4. Check for correct values to test if db operations are working.

## Asymptotic analysis

### For fetch APIs with id:

    1. Without indexing of the id column the complexity will O(n), as postgres will perform a sequential scan.
    2. With indexing, postgres will use a B-tree, so the lookup time will be log(n)

### For fetch APIs with type:

    1. Without indexing of the id column the complexity will O(n), as postgres will perform a sequential scan.
    2. With Indexing:
        a. If the cardinality of the column is high and because search for type is equality based, hash indexes might be suitable. So the time complexity for locating the rows is constant O(1) and based on number rows K the total time complexity will be O(K + 1) = O(K)

### For Fetching Sum:

    1. If id and parent_id columns are indexed, then as we are using recursive CTE, the lookup for columns will be log(n) and if the K is total number of child, grandchild, so on nodes then the time complexity will be
    O(log(n) * K)
    2. As the query is complex, if #reads >>> #writes then we can use materialized views to precompute the result and directly search the view. and refresh the view during modifications. We can also add an index on materialized view, so time complexity will O(log(n)).

### For upsert

    1. With indexing, lookup will be log(n), insert will be O(1) + O(log n) (insertion into B-tree index), if update happens then O(1), total complexity will be O(log(n))
