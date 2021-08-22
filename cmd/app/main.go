package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.n.xiaomi.com/op-basic/event_server/libs/config"
	mylog "git.n.xiaomi.com/op-basic/event_server/libs/log"
	"git.n.xiaomi.com/op-basic/event_server/users"
	"git.n.xiaomi.com/op-basic/event_server/users/auth"
	httpserver "git.n.xiaomi.com/op-basic/event_server/users/auth/transport/http"
	"git.n.xiaomi.com/op-basic/event_server/users/db"

	"github.com/spf13/cobra"
)

type HostServerOptions struct {
	ListenAddr string
	Debug      bool
	Env        string
}

type ServerConfig struct {
	DB  db.Config
	Log mylog.Config
	// Transport TransportConfig
	// Redis     db.RedisConfig
}

func init() {
	rootCmd.AddCommand(ServeCommand())
}

func ServeCommand() *cobra.Command {
	o := &HostServerOptions{}
	cmd := cobra.Command{
		Use:   "web",
		Short: "start the host-apply host creation server",
		Run: func(cmd *cobra.Command, args []string) {
			startUserServer(o)
		},
	}
	cmd.Flags().StringVar(&o.ListenAddr, "listen-addr", "127.0.0.1:5000", "listen address. for example: --listen-addr 127.0.0.1:5000")
	cmd.Flags().BoolVar(&o.Debug, "debug", false, "enable debug-level logging. for example: --debug=true")
	cmd.Flags().StringVar(&o.Env, "config", "staging", "config environment, does not include extension. for example: staging, product")
	return &cmd
}

func startUserServer(opt *HostServerOptions) {
	cfg := ServerConfig{}
	err := config.Init(opt.Env, &cfg)
	if err != nil {
		panic(fmt.Sprintf("init config failed: %s\n", err.Error()))
	}

	//init log
	log, err := mylog.NewLogger(cfg.Log)
	if err != nil {
		panic(fmt.Sprintf("init logger failed: %s\n", err.Error()))
	}
	repo, err := db.NewPostgres(cfg.DB)
	if err != nil {
		log.Fatal("failed to init db")
	}
	pwdSec := users.NewPasswordSecurity()
	authSvc := auth.New(repo, pwdSec)
	srv := httpserver.NewHTTPServer(authSvc, log)
	server := http.Server{
		Addr:    opt.ListenAddr,
		Handler: srv,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Shutdown Server ...", <-quit)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 2 seconds.
	select {
	case sig := <-quit:
		log.Info("received signal %s", sig)
	case <-ctx.Done():
		log.Info("timeout of 2 seconds.")
	}
	log.Info("Server exiting")
}
