package main

type Ingester struct {
	PostgresClient *PostgresClient
	Conf           *Configuration
}

func (i *Ingester) Start() {

}

func NewIngester(conf *Configuration) *Ingester {
	i := new(Ingester)
	i.Conf = conf

	// Postgres
	/*
		i.PostgresClient = NewPostgresClient(i.Conf.PGHost, i.Conf.PGPort,
			i.Conf.PGUser, i.Conf.PGPassword, i.Conf.PGDbname)
	*/

	return i
}
