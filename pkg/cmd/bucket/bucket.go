package bucket

import (
	"fmt"
	"github.com/namhyun-gu/brick/internal/bucket"
	"github.com/namhyun-gu/brick/pkg/cmdutil"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	Path string
}

func NewCmdBucket(factory *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "bucket <command>",
	}

	cmd.AddCommand(NewCmdBucketAdd(factory))
	cmd.AddCommand(NewCmdBucketRemove(factory))
	cmd.AddCommand(NewCmdBucketList(factory))
	return cmd
}

func NewCmdBucketAdd(factory *cmdutil.Factory) *cobra.Command {
	opts := &AddOptions{}
	cmd := &cobra.Command{
		Use:  "add [owner:repo@branch]",
		Args: cobra.MinimumNArgs(1),
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

			bucketExpression := args[0]
			newBucket := bucket.NewBucket(bucketExpression, opts.Path)
			err = repository.Save(newBucket)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(factory.IO.Out, "Added %s to bucket.\n", newBucket.Id())
			return err
		},
	}
	cmd.Flags().StringVarP(&opts.Path, "path", "p", "", "Bucket root path")
	return cmd
}

func NewCmdBucketRemove(factory *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "remove [owner:repo]",
		Args: cobra.MinimumNArgs(1),
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

			bucketId := args[0]
			err = repository.Remove(bucketId)
			if err != nil {
				return err
			}

			fmt.Fprintf(factory.IO.Out, "Removed %s to bucket.\n", bucketId)
			return err
		},
	}
	return cmd
}

func NewCmdBucketList(factory *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
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

			if len(buckets) == 0 {
				fmt.Fprintf(factory.IO.Out, "No buckets.")
			} else {
				for _, b := range buckets {
					fmt.Fprintf(factory.IO.Out, b.String())
				}
				fmt.Fprintf(factory.IO.Out, "\nFound %d buckets.\n", len(buckets))
			}
			return nil
		},
	}
	return cmd
}
