package persistentconfig

import (
	"testing"

	"github.com/spf13/afero"
	"gopkg.in/ini.v1"
)

func TestPersistentConfigEnsureConfigFiles(t *testing.T) {
	client := ClientPersistentConfig{
		Fs: afero.NewMemMapFs(), // Using an in-memory file system
	}
	err := client.ensureConfigFiles()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if files are created
	exists, err := afero.Exists(client.Fs, client.configFile)
	if !exists || err != nil {
		t.Fatalf("Expected config file to exist")
	}

	exists, err = afero.Exists(client.Fs, client.credentialsFile)
	if !exists || err != nil {
		t.Fatalf("Expected credentials file to exist")
	}
}

func TestPersistentConfigSetCredentialsPermissions(t *testing.T) {
	client := ClientPersistentConfig{
		Fs: afero.NewMemMapFs(), // Using an in-memory file system
	}
	err := client.ensureConfigFiles()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = client.setCredentialsPermissions(client.credentialsFile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	info, err := client.Fs.Stat(client.credentialsFile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if info.Mode().Perm() != 0600 {
		t.Fatalf("Expected 0600 permissions, got %v", info.Mode().Perm())
	}
}

func TestPersistentConfigReadAndWriteConfigFile(t *testing.T) {
	client := ClientPersistentConfig{
		Fs: afero.NewMemMapFs(), // Using an in-memory file system
	}
	err := client.ensureConfigFiles()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	cfg := ini.Empty()
	section, err := cfg.NewSection(DefaultConfigProfile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	key, err := section.NewKey("key", "value")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = client.WriteConfigFile(cfg)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	readCfg, err := client.ReadConfigFile()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if readCfg.Section(DefaultConfigProfile).Key("key").String() != key.String() {
		t.Fatalf("Expected key to be %s, got %s", key, readCfg.Section("default").Key("key"))
	}
}
