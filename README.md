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
