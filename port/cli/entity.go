package cli

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *PortClient) ReadEntity(ctx context.Context, id string, blueprint string) (*Entity, error) {
	url := "v1/blueprints/{blueprint}/entities/{identifier}"
	resp, err := c.Client.R().
		SetHeader("Accept", "application/json").
		SetQueryParam("exclude_calculated_properties", "true").
		SetPathParam(("blueprint"), blueprint).
		SetPathParam("identifier", id).
		Get(url)
	if err != nil {
		return nil, err
	}
	var pb PortBody
	err = json.Unmarshal(resp.Body(), &pb)
	if err != nil {
		return nil, err
	}
	return &pb.Entity, nil
}

func (c *PortClient) CreateEntity(ctx context.Context, e *Entity) (*Entity, error) {
	url := "v1/blueprints/{blueprint}/entities"
	pb := &PortBody{}
	resp, err := c.Client.R().
		SetBody(e).
		SetPathParam(("blueprint"), e.Blueprint).
		SetQueryParam("upsert", "true").
		SetResult(&pb).
		Post(url)
	if err != nil {
		return nil, err
	}
	if !pb.OK {
		return nil, fmt.Errorf("failed to create entity, got: %s", resp.Body())
	}
	return &pb.Entity, nil
}

func (c *PortClient) DeleteEntity(ctx context.Context, id string, blueprint string) error {
	url := "v1/blueprints/{blueprint}/entities/{identifier}"
	pb := &PortBody{}
	resp, err := c.Client.R().
		SetHeader("Accept", "application/json").
		SetPathParam("blueprint", blueprint).
		SetPathParam("identifier", id).
		SetResult(pb).
		Delete(url)
	if err != nil {
		return err
	}
	if !pb.OK {
		return fmt.Errorf("failed to delete entity, got: %s", resp.Body())
	}
	return nil
}
