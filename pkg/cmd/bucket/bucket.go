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
			bucketExpression := args[0]
			newBucket := bucket.NewBucket(bucketExpression, opts.Path)
			bucketRepository := factory.BucketRepository
			err := bucketRepository.Save(newBucket)
			if err != nil {
				return err
			}
			fmt.Printf("Added %s to bucket.\n", newBucket.Id())
			return nil
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
			bucketId := args[0]
			bucketRepository := factory.BucketRepository
			err := bucketRepository.Remove(bucketId)
			if err != nil {
				return err
			}
			fmt.Printf("Removed %s to bucket.\n", bucketId)
			return nil
		},
	}
	return cmd
}

func NewCmdBucketList(factory *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			bucketRepository := factory.BucketRepository
			buckets, err := bucketRepository.Read()
			if err != nil {
				return err
			}

			if len(buckets) == 0 {
				fmt.Println("No buckets.")
			} else {
				for _, b := range buckets {
					fmt.Println(b.String())
				}
				fmt.Printf("\nFound %d buckets.\n", len(buckets))
			}
			return nil
		},
	}
	return cmd
}
