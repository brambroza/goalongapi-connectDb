package db

import (
    "context"
    "database/sql" 
    _ "github.com/denisenkom/go-mssqldb"
)


func ExecuteQuery(ctx context.Context, db *sql.DB, query string) ([][]interface{}, error) {
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    columns, err := rows.Columns()
    if err != nil {
        return nil, err
    }

    result := make([][]interface{}, 0)

    values := make([]interface{}, len(columns))
    for i := range values {
        values[i] = new(interface{})
    }

    for rows.Next() {
        if err := rows.Scan(values...); err != nil {
            return nil, err
        }

        row := make([]interface{}, len(columns))
        for i, value := range values {
            row[i] = *(value.(*interface{}))
        }

        result = append(result, row)
    }

    return result, nil
}
