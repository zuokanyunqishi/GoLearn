package service

import "GoLearn/chat/interfaces"

func (c *Client) CreateConnection(openWorkPool bool) interfaces.Connection {
	return c.ClientServe.CreateConnection(openWorkPool)
}
