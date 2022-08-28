package img

// NewGetter constructs a Getter based on the GetterType argument.
func NewGetter(t GetterType) Getter {
	switch t {
	case LocalGetter:
		return NewLocalGetter()
	}
	return NewLocalGetter()
}

// NewRepository constructs a Repository based on the RepositoryType argument.
func NewRepository(t RepositoryType) Repository {
	switch t {
	case InMemoryRepository:
		return NewInMemoryRepository(DefaultImages)
	}
	return NewInMemoryRepository(DefaultImages)
}

// NewService constructs a Service.
func NewService(
	repository Repository,
	uploader Uploader,
	baseURL string,
	maxWidth int,
	maxHeight int,
) *Service {
	return &Service{
		repository: repository,
		uploader:   uploader,
		baseURL:    baseURL,
		maxWidth:   maxWidth,
		maxHeight:  maxHeight,
	}
}

// NewUploader constructs an UploaderType based on the Type argument.
func NewUploader(t UploaderType, basePath string) Uploader {
	switch t {
	case LocalUploader:
		return NewLocalUploader(basePath)
	case NoopUploader:
		return NewNoopUploader()
	}
	return NewLocalUploader(basePath)
}
