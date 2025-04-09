package main

import (
	"context"
	"log"
	"log/slog"
	
	_ "github.com/whosonfirst/go-whosonfirst-iterate-organization"
	_ "github.com/whosonfirst/go-writer-featurecollection/v3"
	
	"github.com/whosonfirst/go-whosonfirst-iterwriter/app/iterwriter"
)

func main() {

	ctx := context.Background()
	err := iterwriter.Run(ctx, slog.Default())

	if err != nil {
		log.Fatalf("Failed to run iterwriter, %v", err)
	}

}
