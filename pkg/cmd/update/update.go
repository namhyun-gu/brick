package update

import (
	"fmt"
	"github.com/namhyun-gu/brick/internal/bucket"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
	"path/filepath"
)

func NewCmdUpdate(factory *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "update",
		RunE: func(cmd *cobra.Command, args []string) error {
			repository := factory.BucketRepository
			buckets, err := repository.Read()
			if err != nil {
				return err
			}

			if buckets == nil {
				buckets, err = cmdutil.InitBucket(repository)
				if err != nil {
					return err
				}
			}

			for _, bucketObj := range buckets {
				fmt.Printf("Updating %s...", bucketObj.Id())
				err := bucket.WriteCache(factory.Client, filepath.Dir(factory.Executable), bucketObj)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	return cmd
}
