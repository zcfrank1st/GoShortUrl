package define

type Engine interface {
    GetUrl(id int) (string, error)
    StoreUrl(url string) (int, error)
}

type MysqlEngine struct {}

func (my *MysqlEngine) GetUrl(id int) (string, error) {
    // todo
    return nil, nil
}

func (my *MysqlEngine) StoreUrl(url string) (int, error) {
    // todo
    return nil, nil
}