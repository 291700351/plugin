package plugin_interface

type PluginRuntime struct {
	RuntimeDir string
}

func NewPluginRuntime(runtimeDir string) *PluginRuntime {
	return &PluginRuntime{RuntimeDir: runtimeDir}
}
