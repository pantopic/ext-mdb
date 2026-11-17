package wazero_lmdb

type Option func(*hostModule)

func WithDirLocal(dir string) Option {
	return func(h *hostModule) {
		h.dirLocal = dir
	}
}
