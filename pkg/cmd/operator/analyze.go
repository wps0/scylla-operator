package operator

import (
	"context"
	"errors"
	"fmt"

	"github.com/scylladb/scylla-operator/pkg/analyze"
	"github.com/scylladb/scylla-operator/pkg/analyze/snapshot"
	scyllaversioned "github.com/scylladb/scylla-operator/pkg/client/scylla/clientset/versioned"
	"github.com/scylladb/scylla-operator/pkg/genericclioptions"
	scyllaScheme "github.com/scylladb/scylla-operator/pkg/scheme"
	"github.com/scylladb/scylla-operator/pkg/version"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	apierrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/kubernetes"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog/v2"
	"k8s.io/kubectl/pkg/util/templates"
	"os"
)

var (
	analyzeLongDescription = templates.LongDesc(`
		analyze traverses associated resources and diagnoses common issues.

		This command is experimental and subject to change without notice.
	`)
)

type AnalyzeOptions struct {
	genericclioptions.ClientConfig

	ArchivePath           string
	DisableStrictEncoding bool

	kubeClient   *kubernetes.Clientset
	scyllaClient *scyllaversioned.Clientset
}

func NewAnalyzeOptions(streams genericclioptions.IOStreams) *AnalyzeOptions {
	return &AnalyzeOptions{
		ClientConfig: genericclioptions.NewClientConfig("scylla-operator-analyze"),
	}
}

func NewAnalyzeCmd(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewAnalyzeOptions(streams)

	cmd := &cobra.Command{
		Use:   "analyze",
		Short: "Run scylla-operator analyze.",
		Long:  analyzeLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := o.Validate()
			if err != nil {
				return err
			}

			err = o.Complete()
			if err != nil {
				return err
			}

			err = o.Run(streams, cmd)
			if err != nil {
				return err
			}

			return nil
		},
	}

	o.AddFlags(cmd)

	return cmd
}

func (o *AnalyzeOptions) AddFlags(cmd *cobra.Command) {
	o.ClientConfig.AddFlags(cmd)

	cmd.Flags().StringVarP(&o.ArchivePath, "archive-path", "", o.ArchivePath, "Path to a compressed must-gather archive or a directory having must-gather structure")
	cmd.Flags().BoolVarP(&o.DisableStrictEncoding, "disable-strict-encoding", "", false, "Disable strict mode in deserializer used for parsing archive")
}

func (o *AnalyzeOptions) Validate() error {
	var errs []error

	if len(o.ArchivePath) > 0 {
		_, err := os.Stat(o.ArchivePath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				errs = append(errs, fmt.Errorf("archive path %q does not exist", o.ArchivePath))
			} else {
				errs = append(errs, fmt.Errorf("can't stat archive path %q", o.ArchivePath))
			}
		}
	} else {
		errs = append(errs, o.ClientConfig.Validate())
	}

	if len(o.Kubeconfig) != 0 && len(o.ArchivePath) != 0 {
		errs = append(errs, fmt.Errorf("kubeconfig and archive-path can't both be set"))
	}

	return apierrors.NewAggregate(errs)
}

func (o *AnalyzeOptions) Complete() error {
	if len(o.ArchivePath) != 0 {
		return nil
	}

	err := o.ClientConfig.Complete()
	if err != nil {
		return err
	}

	o.kubeClient, err = kubernetes.NewForConfig(o.ProtoConfig)
	if err != nil {
		return fmt.Errorf("can't build kubernetes clientset: %w", err)
	}
	o.scyllaClient, err = scyllaversioned.NewForConfig(o.RestConfig)
	if err != nil {
		return fmt.Errorf("can't build scylla clientset: %w", err)
	}

	return nil
}

func (o *AnalyzeOptions) Run(streams genericclioptions.IOStreams, cmd *cobra.Command) error {
	klog.Infof("%s version %s", cmd.Name(), version.Get())
	cliflag.PrintFlags(cmd.Flags())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var clusterSnapshot snapshot.Snapshot
	var err error
	if len(o.ArchivePath) > 0 {
		var codecFactory serializer.CodecFactory
		if o.DisableStrictEncoding {
			codecFactory = serializer.NewCodecFactory(scyllaScheme.Scheme, serializer.DisableStrict)
		} else {
			codecFactory = serializer.NewCodecFactory(scyllaScheme.Scheme, serializer.EnableStrict)
		}
		fs := os.DirFS(o.ArchivePath)
		clusterSnapshot, err = snapshot.NewSnapshotFromFS(fs, codecFactory.UniversalDeserializer())
		if err != nil {
			return fmt.Errorf("can't build data source from must-gather: %w", err)
		}
	} else {
		clusterSnapshot, err = snapshot.NewSnapshotFromListers(ctx, o.kubeClient, o.scyllaClient)
		if err != nil {
			return fmt.Errorf("can't build data source from clients: %w", err)
		}
	}

	return analyze.Analyze(ctx, clusterSnapshot)
}
