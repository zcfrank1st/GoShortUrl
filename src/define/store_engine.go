package define

import "log"

type Engine interface {
    GetUrl(id int64) (string, error)
    StoreUrl(url string) (int64, error)
}

type MysqlEngine struct {}

func (my *MysqlEngine) GetUrl(id int64) (string, error) {
    var url []byte
    row := Db.QueryRow("select origin_url from short where id = ?", id)
    err := row.Scan(&url)
    go my.updateModifyTime(id)
    if err != nil {
        log.Println(err)
        return "", err
    }
    return string(url), nil
}

func (my *MysqlEngine) StoreUrl(url string) (int64, error) {
    res, err := Db.Exec("insert into short (origin_url, last_modify_time) VALUES (?, now())", url)

    if err != nil {
        log.Println(err)
        return -1, err
    }

    lastInsertId, err1 := res.LastInsertId()
    if err1 != nil {
        log.Println(err)
        return -1, err
    }

    return lastInsertId, nil
}

func (my *MysqlEngine) updateModifyTime(id int64) {
    if _, err := Db.Exec("update short set last_modify_time = now() where id = ?", id); err != nil {
        log.Println("update modify time failed")
    }
}