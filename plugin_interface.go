package plugin_interface

type PluginInterface interface {
	PackageName() string
	PluginName() string
	VersionCode() uint64
	VersionName() string

	OnCreate(runtime *PluginRuntime)
	OnProcess()
}
