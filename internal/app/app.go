package app

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"

	"github.com/s02190058/reverse-proxy-cache-client/internal/grpc"
	thumbnailpb "github.com/s02190058/reverse-proxy-cache/gen/go/thumbnail/v1"
	"golang.org/x/exp/slog"
)

const (
	perm = 0644
)

var (
	ErrInvalidURL = errors.New("invalid youtube url")
	ErrInternal   = errors.New("internal error")
)

// App is an engine.
type App struct {
	prompt     string
	dir        string
	async      bool
	l          *slog.Logger
	grpcClient thumbnailpb.ThumbnailServiceClient
	re         *regexp.Regexp
	closeFunc  func()
}

// New creates an rpc-cli instance.
func New(host, port string, dir string, async bool) (*App, error) {
	l := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	grpcClient, closeFunc, err := grpc.NewClient(host, port)
	if err != nil {
		return nil, err
	}

	return &App{
		prompt:     fmt.Sprintf("%s:%s> ", host, port),
		dir:        dir,
		async:      async,
		l:          l,
		grpcClient: grpcClient,
		re:         regexp.MustCompile(`^(https?://)?www\.youtube\.com/watch\?v=[a-zA-Z0-9_-]{11}`),
		closeFunc: func() {
			if err = closeFunc(); err != nil {
				l.Error("can't close conn", slogErr(err))
			}
		},
	}, nil
}

// Run launches the client.
func (a *App) Run() {
	defer a.closeFunc()

	scanner := bufio.NewScanner(os.Stdout)
	fmt.Print(a.prompt)
	for scanner.Scan() {
		videoURL := scanner.Text()

		if a.async {
			go func(videoURL string) {
				if err := a.handle(videoURL); err != nil {
					a.l.Error("request processing error", slogErr(err))
				}
			}(videoURL)
		} else {
			if err := a.handle(videoURL); err != nil {
				a.l.Error("request processing error", slogErr(err))
				if !errors.Is(err, ErrInvalidURL) {
					break
				}
			}
		}

		fmt.Print(a.prompt)
	}
	if err := scanner.Err(); err != nil {
		a.l.Error("Scanner.Err", slog.String("error", err.Error()))
	}
}

func (a *App) handle(videoURL string) error {
	if !a.re.MatchString(videoURL) {
		return ErrInvalidURL
	}

	videoID, err := videoIDFromURL(videoURL)
	if err != nil {
		return err
	}

	resp, err := a.grpcClient.Download(context.Background(), &thumbnailpb.DownloadThumbnailRequest{
		VideoID: videoID,
		Type:    thumbnailpb.ThumbnailType_DEFAULT,
	})
	if err != nil {
		return err
	}

	filename := videoID + ".jpg"
	path := filepath.Join(a.dir, filename)
	if err = os.WriteFile(path, resp.Image, perm); err != nil {
		return err
	}

	return nil
}

// slogErr generates slog attribute for an error
func slogErr(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

// videoIDFromURL pulls video id form the url.
func videoIDFromURL(videoURL string) (string, error) {
	u, err := url.Parse(videoURL)
	if err != nil {
		return "", ErrInternal
	}

	return u.Query().Get("v"), nil
}
