package envloader

import (
	"context"
	"fmt"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"

	etcd "go.etcd.io/etcd/client/v3"
)

const (
	Yaml = "application/x-yaml"
	JSON = "application/json"
)

type (
	EnvLoaderOption func(*EnvLoader) error
	EnvLoader       struct {
		File *FileLoader
		Etcd *EtcdLoader
	}

	FileLoader struct {
		Path string
	}

	EtcdLoader struct {
		Prefix    string
		Endpoints []string
	}
)

func New(opts ...EnvLoaderOption) (*EnvLoader, error) {
	envLoader := &EnvLoader{}
	for _, opt := range opts {
		err := opt(envLoader)
		if err != nil {
			return nil, err
		}
	}

	return envLoader, nil
}

func (e *EnvLoader) Validate(config interface{}) error {
	rt := reflect.TypeOf(config)
	if rt.Kind() != reflect.Ptr || rt.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("config must be a pointer to struct")
	}

	if e.File != nil {
		if !path.IsAbs(e.File.Path) {
			return fmt.Errorf("filepath must be absolute path")
		}

		if !strings.Contains(e.File.Path, ".env") {
			return fmt.Errorf("filepath must be .env file")
		}
	}

	return nil
}

func WithFile(filepath string) EnvLoaderOption {
	return func(e *EnvLoader) error {
		e.File = &FileLoader{Path: filepath}
		return nil
	}
}

func WithEtcd(prefix string, endpoints []string) EnvLoaderOption {
	return func(e *EnvLoader) error {
		e.Etcd = &EtcdLoader{Prefix: prefix, Endpoints: endpoints}
		return nil
	}
}

func Load(config interface{}, opts ...EnvLoaderOption) error {
	options, err := New(opts...)
	if err != nil {
		return err
	}

	err = options.Validate(config)
	if err != nil {
		return err
	}

	switch {
	case options.File != nil:
		err = LoadFromFile(options.File.Path, config)
	case options.Etcd != nil:
		err = LoadFromEtcd(options.Etcd.Prefix, options.Etcd.Endpoints, config)
	default:
		err = LoadFromEnv(config)
	}

	return err
}

func LoadFromFile(filepath string, config interface{}) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

    configMap := make(map[string][]byte)
	for _, line := range strings.Split(string(file), "\n") {
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		configMap[key] = []byte(value)
	}

	return setConfig(config, configMap)
}

func LoadFromEtcd(prefix string, endpoints []string, config interface{}) error {
	client, err := etcd.New(etcd.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	defer client.Close()

	resp, err := client.Get(context.Background(), prefix, etcd.WithPrefix())
	if err != nil {
		return err
	}

	if len(resp.Kvs) == 0 {
		return nil
	}

	configMap := make(map[string][]byte)
	for _, kv := range resp.Kvs {
		keyWithoutPrefix := strings.TrimPrefix(string(kv.Key), fmt.Sprintf("%s/", prefix))
		configMap[keyWithoutPrefix] = kv.Value
	}

	return setConfig(config, configMap)
}

func LoadFromEnv(config interface{}) error {
    configMap := make(map[string][]byte)
    for _, env := range os.Environ() {
        parts := strings.Split(env, "=")
        if len(parts) != 2 {
            continue
        }

        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])

        configMap[key] = []byte(value)
    }

    return setConfig(config, configMap)
}

func setConfig(config interface{}, configMap map[string][]byte) error {
	rv := reflect.ValueOf(config)
	rt := reflect.TypeOf(config)

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		tag := rt.Field(i).Tag.Get("env")

		if field.Kind() == reflect.Struct {
			err := setConfig(field.Addr().Interface(), configMap)
			if err != nil {
				return err
			}
		}

		if tag == "" {
			continue
		}

		fieldValue, ok := configMap[tag]
		if !ok {
			continue
		}

		fieldValueParsed, err := parseString(string(fieldValue), field.Kind())
		if err != nil {
			return err
		}

        if isZero(fieldValueParsed) {
            continue
        }

		switch field.Kind() {
		case reflect.String:
			field.SetString(fieldValueParsed.(string))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(fieldValueParsed.(int64))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetUint(fieldValueParsed.(uint64))
		case reflect.Float32, reflect.Float64:
			field.SetFloat(fieldValueParsed.(float64))
		case reflect.Bool:
			field.SetBool(fieldValueParsed.(bool))
		default:
			return fmt.Errorf("unsupported type %s", field.Kind().String())
		}
	}

	return nil
}

func isZero(value interface{}) bool {
    return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

func parseString(value string, kind reflect.Kind) (interface{}, error) {
	switch kind {
	case reflect.String:
		return value, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.ParseInt(value, 10, 64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.ParseUint(value, 10, 64)
	case reflect.Float32, reflect.Float64:
		return strconv.ParseFloat(value, 64)
	case reflect.Bool:
		return strconv.ParseBool(value)
	default:
		return nil, fmt.Errorf("unsupported type %s", kind.String())
	}
}
