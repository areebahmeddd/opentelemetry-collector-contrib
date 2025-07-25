// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				HostArch:           ResourceAttributeConfig{Enabled: true},
				HostCPUCacheL2Size: ResourceAttributeConfig{Enabled: true},
				HostCPUFamily:      ResourceAttributeConfig{Enabled: true},
				HostCPUModelID:     ResourceAttributeConfig{Enabled: true},
				HostCPUModelName:   ResourceAttributeConfig{Enabled: true},
				HostCPUStepping:    ResourceAttributeConfig{Enabled: true},
				HostCPUVendorID:    ResourceAttributeConfig{Enabled: true},
				HostID:             ResourceAttributeConfig{Enabled: true},
				HostInterface:      ResourceAttributeConfig{Enabled: true},
				HostIP:             ResourceAttributeConfig{Enabled: true},
				HostMac:            ResourceAttributeConfig{Enabled: true},
				HostName:           ResourceAttributeConfig{Enabled: true},
				OsBuildID:          ResourceAttributeConfig{Enabled: true},
				OsDescription:      ResourceAttributeConfig{Enabled: true},
				OsName:             ResourceAttributeConfig{Enabled: true},
				OsType:             ResourceAttributeConfig{Enabled: true},
				OsVersion:          ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				HostArch:           ResourceAttributeConfig{Enabled: false},
				HostCPUCacheL2Size: ResourceAttributeConfig{Enabled: false},
				HostCPUFamily:      ResourceAttributeConfig{Enabled: false},
				HostCPUModelID:     ResourceAttributeConfig{Enabled: false},
				HostCPUModelName:   ResourceAttributeConfig{Enabled: false},
				HostCPUStepping:    ResourceAttributeConfig{Enabled: false},
				HostCPUVendorID:    ResourceAttributeConfig{Enabled: false},
				HostID:             ResourceAttributeConfig{Enabled: false},
				HostInterface:      ResourceAttributeConfig{Enabled: false},
				HostIP:             ResourceAttributeConfig{Enabled: false},
				HostMac:            ResourceAttributeConfig{Enabled: false},
				HostName:           ResourceAttributeConfig{Enabled: false},
				OsBuildID:          ResourceAttributeConfig{Enabled: false},
				OsDescription:      ResourceAttributeConfig{Enabled: false},
				OsName:             ResourceAttributeConfig{Enabled: false},
				OsType:             ResourceAttributeConfig{Enabled: false},
				OsVersion:          ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{}))
			require.Emptyf(t, diff, "Config mismatch (-expected +actual):\n%s", diff)
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}
