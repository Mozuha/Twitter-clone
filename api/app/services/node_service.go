package services

import (
	"app/ent"
	"context"
)

type nodeService struct {
	client *ent.Client
}

func (n *nodeService) Node(ctx context.Context, id int) (ent.Noder, error) {
	return n.client.Noder(ctx, id)
}

func (n *nodeService) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return n.client.Noders(ctx, ids)
}
