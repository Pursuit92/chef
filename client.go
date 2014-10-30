package chef

import "fmt"

type ClientService struct {
	client *Client
}

// Client represents the native Go version of the deserialized Client type
type ClientObj struct {
	Name        string `json:"name"`
	ClientName  string `json:"clientname"`
	OrgName     string `json:"orgname"`
	Validator   bool   `json:"validator"`
	Certificate string `json:"certificate"`
	Admin       bool   `json:"admin"`
	PrivateKey  string `json:"private_key"`
}

type ClientResult struct {
	Uri string `json:"uri"`
}

// NewClient is the Client constructor method
func NewClientObj(name string) (clients ClientObj) {
	clients = ClientObj{
		Name: name,
	}
	return
}

// List lists the nodes in the Chef server.
//
// Chef API docs: http://docs.opscode.com/api_chef_server.html#id25
func (e *ClientService) List() (data map[string]string, err error) {
	err = e.client.magicRequestDecoder("GET", "clients", nil, &data)
	return
}

// Get gets a node from the Chef server.
//
// Chef API docs: http://docs.opscode.com/api_chef_server.html#id28
func (e *ClientService) Get(name string) (client ClientObj, err error) {
	url := fmt.Sprintf("nodes/%s", name)
	err = e.client.magicRequestDecoder("GET", url, nil, &client)
	return
}

// Post creates a Client on the chef server
//
// Chef API docs: https://docs.getchef.com/api_chef_server.html#id39
func (e *ClientService) Post(client ClientObj) (data *ClientResult, err error) {
	body, err := JSONReader(client)
	if err != nil {
		return
	}

	err = e.client.magicRequestDecoder("POST", "clients", body, &data)
	return
}

// Put updates a node on the Chef server.
//
// Chef API docs: http://docs.getchef.com/api_chef_server.html#id42
func (e *ClientService) Put(n ClientObj) (node ClientObj, err error) {
	url := fmt.Sprintf("clients/%s", n.Name)
	body, err := JSONReader(n)
	if err != nil {
		return
	}

	err = e.client.magicRequestDecoder("PUT", url, body, &node)
	return
}

// Delete removes a node on the Chef server
//
// Chef API docs: https://docs.getchef.com/api_chef_server.html#id40
func (e *ClientService) Delete(name string) (err error) {
	err = e.client.magicRequestDecoder("DELETE", "clients/"+name, nil, nil)
	return
}
