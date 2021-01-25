package v1

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/marshaler"
	"github.com/eko/gocache/store"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evcache"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evsmtp"
)

type CacheInterface interface {
	Checker(checker evsmtp.Checker) evsmtp.Checker
	Validator(validator ev.Validator) ev.Validator
}

func NewMemcachedCache(servers []string, optChecker, optValidator *store.Options) CacheInterface {
	memcached := memcache.New(servers...)

	memcachedStore := store.NewMemcache(memcached, nil)
	marshaller := marshaler.New(memcachedStore)

	return &Cache{
		marshaller:   marshaller,
		optChecker:   optChecker,
		optValidator: optValidator,
	}
}

type Cache struct {
	marshaller   *marshaler.Marshaler
	optChecker   *store.Options
	optValidator *store.Options
}

func (m *Cache) Checker(checker evsmtp.Checker) evsmtp.Checker {
	return evsmtp.NewCheckerCacheRandomRCPT(
		checker.(evsmtp.CheckerWithRandomRCPT),
		evcache.NewCacheMarshaller(
			m.marshaller,
			func() interface{} {
				return new([]error)
			},
			m.optChecker,
		),
		evsmtp.DefaultRandomCacheKeyGetter,
	)
}
func (m *Cache) Validator(validator ev.Validator) ev.Validator {
	return ev.NewCacheDecorator(
		validator,
		evcache.NewCacheMarshaller(
			m.marshaller,
			func() interface{} {
				return new(ev.ValidationResult)
			},
			m.optValidator,
		),
		ev.EmailCacheKeyGetter,
	)
}

func NewRistrettoCache(cfgChecker *ristretto.Config, cfgValidator *ristretto.Config, optChecker, optValidator *store.Options) CacheInterface {
	return &RistrettoCache{
		marshallerChecker:   getRistrettoMarshaller(cfgChecker),
		marshallerValidator: getRistrettoMarshaller(cfgValidator),
		optChecker:          optChecker,
		optValidator:        optValidator,
	}
}

func getRistrettoMarshaller(cfg *ristretto.Config) *marshaler.Marshaler {
	ristrettoCache, err := ristretto.NewCache(cfg)
	if err != nil {
		panic(err)
	}

	ristrettoStore := store.NewRistretto(ristrettoCache, nil)
	return marshaler.New(ristrettoStore)
}

type RistrettoCache struct {
	marshallerChecker   *marshaler.Marshaler
	marshallerValidator *marshaler.Marshaler
	optChecker          *store.Options
	optValidator        *store.Options
}

func (m *RistrettoCache) Checker(checker evsmtp.Checker) evsmtp.Checker {
	return evsmtp.NewCheckerCacheRandomRCPT(
		checker.(evsmtp.CheckerWithRandomRCPT),
		evcache.NewCacheMarshaller(
			m.marshallerChecker,
			func() interface{} {
				return new([]error)
			},
			m.optChecker,
		),
		evsmtp.DefaultRandomCacheKeyGetter,
	)
}
func (m *RistrettoCache) Validator(validator ev.Validator) ev.Validator {
	return ev.NewCacheDecorator(
		validator,
		evcache.NewCacheMarshaller(
			m.marshallerValidator,
			func() interface{} {
				return new(ev.ValidationResult)
			},
			m.optValidator,
		),
		ev.EmailCacheKeyGetter,
	)
}
