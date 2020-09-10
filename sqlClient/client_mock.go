package sqlClient

type clientMock struct {
	mocks map[string]Mock
}

type Mock struct {
	Query string
	Args  []interface{}
	Error error
}

func AddMock(m Mock) {
	client, ok := dbClient.(*clientMock)
	if ok {
		if client.mocks == nil {
			client.mocks = make(map[string]Mock, 0)
		}
		client.mocks[m.Query] = m
	}
}

func (c *clientMock) Query(query string, args ...interface{}) (rows, error) {
	//todo
	return nil, nil
}
