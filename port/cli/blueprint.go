package cli

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *PortClient) ReadBlueprint(ctx context.Context, id string) (*Blueprint, int, error) {
	pb := &PortBody{}
	url := "v1/blueprints/{identifier}"
	resp, err := c.Client.R().
		SetContext(ctx).
		SetHeader("Accept", "application/json").
		SetQueryParam("exclude_calculated_properties", "true").
		SetResult(pb).
		SetPathParam("identifier", id).
		Get(url)
	if err != nil {
		return nil, resp.StatusCode(), err
	}
	if !pb.OK {
		return nil, resp.StatusCode(), fmt.Errorf("failed to read blueprint, got: %s", resp.Body())
	}
	return &pb.Blueprint, resp.StatusCode(), nil
}

func (c *PortClient) CreateBlueprint(ctx context.Context, b *Blueprint) (*Blueprint, error) {
	url := "v1/blueprints"
	resp, err := c.Client.R().
		SetBody(b).
		SetContext(ctx).
		Post(url)
	if err != nil {
		return nil, err
	}
	var pb PortBody
	err = json.Unmarshal(resp.Body(), &pb)
	if err != nil {
		return nil, err
	}
	if !pb.OK {
		return nil, fmt.Errorf("failed to create blueprint, got: %s", resp.Body())
	}
	return &pb.Blueprint, nil
}

func (c *PortClient) UpdateBlueprint(ctx context.Context, b *Blueprint, id string) (*Blueprint, error) {
	url := "v1/blueprints/{identifier}"
	resp, err := c.Client.R().
		SetBody(b).
		SetContext(ctx).
		SetPathParam("identifier", id).
		Put(url)
	if err != nil {
		return nil, err
	}
	var pb PortBody
	err = json.Unmarshal(resp.Body(), &pb)
	if err != nil {
		return nil, err
	}
	if !pb.OK {
		return nil, fmt.Errorf("failed to create blueprint, got: %s", resp.Body())
	}
	return &pb.Blueprint, nil
}

func (c *PortClient) DeleteBlueprint(ctx context.Context, id string) error {
	url := "v1/blueprints/{identifier}"
	resp, err := c.Client.R().
		SetContext(ctx).
		SetHeader("Accept", "application/json").
		SetPathParam("identifier", id).
		Delete(url)
	if err != nil {
		return err
	}
	responseBody := make(map[string]interface{})
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		return err
	}
	if !(responseBody["ok"].(bool)) {
		return fmt.Errorf("failed to delete blueprint. got:\n%s", string(resp.Body()))
	}
	return nil
}
