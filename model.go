package main

import "database/sql"

type product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
    return db.QueryRow("select name, price from product where id=?", p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updateProduct(db *sql.DB) error {
    _, err := db.Exec("update product set name=?, price=? where id=?", p.Name, p.Price, p.ID)
    return err
}

func (p *product) deleteProduct(db *sql.DB) error {
    _, err := db.Exec("delete from product where id=?", p.ID)
    return err
}

func (p *product) createProduct(db *sql.DB) error {
    res, err := db.Exec("insert into product(name, price) values(?, ?)", p.Name, p.Price)

    if err != nil {
        return err
    }

    id, err := res.LastInsertId()
    if err != nil {
        return err
    }

    p.ID = int(id)
	return nil
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
    rows, err := db.Query("select id, name, price from product limit ? offset ?", count, start)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    products := []product{}

    for rows.Next() {
        var p product
        if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
            return nil, err
        }
        products = append(products, p)
    }

    return products, nil
}
