package services

import (
	"context"
	"fmt"
)

type ScanService struct{}

func (service *ScanService) Run(ctx context.Context) error {
	fmt.Println("test")
	return nil
}
