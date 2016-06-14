// +build bindata

package modules

import (
	"gopkg.in/macaron.v1"
	"github.com/ignatov/go-reactjs-example/public"
	"github.com/ignatov/go-reactjs-example/templates"
	"github.com/go-macaron/bindata"
)

var Public = macaron.Static("public",
	macaron.StaticOptions{
		Prefix: "-",
		FileSystem: bindata.Static(bindata.Options{
			Asset:      public.Asset,
			AssetDir:   public.AssetDir,
			AssetNames: public.AssetNames,
			AssetInfo:  public.AssetInfo,
			Prefix:     "",
		}),
	})

var Renderer = macaron.Renderer(macaron.RenderOptions{
	TemplateFileSystem: bindata.Templates(bindata.Options{
		Asset:      templates.Asset,
		AssetDir:   templates.AssetDir,
		AssetNames: templates.AssetNames,
		AssetInfo:  templates.AssetInfo,
		Prefix:     "",
	}),
})
