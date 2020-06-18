package fuse

import (
	"context"
	"sync"

	"github.com/FleekHQ/space-poc/core/spacefs"

	"github.com/FleekHQ/space-poc/config"
	"github.com/FleekHQ/space-poc/core/libfuse"
	"github.com/FleekHQ/space-poc/core/store"
	"github.com/FleekHQ/space-poc/log"
)

// Controller is the space domain controller for managing the libfuse VFS.
// It is used by the grpc server and app/daemon generally
type Controller struct {
	cfg       config.Config
	vfs       *libfuse.VFS
	store     store.Store
	isServed  bool
	mountLock sync.RWMutex
}

func NewController(
	ctx context.Context,
	cfg config.Config,
	store store.Store,
	sfs *spacefs.SpaceFS,
) *Controller {
	mountPath := cfg.GetString(config.FuseMountPath, "~/")
	vfs := libfuse.NewVFileSystem(ctx, mountPath, sfs)
	return &Controller{
		cfg:       cfg,
		store:     store,
		vfs:       vfs,
		isServed:  false,
		mountLock: sync.RWMutex{},
	}
}

// ShouldMount check the store and config to determine if the libfuse drive was previously mounted
func (s *Controller) ShouldMount() bool {
	mountFuseDrive, err := s.store.Get([]byte(config.MountFuseDrive))
	if err == nil {
		return string(mountFuseDrive) == "true"
	} else {
		log.Debug("Error fetching mountFuseDrive state: %s\n", err.Error())
	}

	return s.cfg.GetString(config.MountFuseDrive, "false") == "true"
}

// Mount mounts the vfs drive and immediately serves the handler.
// It starts the Fuse Server in the background
func (s *Controller) Mount() error {
	s.mountLock.Lock()
	defer s.mountLock.Unlock()

	if s.vfs.IsMounted() {
		return nil
	}

	if err := s.vfs.Mount(
		s.cfg.GetString(config.FuseDriveName, "FleekSpace"),
	); err != nil {
		return err
	}

	// persist mount state to store to trigger remount on restart
	if err := s.store.Set([]byte(config.MountFuseDrive), []byte("true")); err != nil {
		return err
	}

	s.serve()
	return nil
}

func (s *Controller) serve() {
	if s.isServed {
		return
	}

	go func() {
		s.isServed = true
		defer func() {
			s.isServed = false
		}()

		// this blocks and unblocks when vfs.Unmount() is called
		// or some external thing happens like user unmounting the drive
		err := s.vfs.Serve()
		if err != nil {
			log.Error("error ending fuse server", err)
		}
	}()
}

func (s *Controller) IsMounted() bool {
	s.mountLock.RLock()
	defer s.mountLock.RUnlock()
	return s.vfs.IsMounted()
}

func (s *Controller) Unmount() error {
	s.mountLock.Lock()
	defer s.mountLock.Unlock()
	if !s.vfs.IsMounted() {
		return nil
	}

	// persist mount state to store to trigger remount on restart
	if err := s.store.Set([]byte(config.MountFuseDrive), []byte("false")); err != nil {
		return err
	}

	return s.vfs.Unmount()
}