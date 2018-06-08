package goconf

import (
	"testing"
	"path/filepath"
)

type Database struct {
	Driver string `json:"driver"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
}

type ConfigDemo struct {
	Database Database `json:"database"`
}

func TestNewConfig(t *testing.T) {
	_, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewConfig("./testdata2/")
	if err == nil {
		t.Fatalf("testdata2 must return not found error")
	}
}

func TestConfig_Get(t *testing.T) {
	conf, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}

	if ret, err := conf.Get("dev.name.last"); err != nil || ret.String() != "Anderson" {
		t.Errorf("get key %s test error", "dev.name.last")
	}

	if ret, err := conf.Get("dev.age"); err != nil || ret.Int() != 37 {
		t.Errorf("get key %s test error", "dev.age")
	}

	if ret, err := conf.Get("dev.childen"); err != nil || ret.IsArray() {
		t.Errorf("get key %s test error", "dev.childen")
	}

	if ret, err := conf.Get("dev.friends.1.age"); err != nil || ret.Int() != 68 {
		t.Errorf("get key %s test error", "dev.friends.1.age")
	}

	if ret, err := conf.Get("prod.widget.window.width"); err != nil || ret.Int() != 500 {
		t.Errorf("get key %s test error", "prod.window.width")
	}

	if ret, err := conf.Get("prod.image2.alignment"); err != nil || ret.String() != "" {
		t.Errorf("get key %s value must empty", "prod.image2.alignment")
	}

	if _, err := conf.Get("dev2.notfound"); err == nil {
		t.Errorf("get key %s must return err", "dev2.notfound")
	}
}

func TestConfig_MustGet(t *testing.T) {
	conf, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}

	if conf.MustGet("dev.name.last").String() != "Anderson" {
		t.Errorf("must get key %s test error", "dev.name.last")
	}

	if conf.MustGet("dev.age").Int() != 37 {
		t.Errorf("must get key %s test error", "dev.age")
	}

	if conf.MustGet("dev.childen").IsArray() {
		t.Errorf("must get key %s test error", "dev.childen")
	}

	if conf.MustGet("dev.friends.1.age").Int() != 68 {
		t.Errorf("must get key %s test error", "dev.friends.1.age")
	}

	if conf.MustGet("prod.widget.window.width").Int() != 500 {
		t.Errorf("must get key %s test error", "prod.window.width")
	}

	if conf.MustGet("prod.image2.alignment").String() != "" {
		t.Errorf("must get key %s value must empty", "prod.image2.alignment")
	}

	if conf.MustGet("dev.name.notfound").String() != "" {
		t.Errorf("must get key %s must return empty", "dev.name.notfound")
	}

	if conf.MustGet("dev.name.notfound2").Int() != 0 {
		t.Errorf("must get key %s must return 0", "dev.name.notfound2")
	}

	if conf.MustGet("dev2.notfound").String() != "" {
		t.Errorf("must get key %s must return empty", "dev2.notfound")
	}

	if conf.MustGet("dev2.notfound2").Int() != 0 {
		t.Errorf("must get key %s must return 0", "dev2.notfound2")
	}
}

func TestConfig_Cache(t *testing.T) {
	conf, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}

	name := conf.MustGet("dev.name.first").String()

	file := filepath.Join(conf.Path, "dev"+conf.Ext)

	if !conf.cache[file].IsObject() {
		t.Fatalf("config cache miss")
	}

	cacheName := conf.cache[file].Get("name.first").String()

	if name != "" && name != cacheName {
		t.Fatalf("cache value not match [%s:%s]", name, cacheName)
	}
}

func TestConfig_Unmarshal(t *testing.T) {
	conf, err := NewConfig("./testdata/")
	if err != nil {
		t.Fatal(err)
	}

	app := &ConfigDemo{}
	err = conf.Unmarshal("json5", app)

	if err != nil {
		t.Fatal(err)
	}

	if app.Database.Host != "localhost" {
		t.Fatalf("Unmarshal struct host must return %s", "localhost")
	}

	if app.Database.Port != 3306 {
		t.Fatalf("Unmarshal struct port must return %d", 3306)
	}
	//fmt.Println(app)
}
